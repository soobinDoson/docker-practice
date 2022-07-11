package main

import (
	"context"

	pb "github.com/soobinDoson/docker-practice.git/proto"
)

const DEFAULT_LIMIT = 30

func (u *User) ListUserPartners(ctx context.Context, req *pb.UserPartnerRequest) (*pb.UserPartners, error) {
	if req.GetLimit() == 0 {
		req.Limit = DEFAULT_LIMIT
	}
	up, err := u.db.ListUserPartner(req)
	if err != nil {
		return nil, err
	}
	uPs := &pb.UserPartners{
		UserPartners: up,
	}
	return uPs, nil
}
