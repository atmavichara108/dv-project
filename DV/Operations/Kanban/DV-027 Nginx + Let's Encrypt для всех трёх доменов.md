---
type: task
id: DV-027
title: Nginx reverse proxy + SSL для re-search.wiki, www, meet
status: backlog
priority: P0
effort: S
epic: Self-hosted Infra
assignee: max
created: 2026-05-11
tags:
  - infra
  - nginx
  - ssl
depends_on:
  - DV-008
  - DV-011
  - DV-010
---

## Goal
Два конфига Nginx, оба с HTTPS, проксируют на 8787 (dv-hub) и 3010 (MiroTalk).

## Definition of Done
- /etc/nginx/sites-available/{re-search.wiki, meet.re-search.wiki}
- Симлинки в sites-enabled, default удалён
- `nginx -t` ок
- certbot выдал сертификаты на 3 имени, авто-renewal в cron работает
- WebSocket upgrade headers стоят (нужны и для Hono, и для MiroTalk)
- HTTP→HTTPS редирект работает

## Notes
- HTTP/2 включён
- Проверить `ssl_protocols TLSv1.2 TLSv1.3`
