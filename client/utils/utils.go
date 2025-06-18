package utils

import (
	"bufio"
	"bytes"
	"client/globals"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

type Mensaje struct {
	Mensaje string `json:"mensaje"`
}

type Paquete struct {
	Valores []string `json:"valores"`
}

func IniciarConfiguracion(filePath string) *globals.Config {
	// Recibe el path del archivo de configuración
	// y devuelve un puntero a una estructura Config

	var config *globals.Config
	// Inicializamos la variable config para que tome el formato
	// de la estructura globals.Config
	// Es decir, un struct que guarda ip, puerto y mensaje.

	configFile, err := os.Open(filePath)
	// os.Open abre el archivo especificado por filePath.
	// Si el archivo se abre correctamente, se asigna a la variable configFile.
	// Si el archivo no existe o hay un error al abrirlo,
	// se asigna el error a la variable err.

	if err != nil {
		log.Fatal(err.Error())
	}
	// Si hay error, el log mostrara el error
	// y al ser fatal, el programa se detendrá.
	// err.Error() devuelve una cadena que describe el error.

	defer configFile.Close()
	// defer se usa para asegurarse de que el archivo se cierre
	// al final de la función, independientemente de si ocurre un error o no.

	jsonParser := json.NewDecoder(configFile)
	// Se crea un decoder con la función json.NewDecoder
	// que leerá el archivo configFile y se asigna a jsonParser.
	// siendo configFile el archivo de configuración que se abrió previamente.

	jsonParser.Decode(&config) 
	// Decode hace que el jsonParser decodifique el contenido del archivo abierto
	// y vaya llenando la estructura config con los datos decodificados.

	return config
	// Finalmente, se devuelve el puntero a la estructura config
	// que contiene los datos del archivo de configuración.
}

func LeerConsola() {
	// Leer de la consola
	reader := bufio.NewReader(os.Stdin)
	// Se guarda en reader un nuevo lector de búfer
	// NewReader crea un nuevo lector de búfer
	// el cual lee desde la entrada estándar (os.Stdin).
	// En este caso, la entrada estándar es la consola.

	log.Println("Ingrese los mensajes")
	// Se loggea un mensaje pidiendo al usuario que ingrese los mensajes.
	// Println agrega un salto de línea al final del mensaje.

	var text string

	for {
		text, _ = reader.ReadString('\n')
		// Se guarda en text la cadena de texto leída desde la consola.
		// ReadString lee hasta que encuentra el carácter especificado
		// En este caso, lee hasta que encuentra un salto de línea ('\n').
		
		if text == "\n" {
			break
			// Si el texto ingresado es un salto de línea,
			// se sale del bucle y no se loggea.
		} else {
			log.Print(text)
			// Se loggea el texto ingresado por el usuario.
			// No se agrega un salto de línea al final del mensaje.
		}
	}
}

func GenerarYEnviarPaquete() {
	paquete := Paquete{}
	// Leemos y cargamos el paquete

	log.Printf("paqute a enviar: %+v", paquete)
	// Enviamos el paqute
}

func EnviarMensaje(ip string, puerto int, mensajeTxt string) {
	mensaje := Mensaje{Mensaje: mensajeTxt}
	body, err := json.Marshal(mensaje)
	if err != nil {
		log.Printf("error codificando mensaje: %s", err.Error())
	}

	url := fmt.Sprintf("http://%s:%d/mensaje", ip, puerto)
	resp, err := http.Post(url, "application/json", bytes.NewBuffer(body))
	if err != nil {
		log.Printf("error enviando mensaje a ip:%s puerto:%d", ip, puerto)
	}

	log.Printf("respuesta del servidor: %s", resp.Status)
}

func EnviarPaquete(ip string, puerto int, paquete Paquete) {
	body, err := json.Marshal(paquete)
	if err != nil {
		log.Printf("error codificando mensajes: %s", err.Error())
	}

	url := fmt.Sprintf("http://%s:%d/paquetes", ip, puerto)
	resp, err := http.Post(url, "application/json", bytes.NewBuffer(body))
	if err != nil {
		log.Printf("error enviando mensajes a ip:%s puerto:%d", ip, puerto)
	}

	log.Printf("respuesta del servidor: %s", resp.Status)
}

func ConfigurarLogger() {
	logFile, err := os.OpenFile("tp0.log", os.O_CREATE|os.O_APPEND|os.O_RDWR, 0666)
	// os.O_CREATE: // Crea el archivo si no existe.
	// os.O_APPEND: // Agrega al final del archivo.
	// os.O_RDWR: // Abre el archivo para lectura y escritura.
	// 0666: Permisos de lectura y escritura para todos los usuarios.
	// Si no hay errores, se asigna el archivo a logFile.
	// Si hay un error, se guarda el mismo en la variable err.

	if err != nil {
		panic(err)
	}
	// Si hay error, panic imprime el error y termina la ejecución del programa.

	mw := io.MultiWriter(os.Stdout, logFile)
	// Se guarda en mw un MultiWriter que escribe tanto 
	// en la salida estándar (os.Stdout) como en el archivo logFile.

	log.SetOutput(mw)
	// SetOutput hace que todos los mensajes de log 
	// se escriban usando el MultiWriter mw.
	// Es decir, de ahora en adelante, 
	// todos los log se mostrarán tanto en la consola 
	// y se guardarán en el archivo de logs "tp0.log".
}
