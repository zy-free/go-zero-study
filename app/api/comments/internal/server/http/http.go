package http

import (
	"go-zero-study/app/api/comments/internal/service"
	nmd "go-zero-study/core/metadata"
	"go-zero-study/rest"
	"net/http"
)


func greetMiddleware1(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// todo remove
		ctx := nmd.NewContext(r.Context(), nmd.MD{nmd.Mid: "4", nmd.Color: r.Header.Get("color")})
		r = r.WithContext(ctx)
		next(w, r)
	}
}


func RegisterHandlers(engine *rest.Server, serverCtx *service.ServiceContext) {
	engine.Use(greetMiddleware1)
	engine.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodPost,
				Path:    "/comments",
				Handler: createCommentHandler(serverCtx),
			},
			{
				Method: http.MethodGet,
				Path: "/comments/:room_id",
				Handler: batchQueryCommentsHandler(serverCtx),
			},
		},
	)
}
