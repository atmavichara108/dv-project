---
type: task
id: DV-010
title: "DNS A-записи: @, www, meet → IP сервера"
status: backlog
priority: P0
effort: XS
epic: "Self-hosted Infra"
assignee: max
created: 2026-05-11
tags: [infra, dns, p0]
depends_on: [DV-005, DV-006]
---

## Goal
Настроить Advanced DNS в Namecheap.

## Definition of Done
- A @ → IP
- A www → IP
- A meet → IP
- `dig re-search.wiki` и `dig meet.re-search.wiki` возвращают правильный IP
- TTL поставлен 5 мин на время отладки, потом 30 мин

## Notes
- Записи для Resend (SPF/DKIM/DMARC) — в отдельной задаче DV-026
