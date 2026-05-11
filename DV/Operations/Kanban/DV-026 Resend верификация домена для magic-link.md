---
type: task
id: DV-026
title: "Resend: верифицировать re-search.wiki для magic-link email"
status: backlog
priority: P1
effort: XS
epic: "Self-hosted Infra"
assignee: max
created: 2026-05-11
tags: [email, p1]
depends_on: [DV-010]
---

## Goal
Slать magic-link с адреса вроде noreply@re-search.wiki.

## Definition of Done
- Домен добавлен в Resend
- SPF, DKIM, DMARC записи добавлены в Namecheap Advanced DNS
- Статус "Verified" в Resend
- Тестовое письмо доходит до Gmail и Proton без спама
- RESEND_API_KEY в .env на сервере
