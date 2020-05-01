package structs

type Student struct {
	Id int `gorm:"primary_key" json:"id"`
	Name string `gorm:"type:varchar(20);column:name" json:"name"`
	Gender int `gorm:"type:int;column:gender"`
	Age int `gorm:"type:int;column:age"`
	City string `gorm:"varchar(20);column:city"`
	School string `gorm:"varchar(20);column:school"`
}
