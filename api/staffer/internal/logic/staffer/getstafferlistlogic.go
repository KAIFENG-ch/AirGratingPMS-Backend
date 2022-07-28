package staffer

import (
	"context"

	"air-grating-pms-backend/api/staffer/internal/svc"
	"air-grating-pms-backend/api/staffer/internal/types"
	"air-grating-pms-backend/api/staffer/utils/convert"
	"air-grating-pms-backend/rpc/staffer/pb"
	"air-grating-pms-backend/utils"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetStafferListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetStafferListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetStafferListLogic {
	return &GetStafferListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetStafferListLogic) GetStafferList(req *types.GetStafferListReq) (resp *types.GetStafferListReply, err error) {
	list, err := l.svcCtx.StafferRPC.FindListByEnterprise(l.ctx, &pb.FindListByEnterpriseReq{
		EnterpriseId: utils.GetEnterpriseId(l.ctx),
		PageSize:     req.PageSize,
		PageNumber:   req.PageNumber,
	})
	if err != nil {
		return nil, err
	}

	listResp, err := convert.ListConvert(list)

	return &types.GetStafferListReply{
		Message:     "OK",
		StafferList: *listResp,
	}, err
}
