package logic

import (
	"context"

	"Admonitor-api/server/internal/svc"
	"Admonitor-api/server/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetStatisticsLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetStatisticsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetStatisticsLogic {
	return &GetStatisticsLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetStatisticsLogic) GetStatistics(req *types.GetStatisticsRequest) (resp *types.Response, err error) {
	// todo: add your logic here and delete this line

	return
}
