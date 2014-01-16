package main

import (
	"log"
	"os"
	"os/signal"
	"runtime"
	"runtime/debug"
	"sync"
	"syscall"
	"time"
)

func DoGC() {
	runtime.GC()
	debug.FreeOSMemory()
	log.Println("Current Routines", runtime.NumGoroutine())
}

var stoped = false

func GoroutineTest() {
	count := 0
	for {
		if stoped {
			break
		}
		wg := sync.WaitGroup{}
		times := 400000
		for i := 0; i < times; i++ {
			wg.Add(1)
			go func() {
				// Nothing to do
				time.Sleep(10 * time.Second)
				wg.Done()
			}()
		}
		wg.Wait()
		count = count + times
		log.Println("You have created ", count, "routines")
	}
}

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	go GoroutineTest()
	ch := make(chan os.Signal)
	signal.Notify(ch, syscall.SIGINT, syscall.SIGTERM)
	log.Println(<-ch)
	stoped = true
	time.Sleep(3 * time.Second)
	for i := 0; i < 10; i++ {
		time.Sleep(2 * time.Second)
		DoGC()
	}
}
