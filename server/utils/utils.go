package utils

import (
	"encoding/json"
	"log"
	"net/http"
)

type Mensaje struct {
	Mensaje string `json:"mensaje"`
}

type Paquete struct {
	Valores []string `json:"valores"`
}

func RecibirPaquetes(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var paquete Paquete
	err := decoder.Decode(&paquete)
	if err != nil {
		log.Printf("error al decodificar mensaje: %s\n", err.Error())
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("error al decodificar mensaje"))
		return
	}

	log.Println("me llego un paquete de un cliente")
	log.Printf("%+v\n", paquete)

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("ok"))
}

func RecibirMensaje(w http.ResponseWriter, r *http.Request) {
	// RecibirMensaje recibe 2 argumentos:
	// w es un http.ResponseWriter que se usa para enviar
	// la respuesta al cliente.
	// r es un *http.Request que contiene la información de la
	// petición HTTP recibida.

	decoder := json.NewDecoder(r.Body)
	// Se crea un nuevo decodificador de JSON
	// que usará el body del request recibido.

	var mensaje Mensaje
	// Se declara una variable mensaje de tipo struct Mensaje

	err := decoder.Decode(&mensaje)
	// El decoder decodifica el body del request
	// y lo guarda en la variable mensaje.

	if err != nil {
		log.Printf("Error al decodificar mensaje: %s\n", err.Error())
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Error al decodificar mensaje"))
		return
	}
	// Si hay un error al decodificar el mensaje,
	// se loggea un mensaje personalizado con la descripción del error,
	// WhiteHeader cambia el estado del ResponseWriter a 400 (Bad Request),
	// Write escribe una respuesta al cliente en el ResponseWriter.
	// Luego se sale de la función con return.

	log.Println("Me llegó un mensaje de un cliente")
	// Si no hay error, se loggea un mensaje personalizado
	// con salto de linea al final.

	log.Printf("%+v\n", mensaje)
	// Se loggea el mensaje recibido.
	// %+v imprime todos los campos del struct mensaje con sus valores.

	w.WriteHeader(http.StatusOK)
	// Se cambia el estado del ResponseWriter a 200 (OK).

	w.Write([]byte("ok"))
	// Se escribe la respuesta al cliente en el ResponseWriter.
}
