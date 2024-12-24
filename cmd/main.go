package main

import (
	"time"

	"go-backend/database"

	user "go-backend/internal/model"

	"go-backend/internal/router"

	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"gorm.io/gorm"

	"math/rand"
)

var (
	counter = promauto.NewCounter(prometheus.CounterOpts{
		Name: "devbulls_counter",
		Help: "Counting the total number of requets being handled",
	})

	gauge = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Name: "devbulls_gauge",
		Help: "Monitoring node usage",
	}, []string{"node", "namespace"})
)

func RecordMetrics() {
	go func() {
		for {
			counter.Inc()
			gauge.WithLabelValues("node-1", "namespace-b").Set(rand.Float64())
			time.Sleep(time.Second * 5)
		}
	}()
}

var db *gorm.DB

func main() {
	RecordMetrics()

	// Initialize db connection
	db = database.GetDb()

	db.AutoMigrate(&user.User{})

	r := router.SetupRouter(db)

	r.GET("/metrics", gin.WrapH(promhttp.Handler()))

	r.Run()
}
