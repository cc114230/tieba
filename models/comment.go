package models

import "time"

type Comment struct {
	ID          int64     `json:"id,string" db:"comment_id"`
	PostID      int64     `json:"post_id" db:"post_id" binding:"required"`
	Content     string    `json:"content" db:"content" binding:"required"`
	CreateTime  time.Time `json:"create_time" db:"create_time"`
	CommenterID int64     `json:"commenter_id" db:"commenter_id"`
	Commenter   string    `json:"commenter" db:"commenter"`
}
type CommentDetail struct {
	PostID      string    `json:"post_id" db:"post_id"`
	Content     string    `json:"content" db:"content" `
	CommenterID string    `json:"commenter_id" db:"commenter_id"`
	Commenter   string    `json:"commenter" db:"commenter"`
	CreateTime  time.Time `json:"create_time" db:"create_time"`
}
