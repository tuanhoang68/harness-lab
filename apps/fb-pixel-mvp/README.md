# fb-pixel-mvp — app build bằng harness (dogfood exp-001)

App **Facebook Pixel for Shopify (MVP)** — xây bằng meta-harness để làm bằng chứng.
Spec: [`../../experiments/exp-001-intake-fb-pixel/`](../../experiments/exp-001-intake-fb-pixel/)
(App Spec + Design Doc + Feature Specs F1/F2/F3).

Stack (theo Design Doc): Go · Gin · go-shopify/v4 · GORM+MySQL · OAuth state stateless.

## Trạng thái build
| Feature | Trạng thái |
|---------|-----------|
| **F1** Cài app (OAuth) | ✅ **LIVE** — cài thật trên `tuanhoangpc-2.myshopify.com`, token `shpua_…` lưu vào DB (HMAC+state verify pass với Shopify thật) · 15 unit test green |
| **F2** Cấu hình Pixel | 🟢 **wired & runs** — `pixel` (validate + repo upsert) + `web.SavePixel` · POST→saved/400 verify thật |
| **F3** Nhúng Pixel | ✅ **LIVE (🅐)** — extension deploy thật (fbpx-mvp-6) · webPixelCreate ok (web_pixel_id thật) · storefront `tuanhoangpc-2` fire `page_viewed` → backend log nhiều event |

**Tổng: 24 test green** (shopifyauth · shop · pixel · web · webpixel). Cả 3 feature wired; binary build & chạy.

## Chạy test
```bash
go test ./...        # 15 test green
```

## Chạy app (zero-setup, DB = SQLite)
```bash
go build -o fb-pixel-mvp . && PORT=8099 ./fb-pixel-mvp
# kiểm thử nhanh:
curl localhost:8099/healthz                              # ok
curl -i "localhost:8099/install?shop=demo.myshopify.com" # 302 → Shopify OAuth + state ký
curl -X POST localhost:8099/settings/pixel -d shop=demo.myshopify.com -d pixel_id=1234567890 # lưu Pixel ID
curl "localhost:8099/settings/pixel?shop=demo.myshopify.com"  # đọc Pixel ID đã lưu (200) / "chưa cấu hình" (404)
```
> Config qua env: `SHOPIFY_API_KEY` · `SHOPIFY_API_SECRET` · `SHOPIFY_SCOPES` · `APP_URL` · `PORT` · `DB_PATH`.
> Prod: đổi GORM driver SQLite → MySQL theo Design Doc. E2E OAuth thật (grant) cần Shopify
> app creds + public URL — làm khi có.
