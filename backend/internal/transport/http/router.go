package http

import (
	"github.com/gin-gonic/gin"
	// Importar os handlers/controllers
	// Exemplo: "seu-nome-de-projeto/api/internal/transport/http"
	// Os controladores precisam ser criados (ex: user_handler.go, tenant_handler.go)
)

// O nome da sua pasta de projeto Go deve ser usado no import
const API_PREFIX = "/api" // Simplificando a lógica NODE_ENV ? '' : '/api'

// A função setupRoutes é chamada no main.go
func SetupRoutes(router *gin.Engine) {
	// --- Middlewares (os imports precisam ser configurados para os arquivos Go)
	// isAuthenticated := middleware.IsAuthenticated()

	// Simulação dos Controllers (Handlers em Go)
	authHandler := AuthHandler{}
	crudTenantHandler := CRUDTenantHandler{}
	registroHandler := RegistroHandler{}

	// --- Rotas Abertas (Sem Autenticação)

	// router.post(API + '/registro', new RegistroController().create)
	router.POST(API_PREFIX+"/registro", registroHandler.Create)

	// router.post(API + '/session', new AuthUserController().handle)
	router.POST(API_PREFIX+"/session", authHandler.Handle)

	// --- Rotas Protegidas (Exigindo o Middleware de Autenticação)
	// Simulamos a proteção: router.Group("/", isAuthenticated)
	// Na prática, você define o middleware aqui:
	// protected := router.Group(API_PREFIX, isAuthenticated)
	protected := router.Group(API_PREFIX)

	{
		// Registro
		// router.put(API + '/registro', isAuthenticated, new RegistroController().update)
		protected.PUT("/registro", registroHandler.Update)

		// Tenant
		// router.get(API + '/tenant', isAuthenticated, new CRUDTenantController().detail)
		protected.GET("/tenant", crudTenantHandler.Detail)

		// ... e assim por diante para todas as suas 30+ rotas ...
	}
}

// Para que a função setupRoutes possa ser chamada pelo main.go,
// ela precisa ser exportada. No entanto, se o main.go estiver no cmd/api
// e este arquivo estiver em internal/transport/http, você precisa
// ter certeza de que o `main.go` importa e chama esta função.

// No Go idiomático, a função principal estaria no pacote 'main'
// (como está), e a função de roteamento ficaria assim:
// func setupRoutes(router *gin.Engine) { ... }
// E no main.go, você teria: http.setupRoutes(router) (após ajustar os imports).
