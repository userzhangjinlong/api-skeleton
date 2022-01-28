package Model

type User struct {
   
	//用户编号  
	 Id int32 `gorm:"id" json:"id"` 
  
	//用户名  
	 Username string `gorm:"username" json:"username"` 
  
	//密码  
	 Password string `gorm:"password" json:"password"` 
  
	//邮箱  
	 Email string `gorm:"email" json:"email"` 
  
	//年龄  
	 Age int8 `gorm:"age" json:"age"` 
  
	//电话  
	 Tel string `gorm:"tel" json:"tel"` 
  
	//地址  
	 Addr string `gorm:"addr" json:"addr"` 
  
	//身份证号  
	 Card string `gorm:"card" json:"card"` 
  
	//创建时间  
	 CreateTime int32 `gorm:"createTime" json:"createTime"` 
  
	//更新时间  
	 UpdateTime int32 `gorm:"updateTime" json:"updateTime"` 

}

func (model *User) TableName() string {
	return "user"
}