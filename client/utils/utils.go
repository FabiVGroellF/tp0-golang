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
	// y devuelve un puntero a una estructura Config.

	var config *globals.Config
	// Inicializamos la variable config para que tome el formato
	// de la estructura globals.Config.
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
	// que leerá el archivo configFile y se asigna a jsonParser,
	// siendo configFile el archivo de configuración que se abrió previamente.

	jsonParser.Decode(&config) 
	// Decode hace que el jsonParser decodifique el contenido del archivo
	// abierto y vaya llenando la estructura config con los datos decodificados.

	return config
	// Finalmente, se devuelve el puntero a la estructura config
	// que contiene los datos del archivo de configuración.
}

func LeerConsola() Paquete {
	// Lee la consola hasta que el usuario ingrese un salto de línea
	// y devuelve un paquete de tipo struct Paquete.

	paquete := Paquete{
		Valores: []string{},
	}
	// Inicializamos un nuevo paquete de tipo struct Paquete.
	// El campo Valores es un slice de strings que se inicializa como vacío.

	reader := bufio.NewReader(os.Stdin)
	// NewReader crea un nuevo lector de búfer (bufio)
	// el cual lee desde la entrada estándar (os.Stdin).
	// En este caso, la entrada estándar es la consola.

	log.Println("Ingrese los mensajes")
	// Se loggea un mensaje pidiendo al usuario que ingrese los mensajes.
	// Println agrega un salto de línea al final del mensaje.

	for {
		// Bucle while para leer mensajes de la consola
		// mientras el usuario no ingrese un salto de línea.

		text, _ := reader.ReadString('\n')
		// Se guarda en text la cadena de texto leída desde la consola.
		// ReadString lee hasta que encuentra el carácter especificado
		// En este caso, lee hasta que encuentra un salto de línea ('\n').

		if text == "\n" {
			break
			// Si el texto ingresado es un salto de línea,
			// se sale del bucle y no se loggea.

		} else {
			text = text[:len(text)-1]
			// Se elimina el último carácter de text, que es el salto de línea.

			log.Print(text)
			// Se loggea el texto ingresado por el usuario.
			// No se agrega un salto de línea al final del mensaje.

			paquete.Valores = append(paquete.Valores, text)
			// Se agrega el texto ingresado al slice Valores del paquete.
		}
	}
	return paquete
}

func GenerarYEnviarPaquete() {
	paquete := LeerConsola()
	// Se lee de la consola y se carga lo leído en el paquete.

	log.Printf("Paquete a enviar: %+v", paquete)
	// Se loggea el paquete que se va a enviar.

	EnviarPaquete(globals.ClientConfig.Ip, globals.ClientConfig.Puerto, paquete)
	// Se envía el paquete.
}

func EnviarMensaje(ip string, puerto int, mensajeTxt string) {
	// Recibe la IP, el puerto y el mensaje a enviar.

	mensaje := Mensaje{Mensaje: mensajeTxt}
	// Se crea un nuevo mensaje de tipo struct Mensaje
	// donde el campo Mensaje se inicializa con el mensajeTxt recibido.

	body, err := json.Marshal(mensaje)
	// json.Marshal convierte el argumento a formato JSON.
	// Si no hay errores, se guarda el resultado en body.
	// Si hay un error, se guarda en err.

	if err != nil {
		log.Printf("Error codificando mensaje: %s", err.Error())
	}
	// Si hay error, se loggea un mensaje personalizado
	// que incluye la descripción del error.
	// err.Error() devuelve una cadena que describe el error.

	url := fmt.Sprintf("http://%s:%d/mensaje", ip, puerto)
	// Se construye la URL a la que se enviará el mensaje.
	// Sprintf permite imprimir usando placeholders.
	// La dirección "/mensaje" es la dirección de la ruta
	// donde el servidor espera recibir mensajes.

	resp, err := http.Post(url, "application/json", bytes.NewBuffer(body))
	// http.Post se usar para enviar datos a un servidor HTTP.
	// Recibe la URL por donde se enviará el mensaje,
	// recibe además el tipo de contenido, en este caso "application/json",
	// que significa que son datos en formato JSON,
	// y tambien recibe un buffer de bytes donde se guarda
	// el contenido de body, que es el mensaje codificado en JSON.
	// Si la petición se envía correctamente, se guarda la respuesta
	// que manda el servidor en la variable resp.
	// Si hay un error al enviar la petición, se guarda en err.

	if err != nil {
		log.Printf("Error enviando mensaje a ip:%s puerto:%d", ip, puerto)
	}
	// Si hay un error al enviar el mensaje, se loggea un mensaje personalizado
	// que incluye la IP y el puerto a los que se intentó enviar el mensaje.

	log.Printf("Respuesta del servidor: %s", resp.Status)
	// Se loggea un mensaje personalizado con el estado de la respuesta
	// del servidor a la petición realizada.
}

func EnviarPaquete(ip string, puerto int, paquete Paquete) {
	// Recibe la IP, el puerto y el paquete a enviar.

	body, err := json.Marshal(paquete)
	// json.Marshal convierte el paquete a formato JSON.
	// Si no hay errores, se guarda el resultado en body.
	// Si hay un error, se guarda en err.

	if err != nil {
		log.Printf("Error codificando paquete: %s", err.Error())
	}
	// Si hay error, se loggea un mensaje personalizado
	// que incluye la descripción del error.
	// err.Error() devuelve una cadena que describe el error.

	url := fmt.Sprintf("http://%s:%d/paquetes", ip, puerto)
	// Se construye la URL a la que se enviará el paquete.
	// Sprintf permite imprimir usando placeholders.
	// La dirección "/paquetes" es la dirección de la ruta
	// donde el servidor espera recibir paquetes.

	resp, err := http.Post(url, "application/json", bytes.NewBuffer(body))
	// http.Post se usar para enviar datos a un servidor HTTP.
	// Recibe la URL por donde se enviará el paquete,
	// recibe además el tipo de contenido, en este caso "application/json",
	// que significa que son datos en formato JSON,
	// y también recibe un buffer de bytes donde se guarda
	// el contenido de body, que es el paquete codificado en JSON.
	// Si la petición se envía correctamente, se guarda la respuesta
	// que manda el servidor en la variable resp.
	// Si hay un error al enviar la petición, se guarda en err.

	if err != nil {
		log.Printf("Error enviando paquete a ip:%s puerto:%d", ip, puerto)
	}
	// Si hay un error al enviar el paquete, se loggea un mensaje personalizado
	// que incluye la IP y el puerto a los que se intentó enviar el paquete.

	log.Printf("Respuesta del servidor: %s", resp.Status)
	// Se loggea un mensaje personalizado con el estado de la respuesta
	// del servidor a la petición realizada.
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
