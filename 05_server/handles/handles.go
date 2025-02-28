package handles

import (
	"fmt"
	"io"
	"net/http"
)

func Hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "hello\n")
	// io.WriteString(w, "This is my website!\n") //otra forma de enviar una respuesta
}

func Query(w http.ResponseWriter, r *http.Request) {

	hasFirst := r.URL.Query().Has("first")
	first := r.URL.Query().Get("first")
	hasSecond := r.URL.Query().Has("second")
	second := r.URL.Query().Get("second")

	fmt.Printf("Request. first(%t)=%s, second(%t)=%s\n", hasFirst, first, hasSecond, second)

}

func AccederBody(w http.ResponseWriter, r *http.Request) {

	body, err := io.ReadAll(r.Body)
	if err != nil {
		fmt.Printf("No se puede leer el body: %s\n", err)
	}

	fmt.Printf(" body:\n%s\n", body)

}

func EncabezadoYCodigoEstado(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("x-missing-field", "UnValor")
	w.WriteHeader(http.StatusBadRequest)

}
