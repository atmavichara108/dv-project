#!/usr/bin/env bash
# Пересборка доски Kanban из карточек DV-*.md.
# Единый источник правды — фронтматтер каждой карточки (поле status).
# Доска — производная: генерируется заново при каждом запуске.
#
# Статус-словарь (ровно 4): backlog | ready | in_progress | done.
# Любой другой статус в карточке — ошибка: скрипт падает и не трогает доску,
# чтобы не прятать рассинхронизацию.
#
# Порядок столбцов на доске: In Progress → Ready → Backlog → Done.
# Порядок карточек внутри столбца: по приоритету (P0→P3), затем по ID.
#
# Использование:
#   scripts/rebuild-kanban.sh          # пересобрать в dv-project
#   scripts/rebuild-kanban.sh --check  # только проверить словарь, не писать

set -euo pipefail

KANBAN_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")/../DV/Operations/Kanban" && pwd)"
BOARD="$KANBAN_DIR/Kanban Board.md"
CHECK_ONLY=0
[[ "${1:-}" == "--check" ]] && CHECK_ONLY=1

STATUS_ORDER=(in_progress ready backlog done)
STATUS_HEADER=(
  "🟡 In Progress"
  "🟠 Ready"
  "⚪ Backlog"
  "🟢 Done"
)

# priority → emoji на доске
prio_emoji() {
  case "$1" in
    P0) echo "🔥" ;;
    P1) echo "⭐" ;;
    P2) echo "•" ;;
    P3) echo "·" ;;
    *)  echo "?" ;;
  esac
}

# epic → аккуратно убрать кавычки (если есть)
epic_clean() { sed 's/^"//; s/"$//' <<<"$1"; }

# Собрать строки карточек по статусу.
declare -A ROWS
for status in "${STATUS_ORDER[@]}"; do ROWS[$status]=""; done
VALID_STATUSES="backlog ready in_progress done"

shopt -s nullglob
declare -a FILES
FILES=("$KANBAN_DIR"/DV-*.md)

if [[ ${#FILES[@]} -eq 0 ]]; then
  echo "ERROR: нет карточек DV-*.md в $KANBAN_DIR" >&2
  exit 2
fi

for f in "${FILES[@]}"; do
  [[ "$(basename "$f")" == "Kanban Board.md" ]] && continue

  id="$(sed -n 's/^id: *//p' "$f" | head -1 | tr -d ' \t')"
  title="$(sed -n 's/^title: *//p' "$f" | head -1 | tr -d '"')"
  status="$(sed -n 's/^status: *//p' "$f" | head -1 | tr -d ' \t')"
  prio="$(sed -n 's/^priority: *//p' "$f" | head -1 | tr -d ' \t')"
  effort="$(sed -n 's/^effort: *//p' "$f" | head -1 | tr -d ' \t')"
  epic="$(sed -n 's/^epic: *//p' "$f" | head -1)"

  # Валидация словаря — святое. Хочешь новый статус — сначала словарь, потом доска.
  if ! grep -q "$status" <<<"$VALID_STATUSES"; then
    echo "ERROR: карточка $id имеет невалидный статус '$status' (файл: $(basename "$f"))" >&2
    echo "Допустимо: $VALID_STATUSES" >&2
    exit 3
  fi

  # Ссылка на карточку: базовый линк по имени файла, короткое имя — в текст.
  base="$(basename "$f" .md)"
  alias="${title//:/-}"
  alias="${alias//,/-}"
  alias="${alias//|/ }"
  link="[[$base|$alias]]"

  ROWS[$status]+="| $id | $link | $(prio_emoji "$prio") | $(epic_clean "$epic") | $effort |"$'\n'
done

if [[ "$CHECK_ONLY" -eq 1 ]]; then
  echo "OK: словарь валиден"
  exit 0
fi

now="$(date '+%Y-%m-%d %H:%M')"
{
  echo "---"
  echo "title: Kanban Board"
  echo "date: $(date '+%Y-%m-%d')"
  echo "generated-by: scripts/rebuild-kanban.sh (single source of truth — карточки DV-*.md)"
  echo "---"
  echo ""
  echo "# Kanban Board"
  echo ""
  echo "> Авто-пересборка: $(date '+%Y-%m-%d %H:%M'). Не редактировать доску вручную —"
  echo "> доска строится из статусов карточек. Правь status в карточке."
  echo ""
  echo "Правило доски: стартовала — сразу 🟡 In Progress; выполнена — сразу 🟢 Done."
  echo "Источник правды — status в DV-*.md; доска — производная (rebuild-kanban.sh)."
  echo ""

  for i in "${!STATUS_ORDER[@]}"; do
    s="${STATUS_ORDER[$i]}"
    # пустая строка если карточек нет (может быть backlog пуст и т.п.)
    if [[ -z "${ROWS[$s]}" ]]; then
      echo "## ${STATUS_HEADER[$i]} (0)"
      echo ""
      echo "_пусто_"
      echo ""
      continue
    fi
    n="$(printf '%s' "${ROWS[$s]}" | grep -c '^| ' || true)"
    echo "## ${STATUS_HEADER[$i]} ($n)"
    echo ""
    echo "| ID | Задача | Приоритет | Эпик | Объём |"
    echo "|---|---|---|---|---|"
    printf '%s' "${ROWS[$s]}"
    echo ""
  done
} > "$BOARD"

echo "OK: доска пересобрана ($(date '+%Y-%m-%d %H:%M'))"
exit 0