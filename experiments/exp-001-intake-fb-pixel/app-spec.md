# App Spec (PRD): Facebook Pixel for Shopify — MVP

> Instance theo [`contracts/app-spec.template.md`](../../contracts/app-spec.template.md).
> Lấp dần qua phỏng vấn `/make` (B1→B7). 🔒 = bắt buộc cho Definition of Ready.

## 0. Metadata  ✅ *(khóa ở B1)*
- **ID:** `exp-001 / app: fb-pixel-mvp`
- **🔗 Design Doc (kỹ thuật):** [`./design-doc.md`](./design-doc.md)
- **Tác giả / Người duyệt:** domain expert (user) / ⬜ chờ duyệt
- **Trạng thái:** Draft → **Ready** ✅ → Approved *(chờ duyệt cuối)*
- **Oracle:** `PremiUs/facebook-shopify-backend` — modules `install · session · shop · pixel · shopsetting · webpixel`
- **B1 Router:** độ cao = **APP** (`/make-app`, có decompose) · độ chín = **RÕ** → solver **extract**

---

## 1. Bối cảnh & Vấn đề  🔒  ✅
Merchant Shopify muốn đo lường & tối ưu quảng cáo Facebook, nhưng gắn Facebook Pixel
vào storefront thủ công (sửa theme) thì khó, dễ sai, mất dữ liệu chuyển đổi. → Thiếu
một cách **đơn giản, đáng tin** để thu event hành vi người mua gửi về Facebook.

## 2. Mục tiêu & Chỉ số thành công  🔒  ✅
- **Mục tiêu:** merchant gắn 1 FB Pixel vào storefront trong vài phút, **không đụng code theme**.
- **Success metric (SMART):**
  - Cài + cấu hình xong **< 5 phút**.
  - Sau khi lưu Pixel ID, truy cập storefront → **≥1 event PageView** xuất hiện trong
    **Facebook Events Manager**.

## 3. Definition of Done (cấp app)  🔒  ✅
- [ ] Merchant cài app vào shop qua **OAuth** → app lưu được `shop` + access token.
- [ ] Merchant **nhập + lưu 1 Pixel ID** trong admin → reload trang vẫn còn.
- [ ] Truy cập storefront → **Web Pixel chạy** → **PageView** fire về FB đúng Pixel ID.
- [ ] **Kiểm chứng** được bằng Facebook Pixel Helper / Events Manager (thấy PageView).

## 4. Phạm vi — MoSCoW  🔒  ✅
- **Must:**
  - OAuth cài app + lưu shop/token.
  - Nhập & lưu 1 Pixel ID trong admin.
  - **Validate định dạng Pixel ID** (chuỗi số) — *(đẩy từ Could lên Must: nhập sai → pixel câm, khó debug)*.
  - Web Pixel nhúng → fire PageView lên storefront.
- **Should / Could:** hiện trạng thái "pixel đang hoạt động" trong admin.
- **Won't (ngoài MVP):** billing/plan/charge · OAuth-connect tài khoản Facebook ·
  multi-pixel · CAPI server-side · consent mode · catalog/webhook sync · dashboard ·
  voucher · uninstall cleanup · i18n đầy đủ · Sentry/apperror production-grade.

## 5. Ràng buộc · Giả định · Phụ thuộc  🔒  ✅
- **Ràng buộc:** nền Shopify — bắt buộc OAuth chuẩn + **Web Pixel Extension API** (không
  sửa theme tay); Pixel ID nhập tay dạng string (không kết nối tài khoản Facebook); MVP
  ưu tiên happy path (chưa ràng buộc Sentry/i18n). *(Tech stack cụ thể → Design Doc.)*
- **Giả định:** merchant đã có sẵn 1 Facebook Pixel ID; shop dùng theme/store hỗ trợ
  Web Pixel (Online Store 2.0 / checkout extensibility); merchant có quyền cài app.
- **Phụ thuộc:** Shopify OAuth & App API · Shopify Web Pixel Extension · Facebook (đầu
  nhận event) · datastore lưu shop + pixel config.

## 6. Phân rã: Epic → User Stories  🔒  ✅  *(đầu vào cho μ-decompose)*
| # | User Story | Giá trị | Phụ thuộc | Module oracle |
|----|-----------|---------|-----------|---------------|
| **F1** | As a merchant, I want cài app qua **OAuth**, so that shop được kết nối (app có token). | shop kết nối, có token thao tác | — | `install · session · shop` |
| **F2** | As a merchant, I want **nhập & lưu 1 Pixel ID** (đã validate), so that pixel được cấu hình đúng. | cấu hình pixel đúng, tránh nhập sai | F1 | `pixel · shopsetting` |
| **F3** | As a merchant, I want pixel **tự nhúng vào storefront**, so that PageView tự fire về Facebook (không sửa theme). | thu event tự động | F2 | `webpixel` |

Phụ thuộc tuyến tính: **F1 → F2 → F3**.

## 7. Rủi ro & Câu hỏi mở  ✅
- **R1** Web Pixel Extension — luồng đăng ký (Shopify CLI extension); cách truyền Pixel ID
  vào extension qua settings → xác minh khi build F3.
- **R2** Validate Pixel ID: MVP chỉ check **định dạng** (chuỗi số), KHÔNG verify với
  Facebook → chấp nhận rủi ro "đúng format nhưng pixel không tồn tại".
- **R3** Theme cũ không hỗ trợ Web Pixel → ngoài giả định; MVP không xử lý fallback.
- **R4** Event: MVP chỉ **PageView**; Purchase/AddToCart để vòng sau.

## 8. Tham chiếu
- Design Doc (HOW): [`./design-doc.md`](./design-doc.md)
- Oracle repo: `PremiUs/facebook-shopify-backend`

---

## ✅ Definition of Ready (GATE) — chưa đạt (đang lấp)
- [x] §1 Vấn đề rõ, không lẫn giải pháp
- [x] §2 ≥1 success metric đo được
- [x] §3 DoD có tiêu chí nghiệm thu được
- [x] §4 có Must và Won't
- [x] §5 nêu ràng buộc/giả định/phụ thuộc
- [x] §6 phân rã ≥1 story đạt INVEST

**→ DoR 6/6 ✅ — App Spec đạt Definition of Ready.**
