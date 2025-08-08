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