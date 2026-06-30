# exp-001 — Cổng intake (DoR) với case FB Pixel MVP

## Mục tiêu thí nghiệm
Chứng minh viên `μ-intake` + cổng **Definition of Ready** hoạt động: nhiều kiểu đầu vào
→ cùng ra một spec đạt DoR, đối chiếu được với app thật.

## Case
Clone MVP app **Facebook Pixel for Shopify** (oracle: `PremiUs/facebook-shopify-backend`).
Phạm vi: 3 feature — F1 cài app (OAuth) · F2 cấu hình Pixel · F3 nhúng Pixel.

## Phỏng vấn /make — lấp App Spec từng bước (tay trước, codify sau)
- [x] **B1** Router → `/make-app` + solver `extract` (khóa ở [`app-spec.md`](app-spec.md) §0)
- [x] **B2** Bối cảnh & Mục tiêu + success metric (PRD §1–2)
- [x] **B3** Definition of Done (§3)
- [x] **B4** Phạm vi MoSCoW (§4) — validate Pixel ID đẩy lên Must
- [x] **B5** Ràng buộc/Giả định/Phụ thuộc (§5)
- [x] **B6** Phân rã Epic→Stories (§6)
- [x] **B7** Rủi ro & câu hỏi mở (§7) → **DoR 6/6 → App Spec Ready** ✅
- [x] (kèm) Design Doc HOW → [`design-doc.md`](design-doc.md) (stack/kiến trúc/ER/flows từ oracle; chờ duyệt ⚑)

## Vòng 3 — μ-decompose (App Spec → Feature Specs)
- [x] Chẻ §6 thành 3 skeleton: [`F1`](feature-specs/f1-oauth-install.md) ·
  [`F2`](feature-specs/f2-config-pixel.md) · [`F3`](feature-specs/f3-inject-pixel.md)
- [x] Elaborate F1·F2·F3: lấp Gherkin AC + ước lượng → cả 3 đạt DoR (Ready)
- [ ] Chạy 2 đầu vào trong `inputs/` (good-doc / vague) → đối chiếu (chứng minh cổng)
- [ ] Codify B1–B7 + decompose thành skill `/intake`, `/make`, `/make-app`

## Kết quả & nhận xét
*(điền sau khi chạy)*

## Phát hiện về contract
- App Spec cần mục **"Phân rã (app→feature)"** mà Feature Spec không có → đã tách 3
  template riêng (`app-spec` / `feature-spec` / `design-doc`) trong `contracts/`. ✅
