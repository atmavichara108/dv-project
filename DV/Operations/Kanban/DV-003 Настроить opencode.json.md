---
type: task
id: DV-003
title: Настроить opencode.json
status: done
priority: P0
effort: S
epic: Product Vision Docs
sprint:
assignee: max
created: 2026-05-10
due: 2026-05-13
tags:
  - opencode
  - config
file: DV/Operations/Kanban/DV-002 Подключить dv-project как git submodule.md
---


## Цель
Подключить продуктовый контекст из docs/ и context/ как instructions для всех агентов opencode. Прописать базовые permissions.

## Definition of Done
- [ ] В корне dv-hub создан `opencode.json` с правильным $schema
- [ ] В `instructions` подключены: docs/product-vision.md, architecture.md, glossary.md, roadmap.md
- [ ] В `instructions` подключены ключевые файлы из context/DV/Movement и context/DV/Structure
- [ ] Permissions: edit=ask, bash=ask с whitelist на git status/diff и npm run *
- [ ] Запуск `opencode` в проекте показывает в логах загруженные instructions

## Зависимости
- Блокирует: [[DV-004 Создать промпты агентов opencode]]
- Зависит от: [[DV-001 Создать product-vision и архитектурные документы]], [[DV-002 Подключить dv-project как git submodule]]

## Заметки
Следить, чтобы в instructions не попадали тяжёлые файлы (>50KB) — раздувают каждый запрос.
