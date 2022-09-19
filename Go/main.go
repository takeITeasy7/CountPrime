package main

import (
	"fmt"
	"strconv"
	"sync"
	"time"
)

// 求める素数の最大値
const MAX = 1000000

func main() {
	// 時間計測開始
	start := time.Now()

	// 直列
	primeTest()

	// 並列
	// primeTestParallel()

	// 時間計測終了
	end := time.Now()
	delta := (end.Sub(start)).Seconds()
	// 結果を表示
	fmt.Println("かかった時間は" + strconv.FormatFloat(delta, 'f', 8, 64) + "秒です。")
}

// MAXまでの素数を求める（並列）
func primeTestParallel() {
	var wg sync.WaitGroup
	for i := 2; i <= MAX; i++ {
		wg.Add(1)
		go func(target int) {
			defer wg.Done()
			judgePrime(target)
		}(i)
	}
	wg.Wait()
}

// MAXまでの素数を求める（直列）
func primeTest() {
	for i := 2; i <= MAX; i++ {
		judgePrime(i)
	}
}

// 素数を求める
func judgePrime(target int) {
	for i := 2; i <= target; i++ {
		if target == i {
			// fmt.Println(target)	// 計測時はコメントアウト
		} else if target%i == 0 {
			break
		}
	}
}
