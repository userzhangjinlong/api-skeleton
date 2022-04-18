package ApiRequest

type LoginForm struct {
	Username string `form:"username" binding:"required,regPhone"`
	Password string `form:"password" binding:"required,min=6"`
}
