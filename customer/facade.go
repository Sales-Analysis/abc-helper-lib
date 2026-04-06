package customer

// Facade is a reserved entry point for customer-focused analytics such as RFM and CLV.
type Facade struct{}

func New() *Facade {
	return &Facade{}
}
