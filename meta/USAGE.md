# USAGE — Cách dùng harness (flow-first)

> Cách "gõ gì → ra sản phẩm" cho 3 case: **1 feature · nhiều feature · app mới từ đầu**.
> Chi tiết recipe: [`meta/.claude/skills/`](.claude/skills/) · Loại spec: [`knowledge/glossary.md`](../knowledge/glossary.md)
> · Trạng thái dự án: [`PLAN.md`](../PLAN.md).

Nhớ 1 điều: **bạn nói *cái gì cần* + duyệt cổng; harness lo *làm thế nào*.** Luôn bắt đầu bằng `/make`.

---

## 0 · Cửa vào — router

```
        bạn gõ /make
             │
             ▼
      ┌──────────────┐
      │  độ cao?     │
      └──────┬───────┘
       ┌─────┴──────┐
       ▼            ▼
  app mới       feature vào
  chưa có code  app có sẵn
       │            │
       ▼            ▼
  /make-app    /make-feature
   (Case 3)     (Case 1/2)
```

---

## Xương sống chung (cả 3 case chảy vào đây)

```
 intake ─►( DoR? )─►[μ-decompose]─► build TDD ─►( DoD? )─► verify ─► DONE
  spec      ▲ │        chẻ feature   RED→GREEN    ▲ │      chạy thật
            │ │                                   │ │
         thiếu└─ quay lại hỏi                  test └─ đỏ→sửa code
       (chưa đủ rõ để code)                  (chưa pass hết)

  ●═══ 2 cổng chặn: DoR = đủ rõ để CODE?   DoD = mọi AC PASS?
  ●═══ mỗi ───► có checkpoint: harness DỪNG, chờ bạn "gật"
```

- **DoR** (Definition of Ready) = cổng VÀO: spec đủ rõ để bắt đầu code (User Story + Gherkin AC + INVEST).
- **DoD** (Definition of Done) = cổng RA: mọi Acceptance Criteria pass + test/vet xanh.

---

## Case 1 · 1 tính năng  (`/make-feature`)

```
 /make-feature "thêm X"
        │
        ▼
 [1] harness ĐỌC convention app ........................ (tự làm)
        │
        ▼
 [2] soạn nháp Feature Spec ─► hỏi điểm mơ hồ ─► 🙋 BẠN GẬT
        │                                         (gate DoR)
        ▼
 [3] mỗi AC:  viết test ─► RED ✗ ─► code ─► GREEN ✓ ─► 🙋 BẠN GẬT
        │
        ▼
 [4] go test/vet xanh + chạy binary thật
        │
        ▼
   spec: Draft ─► Ready ─► DONE ✅      ⟵ gật ~2–3 lần
```

---

## Case 2 · nhiều feature cùng lúc

```
       các feature LIÊN QUAN?
        ┌──────────┴───────────┐
       CÓ                      KHÔNG (rời rạc)
        │                        │
   μ-decompose              lặp /make-feature
        │                        │
        ▼                        ▼
   ┌─Fa─┬─Fb─┬─Fc─┐         Fx · Fy · Fz  (độc lập
   xếp THEO dependency       → chạy song song được)
        │
        ▼
   build lần lượt, mỗi feature 1 cổng DoD
        │
        ▼  THEO DÕI cái nào xong — KHÔNG đọc code:
   ┌─────────────────────────────────────────────┐
   │ PLAN.md       → Fa ✅  Fb 🚧  Fc ⏳  (mắt người) │
   │ Feature Spec  → field Draft/Ready/Done         │
   │ go test -run TestFx → xanh=xong đỏ=dở vắng=chưa │  ◄── CHÂN LÝ
   └─────────────────────────────────────────────┘
```

---

## Case 3 · app mới từ đầu  (`/make-app`)

```
 /make-app "dựng app X" / "clone oracle Y"
        │
        ▼
 B1 router: đầu vào RÕ (extract) hay MƠ HỒ (brainstorm)?
        │
        ▼
 B2─B7 intake APP SPEC (WHAT/WHY) ── mỗi mục: nháp ─► 🙋 GẬT ─► tick
   §1 bối cảnh ·§2 mục tiêu+metric ·§3 DoD ·§4 phạm vi ·§5 ràng buộc ·§6 epic→stories
        │
        ▼
   ( DoR: đủ 6 mục 🔒? ) ──► đẻ DESIGN DOC (HOW: stack/data/flow)
        │
        ▼
 μ-decompose: §6 ──► Fa·Fb·Fc (mỗi story 1 Feature Spec, Draft)
        │
        ▼
 per-feature: Gherkin AC + ước lượng ─► 🙋 GẬT ─► mỗi spec Ready
        │
        ▼
 ════ từ đây GIỐNG Case 2: build TDD từng feature ─► verify ─► APP CHẠY ════
```

---

## So 3 case trên cùng 1 trục

```
                intake          decompose      build      verify
 Case1 feature  ─ (đã có)  ───  (bỏ qua)  ───  TDD   ───  chạy   │ gật ~2–3
 Case2 nhiều    ─ (đã có)  ───  CHẺ Fx    ───  TDDx  ───  chạy   │ gật / feature
 Case3 app mới  ─ APP SPEC ───  CHẺ Fx    ───  TDDx  ───  chạy   │ gật nhiều
                  ▲đẻ mới        ▲+Design Doc
```

---

## Ai làm gì (mọi case)

```
   BẠN  ──► nói "cái gì cần"  +  duyệt cổng (gật/sửa §X)
                                        │
   HARNESS ◄────────────────────────────┘
        hỏi đúng thứ tự · soạn nháp spec · check DoR
        · chẻ feature · viết test trước · build · verify · cập nhật PLAN

   ⛔ chưa Ready → KHÔNG xuống build      ⛔ test chưa xanh → KHÔNG gọi Done
```

---

## Sample thật — F4 chạy qua Case 1  (2026-06-30)

> Feature: `GET /settings/pixel?shop=<domain>` (đọc Pixel ID, đối xứng POST đã có) vào app
> `apps/fb-pixel-mvp`. Spec: [`f4-get-pixel.md`](../experiments/exp-001-intake-fb-pixel/feature-specs/f4-get-pixel.md).

```
 INPUT bạn đưa:
   "thêm GET /settings/pixel trả Pixel ID đang lưu; chưa cấu hình thì báo;
    đối xứng POST; dùng template contracts/ + convention app"
        │
 [1] harness ĐỌC settings.go·pixel.go·web.go·main.go ─ rút convention:
        Gin handler · plain text · pixel.Repository.GetByShop (not-found=ErrRecordNotFound)
        │
 [2] hỏi 2 điểm mơ hồ (AskUserQuestion):
        format khi CÓ?  → 🙋 plain text
        khi CHƯA có?    → 🙋 404 "chưa cấu hình"
     ─► soạn Feature Spec (US + 3 Gherkin AC) ─► gate DoR ─► 🙋 "gật, build TDD"
        │
 [3] TDD 3 AC:
        viết TestGetPixel_ConfiguredReturnsID / _NotConfigured404 / _InvalidShop400
          └─ RED ✗  "h.GetPixel undefined"   (đúng lý do: feature thiếu)
        viết handler GetPixel (map 400/404/500/200)
          └─ GREEN ✓  3 PASS + full suite xanh
        nối route r.GET("/settings/pixel", h.GetPixel) vào main.go
        │
 [4] verify chạy binary thật (cổng 8100, né server LIVE 8099):
        chưa có → 404 "chưa cấu hình"   seed POST → 200 "saved"
        đã có   → 200 "9876543210987"   shop sai  → 400 "invalid shop"
        │
     ─► spec Draft→Ready→DONE ✅ · cập nhật PLAN.md + README · gật tất cả: 3 lần
```

OUTPUT cuối: `settings.go` (+handler) · `settings_test.go` (+3 test) · `main.go` (+route) ·
`f4-get-pixel.md` (Done) — `go test ./...` + `go vet` xanh, hành vi xác minh trên binary thật.
