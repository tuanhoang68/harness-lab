package pixel

import "testing"

// F2 Scenario 2 — validate chặn định dạng sai; Scenario 1 — định dạng đúng thì qua.
func TestValidatePixelID(t *testing.T) {
	cases := []struct {
		name string
		in   string
		ok   bool
	}{
		{"hợp lệ 16 chữ số", "1234567890123456", true},
		{"hợp lệ 15 chữ số", "123456789012345", true},
		{"có chữ cái", "12345abc90123", false},
		{"rỗng", "", false},
		{"quá ngắn", "123", false},
		{"có khoảng trắng", "123 456 789012", false},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			if got := ValidatePixelID(c.in); got != c.ok {
				t.Errorf("ValidatePixelID(%q) = %v, want %v", c.in, got, c.ok)
			}
		})
	}
}
