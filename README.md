

**基于 gin+gorm+mysql 的一个一言项目**
尝试了一下jwt和redis，还没实现读写分离，仅仅简单运用了下redis
尝试工程化项目目录结构
错误码是别的项目搬过来的，尝试应用了一下，但还不是很熟练

# 项目运行✨
本项目采用GOMODULE管理依赖
## 手动运行
```go
cd ./cmd/single-sentence
go run single-sentence.go
```

# 主要功能

- 用户注册登录(jwt-go)
- 用户基本信息修改，修改密码
- 句子添加删除修改
- 随机/指定类型获取句子
- 不完善的管理员

# 项目规划
- [] 完善错误检查管理，错误码做好对应
- [] 实现读写分离（这个项目好像有点不好搞）
- [] review数量还没做
- [] 前端
- [] 用户与句子的一对多关系
- [] 。。。。。。

# 主要依赖
| 名称           | 版本   |
|--------------|---------|
| golang       | 1.21.0  |
| gin          | v1.9.1  |
| gorm         | v1.25.5 |
| mysql        | V8.0.28 |
| redis        | v7.2.2  |
| jwt-go       | v3.2.0  |
| redigo       | v1.8.9 |


# 项目结构
```
├── api             #控制业务模块流程
├── cmd             # 程序入口
├── config          # 配置文件
├── middleware      # 中间件
├── pkg
│  ├── e            # 错误码
│  └── pkgX         # 对应包的工具函数
├── repository       
│   ├── dao       # dao层，对数据库进行操作
│   └── model     # 定义数据库的模型
├── routes          # 路由
├── serializer      # 与前端交互的json规范
├── service         # 控制业务模块逻辑应用
```

# 配置文件

`config/config.json` 文件配置
# 如何导入并测试接口
导入single-sentence.postman_collection.json这个就行
