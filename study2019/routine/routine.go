package routine

import (
	"fmt"
	"sync"
	"time"
)

func Say(wait *sync.WaitGroup, s string) {
	defer wait.Done()

	for i := 0; i < 5; i++ {
		fmt.Println(s)
		time.Sleep(500 * time.Millisecond)
	}

	wait.Done()
}
