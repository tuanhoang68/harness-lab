# harness-lab

Phòng thí nghiệm nghiên cứu **harness engineering với Claude** — xây dựng mọi thứ
*bao quanh* model/agent để tự động hóa việc code: nhanh nhất có thể, quy trình
thống nhất, không phải prompt đơn điệu lặp đi lặp lại.

> 🧭 **Lộ trình & vị trí hiện tại:** [`PLAN.md`](PLAN.md) (bản đồ sống).

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

## Bản đồ lab

```
        knowledge/  +  contracts/        ← lý thuyết + cổng chuẩn (dùng chung)
               │
               ▼
        ┌──────────────┐
        │    meta/     │   ★ TRỌNG TÂM: meta-harness trên Claude Code
        └──────┬───────┘
               │ chạy quy trình meta để build
               ▼
        ┌──────────────┐
        │    apps/     │   ◀── đầu ra của meta: app build thật để chứng minh
        └──────────────┘

        ┌──────────────┐
        │   engines/   │   💤 nhánh A (để dành): harness tự xây bằng Go
        └──────────────┘
```

## Thư mục

| Thư mục       | Vai trò                                                             |
|---------------|--------------------------------------------------------------------|
| `knowledge/`  | Kiến thức chung nhất — lý thuyết, nguyên lý, **từ điển spec** ([glossary](knowledge/glossary.md)), ADR. |
| `contracts/`  | Template **3 loại spec** chuẩn: App Spec · Feature Spec · Design Doc. |
| `meta/`       | ★ Trọng tâm. Meta-harness trên Claude Code: solvers + skills.       |
| `apps/`       | Đầu ra của `meta/` — app build thật, làm bằng chứng harness hoạt động. |
| `engines/`    | 💤 Để dành. Nhánh A: harness tự xây từ Claude API (Go).            |
| `experiments/`| Nhật ký thực nghiệm — input → output → nhận xét. Bằng chứng, không lý thuyết suông. |

## Nguyên lý cốt lõi

> **Cố định các "khớp nối" (contract), thay người làm bên trong (solver).**
> Kiến trúc không đổi. Đầu vào dễ hay khó chỉ đổi *ai* được gọi vào từng ô.

Chi tiết: [`knowledge/principles.md`](knowledge/principles.md),
[`knowledge/architecture.md`](knowledge/architecture.md).

## Trạng thái

- [x] Khung lab + lý thuyết nền (MVP vòng 1)
- [x] Bộ 3 template spec chuẩn (App/Feature/Design) + từ điển phân loại
- [ ] Cổng DoR hoàn chỉnh + skill `/intake`
- [ ] Solver `extract` + `brainstorm`
- [ ] App demo đầu tiên trong `apps/`
- [ ] (để dành) move 2 project Go vào `engines/`
