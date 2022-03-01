package ApiSkeleton

type User struct {

	//用户编号
	Id int `gorm:"column:id" json:"id"`

	//用户名
	Username string `gorm:"column:username" json:"username"`

	//密码
	Password string `gorm:"column:password" json:"password"`

	//邮箱
	Email string `gorm:"column:email" json:"email"`

	//年龄
	Age int8 `gorm:"column:age" json:"age"`

	//电话
	Tel string `gorm:"column:tel" json:"tel"`

	//地址
	Addr string `gorm:"column:addr" json:"addr"`

	//身份证号
	Card string `gorm:"column:card" json:"card"`

	//创建时间
	CreateTime int64 `gorm:"column:createTime" json:"createTime"`

	//更新时间
	UpdateTime int64 `gorm:"column:updateTime" json:"updateTime"`
}

func (model *User) TableName() string {
	return "user"
}
