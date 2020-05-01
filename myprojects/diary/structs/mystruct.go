package structs

//部分操作完成后的返回结构
type Tip struct {
	Code int `json:"code"`
	Msg string `json:"msg"`
}

//账户所对应的类
type Account struct {
	Id int `gorm:"primary_key;column:id"`
	Username string `gorm:"type:varchar(32);column:username" json:"username"`
	Passwords string `gorm:"varchar(32);column:passwords" json:"passwords"`
	Createtime int64 `gorm:"type:bigint;column:createtime" json:"createtime"`
}

//记事内容上传
type Diary struct {
	Id int `gorm:"primary_key;column:id"`
	Username string `gorm:"type:varchar(32);column:username" json:"username"`
	Diaryid int `gorm:"type:integer;column:diaryid" json:"diaryid"`
	Title string `gorm:"type:varchar(64);column:title" json:"title"`
	Content string `gorm:"type:varchar(1024);column:content" json:"content"`
	Createtime int64 `gorm:"type:bigint;column:createtime" json:"createtime"`
	Lastupdatetime int64 `gorm:"type:bigint;column:lastupdatetime" json:"lastupdatetime"`
}
