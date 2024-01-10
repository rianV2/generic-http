// You can edit this code!
// Click here and start typing.
package main

import "fmt"

type Person struct {
	ID   int
	Name string
}

type Query interface {
	save(person Person) error
}

type QueryEmpty struct{}

type QueryData struct {
	ID int
}

func (db *QueryEmpty) save(person Person) error {
	// Implementation for saving a person to the local database
	fmt.Printf("Saving person in Local DB: %+v\n", person)

	return nil
}

func (db *QueryData) save(person Person) error {
	db.ID = person.ID
	// Implementation for saving a person to the map in meory
	fmt.Printf("Saving person in memory DB: %+v\n", person)
	return nil
}

type DBPointer[T any] interface {
	*T
	Query
}

type StandardServe[T any, dbPointer DBPointer[T]] struct{}

func (h StandardServe[T, dbPointer]) Serve(person Person) error {
	var db dbPointer = new(T)
	db.save(person)
	return nil
}

func main() {
	person := Person{
		ID:   101,
		Name: "John D",
	}

	queryData := StandardServe[QueryData, *QueryData]{}
	queryEmpty := StandardServe[QueryEmpty, *QueryEmpty]{}

	queryData.Serve(person)

	queryEmpty.Serve(person)
}
