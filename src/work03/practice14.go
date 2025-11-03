package practice14

import (
	"fmt"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"gorm.io/gorm/logger"
)

/**
题目1：基本CRUD操作
假设有一个名为 students 的表，包含字段 id （主键，自增）、 name （学生姓名，字符串类型）、 age （学生年龄，整数类型）、 grade （学生年级，字符串类型）。
要求 ：
编写SQL语句向 students 表中插入一条新记录，学生姓名为 "张三"，年龄为 20，年级为 "三年级"。
编写SQL语句查询 students 表中所有年龄大于 18 岁的学生信息。
编写SQL语句将 students 表中姓名为 "张三" 的学生年级更新为 "四年级"。
编写SQL语句删除 students 表中年龄小于 15 岁的学生记录。
 * @return:
*/
// Student 结构体映射students表
type Student struct {
	ID    uint   `gorm:"primaryKey;autoIncrement" json:"id"`
	Name  string `gorm:"size:100;not null" json:"name"`
	Age   int    `gorm:"not null" json:"age"`
	Grade string `gorm:"size:100;not null" json:"grade"`
}

// ===== 题目2：事务语句 =====
/**
题目2：事务语句
假设有两个表： accounts 表（包含字段 id 主键， balance 账户余额）和 transactions 表（包含字段 id 主键， from_account_id 转出账户ID， to_account_id 转入账户ID， amount 转账金额）。
要求 ：
编写一个事务，实现从账户 A 向账户 B 转账 100 元的操作。在事务中，需要先检查账户 A 的余额是否足够，如果足够则从账户 A 扣除 100 元，向账户 B 增加 100 元，并在 transactions 表中记录该笔转账信息。如果余额不足，则回滚事务。
 * @return:
*/

// Account 结构体映射accounts表
type Account struct {
	ID      uint    `gorm:"primaryKey" json:"id"`
	Balance float64 `gorm:"not null;default:0" json:"balance"`
}

// Transaction 结构体映射transactions表
type Transaction struct {
	ID            uint    `gorm:"primaryKey;autoIncrement" json:"id"`
	FromAccountID uint    `gorm:"not null" json:"from_account_id"`
	ToAccountID   uint    `gorm:"not null" json:"to_account_id"`
	Amount        float64 `gorm:"not null" json:"amount"`
}

// InitDB 初始化数据库连接
func InitDB(dsn string) (*gorm.DB, error) {
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info), // 启用SQL日志
	})
	if err != nil {
		return nil, err
	}

	// 自动迁移表结构
	err = db.AutoMigrate(&Student{}, &Account{}, &Transaction{})
	if err != nil {
		return nil, err
	}

	return db, nil
}

// DemoCRUDOperations 演示基本的CRUD操作（GORM方式）
func DemoCRUDOperations(db *gorm.DB) {
	fmt.Println("===== 题目1：基本CRUD操作（GORM实现） =====")

	// 1. 插入操作：向students表中插入一条新记录
	fmt.Println("1. 插入操作:")
	student := Student{Name: "张三", Age: 20, Grade: "三年级"}
	result := db.Create(&student)
	if result.Error != nil {
		log.Printf("插入失败: %v\n", result.Error)
	} else {
		fmt.Printf("插入成功！ID: %d, 姓名: %s, 年龄: %d, 年级: %s\n\n",
			student.ID, student.Name, student.Age, student.Grade)
	}

	// 插入更多测试数据
	testStudents := []Student{
		{Name: "李四", Age: 19, Grade: "二年级"},
		{Name: "王五", Age: 21, Grade: "四年级"},
		{Name: "赵六", Age: 17, Grade: "一年级"},
		{Name: "钱七", Age: 22, Grade: "四年级"},
	}
	db.Create(&testStudents)

	// 2. 查询操作：查询年龄大于18岁的学生信息
	fmt.Println("2. 查询操作 - 年龄大于18岁的学生:")
	var students []Student
	result = db.Where("age > ?", 18).Find(&students)
	if result.Error != nil {
		log.Printf("查询失败: %v\n", result.Error)
	} else {
		for _, s := range students {
			fmt.Printf("  ID: %d, 姓名: %s, 年龄: %d, 年级: %s\n",
				s.ID, s.Name, s.Age, s.Grade)
		}
		fmt.Println()
	}

	// 3. 更新操作：将姓名为"张三"的学生年级更新为"四年级"
	fmt.Println("3. 更新操作:")
	result = db.Model(&Student{}).Where("name = ?", "张三").Update("grade", "四年级")
	if result.Error != nil {
		log.Printf("更新失败: %v\n", result.Error)
	} else {
		fmt.Printf("更新成功！影响行数: %d\n\n", result.RowsAffected)
	}

	// 4. 删除操作：删除年龄小于15岁的学生记录
	fmt.Println("4. 删除操作:")
	result = db.Where("age < ?", 15).Delete(&Student{})
	if result.Error != nil {
		log.Printf("删除失败: %v\n", result.Error)
	} else {
		fmt.Printf("删除成功！影响行数: %d\n\n", result.RowsAffected)
	}

	// 显示最终所有学生
	fmt.Println("5. 最终所有学生信息:")
	var allStudents []Student
	db.Find(&allStudents)
	for _, s := range allStudents {
		fmt.Printf("  ID: %d, 姓名: %s, 年龄: %d, 年级: %s\n",
			s.ID, s.Name, s.Age, s.Grade)
	}
	fmt.Println()
}

// DemoTransaction 演示转账事务操作（GORM方式）
func DemoTransaction(db *gorm.DB) {
	fmt.Println("===== 题目2：事务语句（GORM实现） =====")

	// 准备测试数据
	accounts := []Account{
		{ID: 1, Balance: 500.0},
		{ID: 2, Balance: 300.0},
		{ID: 3, Balance: 200.0},
	}

	// 使用Clauses确保插入或更新
	for _, account := range accounts {
		db.Clauses(clause.OnConflict{
			UpdateAll: true,
		}).Create(&account)
	}

	// 显示转账前的账户余额
	fmt.Println("转账前账户余额:")
	var beforeAccounts []Account
	db.Find(&beforeAccounts)
	for _, acc := range beforeAccounts {
		fmt.Printf("  账户ID: %d, 余额: %.2f\n", acc.ID, acc.Balance)
	}
	fmt.Println()

	// 执行转账操作
	fromAccountID := uint(1)
	toAccountID := uint(2)
	amount := 100.0

	fmt.Printf("执行转账: 从账户 %d 向账户 %d 转账 %.2f 元\n\n", fromAccountID, toAccountID, amount)

	// 使用GORM事务
	err := db.Transaction(func(tx *gorm.DB) error {
		// 1. 查询转出账户并加锁
		var fromAccount Account
		if err := tx.Clauses(clause.Locking{Strength: "UPDATE"}).
			First(&fromAccount, fromAccountID).Error; err != nil {
			return fmt.Errorf("查询转出账户失败: %v", err)
		}

		// 2. 检查余额是否足够
		if fromAccount.Balance < amount {
			return fmt.Errorf("余额不足: 账户 %d 余额 %.2f < 转账金额 %.2f",
				fromAccount.ID, fromAccount.Balance, amount)
		}

		// 3. 从转出账户扣除金额
		if err := tx.Model(&Account{}).
			Where("id = ?", fromAccountID).
			Update("balance", gorm.Expr("balance - ?", amount)).Error; err != nil {
			return fmt.Errorf("扣款失败: %v", err)
		}

		// 4. 向转入账户增加金额
		if err := tx.Model(&Account{}).
			Where("id = ?", toAccountID).
			Update("balance", gorm.Expr("balance + ?", amount)).Error; err != nil {
			return fmt.Errorf("存款失败: %v", err)
		}

		// 5. 记录交易
		transaction := Transaction{
			FromAccountID: fromAccountID,
			ToAccountID:   toAccountID,
			Amount:        amount,
		}
		if err := tx.Create(&transaction).Error; err != nil {
			return fmt.Errorf("记录交易失败: %v", err)
		}

		fmt.Println("  转账操作步骤完成")
		return nil // 返回nil提交事务
	})

	if err != nil {
		fmt.Printf("❌ 转账失败: %v\n\n", err)
	} else {
		fmt.Println("✅ 转账成功！\n")

		// 显示转账后的账户余额
		fmt.Println("转账后账户余额:")
		var afterAccounts []Account
		db.Find(&afterAccounts)
		for _, acc := range afterAccounts {
			fmt.Printf("  账户ID: %d, 余额: %.2f\n", acc.ID, acc.Balance)
		}
		fmt.Println()

		// 显示交易记录
		fmt.Println("交易记录:")
		var transactions []Transaction
		db.Order("id DESC").Limit(5).Find(&transactions)
		for _, tx := range transactions {
			fmt.Printf("  交易ID: %d, 从账户: %d, 到账户: %d, 金额: %.2f\n",
				tx.ID, tx.FromAccountID, tx.ToAccountID, tx.Amount)
		}
	}
}

// DemoAdvancedCRUD 演示更多GORM高级CRUD操作
func DemoAdvancedCRUD(db *gorm.DB) {
	fmt.Println("===== 高级CRUD操作演示 =====")

	// 批量插入
	students := []Student{
		{Name: "孙八", Age: 18, Grade: "二年级"},
		{Name: "周九", Age: 20, Grade: "三年级"},
		{Name: "吴十", Age: 19, Grade: "二年级"},
	}
	db.CreateInBatches(students, 2) // 每批插入2条
	fmt.Println("批量插入完成")

	// 复杂查询
	fmt.Println("\n复杂查询示例:")

	// 查询年龄在18-20之间，年级包含"二"的学生
	var filteredStudents []Student
	db.Where("age BETWEEN ? AND ?", 18, 20).
		Where("grade LIKE ?", "%二%").
		Order("age DESC").
		Find(&filteredStudents)

	fmt.Println("年龄18-20且年级包含'二'的学生:")
	for _, s := range filteredStudents {
		fmt.Printf("  %s, 年龄: %d, 年级: %s\n", s.Name, s.Age, s.Grade)
	}

	// 统计查询
	var count int64
	var averageAge float64

	db.Model(&Student{}).Count(&count)
	db.Model(&Student{}).Select("AVG(age)").Scan(&averageAge)

	fmt.Printf("\n统计信息: 总学生数: %d, 平均年龄: %.2f\n", count, averageAge)

	// 事务中的复杂操作
	fmt.Println("\n事务中的批量操作:")
	err := db.Transaction(func(tx *gorm.DB) error {
		// 批量更新：将所有二年级学生的年龄增加1
		if err := tx.Model(&Student{}).
			Where("grade = ?", "二年级").
			Update("age", gorm.Expr("age + ?", 1)).Error; err != nil {
			return err
		}

		// 删除特定条件的学生
		if err := tx.Where("age > ?", 25).Delete(&Student{}).Error; err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		fmt.Printf("批量操作失败: %v\n", err)
	} else {
		fmt.Println("批量操作成功")
	}
}

// Example 主执行函数
func Example() {
	// 数据库连接DSN
	dsn := "username:password@tcp(localhost:3306)/school?charset=utf8mb4&parseTime=True&loc=Local"

	// 初始化数据库连接
	db, err := InitDB(dsn)
	if err != nil {
		log.Fatal("数据库连接失败:", err)
	}

	// 执行基本CRUD操作演示
	DemoCRUDOperations(db)

	// 执行事务操作演示
	DemoTransaction(db)

	// 执行高级CRUD操作演示
	DemoAdvancedCRUD(db)
}
