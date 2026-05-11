---
type: task
id: DV-008
title: "Деплой dv-hub: Hono Node adapter + PM2"
status: backlog
priority: P0
effort: M
epic: "Self-hosted Infra"
assignee: max
created: 2026-05-11
tags: [infra, deploy, p0]
depends_on: [DV-006a, DV-007]
---

## Goal
Перевести Hono-приложение с Cloudflare Workers на Node.js runtime, запустить через PM2.

## Definition of Done
- Добавлен `@hono/node-server` как entry point (`src/server.ts`)
- Bindings из wrangler заменены на env-переменные + better-sqlite3 (или libSQL)
- `npm run build` собирает Node-бандл
- pm2 start, pm2 save, pm2 startup настроены
- Приложение слушает 8787 на localhost
- curl localhost:8787 возвращает ответ

## Notes
- Это самая большая задача этапа — D1 → SQLite-файл требует переписать DB-слой (см. DV-007)
- Раздели на подзадачи если затянется
