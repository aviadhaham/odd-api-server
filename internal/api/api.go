package api

import (
	"io"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/aviadhaham/odd-api-server/internal/utils"
	"github.com/gin-gonic/gin"
)

type server struct {
	port   string
	logger io.Writer
}

func NewServer(port string, logger io.Writer) *server {
	return &server{
		port:   port,
		logger: logger,
	}
}

func NewRouter(logger io.Writer) *gin.Engine {
	r := gin.Default()
	api := r.Group("/")
	{
		api.GET("/odd", func(c *gin.Context) {
			rand_int_num := utils.GetRandomOddNumber()
			num := strconv.Itoa(rand_int_num)
			res := gin.H{"number": num}
			c.JSON(http.StatusOK, res)

			log.SetOutput(logger)
			log.Printf("Response: %s", num)
		})

		api.GET("/ready", func(c *gin.Context) {
			oddFilePath := "/tmp/odd-logs.txt"
			if _, err := os.Stat(oddFilePath); os.IsNotExist(err) {
				c.JSON(http.StatusServiceUnavailable, gin.H{
					"status": "not ready",
					"error":  "odd file isn't found",
				})
				return
			}
			c.JSON(http.StatusOK, gin.H{
				"status": "ready",
			})
		})
	}

	return r
}

func (s *server) Run() {
	r := NewRouter(s.logger)

	err := r.Run(":" + s.port)
	if err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
