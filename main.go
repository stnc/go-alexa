package main

import (
	"fmt"

	"avia/goalexa"
	"net/http"
)

//import 	"github.com/aivahealth/goalexa/alexaapi"

func main() {
	test := goalexa.NewSkill("amzn1.ask.skill.8271ee57-716d-46db-bf6d-684e27ca4052")

	http.HandleFunc("/check", test.ServeHTTP)
	var port string = "8090"
	fmt.Println("server running localhost:" + port)
	http.ListenAndServe(":"+port, nil)

}
