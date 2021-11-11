package main

import (
	"log"
	"time"

	"github.com/robfig/cron"

	"gin-simple-base/models"
)

func main() {
	log.Println("Starting...")

	c := cron.New()
	c.AddFunc("* * * * * *", func() { //AddFunc 会向 Cron job runner 添加一个 func ，以按给定的时间表运行
		log.Println("Run models.CleanAllTag...")
		var u models.User
		models.UpdateUserScript(u)
	})
	c.AddFunc("* * * * * *", func() {
		log.Println("Run models.CleanAllArticle...")
		u := map[string]interface{}{}
		models.CleanUserScript(u)
	})
	//在当前执行的程序中启动 Cron 调度程序。其实这里的主体是 goroutine + for + select + timer 的调度控制
	c.Start()

	t1 := time.NewTimer(time.Second * 10)
	for {
		select {
		case <-t1.C:
			t1.Reset(time.Second * 10) //会重置定时器，让它重新开始计时
		}
	}
}
