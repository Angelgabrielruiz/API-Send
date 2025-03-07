package main

import (
	"log"
	"time"
	"github.com/joho/godotenv"
	"github.com/gin-gonic/gin"
	"github.com/gin-contrib/cors"
	pagosInfra "Send/src/Pagos/infraestructure"
	"Send/src/Pagos/infraestructure/adapters"
)

func main() {
	// Cargar archivo .env
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error al cargar el archivo .env: %v", err)
	}

	// Conectar a RabbitMQ
	rabbitConn, err := adapters.NewRabbitMQConnection()
	if err != nil {
		log.Fatalf("Error al conectar con RabbitMQ: %v", err)
	}

	// Configurar el servidor Gin
	r := gin.Default()

	// Configurar CORS
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:5173"}, 
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	// Configurar las rutas
	pagosInfra.SetupRoutesPago(r, rabbitConn)

	// Iniciar el servidor en el puerto 8080
	if err := r.Run(":8080"); err != nil {
		log.Fatalf("Error al iniciar el servidor: %v", err)
	} 
}
