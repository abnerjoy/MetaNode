package _sqlxTest

import (
	"fmt"

	"github.com/jmoiron/sqlx"
)

type employee struct {
	Id         int
	Name       string
	Department string
	Salary     float64
}

func Run(db *sqlx.DB) {
	//createTableSQL := ` //CREATE TABLE IF NOT EXISTS employee (
	//    id INT AUTO_INCREMENT PRIMARY KEY,
	//    name VARCHAR(100),
	//    department VARCHAR(100),
	//    salary FLOAT
	//)`
	//_, err := db.Exec(createTableSQL)
	//if err != nil {
	//	return
	//}

	// 插入单条数据
	//insertSQL := "INSERT INTO employee (name, department, salary) VALUES (?, ?, ?)"
	//_, err := db.Exec(insertSQL, "张三", "技术部", 8000.0)
	//if err != nil {
	//	panic(err)
	//}
	//
	//// 批量插入数据
	//employees := []employee{
	//	{Name: "李四", Department: "技术部", Salary: 8000.0},
	//	{Name: "王五", Department: "销售部", Salary: 7000.0},
	//	{Name: "赵六", Department: "人事部", Salary: 6500.0},
	//}
	//
	//for _, emp := range employees {
	//	_, err = db.Exec(insertSQL, emp.Name, emp.Department, emp.Salary)
	//	if err != nil {
	//		panic(err)
	//	}
	//}

	var employees []employee
	err := db.Select(&employees, "SELECT * FROM employee where department=? and name=?", "技术部", "张三")
	if err != nil {
		panic(err)
	}
	for _, emp := range employees {
		fmt.Printf("ID: %d, Name: %s, Department: %s, Salary: %.2f\n", emp.Id, emp.Name, emp.Department, emp.Salary)
	}

	var employee employee
	err = db.Get(&employee, "select * from employee order by salary desc limit 1")
	if err != nil {
		panic(err)
	}
	fmt.Printf("ID: %d, Name: %s, Department: %s, Salary: %.2f\n", employee.Id, employee.Name, employee.Department, employee.Salary)
}
