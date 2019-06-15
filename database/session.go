package database

import (
	"fmt"
	"log"
	"os"

	"github.com/go-bongo/bongo"
)

type Connection struct {
	Session *bongo.Connection
}

func Connect() Connection {
	db_host := os.Getenv("db_host")
	db_port := os.Getenv("db_port")
	db_name := os.Getenv("db_name")

	config := &bongo.Config{
		ConnectionString: db_host + ":" + db_port,
		Database:         db_name,
	}
	fmt.Println("Connect " + db_host + ":" + db_port)
	fmt.Println("to db : " + db_name)

	connection, err := bongo.Connect(config)

	if err != nil {
		log.Fatal(err)
	}

	return Connection{Session: connection}
}
