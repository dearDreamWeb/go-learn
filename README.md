# gin框架学习
该仓库仅用来学习

# 安装依赖
在项目中打开终端输入`go mod tidy`，安装依赖
```
go mod tidy
```
想要安装自己需要的依赖
```
go get packageName
```
`packageName`代表要下载的依赖包

# 目录结构

```text
├── config       配置文件
├── controller   接口控制层
│         ├── v1 接口的v1版本
│         └── v2 接口的v2版本
├── logs         打印的日志文件
├── middleware   中间件，比如：日志、jwt验证
├── model        数据库
├── router       路由
└── utils        工具库
```
