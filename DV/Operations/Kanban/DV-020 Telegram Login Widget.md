---
type: task
id: DV-020
title: "Решить судьбу Telegram Login Widget"
status: backlog
priority: P2
effort: S
epic: "Roles & Auth"
sprint: 
assignee: max
created: 2026-05-10
due: 
tags: [auth, decision]
---

## Цель
TG Login Widget удобен, но привязывает к Telegram-инфраструктуре и требует publicly reachable callback. После переезда возможны нюансы. Решить — оставлять или нет.

## Definition of Done
- [ ] Принято решение: оставить как опцию / убрать совсем / оставить только для админов
- [ ] Решение зафиксировано в ADR
- [ ] Если убрали — magic-link через email становится единственным методом, проверена доставка через Resend (или замена)
- [ ] Если оставили — проверено, что widget работает с нового домена

## Зависимости
- Зависит от: [[DV-010 DNS-записи для re-search.wiki]]
