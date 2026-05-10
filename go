
#!/bin/bash
set -e

# ============================================================
# STEP 1: Переименование файлов (сначала файлы, потом папки — снизу вверх)
# ============================================================

# --- ДВ/Движение ---
mv "ДВ/Движение/Основа ДВ.md" "ДВ/Движение/DV Foundation.md" 2>/dev/null || true

# --- ДВ/Структура ---
mv "ДВ/Структура/Consent.md" "ДВ/Структура/Consent.md" 2>/dev/null || true  # уже на английском
mv "ДВ/Структура/S3 интеграция.md" "ДВ/Структура/S3 Integration.md" 2>/dev/null || true
mv "ДВ/Структура/Анатомия ячейки.md" "ДВ/Структура/Cell Anatomy.md" 2>/dev/null || true
mv "ДВ/Структура/Глоссарий S3.md" "ДВ/Структура/S3 Glossary.md" 2>/dev/null || true
mv "ДВ/Структура/Жизненный цикл.md" "ДВ/Структура/Lifecycle.md" 2>/dev/null || true
mv "ДВ/Структура/Карта ячеек.md" "ДВ/Структура/Cells Map.md" 2>/dev/null || true
mv "ДВ/Структура/Логбук.md" "ДВ/Структура/Logbook.md" 2>/dev/null || true
mv "ДВ/Структура/Матрица ролей — шаблон.md" "ДВ/Структура/Roles Matrix — Template.md" 2>/dev/null || true
mv "ДВ/Структура/Онбординг-сценарий.md" "ДВ/Структура/Onboarding Scenario.md" 2>/dev/null || true
mv "ДВ/Структура/Паспорт ячейки — шаблон.md" "ДВ/Структура/Cell Passport — Template.md" 2>/dev/null || true
mv "ДВ/Структура/Протокол обратной связи.md" "ДВ/Структура/Feedback Protocol.md" 2>/dev/null || true
mv "ДВ/Структура/Протокол предложений.md" "ДВ/Структура/Proposal Protocol.md" 2>/dev/null || true
mv "ДВ/Структура/Протокол принятия решений.md" "ДВ/Структура/Decision Protocol.md" 2>/dev/null || true
mv "ДВ/Структура/Роли.md" "ДВ/Структура/Roles.md" 2>/dev/null || true
mv "ДВ/Структура/Ротация.md" "ДВ/Структура/Rotation.md" 2>/dev/null || true
mv "ДВ/Структура/Сеть ячеек.md" "ДВ/Структура/Cells Network.md" 2>/dev/null || true
mv "ДВ/Структура/Шаблон драйвера.md" "ДВ/Структура/Driver Template.md" 2>/dev/null || true
mv "ДВ/Структура/Шаблон соглашения.md" "ДВ/Структура/Agreement Template.md" 2>/dev/null || true

# --- ДВ/Сайт ---
mv "ДВ/Сайт/UX-заметки.md" "ДВ/Сайт/UX Notes.md" 2>/dev/null || true
mv "ДВ/Сайт/Архитектура сайта.md" "ДВ/Сайт/Site Architecture.md" 2>/dev/null || true
mv "ДВ/Сайт/Бэклог фич.md" "ДВ/Сайт/Feature Backlog.md" 2>/dev/null || true
mv "ДВ/Сайт/Деплой и инфра.md" "ДВ/Сайт/Deploy and Infra.md" 2>/dev/null || true
mv "ДВ/Сайт/Интеграция S3.md" "ДВ/Сайт/S3 Integration.md" 2>/dev/null || true
mv "ДВ/Сайт/ключи-пароли.mdenc" "ДВ/Сайт/keys-passwords.mdenc" 2>/dev/null || true

# --- ДВ/Сообщество ---
mv "ДВ/Сообщество/Ивенты.md" "ДВ/Сообщество/Events.md" 2>/dev/null || true
mv "ДВ/Сообщество/Коммуникации.md" "ДВ/Сообщество/Communications.md" 2>/dev/null || true
mv "ДВ/Сообщество/Онбординг.md" "ДВ/Сообщество/Onboarding.md" 2>/dev/null || true
mv "ДВ/Сообщество/Участники.md" "ДВ/Сообщество/Members.md" 2>/dev/null || true

# --- ДВ/Контент ---
mv "ДВ/Контент/Воронка публикации.md" "ДВ/Контент/Publishing Funnel.md" 2>/dev/null || true
mv "ДВ/Контент/Контент-план.md" "ДВ/Контент/Content Plan.md" 2>/dev/null || true
mv "ДВ/Контент/Медиа-инструменты.md" "ДВ/Контент/Media Tools.md" 2>/dev/null || true
mv "ДВ/Контент/Темы в работе.md" "ДВ/Контент/Topics in Progress.md" 2>/dev/null || true
mv "ДВ/Контент/Форматы.md" "ДВ/Контент/Formats.md" 2>/dev/null || true

# --- ДВ/Исследования ---
mv "ДВ/Исследования/Воркфлоу исследования.md" "ДВ/Исследования/Research Workflow.md" 2>/dev/null || true
mv "ДВ/Исследования/Текущие исследования/ИИ и будущее.md" "ДВ/Исследования/Текущие исследования/AI and Future.md" 2>/dev/null || true
mv "ДВ/Исследования/Текущие исследования/Мужское и женское.md" "ДВ/Исследования/Текущие исследования/Masculine and Feminine.md" 2>/dev/null || true
mv "ДВ/Исследования/Текущие исследования/Экология.md" "ДВ/Исследования/Текущие исследования/Ecology.md" 2>/dev/null || true

# --- ДВ/Операционка ---
mv "ДВ/Операционка/Дорожная карта.md" "ДВ/Операционка/Roadmap.md" 2>/dev/null || true
mv "ДВ/Операционка/Метрики.md" "ДВ/Операционка/Metrics.md" 2>/dev/null || true
mv "ДВ/Операционка/Монетизация.md" "ДВ/Операционка/Monetization.md" 2>/dev/null || true
mv "ДВ/Операционка/Риски и блокеры.md" "ДВ/Операционка/Risks and Blockers.md" 2>/dev/null || true
mv "ДВ/Операционка/Спринты.md" "ДВ/Операционка/Sprints.md" 2>/dev/null || true

# ============================================================
# STEP 2: Переименование папок (снизу вверх по вложенности)
# ============================================================

mv "ДВ/Исследования/Текущие исследования" "ДВ/Исследования/Current Research" 2>/dev/null || true
mv "ДВ/Движение" "ДВ/Movement" 2>/dev/null || true
mv "ДВ/Структура" "ДВ/Structure" 2>/dev/null || true
mv "ДВ/Сайт" "ДВ/Site" 2>/dev/null || true
mv "ДВ/Сообщество" "ДВ/Community" 2>/dev/null || true
mv "ДВ/Контент" "ДВ/Content" 2>/dev/null || true
mv "ДВ/Исследования" "ДВ/Research" 2>/dev/null || true
mv "ДВ/Операционка" "ДВ/Operations" 2>/dev/null || true
mv "ДВ" "DV" 2>/dev/null || true

# Canvas файлы
mv "Экосистема.canvas" "Ecosystem.canvas" 2>/dev/null || true
mv "Структура.canvas" "Structure.canvas" 2>/dev/null || true

# ============================================================
# STEP 3: Патч Canvas файлов (замена путей)
# ============================================================

# --- Ecosystem.canvas ---
sed -i \
  -e 's|ДВ/Структура/Карта ячеек.md|DV/Structure/Cells Map.md|g' \
  -e 's|ДВ/Структура/Протокол принятия решений.md|DV/Structure/Decision Protocol.md|g' \
  -e 's|ДВ/Структура/Протокол обратной связи.md|DV/Structure/Feedback Protocol.md|g' \
  -e 's|ДВ/Структура/Шаблон соглашения.md|DV/Structure/Agreement Template.md|g' \
  -e 's|ДВ/Структура/Роли.md|DV/Structure/Roles.md|g' \
  -e 's|ДВ/Структура/Протокол предложений.md|DV/Structure/Proposal Protocol.md|g' \
  -e 's|ДВ/Структура/Шаблон драйвера.md|DV/Structure/Driver Template.md|g' \
  -e 's|ДВ/Структура/Логбук.md|DV/Structure/Logbook.md|g' \
  -e 's|ДВ/Движение/Манифест.md|DV/Movement/DV Foundation.md|g' \
  -e 's|ДВ/Движение/Цели.md|DV/Movement/DV Foundation.md|g' \
  -e 's|ДВ/Движение/Принципы.md|DV/Movement/DV Foundation.md|g' \
  -e 's|ДВ/Движение/Этика и ответственность.md|DV/Movement/DV Foundation.md|g' \
  -e 's|ДВ/Операционка/Дорожная карта.md|DV/Operations/Roadmap.md|g' \
  -e 's|ДВ/Операционка/Спринты.md|DV/Operations/Sprints.md|g' \
  -e 's|ДВ/Операционка/Метрики.md|DV/Operations/Metrics.md|g' \
  -e 's|ДВ/Операционка/Монетизация.md|DV/Operations/Monetization.md|g' \
  -e 's|ДВ/Операционка/Риски и блокеры.md|DV/Operations/Risks and Blockers.md|g' \
  -e 's|ДВ/Сайт/Архитектура сайта.md|DV/Site/Site Architecture.md|g' \
  -e 's|ДВ/Сайт/Интеграция S3.md|DV/Site/S3 Integration.md|g' \
  -e 's|ДВ/Сайт/Деплой и инфра.md|DV/Site/Deploy and Infra.md|g' \
  -e 's|ДВ/Сайт/Бэклог фич.md|DV/Site/Feature Backlog.md|g' \
  -e 's|ДВ/Сайт/UX-заметки.md|DV/Site/UX Notes.md|g' \
  -e 's|ДВ/Сообщество/Участники.md|DV/Community/Members.md|g' \
  -e 's|ДВ/Сообщество/Коммуникации.md|DV/Community/Communications.md|g' \
  -e 's|ДВ/Сообщество/Онбординг.md|DV/Community/Onboarding.md|g' \
  -e 's|ДВ/Сообщество/Ивенты.md|DV/Community/Events.md|g' \
  -e 's|ДВ/Контент/Контент-план.md|DV/Content/Content Plan.md|g' \
  -e 's|ДВ/Контент/Форматы.md|DV/Content/Formats.md|g' \
  -e 's|ДВ/Контент/Темы в работе.md|DV/Content/Topics in Progress.md|g' \
  -e 's|ДВ/Контент/Воронка публикации.md|DV/Content/Publishing Funnel.md|g' \
  -e 's|ДВ/Контент/Медиа-инструменты.md|DV/Content/Media Tools.md|g' \
  -e 's|ДВ/Исследования/Воркфлоу исследования.md|DV/Research/Research Workflow.md|g' \
  -e 's|ДВ/Исследования/Синтезы.md|DV/Research/Research Workflow.md|g' \
  -e 's|ДВ/Исследования/Эксперты.md|DV/Research/Research Workflow.md|g' \
  -e 's|ДВ/Исследования/Текущие исследования/Экология.md|DV/Research/Current Research/Ecology.md|g' \
  -e 's|ДВ/Исследования/Текущие исследования/ИИ и будущее.md|DV/Research/Current Research/AI and Future.md|g' \
  -e 's|ДВ/Исследования/Текущие исследования/Мужское и женское.md|DV/Research/Current Research/Masculine and Feminine.md|g' \
  "Ecosystem.canvas"

# --- Structure.canvas ---
sed -i \
  -e 's|ДВ/Структура/Анатомия ячейки.md|DV/Structure/Cell Anatomy.md|g' \
  -e 's|ДВ/Структура/Жизненный цикл.md|DV/Structure/Lifecycle.md|g' \
  -e 's|ДВ/Структура/Паспорт ячейки — шаблон.md|DV/Structure/Cell Passport — Template.md|g' \
  -e 's|ДВ/Структура/Роли.md|DV/Structure/Roles.md|g' \
  -e 's|ДВ/Структура/Матрица ролей — шаблон.md|DV/Structure/Roles Matrix — Template.md|g' \
  -e 's|ДВ/Структура/Ротация.md|DV/Structure/Rotation.md|g' \
  -e 's|ДВ/Структура/S3 интеграция.md|DV/Structure/S3 Integration.md|g' \
  -e 's|ДВ/Структура/Глоссарий S3.md|DV/Structure/S3 Glossary.md|g' \
  -e 's|ДВ/Структура/Онбординг-сценарий.md|DV/Structure/Onboarding Scenario.md|g' \
  -e 's|ДВ/Структура/Сеть ячеек.md|DV/Structure/Cells Network.md|g' \
  "Structure.canvas"

# ============================================================
# STEP 4: Патч wikilinks внутри .md файлов
# ============================================================

find DV -name "*.md" -exec sed -i \
  -e 's|\[\[Основа ДВ\]\]|[[DV Foundation]]|g' \
  -e 's|\[\[Карта ячеек\]\]|[[Cells Map]]|g' \
  -e 's|\[\[Протокол принятия решений\]\]|[[Decision Protocol]]|g' \
  -e 's|\[\[Протокол предложений\]\]|[[Proposal Protocol]]|g' \
  -e 's|\[\[Протокол обратной связи\]\]|[[Feedback Protocol]]|g' \
  -e 's|\[\[Шаблон драйвера\]\]|[[Driver Template]]|g' \
  -e 's|\[\[Шаблон соглашения\]\]|[[Agreement Template]]|g' \
  -e 's|\[\[Логбук\]\]|[[Logbook]]|g' \
  -e 's|\[\[Роли\]\]|[[Roles]]|g' \
  -e 's|\[\[Ротация\]\]|[[Rotation]]|g' \
  -e 's|\[\[Анатомия ячейки\]\]|[[Cell Anatomy]]|g' \
  -e 's|\[\[Жизненный цикл\]\]|[[Lifecycle]]|g' \
  -e 's|\[\[Сеть ячеек\]\]|[[Cells Network]]|g' \
  -e 's|\[\[S3 интеграция\]\]|[[S3 Integration]]|g' \
  -e 's|\[\[Глоссарий S3\]\]|[[S3 Glossary]]|g' \
  -e 's|\[\[Онбординг-сценарий\]\]|[[Onboarding Scenario]]|g' \
  -e 's|\[\[Паспорт ячейки — шаблон\]\]|[[Cell Passport — Template]]|g' \
  -e 's|\[\[Матрица ролей — шаблон\]\]|[[Roles Matrix — Template]]|g' \
  -e 's|\[\[Архитектура сайта\]\]|[[Site Architecture]]|g' \
  -e 's|\[\[Бэклог фич\]\]|[[Feature Backlog]]|g' \
  -e 's|\[\[Интеграция S3\]\]|[[S3 Integration]]|g' \
  -e 's|\[\[UX-заметки\]\]|[[UX Notes]]|g' \
  -e 's|\[\[Деплой и инфра\]\]|[[Deploy and Infra]]|g' \
  -e 's|\[\[Участники\]\]|[[Members]]|g' \
  -e 's|\[\[Онбординг\]\]|[[Onboarding]]|g' \
  -e 's|\[\[Коммуникации\]\]|[[Communications]]|g' \
  -e 's|\[\[Ивенты\]\]|[[Events]]|g' \
  -e 's|\[\[Контент-план\]\]|[[Content Plan]]|g' \
  -e 's|\[\[Темы в работе\]\]|[[Topics in Progress]]|g' \
  -e 's|\[\[Форматы\]\]|[[Formats]]|g' \
  -e 's|\[\[Воронка публикации\]\]|[[Publishing Funnel]]|g' \
  -e 's|\[\[Медиа-инструменты\]\]|[[Media Tools]]|g' \
  -e 's|\[\[Воркфлоу исследования\]\]|[[Research Workflow]]|g' \
  -e 's|\[\[ИИ и будущее\]\]|[[AI and Future]]|g' \
  -e 's|\[\[Мужское и женское\]\]|[[Masculine and Feminine]]|g' \
  -e 's|\[\[Экология\]\]|[[Ecology]]|g' \
  -e 's|\[\[Дорожная карта\]\]|[[Roadmap]]|g' \
  -e 's|\[\[Спринты\]\]|[[Sprints]]|g' \
  -e 's|\[\[Метрики\]\]|[[Metrics]]|g' \
  -e 's|\[\[Монетизация\]\]|[[Monetization]]|g' \
  -e 's|\[\[Риски и блокеры\]\]|[[Risks and Blockers]]|g' \
  -e 's|\[\[Consent\]\]|[[Consent]]|g' \
  {} +

echo "Done. Структура после переименования:"
find . -not -path './.git/*' -not -path './.obsidian/*' -not -name '.gitignore' | sort
