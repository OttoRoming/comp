package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"sync"
	"sync/atomic"
	"time"
)

func checkNumber(num int) int {
	str := strconv.Itoa(num)
	for i := 1; i < len(str); i++ {
		if len(str)%i != 0 {
			continue
		}

		times := len(str) / i
		part := str[:len(str)/times]

		if strings.Repeat(part, times) == str {
			return num
		}
	}
	return 0
}

func checkRange(result *uint64, r string) {
	split := strings.Split(r, "-")
	start, _ := strconv.Atoi(split[0])
	end, _ := strconv.Atoi(split[1])

	var num uint64 = 0

	for i := start; i <= end; i++ {
		num += uint64(checkNumber(i))
	}

	atomic.AddUint64(result, num)
}

func solve(input string) uint64 {
	var result uint64 = 0

	var wg sync.WaitGroup
	for r := range strings.SplitSeq(input, ",") {
		wg.Add(1)
		go func(rr string) {
			defer wg.Done()
			checkRange(&result, rr)
		}(r)
	}

	wg.Wait()

	return result
}

func main() {
	start := time.Now().UnixMilli()
	input := os.Args[1]
	fmt.Println(solve(input))
	end := time.Now().UnixMilli()
	fmt.Printf("time: %dms\n", end-start)
}
