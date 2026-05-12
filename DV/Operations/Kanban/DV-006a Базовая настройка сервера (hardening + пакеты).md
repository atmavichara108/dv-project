---
type: task
id: DV-006a
title: "Базовая настройка сервера: SSH ключи, ufw, пакеты"
status: ready
priority: P0
effort: S
epic: Self-hosted Infra
assignee: max
created: 2026-05-11
tags:
  - infra
  - security
  - p0
---

## Goal
Подготовить Ubuntu 22.04 к продакшен-нагрузке: SSH-ключи, файрвол, базовые пакеты, swap.

## Definition of Done
- SSH вход по паролю отключён, root-логин ограничен
- ufw настроен: 22, 80, 443, 3010, 40000-40100 tcp/udp
- Установлены: build-essential, git, curl, wget, unzip, nginx, snapd, ffmpeg
- Node.js LTS через nvm, pm2 глобально, certbot через snap
- Swap 2 GB (на 5 GB RAM не помешает для mediasoup)
- fail2ban установлен

## Notes
- Команды — в плане, этап 4
- Создать non-root пользователя `dv`, всё дальше делать от него
