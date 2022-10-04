package main

import (
	"log"
	"reflect"
	"testing"

	"github.com/tjarratt/babble"
)

func getData() []TestType {
	babbler := babble.NewBabbler()
	r := make([]TestType, 0)
	fields := reflect.VisibleFields(reflect.TypeOf(struct{ TestType }{}))
	for i := 0; i < 1e3; i++ {
		for _, field := range fields {
			var x = TestType{}
			x[field.Name]
		}
	}
}
func TestMarshalJSONA(t *testing.T) {
	x := TestType{}
	b, err := x.MarshalJSONA()
	if err != nil {
		t.Errorf("Error marshalling")
	}
	log.Println(string(b))
	c, err := x.MarshalJSONB()
	if err != nil {
		t.Errorf("Error marshalling")
	}
	log.Println(string(c))
}

func BenchmarkMarshalJSONA(b *testing.B) {
	x := TestType{}
	m
	_, err := x.MarshalJSONA()
	if err != nil {
		b.Errorf("Error marshalling")
	}
}

func BenchmarkMarshalJSONB(b *testing.B) {
	x := TestType{}
	_, err := x.MarshalJSONB()
	if err != nil {
		b.Errorf("Error marshalling")
	}
}
