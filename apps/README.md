# apps/ — đầu ra của meta-harness

Nơi chứa **app build thật bằng `meta/`**. Mục đích không phải làm sản phẩm, mà là
**bằng chứng**: meta-harness có thật sự đẻ ra được app chạy được — nhanh, đúng quy trình.

```
   meta/ (chạy quy trình)  ──build──▶  apps/<tên-app>
                                          │
                                          ▼
                              chạy được = harness đạt ✅
```

## Quy ước

- Mỗi app = 1 thư mục con.
- Mỗi app nên có ghi chú: dùng App Spec nào, qua những solver/skill nào, mất bao lâu
  → đối chiếu với `experiments/` để đo harness tốt lên hay không.

## Trạng thái

Chưa có app nào. App demo đầu tiên thuộc vòng sau (sau khi `/intake` chạy được).
