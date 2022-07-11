package main

import (
	"flag"
	"fmt"
	"log"
	"net"
	"net/http"
	"runtime/debug"

	"github.com/gin-gonic/gin"
	"github.com/soobinDoson/docker-practice.git/db"
	pb "github.com/soobinDoson/docker-practice.git/proto"
	"google.golang.org/grpc"
)

type User struct {
	db IDatabase
}

type IDatabase interface {
	ListUserPartner(rq *pb.UserPartnerRequest) ([]*pb.UserPartner, error)
}

func initServe() *User {
	log.SetFlags(log.Lshortfile)
	d := &db.DB{}
	if err := d.ConnectDb("host=127.0.0.1 user=postgres password=123 dbname=docker_practice port=5432 sslmode=disable"); err != nil {
		debug.PrintStack()
		log.Panicln(err)
	}
	log.Println("Connect db success!")
	return &User{
		db: d,
	}
}

func (r *Router) HttpRouter(u *User) error {
	// ro := r.route
	// r.router()
	ro := gin.Default()
	ro.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Hello, ready to serve!",
		})
	})
	ro.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
	ro.GET("/user-partner", r.handleListUserPartner)
	go ro.Run(":3001")
	return nil
}

func StartGRPCServe(port int, u *User) error {
	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", port))
	if err != nil {
		log.Fatalln("failed to listen: %v", err)
	}
	var opts []grpc.ServerOption
	grpcServer := grpc.NewServer(opts...)
	pb.RegisterUserServiceServer(grpcServer, u)
	grpcServer.Serve(lis)
	return nil
}
