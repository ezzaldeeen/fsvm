package order

import "github.com/ezzaldeeen/fsvm/eventstore"

// Service ...
type Service struct {
	storage eventstore.EventStore
}

// NewService factory function to instantiate new service
func NewService(storage eventstore.EventStore) *Service {
	return &Service{
		storage: storage,
	}
}

// SelectProduct use case for processing select event
func (s Service) SelectProduct(productID string) error {
	return nil
}

// DepositMoney use case for processing deposit event
func (s Service) DepositMoney(amount uint) error {
	return nil
}

// CancelOrder use case for processing cancel event
func (s Service) CancelOrder(orderID string) error {
	return nil
}
