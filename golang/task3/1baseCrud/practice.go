package _baseCrud

import (
	"gorm.io/gorm"
)

type Student struct {
	Id    int
	Name  string `gorm:"type:varchar(100)"`
	Age   int
	Grade string `gorm:"type:varchar(100)"`
}

func Run(db *gorm.DB) {
	//db.AutoMigrate(&Student{})

	//var student Student = Student{
	//	Name:  "张三",
	//	Age:   20,
	//	Grade: "三年级",
	//}
	//db.Create(&student)

	//var students []Student
	//db.Debug().Find(&students, "age > ?", 18)
	//for _, row := range students {
	//	fmt.Println(row)
	//}

	//db.Debug().Model(&Student{}).Where("name=?", "张三").Update("Grade", "四年级")
	db.Debug().Where("age<?", 15).Delete(&Student{})
}
