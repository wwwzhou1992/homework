package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

// 数据库初始化
func initDB() (*sqlx.DB, error) {
	db, err := sqlx.Connect("mysql", "root:password@tcp(localhost:3306)/testdb?parseTime=true")
	if err != nil {
		return nil, err
	}

	// 创建测试表
	schema := `
    CREATE TABLE IF NOT EXISTS employees (
        id INT AUTO_INCREMENT PRIMARY KEY,
        name VARCHAR(100) NOT NULL,
        department VARCHAR(100) NOT NULL,
        salary INT NOT NULL
    );
    
    CREATE TABLE IF NOT EXISTS books (
        id INT AUTO_INCREMENT PRIMARY KEY,
        title VARCHAR(200) NOT NULL,
        author VARCHAR(100) NOT NULL,
        price DECIMAL(10,2) NOT NULL
    );`

	_, err = db.Exec(schema)
	if err != nil {
		return nil, err
	}

	return db, nil
}

// 插入测试数据
func insertTestData(db *sqlx.DB) error {
	// 插入员工数据
	employees := []Employee{
		{Name: "张三", Department: "技术部", Salary: 8000},
		{Name: "李四", Department: "技术部", Salary: 9500},
		{Name: "王五", Department: "市场部", Salary: 7000},
		{Name: "赵六", Department: "技术部", Salary: 12000},
	}

	for _, emp := range employees {
		_, err := db.NamedExec(
			"INSERT INTO employees (name, department, salary) VALUES (:name, :department, :salary)",
			emp)
		if err != nil {
			return err
		}
	}

	// 插入书籍数据
	books := []Book{
		{Title: "Go语言编程", Author: "张三", Price: 65.50},
		{Title: "数据库设计", Author: "李四", Price: 45.00},
		{Title: "Web开发实战", Author: "王五", Price: 78.90},
		{Title: "算法导论", Author: "赵六", Price: 99.99},
	}

	for _, book := range books {
		_, err := db.NamedExec(
			"INSERT INTO books (title, author, price) VALUES (:title, :author, :price)",
			book)
		if err != nil {
			return err
		}
	}

	return nil
}

func main() {
	db, err := initDB()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// 插入测试数据
	err = insertTestData(db)
	if err != nil {
		log.Fatal(err)
	}

	// 执行题目1的查询
	fmt.Println("=== 题目1：员工查询 ===")

	// 查询技术部员工
	var techEmployees []Employee
	err = db.Select(&techEmployees,
		"SELECT id, name, department, salary FROM employees WHERE department = ?", "技术部")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("技术部员工:")
	for _, emp := range techEmployees {
		fmt.Printf("  %s - 薪资: %d\n", emp.Name, emp.Salary)
	}

	// 查询最高薪资员工
	var topEmployee Employee
	err = db.Get(&topEmployee,
		"SELECT id, name, department, salary FROM employees ORDER BY salary DESC LIMIT 1")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("最高薪资员工: %s, 薪资: %d\n", topEmployee.Name, topEmployee.Salary)

	// 执行题目2的查询
	fmt.Println("\n=== 题目2：书籍查询 ===")

	var expensiveBooks []Book
	err = db.Select(&expensiveBooks,
		"SELECT id, title, author, price FROM books WHERE price > ?", 50.0)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("价格大于50元的书籍:")
	for _, book := range expensiveBooks {
		fmt.Printf("  《%s》 - %s - ￥%.2f\n", book.Title, book.Author, book.Price)
	}
}
