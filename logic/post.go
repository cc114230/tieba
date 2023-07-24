package logic

import (
	"go.uber.org/zap"
	"tieba/dao/mysql"
	"tieba/dao/redis"
	"tieba/models"
	"tieba/pkg/snowflake"
)

func CreatePost(p *models.Post) (err error) {
	// 1. 生成post id
	p.ID = snowflake.GenID()
	// 2. 保存到数据库
	err = mysql.CreatePost(p)
	if err != nil {
		return err
	}
	err = redis.CreatePost(p.ID)
	return
}

func GetPostById(pid int64) (data *models.ApiPostDetail, err error) {
	//data = new(models.ApiPostDetail)
	//组合接口想用的数据
	post, err := mysql.GetPostById(pid)
	if err != nil {
		zap.L().Error("mysql.GetPostById(pid) failed", zap.Error(err))
		return
	}
	//根据作者id查询作者账号名称
	user, err := mysql.GetUserById(post.AuthorID)
	if err != nil {
		zap.L().Error("mysql.GetUserById(post.AuthorID) failed", zap.Error(err))
		return
	}

	//根据社区id查询社区详情信息
	community, err := mysql.GetCommunityDetailByID(post.CommunityID)
	if err != nil {
		zap.L().Error("mysql.GetCommunityDetailByID(post.CommunityID) failed", zap.Error(err))
		return
	}
	data = &models.ApiPostDetail{
		AuthorName:      user.Username,
		Post:            post,
		CommunityDetail: community,
	}

	//data.AuthorName = user.Username
	//data.Post = post
	//data.CommunityDetail = community

	return
}

// GetPostList 获取帖子列表
func GetPostList() (data []*models.ApiPostDetail, err error) {
	posts, err := mysql.GetPostList()
	if err != nil {
		return nil, err
	}
	data = make([]*models.ApiPostDetail, 0, len(posts))

	for _, post := range posts {
		//根据作者id查询作者账号名称
		user, errs := mysql.GetUserById(post.AuthorID)
		if errs != nil {
			zap.L().Error("mysql.GetUserById(post.AuthorID) failed",
				zap.Error(errs))
			continue
		}

		//根据社区id查询社区详情信息
		community, err := mysql.GetCommunityDetailByID(post.CommunityID)
		if err != nil {
			zap.L().Error("mysql.GetCommunityDetailByID(post.CommunityID) failed",
				zap.Error(err))
			continue
		}
		postDetail := &models.ApiPostDetail{
			AuthorName:      user.Username,
			Post:            post,
			CommunityDetail: community,
		}
		data = append(data, postDetail)
	}
	return
}
