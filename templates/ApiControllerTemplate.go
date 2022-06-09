package templates

var APIControllerGin = `package controllers
import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/karim-w/httputils"
	"go.uber.org/fx"
)

type {{.Name}}Controller interface {
	//router
	SetupRoutes(r *gin.RouterGroup)
}

type {{.name}}ControllerImpl struct {

}
//=============================================	   Constructor and DI		========================================================

var _ {{.Name}}Controller = (*{{.name}}ControllerImpl)(nil)

func {{.Name}}ControllerProvider() *{{.name}}ControllerImpl {
	return &{{.name}}ControllerImpl{}
}

var {{.Name}}ControllerDependency = fx.Option(fx.Provide({{.Name}}ControllerProvider))

//=============================================	  	 Router Functions		========================================================

func (c *{{.name}}ControllerImpl) SetupRoutes(r *gin.RouterGroup) {

}

//============================================= 		Controller Functions	========================================================

`
