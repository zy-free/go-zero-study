package http

import (
	"fmt"
	"github.com/pkg/errors"
	"go-zero-study/app/api/member/internal/model"
	"go-zero-study/core/ecode"
	"net/http"

	"go-zero-study/app/api/member/internal/service"
	"go-zero-study/rest/httpx"
)

func loginMember(ctx *service.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req model.GetMemberByIDReq
		if err := httpx.Parse(r, &req); err != nil {
			err = errors.Wrapf(ecode.RequestErr, "loginMember Parse error(%s)", err.Error())
			httpx.JSON(r.Context(), w, nil, err)
			return
		}
		l := service.NewMemberLogic(r.Context(), ctx)
		resp, err := l.LoginMember(req)
		httpx.JSON(r.Context(), w, resp, err)
	}
}

func getMemberByID(ctx *service.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req model.GetMemberByIDReq
		if err := httpx.Parse(r, &req); err != nil {
			err = errors.Wrapf(ecode.RequestErr, "getMemberByID Parse error(%s)", err.Error())
			httpx.JSON(r.Context(), w, nil, err)
			return
		}
		l := service.NewMemberLogic(r.Context(), ctx)
		resp, err := l.GetMemberByID(req)
		httpx.JSON(r.Context(), w, resp, err)
	}
}

func getMemberByPhone(ctx *service.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req model.GetMemberByPhoneReq
		if err := httpx.Parse(r, &req); err != nil {
			err = errors.Wrapf(ecode.RequestErr, "getMemberByPhone Parse error(%s)", err.Error())
			httpx.JSON(r.Context(), w, nil, err)
			return
		}
		l := service.NewMemberLogic(r.Context(), ctx)
		resp, err := l.GetMemberByPhone(req)
		httpx.JSON(r.Context(), w, resp, err)
	}
}

func getMemberMaxAge(ctx *service.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := service.NewMemberLogic(r.Context(), ctx)
		resp, err := l.GetMemberMaxAge()
		httpx.JSON(r.Context(), w, resp, err)
	}
}

func queryMemberByName(ctx *service.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req model.QueryMemberByNameReq
		if err := httpx.Parse(r, &req); err != nil {
			err = errors.Wrapf(ecode.RequestErr, "QueryMemberByName Parse error(%s)", err.Error())
			httpx.JSON(r.Context(), w, nil, err)
			return
		}
		l := service.NewMemberLogic(r.Context(), ctx)
		resp, err := l.QueryMemberByName(req)
		httpx.JSON(r.Context(), w, resp, err)
	}
}

func queryMemberByIDs(ctx *service.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req model.QueryMemberByIdsReq
		if err := httpx.Parse(r, &req); err != nil {
			err = errors.Wrapf(ecode.RequestErr, "queryMemberByIDs Parse error(%s)", err.Error())
			httpx.JSON(r.Context(), w, nil, err)
			return
		}
		l := service.NewMemberLogic(r.Context(), ctx)
		resp, err := l.QueryMemberByIds(req)
		httpx.JSON(r.Context(), w, resp, err)
	}
}

func addMember(ctx *service.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req model.AddMemberReq
		if err := httpx.Parse(r, &req); err != nil {
			err = errors.Wrapf(ecode.RequestErr, "addMember Parse error(%s)", err.Error())
			httpx.JSON(r.Context(), w, nil, err)
			return
		}
		l := service.NewMemberLogic(r.Context(), ctx)
		resp, err := l.AddMember(req)
		httpx.JSON(r.Context(), w, resp, err)
	}
}

func initMember(ctx *service.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req model.InitMemberReq
		if err := httpx.Parse(r, &req); err != nil {
			err = errors.Wrapf(ecode.RequestErr, "InitMemberReq Parse error(%s)", err.Error())
			httpx.JSON(r.Context(), w, nil, err)
			return
		}
		l := service.NewMemberLogic(r.Context(), ctx)
		err := l.InitMember(req)
		httpx.JSON(r.Context(), w, nil, err)
	}
}

func batchAddMember(ctx *service.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req model.BatchAddMemberReq
		if err := httpx.Parse(r, &req); err != nil {
			err = errors.Wrapf(ecode.RequestErr, "batchAddMember Parse error(%s)", err.Error())
			httpx.JSON(r.Context(), w, nil, err)
			return
		}
		l := service.NewMemberLogic(r.Context(), ctx)
		err := l.BatchAddMember(req)
		httpx.JSON(r.Context(), w, nil, err)
	}
}

func updateMember(ctx *service.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req model.UpdateMemberReq
		if err := httpx.Parse(r, &req); err != nil {
			err = errors.Wrapf(ecode.RequestErr, "updateMember Parse error(%s)", err.Error())
			httpx.JSON(r.Context(), w, nil, err)
			return
		}
		l := service.NewMemberLogic(r.Context(), ctx)
		err := l.UpdateMember(req)
		httpx.JSON(r.Context(), w, nil, err)
	}
}

func updateSomeMember(ctx *service.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req model.UpdateSomeMemberReq
		if err := httpx.Parse(r, &req); err != nil {
			err = errors.Wrapf(ecode.RequestErr, "updateSomeMember Parse error(%s)", err.Error())
			httpx.JSON(r.Context(), w, nil, err)
			return
		}
		fmt.Println(req)
		l := service.NewMemberLogic(r.Context(), ctx)
		err := l.UpdateSomeMember(req)
		httpx.JSON(r.Context(), w, nil, err)
	}
}

func setMember(ctx *service.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req model.SetMemberReq
		if err := httpx.Parse(r, &req); err != nil {
			err = errors.Wrapf(ecode.RequestErr, "setMember Parse error(%s)", err.Error())
			httpx.JSON(r.Context(), w, nil, err)
			return
		}
		l := service.NewMemberLogic(r.Context(), ctx)
		err := l.SetMember(req)
		httpx.JSON(r.Context(), w, nil, err)
	}
}

func sortMember(ctx *service.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req model.SortMemberReq
		if err := httpx.Parse(r, &req); err != nil {
			err = errors.Wrapf(ecode.RequestErr, "sortMember Parse error(%s)", err.Error())
			httpx.JSON(r.Context(), w, nil, err)
			return
		}
		l := service.NewMemberLogic(r.Context(), ctx)
		err := l.SortMember(req)
		httpx.JSON(r.Context(), w, nil, err)
	}
}


func delMember(ctx *service.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req model.DelMemberReq
		if err := httpx.Parse(r, &req); err != nil {
			err = errors.Wrapf(ecode.RequestErr, "delMember Parse error(%s)", err.Error())
			httpx.JSON(r.Context(), w, nil, err)
			return
		}
		l := service.NewMemberLogic(r.Context(), ctx)
		err := l.DelMember(req)
		httpx.JSON(r.Context(), w, nil, err)
	}
}
