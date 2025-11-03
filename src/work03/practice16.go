package main

/*## 进阶gorm
### 题目1：模型定义
- 假设你要开发一个博客系统，有以下几个实体： User （用户）、 Post （文章）、 Comment （评论）。
  - 要求 ：
    - 使用Gorm定义 User 、 Post 和 Comment 模型，其中 User 与 Post 是一对多关系（一个用户可以发布多篇文章）， Post 与 Comment 也是一对多关系（一篇文章可以有多个评论）。
    - 编写Go代码，使用Gorm创建这些模型对应的数据库表。
### 题目2：关联查询
- 基于上述博客系统的模型定义。
  - 要求 ：
    - 编写Go代码，使用Gorm查询某个用户发布的所有文章及其对应的评论信息。
    - 编写Go代码，使用Gorm查询评论数量最多的文章信息。
### 题目3：钩子函数
- 继续使用博客系统的模型。
  - 要求 ：
    - 为 Post 模型添加一个钩子函数，在文章创建时自动更新用户的文章数量统计字段。
    - 为 Comment 模型添加一个钩子函数，在评论删除时检查文章的评论数量，如果评论数量为 0，则更新文章的评论状态为 "无评论"。*/
import (
	"fmt"
	"log"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// User 用户模型
type User struct {
	ID         uint   `gorm:"primaryKey"`
	Name       string `gorm:"size:100;not null"`
	Email      string `gorm:"size:100;uniqueIndex;not null"`
	Posts      []Post `gorm:"foreignKey:UserID"`
	PostsCount int    `gorm:"default:0"` // 用于统计用户文章数量
	CreatedAt  time.Time
	UpdatedAt  time.Time
}

// Post 文章模型
type Post struct {
	ID            uint      `gorm:"primaryKey"`
	Title         string    `gorm:"size:200;not null"`
	Content       string    `gorm:"type:text;not null"`
	UserID        uint      `gorm:"not null;index"`
	User          User      `gorm:"foreignKey:UserID"`
	Comments      []Comment `gorm:"foreignKey:PostID"`
	CommentCount  int       `gorm:"default:0"`             // 评论数量统计
	CommentStatus string    `gorm:"size:20;default:'无评论'"` // 评论状态
	CreatedAt     time.Time
	UpdatedAt     time.Time
}

// Comment 评论模型
type Comment struct {
	ID        uint   `gorm:"primaryKey"`
	Content   string `gorm:"type:text;not null"`
	PostID    uint   `gorm:"not null;index"`
	Post      Post   `gorm:"foreignKey:PostID"`
	UserID    uint   `gorm:"not null;index"`
	User      User   `gorm:"foreignKey:UserID"`
	CreatedAt time.Time
}

var db *gorm.DB

func initDB() {
	dsn := "username:password@tcp(localhost:3306)/blog_system?charset=utf8mb4&parseTime=True&loc=Local"

	var err error
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		log.Fatal("连接数据库失败:", err)
	}

	fmt.Println("数据库连接成功")
}

// 创建数据库表
func createTables() {
	// 自动迁移模式，创建表并添加外键约束
	err := db.AutoMigrate(&User{}, &Post{}, &Comment{})
	if err != nil {
		log.Fatal("创建表失败:", err)
	}

	fmt.Println("数据库表创建成功")
}

func main() {
	initDB()
	createTables()

	// 创建测试数据
	createTestData()
}

// 创建测试数据
func createTestData() {
	// 创建用户
	users := []User{
		{Name: "张三", Email: "zhangsan@example.com"},
		{Name: "李四", Email: "lisi@example.com"},
		{Name: "王五", Email: "wangwu@example.com"},
	}

	for i := range users {
		result := db.Create(&users[i])
		if result.Error != nil {
			log.Printf("创建用户失败: %v", result.Error)
		}
	}

	fmt.Println("测试数据创建完成")
}
