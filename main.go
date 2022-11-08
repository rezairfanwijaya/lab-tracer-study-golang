package main

import (
	"fmt"
	"log"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/rezairfanwijaya/lab-tracer-study-golang/coordinat"
	"github.com/rezairfanwijaya/lab-tracer-study-golang/handler"
	"github.com/rezairfanwijaya/lab-tracer-study-golang/helper"
	"github.com/rezairfanwijaya/lab-tracer-study-golang/tracer"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	// get env value
	env, err := helper.GetENV()
	if err != nil {
		log.Println(err)
	}
	username := env["DATABASE_USERNAME"]
	host := env["DATABASE_HOST"]
	port := env["DATABASE_PORT"]
	dbName := env["DATABASE_NAME"]

	// open onnection
	dsn := fmt.Sprintf("%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", username, host, port, dbName)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err.Error())
	}

	// migration table
	if err = db.AutoMigrate(&tracer.Tracer{}); err != nil {
		log.Fatal(err)
		return
	}

	// coordinate service
	coordinateService := coordinat.NewCoordinateService()

	// tracer repo
	tracerRepo := tracer.NewTracerRepository(db)
	tracerService := tracer.NewTracerService(tracerRepo, coordinateService)
	tracerHandler := handler.NewTracerHandler(tracerService)

	// router
	router := gin.Default()
	// enable cors
	router.Use(cors.Default())

	// endpoint
	router.GET("/tracers", tracerHandler.GetAll)
	router.POST("/tracer", tracerHandler.SaveTracer)
	router.GET("/tracers/city", tracerHandler.GetAllTracerBaseCity)

	// run  server
	if err = router.Run("localhost:9090"); err != nil {
		log.Fatal(err)
	}

}
