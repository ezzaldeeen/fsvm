package engine

import "testing"

func TestDispensing_Cancel(t *testing.T) {
	type fields struct {
		engine *Engine
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		{
			name:    "invalid cancellation from dispensing state",
			fields:  fields{engine: NewEngine()},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := Dispensing{
				engine: tt.fields.engine,
			}
			if err := s.Cancel(); (err != nil) != tt.wantErr {
				t.Errorf("Cancel() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestDispensing_Deposit(t *testing.T) {
	type fields struct {
		engine *Engine
	}
	type args struct {
		in0 float64
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name:    "invalid deposit in dispensing state",
			fields:  fields{engine: NewEngine()},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := Dispensing{
				engine: tt.fields.engine,
			}
			if err := s.Deposit(tt.args.in0); (err != nil) != tt.wantErr {
				t.Errorf("Deposit() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestDispensing_Dispense(t *testing.T) {
	type fields struct {
		engine  *Engine
		basket  []Item
		balance float64
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		{
			name: "valid dispense from dispensing state",
			fields: fields{
				engine:  NewEngine(),
				basket:  []Item{{price: 2}},
				balance: 4,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.fields.engine.basket = tt.fields.basket
			tt.fields.engine.balance = tt.fields.balance
			s := Dispensing{
				engine: tt.fields.engine,
			}
			if err := s.Dispense(); (err != nil) != tt.wantErr {
				t.Errorf("Dispense() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestDispensing_Select(t *testing.T) {
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
			name:    "invalid select in dispensing state",
			fields:  fields{engine: NewEngine()},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := Dispensing{
				engine: tt.fields.engine,
			}
			if err := s.Select(tt.args.in0); (err != nil) != tt.wantErr {
				t.Errorf("Select() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
