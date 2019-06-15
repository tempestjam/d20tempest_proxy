package main

import (
	fmt "fmt"
	http "net/http"

	"github.com/joho/godotenv"
)

func main() {
	fmt.Println("=====================================")
	router := createRouter()
	host := "0.0.0.0"
	port := "8000"

	e := godotenv.Load()
	if e != nil {
		fmt.Print(e)
	}

	//dbSession := database.Connect()
	//results := dbSession.Session.Collection("user").Find(bson.M{})
	//user := &database.User{}
	//for results.Next(user) {
	//	fmt.Println(user.Mail)
	//}
	//headersOk := handlers.AllowedHeaders([]string{"X-Requested-With"})
	//originsOk := handlers.AllowedOrigins([]string{os.Getenv("ORIGIN_ALLOWED")})
	//methodsOk := handlers.AllowedMethods([]string{"GET", "HEAD", "POST", "PUT", "OPTIONS", "WS"})

	fmt.Println("Ready -> Listen to " + host + ":" + port)
	//err := http.ListenAndServe(host+":"+port, handlers.CORS(originsOk, headersOk, methodsOk)(router))
	err := http.ListenAndServe(host+":"+port, router)
	if err != nil {
		fmt.Print(err)
	}
}
