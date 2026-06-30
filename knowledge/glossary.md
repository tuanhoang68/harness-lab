# Từ điển & Phân loại Spec — NGUỒN CHÂN LÝ

> Khi bối rối "feature spec vs ready spec vs app spec là gì, khác nhau ra sao" → **đọc file này.**
> Đây là tài liệu định nghĩa chuẩn cho cả dự án.

## Nguyên tắc nền: LOẠI ≠ TRẠNG THÁI (2 trục độc lập)

Lẫn lộn hay xảy ra vì gộp hai thứ khác nhau làm một. Tách rõ:

```
   TRỤC 1 — LOẠI spec (theo độ cao / vai trò)
   ┌──────────────┬───────────────┬──────────────┐
   │  App Spec    │ Feature Spec  │  Design Doc   │
   │  Epic/PRD    │  User Story   │  TDD          │
   │  WHAT/WHY    │  WHAT/WHY     │  HOW          │
   └──────────────┴───────────────┴──────────────┘

   TRỤC 2 — TRẠNG THÁI spec (vòng đời)
   Draft ──(cổng DoR)──► READY ──(làm)──► In progress ──(cổng DoD)──► Done
                          ▲                                            ▲
              "Ready" = đã qua Definition of Ready        "Done" = đã qua Definition of Done

   ⇒ "Ready" và "Done" là TRẠNG THÁI, áp cho BẤT KỲ loại spec nào.
     KHÔNG có loại spec nào tên là "Ready Spec".
```

## Trục 1 — Ba LOẠI spec

| Loại | Là gì | Độ cao | Chuẩn ngành | Template |
|------|-------|--------|-------------|----------|
| **App Spec** | Epic/PRD: vấn đề, mục tiêu, scope, **phân rã ra feature** | App / dự án | PRD · SMART · MoSCoW · Working Backwards · Epic | `contracts/app-spec.template.md` |
| **Feature Spec** | User Story + acceptance criteria | Feature | Connextra · Gherkin · INVEST · 3C's | `contracts/feature-spec.template.md` |
| **Design Doc** | Kỹ thuật: tech stack, mô hình dữ liệu, các luồng | App (đi kèm App Spec) | Design Doc/RFC · C4 · 4+1 · ADR · ER | `contracts/design-doc.template.md` |

**Quan hệ:**
```
   App Spec  ──μ-decompose──►  nhiều Feature Spec
      ⇅ (đi kèm, giải phần HOW)
   Design Doc
```

## Trục 2 — TRẠNG THÁI spec & hai cổng

| Trạng thái | Nghĩa | Cổng để vào |
|------------|-------|-------------|
| Draft | đang soạn, chưa đủ | — |
| **Ready** | đủ rõ để bắt tay build | **Definition of Ready (DoR)** |
| In progress | đang build | — |
| **Done** | hoàn thành, nghiệm thu xong | **Definition of Done (DoD)** |

- **Definition of Ready (DoR):** checklist "đủ rõ để build chưa". Cả App Spec lẫn Feature
  Spec đều có DoR riêng (xem cuối mỗi template).
- **Definition of Done (DoD):** checklist "coi như xong chưa".

## ⚠️ Bẫy tên gọi đã từng gặp (ghi lại để khỏi lặp)

```
   ❌ "Ready Spec" KHÔNG phải loại spec thứ 4.
      = cách nói tắt cho "một spec (app/feature) đang ở trạng thái Ready".
   ❌ File template TỪNG tên `ready-spec.template.md` → khiến tưởng "ready spec" = "feature spec".
      ✅ Đã đổi thành `feature-spec.template.md` (2026-06-28).
   👉 Từ nay: "Ready"/"Done" = TRẠNG THÁI (trục 2). Tên LOẠI chỉ gồm:
      App Spec · Feature Spec · Design Doc.
```

## Flow — một spec chảy qua harness

```
                       đầu vào thô
                            │
                            ▼  μ-intake (router + extract/brainstorm)
             ┌──────────────┴───────────────┐
         độ cao = APP                    độ cao = FEATURE
             │                                │
             ▼                                ▼
     📄 App Spec  ⇄  📐 Design Doc        📄 Feature Spec
             │                                │
     ─ ─ Definition of Ready (DoR) ─ ─ ─ ─ ─ ─ ─ ─ ─   ← cổng "đủ rõ"
             │ ✅ READY                       │ ✅ READY
             ▼  μ-decompose (chỉ cấp app)      │
     nhiều 📄 Feature Spec ──(mỗi cái qua DoR)─┤
                                               ▼
                                μ-build → μ-verify → μ-review
                                               │
                                  ─ Definition of Done ─ ► ✅ DONE
```

## Thuật ngữ nhanh
- **Epic / User Story / Task:** phân cấp Agile. App Spec ≈ Epic; Feature Spec ≈ User Story.
- **μ-intake, μ-decompose, μ-build, μ-verify, μ-review:** các micro-harness (xem
  [`principles.md`](principles.md) nguyên lý 6).
- **Recipe:** công thức xâu chuỗi micro-harness — `/make`, `/make-app`, `/make-feature`.
- **Oracle:** bản tham chiếu để chấm "done" (vd repo PremiUs cho case FB Pixel).
