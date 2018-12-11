package main

import (
	"myproject/fetcher/db"
	"myproject/kmq"
	"fmt"

	"github.com/s3dteam/go-toolkit/log/logruslogger"
)

func main() {
	loger := logruslogger.GetLoggerWithOptions("producer", nil)
	p, err := kmq.NewProducer([]string{"106.75.25.3:8620"}, loger)
	if err != nil {
		fmt.Println("start error!", err)
	}

	testData := db.Balance{
		UserAddress: "0x123",
	}

	p.SendMsg("top1", testData)

	p.Stop()
}
