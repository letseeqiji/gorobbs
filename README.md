<div align="center">
<img alt="Gorobbs" width="150rpm"  src="https://github.com/letseeqiji/gorobbs/blob/master/doc/gr.png">
 <br>
 轻而快，为实用而构建
</div>


<br><br>

</p>

## 简介

[Gorobbs](https://github.com/letseeqiji/gorobbs) 是一款轻巧的内置了全文搜索引擎的的BBS系统, 专为普通用户设计，开箱即用，无需复杂配置。我们的目标是打造最轻量化的分布式BBS系统！

## 案例

* [新书来了](https://www.xinshulaile.com)
* 妙记英语
![console](https://github.com/letseeqiji/gorobbs/blob/master/doc/miaoji.jpg)

## 功能

* 多用户BBS
* 自定义导航
* 多主题 / 多语言
* MySQL + Redis
* 内置轻巧的全文搜索引擎
* 良好的SEO优化

## 界面

### 首页

![start](https://github.com/letseeqiji/gorobbs/blob/master/doc/index.png)

### 登录后效果

![start](https://github.com/letseeqiji/gorobbs/blob/master/doc/login.png)

### 手机版

![start](https://github.com/letseeqiji/gorobbs/blob/master/doc/mobile.png)

### 编辑帖子

![console](https://github.com/letseeqiji/gorobbs/blob/master/doc/thread.png)

### 帖子详情

![post](https://github.com/letseeqiji/gorobbs/blob/master/doc/detail.png)

### 后台某页面

![post](https://github.com/letseeqiji/gorobbs/blob/master/doc/tags.png)

## 新增界面

### 第三方登录：微信登录

![post](https://github.com/letseeqiji/gorobbs/blob/master/doc/wechatlogin.png)

## 安装

### 项目依赖包
 * github.com/gin-gonic/gin
 * gopkg.in/gomail.v2
 * github.com/tommy351/gin-sessions
 * github.com/sirupsen/logrus
 * github.com/rifflock/lfshook
 * github.com/unknwon/com
 * github.com/astaxie/beego/validation
 * github.com/aviddiviner/gin-limit
 * github.com/huichen/wukong
 * github.com/go-ini/ini
 * github.com/gomodule/redigo/redis
 * github.com/lestrrat-go/file-rotatelogs
 * github.com/jinzhu/gorm
 * github.com/jinzhu/gorm/dialects/mysql
 * github.com/dgrijalva/jwt-go

### 项目环境依赖
 * golang 1.11 and above
 * mysql 5.6 and above
 * redis 5 and above
 * 若安装在生产环境，推荐使用Nginx1.16

### 本地试用


#### 方案一: 使用 gomod

- 在具有读写权限的目录克隆代码: git clone https://github.com/letseeqiji/gorobbs.git；
- 执行命令:go mod tidy;
- 解压静态包：打开 gorobb/static 将static.zip解压到static目录;
- 配置文件: 打开 gorobbs/conf/app.ini 并配置数据库和redis;
- 导入sql数据: 导入 gorobbs/gorobbs.sql 到本地MYSQL数据库;
- 运行: 进入 gorobbs 目录，运行命令: go run main.go;
- 访问地址: http://127.0.0.1:9000  端口号在配置文件中可以配置;
- 测试用用户名和密码:  地址:admin@local.com   密码:123456;

#### 方案二: 不使用 gomod

- 切换目录: 首先进入到本地的GOPATH目录；
- 克隆代码: git clone https://github.com/letseeqiji/gorobbs.git；
- 进入gorobbs目录，可以将目录下vendor目录的内容复制到GOPATH目录;
- 解压静态包：打开 gorobb/static 将static.zip解压到static目录;
- 配置文件: 打开 gorobbs/conf/app.ini 并配置数据库和redis;
- 导入sql数据: 导入 gorobbs/gorobbs.sql 到本地MYSQL数据库;
- 运行: 进入 gorobbs 目录，运行命令: go run main.go;
- 访问地址: http://127.0.0.1:9000  端口号在配置文件中可以配置;
- 测试用用户名和密码:  地址:admin@local.com   密码:123456;

#### 方案三: 使用docker

- docker容器下载：链接: https://pan.baidu.com/s/1eYq51V2ZRxIVeE9dEDtOWw 提取码: tm1c ；
- 使用dokcer import 命令加载本地容器；
- 内部已经安装好了mysql redis 和 gorobbs可执行文件;
- 数据库：端口映射9005:3306，user:root, password:abc123456;
- gorobbs: 端口映射9006:9006;
- 容器内文件路径，gorobbs：/home/gobbs, redis:/usr/local/redis;
- 测试用用户名和密码:  地址:admin@local.com   密码:123456;


## 文档

* [《提问的智慧》精读注解版](https://#)
* [用户指南](https://#)
* [开发指南](https://#)
* [主题开发指南](https://#)
* [贡献指南](https://#)

## 社区

* [讨论区](https://#)
* [报告问题](https://#)

## 授权

Gorobbs 使用 [BSD](https://#) 开源协议。

## 鸣谢

* [jQuery](https://github.com/jquery/jquery)：JavaScript 工具库，用于主题页面
* [Gin](https://github.com/gin-gonic/gin)：又快又好用的 golang HTTP web 框架
* [GORM](https://github.com/jinzhu/gorm)：极好的 golang ORM 库


---

## 开源项目推荐

