package main

/*## 进阶gorm

### 题目3：钩子函数
- 继续使用博客系统的模型。
  - 要求 ：
    - 为 Post 模型添加一个钩子函数，在文章创建时自动更新用户的文章数量统计字段。
    - 为 Comment 模型添加一个钩子函数，在评论删除时检查文章的评论数量，如果评论数量为 0，则更新文章的评论状态为 "无评论"。*/
import (
	"fmt"
	"log"

	"gorm.io/gorm"
)

// Post 模型的钩子函数

// BeforeCreate 在创建文章前更新用户文章计数
func (p *Post) BeforeCreate(tx *gorm.DB) (err error) {
	// 更新用户的文章数量
	result := tx.Model(&User{}).Where("id = ?", p.UserID).
		Update("posts_count", gorm.Expr("posts_count + ?", 1))

	if result.Error != nil {
		return fmt.Errorf("更新用户文章计数失败: %v", result.Error)
	}

	if result.RowsAffected == 0 {
		return fmt.Errorf("用户不存在")
	}

	return nil
}

// AfterCreate 在创建文章后初始化评论状态
func (p *Post) AfterCreate(tx *gorm.DB) (err error) {
	// 确保新文章的评论状态正确
	if p.CommentCount == 0 {
		tx.Model(p).Update("comment_status", "无评论")
	} else {
		tx.Model(p).Update("comment_status", "有评论")
	}
	return nil
}

// BeforeDelete 在删除文章前更新用户文章计数
func (p *Post) BeforeDelete(tx *gorm.DB) (err error) {
	// 减少用户的文章数量
	result := tx.Model(&User{}).Where("id = ?", p.UserID).
		Update("posts_count", gorm.Expr("posts_count - ?", 1))

	if result.Error != nil {
		return fmt.Errorf("更新用户文章计数失败: %v", result.Error)
	}

	return nil
}

// Comment 模型的钩子函数

// AfterCreate 在创建评论后更新文章评论计数和状态
func (c *Comment) AfterCreate(tx *gorm.DB) (err error) {
	// 更新文章的评论数量
	result := tx.Model(&Post{}).Where("id = ?", c.PostID).
		Update("comment_count", gorm.Expr("comment_count + ?", 1))

	if result.Error != nil {
		return fmt.Errorf("更新文章评论计数失败: %v", result.Error)
	}

	// 更新文章的评论状态
	result = tx.Model(&Post{}).Where("id = ?", c.PostID).
		Update("comment_status", "有评论")

	if result.Error != nil {
		return fmt.Errorf("更新文章评论状态失败: %v", result.Error)
	}

	return nil
}

// AfterDelete 在删除评论后检查文章评论数量
func (c *Comment) AfterDelete(tx *gorm.DB) (err error) {
	// 查询当前文章的评论数量
	var commentCount int64
	err = tx.Model(&Comment{}).Where("post_id = ?", c.PostID).Count(&commentCount).Error
	if err != nil {
		return fmt.Errorf("查询评论数量失败: %v", err)
	}

	// 更新文章的评论数量
	result := tx.Model(&Post{}).Where("id = ?", c.PostID).
		Update("comment_count", commentCount)

	if result.Error != nil {
		return fmt.Errorf("更新文章评论计数失败: %v", result.Error)
	}

	// 如果评论数量为0，更新评论状态
	if commentCount == 0 {
		result = tx.Model(&Post{}).Where("id = ?", c.PostID).
			Update("comment_status", "无评论")

		if result.Error != nil {
			return fmt.Errorf("更新文章评论状态失败: %v", result.Error)
		}
	}

	return nil
}

// 测试钩子函数
func testHooks() {
	fmt.Println("\n=== 测试钩子函数 ===")

	// 获取一个用户来创建新文章
	var user User
	db.First(&user)

	// 创建新文章测试 BeforeCreate 钩子
	newPost := Post{
		Title:   "测试钩子函数的文章",
		Content: "这篇文章用于测试GORM的钩子函数...",
		UserID:  user.ID,
	}

	result := db.Create(&newPost)
	if result.Error != nil {
		log.Printf("创建文章失败: %v", result.Error)
	} else {
		fmt.Println("文章创建成功，钩子函数已执行")

		// 验证用户文章计数是否更新
		var updatedUser User
		db.First(&updatedUser, user.ID)
		fmt.Printf("用户文章数量: %d\n", updatedUser.PostsCount)
	}

	// 为文章创建评论测试 Comment 的 AfterCreate 钩子
	newComment := Comment{
		Content: "测试评论钩子函数",
		PostID:  newPost.ID,
		UserID:  user.ID,
	}

	result = db.Create(&newComment)
	if result.Error != nil {
		log.Printf("创建评论失败: %v", result.Error)
	} else {
		fmt.Println("评论创建成功，钩子函数已执行")

		// 验证文章评论计数和状态
		var updatedPost Post
		db.First(&updatedPost, newPost.ID)
		fmt.Printf("文章评论数量: %d, 评论状态: %s\n",
			updatedPost.CommentCount, updatedPost.CommentStatus)
	}

	// 删除评论测试 AfterDelete 钩子
	result = db.Delete(&newComment)
	if result.Error != nil {
		log.Printf("删除评论失败: %v", result.Error)
	} else {
		fmt.Println("评论删除成功，钩子函数已执行")

		// 验证文章评论计数和状态
		var finalPost Post
		db.First(&finalPost, newPost.ID)
		fmt.Printf("删除后文章评论数量: %d, 评论状态: %s\n",
			finalPost.CommentCount, finalPost.CommentStatus)
	}

	// 删除文章测试 BeforeDelete 钩子
	result = db.Delete(&newPost)
	if result.Error != nil {
		log.Printf("删除文章失败: %v", result.Error)
	} else {
		fmt.Println("文章删除成功，钩子函数已执行")

		// 验证用户文章计数是否更新
		var finalUser User
		db.First(&finalUser, user.ID)
		fmt.Printf("删除后用户文章数量: %d\n", finalUser.PostsCount)
	}
}

// 完整的main函数
func main() {
	initDB()
	createTables()

	// 创建基础测试数据
	createTestData()
	createPostsAndComments()

	// 运行关联查询测试
	fmt.Println("=== 关联查询测试 ===")
	user, _ := getUserPostsWithComments(1)
	fmt.Printf("用户 %s 有 %d 篇文章\n", user.Name, len(user.Posts))

	popularPost, _ := getMostCommentedPost()
	fmt.Printf("最热门文章: %s (评论数: %d)\n", popularPost.Title, popularPost.CommentCount)

	// 运行钩子函数测试
	testHooks()
}

// 额外的工具函数：获取用户统计信息
func getUserStats(userID uint) (map[string]interface{}, error) {
	var user User
	err := db.Preload("Posts").Preload("Posts.Comments").First(&user, userID).Error
	if err != nil {
		return nil, err
	}

	totalComments := 0
	for _, post := range user.Posts {
		totalComments += len(post.Comments)
	}

	stats := map[string]interface{}{
		"user_name":         user.Name,
		"posts_count":       len(user.Posts),
		"total_comments":    totalComments,
		"posts_count_field": user.PostsCount, // 从字段读取
	}

	return stats, nil
}
