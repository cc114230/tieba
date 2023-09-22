package mysql

import "tieba/models"

func Comment(p *models.Comment) (err error) {
	sqlStr := `insert into comment(comment_id,post_id,content,commenter,commenter_id)
	values (?, ?, ?, ?,?)
	`
	_, err = db.Exec(sqlStr, p.ID, p.PostID, p.Content, p.Commenter, p.CommenterID)
	return
}

func GetCommentList() (comments []*models.CommentDetail, err error) {
	sqlStr := `select 
	post_id, content, commenter_id, commenter,create_time
	from comment
	ORDER BY create_time
	DESC 
-- 	limit ?,?
	`
	comments = make([]*models.CommentDetail, 0, 2)
	err = db.Select(&comments, sqlStr)
	//err = db.Select(&posts, sqlStr, (page-1)*size, size)
	return
}
