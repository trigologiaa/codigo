# Taller de Go

Pasos para empezar a probar el código en este repositorio:

1. Clonar el repositorio y moverse a la carpeta del taller:

   ```bash
   git clone git@github.com:untref-ayp2/taller-GO.git
   cd taller-GO
   ```

2. Ejecutar el ejemplo que quieran. Para esto hay muchisimas formas de hacerlo:

   ```bash
   $ go run ./00-hola/main.go
   ¡Hola, Mundo!
   ```

   ```bash
   $ go run ./00-hola
   ¡Hola, Mundo!
   ```

   ```bash
   $ go run ./00-hola/.
   ¡Hola, Mundo!
   ```

   ```bash
   $ go run github.com/untref-ayp2/taller-GO/00-hola
   ¡Hola, Mundo!
   ```

   ```bash
   $ cd 00-hola
   00-hola $ go run .
   ¡Hola, Mundo!
   ```

   ```bash
   00-hola $ go run main.go
   ¡Hola, Mundo!
   ```

   ```bash
   00-hola $ go run github.com/untref-ayp2/taller-GO/00-hola
   ¡Hola, Mundo!
   ```

   ```bash
   00-hola $ cd ../07-punteros # desde cualquier carpeta dentro del módulo
   07-punteros $ go run github.com/untref-ayp2/taller-GO/00-hola
   ¡Hola, Mundo!
   ```

## Solucionando problemas

### Requisitos mínimos

En este punto, damos por descartado que en sus entornos pueden ejecutar comandos
de Go en una terminal. Si este no fuera el caso, deberian solicitar asistencia
en el Slack de la materia.

Como requerimiento base, deberían poder ejecutar el siguiente comando:

```bash
$ go version
go version go1.22.0 <nombre-de-la-arquitectura>
```

Como recomendación utilicen la última versión disponible, que al día de hoy es
1.22, o de mínima la version 1.20 que es la versión donde se introdujo el
soporte para módulos.

### Problemas al ejecutar `go run`

Esto fue sucediendo a algunos estudiantes, principalmente en entornos Windows.
Aunque no es exclusivo de este sistema operativo.

Si cuando quieren ejecutar uno de los ejemplos del taller, puede pasar que les
diga esto:

```bash
taller-GO/07-punteros $ go run main.go
main.go:6:2: cannot find package "github.com/untref-ayp2/taller-GO/07-punteros/punteros" in any of:
        /usr/local/go/src/github.com/untref-ayp2/taller-GO/07-punteros/punteros (from $GOROOT)
        /Users/tiagox/go/src/github.com/untref-ayp2/taller-GO/07-punteros/punteros (from $GOPATH)
```

Esto podría deberse a que el soporte para módulos esté desactivado. Para
verificar esto deberiamos chequear la variable de entorno de Go `GO111MODULE`, y
para eso vamos a usar `go env` que es la herramienta para administrar variables
de entorno relacionadas a Go.

Primero veamos que valor tiene:

```bash
taller-GO/07-punteros $ go env GO111MODULE
off
```

El comando muestra el valor en `off`, estamos en problemas, pero la solución es
muy simple. Solamente debemos actualizar la variable con un nuevo valor, usando
la opción `-w`:

```bash
taller-GO/07-punteros $ go env -w GO111MODULE=''
```

Con esto estamos dejando que use el valor por defecto de la variable que es `auto`.

Si ahora ejecuto el código nuevamente:

```bash
taller-GO/07-punteros $ go run main.go
x =  12
p2 =  0x1400009c008
y =  10
p2 =  0x1400009c010
La dirección de memoria del numero es:  0x1400009c008
x =  13
x =  13
```

Todo debería funcionar según lo esperado.