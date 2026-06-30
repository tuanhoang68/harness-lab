# Bậc 1 — Cài app THẬT lên dev store (OAuth install live)

Mục tiêu: cài app lên một Shopify dev store thật qua OAuth → token được lưu.
Nghiệm thu DoD live của **F1**.

## A. BẠN làm (tài khoản & creds — chỉ con người làm được)

1. **Partner account:** đăng ký tại https://partners.shopify.com (miễn phí).
2. **Tạo app:** Partners → Apps → *Create app* → *Create app manually* → đặt tên.
   → Lấy **Client ID** (API key) + **Client secret** (API secret).
3. **Dev store:** Partners → Stores → *Add store* → *Development store*.
   → Ghi lại domain `<your-store>.myshopify.com`.
4. **Tunnel** (chọn 1, để Shopify gọi được vào localhost):
   - ngrok: `ngrok http 8099` → copy URL `https://xxxx.ngrok-free.app`
   - hoặc: `cloudflared tunnel --url http://localhost:8099`
5. **App settings** (Partners → app → *Configuration*):
   - **App URL:** `https://<tunnel>/`
   - **Allowed redirection URL(s):** `https://<tunnel>/auth/callback`
   - (tùy chọn) tắt *Embed app in Shopify admin* cho đơn giản.

## B. ĐÃ chuẩn bị (code — harness)

- Backend install-ready: scope hợp lệ (`read_products`), log khi cài xong.
- Routes: `/install` · `/auth/callback` · `/` · `/healthz`.
- HMAC + state verify đúng canonical của Shopify (24 unit test).

## C. Chạy & nghiệm thu

```bash
# 1) Tunnel — terminal 1
ngrok http 8099                      # copy https URL = <tunnel>

# 2) Server — terminal 2
cd apps/fb-pixel-mvp
SHOPIFY_API_KEY=<client_id> \
SHOPIFY_API_SECRET=<client_secret> \
APP_URL=https://<tunnel> \
PORT=8099 \
go run .

# 3) Cài app — mở trình duyệt
https://<tunnel>/install?shop=<your-store>.myshopify.com
#   → màn OAuth grant của Shopify → Install
#   → redirect về "/" hiện: fb-pixel-mvp running (shop=...)
```

**✅ Nghiệm thu F1 LIVE:** console server in
`OAuth install completed: shop=<your-store>.myshopify.com (access token saved)`.

## Gỡ lỗi thường gặp
- `redirect_uri is not whitelisted` → App URL / redirect chưa khớp `<tunnel>` (mục A5).
- `invalid_scope` → đặt `SHOPIFY_SCOPES=read_products`.
- Tunnel đổi URL mỗi lần chạy ngrok free → phải cập nhật lại A5 mỗi lần.
- Token lưu vào `fb-pixel.db` (SQLite) cạnh binary.
