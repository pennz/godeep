package main

import (
	"testing"
)

func Test_helloWorld(t *testing.T) {
	tests := []struct {
		name string
		want int
	}{
		// TODO: Add test cases.
		{"mock", 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := helloWorld(); got != tt.want {
				t.Errorf("helloWorld() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLinkedList_Init(t *testing.T) {
	type fields struct {
		val  int
		next *LinkedList
	}
	tests := []struct {
		name   string
		fields fields
	}{
		// TODO: Add test cases.
		{"zero", fields{0, nil}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			head := &LinkedList{
				val:  tt.fields.val,
				next: tt.fields.next,
			}
			head.Init()
		})
	}
}

func TestLinkedList_AddAt(t *testing.T) {
	type fields struct {
		val  int
		next *LinkedList
	}
	type args struct {
		pos int
		val int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
		{"zero", fields{0, nil}, args{0, -1}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			head := &LinkedList{
				val:  tt.fields.val,
				next: tt.fields.next,
			}
			head.AddAt(tt.args.pos, tt.args.val)
			if got := head.next.val; got != -1 || head.next.next != nil || head.val != 1 {
				t.Errorf("AddAt function has issue")
			}
		})
	}
}
