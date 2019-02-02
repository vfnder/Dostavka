package dostavkarepo

import (
	"github.com/jinzhu/gorm"
)

//Dish
/*type Dish struct {
	gorm.Model
	ID   int
	Name string
	Cost int
	IMG  string
}*/

func getDishes() {
	db, err := gorm.Open("mssql", "sqlserver=WIN-RO24Q0C93S6\\SQLEXPRESS;username=WIN-RO24Q0C93S6\\User;password='';database=FoodlandDelivery")
	if err != nil {
		panic("failed to connect database")
	}
	defer db.Close()
}
