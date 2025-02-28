package main

import (
	"errors"
	"fmt"
	"net/http"
	"os"
	"server/handles"
)

func main() {

	http.HandleFunc("/hello", handles.Hello)
	http.HandleFunc("/query", handles.Query) //http://localhost:8080/query?first=1&second=hola
	http.HandleFunc("/body", handles.AccederBody)
	http.HandleFunc("/encabezado", handles.EncabezadoYCodigoEstado)

	fmt.Println("Server montado")

	err := http.ListenAndServe(":8080", nil)

	if errors.Is(err, http.ErrServerClosed) {
		fmt.Printf("server closed\n")

	} else if err != nil {
		fmt.Printf("error starting server: %s\n", err)
		os.Exit(1)
	}
}
