package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"sync"
	"telegram_bots/api"
	"telegram_bots/set_rand"
	"time"
)
// 定义接收数据的结构体
type Sendmessage struct {
	// binding:"required"修饰的字段，若接收为空值，则报错，是必须字段
	Message string `form:"message" json:"message" uri:"message" xml:"message" binding:"required"`
}
var chandata = make(chan string, 100)
var wg sync.WaitGroup
var lock sync.Mutex

func handfunc(c *gin.Context){
	var sendmassage Sendmessage
	if err := c.ShouldBindUri(&sendmassage); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	massage :=sendmassage.Message
	select {
	case chandata <- massage:
		break
	default:
		fmt.Println("管道已经满了，正在丢弃信息：", massage)
		break
	}
	c.String(http.StatusOK, "发送的消息为 ："+ massage)
}
func goroutinePool(n int) {
	wg.Add(n)
	for i := 1; i <= n; i++ {
		go func(gi int) {
			fmt.Printf("第%d个goroutine启动！！！\n", gi)
			for {
				select {
				case data := <-chandata:
					fmt.Printf("我是第%d 个goroutine，从管道中取出了消息：%s\n", gi, data)
					lock.Lock() // 加锁
					botid :=set_rand.Rand_bot_id()
					fmt.Println("发消息的随机机器人为 :"+botid)
          //chat_ID 自己创建telegram群的id 
					api.SendMessage(data,botid,chat_ID)
					lock.Unlock() // 解锁
				default:
					time.Sleep(time.Second * 1)
				}
			}
		}(i)
		wg.Done()
	}
	wg.Wait()
}

func main(){

	goroutinePool(5)
	r := gin.Default()
	r.GET("/v1/sendmassage/:message",handfunc)
	//监听端口默认为8080
	r.Run(":8888")
}
