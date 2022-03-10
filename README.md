# lzq.admin

#### 介绍
Golang语言实现的多租户SaaS系统，项目基于gin、xorm、jwt、swagger、vue实现的前后端分离多租户后台管理系统，主要功能有多租户、用户管理、RBAC权限管理、按钮级别的权限校验及redis缓存，日志支持Elasticsearch

#### 功能
1. 多租户：同套代码，同套部署，根据租户进行数据隔离
2. 系统用户管理
3. 菜单管理
4. 操作权限管理：支持菜单、按钮级别的权限校验
5. 角色管理：支持账号多角色授权
6. 日志系统：支持本地日志文件和Elasticsearch

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

#### 可能出现的问题及解决办法
##### 1. 报node-sass错误
解决办法：
``` 
npm config set sass_binary_site=https://npm.taobao.org/mirrors/node-sass 
```

#### 界面展示
![](/tools/md-images/1646896434.png)
![](/tools/md-images/1646896535.jpg)
![](/tools/md-images/1646896551.jpg)
![](/tools/md-images/1646896572.jpg)

