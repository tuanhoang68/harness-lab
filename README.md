# harness-lab — Đồ án tốt nghiệp

**Meta-harness trên Claude Code** (quy trình tự động hóa việc code) **+ một Shopify App
build bằng chính harness đó** và **chạy LIVE thật trên store**.

> Không chỉ là 1 Shopify app — mà là *cỗ máy quy trình* (`/make`) đẻ ra app đó, có test,
> đúng chuẩn, mọi lần như một. App `fb-pixel-mvp` là **bằng chứng** harness hoạt động.

---

## 🏆 Kết quả nổi bật

```
   • Shopify Facebook Pixel App — 3 feature CHẠY LIVE THẬT trên store tuanhoangpc-2:
       F1 cài app (OAuth)         → token thật lưu DB
       F2 cấu hình Pixel ID       → webPixelCreate thật (web_pixel_id gid://shopify/WebPixel/…)
       F3 nhúng Web Pixel         → storefront fire page_viewed → backend nhận event thật
   • 28 unit test (TDD RED→GREEN) · go vet sạch · binary chạy được
   • F4 (đọc Pixel ID) do MỘT AGENT KHÁC tự build qua /make-feature — chứng minh harness "tự lái"
   • Harness tái dùng: /make · /make-app · /make-feature (skill ở meta/.claude/skills)
```

---

## 🎤 Trình bày ngắn (6 mục theo đề bài)

**1. Mục tiêu của app**
App `fb-pixel-mvp`: cho merchant Shopify gắn 1 Facebook Pixel vào storefront để thu event
(PageView) **trong vài phút, không phải sửa code theme**.

**2. Kiến thức khóa học đã áp dụng** *(→ bằng chứng trong repo)*

| Kiến thức | Bằng chứng |
|-----------|-----------|
| Planning | [`PLAN.md`](PLAN.md) — bản đồ sống, cập nhật liên tục |
| CLAUDE.md / context cho AI | [`CLAUDE.md`](CLAUDE.md) — luật project cho Claude |
| Harness workflow | `meta/.claude/skills/` — recipe `/make · /make-app · /make-feature` |
| Prompt/context engineering | [`contracts/`](contracts/) (template chuẩn) + [`knowledge/glossary.md`](knowledge/glossary.md) |
| Delegation | **F4**: agent trắng context tự build feature qua skill ([f4](experiments/exp-001-intake-fb-pixel/feature-specs/f4-get-pixel.md)) |
| Verification | F1–F3 nghiệm thu **LIVE** trên store thật (curl / log / DB) |
| Testing | 28 test TDD ([apps/fb-pixel-mvp](apps/fb-pixel-mvp)) — viết test trước (RED) → code (GREEN) |
| Guardrails | Cổng DoR/DoD · verify HMAC · chống open-redirect · dependency-injection để test offline |

**3. AI đã hỗ trợ phần nào**
Gần như **toàn bộ pipeline**: phỏng vấn lấp spec → Design Doc → decompose feature → viết code
theo TDD → deploy Shopify CLI → debug. Con người (mình) đóng vai **domain expert + duyệt cổng**;
harness lo phần lặp lại.

**4. Kết quả đạt được**
App 3 feature chạy LIVE end-to-end trên store thật + 1 feature do agent tự build + một bộ
harness tái dùng được + tài liệu quy trình đầy đủ.

**5. Khó khăn & cách xử lý**
- **R1**: Web Pixel sandbox "strict" KHÔNG load được `fbq` client-side → **điều chỉnh DoD
  trung thực** sang phương án fetch event về backend (🅐), thay vì giả vờ.
- **Node cũ** không chạy Shopify CLI mới → cài node 22 qua **nvm** (không sudo).
- **Tunnel đổi URL** mỗi lần → viết `sync-tunnel.sh` đồng bộ tự động.

**6. Bài học**
- Harness = **bộ khung + kỷ luật** giúp việc-làm-tay ra chuẩn; rồi **codify** thành tự động.
- **Tay trước, codify sau** · **verify bằng bằng chứng thật**, không tin suông.
- Test một skill còn **tự mài sắc chính skill** (vòng RED→GREEN: phát hiện hỏi trống → ép
  dùng câu hỏi có option + lý do).

---

## 🗺️ Mô hình & bản đồ repo

```
   TẦNG 3 META-HARNESS CỦA TA   ← repo này
     └ bọc quanh → TẦNG 2 CLAUDE CODE → TẦNG 1 MODEL
```

| Thư mục | Vai trò |
|---------|---------|
| [`knowledge/`](knowledge/) | Lý thuyết · nguyên lý · từ điển spec (glossary) |
| [`contracts/`](contracts/) | Template 3 loại spec: App Spec · Feature Spec · Design Doc |
| [`meta/`](meta/) | ★ Harness: skills `/make·/make-app·/make-feature` |
| [`apps/fb-pixel-mvp/`](apps/fb-pixel-mvp/) | **Shopify app** build bằng harness (chạy LIVE) |
| [`experiments/`](experiments/) | Nhật ký build app (App Spec · Design Doc · Feature Specs F1–F4) |
| [`engines/`](engines/) | 💤 Nhánh A để dành: harness tự xây bằng Go |

## ▶️ Chạy thử app (zero-setup)
```bash
cd apps/fb-pixel-mvp
cp .env.example .env      # điền creds nếu chạy OAuth thật
go test ./...             # 28 test xanh
go run .                  # server :8099 (đọc .env)
```
Chạy LIVE trên Shopify thật: xem [`apps/fb-pixel-mvp/RUN-LIVE.md`](apps/fb-pixel-mvp/RUN-LIVE.md)
và [`RUN-LIVE-BAC2.md`](apps/fb-pixel-mvp/RUN-LIVE-BAC2.md).

## Nguyên lý cốt lõi
> **Cố định contract, thay solver.** Kiến trúc không đổi; đổi bài toán = đổi người giải.
> Chi tiết: [`knowledge/principles.md`](knowledge/principles.md) · [`architecture.md`](knowledge/architecture.md).
