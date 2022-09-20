package main

import (
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/pyroscope-io/client/pyroscope"
)

func init() {
	startPyroscopeClient()
}

func fibonacciRecursion(n int) int {
	if n <= 1 {
		return n
	}
	return fibonacciRecursion(n-1) + fibonacciRecursion(n-2)
}

func main() {
	router := gin.New()
	router.GET("/fibo/:number", func(c *gin.Context) {
		number := c.Param("number")
		n, err := strconv.Atoi(number)
		if err != nil {
			c.Status(400)
			return
		}
		start := time.Now()
		result := fibonacciRecursion(n)
		elapsed := time.Since(start)
		c.JSON(200,
			gin.H{
				"givin":      n,
				"result":     result,
				"elapsed_ms": elapsed.Milliseconds(),
			},
		)
	})
	router.Run(":8080")
}

func startPyroscopeClient() {
	pyroscope.Start(pyroscope.Config{
		ApplicationName: "simple.golang.app",

		// replace this with the address of pyroscope server
		ServerAddress: "http://localhost:4040",

		// you can disable logging by setting this to nil
		Logger: pyroscope.StandardLogger,

		// optionally, if authentication is enabled, specify the API key:
		// AuthToken: os.Getenv("PYROSCOPE_AUTH_TOKEN"),

		ProfileTypes: []pyroscope.ProfileType{
			// these profile types are enabled by default:
			pyroscope.ProfileCPU,
			pyroscope.ProfileAllocObjects,
			pyroscope.ProfileAllocSpace,
			pyroscope.ProfileInuseObjects,
			pyroscope.ProfileInuseSpace,

			// these profile types are optional:
			pyroscope.ProfileGoroutines,
			pyroscope.ProfileMutexCount,
			pyroscope.ProfileMutexDuration,
			pyroscope.ProfileBlockCount,
			pyroscope.ProfileBlockDuration,
		},
	})

}
