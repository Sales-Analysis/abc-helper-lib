package validation

import (
	"fmt"
	"math"
)

func Invalidf(format string, args ...any) error {
	return fmt.Errorf("invalid input: "+format, args...)
}

func RequirePercent(field string, value float64) error {
	if err := requireFiniteFloat(field, value); err != nil {
		return err
	}
	if value <= 0 || value > 100 {
		return Invalidf("%s must be in (0, 100]", field)
	}
	return nil
}

func RequirePercentAllowZero(field string, value float64) error {
	if err := requireFiniteFloat(field, value); err != nil {
		return err
	}
	if value < 0 || value > 100 {
		return Invalidf("%s must be in [0, 100]", field)
	}
	return nil
}

func RequirePositiveFloat(field string, value float64) error {
	if err := requireFiniteFloat(field, value); err != nil {
		return err
	}
	if value <= 0 {
		return Invalidf("%s must be greater than 0", field)
	}
	return nil
}

func RequireNonNegativeFloat(field string, value float64) error {
	if err := requireFiniteFloat(field, value); err != nil {
		return err
	}
	if value < 0 {
		return Invalidf("%s must be greater than or equal to 0", field)
	}
	return nil
}

func RequirePositiveInt(field string, value int) error {
	if value <= 0 {
		return Invalidf("%s must be greater than 0", field)
	}
	return nil
}

func RequireNonNegativeInt(field string, value int) error {
	if value < 0 {
		return Invalidf("%s must be greater than or equal to 0", field)
	}
	return nil
}

func RequireGreaterFloat(field string, value float64, otherField string, otherValue float64) error {
	if err := requireFiniteFloat(field, value); err != nil {
		return err
	}
	if err := requireFiniteFloat(otherField, otherValue); err != nil {
		return err
	}
	if value <= otherValue {
		return Invalidf("%s must be greater than %s", field, otherField)
	}
	return nil
}

func RequireGreaterInt(field string, value int, otherField string, otherValue int) error {
	if value <= otherValue {
		return Invalidf("%s must be greater than %s", field, otherField)
	}
	return nil
}

func RequireRatio(field string, value float64) error {
	if err := requireFiniteFloat(field, value); err != nil {
		return err
	}
	if value < 0 {
		return Invalidf("%s must be in [0, 1] or [0, 100]", field)
	}
	if value > 100 {
		return Invalidf("%s must be in [0, 1] or [0, 100]", field)
	}
	return nil
}

func requireFiniteFloat(field string, value float64) error {
	if math.IsNaN(value) || math.IsInf(value, 0) {
		return Invalidf("%s must be finite", field)
	}
	return nil
}
