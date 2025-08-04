package abc_test

import (
	"testing"
	"github.com/Sales-Analysis/abc-helper-lib/abc"
)


func TestABC(t *testing.T) {

	products := []abc.Product{
		{SKU: "T001", Name: "Клавиатура механическая", Quantity: 50, Price: 80.00},
		{SKU: "T002", Name: "Мышь беспроводная", Quantity: 150, Price: 30.00},
		{SKU: "T003", Name: "USB-хаб 7 портов", Quantity: 80, Price: 25.00},
		{SKU: "T004", Name: "Кабель HDMI 2м", Quantity: 300, Price: 10.00},
		{SKU: "T005", Name: "Веб-камера Full HD", Quantity: 40, Price: 60.00},
		{SKU: "T006", Name: "Коврик для мыши XL", Quantity: 250, Price: 15.00},
		{SKU: "T007", Name: "Карта памяти microSD 64GB", Quantity: 200, Price: 12.00},
		{SKU: "T008", Name: "Батарейки AA (упаковка 4 шт)", Quantity: 500, Price: 3.00},
		{SKU: "T009", Name: "Подставка для ноутбука", Quantity: 20, Price: 45.00},
		{SKU: "T010", Name: "Чистящий спрей для экрана", Quantity: 400, Price: 4.00},
	}

	instance := abc.New()
	t.Log(instance.Calculate(products))

}
