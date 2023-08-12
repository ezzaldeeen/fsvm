package engine

import "testing"

func TestSelecting_Cancel(t *testing.T) {
	type fields struct {
		engine *Engine
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		{
			name:    "invalid cancel from selecting state",
			fields:  fields{engine: NewEngine()},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := Selecting{
				engine: tt.fields.engine,
			}
			if err := s.Cancel(); (err != nil) != tt.wantErr {
				t.Errorf("Cancel() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestSelecting_Deposit(t *testing.T) {
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
			name:    "invalid deposit from selecting state",
			fields:  fields{engine: NewEngine()},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := Selecting{
				engine: tt.fields.engine,
			}
			if err := s.Deposit(tt.args.in0); (err != nil) != tt.wantErr {
				t.Errorf("Deposit() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestSelecting_Dispense(t *testing.T) {
	type fields struct {
		engine *Engine
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		{
			name:    "invalid dispense from selecting state",
			fields:  fields{engine: NewEngine()},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := Selecting{
				engine: tt.fields.engine,
			}
			if err := s.Dispense(); (err != nil) != tt.wantErr {
				t.Errorf("Dispense() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestSelecting_Select(t *testing.T) {
	type fields struct {
		engine    *Engine
		inventory map[string]Item
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
			name:    "item not found in selecting state",
			fields:  fields{engine: NewEngine()},
			args:    args{itemID: "fake-not-found-item-id"},
			wantErr: true,
		},
		{
			name: "item found in selecting state",
			fields: fields{
				engine: NewEngine(),
				inventory: map[string]Item{
					"fake-found-item-id": {name: "fake-name"},
				},
			},
			args:    args{itemID: "fake-found-item-id"},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		tt.fields.engine.inventory = tt.fields.inventory
		t.Run(tt.name, func(t *testing.T) {
			s := Selecting{
				engine: tt.fields.engine,
			}
			if err := s.Select(tt.args.itemID); (err != nil) != tt.wantErr {
				t.Errorf("Select() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
