# DV Project — Obsidian Vault

Рабочее пространство проекта «Дискуссионные Вечера». Здесь ведётся вся проектная документация: структура движения, материалы, исследования, операционка.

Репозиторий используется как совместный Obsidian-волт через плагин Obsidian Git.

---

## Быстрый старт

### Что нужно заранее

1. Git — установленный на компьютере
2. Obsidian — https://obsidian.md
3. Аккаунт GitHub (для редакторов)

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

Открыть терминал (на Windows — Git Bash) и выполнить:

    cd ~/Documents
    git clone https://github.com/atmavichara108/dv-project.git

### Шаг 2. Открытие в Obsidian

Obsidian → Open folder as vault → выбрать папку dv-project.

Плагин Obsidian Git уже в репозитории. Зайти в Settings → Community plugins и убедиться что Obsidian Git активен (фиолетовый тумблер).

### Шаг 3. Авторизация GitHub (при первом push)

При первом взаимодействии с GitHub система попросит авторизацию.

**Linux:** откроется браузер для OAuth, либо запросит Personal Access Token. Создать токен: github.com → Settings → Developer Settings → Personal access tokens → Tokens (classic). Scope: repo. Скопировать токен, использовать как пароль в терминале.

**macOS:** откроется окно Keychain или OAuth в браузере.

**Windows:** откроется Git Credential Manager, авторизация через браузер.

---

## Настройка плагина Obsidian Git

### Для редакторов (push + pull)

Требование: быть добавленным как collaborator в репозиторий. Скинь свой GitHub username Максу — он добавит.

Зайти в Settings → Obsidian Git и выставить:

**Секция Automatic:**

| Параметр | Значение |
|---|---|
| Split timers for automatic commit and sync | OFF |
| Auto commit-and-sync interval (minutes) | 35 |
| Auto commit-and-sync after stopping file edits | OFF |
| Auto commit-and-sync after latest commit | OFF |
| Auto push interval (minutes) | 0 |
| Auto pull interval (minutes) | 35 |
| Auto commit-and-sync only staged files | OFF |
| Specify custom commit message on auto commit-and-sync | OFF |
| Commit message on auto commit-and-sync | {{hostname}}: {{date}} |

**Секция Commit message:**

| Параметр | Значение |
|---|---|
| Commit message on manual commit | {{hostname}}: {{date}} |
| {{date}} placeholder format | YYYY-MM-DD HH:mm:ss |
| {{hostname}} placeholder replacement | твой ник (например: max, dima, anna) |
| List filenames affected by commit in the commit body | ON |

**Секция Commit-and-sync:**

| Параметр | Значение |
|---|---|
| Push on commit-and-sync | ON |
| Pull on commit-and-sync | ON |

**Секция Pull:**

| Параметр | Значение |
|---|---|
| Merge strategy | Merge |
| Merge strategy on conflicts | None (git default) |
| Pull on startup | ON |

**Секция Hunk management:**

| Параметр | Значение |
|---|---|
| Signs | OFF |
| Hunk commands | OFF |
| Status bar with summary of line changes | Disabled |

**Секция Line author information:**

| Параметр | Значение |
|---|---|
| Show commit authoring information next to each line | OFF |

**Секция History view:**

| Параметр | Значение |
|---|---|
| Show Author | Show |
| Show Date | ON |

**Секция Source control view:**

| Параметр | Значение |
|---|---|
| Automatically refresh source control view on file changes | ON |
| Source control view refresh interval | 7000 |

**Секция Miscellaneous:**

| Параметр | Значение |
|---|---|
| Diff view style | Split |
| Disable informative notifications | OFF |
| Disable error notifications | OFF |
| Hide notifications for no changes | OFF |
| Show status bar | ON |
| File menu integration | ON |
| Show branch status bar | ON |
| Show the count of modified files in the status bar | OFF |

**Секция Commit author:**

| Параметр | Значение |
|---|---|
| Author name for commit | оставить пустым (возьмёт из git config) |
| Author email for commit | оставить пустым (возьмёт из git config) |

**Секция Advanced:** оставить всё по умолчанию.

### Для наблюдателей (только pull)

Не нужно быть collaborator — репозиторий публичный. Всё то же самое что выше, но с тремя отличиями:

| Параметр | Значение |
|---|---|
| Auto commit-and-sync interval (minutes) | 0 (выключено) |
| Push on commit-and-sync | OFF |
| Disable on this device | OFF |

Наблюдатель только получает обновления каждые 35 минут и при запуске, но ничего не отправляет.

### Если наблюдатель случайно что-то отредактировал

Не страшно — push отключён, ничего не уйдёт в репозиторий. Но локальные изменения могут мешать при следующем pull. Чтобы сбросить, открыть терминал:

    cd ~/Documents/dv-project
    git checkout .
    git pull

---

## Правила совместной работы

**Один файл — один автор в момент времени.** Перед редактированием крупного файла — написать в чат: «я работаю над таким-то файлом».

**Canvas — только один человек за раз.** Canvas в Obsidian — это JSON-файл. Любое одновременное редактирование даст конфликт который сложно разрешить.

**Частые маленькие правки лучше редких больших.** Для этого и нужен 5-минутный автоинтервал.

### Что делать при конфликте

Конфликт — это когда два человека отредактировали один файл одновременно и Git не смог автоматически объединить изменения. Obsidian Git покажет уведомление.

**Самый простой способ — выбрать чью-то версию целиком.**

Чтобы оставить свою версию (отбросить чужие изменения в этом файле), открыть терминал:

    cd ~/Documents/dv-project
    git checkout --ours "путь/к/файлу.md"
    git add .
    git commit -m "resolve conflict: kept my version"
    git push

Чтобы оставить чужую версию (отбросить свои изменения в этом файле):

    cd ~/Documents/dv-project
    git checkout --theirs "путь/к/файлу.md"
    git add .
    git commit -m "resolve conflict: kept their version"
    git push

Если нужно объединить обе версии вручную — открыть конфликтный файл в любом текстовом редакторе (не в Obsidian). Внутри будет что-то такое:

    <<<<<<< HEAD
    Здесь твой текст
    =======
    Здесь чужой текст
    >>>>>>> main

Нужно оставить правильный текст и удалить три служебные строки (с угловыми скобками и знаками равенства). После этого сохранить файл и в терминале:

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
