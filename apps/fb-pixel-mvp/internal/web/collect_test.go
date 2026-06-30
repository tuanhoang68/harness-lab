package web

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
)

func serveCollect(h *Handler, body string) *httptest.ResponseRecorder {
	gin.SetMode(gin.TestMode)
	r := gin.New()
	r.POST("/collect", h.Collect)
	req := httptest.NewRequest(http.MethodPost, "/collect", strings.NewReader(body))
	req.Header.Set("Content-Type", "text/plain")
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	return w
}

// F3 (🅐) — backend nhận event từ web pixel storefront: trả 200 + CORS để fetch cross-origin được.
func TestCollect_Accepts200WithCORS(t *testing.T) {
	w := serveCollect(testHandler(), `{"event":"page_viewed","accountID":"1234567890123456"}`)
	if w.Code != http.StatusOK {
		t.Fatalf("status = %d, want 200", w.Code)
	}
	if w.Header().Get("Access-Control-Allow-Origin") != "*" {
		t.Errorf("thiếu CORS header Access-Control-Allow-Origin: *")
	}
}
