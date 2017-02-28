//my api controllers
package controllers

import (
	"encoding/json"
	"errors"
	"fmt"
	"strconv"

	"github.com/matthew/beego/lessonsbasketapi/models"

	"github.com/astaxie/beego"

	"github.com/matthew/beego/lessonsbasketapi/tools"

	"github.com/astaxie/beego/validation"
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

type JsonError struct {
	ErrorStatus string "json:result"
}

type JsonResult struct {
	Result string "json:result"
}

func (c *LessonScreenController) GetScreensByLesson() {
	idStr := c.Ctx.Input.Param(":id") //lesson id
	id, _ := strconv.Atoi(idStr)
	v, err := models.GetScreensByLessonId(id)
	if err != nil {
		jsonerror := &JsonError{
			ErrorStatus: err.Error(),
		}
		c.Data["json"] = jsonerror
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
		jsonerror := &JsonError{
			ErrorStatus: err.Error(),
		}
		c.Data["json"] = jsonerror
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
		jsonerror := &JsonError{
			ErrorStatus: err.Error(),
		}
		c.Data["json"] = jsonerror
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
		jsonerror := &JsonError{
			ErrorStatus: err.Error(),
		}
		c.Data["json"] = jsonerror
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
		jsonerror := &JsonError{
			ErrorStatus: err.Error(),
		}
		c.Data["json"] = jsonerror
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
			jsonerror := &JsonError{
				ErrorStatus: err.Error(),
			}
			c.Data["json"] = jsonerror
		}
	} else {
		jsonerror := &JsonError{
			ErrorStatus: err.Error(),
		}
		c.Data["json"] = jsonerror
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
		//update token
		user.Token = token
		err = models.UpdateUsersById(user)
		if err == nil {
			loginresult.Result = "success"
			loginresult.Token = token
		} else {
			loginresult.Result = "fail"
			loginresult.Token = ""
		}

	} else {
		loginresult.Result = "fail"
		loginresult.Token = ""
	}
	c.Data["json"] = loginresult
	c.ServeJSON()
}

func (c *LoginController) GetProfile() {
	token := c.Ctx.Input.Param(":token")
	sessionToken := c.GetSession("token")
	if sessionToken == nil {
		//session expires
		jsonerror := &JsonError{
			ErrorStatus: errors.New("Token has expired,please login again!").Error(),
		}
		c.Data["json"] = jsonerror
	} else {
		if token == sessionToken.(string) {
			//Get profile
			profile, err := models.GetProfileByToken(token)
			if err == nil {
				c.Data["json"] = profile
			} else {
				jsonerror := &JsonError{
					ErrorStatus: err.Error(),
				}
				c.Data["json"] = jsonerror
			}
		} else {
			jsonerror := &JsonError{
				ErrorStatus: errors.New("Token is not correct,please login again!").Error(),
			}
			c.Data["json"] = jsonerror
		}
	}
	c.ServeJSON()
}

func (c *LoginController) Logout() {
	token := c.Ctx.Input.Param(":token")
	sessionToken := c.GetSession("token")
	if sessionToken == nil {
		//session expires
		jsonerror := &JsonResult{
			Result: "fail",
		}
		c.Data["json"] = jsonerror
	} else {
		if token == sessionToken.(string) {
			//Delete token
			c.DelSession("token")
			jsonerror := &JsonResult{
				Result: "success",
			}
			c.Data["json"] = jsonerror
		} else {
			jsonerror := &JsonResult{
				Result: "fail",
			}
			c.Data["json"] = jsonerror
		}
	}
	c.ServeJSON()
}

//user register and validate
func (c *LoginController) Register() {
	var v models.Users
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &v); err == nil {
		//check validity todo:valid form according to requirement
		valid := validation.Validation{}
		valid.Required(v.Firstname, "firstname")

		if _, err := models.AddUsers(&v); err == nil {
			c.Ctx.Output.SetStatus(201)
			jsonerror := &JsonResult{
				Result: "success",
			}
			c.Data["json"] = jsonerror
		} else {
			jsonerror := &JsonResult{
				Result: "fail",
			}
			c.Data["json"] = jsonerror
		}
	} else {
		jsonerror := &JsonResult{
			Result: "fail",
		}
		c.Data["json"] = jsonerror
	}
	c.ServeJSON()
}
