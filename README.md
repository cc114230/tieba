# 基于go+gin实现的仿百度贴吧论坛项目
## 功能如下：
### 用户注册登录，选择社区发帖，点赞点踩，用户签到，用户评论。。。
### 后续功能完善中


## 请按如下顺序启动项目

1. 根据实际情况修改 conf/config.yaml 文件中 MySQL 和 Redis 部分的配置
2. 连接上你的MySQL数据库，自行建库，按顺序依次执行model文件夹下的table.sql文件，完成建表和导入初始数据
3. 执行 `go run main.go`，启动程序
### 使用postman进行接口测试
![postman](https://img1.imgtp.com/2023/09/22/Mw34C9PE.png "postman图片")
## 注意事项

缺少包依赖执行 `go mod tidy `


