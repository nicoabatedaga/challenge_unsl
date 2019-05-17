package main

import (
	"reflect"
	"testing"
)

func TestOrdenarYSumar(t *testing.T) {
	tests := []struct {
		name    string
		numeros []int64
		want    []int64
		want1   []int64
		want2   []int64
	}{
		{"test1", []int64{12, 21, 33, 4, 1}, []int64{4, 14}, []int64{1, 21, 33}, []int64{5, 35, 33}},
		{"test2", []int64{12, 21, 33, 4}, []int64{4, 12}, []int64{21, 33}, []int64{25, 35}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1, got2 := OrdenarYSumar(tt.numeros)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("OrdenarYSumar() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("OrdenarYSumar() got1 = %v, want %v", got1, tt.want1)
			}
			if !reflect.DeepEqual(got2, tt.want2) {
				t.Errorf("OrdenarYSumar() got2 = %v, want %v", got2, tt.want2)
			}
		})
	}
}
