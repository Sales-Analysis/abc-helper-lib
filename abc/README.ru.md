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
```

### Пример использования

```go
package main

import (
    "fmt"
    "gitlab.com/username/abc-helper-lib/abc"
)

func main() {
    products := []abc.Product{
        {SKU: "001", Name: "Товар A", Quantity: 10, Price: 10},
        {SKU: "002", Name: "Товар B", Quantity: 5, Price: 20},
        {SKU: "003", Name: "Товар C", Quantity: 1, Price: 5},
    }

    analysis := abc.New()
    analysis.Calculate(products)

    for _, r := range analysis.Result {
        fmt.Printf("%+v\n", r)
    }
}
```

### Запуск тестов

```bash
go test ./abc
```
