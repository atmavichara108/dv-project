

```markdown
# DV Project — Obsidian Vault

Рабочее пространство проекта «Дискуссионные Вечера».
Здесь ведётся вся проектная документация: структура движения, материалы, исследования, операционка.

Репозиторий используется как совместный Obsidian-волт через плагин **Obsidian Git**.

---

## Быстрый старт

### Что нужно заранее

1. **Git** — установленный на компьютере
2. **Obsidian** — https://obsidian.md
3. **Аккаунт GitHub** (для редакторов)

### Установка Git

**Linux (Manjaro / Arch / Ubuntu):**
```bash
# Arch / Manjaro
sudo pacman -S git

# Ubuntu / Debian
sudo apt install git
```

**macOS:**
```bash
brew install git
```
Если нет Homebrew:
```bash
/bin/bash -c "$(curl -fsSL https://raw.githubusercontent.com/Homebrew/install/HEAD/install.sh)"
```

**Windows:**

Скачать установщик: https://git-scm.com/download/win
Установить с настройками по умолчанию. Убедиться, что Git добавлен в PATH (это галочка в установщике).

### Настройка Git (все ОС)

```bash
git config --global user.name "Твоё Имя"
git config --global user.email "твой@email.com"
```

---

## Подключение к волту

### Шаг 1. Клонирование

```bash
cd ~/Documents
git clone https://github.com/atmavichara108/dv-project.git
```

### Шаг 2. Открытие в Obsidian

Obsidian → **Open folder as vault** → выбрать папку `dv-project`.

Плагин Obsidian Git уже в репозитории. Зайти в **Settings → Community plugins** и убедиться, что **Obsidian Git** активен (фиолетовый тумблер).

### Шаг 3. Авторизация GitHub (при первом push)

При первом взаимодействии с GitHub — система попросит авторизацию.

**Linux:** откроется браузер для OAuth, либо запросит Personal Access Token (PAT). Создать: github.com → Settings → Developer Settings → Personal access tokens → Tokens (classic). Scope: `repo`. Скопировать токен, использовать как пароль.

**macOS:** откроется окно Keychain / OAuth в браузере.

**Windows:** откроется Git Credential Manager, авторизация через браузер.

---

## Режим: Редактор (push + pull)

Для тех, кто вносит изменения в документацию.

**Требование:** быть добавленным как collaborator. Скинь свой GitHub username Максу.

### Настройки плагина

Settings → Obsidian Git:

| Параметр | Значение |
|----------|----------|
| Auto backup every X minutes | `5` |
| Auto pull interval (minutes) | `5` |
| Pull updates on startup | `ON` |
| Pull before push | `ON` |
| Sync method | `Merge` |
| Disable push | `OFF` |
| Commit message | `{{hostname}}: {{date}}` |

Формат `{{hostname}}: {{date}}` нужен, чтобы в истории коммитов было видно кто и когда редактировал.

### Рабочий цикл

Всё автоматизировано. Каждые 5 минут плагин:
1. Подтягивает чужие изменения (pull)
2. Коммитит и отправляет твои (push)

Ручные команды (через палитру `Ctrl+P` / `Cmd+P`):
- `Obsidian Git: Commit all changes` — закоммитить сейчас
- `Obsidian Git: Push` — отправить сейчас
- `Obsidian Git: Pull` — получить изменения сейчас

---

## Режим: Наблюдатель (только pull)

Для тех, кто читает документацию без редактирования.

**Не нужно** быть collaborator — репозиторий публичный.

### Настройки плагина

Settings → Obsidian Git:

| Параметр | Значение |
|----------|----------|
| Auto backup every X minutes | `0` (выключено) |
| Auto pull interval (minutes) | `5` |
| Pull updates on startup | `ON` |
| Disable push | `ON` |

### Если случайно что-то отредактировал

Локальные изменения не уйдут в репозиторий (push отключён), но могут мешать при pull. Сброс:

```bash
cd ~/Documents/dv-project
git checkout .
git pull
```

---

## Правила совместной работы

### 1. Один файл — один автор в момент времени
Перед редактированием крупного файла — написать в чат: «я работаю над [название файла]».

### 2. Canvas — только один человек за раз
Canvas в Obsidian — это JSON. Любое одновременное редактирование гарантированно даст конфликт.

### 3. Частые маленькие коммиты лучше редких больших
Для этого и нужен 5-минутный автоинтервал.

### 4. Что делать при конфликте
Открыть конфликтный файл в любом текстовом редакторе. Найти маркеры:
```
<<<<<<< HEAD
твоя версия
=======
чужая версия
>>>>>>> branch
```
Выбрать правильную версию, удалить маркеры, сохранить. Потом:
```bash
git add .
git commit -m "resolve conflict"
git push
```

---

## Структура проекта

```
dv-project/
├── .obsidian/              # Настройки Obsidian и плагинов
├── ДВ/
│   ├── Движение/           # Идеология: манифест, цели, принципы, этика
│   ├── Структура/          # Ячейки, роли, протоколы, S3
│   ├── Сайт/               # DV Hub: архитектура, бэклог, UX
│   ├── Контент/            # Производство контента
│   ├── Сообщество/         # Комьюнити, онбординг, ивенты
│   ├── Исследования/       # Темы, эксперты, синтезы
│   └── Операционка/        # Дорожная карта, спринты, метрики
├── Экосистема.canvas       # Ментальная карта всего проекта
├── Структура.canvas        # Детализация структуры движения
└── .gitignore
```

---

## Связанные ресурсы

- **DV Hub (сайт):** https://dv-hub.pages.dev
- **Репозиторий сайта:** https://github.com/atmavichara108/dv-hub
- **Концептуальные документы:** https://docs.google.com/document/d/1Bi394WoKqc_6egxHPK5qmF4uIPT_49xID5I96kLUfY4
```
