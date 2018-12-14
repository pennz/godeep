package main

import (
	"testing"
)

func Test_main(t *testing.T) {
	tests := []struct {
		name string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			main()
		})
	}
}

func TestMALG_Calculate(t *testing.T) {
	type fields struct {
		J11        float64
		J21        float64
		J22        float64
		J12        float64
		R1         float64
		R2         float64
		_MALGModel _MALGModel
	}
	tests := []struct {
		name   string
		fields fields
		want   float64
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &MALG{
				J11:        tt.fields.J11,
				J21:        tt.fields.J21,
				J22:        tt.fields.J22,
				J12:        tt.fields.J12,
				R1:         tt.fields.R1,
				R2:         tt.fields.R2,
				_MALGModel: tt.fields._MALGModel,
			}
			if got := m.Calculate(); got != tt.want {
				t.Errorf("MALG.Calculate() = %v, want %v", got, tt.want)
			}
		})
	}
}
