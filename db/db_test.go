package db

import (
	"log"
	"testing"

	pb "github.com/soobinDoson/docker-practice.git/proto"
)

func Test_connection(t *testing.T) {
	d := &DB{}
	err := d.ConnectDb("host=127.0.0.1 user=postgres password=123 dbname=docker_practice port=5432 sslmode=disable")
	if err != nil {
		log.Print(err)
	}
}

func Test_listPartners(t *testing.T) {
	d := &DB{}
	err := d.ConnectDb("host=127.0.0.1 user=postgres password=123 dbname=docker_practice port=5432 sslmode=disable")
	if err != nil {
		log.Print(err)
	}
	list, err := d.ListUserPartner(&pb.UserPartnerRequest{Limit: 5})
	log.Print("list: ", list)
}
