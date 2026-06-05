---
type: task
id: DV-006a
title: "Базовая настройка сервера: SSH ключи, ufw, пакеты"
status: done
priority: P0
effort: S
epic: Self-hosted Infra
assignee: max
created: 2026-05-11
tags:
  - infra
  - security
  - p0
due: 2026-06-05
---

## Goal
Пол: Ubuntu 24.04 LTS (вместо 22.04)
Подготовить Ubuntu 24.04 LTS к продакшен-нагрузке: SSH-ключи, файрвол, базовые пакеты, swap.

## Definition of Done
- SSH вход по паролю отключён, root-логин ограничен
- ufw настроен: 28108, 80, 443, 3010, 40000-40100 tcp/udp
- Установлены: build-essential, git, curl, wget, unzip, nginx, snapd, ffmpeg
- Node.js LTS через nvm, pm2 глобально, certbot через snap
- fail2ban установлен

## Notes
- Команды — в плане, этап 4
- Создать non-root пользователя `dv`, всё дальше делать от него
