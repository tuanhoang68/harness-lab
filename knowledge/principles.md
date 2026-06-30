# Nguyên lý harness engineering

## Nguyên lý trung tâm

> **Cố định các "khớp nối" (contract), thay người làm bên trong (solver).**
> Kiến trúc không đổi. Đổi bài toán = đổi *người giải* + *đường đi*, KHÔNG đổi xương sống.

Đây là *strategy pattern* / *ports-and-adapters* áp cho agent harness.

```
   ❌ Sai:  "đổi bài toán → xây lại pipeline"
   ✅ Đúng: "pipeline = KHUNG có các Ô trống + Cổng kiểm.
            Đổi bài toán → chỉ cắm 'người giải' khác vào Ô."

   ┌──────┐  ┌──────┐  ┌──────┐  ┌──────┐
   │ Ô 1  │─▶│ Ô 2  │─▶│ Ô 3  │─▶│ Ô 4  │
   └──┬───┘  └──┬───┘  └──┬───┘  └──┬───┘
   cắm ai?   cắm ai?   cắm ai?   cắm ai?
      ▼         ▼         ▼         ▼
  [triage]  [go-coder] [test-run] [reviewer]
  [brainstorm][react-coder] ...    [security]
```

## Cái gì bất biến, cái gì thay đổi

```
┌─────────────────────────────────────────────────────────┐
│  BẤT BIẾN (xương sống — thiết kế 1 lần):                 │
│    • Các Ô (stages)                                      │
│    • Các Cổng / artifact chuẩn giữa các ô (spec Ready…)  │
│    • Luật: cổng chưa đủ → không qua ô sau                │
│                                                          │
│  THAY ĐỔI tự do (cắm nóng — đổi mỗi bài toán):           │
│    • Solver nào cắm vào ô (Go/React, extract/brainstorm) │
│    • Đường đi (router chọn nhánh theo độ chín đầu vào)   │
└─────────────────────────────────────────────────────────┘
```

## Các nguyên lý phụ trợ

1. **Nhốt hỗn loạn ở cổng đầu.** Đầu vào đời thực hỗn loạn vô hạn → ép về một
   spec đạt Definition of Ready tại Ô 1. Mọi ô sau sống trong thế giới sạch.

2. **Chất lượng agent = chất lượng model × chất lượng harness.** Model ai cũng gọi
   được; lợi thế nằm ở harness. Đừng fine-tune khi mới bắt đầu — 90% giá trị ở
   context + tools + loop.

3. **Verification tách đồ chơi khỏi agent thật.** Agent code mà không tự chạy được
   test/build thì không đáng tin. Feedback loop (chạy → thấy lỗi → sửa) là linh hồn.

4. **Hook vs Skill vs CLAUDE.md** — ba lớp của quy trình thống nhất:
   - **Hook** = tự động (máy chạy, model không cần nhớ).
   - **Skill / slash command** = quy trình đóng gói (gõ 1 lệnh, chạy cả chuỗi).
   - **CLAUDE.md** = luật bất biến (khỏi nhắc lại mỗi lần).

5. **Đo, đừng tin.** Mỗi cải tiến harness phải có thực nghiệm trong `experiments/`
   chứng minh, không tuyên bố suông.

6. **Micro-harness, nối qua contract — toolkit chứ không monolith.**
   Triết lý Unix áp cho harness: *mỗi micro-harness làm MỘT việc, làm cho tốt; nối
   lại bằng đường ống.* Đường ống chính là **contract** (artifact chuẩn). Hai
   micro-harness ghép được **chỉ khi** output của cái này khớp input-contract của
   cái kia — đó là lý do contract phải được cố định trước.

   ```
   μ-decompose     μ-intake        μ-build       μ-verify      μ-review
   app-spec →    đầu vào thô →  feature-spec →   code →        code →
   [feature-    feature-spec     code           kết quả test  duyệt
    specs]
        └── output khớp input của viên kế (contract = mối nối) ──┘
   ```

   - **Mỗi viên** chỉ biết contract của mình, không biết viên trước/sau → thay/bỏ
     thoải mái.
   - **Nối là TÙY CHỌN** = chọn điểm vào dây chuyền. Build app từ đầu: vào từ
     `μ-decompose` rồi mỗi feature chạy chuỗi `intake→build→verify→review`. App đã
     ổn định: bỏ qua decompose, vào thẳng chuỗi feature.
   - Điều kiện để "vào thẳng" hoạt động: mỗi viên phải khởi động được từ một
     **artifact có thật trên đĩa** (vd feature-spec), bất kể do viên trước đẻ ra hay
     người viết tay. Không có contract sống thật → "nối tùy chọn" chỉ là lý thuyết.
   - Việc **xâu chuỗi** nằm ở các "công thức" (recipe = skill chỉ lo nối), tách khỏi
     bản thân các viên:
     - `/make` = cửa vào: hỏi "app hay feature?" rồi dispatch xuống đúng recipe.
     - `/make-app` = `decompose ▸ (mỗi feature: intake▸build▸verify▸review)`
     - `/make-feature` = `intake ▸ build ▸ verify ▸ review`
     - `/intake` = chạy đúng một viên.

   → Hệ quả: "tùy biến cho mọi trường hợp mà không đổi kiến trúc" = **không đổi viên
   Lego, chỉ đổi công thức xâu chuỗi.**

   *(Lưu ý đặt tên: A/B trong lúc thảo luận chỉ để dễ hình dung. Tên chính thức của
   các viên đặt theo chức năng — `μ-decompose`, `μ-intake`, `μ-build`, `μ-verify`,
   `μ-review` — KHÔNG dùng A/B để tránh đụng với `engines/` và `meta/`.)*
