---
type: task
id: DV-041
title: "Правка deploy-mirotalk.sh на сервере (точечный фикс вне репо)"
status: ready
priority: P2
effort: XS
epic: Conferencing
sprint:
assignee: max
created: 2026-11-06
due:
tags: [mirotalk, deploy, debt]
---

## Цель
Фикс под версию 2.7.65 внесён точечно на сервере, а сам скрипт в репозитории остался прежним: переустановка повторит поломку.

## Definition of Done
- [ ] Правка перенесена с сервера в `scripts/deploy-mirotalk.sh`
- [ ] Скрипт прогнан на чистом сервере
- [ ] Правка влита в main

## Зависимости
- Зависит от: [[DV-033 Деплой MiroTalk 2.7.65]]

## Заметки
Ждёт рук исполнителя. Без этого деплой невоспроизводим.
