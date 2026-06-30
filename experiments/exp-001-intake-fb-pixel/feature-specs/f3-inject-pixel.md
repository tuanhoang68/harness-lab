# Feature Spec: F3 — Nhúng Pixel vào storefront

> Instance theo [`contracts/feature-spec.template.md`](../../../contracts/feature-spec.template.md).
> Sinh từ `μ-decompose` của [`../app-spec.md`](../app-spec.md) §6. **AC chờ intake-feature**.

## 0. Metadata
- **ID:** `exp-001 / F3`
- **🔗 Epic cha:** [`../app-spec.md`](../app-spec.md) (story F3)
- **🔗 Design Doc:** [`../design-doc.md`](../design-doc.md)
- **Trạng thái:** Draft → **Ready** ✅ → Done
- **Ước lượng:** M — ~2–3 ngày (R1: luồng Web Pixel Extension chưa rõ, có thể cần spike)

## 1. User Story  🔒
**As a** merchant, **I want** pixel **tự nhúng vào storefront**, **so that** event
PageView tự fire về Facebook mà không cần sửa code theme.

## 2. Acceptance Criteria  🔒  ✅  *(Gherkin)*
```gherkin
Scenario 1 — Pixel fire PageView
  Given shop đã lưu Pixel ID (F2 xong)
  When  khách truy cập một trang storefront
  Then  Web Pixel chạy và gửi event PageView về Facebook với đúng Pixel ID

Scenario 2 — Kiểm chứng được (khớp metric §2 App Spec)
  Given pixel đã nhúng
  When  kiểm tra Facebook Events Manager / Pixel Helper
  Then  thấy event PageView từ đúng Pixel ID

Scenario 3 — Chưa cấu hình pixel thì không fire
  Given shop CHƯA lưu Pixel ID
  When  khách truy cập storefront
  Then  KHÔNG nhúng pixel rỗng, KHÔNG fire event
```

## 3. Phạm vi  🔒
- **Trong:** đăng ký Web Pixel Extension; fire **PageView** với đúng Pixel ID.
- **Ngoài:** event Purchase/AddToCart · CAPI server-side · consent mode.

## 4. Ràng buộc & quy ước
- Shopify **Web Pixel Extension API**. *(luồng đăng ký → Design Doc + R1)*

## 5. Phụ thuộc
- **F2** (cần có Pixel ID đã lưu để truyền vào extension).

## 6/7. Ghi chú trao đổi · Rủi ro
- Oracle: `webpixel`. R1: cần xác minh luồng đăng ký extension + truyền Pixel ID.

---
## ✅ Definition of Ready — ĐẠT ✅  (US ✓ · AC Gherkin ✓ · Estimable ✓ · dep F2 rõ · Testable ✓)

> ⚠️ **R1 xác minh (2026-06-30) → DoD điều chỉnh trung thực:** sandbox strict không load fbq
> client-side. AC gốc "PageView trong Events Manager" KHÔNG khả thi client-side (cần CAPI =
> 🅑, ngoài MVP). Bậc 2 chốt **🅐**: web pixel fetch `page_viewed` → backend `/collect`.

## ✅ Definition of Done — phiên bản 🅐 (MVP-scope) — ĐẠT ✅
- [x] Logic build TDD: `webpixel.Service` activate/skip · backend `/collect` nhận event (25 test).
- [x] **LIVE (2026-06-30):** deploy extension (fbpx-mvp-6) → `webPixelCreate` thật
  (web_pixel_id `gid://shopify/WebPixel/2348187885`) → browse storefront `tuanhoangpc-2`
  (trang chủ + trang sản phẩm) → backend log `PixelEvent received` nhiều lần. **F3 🅐 nghiệm thu live.**
- [x] Scenario 3 (chưa cấu hình → không fire) — verified: trước khi activate, browse KHÔNG có event.
- [ ] (🅑, ngoài MVP) PageView tới Facebook Events Manager — cần CAPI.
