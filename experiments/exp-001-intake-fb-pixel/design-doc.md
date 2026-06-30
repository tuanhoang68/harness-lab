# Technical Design Doc (TDD): Facebook Pixel for Shopify — MVP

> Instance theo [`contracts/design-doc.template.md`](../../contracts/design-doc.template.md).
> Phần **HOW** đi kèm [`app-spec.md`](./app-spec.md). Tech extract từ oracle
> `PremiUs/facebook-shopify-backend`. *(Chờ domain expert duyệt các quyết định ⚑.)*

## 0. Metadata
- **ID:** `exp-001 / app: fb-pixel-mvp (TDD)`
- **🔗 App Spec (PRD):** [`./app-spec.md`](./app-spec.md)
- **Trạng thái:** Draft → **Reviewed** ✅  (⚑ chốt: MySQL · bỏ Redis · OAuth state stateless)

## 1. Bối cảnh kỹ thuật
Backend Go phục vụ app **nhúng** (embedded) trong Shopify Admin: nhận OAuth, lưu
shop + Pixel ID, đăng ký Web Pixel để fire PageView. WHAT/WHY → xem App Spec.

## 2. Tech stack & lý do  *(ADR — extract từ oracle)*
| Lớp | Chọn | Vì sao | Phương án cân nhắc |
|-----|------|--------|--------------------|
| Ngôn ngữ/runtime | **Go 1.24** | khớp oracle, 1 binary, hiệu năng | Node/Remix |
| HTTP framework | **Gin** | oracle dùng, nhẹ, quen team | Echo · Fiber |
| Shopify SDK | **bold-commerce/go-shopify/v4** | oracle dùng; có OAuth + Admin API | tự gọi REST |
| Persistence | **GORM + MySQL** | oracle dùng (giữ để đối chiếu) | Postgres |
| OAuth state | **stateless** (signed `state` param) | bỏ Redis cho gọn MVP | Redis |

## 3. Kiến trúc tổng thể  *(C4 — Context + Container)*
```
  Merchant ──(embedded UI)──► [Backend Gin]
                                 │  ├─► MySQL  (shop · pixel_config)
                                 │  └─► Shopify Admin API (OAuth, đăng ký Web Pixel)
                                 │            via go-shopify
  Storefront(browser) ─► [Web Pixel sandbox của Shopify] ─► Facebook (event)
```
Mỗi module theo tầng (giống oracle): `controller → service → repository(port) → model`
(+ `di.go`, `route.go`). 3 module MVP: `install` · `pixelconfig` · `webpixel`.

## 4. Mô hình dữ liệu  *(ER — mức thực thể)*
| Thực thể | Vai trò | Quan hệ | Lưu |
|----------|---------|---------|-----|
| `Shop` | shop đã cài (domain, **access_token**, installed_at) | 1—1 `PixelConfig` | MySQL |
| `PixelConfig` | pixel_id của shop (+ web_pixel_id sau khi đăng ký) | thuộc `Shop` | MySQL |

> OAuth state **stateless** (signed `state` param) — không bảng Session, không Redis.

## 5. Các luồng chính  *(4+1 scenarios)*
### Flow A — Install (OAuth)  *(F1)*
`/install?shop=` → redirect Shopify grant → `/auth/callback` → **verify HMAC + state** →
đổi code→access_token → upsert `Shop` → redirect về admin.
### Flow B — Config Pixel  *(F2)*
admin `POST /settings/pixel {pixel_id}` → **validate định dạng** → upsert `PixelConfig`.
### Flow C — Inject + Event  *(F3)*
khi có pixel_id → gọi Shopify Admin API **tạo/cập nhật Web Pixel** (settings chứa pixel_id,
lưu `web_pixel_id`) → storefront tải Web Pixel → fire **PageView** → Facebook.

## 6. Tích hợp & giao diện ngoài
- **Shopify OAuth & Admin API** (go-shopify) — auth: OAuth access token.
- **Shopify Web Pixel Extension API** — đăng ký pixel (Flow C).
- **Facebook** — nhận event client-side qua pixel.
- **MySQL** (persistence). OAuth state stateless (signed `state`) — **không Redis** ở MVP.

## 7. Non-goals kỹ thuật & rủi ro
- **Non-goals:** MongoDB (oracle có, MVP không cần) · CAPI server-side · worker/queue ·
  multi-pixel.
- **Rủi ro:**
  - **R1 — ĐÃ XÁC MINH (2026-06-30):** sandbox Web Pixel "strict" **KHÔNG load được fbq
    client-side** → không nhúng Facebook Pixel kiểu cổ điển được. Cách hỗ trợ: extension
    chỉ `fetch` event server-side. ⇒ Bậc 2 (🅐): web pixel fetch `page_viewed` → backend
    `/collect`. Scope thật = `write_pixels,read_customer_events`. Settings extension =
    `{accountID, collectURL}` (backend gửi qua `webPixelCreate`). "Tới Events Manager"
    cần **CAPI server-side** (🅑) — ngoài MVP.
  - **Bảo mật:** verify HMAC + state ở callback là **bắt buộc** (khớp F1 Scenario 4).
