package server

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
	"github.com/klasrak/data-integration/config"
	"github.com/klasrak/data-integration/mongo"
	diRedis "github.com/klasrak/data-integration/redis"
	mongoDriver "go.mongodb.org/mongo-driver/mongo"
)

type Server struct {
	Router      *gin.Engine
	MongoClient *mongoDriver.Client
	RedisClient *redis.Client
	Env         *config.Env
}

var HttpServer Server

func (s *Server) Initialize(mongoClient *mongoDriver.Client, redisClient *redis.Client, cfg *config.Env) {
	s.Router = gin.Default()
	s.Env = cfg
	s.MongoClient = mongoClient
	s.RedisClient = redisClient

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

	redisClient := diRedis.NewRedisDB(cfg.REDIS_HOST, cfg.REDIS_PORT)

	HttpServer.Initialize(mongoClient, redisClient, cfg)

	HttpServer.Run(appAddr)
}
