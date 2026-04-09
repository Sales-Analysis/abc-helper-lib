# Обзор API Перед V1

На этой странице зафиксировано текущее состояние публичного API
`abc-helper-lib` и список решений, которые нужно закрыть до обязательства по
стабильности в `v1.0.0`.

## Текущий Уровень Стабильности

- Модуль все еще находится в фазе `v0.x`.
- Публичные API уже пригодны к использованию и документированы, но пока не
  заморожены.
- В minor-релизах еще допустимы уточнения exported structs, defaults,
  validation rules и output conventions, если это отражено в `CHANGELOG.md`.

## Что Уже Выглядит Зрелым

- Корневая orchestration-точка через `analytics.New()`.
- Доменные facade через `Inventory()` и `Customer()`.
- Stateless package-level entry points через `Analyze(ctx, input)`.
- Явные `Input` и `Output` структуры по всем пакетам.
- Консистентное использование `Items` для товарных входов и `Customers` для клиентских.
- Сохранение `OriginalIndex`, когда результат переставляется относительно входа.
- `Summary` у классификационных анализов и у анализов с явными dataset-level aggregate-метриками.

## Deprecated Surface

- `abc.New()`
- `(*abc.ABC).Calculate(...)`
- `abc.Input.Products` вместо `abc.Input.Items`

Эти API пока оставлены ради обратной совместимости, но новые интеграции не
должны на них опираться.

## Политика Совместимости Внутри v0.x

- Patch-релизы не должны намеренно ломать exported API.
- Minor-релизы пока могут добавлять поля, ужесточать validation или уточнять
  output ordering, если это улучшает консистентность и описано в документации.
- Deprecated API нельзя удалять молча.
- Любое изменение, влияющее на совместимость, должно попадать в
  `CHANGELOG.md` и двуязычную документацию в `docs/`.

## Открытые Решения Перед v1.0.0

1. Решить, будет ли legacy stateful API пакета `abc` удален или сохранен как тонкий compatibility layer.
2. Решить, остается ли `abc.TotalRevenue` долгосрочным convenience-алиасом рядом с `abc.Output.Summary`.
3. Решить, насколько далеко расширять `Summary` на non-classification outputs вроде `EOQ`, `CLV`, `GM/Contribution` и `Service Level`.
4. Решить, достаточно ли generic `"invalid input: ..."` ошибок или библиотеке нужно поднимать typed validation errors в публичный контракт.
5. Закрыть слой robustness:
   fuzz tests, property-based tests и более широкий набор edge-case datasets.
6. Автоматизировать release flow от version tag до GitHub Release notes.
7. Опубликовать явное заявление о backward compatibility для линии после `v1.0.0`.

## Рекомендуемый Чеклист Готовности К v1

- Финализировать судьбу deprecated entry points пакета `abc`.
- Заморозить naming и output conventions по всем пакетам.
- Расширить edge-case coverage за пределы базовых unit tests.
- Держать green golden tests и benchmarks по репрезентативным анализам.
- Подготовить compatibility и migration notes под последние пред-v1 API-изменения.
- Ставить первый `v1`-тег только после закрытия всех review-пунктов выше.
