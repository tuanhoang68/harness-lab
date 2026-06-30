# claude-harness — Thiết kế (project học harness)

**Ngày:** 2026-06-18
**Mục tiêu học:** Hiểu một agent harness (như Claude Code) chạy thế nào bên trong, bằng cách **tự viết vòng lặp agent** trong Go.

## 1. Tổng quan

Một CLI Go tên `claude-harness`. Người dùng gõ yêu cầu; harness gọi Claude API; Claude
quyết định gọi tool; **harness tự thực thi tool** rồi đẩy kết quả về; lặp lại cho tới khi
Claude trả lời xong.

Bài học cốt lõi là **vòng lặp agent** và ranh giới trách nhiệm: Claude *phát ra ý định*
gọi tool (block `tool_use`); harness là phần *thực thi, kiểm soát, và nối kết quả* trở lại
hội thoại (block `tool_result`). Claude không tự chạy gì cả.

## 2. Lựa chọn kỹ thuật

- **Ngôn ngữ:** Go (Go 1.25), khớp workspace hiện tại.
- **Hướng:** Dùng Go SDK chính thức `github.com/anthropics/anthropic-sdk-go` để lo phần
  HTTP + marshal JSON, nhưng **tự viết vòng lặp thủ công** (không dùng `BetaToolRunner`) —
  vì cái loop chính là thứ đáng học.
- **Model:** `claude-opus-4-8` (mặc định của SDK qua `anthropic.ModelClaudeOpus4_8`).
- **API key:** đọc từ biến môi trường `ANTHROPIC_API_KEY` (SDK tự đọc). Không hardcode.

## 3. Kiến trúc & cấu trúc thư mục

```
claude-harness/
├── go.mod
├── main.go            # CLI/REPL: đọc API key, nhận input, gọi agent, in kết quả
├── agent/
│   └── agent.go       # Agent struct + vòng lặp Run() (phần lõi, dùng SDK)
├── tools/
│   ├── tool.go        # Interface Tool + Registry dispatch
│   ├── readfile.go    # tool read_file
│   └── bash.go        # tool bash (có cổng xác nhận)
└── README.md
```

Ba đơn vị tách bạch, mỗi đơn vị một trách nhiệm rõ ràng:

- **tools** — mỗi tool tự mô tả (tên, mô tả, JSON Schema input) và tự thực thi.
- **agent** — vòng lặp, quản lý lịch sử hội thoại, dịch tool sang định dạng SDK, dispatch.
- **main** — CLI/REPL và việc nối cổng xác nhận bash vào stdin.

### Quyết định thiết kế: tools KHÔNG phụ thuộc SDK

Interface `Tool` trả về JSON Schema dưới dạng `map[string]any` thuần, *không* dùng kiểu của
Anthropic SDK. Tầng `agent` mới là nơi dịch các mô tả này thành `anthropic.ToolParam` và
định tuyến các block `tool_use` ngược về. Lợi ích: đọc/test/suy luận một tool mà không cần
nghĩ tới giao thức wire; tầng tools build và test được offline (không cần SDK, không cần mạng).

## 4. Vòng lặp agent (phần đáng học nhất)

Trong `agent.Run(ctx, userInput)`:

1. Thêm tin nhắn user vào `messages`.
2. **Lặp:**
   - `resp := client.Messages.New(ctx, {Model, MaxTokens, Messages, Tools})`.
   - Append `resp.ToParam()` vào `messages` (giữ nguyên cả các block `tool_use`).
   - Duyệt `resp.Content`: in `TextBlock`; với mỗi `ToolUseBlock` →
     `registry.Dispatch(name, input)` → gom `anthropic.NewToolResultBlock(id, result, isError)`.
   - Nếu `resp.StopReason != anthropic.StopReasonToolUse` → **dừng**, trả về text cuối.
   - Ngược lại: append các `tool_result` thành một tin nhắn user, lặp tiếp.
3. `messages` được giữ qua các lượt REPL → hội thoại đa lượt.

Harness in rõ từng bước (Claude gọi tool gì, input ra sao, kết quả thế nào) để người học
*thấy* vòng lặp chạy.

## 5. Tool: interface & an toàn

```go
type Tool interface {
    Name() string
    Description() string
    InputSchema() map[string]any
    Execute(ctx context.Context, input json.RawMessage) (string, error)
}
```

- **`read_file`** — input `{path}`, trả nội dung file.
- **`bash`** — input `{command}`, chạy `bash -c`, trả stdout+stderr gộp. **Có cổng xác nhận:**
  trước khi chạy, harness gọi một `Confirmer` (mặc định hỏi y/N trên terminal). Cổng này được
  *tiêm vào* (injectable) nên test được và main nối vào stdin. Đây chính là bài học "tại sao
  nâng một hành động thành dedicated tool" — để harness có thể *chặn và kiểm soát* hành động
  khó đảo ngược, thay vì để model chạy tùy ý.

Lỗi tool được báo về model dưới dạng `tool_result` có `is_error = true` để model tự điều chỉnh,
thay vì làm sập vòng lặp.

## 6. Phạm vi

- **v1 (làm trước):** non-streaming; vòng lặp thủ công; 2 tool `read_file` + `bash`;
  REPL đa lượt; cổng xác nhận bash; in chi tiết để thấy loop.
- **Mở rộng (tùy chọn, sau):** streaming output; adaptive thinking (`display: "summarized"`);
  prompt caching; thêm tool (vd `write_file`).

## 7. Cấu hình, lỗi, test

- **Cấu hình:** `ANTHROPIC_API_KEY` từ env. Model là hằng trong code.
- **Test:** unit test cho tầng tools (`read_file` trên file tạm; `bash` với `echo` qua một
  `Confirmer` giả; dispatch tool không tồn tại trả về lỗi). Vòng lặp gọi API thật thì chạy
  thử end-to-end thủ công.

## 8. Kế hoạch theo checkpoint

1. **Checkpoint 1 (hiện tại):** spec + khung project + tầng tools (offline, có test).
2. **Checkpoint 2:** tầng `agent` — vòng lặp thủ công dùng SDK.
3. **Checkpoint 3:** REPL trong `main.go` nối tất cả lại, chạy end-to-end.
4. **(Tùy chọn) Checkpoint 4+:** các tính năng mở rộng ở mục 6.
