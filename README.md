# lzq.admin

#### 介绍
Golang语言实现的多租户SaaS系统，项目基于gin、xorm、jwt、swagger、vue实现的前后端分离多租户后台管理系统，主要功能有多租户、公司、部门、用户管理、RBAC权限管理、按钮级别的权限校验及redis缓存、七牛云存储、接口请求日志、登录日志等，日志支持Elasticsearch

#### 功能
1. 多租户：同套代码，同套部署，根据租户进行数据隔离
2. 公司、部门
3. 系统用户管理
4. 菜单管理
5. 操作权限管理：支持菜单、按钮级别的权限校验
6. 角色管理：支持账号多角色授权
7. 日志系统：支持本地日志文件和Elasticsearch
8. 配置：七牛云存储等各种配置
9. 接口请求日志、登录日志
10. 字典功能

#### 交流群
1. QQ群：823184304

#### 软件架构
软件架构说明


#### 安装教程
##### 开发环境搭建
1. 安装MySql数据库>=5.7（必须）
2. 安装Redis数据库（必须）
3. 安装ELK
4. 安装Node（必须）

##### 后端
1. 安装golang及搭建golang开发环境
2. 安装Goland或VisualStudio Code开发工具
3. 复制tools目录下的hsf_basic_dev.sql脚本到自己的mysql数据库中执行
4. 使用golang开发工具运行代码或者进入项目的lzq-admin目录中，使用操作系统命令方式运行golang项目(go run main.go)

##### 前端
```
1. cd lzq-admin-vue
2. npm install --registry=https://registry.npm.taobao.org
3. npm run dev
```
##### 数据库
1. 建好mysql数据库，创建一个utf8的数据库
2. 在配置文件中配置好数据库信息
3. 生产数据库及表格：将配置文件中的database-》is_migration改成true，表示ORM根据当前model结构体和数据库对比，执行更新数据库表格的脚本
4. 初始化数据库脚本表格，根据文件夹tools-》database-init中的sql文件初始化数据

#### 可能出现的问题及解决办法
##### 1. 报node-sass错误
解决办法：
``` 
npm config set sass_binary_site=https://npm.taobao.org/mirrors/node-sass 
```

#### 界面展示
![](https://gitee.com/htd_laozhiqing/lzq.admin/blob/master/tools/md-images/1646896434.png)
![](https://gitee.com/htd_laozhiqing/lzq.admin/blob/master/tools/md-images/1646896535.jpg)
![](https://gitee.com/htd_laozhiqing/lzq.admin/blob/master/tools/md-images/1646896551.jpg)
![](https://gitee.com/htd_laozhiqing/lzq.admin/blob/master/tools/md-images/1654616733315.jpg)
![](https://gitee.com/htd_laozhiqing/lzq.admin/blob/master/tools/md-images/1654616495767.jpg)
![](https://gitee.com/htd_laozhiqing/lzq.admin/blob/master/tools/md-images/1653234523(1).jpg)
![](https://gitee.com/htd_laozhiqing/lzq.admin/blob/master/tools/md-images/16532343051.jpg)
![](https://gitee.com/htd_laozhiqing/lzq.admin/blob/master/tools/md-images/16532342055.jpg)
![](https://gitee.com/htd_laozhiqing/lzq.admin/blob/master/tools/md-images/16532343631.jpg)
![](https://gitee.com/htd_laozhiqing/lzq.admin/blob/master/tools/md-images/16532343631.jpg)
<img src="https://gitee.com/htd_laozhiqing/lzq.admin/blob/master/tools/md-images/16532343631.jpg" />
<img src="/tools/md-images/16532343631.jpg" />




