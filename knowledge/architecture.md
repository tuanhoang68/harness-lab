# Kiến trúc harness — xương sống

Tài liệu này mô tả **phần BẤT BIẾN** của harness: các Ô (stages) và các Cổng (gates).
Phần thay đổi (solver, đường đi) nằm ở `meta/` và được mô tả ở
[`principles.md`](principles.md).

## 1. Xương sống: 4 Ô + 3 Cổng

```
┌──────────┐   ┌──────────┐   ┌──────────┐   ┌──────────┐
│  Ô 1     │──▶│  Ô 2     │──▶│  Ô 3     │──▶│  Ô 4     │
│ INTAKE   │ ▣ │ BUILD    │ ▣ │ VERIFY   │ ▣ │ REVIEW   │
└──────────┘   └──────────┘   └──────────┘   └──────────┘

▣ = CỔNG (gate): artifact chuẩn phải có thì cổng mới mở
```

| Ô | Tên | Nhận | Trả ra | Cổng ra |
|---|-----|------|--------|---------|
| 1 | INTAKE | đầu vào thô (doc / 1 dòng / rỗng) | spec đã chuẩn hóa | **DoR** ✅ (spec → Ready) |
| 2 | BUILD  | spec đã Ready | code thay đổi | build pass |
| 3 | VERIFY | code thay đổi | kết quả test/lint | test xanh |
| 4 | REVIEW | code đã verify | nhận xét + duyệt | review pass |

> Khung này **không bao giờ đổi**. Bài toán Go hay React, doc xịn hay doc rác —
> chỉ đổi cái cắm vào ô (xem `principles.md`).

## 2. Cổng trung tâm: Definition of Ready (DoR)

> "Ready" là **trạng thái**, không phải loại spec. Cổng này nhận **App Spec** *hoặc*
> **Feature Spec** và chỉ cho qua khi đủ DoR. Phân loại spec đầy đủ: [`glossary.md`](glossary.md).

Mọi đầu vào hỗn loạn phải đi qua **một cổng duy nhất** rồi mới được vào Ô 2.

```
   NGÀY ĐẸP TRỜI                    NGÀY XẤU TRỜI
   doc 20 trang rất clear      PO ném "làm cái auto-translate đi"
        │                                │
        ▼ solver: EXTRACT                ▼ solver: BRAINSTORM
        │ (rút gọn doc → spec)           │ (hỏi lại, đào sâu → spec)
        └──────────────┐   ┌─────────────┘
                       ▼   ▼
              ╔═══════════════════╗
              ║  📋 SPEC đã READY  ║  ◀── CỔNG CHUNG (định dạng DUY NHẤT)
              ╚═════════╤═════════╝
                        ▼
          Ô 2,3,4 KHÔNG BIẾT đầu vào hôm nay là doc xịn hay 1 dòng.
          Chúng chỉ thấy spec-đã-Ready giống hệt nhau.
```

→ **Đầu vào biến thiên vô hạn, nhưng bị ép về một artifact chuẩn trước khi xuống
dây chuyền. Mọi hỗn loạn bị nhốt ở Ô 1.** Phần còn lại sống trong thế giới sạch sẽ.

Cổng DoR định nghĩa ở cuối từng template ([`app-spec`](../contracts/app-spec.template.md) ·
[`feature-spec`](../contracts/feature-spec.template.md)). Phân loại spec: [`glossary.md`](glossary.md).

## 3. Router: tự chọn người giải theo độ chín đầu vào

```
              ┌─────────────────────────┐
   đầu vào ──▶│   🚦 ROUTER (triage)     │  đo "độ chín" của đầu vào
              └───────────┬─────────────┘
          ┌───────────────┼───────────────┐
          ▼ rõ ràng       ▼ mơ hồ          ▼ rỗng/sai
    [EXTRACT solver]  [BRAINSTORM solver]  [chặn, hỏi PO 2-3 câu]
          └───────────────┴───────────────┘
                          ▼
                   📋 SPEC đã READY  (luôn ra cùng 1 thứ)
```

## 3b. Hai độ cao: feature vs app — và lớp DECOMPOSE

Kiến trúc *fractal*: một app là **một cây feature**, cùng bộ cổng lặp lại ở mỗi nút.
Build app từ đầu chỉ cần **thêm một viên** phía trước: `μ-decompose`.

```
   BUILD 1 FEATURE                 BUILD 1 BASE APP
   ───────────────                 ────────────────
   feature-spec                    app-spec
        │                              │
        │                              ▼ μ-decompose  ◀── viên THÊM (app altitude)
        │                          [nhiều feature-spec]
        │                              │
        ▼                         ┌────┴────┬────┐
   Intake→Build→...               ▼         ▼    ▼
                              feature   feature feature
                              (mỗi cái = chuỗi cũ, KHÔNG đổi)
```

`μ-decompose` không phải harness mới — chỉ là thêm một stage/solver đúng nguyên lý
"cố định contract, thay solver". Cổng của nó: `app-spec → [feature-spec]`.

## 3c. Toolkit, không monolith — nối qua contract

Các Ô là **micro-harness rời**, mỗi viên một việc, nối nhau qua contract. Việc
**xâu chuỗi** tách ra thành các "công thức" (recipe):

```
   /make         = CỬA VÀO: hỏi "app hay feature?" rồi dispatch xuống dưới
   /make-app     = decompose ▸ (mỗi feature: intake▸build▸verify▸review)
   /make-feature = intake ▸ build ▸ verify ▸ review
   /intake       = chạy đúng một viên
```

→ "Nối hay không" = chọn điểm vào dây chuyền. Chi tiết & ràng buộc: xem nguyên lý 6
trong [`principles.md`](principles.md).

## 4. Map ra Claude Code

| Khái niệm kiến trúc | Hiện thực trong Claude Code |
|---------------------|------------------------------|
| Cổng / artifact chuẩn | format trong `contracts/` + nhắc trong `CLAUDE.md` |
| Router + Solver | Skill / slash command (`meta/.claude/skills/`) |
| Gate tự động (fmt, test) | Hook trong `settings.json` |
| Tách việc nặng | Subagent |
| Trí nhớ xuyên phiên | Memory |
