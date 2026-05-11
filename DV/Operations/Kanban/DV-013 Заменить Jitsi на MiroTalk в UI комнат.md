---
type: task
id: DV-013
title: "Заменить Jitsi на MiroTalk в UI комнат сайта"
status: backlog
priority: P1
effort: S
epic: "Conferencing"
sprint: 
assignee: max
created: 2026-05-10
due: 
tags: [frontend, integration]
---

## Цель
Перевести фичу discussion rooms на свой сервис созвонов.

## Definition of Done
- [ ] В `public/static/modules/rooms.js` Jitsi-iframe заменён на MiroTalk-iframe
- [ ] При клике «Присоединиться» комната создаётся с детерминированным room-id (например, hash от topic_id)
- [ ] Передаются user display name из сессии и роль (moderator для админов комнаты)
- [ ] FAQ обновлён: убраны упоминания Jitsi
- [ ] Старые активные комнаты не сломаны (либо мигрированы, либо помечены legacy)

## Зависимости
- Зависит от: [[DV-011 Развернуть MiroTalk SFU]]

## Заметки
Рассмотреть вариант не iframe, а deep-link с открытием в новой вкладке — упрощает интеграцию и не ломается из-за CSP.
