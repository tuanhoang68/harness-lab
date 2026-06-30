# App Spec (cấp APP / EPIC) — Template chuẩn

> Mẫu spec cấp app/dự án. **Ráp từ chuẩn ngành**, không tự chế — để khỏi mất công
> chứng minh tính thực tế với team:
>
> | Phần | Dựa trên chuẩn |
> |------|----------------|
> | Cấu trúc tổng | **PRD** — Product Requirements Document (Atlassian/SVPG) |
> | Phân rã app→feature | **Agile**: Initiative ▸ **Epic** ▸ Story ▸ Task (Scrum/SAFe) |
> | Bối cảnh/vấn đề | **Working Backwards** (Amazon) — bắt đầu từ vấn đề khách hàng |
> | Mục tiêu & metric | **SMART** goals |
> | Ưu tiên phạm vi | **MoSCoW** (Must/Should/Could/Won't) |
> | Cổng "đủ rõ để build" | **Definition of Ready (DoR)** + **DoD** (Scrum) |
> | Chất lượng story con | **INVEST** (Bill Wake) |
>
> Quy ước: 🔒 = mục **bắt buộc** (gate). Thiếu bất kỳ 🔒 nào → chưa "Ready".

---

> **Phạm vi tài liệu:** chỉ **WHAT / WHY** (sản phẩm). Phần **HOW** (tech stack, mô
> hình dữ liệu, các luồng) nằm ở **Design Doc** liên kết bên dưới — để app-spec gọn.

## 0. Metadata
- **ID:** `<app/exp id>`
- **🔗 Design Doc (kỹ thuật):** [`./design-doc.md`](./design-doc.md)  ← *tech stack · data · flows*
- **Tác giả / Người duyệt:** `<...>`
- **Trạng thái:** Draft → **Ready** → Approved
- **Oracle / nguồn tham chiếu:** `<repo, design, doc>`

## 1. Bối cảnh & Vấn đề  🔒  *(Working Backwards)*
> Người dùng là ai, đang gặp vấn đề gì, vì sao đáng giải. **KHÔNG mô tả giải pháp.**

`<...>`

## 2. Mục tiêu & Chỉ số thành công  🔒  *(SMART)*
> Mục tiêu sản phẩm + 1–3 metric **đo được** (Specific, Measurable, Achievable,
> Relevant, Time-bound).

- **Mục tiêu:** `<...>`
- **Success metric:** `<vd: X% ..., N event/ngày, ...>`

## 3. Definition of Done (cấp app)  🔒
> Điều kiện nghiệm thu cả app. Mỗi dòng phải **kiểm chứng được**.

- [ ] `<tiêu chí đo được>`

## 4. Phạm vi  🔒  *(MoSCoW)*
- **Must have:** `<lõi không có thì app vô nghĩa>`
- **Should / Could:** `<có thì tốt>`
- **Won't (Out of scope):** `<chốt rõ KHÔNG làm — chống phình>`

## 5. Ràng buộc · Giả định · Phụ thuộc  🔒
- **Ràng buộc (Constraints):** `<nền tảng, kỹ thuật, pattern bắt buộc>`
- **Giả định (Assumptions):** `<điều coi là đúng>`
- **Phụ thuộc (Dependencies):** `<hệ thống/bên thứ 3>`

## 6. Phân rã: Epic → User Stories  🔒  *(riêng cấp app — đầu vào của μ-decompose)*
> Mỗi dòng = một feature-spec sẽ tạo theo `feature-spec.template.md`. Mỗi story nên đạt
> **INVEST**.

| # | User Story | Giá trị | Phụ thuộc | Module oracle |
|---|-----------|---------|-----------|---------------|
| 1 | `<...>` | `<...>` | — | `<...>` |

## 7. Rủi ro & Câu hỏi mở  *(Risks / Open questions)*
- `<điều chưa chắc, cần xác nhận trước khi build>`

## 8. Tham chiếu
- **Design Doc (HOW):** [`./design-doc.md`](./design-doc.md)
- `<links: oracle repo, design, tài liệu khác>`

---

## ✅ Definition of Ready (GATE) — cổng mở khi TẤT CẢ đúng
- [ ] **§1** Vấn đề rõ, không lẫn giải pháp.
- [ ] **§2** Có ≥1 success metric **đo được** (SMART).
- [ ] **§3** DoD có tiêu chí **nghiệm thu được**.
- [ ] **§4** Có cả **Must** và **Won't/Out-of-scope**.
- [ ] **§5** Nêu được ràng buộc / giả định / phụ thuộc.
- [ ] **§6** Phân rã ra ≥1 story, mỗi story đạt **INVEST**.

→ Thiếu bất kỳ mục nào: **chưa Ready** → quay lại intake đào tiếp. KHÔNG xuống build.

---
*Chuẩn tham chiếu: PRD · Agile Epic/Story (Scrum, SAFe) · SMART · MoSCoW ·
Amazon Working Backwards · Definition of Ready/Done · INVEST.*
