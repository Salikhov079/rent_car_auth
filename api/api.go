package api

import (
	"log"

	"github.com/Salikhov079/rent_car/api/handler"
	"github.com/Salikhov079/rent_car/api/middleware"
	_ "github.com/Salikhov079/rent_car/docs"
	"github.com/casbin/casbin/v2"
	_ "github.com/gin-contrib/cors"

	"github.com/gin-gonic/gin"
	files "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @tite Voting service
// @version 1.0
// @description Voting service
// @host localhost:8081
// @BasePath /
// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authourization

func NewGin(h *handler.Handler) *gin.Engine {

	e, err := casbin.NewEnforcer("config/model.conf", "config/policy.csv")
	if err != nil {
		panic(err)
	}

	err = e.LoadPolicy()
	if err != nil {
		log.Fatal("casbin error load policy: ", err)
		panic(err)
	}

	r := gin.Default()

	r.Use(middleware.NewAuth(e))

	u := r.Group("/user")
	u.POST("/registr", h.RegisterUser)
	u.PUT("/update/:id", h.UpdateUser)
	u.DELETE("/delete/:id", h.DeleteUser)
	u.GET("/getall", h.GetAllUser)
	u.GET("/getbyid/:id", h.GetbyIdUser)
	u.POST("/login", h.LoginUser)
	url := ginSwagger.URL("swagger/doc.json")
	r.GET("/swagger/*any", ginSwagger.WrapHandler(files.Handler, url))
	return r
}
