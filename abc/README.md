## English

### Overview

The `abc` package contains the core logic for ABC analysis:

- Calculates total revenue per product
- Determines percentage and accumulated share
- Assigns ABC categories
- Returns results in a structured format (`ProductResult`)
- Exposes a stateless API for orchestration and composite analyses

### Structures

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

### Usage Example

```go
package main

import (
    "fmt"
    "github.com/Sales-Analysis/abc-helper-lib/abc"
)

func main() {
    products := []abc.Product{
        {SKU: "001", Name: "Product A", Quantity: 10, Price: 10},
        {SKU: "002", Name: "Product B", Quantity: 5, Price: 20},
        {SKU: "003", Name: "Product C", Quantity: 1, Price: 5},
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

The old stateful API based on `abc.New()` and `(*ABC).Calculate(...)` is still available
for backward compatibility, but it is deprecated in favor of `Analyze(...)`.

### Running Tests

```bash
go test ./abc
```
