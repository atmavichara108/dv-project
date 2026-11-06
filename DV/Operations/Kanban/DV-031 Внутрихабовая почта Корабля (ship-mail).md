---
type: task
id: DV-031
title: "Внутрихабовая почта Корабля (ship-mail) — пилот"
status: done
priority: P1
effort: M
epic: "Ship Mail"
sprint:
assignee: max
created: 2026-11-06
due: 
tags: [ship-mail, agents, transport]
---

## Цель
У агентских сессий нет штатного канала переписки: всё шло через чат владельца, и он становился ручным ретранслятором. Сделать файловый транспорт писем между сессиями.

## Definition of Done
- [x] `scripts/ship-mail.sh` — inbox / show / ack / send
- [x] Ящик — директория `~/.local/share/ship-mail/inbox/<sid>/`, письма как `.md` с frontmatter
- [x] Frontmatter: `from`, `to`, `title`, `ref`, `ts`, `id` — id уникален, коллизии исключены
- [x] Спека `docs/specs/ship-mail-pilot.md`
- [x] Адресация только через peer-registry (INC-18) — чужие id из чата не принимаются
- [x] Первый шаг хода агента: `inbox <свой sid>`; после отработки — `ack`

## Зависимости
- Зависит от: [[DV-004 Создать промпты агентов opencode]]

## Заметки
Письмо — транспорт, не память. Канонизация в `docs/decisions.md` — отдельный шаг получателя.
