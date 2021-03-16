package server

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/klasrak/data-integration/config"
	"github.com/klasrak/data-integration/mongo"
	mongoDriver "go.mongodb.org/mongo-driver/mongo"
)

type Server struct {
	Router      *gin.Engine
	MongoClient *mongoDriver.Client
	Env         *config.Env
}

var HttpServer Server

func (s *Server) Initialize(mongoClient *mongoDriver.Client, cfg *config.Env) {
	s.Router = gin.Default()
	s.Env = cfg
	s.MongoClient = mongoClient

	s.InitRoutes()
}

func (s *Server) Run(addr string) {
	s.Router.Run(addr)
}

func Run() {
	HttpServer = Server{}
	appAddr := fmt.Sprintf(":%s", "8080")
	cfg := config.New()

	mongoClient, err := mongo.GetClient(cfg.MONGO_URI)

	if err != nil {
		panic(err.Error())
	}

	HttpServer.Initialize(mongoClient, cfg)

	HttpServer.Run(appAddr)
}
