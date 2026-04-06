# ABC Helper Library

| English | Русский |
| :---: | :---: |
| [**README.md**](README.md) | [**README.ru.md**](README.ru.md) |

---

## Русский

### Описание

`abc-helper-lib` — библиотека на Go для товарной и клиентской аналитики.

Текущий фундамент включает:

- тонкий корневой facade для orchestration;
- доменное разделение на `inventory` и `customer`;
- реализованные товарные анализы: `ABC`, `XYZ`, `ABC-XYZ`.

### Установка

```bash
go get gitlab.com/username/abc-helper-lib@latest
```

### Структура проекта

```
abc-helper-lib/
│
├── abc/                # ABC-анализ
├── xyz/                # XYZ-анализ
├── abcxyz/             # Комбинированный ABC-XYZ
├── inventory/          # Facade товарной аналитики
├── customer/           # Каркас facade клиентской аналитики
├── facade.go           # Корневой analytics facade
├── go.mod
├── README.md           # Общее описание проекта
└── LICENSE
```

### Документация

Подробное описание пакета abc находится в файле
[**abc/README.ru.md**](abc/README.ru.md)
