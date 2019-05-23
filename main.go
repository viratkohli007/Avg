package main

import (
	"fmt"
	"funcs"
	"net/http"
)

func Check(w http.ResponseWriter, r *http.Request) {

	w.Write([]byte("Hello"))
}

func main() {
	port := ":8080"
	http.HandleFunc("/", Check)
	// http.HandleFunc("/adddialogue", funcs.AddDialogue)
	http.HandleFunc("/addcharacter", funcs.AddCharacter)
	err := http.ListenAndServe(port, nil)
	if err != nil {
		fmt.Println("err->", err)
	}
	fmt.Println("Server is running on port ", port)
}
