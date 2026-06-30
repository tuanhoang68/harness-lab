package webpixel

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

// NewShopifyRegistrar trả về Registrar đăng ký Web Pixel qua Shopify Admin GraphQL
// (webPixelCreate). settings gửi: accountID (Pixel ID) + collectURL (backend nhận event).
//
// E2E — chưa unit-test (logic Activate đã test với registrar inject). Field trong `settings`
// phải khớp schema của Web Pixel extension đã deploy (R1).
func NewShopifyRegistrar(collectURL string) Registrar {
	return func(shopDomain, accessToken, pixelID string) (string, error) {
		return registerWebPixel(shopDomain, accessToken, pixelID, collectURL)
	}
}

func registerWebPixel(shopDomain, accessToken, pixelID, collectURL string) (string, error) {
	settings, _ := json.Marshal(map[string]string{"accountID": pixelID, "collectURL": collectURL})
	query, _ := json.Marshal(map[string]string{
		"query": fmt.Sprintf(
			`mutation { webPixelCreate(webPixel: { settings: %q }) { webPixel { id } userErrors { message } } }`,
			string(settings),
		),
	})

	endpoint := "https://" + shopDomain + "/admin/api/2024-01/graphql.json"
	req, err := http.NewRequest(http.MethodPost, endpoint, bytes.NewReader(query))
	if err != nil {
		return "", err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-Shopify-Access-Token", accessToken)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("webPixelCreate status %d", resp.StatusCode)
	}

	var out struct {
		Data struct {
			WebPixelCreate struct {
				WebPixel   struct{ ID string } `json:"webPixel"`
				UserErrors []struct{ Message string } `json:"userErrors"`
			} `json:"webPixelCreate"`
		} `json:"data"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&out); err != nil {
		return "", err
	}
	if len(out.Data.WebPixelCreate.UserErrors) > 0 {
		return "", fmt.Errorf("webPixelCreate: %s", out.Data.WebPixelCreate.UserErrors[0].Message)
	}
	return out.Data.WebPixelCreate.WebPixel.ID, nil
}
