package main

import (
	"testing"
)

func TestMALG_UpdateJ(t *testing.T) {
	type fields struct {
		w1 float64
		w2 float64
		f1 float64
		f2 float64
		f3 float64
	}
	tests := []struct {
		name   string
		fields fields
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &MALG{
				w1: tt.fields.w1,
				w2: tt.fields.w2,
				f1: tt.fields.f1,
				f2: tt.fields.f2,
				f3: tt.fields.f3,
			}
			m.UpdateJ()
		})
	}
}

func TestMALG_UpdateWeight(t *testing.T) {
	type fields struct {
		w1 float64
		w2 float64
		f1 float64
		f2 float64
		f3 float64
	}
	tests := []struct {
		name   string
		fields fields
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &MALG{
				w1: tt.fields.w1,
				w2: tt.fields.w2,
				f1: tt.fields.f1,
				f2: tt.fields.f2,
				f3: tt.fields.f3,
			}
			m.UpdateWeight()
		})
	}
}
