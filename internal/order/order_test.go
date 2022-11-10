package entity

import "testing"

func TestGivenAnEmptyID_WhenCreateANewOrder_ThenShouldReceiveAnError(t *testing.T) {
	order := Order{}
	isValid := order.IsValid()
	if isValid == nil {
		t.Errorf("Should have returned 'invalid id' error, got nil")
	}
	if isValid.Error() != "invalid id" {
		t.Errorf("Should have returned 'invalid id' error, got %s", order.IsValid().Error())
	}
}
func TestGivenAnEmptyPrice_WhenCreateANewOrder_ThenShouldReceiveAnError(t *testing.T) {
	order := Order{ID: "123"}
	isValid := order.IsValid()
	if isValid == nil {
		t.Errorf("Should have returned 'invalid price' error, got nil")
	}
	if isValid.Error() != "invalid price" {
		t.Errorf("Should have returned 'invalid price' error, got %s", order.IsValid().Error())
	}
}
func TestGivenAnEmptyTax_WhenCreateANewOrder_ThenShouldReceiveAnError(t *testing.T) {
	order := Order{ID: "123", Price: 123}
	isValid := order.IsValid()
	if isValid == nil {
		t.Errorf("Should have returned 'invalid tax' error, got nil")
	}
	if isValid.Error() != "invalid tax" {
		t.Errorf("Should have returned 'invalid tax' error, got %s", order.IsValid().Error())
	}
}

func TestGivenValidParams_WhenNewOrderIsCalled_ThenShouldReceiveOrderWithAllParams(t *testing.T) {
	order := Order{
		ID:    "123",
		Price: 10.5,
		Tax:   0.5,
	}

	if order.ID != "123" {
		t.Errorf("Should have the order ID 123")
	}
	if order.Price != 10.5 {
		t.Errorf("Should have the order price set as 10.5")
	}
	if order.Tax != 0.5 {
		t.Errorf("Should have the order tac set as 0.5")
	}

	isValid := order.IsValid()

	if isValid != nil {
		t.Errorf("Should not have returned an error, got %s", isValid.Error())
	}
}

func TestGivenAPriceAndTax_WhenCalculatePriceIsCalled_ThenShouldSetFinalPrice(t *testing.T) {
	order, err := NewOrder("123", 10.0, 2.0)

	if err != nil {
		t.Errorf("Should not thrown an error, got %s", err.Error())
	}

	err = order.CalculateFinalPrice()

	if err != nil {
		t.Errorf("Should not thrown an error, got %s", err.Error())
	}

	if order.FinalPrice != 12.0 {
		t.Errorf("Expected final price to be 12.0, got %b", order.FinalPrice)
	}
}
