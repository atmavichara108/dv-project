# DV Project — Obsidian Vault

Рабочее пространство проекта «Дискуссионные Вечера». Здесь ведётся вся проектная документация: структура движения, материалы, исследования, операционка.

Репозиторий используется как совместный Obsidian-волт через плагин **Obsidian Git**.

---

## Быстрый старт

### Что нужно заранее

1. **Git** — установленный на компьютере
2. **Obsidian** — https://obsidian.md
3. **Аккаунт GitHub** (для редакторов)

### Установка Git

**Linux (Manjaro / Arch):** в терминале выполнить `sudo pacman -S git`

**Linux (Ubuntu / Debian):** в терминале выполнить `sudo apt install git`

**macOS:** в терминале выполнить `brew install git` (нужен Homebrew — https://brew.sh)

**Windows:** скачать установщик с https://git-scm.com/download/win и установить с настройками по умолчанию. Убедиться что стоит галочка «Add to PATH».

### Первоначальная настройка Git (все ОС, в терминале)

    git config --global user.name "Твоё Имя"
    git config --global user.email "твой@email.com"

---

## Подключение к волту

### Шаг 1. Клонирование

Открыть терминал и выполнить:

    cd ~/Documents
    git clone https://github.com/atmavichara108/dv-project.git

На Windows — открыть Git Bash (появится после установки Git) и выполнить то же самое.

### Шаг 2. Открытие в Obsidian

Obsidian → **Open folder as vault** → выбрать папку `dv-project`.

Плагин Obsidian Git уже в репозитории. Зайти в **Settings → Community plugins** и убедиться, что **Obsidian Git** активен (фиолетовый тумблер).

### Шаг 3. Авторизация GitHub (при первом push)

При первом взаимодействии с GitHub система попросит авторизацию.

**Linux:** откроется браузер для OAuth, либо запросит Personal Access Token. Создать токен: github.com → Settings → Developer Settings → Personal access tokens → Tokens (classic). Scope: `repo`. Скопировать токен, использовать как пароль в терминале.

**macOS:** откроется окно Keychain или OAuth в браузере.

**Windows:** откроется Git Credential Manager, авторизация через браузер.

---

## Режим: Редактор (push + pull)

Для тех, кто вносит изменения в документацию.

**Требование:** быть добавленным как collaborator в репозиторий. Скинь свой GitHub username Максу — он добавит.

### Настройки плагина

Зайти в Settings → Obsidian Git и выставить:

| Параметр | Значение |
|---|---|
| Auto backup every X minutes | 5 |
| Auto pull interval (minutes) | 5 |
| Pull updates on startup | ON |
| Pull before push | ON |
| Sync method | Merge |
| Disable push | OFF |
| Commit message | `{{hostname}}: {{date}}` |

Формат `{{hostname}}: {{date}}` — чтобы в истории было видно кто и когда редактировал.

### Рабочий цикл

Всё автоматизировано. Каждые 5 минут плагин подтягивает чужие изменения и отправляет твои.

Ручные команды через палитру (Ctrl+P на Linux/Windows, Cmd+P на Mac):

| Команда | Что делает |
|---|---|
| Obsidian Git: Commit all changes | Закоммитить прямо сейчас |
| Obsidian Git: Push | Отправить прямо сейчас |
| Obsidian Git: Pull | Получить изменения прямо сейчас |

---

## Режим: Наблюдатель (только pull)

Для тех, кто читает документацию без редактирования.

**Не нужно** быть collaborator — репозиторий публичный. Просто клонируй и открой.

### Настройки плагина

Зайти в Settings → Obsidian Git и выставить:

| Параметр | Значение |
|---|---|
| Auto backup every X minutes | 0 (выключено) |
| Auto pull interval (minutes) | 5 |
| Pull updates on startup | ON |
| Disable push | ON |

### Если случайно что-то отредактировал

Не страшно — push отключён, ничего не уйдёт в репозиторий. Но локальные изменения могут мешать при следующем pull. Чтобы сбросить, открой терминал:

    cd ~/Documents/dv-project
    git checkout .
    git pull

---

## Правила совместной работы

**Один файл — один автор в момент времени.** Перед редактированием крупного файла — написать в чат: «я работаю над таким-то файлом».

**Canvas — только один человек за раз.** Canvas в Obsidian — это JSON-файл. Любое одновременное редактирование даст конфликт, который сложно разрешить.

**Частые маленькие правки лучше редких больших.** Для этого и нужен 5-минутный автоинтервал.

### Что делать при конфликте

Конфликт — это когда два человека отредактировали один файл одновременно и Git не смог автоматически объединить изменения. Obsidian Git покажет уведомление.

**Самый простой способ — выбрать одну из версий целиком:**

Если хочешь оставить **свою** версию (отбросить чужие изменения в этом файле):

    cd ~/Documents/dv-project
    git checkout --ours "путь/к/файлу.md"
    git add .
    git commit -m "resolve conflict: kept my version"
    git push

Если хочешь оставить **чужую** версию (отбросить свои изменения в этом файле):

    cd ~/Documents/dv-project
    git checkout --theirs "путь/к/файлу.md"
    git add .
    git commit -m "resolve conflict: kept their version"
    git push

**Если нужно объединить вручную** — открой конфликтный файл в любом текстовом редакторе (не в Obsidian). Ты увидишь что-то вроде:

    <<<<<<< HEAD
    Здесь твой текст
    =======
    Здесь чужой текст
    >>>>>>> main

Оставь нужный текст, удали строки с `<<<<<<<`, `=======` и `>>>>>>>`. Сохрани файл. Потом в терминале:

    cd ~/Documents/dv-project
    git add .
    git commit -m "resolve conflict"
    git push

---

## Структура проекта

    dv-project/
    ├── .obsidian/              — настройки Obsidian и плагинов
    ├── ДВ/
    │   ├── Движение/           — идеология: манифест, цели, принципы
    │   ├── Структура/          — ячейки, роли, протоколы, S3
    │   ├── Сайт/               — DV Hub: архитектура, бэклог, UX
    │   ├── Контент/            — производство контента
    │   ├── Сообщество/         — комьюнити, онбординг, ивенты
    │   ├── Исследования/       — темы, эксперты, синтезы
    │   └── Операционка/        — дорожная карта, спринты, метрики
    ├── Экосистема.canvas       — ментальная карта всего проекта
    ├── Структура.canvas        — детализация структуры движения
    └── .gitignore

---

## Связанные ресурсы

**DV Hub (сайт):** https://dv-hub.pages.dev

**Репозиторий сайта:** https://github.com/atmavichara108/dv-hub

**Концептуальные документы:** https://docs.google.com/document/d/1Bi394WoKqc_6egxHPK5qmF4uIPT_49xID5I96kLUfY4
