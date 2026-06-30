# harness-lab

Phòng thí nghiệm **harness engineering với Claude** — nghiên cứu & xây dựng mọi thứ
*bao quanh* Claude Code (recipe, contract, quy trình) để **tự động hóa việc code**:
nhanh nhất có thể, quy trình thống nhất, không phải prompt đơn điệu lặp đi lặp lại.

> Sản phẩm của lab **không phải một app cụ thể** — mà là *cỗ máy quy trình* (`/make`),
> bộ *chuẩn* (contracts) và *phương pháp luận* để đẻ ra app/feature đúng chuẩn, có test,
> mọi lần như một. Các app trong `apps/` chỉ là **phép thử (dogfood)** để kiểm chứng lab.

> 🧭 Lộ trình & vị trí hiện tại: [`PLAN.md`](PLAN.md) (bản đồ sống).

## Mô hình 3 tầng

```
┌───────────────────────────────────────────────────────┐
│  TẦNG 3 — META-HARNESS CỦA TA   ◀── lab này nghiên cứu │
│  quy trình thống nhất bọc quanh Claude Code            │
│   ┌───────────────────────────────────────────────┐   │
│   │  TẦNG 2 — CLAUDE CODE (Anthropic đã xây)       │   │
│   │   ┌─────────────────────────────────────────┐  │   │
│   │   │  TẦNG 1 — MODEL (Claude Opus)           │  │   │
│   │   └─────────────────────────────────────────┘  │   │
│   └───────────────────────────────────────────────┘   │
└───────────────────────────────────────────────────────┘
```

## Lab tạo ra cái gì (sản phẩm thật của dự án)

- **Họ recipe** (`meta/.claude/skills/`): `/make` · `/make-app` · `/make-feature` — gõ 1
  lệnh, harness tự dẫn `intake → [decompose] → build → verify`.
- **3 template spec chuẩn** ([`contracts/`](contracts/)): App Spec (PRD) · Feature Spec
  (User Story+Gherkin) · Design Doc (TDD) — ráp từ chuẩn ngành.
- **Kiến thức chung** ([`knowledge/`](knowledge/)): nguyên lý, từ điển spec (glossary), kiến trúc.

## Nguyên lý cốt lõi

> **Cố định contract, thay solver.** Kiến trúc không đổi; đổi bài toán = đổi *người giải*
> + *đường đi*. Micro-harness nối qua contract (toolkit, không monolith).

Chi tiết: [`knowledge/principles.md`](knowledge/principles.md) · [`knowledge/architecture.md`](knowledge/architecture.md).

## Bản đồ lab

```
        knowledge/  +  contracts/        ← lý thuyết + cổng chuẩn (dùng chung)
               │
               ▼
        ┌──────────────┐
        │    meta/     │   ★ TRỌNG TÂM: meta-harness trên Claude Code (recipe + solvers)
        └──────┬───────┘
               │ chạy quy trình meta để build
               ▼
        ┌──────────────┐
        │    apps/     │   ◀── ĐẦU RA dogfood: app build thật để CHỨNG MINH harness chạy
        └──────────────┘

        ┌──────────────┐
        │   engines/   │   💤 nhánh A (để dành): harness tự xây bằng Go
        └──────────────┘
```

| Thư mục | Vai trò |
|---------|---------|
| [`meta/`](meta/) | ★ Trọng tâm — meta-harness: recipe `/make*` + solvers |
| [`contracts/`](contracts/) | Template 3 loại spec chuẩn (cổng giữa các bước) |
| [`knowledge/`](knowledge/) | Lý thuyết · nguyên lý · từ điển spec ([glossary](knowledge/glossary.md)) · ADR |
| [`apps/`](apps/) | **Phép thử (dogfood)** — app build bằng lab để kiểm chứng harness |
| [`experiments/`](experiments/) | Nhật ký mỗi lần thử: input → output → nhận xét |
| [`engines/`](engines/) | 💤 Để dành — nhánh A: harness tự xây bằng Go |

## Phép thử đã chạy (kiểm chứng lab hoạt động)

- [`apps/fb-pixel-mvp`](apps/fb-pixel-mvp) — Shopify Facebook Pixel app, build hoàn toàn
  qua quy trình của lab: **28 test TDD**, 3 feature **chạy LIVE thật** trên store.
- **Test "harness tự lái":** feature F4 do **một agent khác (context trắng)** tự build chỉ
  bằng skill `/make-feature` — chứng minh recipe đủ để tái lập quy trình.
- *(Phép thử này cũng được dùng làm đồ án tốt nghiệp khóa học — ghi chú demo:
  [`SUBMISSION.md`](SUBMISSION.md). Đó là một **ứng dụng** của lab, không phải bản thân lab.)*

## Trạng thái

- [x] Nền móng: khung lab + lý thuyết + 3 template chuẩn + taxonomy
- [x] Họ recipe `/make · /make-app · /make-feature` codify (đã chứng minh tự lái qua F4)
- [x] Promote skill ra `~/.claude/skills` (symlink)
- [x] Dogfood `apps/fb-pixel-mvp` — 28 test, F1–F3 LIVE, F4 do agent build
- [ ] Codify stage `μ-design` + `/intake` riêng · hook gate (fmt/test)
