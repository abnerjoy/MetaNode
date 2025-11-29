package _And6reference

import (
	"fmt"

	"gorm.io/gorm"
)

type User struct {
	Id    int
	Name  string
	Posts []Post `gorm:"foreignKey:UserId;references:Id"`
}

type Post struct {
	Id       int
	Title    string
	UserId   int
	Comments []Comment `gorm:"foreignKey:PostId;references:Id"`
}

type Comment struct {
	Id      int
	Content string
	PostId  int
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
	//		},
	//		{
	//			Title: "Post 2",
	//			Comments: []Comment{
	//				{Content: "Comment 3"},
	//				{Content: "Comment 4"},
	//			},
	//		},
	//	},
	//})

	//db.Debug().Create(&User{
	//	Name: "Bob",
	//	Posts: []Post{
	//		{
	//			Title: "Post 3",
	//			Comments: []Comment{
	//				{Content: "Comment 5"},
	//				{Content: "Comment 6"},
	//				{Content: "Comment 7"},
	//			},
	//		},
	//		{
	//			Title: "Post 4",
	//			Comments: []Comment{
	//				{Content: "Comment 8"},
	//				{Content: "Comment 9"},
	//			},
	//		},
	//	},
	//})

	//使用Gorm查询某个用户发布的所有文章及其对应的评论信息。
	//var user User
	//db.Debug().Preload("Posts.Comments").Preload("Posts").Where("name = ?", "Alice").Find(&user)
	//fmt.Println(user)

	//var user User
	//db.Debug(). // 1. 只取 user.name
	//		Preload("Posts", func(db *gorm.DB) *gorm.DB {
	//		return db.Select("title", "user_id", "id") // 2. 只取 post.title + 外键
	//	}).
	//	Preload("Posts.Comments", func(db *gorm.DB) *gorm.DB {
	//		return db.Select("content", "post_id", "id") // 3. 只取 comment.content + 外键
	//	}).Where("name = ?", "Alice").
	//	Find(&user)
	//
	//fmt.Println(user)

	//使用Gorm查询评论数量最多的文章信息。
	var post Post
	db.Debug().Model(&post).
		Joins("JOIN comments ON comments.post_id = posts.id").Group("posts.id").
		Order("count(comments.id) DESC").First(&post)
	fmt.Println(post)

}
