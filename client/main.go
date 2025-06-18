package main

import (
	"client/globals"
	"client/utils"
	"log"
)

func main() {
	utils.ConfigurarLogger()

	// ETAPA 2
	// LOGGING

	// loggear "Hola soy un log" usando la biblioteca log
	log.Println("Hola soy un log")

	// ARCHIVOS DE CONFIGURACION

	// Cargar archivo de configuracion como struct en la variable globals.ClientConfig
	globals.ClientConfig = utils.IniciarConfiguracion("config.json")

	// validar que la config este cargada correctamente
	if globals.ClientConfig == nil {
		log.Fatalf("No se pudo cargar la configuración")
	}
	// Fatalf imprime el mensaje que recibe y termina el programa

	// loggeamos el valor del MENSAJE de la config
	log.Println(globals.ClientConfig.Mensaje)

	// LEER DE LA CONSOLA
	utils.LeerConsola()

	// ADVERTENCIA: Antes de continuar, tenemos que asegurarnos que el servidor esté corriendo para poder conectarnos a él

	// enviar un mensaje al servidor con el valor de la config

	// leer de la consola el mensaje
	

	// generamos un paquete y lo enviamos al servidor
	// utils.GenerarYEnviarPaquete()
}
