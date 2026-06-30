# PLAN — Lộ trình xây harness-lab  *(bản đồ sống, cập nhật liên tục)*

> Cập nhật: **2026-06-30**. Lý thuyết: [`knowledge/`](knowledge/) · Loại spec: [`glossary`](knowledge/glossary.md)
> · App dogfood: [`apps/fb-pixel-mvp/`](apps/fb-pixel-mvp/).

## 🎯 Đích cuối
**Meta-harness trên Claude Code**: gõ `/make` → harness phỏng vấn → chạy quy trình thống nhất
(intake → [decompose] → build → verify) → ra feature/app **nhanh, đúng chuẩn, mọi lần như một**.
Chứng minh bằng dogfood: build thật MVP Facebook Pixel (oracle `PremiUs/facebook-shopify-backend`).

## 📊 Trạng thái tổng quan

| Hạng mục | Trạng thái |
|----------|-----------|
| **HARNESS** — 3 template chuẩn (app/feature/design) + taxonomy | ✅ xong |
| **HARNESS** — họ recipe `/make` · `/make-app` · `/make-feature` (codify) | ✅ xong |
| **HARNESS** — test bộ skill (qua F4 dogfood) + promote ra `~/.claude/skills` | ✅ xong (symlink) |
| **DOGFOOD** — App Spec + Design Doc + 3 Feature Spec (Ready) | ✅ xong |
| **DOGFOOD** — code app (TDD) | ✅ **25 test xanh**, binary chạy |
| **DOGFOOD** — F1 cài app OAuth | ✅ **LIVE trên store thật** |
| **DOGFOOD** — F2 cấu hình Pixel | ✅ **LIVE** — webPixelCreate ok (web_pixel_id thật) |
| **DOGFOOD** — F3 web pixel (🅐) | ✅ **LIVE** — extension deploy + storefront fire → backend nhận event |
| **DOGFOOD** — F4 đọc Pixel ID (`GET /settings/pixel`) | ✅ **Done** — qua `/make-feature`, 3 test TDD + verify binary |

## 🧭 Nguyên tắc xuyên suốt
- **Tay trước, codify sau** · **Từng bước có checkpoint** · **Bằng chứng thực nghiệm** (`experiments/`).

## 🗺️ Các vòng

```
   ✅ V1 Nền móng        khung lab · lý thuyết · 3 template · taxonomy
   ✅ V2 Intake          B1–B7 phỏng vấn tay → App Spec Ready  (codify /intake riêng: tùy chọn)
   ✅ V3 Decompose       App Spec → F1/F2/F3, đều Ready (Gherkin AC)
   ✅ V4 Build           μ-build TDD cho F1/F2/F3 → 25 test xanh, binary chạy
   ✅ V5 Recipe          /make · /make-app · /make-feature codify  ( hook fmt/test: chưa )
   ✅ V6 Dogfood        F1 install LIVE · F2 webPixelCreate LIVE · F3 🅐 web pixel fire storefront LIVE
                        → MVP 3 feature CHẠY THẬT end-to-end trên store tuanhoangpc-2 🎉
```

## 📍 Đang ở đây
**🎉 MVP FB Pixel CHẠY THẬT end-to-end** trên store `tuanhoangpc-2.myshopify.com`:
- F1 OAuth install LIVE (token `shpua_…`) · F2+F3: `shopify app deploy` (fbpx-mvp-6) →
  re-install (scope `write_pixels`) → save pixel → `webPixelCreate` LIVE (web_pixel_id thật) →
  browse storefront → backend nhận `page_viewed` (trang chủ + sản phẩm).
- Tiện ích: `.env` (godotenv) + `sync-tunnel.sh` (sync URL vào .env + shopify.app.toml).
- Còn lại (tùy chọn): test+promote bộ skill `/make*` · 🅑 CAPI→Events Manager (ngoài MVP) ·
  F2 UI embedded · dọn (tắt backend/tunnel, gỡ web pixel/app khỏi store khi xong).

## ⏳ Việc còn lại
```
   CẦN USER (browser/CLI, không tự động được):
     P1 deploy extension (shopify auth login + app deploy)
     P2 re-install app (scope write_pixels)        P3 browse storefront → verify event
   HARNESS (tự làm tiếp được):
     • test bộ skill /make* trên case mới + promote ~/.claude/skills
     • codify stage μ-design + /intake + hook gate (fmt/test)
   NGOÀI MVP:
     • 🅑 CAPI → PageView tới Facebook Events Manager (cần FB access token)
     • F2 nghiệm thu live (UI embedded nhập Pixel ID)
```

## 🔎 Phát hiện kiến trúc (trong quá trình build)
- Pipeline cần thêm stage **μ-design** (đẻ Design Doc) giữa decompose ↔ build → sẽ codify.
- **μ-build = đọc Feature Spec AC → TDD từng AC verify-offline-được → external thì inject dependency.**
- **R1 xác minh:** web pixel sandbox "strict" không load fbq client-side → chỉ fetch server-side
  (DoD F3 điều chỉnh sang 🅐 trung thực).

## 🎮 Thao tác user khi harness đã codify (đích UX)
Thay ~14 lượt chat tay (như exp-001) bằng: `/make-app "<việc>"` → trả lời 2 câu router →
duyệt App Spec ("gật"/"sửa §X") → duyệt decompose. User chỉ **cung cấp sự thật domain + duyệt cổng**;
harness lo phần lặp lại (hỏi đúng thứ tự · soạn nháp · check DoR · chẻ feature · TDD).

## 📒 Nhật ký quyết định
- **2026-06-28:** trọng tâm `meta/`, `engines/` để dành · case dogfood = MVP FB Pixel (oracle PremiUs)
  · tách App Spec (PRD) ⇄ Design Doc (TDD) · trigger `/make` (Anh) · "Ready/Done"=trạng-thái.
- **2026-06-30:** Bậc 1 OAuth install LIVE (cloudflared) · R1 xác minh → F3 chọn 🅐 (không CAPI)
  · nâng node 22 qua nvm cho Shopify CLI 4.3 · backend Bậc 2 verify bằng giả lập storefront.
- **2026-06-30:** **F4** (`GET /settings/pixel`, đọc đối xứng F2) dogfood `/make-feature` thật:
  intake→spec Ready→TDD (RED 3 test→GREEN)→verify binary cổng 8100 (tránh `fbpx-live` 8099).
  Spec: [`feature-specs/f4-get-pixel.md`](experiments/exp-001-intake-fb-pixel/feature-specs/f4-get-pixel.md).
