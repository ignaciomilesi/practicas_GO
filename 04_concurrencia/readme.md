# Concurrencia

GO utiliza concurrencia para ejecutar multiples tareas. Para comprender que es la concurrencia explicaremos que es secuencial, concurrente y paralela:

- Secuencial: Ejecutar una tarea, donde una vez terminada se ejecuta la siguiente.

![secuencial](img/01_secuencial.webp)

Es decir, termina de ejecutar todos los paso de la tarea A para comenzar con la B

- Paralelismo: Ejecutar varias tareas simultáneamente o dicho de otra manera hacer muchas cosas al mismo tiempo.

![paralelismo](img/03_paralelismo.webp)

Ahora las tareas se ejecutan en simultanea, los pasos de la tarea A y B se ejecutan al mismo tiempo, no importa si comenzo o termino la otra tarea.

- Concurrencia: Ejecutar varias tareas mediante periodos de tiempo supuestos y no uno tras otro o dicho de otra manera es tratar con muchas cosas al mismo tiempo

![Concurrecia](img/02_concurrencia.webp)

En otras palabras, no espera a terminar todas los pasos de la tarea A para comenzar con la tarea B, sino que, si encuentra una ventana de tiempo entre los pasos de la tarea A comienza con los pasos de la tarea B y alterna entre una y otra según la disponibilidad.

Puede darnos la ilusión de paralelismo, ya que ambos procesos o tareas hacen un progreso CASI idénticos , aunque en realidad solo ocurre UNA tarea en un momento dado.

## Goroutine

La concurrencia es manejada en Go mediante las goroutines (hilo ligero de ejecución). Para lanzar una goroutine anteponemos la palabra `go` al momento de lanzar la función:

```go
func main(){
    go tarea1()
    go tarea2()
    go tarea3()
}
```
De esta manera, el main (posee su propia Goroutine), la tarea1, tarea2 y tarea3 se ejecutan en su propia Goroutine.

Ya que el main es nuestro hilo principal, al finalizar, finaliza también el programa. En el ejemplo anterior, luego de lanzar las goroutine el programa finalizará sin esperar que finalicen las tarea1, tarea2 o tarea3, ya que no hay nada que impida o que deje esperando a que las goroutine lanzadas finalicen.

Una forma de realizar esta espera es mediante WaitGroup

### WaitGroup 

Un WaitGroup es simplemente un contador que me permite llevar la cuenta de cuantas Goroutine hay activas y esperar a que finalicen. Si tenemos lo siguiente:

```go
for i := 0; i < 10; i++ {
    go func() {
        fmt.Println(i)
    }()
}
```

Veremos no se imprime nada, ya que no hay nada que espere la finalización de las goroutine, para ello uso un WaitGroup como contador

```go
// creamos el contador
var wg sync.WaitGroup

for i := 0; i < 10; i++ {

    // sumamos 1 antes de tirar la goroutine
    wg.Add(1)

    go func() {
        //resta 1 
        defer wg.Done() // el "defer" es para asegurar que sea lo ultimo que se ejecute
        fmt.Println(i)
    }()
}

// bloqueamos hasta que el contador sea 0
wg.Wait()
```
Ahora veremos que se imprimen los números del 0 al 9, debido a que, con `wg.Add(1)`, `wg.Done()` y `wg.Wait()` contamos cuantas goroutine lanzamos, cuando termino su ejecución y bloqueamos esperando a que terminen todas

> <sub>
> Vemos que la salida es desordenada, esto es debido a la concurrencia que ejecuta las tareas según la disponibilidad
> </sub>

Otra forma de esperar a la ejecución de las goroutine, aunque poseen otros usos mas interesantes, es mediante el uso de Channels

## Channels 

Los channels son como tuberías que conectan las goroutine entre sí. Puedes enviar datos a través de un channel y recibirlos en otro. Y lo mejor de todo es que los channels son seguros para la concurrencia, por lo que no tienes que preocuparte por los problemas de sincronización.

Para crear un channel utilizamos la función `make` e indicamos el tipo de dato que pasara por él. Por ejemplo: `c := make(chan int)` creamos el canal `c` por el que puede pasar datos tipo `int`.

Para indicar cuando se coloca y se saca un dato del canal, se utiliza el operador `<-`: 
- `c <- 4` : indica que se envía el valor 4 por el canal
- `<- c`: indica que se extrae el valor del canal

En el ejemplo `ejemplo.CreandoChannel()`:

```go
c := make(chan int)

go func(ch chan int, numero int) {

    ch <- numero
    fmt.Println("Numero Colocado")

}(c, 10)

fmt.Print(<-c)
```

En la función anónima, que se ejecuta en una goroutine independiente, se envía el valor al canal (`ch <- numero`) que es tomado por la goroutine main e impreso (`fmt.Print(<-c)`) 

### Channels de lectura y escritura

El channel indicado anteriormente es bidireccional, permite recibir datos como enviarlos, pero se puede indicar que sea solo de lectura o escritura:

- Channel definido solo escritura: permite solo enviar datos `func Generator(c chan<- int)`

- Channel definido solo lectura: permite solo recibir datos `func Print(c <-chan int)`

Se utiliza cuando define la función, creo los Channel bidireccionales pero al enviarlos a una función, le indico como quiero que se comporte

```go
// define la función e indico que el Channel que pase como parámetro se use como solo para recibir
func tarea(ch <-chan){
    ...
}

// en otra funcion:
ch := make(chan int) // creo el Channel bidireccional

tarea(ch) // ya que lo definí para que se comporte de solo lectura, dentro de la función se comportara como tal
```
De esta forma evito que se envíen o extraigan datos en funciones y Channel que no deseo.

## Bloqueo del channel

En el primer ejemplo podemos observar que el programa no finaliza luego de lanzar la goroutine, sino que la espera. Esto ocurre, debido a que las goroutine, cuando intentan lanzar o tomar un dato de un canal, quedan bloqueada hasta que puedan realizarlo. 

En el ejemplo, la goroutine main quedo bloqueada en `fmt.Print(<-c)` esperando que otra goroutine lanzara un dato. Cuando lo lanzo, tanto la goroutine main como la goroutine que lanzo el dato, pueden continuar con la ejecución

Una goroutine se puede bloquea esperando recibir un dato como también, esperando para enviarlo. Si queremos enviar un dato pero no hay nada para recibirlo, la goroutine quedar bloqueada, si vemos el ejemplo `ejemplo.BloqueoChannelConSleep()`:

```go
c := make(chan int)

go func(ch chan int, numero int) {

    ch <- numero
    fmt.Println("Numero Colocado")

}(c, 10)

time.Sleep(3 * time.Second)
fmt.Print("hola")
```
Lo que se imprimirá sera `hola` pero no `Numero colocada` debido que la goroutine se bloquea en `ch <- numero` al querer lanzar un dato pero no hay otra goroutine para tomarlo.

Esto, si no se maneja con cuidado, puede producir deadlock. En el ejemplo `ejemplo.BloqueoChannelConOtroChannel()`:

```go
dato := make(chan int)
finTarea := make(chan bool)

go func(c_dato chan int, c_finTarea chan bool, numero int) {

    c_dato <- numero
    fmt.Println("Numero Colocado")
    c_finTarea <- true

}(dato, finTarea, 10)

<-finTarea
fmt.Print(<-dato)
```
Tenemos dos canales: el canal `dato` por donde se envía el numero y el canal `finTarea` usado como bandera para saber cuando finalizo la goroutine. El ejemplo produce un error tipo deadlock, ya que, la goroutine main queda bloqueada a la espera de recibir algo por `finTarea` pero la goroutine de la función anónima esta bloqueada a la espera de querer enviar algo por `dato`, una bloquea a la otra.

Otra forma de generar un deadlock es querer sacar un algo de un canal que ya se le ha extraido el dato. Si al ejemplo `ejemplo.CreandoChannel()` le coloco otro `fmt.Print(<-c)` al final:

```go
c := make(chan int)

go func(ch chan int, numero int) {

    ch <- numero
    fmt.Println("Numero Colocado")

}(c, 10)

fmt.Print(<-c)
fmt.Print(<-c)
```

Produzco un deadlock, ya que el segundo print queda esperando un dato que ya fue extraído, pero la razón es la misma que ya comentamos, la goroutine main queda bloqueada en el segundo print a la espera que le envíen algo por el canal, como no hay otra goroutine que lo envié, se produce el deadlock.

Lo mismo ocurriría si realizo un doble envío pero solo lo puedo tomar una vez, en el `ejemplo.DeadlockPorDobleenvio()` se puede comprobar:

```go
c := make(chan int)

go func(ch chan int) {

    fmt.Println(<-ch)

}(c)

c <- 10
c <- 10

time.Sleep(100 * time.Millisecond) //solo para darle tiempo a que se ejecute la goroutine
```

## Unbuffered channels y Buffered channels

Existen  dos tipos de canales: 

- __Unbuffered channels__: la cantidad límite de datos simultáneos, que puede manejar el canal, es 0. Son los que estuvimos viendo anteriormente, al recibir un dato necesita enviarlo, ya que no posee lugar para guardarlo, por ello se bloquean las goroutine si no hay un emisor y un receptor. 


- __Buffered channels__: la cantidad límite de datos simultáneos, que puede manejar el canal, es pasado como argumento extra al crearlo, es decir, pueden almacenar 1 o mas datos. 

Para crear un buffered channel lo indico en el make: `ch := make(chan int, 10)`, para este ejemplo, estoy creando un channel con un buffer de 10, puede contener hasta 10 datos

Este tipo de canal evita el bloqueo al momento del envío pero sigue estando al momento de la extraction: si no hay dato que sacar, la goroutine se bloqueara hasta que aparezca un dato que extraer.

Si vemos el ejemplo `ejemplo.EvitandoBloqueoChannelConOtroChannel()`:

```go
dato := make(chan int, 1) // único cambio
finTarea := make(chan bool)

go func(c_dato chan int, c_finTarea chan bool, numero int) {

    c_dato <- numero
    fmt.Println("Numero Colocado")
    c_finTarea <- true

}(dato, finTarea, 10)

<-finTarea
fmt.Print(<-dato)
```

Es exactamente igual al ejemplo anterior pero se indica que el channel `dato` tiene un buffer de 1. 

Se puede ver que ahora no ocurre el deadlock, debido a que la goroutine al querer enviar el numero al canal `c_dato <- numero`, para este caso y gracias al buffer, encuentra lugar en el channel por lo que lo envía y continua con la goroutine, permitiendo que llegue al final y envié el dato por el channel `finTarea`.

Hay que tener en cuenta que si, el channel llenara el buffer ocurriría lo mismo que para un channel unbuffered, las goroutine se bloquearían esperando lugar para realizar el envió

### Buffered channels como semáforos

Uno de los usos de los channels con buffer es la de limitar la cantidad de procesos concurrentes que se ejecutan al mismo tiempo, actuando como un "semáforo" que regula el flujo de ejecución de las goroutine. 

Con un buffer, puedes predeterminar cuántas goroutine pueden estar activas simultáneamente, asegurando así un control más fino y evitando bloqueos por exceso de carga.

Se aprovecha la características de los channels de bloquear, al momento de querer enviar un dato, la goroutine si no hay espacio para hacerlo. Viendo el ejemplo realizado en `channelsSemaforo.go`:

```go
// el doSomething solo simula una tarea que requiera mucho tiempo

var wg sync.WaitGroup
ch := make(chan int, 3)

for i := 0; i < 40; i++ {
    wg.Add(1)
    ch <- 1 //llenando el channel
    go doSomething(i, &wg, ch)
}

wg.Wait()
```
Al colocar `ch <- 1` antes de lanzar la goroutine, lo estaremos usando como un indicador de cuantas hemos lanzado. Como lo definimos en 3 (`ch := make(chan int, 3)`), al querer lanzar la cuarta, el channel estará lleno, bloqueando la goroutine en ese punto. 

Solo cuando haya lugar en el channel, se podrá continuar y lanzar una nueva goroutine, y esto ocurre cuando una tarea haya finalizado ya que el vaciado del channel se realiza al final del doSomething.

__De esta forma logro que solo 3 goroutine, en simultaneo, se estén ejecutando__

El WaitGroup se encuentra para permitir ejecutar las ultimas tareas, ya que al haber menos de tres, siempre habrá lugar en el channel, lo que me permitirá escarpar del for y terminar la goroutine del main antes que se terminen de ejecutar las otras.

## Worker pools

Están relacionados con los semáforos, buscan gestionar un gran numero de tareas. Imagínate que tienes una serie de tareas concurrentes que quieres realizar, la opción simplista es crear una serie de workers y usarlos de manera concurrente, pero tiene dos grandes **desventajas**: Se estarán **creando workers sin control** y, la segunda, se están **creando y destruyendo workers constantemente**, lo cual puede ser **costoso para tu programa**.

`Worker pool` es un patrón de diseño que suple estas deficiencias, en este, se crea un **número fijo de workers** y se colocan en un ciclo, en el que estarán **escuchando constantemente información de la cola de tareas** (por medio de un channel). De esta manera mantendremos nuestro **manejo de memoria mucho más estable y predecible**, además de que **limitamos el impacto** que ejercerían la **creación y destrucción constantes de workers**.

Viendo el ejemplo realizado en `workerpool.go`:

Defino el worker, para este caso simplemente simula una tarea pesada, que requiere tiempo para procesarse, y devuelve una respuesta:

```go
func worker(id int, jobs <-chan int, results chan<- int) {
	for job := range jobs {
		...
        results <- cuadradoDelNumero
        ...
    }
    fmt.Printf("Worker %d cerrado\n", id)
}
```
El `for` hace que el worker quede en bucle, siempre escuchando si hay trabajo que realizar en la cola, haciendo que, cuando termine una tarea, tome la siguiente disponible. Solo podrá escapar del bucle si se cierra el channel

Se simula una cola de tarea y se establecen cuantos worker habrá 

```go
tasks := []int{2, 3, 4, 5, 7, 10, 12, 40}
numOfWorkers := 3
```
Creo los canales, uno por donde se enviaran las tareas y el otro donde se obtendrán los resultados. Para este caso, los canales tendrán buffer

```go
jobs := make(chan int, len(tasks))
results := make(chan int, len(tasks))
```

Se genera los workers (3 en este caso), estos quedan en bucle esperando las tareas
```go
for w := 1; w <= numOfWorkers; w++ {
    go worker(w, jobs, results)
}
```
Cargo las tareas en el canal y luego cierro el canal, esto permite que el worker termine al acabar las tareas, si no lo colocara los worker seguirían esperando algún dato cuando todo ya se ha procesado y generaria error.
```go
for _, task := range tasks {
    jobs <- task
}
close(jobs)
```
Tomo los resultados. (En este caso solamente los saco del canal)  
```go
for i := 0; i < len(tasks); i++ {
    <-results
}

// solo para dar tiempo a cerrar los worker
time.Sleep(2 * time.Second)
```
El sleep final es para dar tiempo a que se vea que se cierran los worker (ya que estoy trabajando en el main). En algo mas complejo, puede que no lo necesite o necesite emplear un WaitGroup.

En el ejemplo realizado en `workerpool2.go` de desarrolla el mismo programa pero se utilizan canales sin buffer, la diferencia esta en que es necesario generar una goroutine que maneje los resultados (si lo quisiera manejar en el main generaría un deadlock). También se agrego un WaitGroup para asegurar que terminen las goroutine antes de terminar el main.

## Multiplexación 

Si ejecutamos `ejemplo.SinMultiplex()`:

```go
c1 := make(chan int)
c2 := make(chan int)

duration1 := 4 * time.Second
duration2 := 2 * time.Second

go doSomething2(duration1, c1, 1)
go doSomething2(duration2, c2, 2)

fmt.Printf("Channel 1 received: %d\n", <-c1)
fmt.Printf("Channel 2 received: %d\n", <-c2)
```

El channel 1, que tarda 4 segundos en recibir respuesta, se imprime antes que el channel 2, que tarda solo 2 segundos en recibir respuesta. 

Esto ocurre, como ya estuvimos viendo, por que la goroutine main se bloquea en la línea `fmt.Printf("Channel 1 received: %d\n", <-c1)` a la espera de tener un valor que extraer, impidiendo que se llegue a la siguiente linea donde se imprime el channel 2.

Aquí es donde entra en juego el uso de ``select`` para implementar multiplexación. ``Select`` es una estructura similar a un **switch** que te permite escuchar múltiples canales al mismo tiempo, ejecutando el caso que esté listo primero. Permite manejar la llegada de mensajes de manera más flexible que un enfoque secuencial.

El `select` indica que voy a escuchar varios channel al mismo tiempo y con `case` indico que realizo ante la llegada de un dato desde un channel u otro. Si ahora ejecutamos `ejemplo.ConMultiplex()`:

```go
c1 := make(chan int)
c2 := make(chan int)

duration1 := 4 * time.Second
duration2 := 2 * time.Second

go doSomething2(duration1, c1, 1)
go doSomething2(duration2, c2, 2)

for i := 0; i < 2; i++ {
    select {
    case msg1 := <-c1:
        fmt.Printf("Channel 1 received: %d\n", msg1)
    case msg2 := <-c2:
        fmt.Printf("Channel 2 received: %d\n", msg2)
    }
}
```
Veremos que se imprime primero el Channel 2, el de menor tiempo, y luego el Channel 1, el de mayor tiempo.

Para el ``select`` también puedo utilizar `default`, como en el caso del ``switch``, que se ejecutara cuando reviso pero no hay ningún `case` listo aun
