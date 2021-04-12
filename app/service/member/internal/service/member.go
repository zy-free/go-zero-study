package service

import (
	"context"
	"go-zero-study/app/service/member/internal/model"

	"go-zero-study/core/logx"
)

type MemberLogic struct {
	ctx context.Context
	svc *Service
	logx.Logger
}

func NewMemberLogic(ctx context.Context, svc *Service) *MemberLogic {
	return &MemberLogic{
		ctx:    ctx,
		svc:    svc,
		Logger: logx.WithContext(ctx),
	}
}

func (l *MemberLogic) GetMemberByID(id int64) (member *model.Member, err error) {
	return l.svc.MemberDao.FindOne(id)
}

func (l *MemberLogic) GetMemberByPhone(phone string) (member *model.Member, err error) {
	return l.svc.MemberDao.FindOneByPhone(phone)
}

func (l *MemberLogic) GetMemberMaxAge() (age int64, err error) {
	return l.svc.MemberDao.GetMemberMaxAge()
}

func (l *MemberLogic) QueryMemberByName(name string) (members []*model.Member, err error) {
	return l.svc.MemberDao.QueryByName(name)
}

func (l *MemberLogic) QueryMemberByIDs(ids []int64) (res map[int64]*model.Member, err error) {
	return l.svc.MemberDao.QueryMemberByIDs(ids)
}

func (l *MemberLogic) AddMember(member model.Member) (id int64, err error) {
	return l.svc.MemberDao.Insert(member)
}


func (l *MemberLogic) BatchAddMember(members []model.Member) (affectRow int64, err error) {
	return l.svc.MemberDao.BatchAddMember(members)
}

func (l *MemberLogic) InitMember(member model.Member) (err error) {
	return l.svc.MemberDao.Init(member)
}

func (l *MemberLogic) UpdateMember(member model.Member) (err error) {
	return l.svc.MemberDao.Update(member)
}

func (l *MemberLogic) UpdateSome(member model.Member) (err error) {
	return l.svc.MemberDao.UpdateSome(member)
}

func (l *MemberLogic) SetMember(member model.Member) (err error) {
	return l.svc.MemberDao.Set(member)
}

func (l *MemberLogic) SortMember(args model.ArgMemberSort) (err error) {
	return l.svc.MemberDao.SortMember(args)
}

func (l *MemberLogic) DelMember(id int64) (err error) {
	return l.svc.MemberDao.Delete(id)
}
