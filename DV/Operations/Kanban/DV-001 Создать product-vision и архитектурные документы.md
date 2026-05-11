---
type: task
id: DV-001
title: "Создать product-vision и архитектурные документы"
status: backlog
priority: P0
effort: S
epic: "Product Vision Docs"
sprint: 
assignee: max
created: 2026-05-10
due: 
tags: [docs, opencode, foundation]
---

## Цель
Дать opencode-агентам общий продуктовый контекст: что такое DV Hub, какие архитектурные решения приняты, какие термины используются. Без этого агенты пишут технически корректный, но продуктово бессмысленный код.

## Definition of Done
- [ ] В репозитории dv-hub создана папка `docs/`
- [ ] Создан `docs/product-vision.md` (что такое DV Hub, не-цели, ядро воронки, аудитория)
- [ ] Создан `docs/architecture.md` (ADR-001..006: self-host, MiroTalk, Meetily, Twake, чат, deep_research)
- [ ] Создан `docs/glossary.md` (ячейка, тема, материал, комната, синтез, публикация, consent)
- [ ] Создан `docs/roadmap.md` — пока пустой каркас с тремя фазами

## Зависимости
- Блокирует: [[DV-003 Настроить opencode.json]], [[DV-004 Создать промпты агентов opencode]]

## Заметки
Черновики всех файлов уже есть в переписке с агентом-стратегом. Нужно скопировать, проверить актуальность под текущее состояние, закоммитить.
