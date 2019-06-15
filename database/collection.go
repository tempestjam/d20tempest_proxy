package database

//
//import (
//	"fmt"
//	"log"
//
//	"github.com/go-bongo/bongo"
//)
//
//func save(connection *bongo.Connection, collectionName string, data bongo.Document) {
//	err := connection.Collection(collectionName).Save(data)
//
//	if err == nil {
//		return
//	}
//
//	if vErr, ok := err.(*bongo.ValidationError); ok {
//		fmt.Println("Validation errors are:", vErr.Errors)
//		return
//	}
//
//	log.Println(err.Error())
//}
//
//func find(connection *bongo.Connection, collectionName string, filter map[string]interface{}) {
//	results := connection.Collection(collectionName).Find(filter)
//	paginateInfo := results.Paginate
//
//	for results.Next(person) {
//		fmt.Println(person.FirstName)
//	}
//
//}
