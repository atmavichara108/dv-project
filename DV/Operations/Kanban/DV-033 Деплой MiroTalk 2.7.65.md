---
type: task
id: DV-033
title: "Деплой MiroTalk 2.7.65: фикс скрипта (PR #4)"
status: done
priority: P0
effort: M
epic: Conferencing
sprint:
assignee: max
created: 2026-11-06
due:
tags: [mirotalk, deploy, video]
---

## Цель
Версия 2.7.65 требовала ручных правок на сервере, которых не было в репозитории: сервис поднимался только вручную.

## Definition of Done
- [x] IP подставляется в `.env` деплой-скриптом
- [x] `config.js` берётся из шаблона
- [x] `JWT_SECRET` обязателен и генерируется на сервере
- [x] Запуск на Node LTS
- [x] PR #4 влит в main
- [ ] Точечный фикс перенесён с сервера в сам скрипт (см. DV-041)

## Зависимости
- Зависит от: [[DV-011 Развернуть MiroTalk SFU]]
