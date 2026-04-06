## Русский

### Обзор

Пакет `abc` содержит основную логику для проведения ABC-анализа:

- Рассчитывает общую выручку по каждому товару
- Определяет процентную и накопленную долю
- Присваивает ABC-категории
- Возвращает результаты в структурированном виде (`ProductResult`)
- Предоставляет stateless API для orchestration и составных анализов

### Структуры

```go
type Input struct {
    Products   []Product
    Thresholds Thresholds
}

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

type Output struct {
    Results      []ProductResult
    TotalRevenue float64
}
```

### Пример использования

```go
package main

import (
    "fmt"
    "github.com/Sales-Analysis/abc-helper-lib/abc"
)

func main() {
    products := []abc.Product{
        {SKU: "001", Name: "Товар A", Quantity: 10, Price: 10},
        {SKU: "002", Name: "Товар B", Quantity: 5, Price: 20},
        {SKU: "003", Name: "Товар C", Quantity: 1, Price: 5},
    }

    output, err := abc.Analyze(nil, abc.Input{Products: products})
    if err != nil {
        panic(err)
    }

    for _, r := range output.Results {
        fmt.Printf("%+v\n", r)
    }
}
```

### Legacy API

Старый stateful API на базе `abc.New()` и `(*ABC).Calculate(...)` оставлен
для обратной совместимости, но помечен deprecated в пользу `Analyze(...)`.

### Запуск тестов

```bash
go test ./abc
```
