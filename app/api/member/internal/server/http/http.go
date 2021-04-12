package http

import (
	"go-zero-study/app/api/member/internal/service"
	"go-zero-study/core/logx"
	nmd "go-zero-study/core/metadata"
	"go-zero-study/rest"
	"net/http"
)

func greetMiddleware1(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		logx.WithContext(r.Context()).Info("greetMiddleware1 request ... ")
		// todo remove
		ctx := nmd.NewContext(r.Context(), nmd.MD{nmd.Mid: int64(4), nmd.Color: "green"})
		r = r.WithContext(ctx)
		next(w, r)
		logx.WithContext(r.Context()).Info("greetMiddleware1 reponse ... ")
	}
}

func RegisterHandlers(engine *rest.Server, serverCtx *service.ServiceContext) {
	apiPath := "/x/api"
	engine.AddRoutes(
		rest.WithMiddleware(greetMiddleware1,
			[]rest.Route{
				{http.MethodGet, apiPath + "/members/:id", getMemberByID(serverCtx)},
				{http.MethodGet, apiPath + "/members/getByPhone", getMemberByPhone(serverCtx)},
				{http.MethodGet, apiPath + "/members/maxAge", getMemberMaxAge(serverCtx)},
				{http.MethodGet, apiPath + "/members/queryByName", queryMemberByName(serverCtx)},
				{http.MethodGet, apiPath + "/members/queryByIds", queryMemberByIDs(serverCtx)},
				//{http.MethodGet,apiPath + "/members",listMemberByID(serverCtx)},
				{http.MethodPost, apiPath + "/members", addMember(serverCtx)},
				{http.MethodPost, apiPath + "/members/init", initMember(serverCtx)},
				{http.MethodPost, apiPath + "/members/batch", batchAddMember(serverCtx)},
				{http.MethodPost, apiPath + "/members/:id/update", updateMember(serverCtx)},
				{http.MethodPost, apiPath + "/members/:id/updateSome", updateSomeMember(serverCtx)},
				{http.MethodPost, apiPath + "/members/:id/set", setMember(serverCtx)},
				{http.MethodPost,apiPath + "/members/sort",sortMember(serverCtx)},
				{http.MethodDelete, apiPath + "/members/:id", delMember(serverCtx)},

				{http.MethodPost, apiPath + "/favorites", addFavorite(serverCtx)},
				{http.MethodGet, apiPath + "/test/error", errorTest(serverCtx)},

				// 刷新token，用旧token换取新token
				//{http.MethodGet, apiPath + "/refresh_token", refreshToken(serverCtx)},

			}...,
		),
	)
	engine.AddRoutes(
		[]rest.Route{
			{http.MethodPost, apiPath + "/members/:id/login", loginMember(serverCtx)},
		},
	)
}
