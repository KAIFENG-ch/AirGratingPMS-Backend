package logic

import (
	"context"

	"air-grating-pms-backend/rpc/workshop/internal/svc"
	"air-grating-pms-backend/rpc/workshop/types"
	"air-grating-pms-backend/rpc/workshop/utils"

	"github.com/zeromicro/go-zero/core/logx"
)

type FindListByEnterpriseLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewFindListByEnterpriseLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FindListByEnterpriseLogic {
	return &FindListByEnterpriseLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *FindListByEnterpriseLogic) FindListByEnterprise(in *types.FindListByEnterpriseReq) (*types.WorkshopList, error) {
	list, err := l.svcCtx.WorkshopModel.FindListByEnterprise(l.ctx, in.GetEnterpriseId(), (in.GetPageNumber()-1)*in.GetPageSize(), in.GetPageNumber())
	if err != nil {
		return nil, err
	}

	out, err := utils.Convert(&list)
	return out, err
}
