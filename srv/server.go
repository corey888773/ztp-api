package srv

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

type Server struct {
	Router *gin.Engine
}

func NewServer() *Server {
	return &Server{
		Router: gin.Default(),
	}
}

func (s *Server) SetupRouter() {
	s.Router.GET("/dummy", func(context *gin.Context) {
		fmt.Println("Hello dummy")
		context.JSON(http.StatusOK, "Hello dummy")
	})
}

func (s *Server) SetupDatabase(mongoClient *mongo.Client) {
	db := mongoClient.Database("dummy")
	fmt.Println(db.Name())
}

func (s *Server) Start(httpAddress string) error {
	return s.Router.Run(httpAddress)
}
