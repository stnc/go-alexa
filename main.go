package main

import (
	"fmt"
	"go.uber.org/zap"
	"log"

	"avia/goalexa"
	"github.com/joho/godotenv"
	"net/http"
)

// import 	"github.com/aivahealth/goalexa/alexaapi"
func init() {
	//To load our environmental variables.

	if err := godotenv.Load(); err != nil {
		log.Println("no env gotten")
	}

	/* this is server running
		    dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	        if err != nil {
	            log.Fatal(err)
	        }
	        environmentPath := filepath.Join(dir, ".env")
	        err = godotenv.Load(environmentPath)
	        fatal(err)
	*/

}
func main() {

	cfg := zap.NewProductionConfig()
	cfg.OutputPaths = []string{
		"myproject.log",
	}
	cfg.Build()

	test := goalexa.NewSkill("amzn1.ask.skill.8271ee57-716d-46db-bf6d-684e27ca4052")

	http.HandleFunc("/check", test.ServeHTTP)
	test.RegisterHandlers()

	var port string = "8090"
	fmt.Println("server running localhost:" + port)
	http.ListenAndServe(":"+port, nil)

}
