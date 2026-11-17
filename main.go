package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
)

// Estructura que contiene los datos de nuestra web
type Pagina struct {
	Titulo  string
	Mensaje string
	Imagen  string
}

// Cargamos nuestra plantilla HTML
var plantilla = template.Must(template.ParseFiles("templates/index.html"))

// Función que muestra nuestra página web
func inicio(w http.ResponseWriter, r *http.Request) {
	datos := Pagina{
		Titulo:  "Producto 1 - Entorno de desarrollo",
		Mensaje: "Soy alumno de la UOC",
		Imagen:  "/static/img/uoc.png",
	}

	if err := plantilla.Execute(w, datos); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func main() {
	// Configuramos el puerto del servidor
	puerto := os.Getenv("PORT")
	if puerto == "" {
		puerto = "8080"
	}

	// Permitimos que Go sirva imágenes y otros archivos
	fs := http.FileServer(http.Dir("static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	// Ruta para identificar qué pod de Kubernetes responde
	http.HandleFunc("/pod", func(w http.ResponseWriter, r *http.Request) {
		hostname, err := os.Hostname()
		if err != nil {
			http.Error(w, "Error al obtener el pod", http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "text/plain")
		fmt.Fprintf(w, "Hola desde el pod %s\n", hostname)
	})

	// Asociamos la página principal con nuestra función
	http.HandleFunc("/", inicio)

	// Iniciamos el servidor
	log.Printf("Servidor escuchando en http://localhost:%s", puerto)
	log.Fatal(http.ListenAndServe(":"+puerto, nil))
}
