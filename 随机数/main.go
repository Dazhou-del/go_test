package main

import (
	"crypto/rand"
	"fmt"
	"math/big"
	"strconv"
	"sync"
	"time"
)

// IDGenerator 用于生成10位数的自增ID
type IDGenerator struct {
	mu       sync.Mutex // 互斥锁，保证线程安全
	counter  uint64     // 自增计数器
	lastTime int64      // 上一次生成ID的时间戳
	baseTime int64      // 基准时间的时间戳
}

// NewIDGenerator 创建一个新的ID生成器
func NewIDGenerator(baseYear int) *IDGenerator {
	// 使用基准时间作为初始时间戳
	baseTime := time.Date(baseYear, time.January, 1, 0, 0, 0, 0, time.UTC).Unix()
	return &IDGenerator{
		lastTime: baseTime,
		baseTime: baseTime,
	}
}

// NextID 生成下一个10位数的自增ID
func (g *IDGenerator) NextID() uint64 {
	g.mu.Lock()
	defer g.mu.Unlock()

	// 获取当前时间戳
	currentTime := time.Now().Unix()

	// 如果时间没有变化，增加计数器
	if currentTime == g.lastTime {
		g.counter++
	} else {
		// 时间变化，重置计数器
		g.lastTime = currentTime
		g.counter = 0
	}

	// 计算时间戳部分
	timestampPart := currentTime - g.baseTime

	// 确保时间戳部分不会超过10位
	if timestampPart > 999999 {
		timestampPart = 999999
	}

	// 计算最终的10位数ID
	id := uint64(uint64(2024000000) + uint64(timestampPart*1000) + g.counter)

	return id
}

func main() {
	fmt.Println(funcName())
	fmt.Println(funcName())
	fmt.Println(funcName())
	fmt.Println(funcName())
	fmt.Println(funcName())
	fmt.Println(funcName())
}

func funcName() string {
	formattedTime := time.Now().Format("2006010215040501")
	// 提取前4位和后6位

	// 组合前4位和后6位

	//randomNumber := rand.New(rand.NewSource(time.Now().UnixNano())).Int31n(1000000)
	//fmt.Println("ran", randomNumber)

	number, err := generateFixedSizeRandomNumber(10)
	if err != nil {
		return ""
	}
	fmt.Println(number)

	randomNumberStr := strconv.Itoa(number)

	result := formattedTime + randomNumberStr

	return result
}

// generateFixedSizeRandomNumber 生成一个指定位数的随机数
func generateFixedSizeRandomNumber(size int) (int, error) {
	min := intPow(10, size-1)
	max := intPow(10, size) - 1

	// 生成一个在min和max之间的随机数
	randomNum, err := rand.Int(rand.Reader, big.NewInt(int64(max-min+1)))
	if err != nil {
		return 0, err
	}

	return int(randomNum.Int64()) + min, nil
}

func intPow(base, exp int) int {
	result := 1
	for i := 0; i < exp; i++ {
		result *= base
	}
	return result
}
