# Luật chung khi làm việc trong harness-lab

Đây là lab nghiên cứu harness engineering. Khi Claude làm việc trong thư mục này,
tuân theo các quy ước sau.

> **Bắt đầu phiên: đọc [`PLAN.md`](PLAN.md)** — bản đồ sống, biết đang ở đâu / đi đâu.
> Bối rối loại spec → [`knowledge/glossary.md`](knowledge/glossary.md).

## Ngôn ngữ
- Tài liệu, nhận xét, nhật ký: **tiếng Việt**.
- Tên thư mục/file, code, slug: **tiếng Anh** (kebab-case).

## Cách làm việc
- **Từng bước nhỏ, có checkpoint.** Làm một phần → dừng → nói rõ bước kế tiếp →
  chờ duyệt. Không gộp nhiều bước rủi ro vào một lần.
- **Bằng chứng thực nghiệm, không lý thuyết suông.** Mọi tuyên bố "harness tốt lên"
  phải có thực nghiệm trong `experiments/` chứng minh (input → output → nhận xét).
- **Trực quan hóa.** Ưu tiên giải thích bằng sơ đồ / flow ASCII bên cạnh chữ.
- **Hỏi có cấu trúc.** Mọi câu hỏi chốt (chọn hướng, lấp spec) dùng tool **AskUserQuestion**:
  2–4 option, mỗi option có `description` nêu trade-off; cái nên chọn lên đầu kèm "(Recommended)"
  và `description` **phải nói RÕ VÌ SAO** (không chỉ dán chữ); luôn cho nhập tự do. Không hỏi trống văn xuôi.

## Nguyên lý kiến trúc (bất biến)
- **Cố định contract, thay solver.** Đổi bài toán = đổi người giải + đường đi,
  KHÔNG đổi xương sống. Xem `knowledge/principles.md`.
- Mọi đầu vào phải được ép về **một spec chuẩn** đạt **Definition of Ready** trước khi
  xuống dây chuyền. Hỗn loạn bị nhốt ở bước intake.
- **Phân loại spec** (App Spec · Feature Spec · Design Doc; "Ready"/"Done" là trạng thái,
  không phải loại) — nguồn chân lý: `knowledge/glossary.md`. Templates: `contracts/`.

## Khi thêm cái mới
- Một **solver** mới → đặt trong `meta/solvers/<tên>/`, kèm README nói nó nhận gì /
  trả gì (phải khớp một contract trong `contracts/`).
- Một **quyết định thiết kế** đáng nhớ → ghi ADR trong `knowledge/decisions/`.
- Một **lần thử nghiệm** → một thư mục trong `experiments/`, ghi đủ input/output/nhận xét.

## Không làm
- Không đụng `engines/` trừ khi được yêu cầu (nhánh A đang để dành).
- Không tự ý move 2 project Go (`claude-harness`, `minimal-harness`) khi GoLand
  có thể đang index — hỏi trước.
