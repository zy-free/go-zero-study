package http

import (
	"go-zero-study/app/api/comments/internal/service"
	"go-zero-study/rest/httpx"
	"net/http"

	"go-zero-study/app/api/comments/internal/model"
)

func createCommentHandler(ctx *service.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req model.CreateCommentReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.JSON(r.Context()  ,w, nil,err)
			return
		}

		l := service.NewCommentLogic(r.Context(), ctx)
		resp, err := l.CreateComment(req)
		if err != nil {
			httpx.JSON(r.Context()  ,w, nil,err)
		} else {
			httpx.JSON(r.Context(), w, resp, nil)
		}
	}
}

func batchQueryCommentsHandler (ctx *service.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req model.BatchQueryCommentsReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.JSON(r.Context()  ,w, nil,err)
			return
		}

		l := service.NewCommentLogic(r.Context(), ctx)
		resp, err := l.BatchQueryComments(req)
		if err != nil {
			httpx.JSON(r.Context()  ,w, nil,err)
		} else {
			httpx.JSON(r.Context(), w, resp, nil)
		}
	}
}