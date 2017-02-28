//my api controllers
package controllers

import (
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/matthew/beego/lessonsbasketapi/models"

	"github.com/astaxie/beego"

	"github.com/matthew/beego/lessonsbasketapi/tools"
)

type LessonScreenController struct {
	beego.Controller
}

type LessonScreenOptionController struct {
	beego.Controller
}

type LessonAnswerController struct {
	beego.Controller
}

type AnswerResult struct {
	Result string "json:'result'"
}

type LoginController struct {
	beego.Controller
}

type LoginResult struct {
	Result string "json:result"
	Token  string "json:token"
}

func (c *LessonScreenController) GetScreensByLesson() {
	idStr := c.Ctx.Input.Param(":id") //lesson id
	id, _ := strconv.Atoi(idStr)
	v, err := models.GetScreensByLessonId(id)
	if err != nil {
		c.Data["json"] = err.Error()
	} else {
		c.Data["json"] = v
	}
	c.ServeJSON()
}

func (c *LessonScreenController) GetScreenByLessonAndScreen() {
	lidStr := c.Ctx.Input.Param(":lid") //lesson id
	lid, _ := strconv.Atoi(lidStr)
	sidStr := c.Ctx.Input.Param(":sid") //screen id
	sid, _ := strconv.Atoi(sidStr)
	v, err := models.GetScreenByLessonAndScreenId(lid, sid)
	if err != nil {
		c.Data["json"] = err.Error()
	} else {
		c.Data["json"] = v
	}
	c.ServeJSON()
}

func (c *LessonScreenController) GetScreenByLessonAndPosition() {
	idStr := c.Ctx.Input.Param(":id") //lesson id
	id, _ := strconv.Atoi(idStr)
	positionStr := c.Ctx.Input.Param(":position") //position
	position, _ := strconv.Atoi(positionStr)
	v, err := models.GetScreenByLessonAndPositionId(id, position)
	if err != nil {
		c.Data["json"] = err.Error()
	} else {
		c.Data["json"] = v
	}
	c.ServeJSON()
}

func (c *LessonScreenController) GetNextScreenByScreen() {
	lidStr := c.Ctx.Input.Param(":lid") //lesson id
	lid, _ := strconv.Atoi(lidStr)
	sidStr := c.Ctx.Input.Param(":sid") //screen id
	sid, _ := strconv.Atoi(sidStr)
	sid++
	v, err := models.GetScreenByLessonAndScreenId(lid, sid)
	if err != nil {
		c.Data["json"] = err.Error()
	} else {
		c.Data["json"] = v
	}
	c.ServeJSON()
}

func (c *LessonScreenOptionController) GetOptionsByLessonAndScreen() {
	lidStr := c.Ctx.Input.Param(":lid") //lesson id
	lid, _ := strconv.Atoi(lidStr)
	sidStr := c.Ctx.Input.Param(":sid") //screen id
	sid, _ := strconv.Atoi(sidStr)
	v, err := models.GetOptionsByLessonAndScreenId(lid, sid)
	if err != nil {
		c.Data["json"] = err.Error()
	} else {
		c.Data["json"] = v
	}
	c.ServeJSON()
}

func (c *LessonAnswerController) PostLessonAnswer() {
	var v models.Answers
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &v); err == nil {
		if result, err := models.CheckAnswer(&v); err == nil {
			c.Ctx.Output.SetStatus(201)
			var jsonResult AnswerResult
			if result == true {
				jsonResult.Result = "success"
			} else {
				jsonResult.Result = "fail"
			}
			c.Data["json"] = result
		} else {
			c.Data["json"] = err.Error()
		}
	} else {
		c.Data["json"] = err.Error()
	}
	c.ServeJSON()
}

func (c *LoginController) Login() {
	user := &models.Users{}
	user.Email = c.GetString("email", "")
	user.Password = c.GetString("password", "")
	err := models.ValidateUser(user)
	var loginresult LoginResult
	if err == nil {
		token := tools.GenerateToken()
		c.SetSession("token", token)
		fmt.Println("email:", c.GetSession("token"))
		loginresult.Result = "success"
		loginresult.Token = token
	} else {
		loginresult.Result = "fail"
		loginresult.Token = ""
	}
	c.Data["json"] = loginresult
	c.ServeJSON()
}
