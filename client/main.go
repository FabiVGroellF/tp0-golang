package main

import (
	"client/globals"
	"client/utils"
	"log"
)

func main() {
	utils.ConfigurarLogger()

	// ETAPA 2 - COMANDOS BASICOS

	// LOGGING

	// Loggear "Hola soy un log" usando la biblioteca log
	log.Println("Hola soy un log")

	// ARCHIVOS DE CONFIGURACION

	// Cargar archivo de configuracion como struct en la variable globals.ClientConfig
	globals.ClientConfig = utils.IniciarConfiguracion("config.json")

	// Validar que la Config este cargada correctamente
	if globals.ClientConfig == nil {
		log.Fatalf("No se pudo cargar la configuración")
	}
	// Fatalf imprime el mensaje que recibe y termina el programa

	// Loggear el valor del MENSAJE de la Config
	log.Println(globals.ClientConfig.Mensaje)

	// FIN DE ETAPA 2

	// ETAPA 3 - CONEXION CLIENTE-SERVIDOR

	// ADVERTENCIA:
	// Antes de continuar, tenemos que asegurarnos que el servidor 
	// esté corriendo para poder conectarnos a él.

	// Enviar como mensaje al servidor el valor CLAVE de la Config
	utils.EnviarMensaje(globals.ClientConfig.Ip, globals.ClientConfig.Puerto, globals.ClientConfig.Clave)

	// Generar un paquete y enviarlo al servidor
	utils.GenerarYEnviarPaquete()

	// FIN DE ETAPA 3
}
