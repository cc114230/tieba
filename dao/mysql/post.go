package mysql

import (
	"github.com/jmoiron/sqlx"
	"strings"
	"tieba/models"
)

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

func GetPostListByIDs(ids []string) (postList []*models.Post, err error) {
	sqlStr := `select post_id,title,content,author_id,community_id,create_time
			from post 
			where post_id in (?)
			order by FIND_IN_SET(post_id,?)` //根据传进来的的帖子id列表进行排序

	//FIND_IN_SET(value, list)
	//value: 这是要查找的值，通常是一个单个的元素。
	//list: 这是一个以逗号分隔的字符串列表，用于在其中查找值。通常，这是一个逗号分隔的字符串，例如：'1,2,3,4'。

	//sqlx.In 函数的作用是将帖子ID列表转换为适用于SQL查询的格式，并生成带有占位符的查询语句。
	//然后，通过调用 db.Rebind 函数，将占位符替换为实际的帖子ID列表，以保证在SQL查询中使用。
	query, args, err := sqlx.In(sqlStr, ids, strings.Join(ids, ","))
	//如果 ids 是 [1, 2, 3]，那么 strings.Join(ids, ",") 将返回字符串 "1,2,3"。
	//SELECT post_id, title, content, author_id, community_id, create_time
	//FROM post
	//WHERE post_id IN (1, 2, 3)
	//ORDER BY FIND_IN_SET(post_id, '1,2,3')
	if err != nil {
		return
	}
	query = db.Rebind(query)
	err = db.Select(&postList, query, args...)
	return
}
