package main

import (
	"sync"
	"time"

	"github.com/EricYT/promise"
	log "github.com/Sirupsen/logrus"
)

func init() {
	log.SetLevel(log.DebugLevel)
}

func main() {
	log.Println("promise test")

	var wg sync.WaitGroup

	p := promise.NewPromise()

	wg.Add(1)
	go func(p *promise.Promise) {
		log.Println("child wait the set value datatime:", time.Now())
		time.Sleep(time.Second * 3)
		log.Println("waiters length before set value:", p.Size())
		p.Set("hello,world")
		log.Println("goruntine set value datatime: ", time.Now())
		log.Println("waiters length after set value:", p.Size())
		wg.Done()
	}(p)

	wg.Add(10)
	for i := 0; i < 10; i++ {
		go func(index int) {
			//value := p.Get()
			err, value := p.GetTimeout(time.Second * 1)
			if err != nil {
				log.Errorln("test get promise error:", err)
			} else {
				log.Println("routine ", index, " datatime:", time.Now(), "value:", value)
			}
			wg.Done()
		}(i)
	}

	wg.Wait()
	return
}
