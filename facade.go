package analytics

import (
	"github.com/Sales-Analysis/abc-helper-lib/customer"
	"github.com/Sales-Analysis/abc-helper-lib/inventory"
)

// Facade is a thin public entry point over domain-specific analytics packages.
type Facade struct {
	inventory *inventory.Facade
	customer  *customer.Facade
}

func New() *Facade {
	return &Facade{
		inventory: inventory.New(),
		customer:  customer.New(),
	}
}

func (f *Facade) Inventory() *inventory.Facade {
	if f == nil || f.inventory == nil {
		return inventory.New()
	}
	return f.inventory
}

func (f *Facade) Customer() *customer.Facade {
	if f == nil || f.customer == nil {
		return customer.New()
	}
	return f.customer
}
