# meta/ — meta-harness trên Claude Code ★ trọng tâm

Đứng trên vai Claude Code, biến quy trình lặp đi lặp lại thành **quy trình thống nhất**:
gõ 1 lệnh thay vì prompt thủ công 5 lần.

## Thành phần

```
meta/
├── solvers/        🔌 "người giải" cắm rời — mỗi solver khớp 1 contract
│   ├── extract/      doc xịn → spec đã Ready (fast path)
│   └── brainstorm/   doc mơ hồ → spec đã Ready (human-in-loop)
└── .claude/
    └── skills/      ⚡ slash command đóng gói quy trình (/make, /intake, /make-feature)
```

## Quan hệ với phần còn lại của lab

```
   knowledge/ + contracts/  ──(đọc lý thuyết + cổng)──▶  meta/
                                                          │ chạy quy trình
                                                          ▼
                                                        apps/  (đầu ra)
```

## Lộ trình

- [x] Skill **`/make-app`** → `.claude/skills/make-app/SKILL.md` (intake B1–B7 → DoR → decompose)
- [x] Skill **`/make-feature`** → `.claude/skills/make-feature/SKILL.md` (intake-feature → DoR → μ-build TDD)
- [x] Skill **`/make`** → `.claude/skills/make/SKILL.md` (router: hỏi app/feature → dispatch)
- [ ] Skill `/intake` (chạy 1 viên μ-intake riêng)
- [ ] Test bộ skill trên case mới + promote sang `~/.claude/skills/` (gõ được mọi phiên)
- [ ] Solver `extract`/`brainstorm` tách riêng · stage `μ-design` · Hook gate (fmt/test)
