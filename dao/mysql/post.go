package mysql

import "tieba/models"

func CreatePost(p *models.Post) (err error) {
	sqlStr := `insert into post(
	post_id, title, content, author_id, community_id)
	values (?, ?, ?, ?, ?)
	`
	_, err = db.Exec(sqlStr, p.ID, p.Title, p.Content, p.AuthorID, p.CommunityID)
	return
}

func GetPostById(pid int64) (post *models.Post, err error) {
	post = new(models.Post)
	sqlStr := `select 
	post_id, title, content, author_id, community_id, create_time
	from post
	where post_id = ?`
	err = db.Get(post, sqlStr, pid)
	return
}

func GetPostList() (posts []*models.Post, err error) {
	sqlStr := `select 
	post_id, title, content, author_id, community_id, create_time
	from post
	ORDER BY create_time
	DESC 
-- 	limit ?,?
	`
	posts = make([]*models.Post, 0, 2)
	err = db.Select(&posts, sqlStr)
	//err = db.Select(&posts, sqlStr, (page-1)*size, size)
	return
}
