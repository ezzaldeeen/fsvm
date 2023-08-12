package engine

import "testing"

func TestDepositing_Cancel(t *testing.T) {
	type fields struct {
		engine *Engine
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		{
			name:    "valid cancellation",
			fields:  fields{engine: NewEngine()},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := Depositing{
				engine: tt.fields.engine,
			}
			if err := s.Cancel(); (err != nil) != tt.wantErr {
				t.Errorf("Cancel() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestDepositing_Deposit(t *testing.T) {
	type fields struct {
		engine *Engine
		basket []Item
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
			name: "invalid depositing the amount less or equal to zero",
			fields: fields{
				engine: NewEngine(),
			},
			args:    args{amount: -5},
			wantErr: true,
		},
		{
			name: "invalid depositing the amount less than required",
			fields: fields{
				engine: NewEngine(),
				basket: []Item{{price: 5}},
			},
			args:    args{amount: 2},
			wantErr: true,
		},
		{
			name: "valid depositing the amount is enough",
			fields: fields{
				engine: NewEngine(),
				basket: []Item{{price: 5}},
			},
			args:    args{amount: 5},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.fields.engine.basket = tt.fields.basket
			s := Depositing{
				engine: tt.fields.engine,
			}
			if err := s.Deposit(tt.args.amount); (err != nil) != tt.wantErr {
				t.Errorf("Deposit() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestDepositing_Dispense(t *testing.T) {
	type fields struct {
		engine *Engine
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		{
			name:    "invalid dispensing",
			fields:  fields{engine: NewEngine()},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := Depositing{
				engine: tt.fields.engine,
			}
			if err := s.Dispense(); (err != nil) != tt.wantErr {
				t.Errorf("Dispense() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestDepositing_Select(t *testing.T) {
	type fields struct {
		engine *Engine
	}
	type args struct {
		in0 string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name:    "invalid selection",
			fields:  fields{engine: NewEngine()},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := Depositing{
				engine: tt.fields.engine,
			}
			if err := s.Select(tt.args.in0); (err != nil) != tt.wantErr {
				t.Errorf("Select() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
