## Русский

### Обзор
Пакет `abc` содержит основную логику для проведения ABC-анализа:
- Рассчитывает общую выручку по каждому товару
- Определяет процентную и накопленную долю
- Присваивает ABC-категории
- Возвращает результаты в структурированном виде (`ProductResult`)

### Структуры
```go
type Product struct {
    SKU      string
    Name     string
    Quantity int
    Price    float64
}

type ProductResult struct {
    SKU              string
    Name             string
    Quantity         int
    PriceUnit        float64
    PriceTotal       float64
    ShareTotal       float64
    ShareAccumulated float64
    Group            string
}