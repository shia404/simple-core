package position

import (
	"net/http"

	"github.com/suyuan32/simple-admin-core/service/admin/api/internal/logic/position"
	"github.com/suyuan32/simple-admin-core/service/admin/api/internal/svc"
	"github.com/suyuan32/simple-admin-core/service/admin/api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

// swagger:route post /position position GetPositionById
//
// Get position by ID | 通过ID获取职位
//
// Get position by ID | 通过ID获取职位
//
// Parameters:
//  + name: body
//    require: true
//    in: body
//    type: IDReq
//
// Responses:
//  200: PositionInfoResp

func GetPositionByIdHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.IDReq
		if err := httpx.Parse(r, &req, true); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := position.NewGetPositionByIdLogic(r.Context(), svcCtx)
		resp, err := l.GetPositionById(&req)
		if err != nil {
			err = svcCtx.Trans.TransError(r.Context(), err)
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
