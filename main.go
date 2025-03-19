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
	
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error al cargar el archivo .env: %v", err)
	}

	
	rabbitConn, err := adapters.NewRabbitMQConnection()
	if err != nil {
		log.Fatalf("Error al conectar con RabbitMQ: %v", err)
	}

	
	r := gin.Default()

	
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:5173"}, 
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	
	pagosInfra.SetupRoutesPago(r, rabbitConn)

	
	if err := r.Run(":8080"); err != nil {
		log.Fatalf("Error al iniciar el servidor: %v", err)
	} 
}
