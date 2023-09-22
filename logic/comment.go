package logic

import (
	"tieba/dao/mysql"
	"tieba/models"
	"tieba/pkg/snowflake"
)

func Comment(p *models.Comment) (err error) {
	//1.生成comment id
	p.ID = snowflake.GenID()
	// 2. 保存到数据库
	err = mysql.Comment(p)
	return
}
func GetCommentList() (data []*models.CommentDetail, err error) {
	comments, err := mysql.GetCommentList()
	if err != nil {
		return nil, err
	}
	data = make([]*models.CommentDetail, 0, len(comments))

	for _, comment := range comments {
		postId := comment.PostID
		content := comment.Content
		commenterId := comment.CommenterID
		commenter := comment.Commenter
		createTime := comment.CreateTime
		commentDetail := &models.CommentDetail{
			PostID:      postId,
			Content:     content,
			CommenterID: commenterId,
			Commenter:   commenter,
			CreateTime:  createTime,
		}
		data = append(data, commentDetail)
	}
	return
}
