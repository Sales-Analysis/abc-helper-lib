## English

### Overview

The `abc` package contains the core logic for ABC analysis:

- Calculates total revenue per product
- Determines percentage and accumulated share
- Assigns ABC categories
- Returns results in a structured format (`ProductResult`)

### Structures

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

### Usage Example

```go
package main

import (
    "fmt"
    "gitlab.com/username/abc-helper-lib/abc"
)

func main() {
    products := []abc.Product{
        {SKU: "001", Name: "Product A", Quantity: 10, Price: 10},
        {SKU: "002", Name: "Product B", Quantity: 5, Price: 20},
        {SKU: "003", Name: "Product C", Quantity: 1, Price: 5},
    }

    analysis := abc.New()
    analysis.Calculate(products)

    for _, r := range analysis.Result {
        fmt.Printf("%+v\n", r)
    }
}
```

### Running Tests

```bash
go test ./abc
```
