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
