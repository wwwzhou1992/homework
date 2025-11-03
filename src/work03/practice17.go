package main

/*## 进阶gorm
### 题目1：模型定义

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
)

// 查询某个用户发布的所有文章及其对应的评论信息
func getUserPostsWithComments(userID uint) (User, error) {
	var user User

	// 预加载用户的所有文章及文章的评论
	err := db.Preload("Posts.Comments").Preload("Posts.User").
		Where("id = ?", userID).First(&user).Error
	if err != nil {
		return User{}, fmt.Errorf("查询用户文章失败: %v", err)
	}

	return user, nil
}

// 查询评论数量最多的文章信息
func getMostCommentedPost() (Post, error) {
	var post Post

	// 方法1: 使用排序
	err := db.Preload("User").Preload("Comments").
		Order("comment_count DESC").First(&post).Error
	if err != nil {
		return Post{}, fmt.Errorf("查询最热门文章失败: %v", err)
	}

	return post, nil
}

// 查询评论数量最多的文章信息（使用原生SQL）
func getMostCommentedPostRaw() (Post, error) {
	var post Post

	err := db.Raw(`
			SELECT p.*, COUNT(c.id) as comment_count
			FROM posts p
			LEFT JOIN comments c ON p.id = c.post_id
			GROUP BY p.id
			ORDER BY comment_count DESC
			LIMIT 1
		`).Scan(&post).Error
	if err != nil {
		return Post{}, fmt.Errorf("查询最热门文章失败: %v", err)
	}

	return post, nil
}

// 创建测试文章和评论数据
func createPostsAndComments() {
	// 获取用户
	var users []User
	db.Find(&users)

	if len(users) < 2 {
		log.Fatal("用户数量不足")
	}

	// 为第一个用户创建文章
	posts := []Post{
		{
			Title:   "Go语言入门指南",
			Content: "这是一篇关于Go语言入门的详细指南...",
			UserID:  users[0].ID,
		},
		{
			Title:   "GORM使用教程",
			Content: "本文将详细介绍GORM的使用方法...",
			UserID:  users[0].ID,
		},
		{
			Title:   "数据库设计原则",
			Content: "分享数据库设计的最佳实践...",
			UserID:  users[1].ID,
		},
	}

	for i := range posts {
		result := db.Create(&posts[i])
		if result.Error != nil {
			log.Printf("创建文章失败: %v", result.Error)
		}
	}

	// 为文章创建评论
	comments := []Comment{
		{Content: "很好的文章！", PostID: posts[0].ID, UserID: users[1].ID},
		{Content: "学到了很多，谢谢分享", PostID: posts[0].ID, UserID: users[2].ID},
		{Content: "期待更多关于GORM的内容", PostID: posts[1].ID, UserID: users[2].ID},
		{Content: "数据库设计很重要", PostID: posts[2].ID, UserID: users[0].ID},
	}

	for i := range comments {
		result := db.Create(&comments[i])
		if result.Error != nil {
			log.Printf("创建评论失败: %v", result.Error)
		}
	}

	fmt.Println("文章和评论数据创建完成")
}

func main() {
	initDB()

	// 创建测试数据
	createPostsAndComments()

	// 测试关联查询
	fmt.Println("\n=== 测试关联查询 ===")

	// 查询用户1的所有文章及评论
	user, err := getUserPostsWithComments(1)
	if err != nil {
		log.Printf("查询失败: %v", err)
	} else {
		fmt.Printf("用户 %s 的文章:\n", user.Name)
		for _, post := range user.Posts {
			fmt.Printf("  - 文章: %s (评论数: %d)\n", post.Title, len(post.Comments))
			for _, comment := range post.Comments {
				fmt.Printf("    * 评论: %s\n", comment.Content)
			}
		}
	}

	// 查询评论最多的文章
	popularPost, err := getMostCommentedPost()
	if err != nil {
		log.Printf("查询失败: %v", err)
	} else {
		fmt.Printf("\n最热门的文章: %s (评论数: %d)\n",
			popularPost.Title, popularPost.CommentCount)
	}
}
