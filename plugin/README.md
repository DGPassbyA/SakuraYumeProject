## Token生成插件

使用方法

1. 自行打包main.go 成可执行文件（这里我打包成linux-amd64 的 main 文件）
2. modules 文件夹下新建 sakurayume-web 文件夹，并将 main（如果你是main.exe，请修改sakurayume.py里该文件的路径）和 sakurayume.py 置于其中，在`__bot__.py`中添加 `sakurayume-web`插件
3. 重启机器人
4. 向机器人私聊`获取网页饼干`可以得到回复