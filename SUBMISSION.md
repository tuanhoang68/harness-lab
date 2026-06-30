# Ghi chú demo cuối khóa (đồ án tốt nghiệp)

> **Dự án** = `harness-lab` (cỗ máy quy trình tự động hóa việc code — xem [`README.md`](README.md)).
> **Đồ án/demo** = dùng lab này để build & chạy LIVE một **Shopify app** (`apps/fb-pixel-mvp`)
> làm **phép thử kiểm chứng** harness. App là *ứng dụng* của lab, không phải bản thân lab.

---

## 1. Mục tiêu
- **Dự án:** xây meta-harness trên Claude Code — gõ `/make` để tự dẫn quy trình
  intake → build → verify, ra feature/app đúng chuẩn, có test, mọi lần như một.
- **App demo (phép thử):** `fb-pixel-mvp` — cho merchant Shopify gắn 1 Facebook Pixel vào
  storefront thu event (PageView) trong vài phút, không sửa code theme.

## 2. Kiến thức khóa học đã áp dụng → bằng chứng

| Kiến thức | Bằng chứng trong repo |
|-----------|----------------------|
| Planning | [`PLAN.md`](PLAN.md) — bản đồ sống |
| CLAUDE.md / context cho AI | [`CLAUDE.md`](CLAUDE.md) |
| Harness workflow | `meta/.claude/skills/` — `/make · /make-app · /make-feature` |
| Prompt/context engineering | [`contracts/`](contracts/) + [`knowledge/glossary.md`](knowledge/glossary.md) |
| Delegation | **F4** — agent trắng context tự build feature qua skill |
| Verification | F1–F3 nghiệm thu **LIVE** trên store thật (curl/log/DB) |
| Testing | 28 test TDD (RED→GREEN) — `apps/fb-pixel-mvp` |
| Guardrails | cổng DoR/DoD · verify HMAC · chống open-redirect · dependency-injection |

## 3. AI hỗ trợ phần nào
Gần như **toàn bộ pipeline**: phỏng vấn lấp spec → Design Doc → decompose → code TDD →
deploy Shopify CLI → debug. Con người đóng vai **domain expert + duyệt cổng**; harness lo
phần lặp lại.

## 4. Kết quả đạt được
- Bộ harness tái dùng (`/make*` + 3 template chuẩn) — **đã chứng minh tự lái** qua F4.
- App Shopify 3 feature **chạy LIVE end-to-end** trên store `tuanhoangpc-2`: F1 OAuth install
  → token thật · F2 `webPixelCreate` thật · F3 storefront fire `page_viewed` → backend nhận.
- 28 unit test xanh · `go vet` sạch · binary chạy được.

## 5. Khó khăn & cách xử lý
- **R1:** Web Pixel sandbox "strict" KHÔNG load `fbq` client-side → **điều chỉnh DoD trung
  thực** sang fetch event về backend (🅐), thay vì giả vờ đạt.
- **Node cũ** không chạy Shopify CLI mới → cài node 22 qua **nvm** (không sudo).
- **Tunnel đổi URL** mỗi lần → viết `sync-tunnel.sh` tự đồng bộ `.env` + `shopify.app.toml`.

## 6. Bài học
- Harness = **bộ khung + kỷ luật** giúp việc-làm-tay ra chuẩn; rồi **codify** thành tự động.
- **Tay trước, codify sau** · **verify bằng bằng chứng thật**, không tin suông.
- Test một skill còn **tự mài sắc chính skill** (vòng RED→GREEN): phép thử F4 phát hiện skill
  hỏi trống → vá thành câu hỏi có option + lý do.

---

## Output kèm theo
- **Repo:** https://github.com/tuanhoang68/harness-lab
- **App demo:** [`apps/fb-pixel-mvp`](apps/fb-pixel-mvp) — chạy thử: `cd apps/fb-pixel-mvp && go test ./... && go run .`
- **Quy trình LIVE:** [`RUN-LIVE.md`](apps/fb-pixel-mvp/RUN-LIVE.md) · [`RUN-LIVE-BAC2.md`](apps/fb-pixel-mvp/RUN-LIVE-BAC2.md)
