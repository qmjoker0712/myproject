package profile

import (
	"fmt"
	"net/http"
	_ "net/http/pprof"
)

func Start(address string) {
	go func() {
		fmt.Println("start profile, address:", address)
		if err := http.ListenAndServe(address, nil); err != nil {
			fmt.Println("start profile failed:", err)
		}
	}()
}
