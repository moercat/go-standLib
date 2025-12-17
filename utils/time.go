package utils

import (
	"fmt"
	"time"
)

// TimeFormatting 示例：日期时间格式化和解析
func TimeFormatting() {
	// 当前时间
	now := time.Now()
	fmt.Println("当前时间:", now)

	// 格式化时间
	fmt.Println("格式化时间:", now.Format("2006-01-02 15:04:05"))

	// RFC3339格式
	fmt.Println("RFC3339格式:", now.Format(time.RFC3339))

	// 解析时间字符串
	timeStr := "2023-12-25 10:30:00"
	parsedTime, err := time.Parse("2006-01-02 15:04:05", timeStr)
	if err != nil {
		fmt.Println("解析时间错误:", err)
	} else {
		fmt.Println("解析后的时间:", parsedTime)
	}
}

// TimeCalculation 示例：时间计算
func TimeCalculation() {
	now := time.Now()

	// 计算未来时间
	future := now.AddDate(0, 1, 0) // 一个月后
	fmt.Println("一个月后:", future.Format("2006-01-02 15:04:05"))

	// 计算过去时间
	past := now.AddDate(0, 0, -7) // 一周前
	fmt.Println("一周前:", past.Format("2006-01-02 15:04:05"))

	// 时间间隔
	duration := future.Sub(now)
	fmt.Printf("时间间隔: %v\n", duration)
	fmt.Printf("时间间隔(小时): %v\n", int(duration.Hours()))
}

// TimeTicker 示例：定时器和打点器
func TimeTicker() {
	// 创建一个ticker，每2秒触发一次
	ticker := time.NewTicker(2 * time.Second)
	quit := make(chan bool)

	go func() {
		for {
			select {
			case <-ticker.C:
				fmt.Println("Ticker触发:", time.Now())
			case <-quit:
				ticker.Stop()
				return
			}
		}
	}()

	// 5秒后停止ticker
	time.Sleep(5 * time.Second)
	quit <- true
	fmt.Println("Ticker已停止")
}

// TimeAfter 示例：使用time.After
func TimeAfter() {
	// 模拟定时任务
	fmt.Println("开始等待2秒...")
	<-time.After(2 * time.Second)
	fmt.Println("等待结束")
}
