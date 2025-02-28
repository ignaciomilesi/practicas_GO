## Server HTTP básico

Usando el paquete ``net/http`` ya incluido en Go


Un concepto fundamental en el ``net/http`` son los controladores **(handlers)**. Un controlador es un objeto que implementa la interfaz ``http.Handler`` y es el encargado de implementar la lógica, el siguiente controlador imprimirá _hello_ como salida:

```go
func Hello(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "hello\n")
}
```

El controlador se debe definir con ``http.ResponseWriter`` y ``http.Request``. Estos dos parámetros permiten gestionar las solicitudes HTTP entrantes, leer la información proporcionada por el cliente y crear una respuesta adecuada para enviar. 

``http.Request`` proporciona información sobre la solicitud entrante, mientras que ``http.ResponseWriter`` se utiliza para crear la respuesta.

Luego se enrutan con ``http.HandleFunc`` y se monta el server con `http.ListenAndServe`:
```go
func main() {

	http.HandleFunc("/hello", handles.Hello)

	http.ListenAndServe(":8080", nil)
}
```
El ``http.HandleFunc`` indica que cuando llegue una petición a la dirección _"/hello"_ ejecute el controlador **handles.Hello**. El ``http.ListenAndServe`` monta el servidor en el puerto _8080_.

Ahora modifiquemos el ``http.ListenAndServe`` para que pueda manejar los errores:

```go
err := http.ListenAndServe(":8080", nil)

if errors.Is(err, http.ErrServerClosed) {
	
	fmt.Printf("server closed\n")

} else if err != nil {
	
	fmt.Printf("error starting server: %s\n", err)
	os.Exit(1)
}
```

El primer error que se busca, ``http.ErrServerClosed``, se devuelve cuando se le indica al servidor que se apague o se cierre. Este suele ser un error esperado porque usted mismo apagará el servidor, pero también se puede utilizar para mostrar por qué el servidor se detuvo en la salida. 

En la segunda comprobación de errores, se comprueba si hay algún otro error. Si esto sucede, se imprimirá el error en la pantalla y luego se saldrá del programa con un código de error de __1__ uso de función la ``os.Exit``

###  Multiplexor de servidor personalizado

Cuando iniciamos el servidor HTTP, pasamos a la función ``ListenAndServe`` un valor **nil** para el parametro ``http.Handler`` (``http.ListenAndServe(":8080", nil)``), ello es porque usamos el multiplexor de servidor predeterminado. Pero es posible configurar un **Multiplexor de servidor personalizado** a traves de la utilización de ``http.ServeMux``:

```go
func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/hello", handles.Hello)
	//mux.HandleFunc( ... ) otro controlador
	...

	err := http.ListenAndServe(":8080", mux)
	
	...
}
```

### Query

Para trabajar con las query del path se utiliza `r.URL.Query()`. Hay dos métodos que se pueden utilizar para interactuar con los datos:

- ``Has``: devuelve un valor bool que especifica si la cadena de consulta tiene un valor con la clave proporcionada 
- `Get`: devuelve un string con el valor de la clave proporcionada.

En el ejemplo:

```go
func Query(w http.ResponseWriter, r *http.Request) {

	hasFirst := r.URL.Query().Has("first")
	first := r.URL.Query().Get("first")
	hasSecond := r.URL.Query().Has("second")
	second := r.URL.Query().Get("second")

	fmt.Printf("Request. first(%t)=%s, second(%t)=%s\n", hasFirst, first, hasSecond, second)
}
```

Primero consulto si esta la clave _first_ con `r.URL.Query().Has("first")` y luego tomo su valor, con ``r.URL.Query().Get("first")``. Lo mismo hago con _second_ y lo imprimo en consola

### Lectura del body

El ``http.Request`` se utiliza para acceder a información sobre la solicitud entrante y también incluye una forma de acceder al cuerpo de la solicitud (``Body``).

Para acceder al body simplemente utilizamos la propiedad `r.Body` y utilizamos la función `io.ReadAll` para leerlo y poder utilizarlo

### Responder con encabezados y un código de estado

El servidor utiliza los código de estado para dar a un cliente HTTP una mejor idea de como el servidor tomo la solicitud: si se realizó correctamente o si algo salió mal, en el lado del servidor o en el lado del el cliente.

Otra forma en la que los servidores y clientes HTTP se comunican es mediante campos de encabezado. Un campo de encabezado es una clave y un valor que un cliente o servidor enviará al otro para informarle sobre sí mismo. Hay muchos encabezados que están predefinidos por el protocolo HTTP, como Accept, que un cliente usa para indicarle al servidor el tipo de datos que puede aceptar y comprender.

```go
func EncabezadoYCodigoEstado(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("x-missing-field", "UnValor")
	w.WriteHeader(http.StatusBadRequest)

}
```