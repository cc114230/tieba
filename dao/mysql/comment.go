package mysql

import "tieba/models"

func Comment(p *models.Comment) (err error) {
	sqlStr := `insert into comment(comment_id,post_id,content,commenter,commenter_id)
	values (?, ?, ?, ?,?)
	`
	_, err = db.Exec(sqlStr, p.ID, p.PostID, p.Content, p.Commenter, p.CommenterID)
	return
}
