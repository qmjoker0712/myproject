package utils

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"myproject/node"
)

// StartNode start node
func StartNode(stack *node.Node) {
	if err := stack.Start(); err != nil {
		panic(fmt.Sprintf("Error starting protocol stack: %v", err))
	}
	go func() {
		sigc := make(chan os.Signal, 1)
		signal.Notify(sigc, syscall.SIGINT, syscall.SIGTERM)
		defer signal.Stop(sigc)
		<-sigc
		fmt.Printf("Got interrupt, shutting down...")
		go stack.Stop()
		for i := 5; i > 0; i-- {
			<-sigc
			if i > 1 {
				fmt.Print("Already shutting down, interrupt more to panic.", "times", i-1)
			}
		}
	}()
}
