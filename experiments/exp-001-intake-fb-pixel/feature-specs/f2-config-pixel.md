# Feature Spec: F2 — Cấu hình Pixel ID

> Instance theo [`contracts/feature-spec.template.md`](../../../contracts/feature-spec.template.md).
> Sinh từ `μ-decompose` của [`../app-spec.md`](../app-spec.md) §6. **AC chờ intake-feature**.

## 0. Metadata
- **ID:** `exp-001 / F2`
- **🔗 Epic cha:** [`../app-spec.md`](../app-spec.md) (story F2)
- **🔗 Design Doc:** [`../design-doc.md`](../design-doc.md)
- **Trạng thái:** Draft → **Ready** ✅ → Done
- **Ước lượng:** S — ~1 ngày (1 form + validate + lưu)

## 1. User Story  🔒
**As a** merchant, **I want** nhập & lưu **1 Pixel ID** (đã validate định dạng),
**so that** pixel của tôi được cấu hình đúng.

## 2. Acceptance Criteria  🔒  ✅  *(Gherkin)*
```gherkin
Scenario 1 — Lưu Pixel ID hợp lệ
  Given merchant ở trang settings (shop đã cài app)
  When  nhập Pixel ID đúng định dạng (chuỗi số) và bấm Lưu
  Then  app lưu Pixel ID cho shop
  And   reload trang vẫn hiện đúng Pixel ID

Scenario 2 — Nhập sai định dạng (validate chặn)
  Given merchant ở trang settings
  When  nhập Pixel ID KHÔNG hợp lệ (có chữ / rỗng / quá ngắn) và bấm Lưu
  Then  app TỪ CHỐI lưu
  And   hiện lỗi validate rõ ràng

Scenario 3 — Cập nhật Pixel ID
  Given shop đã có 1 Pixel ID
  When  nhập Pixel ID mới hợp lệ và Lưu
  Then  app ghi đè Pixel ID cũ
```

## 3. Phạm vi  🔒
- **Trong:** trang admin nhập Pixel ID; validate **định dạng** (chuỗi số); lưu & reload còn.
- **Ngoài:** verify pixel với Facebook · multi-pixel · connect tài khoản FB.

## 4. Ràng buộc & quy ước
- Pixel ID nhập tay (string). *(chi tiết lưu trữ → Design Doc)*

## 5. Phụ thuộc
- **F1** (cần shop đã kết nối để gắn pixel config).

## 6/7. Ghi chú trao đổi · Rủi ro
- Oracle: `pixel · shopsetting`. R2: chỉ validate format, không verify FB.

---
## ✅ Definition of Ready — ĐẠT ✅  (US ✓ · AC Gherkin ✓ · Small ✓ · Estimable ✓ · dep F1 rõ · Testable ✓)
## ✅ Definition of Done *(khi build)*
- [ ] 3 scenario AC §2 pass.
- [ ] test/review/merge theo convention dự án.
