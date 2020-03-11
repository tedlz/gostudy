package files

import (
	"fmt"
	"math/rand"
	"time"
)

// Shop *
type Shop struct {
	Verbose        bool
	Cakes          int           // 蛋糕烘焙数
	BakeTime       time.Duration // 烘焙时间
	BakeStdDev     time.Duration // 烘焙时间差
	BakeBuf        int           // 烘焙和上糖之间的缓冲槽
	NumIcers       int           // 上糖的厨师数
	IceTime        time.Duration // 上糖的时间
	IceStdDev      time.Duration // 上糖时间差
	IceBuf         int           // 上糖和雕花之间的缓冲槽
	InscribeTime   time.Duration // 雕花的时间
	InscribeStdDev time.Duration // 雕花的时间差
}

type cake int

func (s *Shop) baker(baked chan<- cake) {
	for i := 0; i < s.Cakes; i++ {
		c := cake(i)
		if s.Verbose {
			fmt.Println("烘焙中……", c)
		}
		work(s.BakeTime, s.BakeStdDev)
		baked <- c
	}
	close(baked)
}

func (s *Shop) icer(iced chan<- cake, baked <-chan cake) {
	for c := range baked {
		if s.Verbose {
			fmt.Println("上糖中……", c)
		}
		work(s.IceTime, s.IceStdDev)
		iced <- c
	}
}

func (s *Shop) inscriber(iced <-chan cake) {
	for i := 0; i < s.Cakes; i++ {
		c := <-iced
		if s.Verbose {
			fmt.Println("雕花中……", c)
		}
		work(s.InscribeTime, s.InscribeStdDev)
		if s.Verbose {
			fmt.Println("完成", c)
		}
	}
}

// Work 方法模拟运行时间
func (s *Shop) Work(runs int) {
	for run := 0; run < runs; run++ {
		baked := make(chan cake, s.BakeBuf)
		iced := make(chan cake, s.IceBuf)
		go s.baker(baked)
		for i := 0; i < s.NumIcers; i++ {
			go s.icer(iced, baked)
		}
		s.inscriber(iced)
	}
}

// work 方法会阻止调用 goroutine 一段时间
// 通常分布在 d 周围
// 标准差为 stddev
func work(d, stddev time.Duration) {
	delay := d + time.Duration(rand.NormFloat64()*float64(stddev))
	time.Sleep(delay)
}
