Builder es un patrón de diseño creacional que nos permite construir objetos complejos paso a paso. El patrón nos permite producir distintas configuraciones de objetos similares. 

En ves de maneja un constructor largo y complejo, utilizo métodos para agregar características a un objeto:

```go
nuevaCasa := House{
    ventanas: 4
    puertas: 2
    habitaciones: 2
    garage: true
    pileta: false
    jardin: true
}

// quedaria reemplazado por:

casaBuilder := Builder{}

casaBuilder.AgregarVentanas(4)
casaBuilder.AgregarPuertas(2)
casaBuilder.AgregarHabitaciones(2)
casaBuilder.AgregarGarage()
casaBuilder.AgregarJardin()

nuevaCasa := casaBuilder.Construir()

```

Para el caso de GO, inicializar una estructura es muy explicito, por lo que, la primera opción parece mejor. Pero esa opción trae el problema de que los campos de la estructura deben ser públicos, si quisiéramos mantenerlos privados tendría q recurrir a una función `make`:

```go
var nuevaCasa House
nuevaCasa.makeHouse(4, 2, 2, true, false, true)
```
Por lo que, la creación del struct ya no es tan clara

Otra ventaja es la de separar los métodos de construcción de la estructura final, haciendo que quede mas limpia para ser utilizada despues