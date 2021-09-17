package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"telegram_bots/api"
	"telegram_bots/jwt"
	"time"
)
// 定义接收数据的结构体
type Sendmessage struct {
	// binding:"required"修饰的字段，若接收为空值，则报错，是必须字段
	Message string `form:"message" json:"message" uri:"message" xml:"message" binding:"required"`
}
var chandata = make(chan string, 100)
var botdata =make(chan string ,3)

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

func recv(gi int,bots []string) {
		fmt.Printf("第%d个goroutine启动！！！\n", gi)
		for {
			select {
			case data := <-chandata:
				fmt.Printf("我是第%d个goroutine，从管道中取出了：%s\n", gi, data)
				select {
				case bdata := <-botdata:
					rdata,err:=api.SendMessage(data,bdata,chat_id)
					if err != nil {
						fmt.Println("Fatal error ", err.Error())
						break
					}
					fmt.Println("机器人为"+bdata+"发送的消息为",data)
					fmt.Println(rdata)
				default:
					for _,v:=range bots{
						botdata<-v
					}
				}
			default:
				time.Sleep(time.Second * 1)
			}

		}
}

func goroutinePool(n int,bots []string) {
	for i := 1; i <= n; i++ {
		go recv(i,bots)
	}
}

func main(){
	bots := make([]string, 0)
	bots = append(bots, "bot_id01","bot_id02","bot_id03")
	fmt.Println(len(botdata))
	goroutinePool(3,bots)
	//jwt测试
	aToken, rToken, _ :=jwt.GenToken(123)
	fmt.Println(aToken)
	fmt.Println(rToken)
	mc,_:=jwt.ParseToken(aToken)
	fmt.Println(mc.UserID,mc)
	//////
	r := gin.Default()
	r.GET("/v1/sendmessage/:message",handfunc)
	//监听端口默认为8080
	r.Run(":8888")
}
