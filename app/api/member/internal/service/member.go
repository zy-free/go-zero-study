package service

import (
	"context"
	"fmt"
	"github.com/pkg/errors"
	"go-zero-study/app/api/member/internal/model"
	memberProto "go-zero-study/app/service/member/api"
	"go-zero-study/core/ecode"
	"go-zero-study/core/logx"
	"go-zero-study/core/stringx"
	"go-zero-study/rest/token"
	"time"
)

type MemberLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *ServiceContext
}

func NewMemberLogic(ctx context.Context, svcCtx *ServiceContext) MemberLogic {
	return MemberLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *MemberLogic) LoginMember(req model.GetMemberByIDReq) (*model.LoginMemberResp, error) {

	now := time.Now().Unix()
	accessExpire := int64(10)
	secret := "ajd1283ida9=12391csi"
	resp, err := l.svcCtx.MemberCli.GetMemberByID(l.ctx, &memberProto.GetMemberByIDReq{
		Id: req.ID,
	})
	if err != nil {
		return nil, ecode.ErrQuery
	}
	payloads := make(map[string]interface{})
	payloads["mid"] = resp.Member.Id
	payloads["name"] = resp.Member.Name
	payloads["phone"] = resp.Member.Phone
	accessToken, err := token.GenToken(now, secret, payloads, accessExpire)
	if err != nil {
		return nil, err
	}
	return &model.LoginMemberResp{
		AccessToken:  accessToken,
		AccessExpire: now + accessExpire,
		RefreshAfter: now + accessExpire/2,
	}, nil
}

func (l *MemberLogic) GetMemberByID(req model.GetMemberByIDReq) (*model.GetMemberResp, error) {
	resp, err := l.svcCtx.MemberCli.GetMemberByID(l.ctx, &memberProto.GetMemberByIDReq{
		Id: req.ID,
	})
	if err != nil {
		return nil, ecode.ErrQuery
	}
	//time.Unix(resp.Member.CreatedAt,0)
	return &model.GetMemberResp{Member: model.Member{
		ID:        resp.Member.Id,
		Phone:     resp.Member.Phone,
		Name:      resp.Member.Name,
		Age:       resp.Member.Age,
		Address:   resp.Member.Address,
		CreatedAt: resp.Member.CreatedAt,
		UpdatedAt: resp.Member.UpdatedAt,
	}}, nil
}

func (l *MemberLogic) GetMemberMaxAge() (*model.GetMemberMaxAgeResq, error) {
	resp, err := l.svcCtx.MemberCli.GetMemberMaxAge(l.ctx, &memberProto.GetMemberMaxAgeReq{
	})
	if err != nil {
		return nil, ecode.ErrQuery
	}
	return &model.GetMemberMaxAgeResq{Age:resp.Age}, nil
}

func (l *MemberLogic) GetMemberByPhone(req model.GetMemberByPhoneReq) (*model.GetMemberResp, error) {
	resp, err := l.svcCtx.MemberCli.GetMemberByPhone(l.ctx, &memberProto.GetMemberByPhoneReq{
		Phone: req.Phone,
	})
	if err != nil {
		return nil, ecode.ErrQuery
	}
	//time.Unix(resp.Member.CreatedAt,0)
	return &model.GetMemberResp{Member: model.Member{
		ID:        resp.Member.Id,
		Phone:     resp.Member.Phone,
		Name:      resp.Member.Name,
		Age:       resp.Member.Age,
		Address:   resp.Member.Address,
		CreatedAt: resp.Member.CreatedAt,
		UpdatedAt: resp.Member.UpdatedAt,
	}}, nil
}

func (l *MemberLogic) QueryMemberByName(req model.QueryMemberByNameReq) (*model.QueryMemberByNameResq, error) {
	resp, err := l.svcCtx.MemberCli.QueryMemberByName(l.ctx, &memberProto.QueryMemberByNameReq{
		Name: req.Name,
	})
	if err != nil {
		return nil, ecode.ErrQuery
	}
	members := make([]model.Member, len(resp.List))
	for _, v := range resp.List {
		members = append(members, model.Member{
			ID:        v.Id,
			Phone:     v.Phone,
			Name:      v.Name,
			Age:       v.Age,
			Address:   v.Address,
			CreatedAt: v.CreatedAt,
			UpdatedAt: v.UpdatedAt,
		})
	}
	return &model.QueryMemberByNameResq{
		Members: members,
	}, nil
}

func (l *MemberLogic) QueryMemberByIds(req model.QueryMemberByIdsReq) (*model.QueryMemberByIdsResp, error) {
	ids, err := stringx.SplitInts(req.IDs)
	if err != nil {
		return nil, ecode.ErrQuery
	}
	resp, err := l.svcCtx.MemberCli.QueryMemberByIDs(l.ctx, &memberProto.QueryMemberByIDsReq{
		Ids: ids,
	})
	fmt.Println(resp)
	if err != nil {
		return nil, ecode.ErrQuery
	}

	list := make(map[int64]*model.Member)
	for k, v := range resp.List {
		list[k] = &model.Member{
			ID:        v.Id,
			Phone:     v.Phone,
			Name:      v.Name,
			Age:       v.Age,
			Address:   v.Address,
			CreatedAt: v.CreatedAt,
			UpdatedAt: v.UpdatedAt,
		}
	}

	return &model.QueryMemberByIdsResp{
		List: list,
	}, nil
}

func (l *MemberLogic) AddMember(req model.AddMemberReq) (*model.AddMemberResp, error) {
	resp, err := l.svcCtx.MemberCli.AddMember(l.ctx, &memberProto.AddMemberReq{
		Phone:   req.Phone,
		Name:    req.Name,
		Age:     req.Age,
		Address: req.Address,
	})
	if err != nil {
		return nil, ecode.ErrInsert
	}
	return &model.AddMemberResp{
		ID: resp.Id,
	}, nil
}

func (l *MemberLogic) BatchAddMember(req model.BatchAddMemberReq) error {
	var args []*memberProto.AddMemberReq
	for _, v := range req.Args {
		args = append(args, &memberProto.AddMemberReq{
			Phone:   v.Phone,
			Name:    v.Name,
			Age:     v.Age,
			Address: v.Address,
		})
	}

	_, err := l.svcCtx.MemberCli.BatchAddMember(l.ctx, &memberProto.BatchAddMemberReq{
		AddMemberReq: args,
	})
	if err != nil {
		return ecode.ErrInsert
	}
	return nil
}

func (l *MemberLogic) InitMember(req model.InitMemberReq) error {
	_, err := l.svcCtx.MemberCli.InitMember(l.ctx, &memberProto.InitMemberReq{
		Phone:   req.Phone,
		Name:    req.Name,
		Age:     req.Age,
		Address: req.Address,
	})
	if err != nil {
		return ecode.ErrInsert
	}
	return nil
}

func (l *MemberLogic) UpdateMember(req model.UpdateMemberReq) error {
	_, err := l.svcCtx.MemberCli.UpdateMember(l.ctx, &memberProto.UpdateMemberReq{
		Id:      req.ID,
		Phone:   req.Phone,
		Name:    req.Name,
		Age:     req.Age,
		Address: req.Address,
	})
	if err != nil {
		return ecode.ErrUpdate
	}
	return nil
}

func (l *MemberLogic) UpdateSomeMember(req model.UpdateSomeMemberReq) error {
	_, err := l.svcCtx.MemberCli.UpdateSomeMember(l.ctx, &memberProto.UpdateSomeMemberReq{
		Id:      req.ID,
		Phone:   req.Phone,
		Name:    req.Name,
		Age:     req.Age,
		Address: req.Address,
	})
	if err != nil {
		return ecode.ErrUpdate
	}
	return nil
}

func (l *MemberLogic) SetMember(req model.SetMemberReq) error {
	_, err := l.svcCtx.MemberCli.SetMember(l.ctx, &memberProto.SetMemberReq{
		Id:      req.ID,
		Phone:   req.Phone,
		Name:    req.Name,
		Age:     req.Age,
		Address: req.Address,
	})
	if err != nil {
		return ecode.ErrUpdate
	}
	return nil
}

func (l *MemberLogic) SortMember(req model.SortMemberReq) error {
	args := make([]*memberProto.SortMember,0)
	for _,v := range req.Args{
		args = append(args,&memberProto.SortMember{
			Id:       v.ID,
			OrderNum: v.OrderNum,
		})
	}
	_, err := l.svcCtx.MemberCli.SortMember(l.ctx, &memberProto.SortMemberReq{
		SortMember:args,
	})
	if err != nil {
		return ecode.ErrUpdate
	}
	return nil
}

func (l *MemberLogic) DelMember(req model.DelMemberReq) error {
	_, err := l.svcCtx.MemberCli.DelMember(l.ctx, &memberProto.DeleteMemberReq{
		Id: req.ID,
	})
	if err != nil {
		return ecode.ErrDelete
	}
	return nil
}

func (l *MemberLogic) AddFavorite(req model.AddFavoriteReq) (*model.AddFavoriteResp, error) {
	resp, err := l.svcCtx.MemberCli.AddFavorite(l.ctx, &memberProto.AddFavoriteReq{
		Mid:  req.Mid,
		Name: req.Name,
	})
	if err != nil {
		return nil, ecode.ErrInsert
	}
	return &model.AddFavoriteResp{
		ID: resp.Id,
	}, nil
}

func (l *MemberLogic) ErrorTest() error {
	_, err := l.svcCtx.MemberCli.ErrorTest(l.ctx, &memberProto.GetMemberByIDReq{})
	return errors.Wrapf(err, "ErrorTest")
}
