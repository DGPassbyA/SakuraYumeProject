# SakuraYumeProject
A Princess Connect Re:Dive clan battle manage website

## Notice

本网站以[HoshinoBot](https://github.com/Ice-Cirno/HoshinoBot)为核心，加入了网页端与机器人进行数据交互，并不是独立的公会战管理系统。

原始的HoshinoBot可以安装[yobot](https://github.com/pcrbot/yobot)作为插件运行网页端会战管理，但是不能很好兼容HoshinoBot本身在QQ上的会战管理系统，因此开发这个项目作为原生HoshinoBot的补充。

## About

#### 前端：vue + axios + mdui

#### 后端：go-gin

## Progress ：Opening

### 3.21 上传基本框架，后端已实现查询出刀记录接口，前端界面框架大致完成，可以调用接口展示出刀记录

### 3.23 重构框架，关于设置的项归类到conf，后端实现出刀接口，实现token生成和验证功能

### 3-25 简单完成前端出刀功能，准备动手解决交互和BOSS信息存储的问题

### 3-27 完成生成Token的插件的编写

### 4-8 处理了一些奇怪的bug，成功部署到服务器上，实现出刀和查刀，删刀需要群里妈的指令

-----------------------

## More

目前数据交互上主要是修改HoshinoBot自身维护的sqlite数据库文件。HoshinoBot本身接收来自mirai或者cqhttp的消息，所以理论上是可以发送相同格式的包来模拟QQ群里发消息来报刀的，不过还没找到文档，先摸鱼🐟。至于安全问题，呃.......

如果要测试，只要把机器人的sqlite文件拉一份过来，改下路径conf里的`DatabasePath`就好

关于token的发放问题，应该会参照yobot的方法，通过qq发放，到时候再写一个HoshinoBot的token生成插件吧哈哈哈哈

## 部署方法

* 前端 修改相关链接 `npm run build` 生成静态文件到dist，然后复制到Apache或者Nginx目录下，保证外网能访问
* 后端 如果前端是https则后端需要布置https，需要自行申请证书。然后运行`go run main.go`即可
* 插件 按普通插件一样装给妈即可，不需要修改`__bot__.py`

