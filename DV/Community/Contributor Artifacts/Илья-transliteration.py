# Contributor Artifact — не канон-зона и не публикуемый код без отдельного вердикта.
# project: dv-hub
# author: Илья (@Snusmumrik324)
# received: 2026-10-24
# committed: 2026-11-06
# provenance: внешний вклад контрибьютора Ильи; зафиксирован ЦУ/Кораблём 2026-11-06
# status: contributor artifact — сохранён для истории/разбора, не в продакшене

"""
Транслитерация кириллицы в латиницу для доменных имен
"""

# Таблица транслитерации (ГОСТ 7.79-2000 ISO 9:1995 с адаптацией)
TRANSLIT_MAP = {
    'а': 'a', 'б': 'b', 'в': 'v', 'г': 'g', 'д': 'd', 'е': 'e', 'ё': 'yo',
    'ж': 'zh', 'з': 'z', 'и': 'i', 'й': 'j', 'к': 'k', 'л': 'l', 'м': 'm',
    'н': 'n', 'о': 'o', 'п': 'p', 'р': 'r', 'с': 's', 'т': 't', 'у': 'u',
    'ф': 'f', 'х': 'kh', 'ц': 'ts', 'ч': 'ch', 'ш': 'sh', 'щ': 'shch',
    'ъ': '', 'ы': 'y', 'ь': '', 'э': 'eh', 'ю': 'yu', 'я': 'ya'
}


def transliterate(text: str) -> str:
    """
    Транслитерирует кириллицу в латиницу в соответствии с ГОСТ 7.79-2000 (ISO 9:1995) с адаптацией
    
    Правила:
    - Все символы переводятся в нижний регистр
    - Твёрдый знак (ъ) и мягкий знак (ь) полностью исключаются
    - Адаптация ГОСТ: ё→yo, й→j, э→eh
    
    Args:
        text: Текст на русском языке
        
    Returns:
        Транслитерированный текст
    """
    result = []
    for char in text.lower():
        if char in TRANSLIT_MAP:
            # Твёрдый и мягкий знаки уже исключены в словаре (пустая строка)
            result.append(TRANSLIT_MAP[char])
        elif char.isalnum():
            result.append(char)
        elif char in ' -_':
            result.append('-')
    return ''.join(result)


def to_domain_safe(text: str) -> str:
    """
    Преобразует текст в безопасное для домена имя
    
    Args:
        text: Исходный текст
        
    Returns:
        Безопасное доменное имя
    """
    # Транслитерация
    result = transliterate(text)
    
    # Удаление множественных дефисов
    while '--' in result:
        result = result.replace('--', '-')
    
    # Удаление дефисов в начале и конце
    result = result.strip('-')
    
    # Ограничение длины (домены обычно до 63 символов, но лучше короче)
    if len(result) > 50:
        result = result[:50].rstrip('-')
    
    return result




