# Feature Spec: F1 — Cài app qua OAuth

> Instance theo [`contracts/feature-spec.template.md`](../../../contracts/feature-spec.template.md).
> Sinh từ `μ-decompose` của [`../app-spec.md`](../app-spec.md) §6. **Acceptance Criteria
> chờ intake-feature** (lấp ở bước elaborate).

## 0. Metadata
- **ID:** `exp-001 / F1`
- **🔗 Epic cha:** [`../app-spec.md`](../app-spec.md) (story F1)
- **🔗 Design Doc:** [`../design-doc.md`](../design-doc.md)
- **Trạng thái:** Draft → **Ready** ✅ → Done
- **Ước lượng:** S — ~1–2 ngày (luồng OAuth chuẩn, ít logic)

## 1. User Story  🔒
**As a** merchant, **I want** cài app vào shop qua **OAuth**, **so that** shop được kết
nối và app có access token để thao tác.

## 2. Acceptance Criteria  🔒  ✅  *(Gherkin)*
```gherkin
Scenario 1 — Cài thành công
  Given shop CHƯA cài app
  When  merchant hoàn tất màn OAuth grant của Shopify
  Then  app lưu shop domain + access token
  And   chuyển merchant về trang admin của app

Scenario 2 — Từ chối quyền
  Given merchant đang ở màn OAuth grant
  When  merchant bấm Cancel
  Then  app KHÔNG lưu gì
  And   hiện thông báo "cài chưa hoàn tất"

Scenario 3 — Cài lại shop đã có
  Given shop ĐÃ cài app trước đó
  When  merchant mở lại link cài
  Then  app KHÔNG tạo bản ghi trùng
  And   cập nhật access token mới rồi vào admin

Scenario 4 — Callback giả mạo (bảo mật)
  Given request callback có HMAC/state KHÔNG hợp lệ
  When  app xử lý callback
  Then  app từ chối (4xx), KHÔNG lưu gì
```

## 3. Phạm vi  🔒
- **Trong:** luồng OAuth chuẩn; lưu `shop` + access token.
- **Ngoài:** billing · uninstall cleanup · connect Facebook.

## 4. Ràng buộc & quy ước
- Shopify OAuth chuẩn. *(chi tiết kỹ thuật → Design Doc)*

## 5. Phụ thuộc
- — (story gốc, không phụ thuộc story khác).

## 6/7. Ghi chú trao đổi · Rủi ro
- Oracle: `install · session · shop`.

---
## ✅ Definition of Ready — ĐẠT ✅  (US ✓ · AC Gherkin ✓ · Small ✓ · Estimable ✓ · Independent ✓ · Testable ✓)
## ✅ Definition of Done — ĐẠT ✅
- [x] Scenario 4 (HMAC/state verify) — unit test xanh.
- [x] **LIVE (2026-06-30):** cài thật trên `tuanhoangpc-2.myshopify.com` qua OAuth →
  HMAC+state verify pass với Shopify thật → token `shpua_…` lưu vào DB. **F1 nghiệm thu live.**
- [ ] Scenario 2 (cancel) — chưa cover (ngoài happy path đã chạy live).
