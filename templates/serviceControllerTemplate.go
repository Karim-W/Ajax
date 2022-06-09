package templates

var SvcTemplate = `package services

type {{.Name}}Service interface {
}

type {{.name}}ServiceImpl struct {

}

//=============================================	   Constructor and DI		========================================================
var _ {{.Name}}Service = (*{{.name}}ServiceImpl)(nil)

 func {{.Name}}ServiceProvider() *{{.name}}ServiceImpl {
	return &{{.name}}ServiceImpl{}
 }

 var {{.Name}}ServiceDependency = fx.Option(fx.Provide({{.Name}}ServiceProvider))

 //=============================================	 	SVC Functions		========================================================


`
