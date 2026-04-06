package analytics_test

import (
	"context"
	"fmt"
	"time"

	analytics "github.com/Sales-Analysis/abc-helper-lib"
	"github.com/Sales-Analysis/abc-helper-lib/abc"
	"github.com/Sales-Analysis/abc-helper-lib/rfm"
)

func ExampleNew_inventoryABC() {
	ctx := context.Background()

	out, err := analytics.New().Inventory().ABC(ctx, abc.Input{
		Products: []abc.Product{
			{SKU: "A", Name: "Item A", Quantity: 10, Price: 100},
			{SKU: "B", Name: "Item B", Quantity: 5, Price: 50},
		},
	})
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("%s %s %.0f\n", out.Results[0].SKU, out.Results[0].Group, out.TotalRevenue)

	// Output:
	// A A 1250
}

func ExampleNew_customerRFM() {
	ctx := context.Background()
	analysisTime := time.Date(2026, time.January, 31, 0, 0, 0, 0, time.UTC)

	out, err := analytics.New().Customer().RFM(ctx, rfm.Input{
		AnalysisTime: analysisTime,
		Customers: []rfm.Customer{
			{
				CustomerID:    "C1",
				Name:          "Alice",
				LastOrderAt:   analysisTime.AddDate(0, 0, -5),
				Orders:        10,
				MonetaryValue: 500,
			},
			{
				CustomerID:    "C2",
				Name:          "Bob",
				LastOrderAt:   analysisTime.AddDate(0, 0, -40),
				Orders:        2,
				MonetaryValue: 100,
			},
		},
	})
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("%s %s\n", out.Results[0].CustomerID, out.Results[0].Segment)

	// Output:
	// C1 Champions
}
