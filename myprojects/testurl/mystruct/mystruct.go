package mystruct

type Tip struct {
	Code int `json:"code"`
	Msg string `json:"msg"`
}

type Result struct {
	Code int `json:"code"`
	Msg string `json:"msg"`
	Data map[string]interface{} `json:"data"`
}

type MyContent struct {
	Id int `gorm:"primary_key;column:id"`
	Username string `gorm:"type:varchar(32);column:username" json:"username"`
	Area string `gorm:"varchar(32);column:passwords" json:"passwords"`
	Createtime int64 `gorm:"type:bigint;column:createtime" json:"createtime"`
	MyWords string
}

type UserInfo struct {
	Username string `bson:"username" json:"username"`
	Password string `bason:"password" json:"password"`
	Argot string 	`bason:"argot" json:"argot"`
	Num int			`bason"num" json:"num"`
}

type User struct {
	Name string `bson:"name" json:"name"`
	Password string `bason:"password" json:"password"`
	Argot string 	`bason:"argot" json:"argot"`
	Num int			`bason"num" json:"num"`
}