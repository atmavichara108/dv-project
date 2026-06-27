---
type: spec
id: DV-009
title: "Настроить ежедневные бэкапы БД с offsite-копией"
status: approved
---

## Product-контекст

DV Hub работает на единственном VPS (Fornex, 40 GB NVMe). SQLite-файл БД — единственное хранилище данных. При сбое диска, случайном удалении или ошибке миграции данные будут потеряны безвозвратно.

Нужен автоматический бэкап с offsite-копией на локальный компьютер через Tailscale VPN.

## Архитектура

### Схема

```
VPS (Fornex, Germany)
  └─ cron: backup.sh (каждое воскресенье в 3:00 МСК)
       ├─ sqlite3 .dump → gzip → dv-hub-db-{timestamp}.sql.gz
       ├─ tar .env → dv-hub-env-{timestamp}.tar.gz
       ├─ tar nginx configs → nginx-conf-{timestamp}.tar.gz
       ├─ объединение → dv-hub-full-{timestamp}.tar.gz
       ├─ rsync через Tailscale → локальный комп
       └─ cleanup: хранить 4 последних копии локально

Локальный комп (домашний ПК)
  └─ /mnt/backups/dv-hub/
       └─ dv-hub-full-*.tar.gz (все копии, чистка руками)
```

### Компоненты

| Компонент | Путь | Назначение |
|-----------|------|------------|
| Скрипт бэкапа | `/opt/dv-hub/scripts/backup.sh` | Запускается cron'ом, делает dump + rsync |
| Скрипт восстановления | `/opt/dv-hub/scripts/restore-backup.sh` | Запускается руками, восстанавливает из архива |
| Локальное хранилище | `/opt/dv-hub/backups/` | Временное хранение на VPS (4 копии) |
| Offsite-хранилище | `/mnt/backups/dv-hub/` на локальном компе | Постоянное хранение |
| Лог | `/opt/dv-hub/backups/backup.log` | Лог всех операций бэкапа |

### Зависимости

- **Tailscale VPN** — для безопасного rsync между VPS и локальным компом (устанавливается руками)
- **SSH-ключ** — публичный ключ VPS добавляется в `~rudra/.ssh/authorized_keys` на локальном компе
- **rsync** — уже установлен на Ubuntu 24.04 по умолчанию
- **sqlite3** — уже установлен на VPS

### Конфигурация

Параметры задаются через переменные окружения в `/opt/dv-hub/.env`:

```bash
TAILSCALE_IP=100.x.x.x       # Tailscale IP локального компа
REMOTE_USER=rudra             # пользователь на локальном компе
REMOTE_PATH=/mnt/backups/dv-hub/  # путь назначения
LOCAL_RETENTION=4             # сколько копий хранить на VPS
```

## Изменяемые файлы

1. `scripts/backup.sh` — новый скрипт бэкапа
2. `scripts/restore-backup.sh` — новый скрипт восстановления
3. `docs/infra-runbook.md` — раздел 6 (Бэкапы и восстановление)

## НЕ изменяемые

- `src/` — никаких изменений в коде приложения
- `migrations/` — не затрагиваем
- `public/` — не затрагиваем

## Post-setup checklist (выполняется руками после деплоя)

- [ ] Tailscale установлен на VPS (`tailscale status`)
- [ ] Tailscale установлен на локальном компе
- [ ] SSH-ключ VPS добавлен в `~rudra/.ssh/authorized_keys` локального компа
- [ ] `TAILSCALE_IP=...` добавлен в `/opt/dv-hub/.env`
- [ ] Директория `/mnt/backups/dv-hub/` создана на локальном компе
- [ ] Cron установлен: `crontab -e` → `0 3 * * 0 /opt/dv-hub/scripts/backup.sh`
- [ ] Ручной тест: `bash /opt/dv-hub/scripts/backup.sh` — архив доезжает до локального компа

## Definition of Done

- [ ] `scripts/backup.sh` создан, `chmod +x`, проверен shell syntax
- [ ] `scripts/restore-backup.sh` создан, `chmod +x`
- [ ] `docs/infra-runbook.md` раздел 6 обновлён (метод, расписание, восстановление, checklist)
- [ ] `context/DV/Operations/Specs/DV-009-spec.md` создан
