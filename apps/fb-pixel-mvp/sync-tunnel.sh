#!/usr/bin/env bash
# sync-tunnel.sh — đồng bộ URL tunnel vào .env (APP_URL) + shopify.app.toml
# (application_url + redirect_urls). Khỏi sửa tay 2 chỗ mỗi lần tunnel đổi.
#
# Dùng:
#   ./sync-tunnel.sh https://abc.trycloudflare.com   # sync 1 URL có sẵn
#   ./sync-tunnel.sh                                  # tự mở cloudflared, lấy URL, sync
set -euo pipefail
cd "$(dirname "$0")"

URL="${1:-}"

if [ -z "$URL" ]; then
  echo "→ Mở cloudflared tunnel (localhost:8099)…"
  LOG="$(mktemp)"
  nohup cloudflared tunnel --url http://localhost:8099 --no-autoupdate >"$LOG" 2>&1 &
  TPID=$!
  disown "$TPID" 2>/dev/null || true
  echo "  cloudflared chạy nền PID=$TPID  (dừng: kill $TPID)"
  for _ in $(seq 1 20); do
    URL="$(grep -oE 'https://[a-z0-9-]+\.trycloudflare\.com' "$LOG" | head -1 || true)"
    [ -n "$URL" ] && break
    sleep 1
  done
  [ -z "$URL" ] && { echo "❌ Không lấy được URL tunnel — xem $LOG"; exit 1; }
fi

URL="${URL%/}"   # bỏ dấu '/' cuối nếu có
echo "→ Tunnel URL: $URL"

# 1) .env : APP_URL
if grep -q '^APP_URL=' .env 2>/dev/null; then
  sed -i -E "s#^APP_URL=.*#APP_URL=${URL}#" .env
else
  echo "APP_URL=${URL}" >> .env
fi

# 2) shopify.app.toml : application_url + redirect_urls
sed -i -E "s#^application_url = \".*\"#application_url = \"${URL}/\"#" shopify.app.toml
sed -i -E "s#https://[a-z0-9-]+\.trycloudflare\.com/auth/callback#${URL}/auth/callback#" shopify.app.toml

echo "✅ Sync xong:"
echo "   .env            APP_URL          = ${URL}"
echo "   shopify.app.toml application_url  = ${URL}/"
echo "   shopify.app.toml redirect_urls    = ${URL}/auth/callback"
echo
echo "Bước tiếp:"
echo "   shopify app deploy     # đẩy URL+scope+extension lên Partner app"
echo "   go run .               # chạy backend (đọc .env)"
