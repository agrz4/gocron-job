package main

import (
	"fmt"
	"time"

	"github.com/go-co-op/gocron"
)

func main() {
	local, err := time.LoadLocation("Asia/Jakarta")
	if err != nil {
		local = time.UTC
	}

	// create a scheduler
	s := gocron.NewScheduler(local)

	// edit for your schedule
	s.Every(1).Day().At("17:55").Do(func() {
		fmt.Println("Testing Task")
	})

	s.StartBlocking()
}
