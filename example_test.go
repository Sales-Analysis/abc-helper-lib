package analytics_test

import (
	"context"
	"fmt"
	"time"

	analytics "github.com/Sales-Analysis/abc-helper-lib"
	"github.com/Sales-Analysis/abc-helper-lib/abc"
	"github.com/Sales-Analysis/abc-helper-lib/churn"
	"github.com/Sales-Analysis/abc-helper-lib/clv"
	"github.com/Sales-Analysis/abc-helper-lib/eoq"
	"github.com/Sales-Analysis/abc-helper-lib/pareto"
	"github.com/Sales-Analysis/abc-helper-lib/rfm"
	"github.com/Sales-Analysis/abc-helper-lib/xyz"
)

func ExampleNew_inventoryABC() {
	ctx := context.Background()

	out, err := analytics.New().Inventory().ABC(ctx, abc.Input{
		Items: []abc.Item{
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

func ExampleNew_inventoryXYZ() {
	ctx := context.Background()

	out, err := analytics.New().Inventory().XYZ(ctx, xyz.Input{
		Items: []xyz.Item{
			{SKU: "A", Name: "Item A", Demands: []float64{10, 10, 10}},
		},
	})
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("%s %s %d\n", out.Results[0].SKU, out.Results[0].Group, out.Summary.XCount)

	// Output:
	// A X 1
}

func ExampleNew_inventoryPareto() {
	ctx := context.Background()

	out, err := analytics.New().Inventory().Pareto(ctx, pareto.Input{
		Items: []pareto.Item{
			{SKU: "A", Name: "Item A", Value: 80},
			{SKU: "B", Name: "Item B", Value: 20},
		},
	})
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("%d %.0f %t\n", out.Summary.TopItemCount, out.Summary.TopItemValueShare, out.Summary.ParetoPrincipleMet)

	// Output:
	// 1 80 true
}

func ExampleNew_inventoryEOQ() {
	ctx := context.Background()

	out, err := analytics.New().Inventory().EOQ(ctx, eoq.Input{
		Items: []eoq.Item{
			{SKU: "A", Name: "Item A", AnnualDemand: 1200, OrderingCost: 50, HoldingCost: 2},
		},
	})
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("%s %.2f %.2f\n", out.Results[0].SKU, out.Results[0].OptimalQuantity, out.Results[0].OrdersPerYear)

	// Output:
	// A 244.95 4.90
}

func ExampleNew_customerCLV() {
	ctx := context.Background()

	out, err := analytics.New().Customer().CLV(ctx, clv.Input{
		Customers: []clv.Customer{
			{
				CustomerID:      "C1",
				Name:            "Customer 1",
				Revenue:         1200,
				Orders:          6,
				PeriodsObserved: 3,
				GrossMarginRate: 0.4,
				RetentionRate:   0.8,
				DiscountRate:    0.1,
				AcquisitionCost: 50,
			},
		},
	})
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("%s %.2f\n", out.Results[0].CustomerID, out.Results[0].PredictedCLV)

	// Output:
	// C1 376.67
}

func ExampleNew_customerChurn() {
	ctx := context.Background()
	analysisTime := time.Date(2026, time.January, 31, 0, 0, 0, 0, time.UTC)

	out, err := analytics.New().Customer().Churn(ctx, churn.Input{
		AnalysisTime: analysisTime,
		Customers: []churn.Customer{
			{
				CustomerID:  "C1",
				Name:        "Customer 1",
				LastOrderAt: analysisTime.AddDate(0, 0, -10),
			},
		},
	})
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("%s %s %.0f\n", out.Results[0].CustomerID, out.Results[0].Status, out.Summary.RetentionRate)

	// Output:
	// C1 Active 100
}
