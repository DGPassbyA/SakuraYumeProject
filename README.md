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

-----------------------

## More

目前数据交互上主要是修改HoshinoBot自身维护的sqlite数据库文件。HoshinoBot本身接收来自mirai或者cqhttp的消息，所以理论上是可以发送相同格式的包来模拟QQ群里发消息来报刀的，不过还没找到文档，先摸鱼🐟。至于安全问题，呃.......

如果要测试，只要把机器人的sqlite文件拉一份过来，改下路径conf里的`DatabasePath`就好

关于token的发放问题，应该会参照yobot的方法，通过qq发放，到时候再写一个HoshinoBot的token生成插件吧哈哈哈哈