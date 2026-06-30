# Bậc 2 — Web Pixel fire trên storefront THẬT (🅐 backend nhận event)

Mục tiêu: deploy Web Pixel extension → vào storefront thật → backend nhận `page_viewed`.

> ⚠️ **R1 đã xác minh:** sandbox "strict" KHÔNG load được fbq client-side. Nên web pixel
> chỉ `fetch` event về backend `/collect`. "Tới Facebook Events Manager" cần CAPI (🅑, ngoài MVP).

## Đã chuẩn bị + ĐÃ VERIFY tự động (harness)
- Backend: route `POST /collect` (CORS) nhận event · scope `write_pixels,read_customer_events`
  · `webPixelCreate` gửi settings `{accountID, collectURL}` khi lưu Pixel ID. **25 test xanh.**
- Extension: `extensions/fb-web-pixel/` (toml + `src/index.js` + package.json) — `npm install` OK,
  `index.js` parse OK.
- `shopify.app.toml` đã viết sẵn (client_id điền sẵn).
- Tooling: node 22 + Shopify CLI 4.3.0 (qua nvm).
- ✅ **Đã GIẢ LẬP storefront bắn `page_viewed` → `/collect` → backend log đúng event.**
  Nghĩa là PHÍA BACKEND đã chắc chắn chạy; nếu live không thấy event → lỗi nằm ở Shopify
  (deploy/activate/storefront), KHÔNG phải backend.

> 💡 🅐 KHÔNG cần Facebook Pixel thật — dùng pixel ID giả `1234567890123456` là đủ chứng minh.

## A. Deploy extension (CLI — bạn chạy, cần login Shopify)

```bash
# shopify nằm dưới nvm node 22 — mở terminal mới là có (nvm default = 22). Kiểm tra:
shopify version          # 4.3.0

cd ~/GolandProjects/harness-lab/apps/fb-pixel-mvp
shopify auth logout 2>/dev/null; shopify app config link   # chọn app "fb-pixel-mvp" (client_id 1275b1e5…)
shopify app deploy                                          # deploy web pixel extension
```
> `config link` tạo `shopify.app.toml`. Nếu deploy báo lỗi format extension → chạy
> `shopify app generate extension --type web_pixel_extension`, rồi copy `src/index.js` của mình vào.

## B. Tunnel + sync URL + chạy backend

```bash
cd ~/GolandProjects/harness-lab/apps/fb-pixel-mvp

# 1) Mở tunnel + tự sync URL vào .env & shopify.app.toml (1 lệnh):
./sync-tunnel.sh                       # tự mở cloudflared, lấy URL, sync 2 file
#    (hoặc nếu đã có tunnel: ./sync-tunnel.sh https://<url>.trycloudflare.com)

# 2) Đẩy URL+scope+extension lên Partner app (KHỎI vào dashboard):
shopify app deploy

# 3) Chạy backend (đọc .env):
go run .
```
> `sync-tunnel.sh` lo việc đổi URL ở cả `.env` (APP_URL) và `shopify.app.toml`
> (application_url + redirect_urls). Mỗi lần tunnel đổi chỉ chạy lại bước 1 + `shopify app deploy`.

## C. Re-install (scope mới) + bật pixel

```bash
# 1) Re-install để token có scope write_pixels:
#    mở: https://<tunnel>/install?shop=tuanhoangpc-2.myshopify.com → Install (chấp nhận quyền mới)

# 2) Lưu Pixel ID → backend gọi webPixelCreate (kích hoạt web pixel cho shop):
curl -X POST https://<tunnel>/settings/pixel \
  -d "shop=tuanhoangpc-2.myshopify.com" -d "pixel_id=1234567890123456"
#    → "saved"  (và web pixel được activate với accountID + collectURL=<tunnel>/collect)
```

## D. Nghiệm thu LIVE

```bash
# Mở storefront thật:
https://tuanhoangpc-2.myshopify.com/
#   (dev store có thể hỏi storefront password — lấy ở Online Store → Preferences)
```
→ Storefront load → `page_viewed` → web pixel fetch về `/collect` →
**console backend in:** `PixelEvent received from storefront: {"event":"page_viewed",...}`
= **F3 (🅐) LIVE done** ✅

## Gỡ lỗi
- `shopify app deploy` lỗi format → dùng `generate extension` rồi copy index.js (xem A).
- Không thấy event ở /collect → web pixel chưa activate (chạy lại bước C2) hoặc storefront chưa load lại.
- Storefront password chặn → tắt ở Online Store → Preferences, hoặc nhập password.
- Tunnel đổi URL → cập nhật Partner app URLs + chạy lại C2 (collectURL theo APP_URL mới).
