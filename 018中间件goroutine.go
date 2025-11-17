package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"time"
)

// 深度解读：Gin中间件中使用goroutine的问题

// WrongGoroutineHandler 1. 错误的示例：在goroutine中直接使用原始上下文
func WrongGoroutineHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println("=== 错误示例开始 ===")
		fmt.Printf("原始上下文地址: %p\n", c)

		// 错误：在goroutine中直接使用原始上下文
		go func() {
			// 模拟异步处理（比如日志记录、数据分析等）
			time.Sleep(200 * time.Millisecond)

			// 问题1：上下文可能已经被回收
			fmt.Printf("在goroutine中访问上下文地址: %p\n", c)
			fmt.Printf("在goroutine中访问URL: %s\n", c.Request.URL.Path)

			// 问题2：可能引发竞态条件
			// 主goroutine可能已经完成了请求处理，释放了上下文
			// 这里访问c的任何数据都可能导致panic或数据不一致
		}()

		c.Next()
		fmt.Println("=== 错误示例结束 ===")
	}
}

// CorrectGoroutineHandler 2. 正确的示例：使用上下文的只读副本
func CorrectGoroutineHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println("=== 正确示例开始 ===")
		fmt.Printf("原始上下文地址: %p\n", c)

		// 正确：创建上下文的只读副本
		copyCtx := c.Copy()
		fmt.Printf("副本上下文地址: %p\n", copyCtx)

		go func(ctx *gin.Context) {
			// 模拟异步处理
			time.Sleep(200 * time.Millisecond)

			// 安全：使用副本上下文
			fmt.Printf("在goroutine中访问副本上下文地址: %p\n", ctx)
			fmt.Printf("在goroutine中访问副本URL: %s\n", ctx.Request.URL.Path)

			// 副本是只读的，不会影响原始请求
			// 可以安全地读取请求信息，但不能修改响应
		}(copyCtx)

		c.Next()
		fmt.Println("=== 正确示例结束 ===")
	}
}

// RaceConditionDemoHandler 3. 演示竞态条件的示例
func RaceConditionDemoHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println("=== 竞态条件演示开始 ===")

		// 模拟多个goroutine同时访问上下文
		for i := 0; i < 3; i++ {
			go func(index int) {
				time.Sleep(time.Duration(index) * 100 * time.Millisecond)

				// 这里可能发生竞态条件
				// 不同的goroutine可能同时访问c的相同数据，导致竞态条件
				fmt.Printf("Goroutine %d 访问URL: %s\n", index, c.Request.URL.Path)
			}(i)
		}

		c.Next()
		fmt.Println("=== 竞态条件演示结束 ===")
	}
}

// SafeConcurrentHandler 4. 使用副本避免竞态条件的示例
func SafeConcurrentHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println("=== 安全并发示例开始 ===")

		// 为每个goroutine创建独立的副本
		for i := 0; i < 3; i++ {
			copyCtx := c.Copy()

			go func(index int, ctx *gin.Context) {
				time.Sleep(time.Duration(index) * 100 * time.Millisecond)

				// 每个goroutine使用自己的副本，避免竞态条件
				fmt.Printf("Goroutine %d 访问副本URL: %s\n", index, ctx.Request.URL.Path)
			}(i, copyCtx)
		}

		c.Next()
		fmt.Println("=== 安全并发示例结束 ===")
	}
}

// AsyncLoggingHandler 5. 实际应用场景：异步日志记录
func AsyncLoggingHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()

		// 创建副本用于异步日志记录
		logCtx := c.Copy()

		// 异步记录请求日志（不影响主请求的响应时间）
		go func(ctx *gin.Context, startTime time.Time) {
			// 模拟日志写入（可能是文件、数据库等）
			time.Sleep(50 * time.Millisecond)

			duration := time.Since(startTime)
			fmt.Printf("[ASYNC LOG] %s %s - %v\n",
				ctx.Request.Method,
				ctx.Request.URL.Path,
				duration)
		}(logCtx, start)

		c.Next()

		// 主请求继续处理，不等待日志记录完成
	}
}

// AsyncDataProcessingHandler 6. 实际应用场景：异步数据处理
func AsyncDataProcessingHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 创建副本用于异步数据处理
		dataCtx := c.Copy()

		// 异步处理数据（比如发送到消息队列、更新统计信息等）
		go func(ctx *gin.Context) {
			// 模拟数据处理
			time.Sleep(100 * time.Millisecond)

			// 安全地读取请求数据
			userAgent := ctx.Request.UserAgent()
			ip := ctx.ClientIP()

			fmt.Printf("[DATA PROCESSING] UserAgent: %s, IP: %s\n", userAgent, ip)
		}(dataCtx)

		c.Next()
	}
}

func main() {
	// 创建一个默认的路由引擎
	r := gin.Default()

	// 注册各种演示中间件
	r.Use(AsyncLoggingHandler())        // 异步日志记录
	r.Use(AsyncDataProcessingHandler()) // 异步数据处理

	// 测试路由
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message":   "pong",
			"timestamp": time.Now().Unix(),
		})
	})

	// 演示正确用法的路由
	r.GET("/correct", CorrectGoroutineHandler(), func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "正确使用goroutine的示例",
		})
	})

	// 演示安全并发的路由
	r.GET("/safe", SafeConcurrentHandler(), func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "安全并发处理示例",
		})
	})

	// 警告：演示错误用法的路由（实际生产环境应避免）
	r.GET("/wrong", WrongGoroutineHandler(), func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "错误使用goroutine的示例（可能引发问题）",
		})
	})

	// 警告：演示竞态条件的路由（实际生产环境应避免）
	r.GET("/race", RaceConditionDemoHandler(), func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "竞态条件演示示例（可能引发问题）",
		})
	})

	fmt.Println("服务器启动在 :8080")
	fmt.Println("访问以下路由进行测试:")
	fmt.Println("  GET /ping      - 基础测试")
	fmt.Println("  GET /correct   - 正确用法演示")
	fmt.Println("  GET /safe      - 安全并发演示")
	fmt.Println("  GET /wrong     - 错误用法演示（谨慎使用）")
	fmt.Println("  GET /race      - 竞态条件演示（谨慎使用）")

	r.Run(":8080")
}
