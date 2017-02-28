package routers

import (
	"github.com/astaxie/beego"
)

func init() {

	beego.GlobalControllerRouter["github.com/matthew/beego/lessonsbasketapi/controllers:AdminsController"] = append(beego.GlobalControllerRouter["github.com/matthew/beego/lessonsbasketapi/controllers:AdminsController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/matthew/beego/lessonsbasketapi/controllers:AdminsController"] = append(beego.GlobalControllerRouter["github.com/matthew/beego/lessonsbasketapi/controllers:AdminsController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/matthew/beego/lessonsbasketapi/controllers:AdminsController"] = append(beego.GlobalControllerRouter["github.com/matthew/beego/lessonsbasketapi/controllers:AdminsController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/matthew/beego/lessonsbasketapi/controllers:AdminsController"] = append(beego.GlobalControllerRouter["github.com/matthew/beego/lessonsbasketapi/controllers:AdminsController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/matthew/beego/lessonsbasketapi/controllers:AdminsController"] = append(beego.GlobalControllerRouter["github.com/matthew/beego/lessonsbasketapi/controllers:AdminsController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/matthew/beego/lessonsbasketapi/controllers:AnswersController"] = append(beego.GlobalControllerRouter["github.com/matthew/beego/lessonsbasketapi/controllers:AnswersController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/matthew/beego/lessonsbasketapi/controllers:AnswersController"] = append(beego.GlobalControllerRouter["github.com/matthew/beego/lessonsbasketapi/controllers:AnswersController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/matthew/beego/lessonsbasketapi/controllers:AnswersController"] = append(beego.GlobalControllerRouter["github.com/matthew/beego/lessonsbasketapi/controllers:AnswersController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/matthew/beego/lessonsbasketapi/controllers:AnswersController"] = append(beego.GlobalControllerRouter["github.com/matthew/beego/lessonsbasketapi/controllers:AnswersController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/matthew/beego/lessonsbasketapi/controllers:AnswersController"] = append(beego.GlobalControllerRouter["github.com/matthew/beego/lessonsbasketapi/controllers:AnswersController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/matthew/beego/lessonsbasketapi/controllers:ImagesController"] = append(beego.GlobalControllerRouter["github.com/matthew/beego/lessonsbasketapi/controllers:ImagesController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/matthew/beego/lessonsbasketapi/controllers:ImagesController"] = append(beego.GlobalControllerRouter["github.com/matthew/beego/lessonsbasketapi/controllers:ImagesController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/matthew/beego/lessonsbasketapi/controllers:ImagesController"] = append(beego.GlobalControllerRouter["github.com/matthew/beego/lessonsbasketapi/controllers:ImagesController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/matthew/beego/lessonsbasketapi/controllers:ImagesController"] = append(beego.GlobalControllerRouter["github.com/matthew/beego/lessonsbasketapi/controllers:ImagesController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/matthew/beego/lessonsbasketapi/controllers:ImagesController"] = append(beego.GlobalControllerRouter["github.com/matthew/beego/lessonsbasketapi/controllers:ImagesController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/matthew/beego/lessonsbasketapi/controllers:LessonsController"] = append(beego.GlobalControllerRouter["github.com/matthew/beego/lessonsbasketapi/controllers:LessonsController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/matthew/beego/lessonsbasketapi/controllers:LessonsController"] = append(beego.GlobalControllerRouter["github.com/matthew/beego/lessonsbasketapi/controllers:LessonsController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/matthew/beego/lessonsbasketapi/controllers:LessonsController"] = append(beego.GlobalControllerRouter["github.com/matthew/beego/lessonsbasketapi/controllers:LessonsController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/matthew/beego/lessonsbasketapi/controllers:LessonsController"] = append(beego.GlobalControllerRouter["github.com/matthew/beego/lessonsbasketapi/controllers:LessonsController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/matthew/beego/lessonsbasketapi/controllers:LessonsController"] = append(beego.GlobalControllerRouter["github.com/matthew/beego/lessonsbasketapi/controllers:LessonsController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/matthew/beego/lessonsbasketapi/controllers:OptionsController"] = append(beego.GlobalControllerRouter["github.com/matthew/beego/lessonsbasketapi/controllers:OptionsController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/matthew/beego/lessonsbasketapi/controllers:OptionsController"] = append(beego.GlobalControllerRouter["github.com/matthew/beego/lessonsbasketapi/controllers:OptionsController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/matthew/beego/lessonsbasketapi/controllers:OptionsController"] = append(beego.GlobalControllerRouter["github.com/matthew/beego/lessonsbasketapi/controllers:OptionsController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/matthew/beego/lessonsbasketapi/controllers:OptionsController"] = append(beego.GlobalControllerRouter["github.com/matthew/beego/lessonsbasketapi/controllers:OptionsController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/matthew/beego/lessonsbasketapi/controllers:OptionsController"] = append(beego.GlobalControllerRouter["github.com/matthew/beego/lessonsbasketapi/controllers:OptionsController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/matthew/beego/lessonsbasketapi/controllers:ScreensController"] = append(beego.GlobalControllerRouter["github.com/matthew/beego/lessonsbasketapi/controllers:ScreensController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/matthew/beego/lessonsbasketapi/controllers:ScreensController"] = append(beego.GlobalControllerRouter["github.com/matthew/beego/lessonsbasketapi/controllers:ScreensController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/matthew/beego/lessonsbasketapi/controllers:ScreensController"] = append(beego.GlobalControllerRouter["github.com/matthew/beego/lessonsbasketapi/controllers:ScreensController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/matthew/beego/lessonsbasketapi/controllers:ScreensController"] = append(beego.GlobalControllerRouter["github.com/matthew/beego/lessonsbasketapi/controllers:ScreensController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/matthew/beego/lessonsbasketapi/controllers:ScreensController"] = append(beego.GlobalControllerRouter["github.com/matthew/beego/lessonsbasketapi/controllers:ScreensController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/matthew/beego/lessonsbasketapi/controllers:UserScoreController"] = append(beego.GlobalControllerRouter["github.com/matthew/beego/lessonsbasketapi/controllers:UserScoreController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/matthew/beego/lessonsbasketapi/controllers:UserScoreController"] = append(beego.GlobalControllerRouter["github.com/matthew/beego/lessonsbasketapi/controllers:UserScoreController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/matthew/beego/lessonsbasketapi/controllers:UserScoreController"] = append(beego.GlobalControllerRouter["github.com/matthew/beego/lessonsbasketapi/controllers:UserScoreController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/matthew/beego/lessonsbasketapi/controllers:UserScoreController"] = append(beego.GlobalControllerRouter["github.com/matthew/beego/lessonsbasketapi/controllers:UserScoreController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/matthew/beego/lessonsbasketapi/controllers:UserScoreController"] = append(beego.GlobalControllerRouter["github.com/matthew/beego/lessonsbasketapi/controllers:UserScoreController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/matthew/beego/lessonsbasketapi/controllers:UsersController"] = append(beego.GlobalControllerRouter["github.com/matthew/beego/lessonsbasketapi/controllers:UsersController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/matthew/beego/lessonsbasketapi/controllers:UsersController"] = append(beego.GlobalControllerRouter["github.com/matthew/beego/lessonsbasketapi/controllers:UsersController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/matthew/beego/lessonsbasketapi/controllers:UsersController"] = append(beego.GlobalControllerRouter["github.com/matthew/beego/lessonsbasketapi/controllers:UsersController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/matthew/beego/lessonsbasketapi/controllers:UsersController"] = append(beego.GlobalControllerRouter["github.com/matthew/beego/lessonsbasketapi/controllers:UsersController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/matthew/beego/lessonsbasketapi/controllers:UsersController"] = append(beego.GlobalControllerRouter["github.com/matthew/beego/lessonsbasketapi/controllers:UsersController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

}
