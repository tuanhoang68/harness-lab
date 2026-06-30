# Technical Design Doc (TDD) — Template chuẩn

> Phần **HOW** của một app — đi kèm (không thay) App Spec (phần WHAT/WHY). **Ráp từ
> chuẩn ngành** để khỏi chứng minh tính thực tế với team:
>
> | Phần | Dựa trên chuẩn |
> |------|----------------|
> | Khung tổng | **Design Doc / RFC** (Google design-doc culture · IETF RFC) |
> | Kiến trúc | **C4 model** (Simon Brown) · **arc42** |
> | Các luồng (nhiều loại) | **4+1 View Model** (Kruchten) — scenarios |
> | Quyết định công nghệ | **ADR** (Michael Nygard) |
> | Mô hình dữ liệu | **ER model** (Chen) |
>
> Mức độ: **đủ để hình dung mọi mặt, KHÔNG đi sâu chi tiết.**

---

## 0. Metadata
- **ID:** `<app/exp id>`
- **🔗 App Spec (PRD) liên quan:** [`./app-spec.md`](./app-spec.md)  ← *what/why ở đây*
- **Tác giả / Reviewer:** `<...>`
- **Trạng thái:** Draft → Reviewed → Approved

> Tài liệu này chỉ chứa **kỹ thuật (HOW)**. Mục tiêu/phạm vi/DoD sản phẩm → xem App Spec.

## 1. Bối cảnh kỹ thuật  *(tóm tắt — không lặp lại PRD)*
> 1–2 câu khung kỹ thuật, trỏ về App Spec cho WHAT/WHY.

`<...>`

## 2. Tech stack & lý do chọn  *(ADR-style)*
| Lớp | Chọn | Vì sao | Phương án đã cân nhắc |
|-----|------|--------|------------------------|
| Ngôn ngữ / runtime | `<...>` | `<...>` | `<...>` |
| Datastore | `<...>` | `<...>` | `<...>` |
| Framework / lib chính | `<...>` | `<...>` | `<...>` |

## 3. Kiến trúc tổng thể  *(C4 — mức Context + Container)*
> Hệ thống gồm những khối (container) nào, mỗi khối trách nhiệm gì, nói chuyện với ai.
> Sơ đồ ASCII hoặc bullet là đủ — chưa cần mức Component/Code.

```
<sơ đồ khối: client ─ backend ─ datastore ─ 3rd-party>
```

## 4. Mô hình dữ liệu  *(ER — mức thực thể)*
> Thực thể chính + quan hệ + lưu ở đâu. Không cần liệt kê đủ cột.

| Thực thể | Vai trò | Quan hệ chính | Lưu ở |
|----------|---------|---------------|-------|
| `<...>` | `<...>` | `<...>` | `<...>` |

## 5. Các luồng chính  *(4+1 scenarios — nhiều loại)*
> Liệt kê các luồng quan trọng, mỗi luồng 3–5 bước (sequence rút gọn).

### Flow A — `<tên>`
1. `<...>`

### Flow B — `<tên>`
1. `<...>`

## 6. Tích hợp & giao diện ngoài
> API / webhook / 3rd-party + cách xác thực.

- `<...>`

## 7. Non-goals kỹ thuật & rủi ro
- **Non-goals:** `<cố tình KHÔNG làm về mặt kỹ thuật>`
- **Rủi ro + giảm thiểu:** `<...>`

## 8. Quyết định kiến trúc *(ADR log)*
> Link tới ADR chi tiết trong `knowledge/decisions/` nếu có.

- `<...>`

---

## ✅ Design Ready (checklist)
- [ ] Tech stack đã chốt + có lý do (§2).
- [ ] Có sơ đồ kiến trúc mức container (§3).
- [ ] Mô hình dữ liệu nêu được thực thể chính (§4).
- [ ] ≥1 luồng chính mô tả được đầu→cuối (§5).
- [ ] Tích hợp ngoài liệt kê đủ (§6).

---
*Chuẩn tham chiếu: Design Doc/RFC · C4 model · arc42 · 4+1 View Model (Kruchten) ·
ADR (Nygard) · ER model (Chen).*
