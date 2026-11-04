package main

import (
	"html/template" // para trabajar con la plantilla HTML
	"log"           // para mostrar mensajes por la terminal
	"net/http"      // para crear el servidor web
)

// Pagina agrupa los datos que se insertan en la plantilla.
// Los nombres empiezan por mayúscula para que la plantilla pueda leerlos.
type Pagina struct {
	Titulo  string
	Mensaje string
	Imagen  string
}

// Se carga la plantilla una sola vez, al arrancar el programa.
// Si el archivo no existe o tiene errores, el programa se detiene.
var plantilla = template.Must(template.ParseFiles("templates/index.html"))

// inicio atiende la ruta principal: rellena la plantilla y la envía al navegador.
func inicio(w http.ResponseWriter, r *http.Request) {
	// Datos que se insertan en los huecos de la plantilla
	datos := Pagina{
		Titulo:  "Producto 1 - Entorno de desarrollo",
		Mensaje: "Soy alumno de la UOC",
		Imagen:  "/static/img/uoc.png",
	}
	// Si algo falla al generar la página, se devuelve un error 500
	if err := plantilla.Execute(w, datos); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func main() {
	// Servidor de ficheros para la carpeta static (aquí está la imagen)
	fs := http.FileServer(http.Dir("static"))
	// Las peticiones que empiezan por /static/ buscan el archivo en esa carpeta.
	// StripPrefix quita /static/ de la dirección antes de buscarlo.
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	// La ruta principal ejecuta la función inicio
	http.HandleFunc("/", inicio)

	// Aviso por la terminal y arranque del servidor en el puerto 8080
	log.Println("Servidor escuchando en http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
