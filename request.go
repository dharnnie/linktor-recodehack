package main

import (
	"fmt"
	"net/http"
	"io"
	"os"
	"log"
	
)
// func main() {
// 	fmt.Println("Hello, playground")
	
// 	response, err := http.Get("https://jsonplaceholder.typicode.com/users")
//         if err != nil {
//                 log.Fatal(err)
//         } else {
//                 defer response.Body.Close()
//                 _, err := io.Copy(os.Stdout, response.Body)
//                 if err != nil {
//                         log.Fatal(err)
//                 }
// 	}

// }

func main() {
	//p := {BankCode: "", ClientSec: "", GrantType: ""}
	var p struct{
		BankCode string `json:"bank_code"`
		AccountNum string `json:"accountnumber"`
	}
	res, err := http.Post("pwcstaging.herokuapp.com/account/validation","application/json",p)
	fmt.Println("Here is the Response: ", res)
}