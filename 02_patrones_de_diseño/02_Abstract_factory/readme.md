Ejemplo del uso del patron de diseño __Abstract Factory__. Crea diferentes equipamientos (``Arma``, ``Armadura``, ``Casco``) para la raza seleccionada `elfos`, `enanos` y `orcos`

A diferencia del patron Factory Method, que la factory entregaba directamente los productos (canales de comunicación), aca la factory nos devuelve otra factory (especifica de la raza seleccionada) y es esta ultima la que genera los productos específicos.

Si quisiera agregar una nueva familia de productos (para una nueva raza), simplemente generaria la nueva raza, que implemente la interface `raza` y la agrego al select del factory para poder utilizarla.

Si quisiera generar un nuevo equipamiento, por ejemplo botas, genero el nuevo equipamiento ( genérico y especifico ) modifico la interface

Las razas se componen de la raza genérica, que permite manejar el caso en que una raza no pueda usar un equipamiento. Como ejemplo esta el caso del orco, que no puede usar casco