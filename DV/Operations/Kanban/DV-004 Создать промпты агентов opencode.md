---
type: task
id: DV-004
title: Создать промпты агентов opencode (plan, build, reviewer, researcher)
status: doing
priority: P0
effort: S
epic: Product Vision Docs
sprint:
assignee: max
created: 2026-05-10
due:
tags:
  - opencode
  - agents
---

## Цель
Разделить роли LLM-агентов: стратег не пишет код, исполнитель не выдумывает архитектуру, ревьюер не редактирует, ресёрчер не вызывает build.

## Definition of Done
- [ ] Создана папка `.opencode/agents/`
- [ ] Создан `plan.md` (mode: primary, edit: deny, температура 0.2)
- [ ] Создан `build.md` (mode: primary, edit: allow, температура 0.1)
- [ ] Создан `reviewer.md` (mode: subagent, read-only)
- [ ] Создан `researcher.md` (mode: subagent, webfetch+websearch)
- [ ] Tab в opencode переключает между plan и build
- [ ] @reviewer и @researcher вызываются через mention

## Зависимости
- Зависит от: [[DV-003 Настроить opencode.json]]

## Заметки
Черновики промптов есть в переписке с агентом. После создания — пройти ими простой кейс (попросить @plan описать рефакторинг главной).
