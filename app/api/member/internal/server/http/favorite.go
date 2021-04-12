package http

import (
	"go-zero-study/app/api/member/internal/model"
	"go-zero-study/app/api/member/internal/service"
	"go-zero-study/core/ecode"
	"go-zero-study/rest/httpx"
	"net/http"
)


func addFavorite(ctx *service.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req model.AddFavoriteReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.JSON(r.Context(),w, nil, ecode.RequestErr)
			return
		}
		l := service.NewMemberLogic(r.Context(), ctx)
		resp, err := l.AddFavorite(req)
		httpx.JSON(r.Context(),w, resp, err)
	}
}