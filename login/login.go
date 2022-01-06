package login

import (
    "fmt"
	"log"
    "net/http"
	//"github.com/gorilla/websocket"
	"io/ioutil"
	"os"
	"witt_backend/settings"
)

func Login(w http.ResponseWriter, r *http.Request) {
	log.Println("FAKE_API : " ,settings.IsFakeAPI())

	if  settings.IsFakeAPI() {
		jsonFile, err := os.Open("./mock_responses/login.json")
		// if we os.Open returns an error then handle it
		if err != nil {
			fmt.Println(err)
		}
		//fmt.Println("Successfully Opened users.json")
		// defer the closing of our jsonFile so that we can parse it later on
		defer jsonFile.Close()

		byteValue, _ := ioutil.ReadAll(jsonFile)
		//log.Println(byteValue)
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintf(w, string(byteValue))

	}else{
		fmt.Fprintf(w, "Logica real contra Hobbes")
	}
}

