package configuration

import (
	"context"

	ent2 "github.com/suyuan32/simple-admin-core/service/admin/rpc/ent"
	configuration2 "github.com/suyuan32/simple-admin-core/service/admin/rpc/ent/configuration"
	"github.com/suyuan32/simple-admin-core/service/admin/rpc/ent/predicate"
	"github.com/suyuan32/simple-admin-core/service/admin/rpc/internal/svc"
	"github.com/suyuan32/simple-admin-core/service/admin/rpc/internal/utils/dberrorhandler"
	"github.com/suyuan32/simple-admin-core/service/admin/rpc/types/core"

	"github.com/suyuan32/simple-admin-common/utils/pointy"
	"github.com/zeromicro/go-zero/core/logx"
)

type GetConfigurationListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetConfigurationListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetConfigurationListLogic {
	return &GetConfigurationListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetConfigurationListLogic) GetConfigurationList(in *core.ConfigurationListReq) (*core.ConfigurationListResp, error) {
	var predicates []predicate.Configuration
	if in.Name != nil {
		predicates = append(predicates, configuration2.NameContains(*in.Name))
	}
	if in.Key != nil {
		predicates = append(predicates, configuration2.KeyContains(*in.Key))
	}
	if in.Category != nil {
		predicates = append(predicates, configuration2.CategoryEQ(*in.Category))
	}

	if in.State != nil {
		predicates = append(predicates, configuration2.StateEQ(*in.State))
	}

	result, err := l.svcCtx.DB.Configuration.Query().Where(predicates...).Page(l.ctx, in.Page, in.PageSize, func(cp *ent2.ConfigurationPager) {
		if in.Category != nil {
			cp.Order = ent2.Asc(configuration2.FieldSort)
		}
	})

	if err != nil {
		return nil, dberrorhandler.DefaultEntError(l.Logger, err, in)
	}

	resp := &core.ConfigurationListResp{}
	resp.Total = result.PageDetails.Total

	for _, v := range result.List {
		resp.Data = append(resp.Data, &core.ConfigurationInfo{
			Id:        &v.ID,
			CreatedAt: pointy.GetPointer(v.CreatedAt.UnixMilli()),
			UpdatedAt: pointy.GetPointer(v.UpdatedAt.UnixMilli()),
			Sort:      &v.Sort,
			State:     &v.State,
			Name:      &v.Name,
			Key:       &v.Key,
			Value:     &v.Value,
			Category:  &v.Category,
			Remark:    &v.Remark,
		})
	}

	return resp, nil
}
