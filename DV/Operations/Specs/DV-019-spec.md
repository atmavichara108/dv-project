---
type: spec
id: DV-019
title: "Урезать ролевую модель до admin/member/guest"
status: approved
---

## Product-контекст
Сейчас 6 ролей (admin, moderator, researcher, expert, guest, public), реально нужны 3. 
product-vision и ADR-008 фиксируют целевое состояние. Избыточные роли создают трение 
в онбординге и усложняют код проверок.

## Маппинг
admin → admin, moderator → admin, researcher → member, expert → guest, public → guest, guest → guest

## Изменяемые файлы
1. migrations/0006_simplify_roles.sql — CHECK constraint + UPDATE
2. public/static/modules/admin.js — ALL_ROLES, ROLE_LABELS, ROLE_COLORS, доступ
3. public/static/modules/auth.js — roleColors, isAdmin
4. public/static/modules/topics.js — строка 319: только admin
5. public/static/modules/media.js — строка 34: только admin
6. public/static/modules/rooms.js — строки 142, 493-494
7. docs/glossary.md — Roles
8. README.md — список ролей
9. context/DV/Operations/Kanban/DV-019.md — обновить DoD

## НЕ изменяемые
src/lib/auth.ts (requireRole универсален), src/routes/api.ts, profile.js (использует admin.js global), utils.js, scripts/

## Definition of Done
- [ ] Миграция 0006 создана и применима
- [ ] Во всех фронтенд-модулях старые роли удалены
- [ ] В админке можно назначить только admin/member/guest
- [ ] npm run build проходит
- [ ] docs/glossary.md и README.md обновлены
