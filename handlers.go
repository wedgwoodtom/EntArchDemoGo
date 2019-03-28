package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Welcome!")
}

func Login(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	fmt.Fprintln(w, "vars:", vars)

	//todoId := vars["todoId"]
	//fmt.Fprintln(w, "Todo show:", todoId)

	//if err := json.NewEncoder(w).Encode(todos); err != nil {
	//	panic(err)
	//}

}
