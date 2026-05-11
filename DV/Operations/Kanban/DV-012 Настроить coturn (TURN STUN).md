---
type: task
id: DV-012
title: "Настроить coturn (TURN/STUN сервер)"
status: backlog
priority: P1
effort: M
epic: "Conferencing"
sprint: 
assignee: max
created: 2026-05-10
due: 
tags: [video, webrtc, coturn]
---

## Цель
Обеспечить связь WebRTC между участниками за NAT/симметричными firewall-ами. Без TURN до 30% соединений не установятся.

## Definition of Done
- [ ] coturn установлен (контейнер или nativ apt)
- [ ] Открыт UDP 3478, TCP 3478, TLS 5349, диапазон relay-портов
- [ ] Long-term credentials или REST-secret настроены
- [ ] MiroTalk использует наш TURN: проверено на двух устройствах в разных сетях
- [ ] Логи coturn пишутся, не растут бесконечно

## Зависимости
- Зависит от: [[DV-006 Поднять VPS]]
- Блокирует: [[DV-011 Развернуть MiroTalk SFU]]

## Заметки
TURN можно вынести на отдельный VPS, если упрёшься в трафик. На MVP — на том же сервере.
