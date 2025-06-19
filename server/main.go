package main

import (
	"net/http"
	"server/utils"
)

func main() {
	mux := http.NewServeMux()
	// mux permite manejar las rutas desde donde
	// se recibirán las peticiones HTTP.

	// ETAPA 3 - CONEXION CLIENTE-SERVIDOR

	// Configurar las rutas y asociar funciones a ellas

	mux.HandleFunc("/paquetes", utils.RecibirPaquetes)
	// HandleFunc asocia una función a una ruta específica.
	// En este caso, cuando se recibe una petición a la ruta "/paquetes",
	// se ejecuta la función utils.RecibirPaquetes.

	mux.HandleFunc("/mensaje", utils.RecibirMensaje)
	// Si se recibe una petición a la ruta "/mensaje",
	// se ejecuta la función utils.RecibirMensaje.

	// Iniciar el servidor HTTP en el puerto 8080

	err := http.ListenAndServe(":8080", mux)
	// ListenAndServe inicia un servidor HTTP en el puerto especificado.
	// El primer argumento es el puerto en el que se escucharán las peticiones,
	// y el segundo argumento es el manejador de peticiones, en este caso, mux.

	if err != nil {
		panic(err)
	}
	// Si hay un error al iniciar el servidor, se imprime el error
	// y se detiene el programa.

	// FIN DE ETAPA 3
}
