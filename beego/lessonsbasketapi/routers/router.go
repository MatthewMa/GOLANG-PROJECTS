// @APIVersion 1.0.0
// @Title beego Test API
// @Description beego has a very cool tools to autogenerate documents for your API
// @Contact astaxie@gmail.com
// @TermsOfServiceUrl http://beego.me/
// @License Apache 2.0
// @LicenseUrl http://www.apache.org/licenses/LICENSE-2.0.html
package routers

import (
	"github.com/matthew/beego/lessonsbasketapi/controllers"

	"github.com/astaxie/beego"
)

func init() {
	ns := beego.NewNamespace("/v1",

		beego.NSNamespace("/admins",
			beego.NSInclude(
				&controllers.AdminsController{},
			),
		),

		beego.NSNamespace("/answers",
			beego.NSInclude(
				&controllers.AnswersController{},
			),
		),

		beego.NSNamespace("/images",
			beego.NSInclude(
				&controllers.ImagesController{},
			),
		),

		beego.NSNamespace("/lessons",
			beego.NSInclude(
				&controllers.LessonsController{},
			),
		),

		beego.NSNamespace("/options",
			beego.NSInclude(
				&controllers.OptionsController{},
			),
		),

		beego.NSNamespace("/screens",
			beego.NSInclude(
				&controllers.ScreensController{},
			),
		),

		beego.NSNamespace("/users_score",
			beego.NSInclude(
				&controllers.UserScoreController{},
			),
		),

		beego.NSNamespace("/users",
			beego.NSInclude(
				&controllers.UsersController{},
			),
		),
	)
	beego.AddNamespace(ns)
	beego.Router("/v1/lessons/:id([0-9]+)/screens", &controllers.LessonScreenController{}, "get,post:GetScreensByLesson")
	beego.Router("/v1/lessons/:lid([0-9]+)/screens/:sid([0-9]+)", &controllers.LessonScreenController{}, "get,post:GetScreenByLessonAndScreen")
	beego.Router("/v1/lessons/:id([0-9]+)/screens/position/:position([0-9]+)", &controllers.LessonScreenController{}, "get,post:GetScreenByLessonAndPosition")
	beego.Router("/v1/lessons/:lid([0-9]+)/screens/nextscreen/sid([0-9]+)", &controllers.LessonScreenController{}, "get,post:GetNextScreenByScreen")
	beego.Router("/v1/lessons/:lid([0-9]+)/screens/:sid([0-9]+)/options", &controllers.LessonScreenOptionController{}, "get,post:GetOptionsByLessonAndScreen")
	beego.Router("/v1/lessons/:lid([0-9]+)/answers", &controllers.LessonAnswerController{}, "post:PostLessonAnswer")
	beego.Router("/auth/login", &controllers.LoginController{}, "post:Login")
	beego.Router("/auth/profile/:token", &controllers.LoginController{}, "get:GetProfile")
	beego.Router("/auth/logout/:token", &controllers.LoginController{}, "get:Logout")
	beego.Router("/auth/register", &controllers.LoginController{}, "post:Register")
}
