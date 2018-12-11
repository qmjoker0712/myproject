package main

import (
	"myproject/fetcher/db"
	"myproject/kmq"
	"encoding/json"
	"fmt"
	"sync"

	"github.com/s3dteam/go-toolkit/log/logruslogger"
)

func main() {
	loger := logruslogger.GetLoggerWithOptions("consumer", nil)
	c, err := kmq.NewConsumer([]string{"106.75.25.3:8620"}, loger, "top1", 0, &kmq.NewestOffset{}, kmq.SubMode_Sync)
	if err != nil {
		loger.Info("start error!", err)
	}

	kmq.SubEvent(db.Balance{}, "top1", func(data []byte) {
		var b db.Balance
		err := json.Unmarshal(data, &b)
		if err != nil {
			loger.Info("Unmarshal failed!")
		} else {
			loger.Info("=========== sub1 get msg: %#v", b)
		}
	})

	kmq.SubEvent(db.Balance{}, "tip1", func(data []byte) {
		var b db.Balance
		err := json.Unmarshal(data, &b)
		if err != nil {
			loger.Info("Unmarshal failed!")
		} else {
			loger.Info("=========== sub2 get msg: %#v", b)
		}
	})

	kmq.SubEvent(db.Balance{}, "tip1", f2)
	kmq.SubEvent(db.Balance{}, "top1", f1)

	c.Start()

	var wg sync.WaitGroup
	wg.Add(1)
	wg.Wait()
}

func f1(data []byte) {
	var b db.Balance
	err := json.Unmarshal(data, &b)
	if err != nil {
		fmt.Println("Unmarshal failed!")
	} else {
		fmt.Println("=========== f1 get msg: %#v", b)
	}
}

func f2(data []byte) {
	var b db.Balance
	err := json.Unmarshal(data, &b)
	if err != nil {
		fmt.Println("Unmarshal failed!")
	} else {
		fmt.Println("=========== f2 get msg: %#v", b)
	}
}
