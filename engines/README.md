# engines/ — nhánh A: harness tự xây 💤 ĐỂ DÀNH

Các "động cơ" harness viết từ Claude API bằng Go — tự dựng vòng lặp agent từ con số 0
(tool layer, loop, context tự quản).

> **Trạng thái: để dành.** Chưa có kế hoạch triển khai trong tương lai gần. Trọng tâm
> hiện tại là `meta/`. Giữ thư mục này như kho tham khảo.

## Đã tập kết (✅ move xong)

| Project | Mô tả | Build |
|---------|-------|-------|
| `claude-harness`  | harness Go đầy đủ: `main.go` + `tools/` (bash, readfile) + `agent/` | ✅ |
| `minimal-harness` | bản tối giản | ✅ |

Đã move từ `GolandProjects/{claude-harness,minimal-harness}` vào đây. `go.mod` dùng
tên module độc lập nên đổi đường dẫn không vỡ build (đã verify lại sau move).

> ⚠️ GoLand có thể còn giữ **run config / recent locations** trỏ đường dẫn cũ — mở lại
> hoặc trỏ lại khi cần. Bản thân code và build không ảnh hưởng.

## Phân biệt với `meta/`

```
   engines/ (nhánh A)   = TỰ XÂY harness từ Claude API   → khó, từ con số 0
   meta/    (nhánh B)   = đứng trên vai Claude Code       → trọng tâm hiện tại
```
