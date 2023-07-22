package order

import "errors"

// Select ...
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

// Deposit ...
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

// Cancel ...
func (o *Order) Cancel() {
	o.raise(Canceled{
		orderID: o.id,
	})
}
