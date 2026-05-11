---
type: task
id: DV-025
title: "Twake Drive для тяжёлых медиа (видео, PDF)"
status: backlog
priority: P3
effort: L
epic: "Phase 2"
sprint: 
assignee: max
created: 2026-05-10
due: 
tags: [storage, infra, twake]
---

## Цель
Не пытаться хранить видеофайлы и тяжёлые PDF в основной БД. Дать командам файловое хранилище с правами доступа.

## Definition of Done
- [ ] Twake Drive развёрнут в docker-compose
- [ ] SSO с DV Hub (или хотя бы общий пользователь admin)
- [ ] В таблице `materials` поле `external_storage_url` для ссылки на файл в Twake
- [ ] UI материала умеет показывать превью PDF и проигрывать видео из Twake
- [ ] Бэкап Twake включён в общую процедуру (см. DV-009)

## Зависимости
- Зависит от: [[DV-009 Бэкапы БД]]

## Заметки
Триггер для перехода: суммарный объём медиа в проекте перевалит за 5 GB или появится первая жалоба на лимиты.
