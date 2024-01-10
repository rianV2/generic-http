package main

import "fmt"

type A struct {
	ID      string
	AMember string
}
type B struct {
	ID      string
	BMember string
}

type GS interface {
	Id() string
	SetId(string)
}

func (s A) Id() string      { return s.ID }
func (s *A) SetId(i string) { s.ID = i }
func (s B) Id() string      { return s.ID }
func (s *B) SetId(i string) { s.ID = i }

type MyStore[T GS] struct {
	values map[string]T
}

func (s *MyStore[T]) add(item T) {
	item.SetId("aa")
	s.values["aa"] = item
}

func main2() {
	var storeA = &MyStore[*A]{}
	storeA.values = make(map[string]*A)
	a := &A{}

	storeA.add(a)

	fmt.Println(a.Id())
}
