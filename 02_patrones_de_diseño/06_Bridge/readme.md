__Bridge__ es un patrón de diseño estructural que te permite dividir una clase grande, o un grupo de clases estrechamente relacionadas, en dos jerarquías separadas (abstracción e implementación) que pueden desarrollarse independientemente la una de la otra.

La Abstracción es una capa de control de alto nivel para una entidad. Esta capa no tiene que hacer ningún trabajo real por su cuenta, sino que debe delegar el trabajo a la capa de implementación.

Para nuestro ejemplo la capa abstracción es Parlante y la de implementación es Reproductor: la capa parlante delega a reproductor todo lo referido a reproducir y cambiar de cancion, mientras que parlante solo se encarga del volumen

```go
type generico struct {
	reproductor reproductor.Reproductor //<-- reproductor contenido en parlante
	volumen     int
}

// carga de la capa de implementacion
func (g *generico) ConectarReproductor(nuevoReproductor reproductor.Reproductor) {
	g.reproductor = nuevoReproductor
}

// func especifico de la capa de abstraccion
func (g *generico) AdjustarVolume(nuevoVolumen int) {
	g.volumen = nuevoVolumen
	fmt.Println("Volumen ajustado a:", nuevoVolumen)
}

// func de parlante que delega la ejecucion a  reproductor
func (g *generico) CargarLista(nuevaLista []string) {
	g.reproductor.CargarLista(nuevaLista)
}

func (g generico) Play() {
	g.reproductor.Play()
}

func (g generico) Pausa() {
	g.reproductor.Pausa()
}

func (g generico) SiguienteCancion() {
	g.reproductor.SiguienteCancion()
}
```

La ventaja, al estar separada la abstracción de la implementación, modificar una no afecta a la otra 