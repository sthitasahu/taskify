package v1


import(
	"github.com/labstack/echo/v4"
	"github.com/sthitasahu/tasker/internal/handler"
	"github.com/sthitasahu/tasker/internal/middleware"
)

func RegisterV1Routes(router *echo.Group,handlers *handler.Handlers,middleware *middleware.Middlewares){
   //register todo routes
   registerTodoRoutes(router,handlers.Todo,handlers.Comment,middleware.Auth)
   //register category routes
   registerCategoryRoutes(router,handlers.Category,middleware.Auth)
   //register comment routes
   registerCommentRoutes(router,handlers.Comment,middleware.Auth)
}