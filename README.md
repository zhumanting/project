# seo api 使用说明  （base on beego）

## 安装 beego 
- go get github.com/astaxie/beego 
- 安装bee工具  go get github.com/beego/bee

## bee api 安装项目
- bee api project_name


## 本地开发调试
- 把工程放到 $GOPATH/src 目录下 （注意：工程名称修改为api）
- 执行 go get 
- bee run 

## 本地获得安装包并运行
进入devops目录， 执行 docker-compose -f debug.yml up -d 

## 持续集成
jenkins  devops/deploy.sh 部署

## utils 目录为工具
- redis  使用方法： redis := utils.RedisConnection()
- mysql 初始化， 直接使用orm操作
- sys_code 统一接口错误码定义 0 成功 其他 具体错误码

## controllers 控制器层
- 增加控制器 ： bee generate controller exsample 
- context 对象
- context 对象是对 Input 和 Output 的封装，里面封装了几个方法：

```
Redirect
Abort
WriteString
GetCookie
SetCookie
```
- c.Ctx下面会有很多方法可以使用  Input  ParseForm GetString GetStrings GetInt32 GetFile ...

## models 模型层， 表模型基本方法, 
- 表名在TableName方法定义
- bee generate model  -fields="name:type"
- bee generate appcode -tables="eps_message"  -driver=mysql -conn="root:123456@tcp(192.168.50.78:3306)/chanzhi"
- orm.NewOrm()   Insert 
- //err = o.QueryTable(new(EpsArticle)).Filter("Id", 1).RelatedSel().One(vv)
- o.Raw("select * from eps_article limit 1").QueryRow(vv)

## service 具体业务逻辑层
- 项目实际使用
- httplib 使用  Post Get Put 
- req.Header() 
- req.Response()
- req.ToJSON

## routers 路由层 
- 推荐每个api单独定义方式：
```
beego.Router("/userlist",&controllers.UserController{},"get:GetAll")
```

```
// URLMapping in controller
//func (c *UserController) URLMapping() {
//	c.Mapping("Post", c.Post)
//	c.Mapping("GetOne", c.GetOne)
//	c.Mapping("GetAll", c.GetAll)
//	c.Mapping("Put", c.Put)
//	c.Mapping("Delete", c.Delete)
//}

```

## 自动api文档
- bee run -gendoc=true -downdoc=true


## 业务&使用
- 登录 controllers/user.go  Login
- 鉴权 
```
func (c *UserController) GetAll() {

    c.Auth()
	...
	...

} 
```

- 登录信息保存在 base_controller 中 
- var MyClaims jwt.MapClaims