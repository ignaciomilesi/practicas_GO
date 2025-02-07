Ejemplo del uso del patron de diseño __Factory Method__. La aplicación crea diferentes canales de comunicación para enviar mensaje, en el ejemplo se usa `SMS`, `Email` y `Carta`

El patrón Factory Method sugiere que, en lugar de llamar al operador new para construir objetos directamente, se invoque a un método fábrica especial. El método fábrica sera el encargado de crear los objetos. Los objetos devueltos por el método fábrica a menudo se denominan productos.

En `canalGenerico` esta la lógica de un canal generico (similar a lo que seria una clase abstracta). 

En `canalesEspecficos` se define los diferentes canales:`SMS`, `Email` y `Carta`. para el caso de carta se prueba una sobreescritura de método y el agregado de un campo extra (para su modificacion posterior)

En `factory` se define la interface `canal` y el metodo factory que se encarga de la generacion de objetos.

Si quiera agregar mas canales simplemente agregaria los struct correspondientes que implemente la interface y agregaria el caso generador en Factory