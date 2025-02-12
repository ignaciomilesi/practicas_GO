Adapter es un patrón de diseño estructural que permite la colaboración entre objetos con interfaces incompatibles.

Originalmente teníamos dos estructuras: agujero.Redondo y pieza.Redondo

```go
// agujero
type Redondo struct {
	radio int
}

func (r *Redondo) SetRadio(NuevoRadio int) {
	r.radio = NuevoRadio
}

func (r Redondo) ComprobarEncaje(piezaRedonda pieza.Redonda) bool {
	return r.radio >= piezaRedonda.GetRadio()
}

```

```go
// pieza
type Redonda struct {
	radio int
}

func (r *Redonda) SetRadio(NuevoRadio int) {
	r.radio = NuevoRadio
}

func (r Redonda) GetRadio() int {
	return r.radio
}

```

El agujero.Redondo posee el método `ComprobarEncaje` que recibe una pieza.Redonda y verifica si entra en el agujero.

Ahora aparece la estructura pieza.Cuadrada: 

```go
//pieza
type Cuadrada struct {
	lado int
}

func (c *Cuadrada) SetLado(NuevoLado int) {
	c.lado = NuevoLado
}

func (c *Cuadrada) GetLado() int {
	return c.lado
}

```

Ya no podría utilizar `ComprobarEncaje` por que solo acepta estructuras tipo pieza.Redonda, por lo que tengo dos caminos:

1. Aplicar una interface, modificar el `ComprobarEncaje` para que acepte la interfaz y agrega a pieza.Cuadrada el método GetRadio (devolverá la diagonal del cuadrado):

```go
// se genera la interface
type ParaEncajar interface {
	GetRadio() int
}
```
```go
// se agrega el metodo a pieza.cuadra
func (c Cuadrada) GetRadio() int {
	return int(math.Sqrt(float64(c.lado)) / 2)
}
```

```go
// modifico el método para que acepte la interface
func (r Redondo) ComprobarEncaje(pieza pieza.ParaEncajar) bool {
	return r.radio >= pieza.GetRadio()
}
```

Esta forma tiene la ventaja de que si aparecen mas piezas diferentes (triangular, ovalada, etc) simplemente hacemos que implemente la interface y podríamos seguir utilizando el método `ComprobarEncaje`.

Pero como desventaja tenemos que tuvimos que modificar el método de la estructura agujero.Redondo y que todas las piezas tienen que tener el método `GetRadio` por lo quedaría "sucia" la estructura

2. Una forma para que quede mas "limpio" el código es generando una función que tome una estructura del tipo
pieza.cuadrada y devuelva una estructura del tipo
pieza.redonda cuyo radio sea el equivalente:

```go
func AdaptadorRedondoCuadrado(c Cuadrada) Redonda {

	return Redonda{
		radio: int(math.Sqrt(float64(c.lado)) / 2),
	}
}
```
y para utilizarla con el `ComprobarEncaje`:

```go
piezaCuad := pieza.Cuadrada{}

piezaCuad.SetLado(10)

fmt.Printf(
    "Encaje Red-Cuad: %t \n",
    agujeroRed.ComprobarEncaje(pieza.AdaptadorRedondoCuadrado(piezaCuad)))

```
_(esta variante es el que quedo guardada en el código)_

Este método tiene la ventaja que no necesito modificar las estructura de la pieza y del agujero, dejando un código mas limpio.

Presentan el inconveniente de que si aparecen mas piezas diferentes (triangular, ovalada, etc) es necesario crear los métodos adaptadores para cada uno y, debido a que el método genera una nueva estructura, si estas son grandes y complejas, puede ser un método poco optimo para implementar