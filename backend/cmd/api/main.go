package main

import (
	"log"
	"os"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"

	. "vecontabgovue/backend/internal/transport/http"
)

// Inicia o servidor e carrega as configurações
func main() {
	// 1. Carregar Variáveis de Ambiente (.env)
	err := godotenv.Load()
	if err != nil {
		log.Println("Aviso: Não foi possível carregar o arquivo .env")
	}

	// 2. Configurar o Gin (pode ser gin.ReleaseMode em produção)
	router := gin.Default()

	// 3. Configuração do CORS (Substitui app.use(cors()))
	// É uma configuração simples, adapte conforme a necessidade real de produção
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"}, // Idealmente, liste domínios específicos em produção
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	// 4. Configurar Rotas (Substitui app.use(router) no server.ts)
	// Aqui chamamos a função que define todas as rotas (que criaremos em router.go)
	SetupRoutes(router)

	// 5. Configurar o Middleware de Tratamento de Erros Global (Substitui o app.use((err: Error, ...)))
	// Em Go, o tratamento de erro é mais granular (por handler), mas podemos ter um middleware
	// customizado antes de iniciar o listen. O gin.Default já tem um tratamento básico.

	// 6. Iniciar o Servidor
	serverPort := os.Getenv("SERVER_PORT")
	if serverPort == "" {
		serverPort = "8080" // Default
	}

	log.Printf("Servidor em execução no porto: %s\n", serverPort)
	if err := router.Run(":" + serverPort); err != nil {
		log.Fatalf("Falha ao iniciar o servidor: %v", err)
	}
}
