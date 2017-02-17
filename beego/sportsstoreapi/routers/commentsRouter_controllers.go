package routers

import (
	"github.com/astaxie/beego"
)

func init() {

	beego.GlobalControllerRouter["github.com/matthew/beego/sportsstoreapi/controllers:PostsController"] = append(beego.GlobalControllerRouter["github.com/matthew/beego/sportsstoreapi/controllers:PostsController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/matthew/beego/sportsstoreapi/controllers:PostsController"] = append(beego.GlobalControllerRouter["github.com/matthew/beego/sportsstoreapi/controllers:PostsController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/matthew/beego/sportsstoreapi/controllers:PostsController"] = append(beego.GlobalControllerRouter["github.com/matthew/beego/sportsstoreapi/controllers:PostsController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/matthew/beego/sportsstoreapi/controllers:PostsController"] = append(beego.GlobalControllerRouter["github.com/matthew/beego/sportsstoreapi/controllers:PostsController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/matthew/beego/sportsstoreapi/controllers:PostsController"] = append(beego.GlobalControllerRouter["github.com/matthew/beego/sportsstoreapi/controllers:PostsController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/matthew/beego/sportsstoreapi/controllers:ProductsController"] = append(beego.GlobalControllerRouter["github.com/matthew/beego/sportsstoreapi/controllers:ProductsController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/matthew/beego/sportsstoreapi/controllers:ProductsController"] = append(beego.GlobalControllerRouter["github.com/matthew/beego/sportsstoreapi/controllers:ProductsController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/matthew/beego/sportsstoreapi/controllers:ProductsController"] = append(beego.GlobalControllerRouter["github.com/matthew/beego/sportsstoreapi/controllers:ProductsController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/matthew/beego/sportsstoreapi/controllers:ProductsController"] = append(beego.GlobalControllerRouter["github.com/matthew/beego/sportsstoreapi/controllers:ProductsController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/matthew/beego/sportsstoreapi/controllers:ProductsController"] = append(beego.GlobalControllerRouter["github.com/matthew/beego/sportsstoreapi/controllers:ProductsController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

}
