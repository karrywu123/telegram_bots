# telegram_bots
用gin框架做一个telegram机器人报警的服务端，
telegram机器人 自己从botfarther上面申请
# 运行之后 从浏览器打开：
http://127.0.0.1:8888/v1/sendmassage/123
telegram群里就发出消息123 
# 可以自己调控并发池的个数 和机器人的个数
并发池的大小 main文件的goroutinePool方法控制 
机器人个数 在set_rand的set_rand 的Rand_bot_id 方法的 reasons切片 增加机器人的id
