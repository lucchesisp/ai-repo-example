package main

import (
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

// Response é a estrutura da resposta JSON
type Response struct {
	Message string `json:"message"`
	Status  string `json:"status"`
}

// customLoggingMiddleware é um middleware personalizado para log
func customLoggingMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()

		// Log da requisição
		log.Printf("Iniciando requisição: %s %s", c.Request.Method, c.Request.URL.Path)

		// Processa a requisição
		c.Next()

		// Log do tempo de resposta
		duration := time.Since(start)
		log.Printf("Requisição concluída: %s %s - Status: %d - Tempo: %v",
			c.Request.Method, c.Request.URL.Path, c.Writer.Status(), duration)
	}
}

// helloHandler para o endpoint /hello
func helloHandler(c *gin.Context) {
	response := Response{
		Message: "Olá! API funcionando perfeitamente com Gin!",
		Status:  "success",
	}

	c.JSON(http.StatusOK, response)
}

func main() {
	// Cria uma instância do Gin
	r := gin.Default()

	// Aplica o middleware personalizado
	r.Use(customLoggingMiddleware())

	// Define o endpoint
	r.GET("/hello", helloHandler)

	// Inicia o servidor
	port := ":8080"
	log.Printf("Servidor Gin rodando na porta %s", port)
	log.Println("Acesse: http://localhost:8080/hello")

	// Inicia o servidor
	r.Run(port)
}
