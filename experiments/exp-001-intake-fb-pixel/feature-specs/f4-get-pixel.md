# Feature Spec: F4 — Đọc Pixel ID đang cấu hình

> Instance theo [`contracts/feature-spec.template.md`](../../../contracts/feature-spec.template.md).
> Đối xứng đọc của [`f2-config-pixel.md`](./f2-config-pixel.md) (ghi). **Đường "warm" /make-feature.**

## 0. Metadata
- **ID:** `exp-001 / F4`
- **🔗 Epic cha:** [`../app-spec.md`](../app-spec.md)
- **🔗 Design Doc:** [`../design-doc.md`](../design-doc.md)
- **Trạng thái:** Draft → Ready → **Done** ✅
- **Ước lượng:** XS — ~½ ngày (1 handler đọc + validate + 3 test)

## 1. User Story  🔒
**As a** merchant (hoặc UI settings), **I want** đọc lại **Pixel ID đang lưu** của shop mình
qua `GET /settings/pixel?shop=<domain>`, **so that** tôi xác nhận được pixel đã cấu hình đúng
(hoặc biết shop **chưa cấu hình**).

## 2. Acceptance Criteria  🔒  *(Gherkin)*
```gherkin
Scenario 1 — Shop đã cấu hình → trả Pixel ID
  Given shop "test-shop.myshopify.com" đã lưu Pixel ID "1234567890123456"
  When  GET /settings/pixel?shop=test-shop.myshopify.com
  Then  status 200
  And   body là đúng chuỗi "1234567890123456" (plain text)

Scenario 2 — Shop chưa cấu hình → báo chưa cấu hình
  Given shop "fresh-shop.myshopify.com" CHƯA lưu Pixel ID nào
  When  GET /settings/pixel?shop=fresh-shop.myshopify.com
  Then  status 404
  And   body chứa "chưa cấu hình"

Scenario 3 — Shop domain sai/thiếu → từ chối
  Given tham số shop rỗng hoặc không phải "<tên>.myshopify.com"
  When  GET /settings/pixel?shop=evil.com  (hoặc thiếu shop)
  Then  status 400
  And   body "invalid shop"
```

## 3. Phạm vi  🔒
- **Trong:** route `GET /settings/pixel`; lấy `shop` từ query; validate domain (reuse
  `validShopDomain`); đọc `pixels.GetByShop`; map kết quả → 200 / 404 / 400.
- **Ngoài:** xác thực shop qua Shopify session token (giữ nguyên giả định MVP như POST —
  shop lấy từ tham số client) · trả thêm field (web_pixel_id…) · UI embedded.

## 4. Ràng buộc & quy ước
- Gin handler, trả **plain text** (đồng nhất `SavePixel`). Không tin shop client (ghi chú
  bảo mật như `settings.go`). DB qua `pixel.Repository`; `GetByShop` not-found →
  `gorm.ErrRecordNotFound` ⇒ 404.

## 5. Phụ thuộc
- **F2** (cùng `pixel.Config` + `pixel.Repository.GetByShop`, đã có). Không thêm phụ thuộc mới.

## 6/7. Ghi chú trao đổi · Rủi ro
- Quyết định (user duyệt): found → 200 plain Pixel ID · chưa có → 404 "chưa cấu hình" ·
  shop sai → 400 "invalid shop". Oracle đọc: `pixel.GetByShop`.
- Rủi ro: lẫn lộn "chưa cấu hình" với lỗi DB khác → chỉ map `ErrRecordNotFound`→404, lỗi
  khác → 500.

---

## ✅ Definition of Ready
- [x] §1 User Story rõ giá trị.  [x] §2 ≥1 AC Gherkin testable (3 scenario, gồm nhánh lỗi).
- [x] §3 phạm vi trong/ngoài rõ.  [x] Small (XS).  [x] Estimable.  [x] Independent (chỉ dựa F2 đã có).
- [x] §0 link Epic cha.

## ✅ Definition of Done
- [x] 3 AC pass (offline httptest + sqlite — `TestGetPixel_*`).  [x] `go test ./...` + `go vet` xanh.  [x] route nối vào `main.go`.
- [x] Chạy thật binary (PORT=8100) chứng minh: AC2 `404 "chưa cấu hình"` · AC1 `200 "9876543210987"` · AC3 `400 "invalid shop"`.
