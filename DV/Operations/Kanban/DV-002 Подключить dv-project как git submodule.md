---
type: task
id: DV-002
title: Подключить dv-project как git submodule
status: doing
priority: P0
effort: XS
epic: Product Vision Docs
sprint:
assignee: max
created: 2026-05-10
due:
tags:
  - git
  - opencode
  - infra
file: DV/Operations/Kanban/DV-003 Настроить opencode.json.md
---

## Цель
Дать opencode-агентам доступ к контексту движения (Манифест, Структура, Воркфлоу) из волта без копирования файлов. Submodule даёт version-pinning и совместимость с CI.

## Definition of Done
- [ ] В корне dv-hub выполнено `git submodule add https://github.com/atmavichara108/dv-project.git context`
- [ ] Submodule закоммичен
- [ ] В `package.json` добавлен скрипт `"context:sync": "cd context && git pull origin main && cd .. && git add context"`
- [ ] В README dv-hub добавлена секция «Cloning» с упоминанием `--recurse-submodules`
- [ ] Проверено: `git clone --recurse-submodules ...` восстанавливает контекст

## Зависимости
- Блокирует: [[DV-003 Настроить opencode.json]]

## Заметки
В .gitignore волта добавить `.obsidian/plugins/obsidian-git/main.js` и подобные тяжёлые бинари — не нужны агентам и шумят контекст.
