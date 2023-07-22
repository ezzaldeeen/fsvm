package order

import "errors"

// Select command for validating, and raising Selected event
func (o *Order) Select(productID string) error {
	if o.state == CANCELED { // todo: is it correct?
		return errors.New("unable to select since the order canceled")
	}
	o.raise(Selected{
		orderID:   o.id,
		productID: productID,
	})
	return nil
}

// Deposit command for validating, and raising Deposited event
func (o *Order) Deposit(amount uint) error {
	if o.state == CANCELED {
		return errors.New("unable to deposit, select first") // todo: change this dummy logic
	}
	o.raise(Deposited{
		orderID: o.id,
		amount:  amount,
	})
	return nil
}

// Cancel command for validating, and raising Canceled event
func (o *Order) Cancel() {
	o.raise(Canceled{
		orderID: o.id,
	})
}
