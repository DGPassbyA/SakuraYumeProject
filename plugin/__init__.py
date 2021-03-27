from nonebot import on_command
import os

@on_command('获取网页饼干', aliases=('获取网页饼干', '获取cookie', '获取网页cookie'), only_to_me=False)
async def get_sakurayume_token(session:CommandSession):
    print(session)
    process = os.popen('go run getToken.go 2247751095 false')
    output = process.read()
    process.close()
    await session.send(output, at_sender=True)