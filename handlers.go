package main

import (
	"fmt"
	"net/http"
)

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Welcome!")
}

func Login(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello There")

	//todoId := vars["todoId"]
	//fmt.Fprintln(w, "Todo show:", todoId)

	//if err := json.NewEncoder(w).Encode(todos); err != nil {
	//	panic(err)
	//}

}
