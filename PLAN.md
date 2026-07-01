# PLAN — Lộ trình PHÁT TRIỂN HARNESS  *(harness-lab, bản đồ sống)*

> Đây là kế hoạch phát triển **bản thân harness** (cỗ máy quy trình `/make`). Các app trong
> [`apps/`](apps/) chỉ là **phép thử kiểm chứng** harness — KHÔNG phải nội dung của plan này.
> Cập nhật: **2026-06-30** · Repo: https://github.com/tuanhoang68/harness-lab

## 🎯 Đích cuối
Meta-harness trên Claude Code giúp dev/team code **tự động hơn · đáng tin hơn · nhất quán hơn**:
gõ `/make` → ra feature/app đúng chuẩn, có test, **ít cần người đứng ở vòng lặp lặp lại**.

## 🧭 Nguyên tắc xuyên suốt
Tay trước–codify sau · Từng bước có checkpoint · Bằng chứng thực nghiệm · Hỏi có cấu trúc (option+lý do).

## 🪜 Thang trưởng thành (xương sống)

```
   L1 ✅ METHODOLOGY    quy trình + template + recipe /make*   (gate THỦ CÔNG, verify TAY)  ◀ ĐANG Ở ĐÂY
   L2 ▢  AUTOMATION     hooks tự test/verify · recipe tự chạy chuỗi · μ-design/verify/review
   L3 ▢  RELIABILITY    adversarial verify · DoD auto-check · regression net · guardrails
   L4 ▢  TEAM SCALE     PR/CI · shared conventions · parallel fan-out · onboarding
   L5 ▢  SELF-IMPROVING đo metrics · vòng RED→GREEN tự mài skill · harness học từ dự án
```
> Hiện harness đã chứng minh **quy trình đúng** (L1) nhưng *người vẫn đứng mọi cổng* + *verify do người chạy tay*.
> Hành trình tới = đẩy người ra khỏi vòng lặp (L2) → thêm lưới an toàn (L3) → nhân cho team (L4) → tự cải tiến (L5).

## ✅ ĐÃ LÀM — L1 nền móng
- **Họ recipe** `/make · /make-app · /make-feature` — codify, **promote `~/.claude/skills`** (gõ mọi phiên).
- **3 template chuẩn** (App/Feature/Design) + **taxonomy/glossary** + cổng **DoR/DoD**.
- Kỷ luật **μ-build = TDD-từ-AC** · dependency-injection để test offline.
- Intake **hỏi có cấu trúc** (AskUserQuestion + option + lý do).

## 🧪 ĐÃ KIỂM CHỨNG (dogfood — bằng chứng L1 chạy thật)
- [`apps/fb-pixel-mvp`](apps/fb-pixel-mvp): Shopify FB Pixel app build hoàn toàn qua harness —
  **28 test TDD**, F1–F3 **chạy LIVE** trên store thật.
- **Harness "tự lái":** F4 do **agent trắng context** tự build chỉ bằng `/make-feature`.
- Chi tiết: [`experiments/exp-001-…`](experiments/exp-001-intake-fb-pixel/) · [`SUBMISSION.md`](SUBMISSION.md).

## 🔄 KẾ TIẾP NGAY (nợ đã lộ ra trong lúc build)
- [ ] **μ-design** codify — pipeline thiếu stage "đẻ Design Doc" (đang làm tay).
- [ ] **/intake** riêng — chạy 1 viên, router suy ngữ cảnh.
- [ ] Nợ nhỏ: F2 verify live (UI embedded) · 🅑 CAPI nếu cần tới Events Manager.

## 🚀 ROADMAP — 3 trục

### Trục 1 · TỰ ĐỘNG HƠN (L2) — đẩy người ra khỏi vòng lặp
- [ ] **① HOOKS** *(đòn bẩy lớn nhất, rẻ nhất)* — post-Edit `.go` → tự `fmt+vet+test`; chặn "done" khi test đỏ.
- [ ] **② μ-verify / μ-review codify** — build xong tự chạy test + subagent review → nổi findings.
- [ ] **③ Recipe tự chạy chuỗi** — `/make-app` loop thật: decompose → build từng feature, ít gate tay.

### Trục 2 · ĐÁNG TIN HƠN (L3) — lưới an toàn
- [ ] **④ Adversarial verify** — subagent phản biện cố BÁC BỎ mỗi AC/finding → giảm sai.
- [ ] **⑤ DoD auto-check** — máy kiểm mọi Acceptance Criteria pass mới cho đóng "Done".
- [ ] **⑥ Regression net + guardrails** — tích lũy test chạy mọi thay đổi; permission cho thao tác nguy hiểm.

### Trục 3 · GIÚP TEAM DEV (L4) — 1 người → cả team
- [ ] **⑦ PR/CI integration** — recipe tự mở PR kèm Feature Spec + test + tóm tắt; DoD = CI gate.
- [ ] **⑧ Shared conventions** — CLAUDE.md/AGENTS.md cấp team; harness ép chuẩn chung.
- [ ] **⑨ Parallel fan-out** — nhiều agent build nhiều feature song song.
- [ ] **⑩ Onboarding tức thì** — dev/agent mới đọc glossary+PLAN là chạy (đã chứng minh qua F4).

### L5 · TỰ CẢI TIẾN
- [ ] **⑪ Đo metrics** — cycle-time · coverage · tỷ lệ rework → chứng minh harness tăng hiệu suất bao nhiêu.
- [ ] **⑫ Vòng RED→GREEN tự động** — test skill bằng subagent định kỳ → tự lộ lỗ hổng → tự vá.

## 🎯 Thứ tự đề xuất (đòn bẩy cao trước)
```
   1) HOOKS ①          rẻ · tự động + đáng tin ngay · dùng mỗi ngày
   2) μ-verify+review ② khép vòng build→verify tự động
   3) PR/CI ⑦          mở đường lên team
   4) Đo metrics ⑪      có số liệu chứng minh giá trị
   … rồi adversarial / parallel / self-improving khi cần độ tin & quy mô cao hơn
```

## 🔎 Phát hiện kiến trúc (rút ra khi build)
- Pipeline cần stage **μ-design** (đẻ Design Doc) giữa decompose ↔ build.
- **μ-build = đọc Feature Spec AC → TDD từng AC verify-offline → external thì inject dependency.**
- Recipe nên đặt tên theo **chức năng** (μ-*), "Ready/Done" = **trạng thái** không phải loại spec.
- **R3 (2026-06-30):** app fb-pixel-mvp **chỉ có `webPixelCreate`, THIẾU `webPixelUpdate`**. Shopify
  cho mỗi app 1 web pixel/shop → khi tunnel đổi (deploy mới), pixel cũ trỏ collectURL CHẾT, `Create`
  báo *"settings already been set"*, app không tự repoint → storefront fire vào URL chết, backend 0 event.
  Lộ khi re-verify live: phải `webPixelUpdate` thủ công bằng GraphQL mới thông. **Cần feature upsert Web
  Pixel** → ứng viên `/make-feature` kế tiếp.

## 🎮 Đích UX (khi harness codify đủ)
`/make "<việc>"` → trả lời ~2 câu router → duyệt spec ("gật"/"sửa §X") → duyệt decompose.
User chỉ **cung cấp sự thật domain + duyệt cổng**; harness lo phần lặp lại.

## 📒 Nhật ký quyết định
- **2026-06-28:** trọng tâm `meta/`; tách App Spec (PRD) ⇄ Design Doc (TDD); trigger `/make`; "Ready/Done"=trạng-thái.
- **2026-06-30 (kiểm chứng):** dogfood FB Pixel — F1–F3 LIVE · R1 xác minh (sandbox không fbq → 🅐) · F4 agent tự build.
- **2026-06-30 (identity):** **KHÔNG đánh tráo** — dự án = harness-lab; app/đồ-án chỉ là phép thử. README mô tả lab, demo để `SUBMISSION.md`.
- **2026-06-30 (roadmap):** chốt thang trưởng thành L1→L5 + 3 trục; đòn bẩy kế tiếp = **Hooks**.
- **2026-06-30 (re-verify live):** deploy tunnel mới + OAuth install lại → F1–F3 LIVE end-to-end
  (token thật · `webPixelUpdate` repoint · storefront `page_viewed` về `/collect`). Lộ R3 (thiếu webPixelUpdate).
