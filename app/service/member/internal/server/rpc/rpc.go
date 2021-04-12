package rpc

import "go-zero-study/app/service/member/internal/service"

type Server struct {
	svc *service.Service
}

func New(svc *service.Service) *Server {
	return &Server{
		svc: svc,
	}
}
