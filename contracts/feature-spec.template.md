# Feature Spec (cấp FEATURE / USER STORY) — Template chuẩn

> Spec cho **một feature** = một User Story. **Ráp từ chuẩn ngành** để khỏi chứng minh
> tính thực tế với team:
>
> | Phần | Dựa trên chuẩn |
> |------|----------------|
> | Khung story | **User Story / Connextra** (Mike Cohn) |
> | Tiêu chí nghiệm thu | **Gherkin** Given/When/Then — BDD (Dan North · Cucumber) |
> | Chất lượng story | **INVEST** (Bill Wake) |
> | Tinh thần story | **3 C's** Card·Conversation·Confirmation (Ron Jeffries) |
> | Cổng vào/ra | **Definition of Ready / Done** (Scrum) |
>
> Quy ước: 🔒 = mục **bắt buộc** (gate). Story sinh ra từ `μ-decompose` của một App Spec.

---

## 0. Metadata
- **ID:** `<story id>`
- **🔗 Thuộc App Spec (Epic cha):** [`./app-spec.md`](./app-spec.md)  ← *traceability*
- **🔗 Design Doc (nếu chạm kỹ thuật):** [`./design-doc.md`](./design-doc.md)
- **Trạng thái:** Draft → **Ready** → In progress → Done
- **Ước lượng (estimate):** `<điểm/size>`  *(INVEST — Estimable)*

## 1. User Story  🔒  *(Card — Connextra)*
> **As a** `<role>`, **I want** `<goal>`, **so that** `<benefit>`.

## 2. Acceptance Criteria  🔒  *(Confirmation — Gherkin/BDD)*
> Mỗi tiêu chí = một kịch bản **Given / When / Then**, nghiệm thu **nhị phân** (pass/fail).

### Scenario 1 — `<tên>`
- **Given** `<bối cảnh>`
- **When** `<hành động>`
- **Then** `<kết quả mong đợi>`

## 3. Phạm vi  🔒
- **Trong:** `<...>`
- **Ngoài (out of scope):** `<...>`

## 4. Ràng buộc & quy ước
- `<pattern/convention dự án; trỏ Design Doc nếu cần chi tiết kỹ thuật>`

## 5. Phụ thuộc (Dependencies)
> Lưu ý INVEST — *Independent*: tối thiểu hoá phụ thuộc.
- `<story/khối cần có trước>`

## 6. Ghi chú trao đổi  *(Conversation — 3 C's)*
> Story là **lời mời trao đổi**, không phải hợp đồng đóng băng (*Negotiable*). Ghi câu
> hỏi / quyết định phát sinh khi bàn ở đây.
- `<...>`

## 7. Rủi ro & câu hỏi mở
- `<...>`

---

## ✅ Definition of Ready (GATE VÀO) — đủ rõ để bắt tay code
- [ ] **§1** User Story đúng dạng As a/I want/So that, **giá trị rõ** *(Valuable)*.
- [ ] **§2** Có ≥1 Acceptance Criteria Gherkin, **kiểm thử được** *(Testable)*.
- [ ] **§3** Phạm vi trong/ngoài rõ.
- [ ] Đủ **nhỏ** để làm trong 1 đơn vị sprint *(Small)*.
- [ ] **Ước lượng được** *(Estimable)*.
- [ ] Phụ thuộc rõ & tối thiểu *(Independent)*.
- [ ] **§0** Có link về Epic cha *(traceable)*.

## ✅ Definition of Done (GATE RA) — coi như xong
- [ ] Tất cả Acceptance Criteria (§2) **pass**.
- [ ] `<test / review / merge theo convention dự án>`.

## 🔎 INVEST self-check *(quét nhanh 6 tiêu chí chất lượng story — Bill Wake)*
- [ ] **I**ndependent — phụ thuộc tối thiểu (§5).
- [ ] **N**egotiable — còn chỗ trao đổi, chưa đóng băng giải pháp (§6).
- [ ] **V**aluable — story nêu rõ "so that <giá trị>" (§1).
- [ ] **E**stimable — ước lượng được (§0).
- [ ] **S**mall — gọn trong 1 đơn vị sprint.
- [ ] **T**estable — mọi AC nghiệm thu nhị phân (§2).

---
*Chuẩn tham chiếu: User Story (Cohn/Connextra) · Gherkin BDD (North/Cucumber) ·
INVEST (Wake) · 3 C's (Jeffries) · Definition of Ready/Done (Scrum).*
