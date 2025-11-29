package _hook

import (
	"gorm.io/gorm"
)

type User struct {
	Id      int
	Name    string
	PostNum int
	Posts   []Post `gorm:"foreignKey:UserId;references:Id"`
}

type Post struct {
	Id           int
	Title        string
	UserId       int
	CommentNum   int
	CommentState string    //无评论，有评论
	Comments     []Comment `gorm:"foreignKey:PostId;references:Id"`
}

type Comment struct {
	Id      int
	Content string
	PostId  int
}

// 为 Post 模型添加一个钩子函数，在文章创建时自动更新用户的文章数量统计字段。
func (p *Post) AfterCreate(tx *gorm.DB) (err error) {
	var user User
	tx.Debug().Model(&user).
		Where("id = ?", p.UserId).Update("post_num", gorm.Expr("post_num + ?", 1))
	return

}

// 为 Comment 模型添加一个钩子函数，在评论删除时检查文章的评论数量，
// 如果评论数量为 0，则更新文章的评论状态为 "无评论"。
func (c *Comment) AfterDelete(tx *gorm.DB) (err error) {
	var post Post
	result := tx.Debug().Model(&post).
		Where("id = ?", c.PostId).
		Update("comment_num", gorm.Expr("GREATEST(comment_num - ?, 0)", 1)) // 防止减成负数
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return nil
	}
	return tx.Debug().Model(&post).
		Where("id = ?", c.PostId).
		Update("comment_state",
			gorm.Expr("CASE WHEN comment_num = 0 THEN '无评论' ELSE '有评论' END")).Error
}

func Run(db *gorm.DB) {
	//db.AutoMigrate(&User{}, &Post{}, &Comment{})
	//db.Create(&User{
	//	Name: "Alice",
	//	Posts: []Post{
	//		{
	//			Title: "Post 1",
	//			Comments: []Comment{
	//				{Content: "Comment 1"},
	//				{Content: "Comment 2"},
	//			},
	//			CommentNum: 2,
	//		},
	//		{
	//			Title: "Post 2",
	//			Comments: []Comment{
	//				{Content: "Comment 3"},
	//				{Content: "Comment 4"},
	//				{Content: "Comment 5"},
	//			},
	//			CommentNum: 3,
	//		},
	//	},
	//})

	var comment Comment
	db.First(&comment, 10)
	db.Delete(&comment)

}
