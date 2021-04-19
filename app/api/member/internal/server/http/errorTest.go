package http

import (
	"go-zero-study/app/api/member/internal/service"
	"go-zero-study/rest/httpx"
	"net/http"
)

func errorTest(ctx *service.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		l := service.NewMemberLogic(r.Context(), ctx)
		err := l.ErrorTest()
		httpx.JSON(r.Context(),w, nil, err)
	}
}

func errorTheadGO(ctx *service.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		l := service.NewMemberLogic(r.Context(), ctx)
		err := l.ErrorTheadGO()
		httpx.JSON(r.Context(),w, nil, err)
	}
}


func errorTheadGroup(ctx *service.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		l := service.NewMemberLogic(r.Context(), ctx)
		err := l.ErrorTheadGroup()
		httpx.JSON(r.Context(),w, nil, err)
	}
}


