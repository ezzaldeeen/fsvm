package engine

import (
	"github.com/ezzaldeeen/fsvm/eventstore"
	"reflect"
	"testing"
)

func TestEngine_Cancel(t *testing.T) {
	type fields struct {
		selecting    State
		depositing   State
		dispensing   State
		currentState State
		balance      float64
		basket       []Item
		inventory    map[string]Item
		history      map[string]eventstore.Events
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		{
			name:    "invalid cancellation in selecting state",
			fields:  fields{currentState: &Selecting{}},
			wantErr: true,
		},
		{
			name:    "invalid cancellation in dispensing state",
			fields:  fields{currentState: &Dispensing{}},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := &Engine{
				selecting:    tt.fields.selecting,
				depositing:   tt.fields.depositing,
				dispensing:   tt.fields.dispensing,
				currentState: tt.fields.currentState,
				balance:      tt.fields.balance,
				basket:       tt.fields.basket,
				inventory:    tt.fields.inventory,
				history:      tt.fields.history,
			}
			if err := e.Cancel(); (err != nil) != tt.wantErr {
				t.Errorf("Cancel() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestEngine_CurrentState(t *testing.T) {
	type fields struct {
		selecting    State
		depositing   State
		dispensing   State
		currentState State
		balance      float64
		basket       []Item
		inventory    map[string]Item
		history      map[string]eventstore.Events
	}
	tests := []struct {
		name   string
		fields fields
		want   State
	}{
		{
			name:   "getting current state",
			fields: fields{currentState: &Dispensing{}},
			want:   &Dispensing{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := &Engine{
				selecting:    tt.fields.selecting,
				depositing:   tt.fields.depositing,
				dispensing:   tt.fields.dispensing,
				currentState: tt.fields.currentState,
				balance:      tt.fields.balance,
				basket:       tt.fields.basket,
				inventory:    tt.fields.inventory,
				history:      tt.fields.history,
			}
			if got := e.CurrentState(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CurrentState() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestEngine_Deposit(t *testing.T) {
	type fields struct {
		selecting    State
		depositing   State
		dispensing   State
		currentState State
		balance      float64
		basket       []Item
		inventory    map[string]Item
		history      map[string]eventstore.Events
	}
	type args struct {
		amount float64
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name:    "invalid deposit from selecting state",
			fields:  fields{currentState: &Selecting{}},
			wantErr: true,
		},
		{
			name:    "invalid deposit from dispensing state",
			fields:  fields{currentState: &Dispensing{}},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := &Engine{
				selecting:    tt.fields.selecting,
				depositing:   tt.fields.depositing,
				dispensing:   tt.fields.dispensing,
				currentState: tt.fields.currentState,
				balance:      tt.fields.balance,
				basket:       tt.fields.basket,
				inventory:    tt.fields.inventory,
				history:      tt.fields.history,
			}
			if err := e.Deposit(tt.args.amount); (err != nil) != tt.wantErr {
				t.Errorf("Deposit() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestEngine_Dispense(t *testing.T) {
	type fields struct {
		selecting    State
		depositing   State
		dispensing   State
		currentState State
		balance      float64
		basket       []Item
		inventory    map[string]Item
		history      map[string]eventstore.Events
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		{
			name:    "invalid dispense from selecting state",
			fields:  fields{currentState: &Selecting{}},
			wantErr: true,
		},
		{
			name:    "invalid dispense from depositing state",
			fields:  fields{currentState: &Depositing{}},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := &Engine{
				selecting:    tt.fields.selecting,
				depositing:   tt.fields.depositing,
				dispensing:   tt.fields.dispensing,
				currentState: tt.fields.currentState,
				balance:      tt.fields.balance,
				basket:       tt.fields.basket,
				inventory:    tt.fields.inventory,
				history:      tt.fields.history,
			}
			if err := e.Dispense(); (err != nil) != tt.wantErr {
				t.Errorf("Dispense() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestEngine_Select(t *testing.T) {
	type fields struct {
		selecting    State
		depositing   State
		dispensing   State
		currentState State
		balance      float64
		basket       []Item
		inventory    map[string]Item
		history      map[string]eventstore.Events
	}
	type args struct {
		itemID string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name:    "invalid select from dispensing state",
			fields:  fields{currentState: &Dispensing{}},
			wantErr: true,
		},
		{
			name:    "invalid select from depositing state",
			fields:  fields{currentState: &Depositing{}},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := &Engine{
				selecting:    tt.fields.selecting,
				depositing:   tt.fields.depositing,
				dispensing:   tt.fields.dispensing,
				currentState: tt.fields.currentState,
				balance:      tt.fields.balance,
				basket:       tt.fields.basket,
				inventory:    tt.fields.inventory,
				history:      tt.fields.history,
			}
			if err := e.Select(tt.args.itemID); (err != nil) != tt.wantErr {
				t.Errorf("Select() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestEngine_addToBalance(t *testing.T) {
	type fields struct {
		selecting    State
		depositing   State
		dispensing   State
		currentState State
		balance      float64
		basket       []Item
		inventory    map[string]Item
		history      map[string]eventstore.Events
	}
	type args struct {
		amount float64
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name:    "invalid amount - below zero",
			args:    args{amount: -1},
			wantErr: true,
		},
		{
			name:    "invalid amount - equal zero",
			args:    args{amount: 0},
			wantErr: true,
		},
		{
			name:    "valid amount - above zero",
			args:    args{amount: 2},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := &Engine{
				selecting:    tt.fields.selecting,
				depositing:   tt.fields.depositing,
				dispensing:   tt.fields.dispensing,
				currentState: tt.fields.currentState,
				balance:      tt.fields.balance,
				basket:       tt.fields.basket,
				inventory:    tt.fields.inventory,
				history:      tt.fields.history,
			}
			if err := e.addToBalance(tt.args.amount); (err != nil) != tt.wantErr {
				t.Errorf("addToBalance() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestEngine_addToBasket(t *testing.T) {
	type fields struct {
		selecting    State
		depositing   State
		dispensing   State
		currentState State
		balance      float64
		basket       []Item
		inventory    map[string]Item
		history      map[string]eventstore.Events
	}
	type args struct {
		itemID string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name:    "invalid item - not in the inventory",
			args:    args{itemID: "not-found-fake-item-id"},
			wantErr: true,
		},
		{
			name: "valid item - exists in the inventory",
			fields: fields{
				inventory: map[string]Item{
					"found-fake-item-id": {
						name: "fake-name",
					},
				},
			},
			args:    args{itemID: "found-fake-item-id"},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := &Engine{
				selecting:    tt.fields.selecting,
				depositing:   tt.fields.depositing,
				dispensing:   tt.fields.dispensing,
				currentState: tt.fields.currentState,
				balance:      tt.fields.balance,
				basket:       tt.fields.basket,
				inventory:    tt.fields.inventory,
				history:      tt.fields.history,
			}
			if err := e.addToBasket(tt.args.itemID); (err != nil) != tt.wantErr {
				t.Errorf("addToBasket() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestEngine_getTotalPrice(t *testing.T) {
	type fields struct {
		selecting    State
		depositing   State
		dispensing   State
		currentState State
		balance      float64
		basket       []Item
		inventory    map[string]Item
		history      map[string]eventstore.Events
	}
	tests := []struct {
		name   string
		fields fields
		want   float64
	}{
		{
			name: "empty basket",
			fields: fields{
				basket: []Item{},
			},
			want: 0,
		},
		{
			name: "basket with single item",
			fields: fields{
				basket: []Item{
					{
						price: 2,
					},
				},
			},
			want: 2,
		},
		{
			name: "basket with multiple items",
			fields: fields{
				basket: []Item{
					{
						price: 2.5,
					},
					{
						price: 3,
					},
				},
			},
			want: 5.5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := &Engine{
				selecting:    tt.fields.selecting,
				depositing:   tt.fields.depositing,
				dispensing:   tt.fields.dispensing,
				currentState: tt.fields.currentState,
				balance:      tt.fields.balance,
				basket:       tt.fields.basket,
				inventory:    tt.fields.inventory,
				history:      tt.fields.history,
			}
			if got := e.getTotalPrice(); got != tt.want {
				t.Errorf("getTotalPrice() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestEngine_reset(t *testing.T) {
	type fields struct {
		selecting    State
		depositing   State
		dispensing   State
		currentState State
		balance      float64
		basket       []Item
		inventory    map[string]Item
		history      map[string]eventstore.Events
	}
	tests := []struct {
		name   string
		fields fields
	}{
		{
			name:   "resetting balance and basket",
			fields: fields{basket: []Item{{name: "fake-name"}}, balance: 4},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := &Engine{
				selecting:    tt.fields.selecting,
				depositing:   tt.fields.depositing,
				dispensing:   tt.fields.dispensing,
				currentState: tt.fields.currentState,
				balance:      tt.fields.balance,
				basket:       tt.fields.basket,
				inventory:    tt.fields.inventory,
				history:      tt.fields.history,
			}
			e.reset()
		})
	}
}

func TestEngine_setState(t *testing.T) {
	type fields struct {
		selecting    State
		depositing   State
		dispensing   State
		currentState State
		balance      float64
		basket       []Item
		inventory    map[string]Item
		history      map[string]eventstore.Events
	}
	type args struct {
		state State
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{
			name:   "change from selecting to depositing state",
			fields: fields{currentState: &Selecting{}},
			args:   args{state: &Depositing{}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := &Engine{
				selecting:    tt.fields.selecting,
				depositing:   tt.fields.depositing,
				dispensing:   tt.fields.dispensing,
				currentState: tt.fields.currentState,
				balance:      tt.fields.balance,
				basket:       tt.fields.basket,
				inventory:    tt.fields.inventory,
				history:      tt.fields.history,
			}
			e.setState(tt.args.state)
		})
	}
}
