package web

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"

	"github.com/gin-gonic/gin"

	"fb-pixel-mvp/internal/shopifyauth"
)

// tokenExchangeFunc đổi authorization code lấy access token cho một shop.
type tokenExchangeFunc func(shopDomain, code string) (string, error)

// Callback xử lý GET /auth/callback của Shopify OAuth:
// verify HMAC → verify state (khớp shop) → đổi code lấy token → lưu shop → redirect.
func (h *Handler) Callback(c *gin.Context) {
	q := c.Request.URL.Query()

	if !shopifyauth.VerifyHMAC(q, h.cfg.Secret) {
		c.String(http.StatusUnauthorized, "invalid hmac")
		return
	}
	stateShop, ok := shopifyauth.VerifyState(q.Get("state"), h.cfg.Secret)
	if !ok || stateShop != q.Get("shop") {
		c.String(http.StatusUnauthorized, "invalid state")
		return
	}

	token, err := h.exchange(q.Get("shop"), q.Get("code"))
	if err != nil {
		c.String(http.StatusBadGateway, "token exchange failed")
		return
	}
	if _, err := h.shops.Upsert(q.Get("shop"), token); err != nil {
		c.String(http.StatusInternalServerError, "save failed")
		return
	}

	log.Printf("OAuth install completed: shop=%s (access token saved)", q.Get("shop"))
	c.Redirect(http.StatusFound, "/?shop="+url.QueryEscape(q.Get("shop")))
}

// exchangeViaShopify gọi Shopify thật để đổi code lấy access token.
// (Phần tích hợp ngoài — verify bằng E2E khi có Shopify creds; logic Callback đã unit-test.)
func (h *Handler) exchangeViaShopify(shopDomain, code string) (string, error) {
	endpoint := "https://" + shopDomain + "/admin/oauth/access_token"
	body, _ := json.Marshal(map[string]string{
		"client_id":     h.cfg.APIKey,
		"client_secret": h.cfg.Secret,
		"code":          code,
	})
	resp, err := http.Post(endpoint, "application/json", bytes.NewReader(body))
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("shopify token exchange status %d", resp.StatusCode)
	}
	var out struct {
		AccessToken string `json:"access_token"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&out); err != nil {
		return "", err
	}
	return out.AccessToken, nil
}
