package web

import (
	regexp "github.com/dlclark/regexp2"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"go_learn/webook/internal/domain"
	"go_learn/webook/internal/service"
	"net/http"
)

// UserHandler 定义所有与User有关的路由
type UserHandler struct {
	svc         *service.UserService
	emailExp    *regexp.Regexp
	passwordExp *regexp.Regexp
	birthdayExp *regexp.Regexp
}

func NewUserHandler(svc *service.UserService) *UserHandler {
	const (
		emailRegexPattern    = "^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\\.[a-zA-Z]{2,}$"
		passwordRegexPattern = "^(?=.*[a-z])(?=.*[A-Z])(?=.*\\d)(?=.*[@$!%*?&])[A-Za-z\\d@$!%*?&]{8,16}$"
		birthdayPattern      = "^(19|20)\\d\\d-(0[1-9]|1[0-2])-(0[1-9]|[12][0-9]|3[01])$"
	)
	emailExp := regexp.MustCompile(emailRegexPattern, regexp.None)
	passwordExp := regexp.MustCompile(passwordRegexPattern, regexp.None)
	birthdayExp := regexp.MustCompile(birthdayPattern, regexp.None)
	return &UserHandler{
		svc:         svc,
		emailExp:    emailExp,
		passwordExp: passwordExp,
		birthdayExp: birthdayExp,
	}

}

func (u *UserHandler) RegisterRoutes(server *gin.Engine) {
	ug := server.Group("/users")
	// /users/signup   /users/login   /users/edit  /users/profile
	ug.POST("/signup", u.Signup)
	ug.POST("/login", u.Login)
	ug.POST("/edit", u.Edit)
	ug.GET("/profile", u.Profile)
}

func (u *UserHandler) Signup(c *gin.Context) {
	type SignupReq struct {
		Email           string `json:"email"`
		Password        string `json:"password"`
		ConfirmPassword string `json:"confirmPassword"`
	}

	var req SignupReq
	// Bind方法会根据Content-type来解析数据到req里面
	// 解析错误，则返回一个4xx错误

	//接收数据
	if err := c.Bind(&req); err != nil {
		return
	}

	// 校验
	ok, err := u.emailExp.MatchString(req.Email)
	if err != nil {
		c.String(http.StatusOK, "系统错误1")
		return
	}
	if !ok {
		c.String(http.StatusOK, "邮箱格式错误")
		return
	}

	if req.ConfirmPassword != req.Password {
		c.String(http.StatusOK, "两次输入的密码不一致")
		return
	}

	ok, err = u.passwordExp.MatchString(req.Password)
	if err != nil {
		c.String(http.StatusOK, "系统错误2")
		return
	}
	if !ok {
		c.String(http.StatusOK, "密码格式错误")
		return
	}

	err = u.svc.Signup(c, domain.User{
		Email:    req.Email,
		Password: req.Password,
	})

	if err == service.ErrUserDuplicateEmail {
		c.String(http.StatusOK, "重复邮箱，请换一个邮箱")
		return
	}
	if err != nil {
		c.String(http.StatusOK, "服务器异常，注册失败")
		return
	}

	c.String(http.StatusOK, "注册成功")

}

func (u *UserHandler) Login(c *gin.Context) {
	type LoginReq struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	var req LoginReq
	if err := c.Bind(&req); err != nil {
		return
	}

	user, err := u.svc.Login(c, req.Email, req.Password)
	if err == service.ErrInvalidUserOrPassword {
		c.String(http.StatusOK, "邮箱或密码不对")
		return
	}
	if err != nil {
		c.String(http.StatusOK, "系统错误")
		return
	}

	//设置session
	sess := sessions.Default(c)
	// 可以随便设置要放在session中的值
	sess.Set("userId", user.Id)
	sess.Save()

	c.String(http.StatusOK, "登录成功")

	return
}

func (u *UserHandler) Edit(c *gin.Context) {
	type EditReq struct {
		Nickname        string `json:"nickname"`
		Birthday        string `json:"birthday"`
		PersonalProfile string `json:"personalProfile"`
	}

	var req EditReq
	if err := c.Bind(&req); err != nil {
		return
	}

	// 校验
	// 校验生日格式
	ok, err := u.birthdayExp.MatchString(req.Birthday)
	if err != nil {
		c.String(http.StatusOK, "系统错误")
		return
	}
	if !ok {
		c.String(http.StatusOK, "生日格式错误")
		return
	}

	if len(req.Nickname) == 0 {
		c.String(http.StatusOK, "昵称不能为空")
		return
	}
	
	if len(req.Nickname) > 60 {
		c.String(http.StatusOK, "昵称过长，不能超过60个英文字符")
		return
	}

	if len(req.PersonalProfile) > 600 {
		c.String(http.StatusOK, "个人简介过长，不能超过600个英文字符")
		return
	}

	sess := sessions.Default(c)
	id := sess.Get("userId")

	err = u.svc.Edit(c, domain.User{
		Id:              id.(int64),
		Nickname:        req.Nickname,
		Birthday:        req.Birthday,
		PersonalProfile: req.PersonalProfile,
	})

	if err != nil {
		return
	}

	c.String(http.StatusOK, "编辑成功")
}

func (u *UserHandler) Profile(c *gin.Context) {
	sess := sessions.Default(c)
	id := sess.Get("userId")
	user, err := u.svc.Profile(c, id.(int64))
	if err != nil {
		c.String(http.StatusOK, "系统错误")
		return
	}
	c.String(http.StatusOK, "profile信息为：{Nickname:%s,Birthday:%s,PersonalProfile:%s", user.Nickname, user.Birthday, user.PersonalProfile)
}
