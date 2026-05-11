---
type: task
id: DV-011
title: "Развернуть MiroTalk SFU на meet.re-search.wiki"
status: backlog
priority: P0
effort: M
epic: "Calls & Transcription"
assignee: max
created: 2026-05-11
tags: [infra, calls, p0]
depends_on: [DV-006a]
---

## Goal
MiroTalk SFU работает в production-режиме за Nginx + HTTPS.

## Definition of Done
- Клон в /opt/mirotalksfu
- .env: ENVIRONMENT=production, SFU_ANNOUNCED_IP=<публичный IP>
- config.js: задан HOST_USERNAME/PASSWORD или OIDC-заглушка
- pm2 start, автозапуск настроен
- Порты 40000–40100 tcp+udp открыты (см. DV-006a)
- Тестовый созвон с двух устройств работает

## Notes
- Memory footprint mediasoup ощутимый — мониторь RAM
- Если 5 GB не хватит при 10+ участниках, апгрейдь VPS
