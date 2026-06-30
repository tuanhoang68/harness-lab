import {register} from '@shopify/web-pixels-extension';

// Web Pixel chạy trong sandbox "strict": KHÔNG load được fbq client-side (R1).
// → Khi storefront có page_viewed, fetch event về backend /collect (Bậc 2 🅐).
// settings.accountID = Facebook Pixel ID; settings.collectURL = backend endpoint
// (cả hai do backend gửi qua webPixelCreate).
register(({analytics, settings}) => {
  const endpoint = settings.collectURL;
  if (!endpoint) return;

  analytics.subscribe('page_viewed', (event) => {
    fetch(endpoint, {
      method: 'POST',
      headers: {'Content-Type': 'text/plain'}, // simple request → tránh CORS preflight
      keepalive: true,
      body: JSON.stringify({
        event: 'page_viewed',
        accountID: settings.accountID,
        url: event?.context?.document?.location?.href,
        ts: event?.timestamp,
      }),
    });
  });
});
