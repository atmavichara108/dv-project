---
type: task
id: DV-019
title: "Урезать ролевую модель до admin/member/guest"
status: backlog
priority: P1
effort: S
epic: "Roles & Auth"
sprint: 
assignee: max
created: 2026-05-10
due: 
tags: [auth, refactor]
---

## Цель
Сейчас 6 ролей, реально нужны 3. Лишняя сложность создаёт трение в онбординге и в коде проверок.

## Definition of Done
- [ ] Миграция БД: значения `researcher`, `expert`, `moderator` → `member`
- [ ] В коде middleware и UI остались только admin/member/guest
- [ ] Документ `docs/roles.md` описывает что может каждая роль
- [ ] Страница профиля показывает корректную роль

## Зависимости
- Блокирует: [[DV-017 FAQ-копирайт]]

## Заметки
Если позже понадобится «эксперт» как видимый ярлык — это можно сделать через тег профиля, не через роль.
