package dictionary

import (
	"net/http"

	"github.com/suyuan32/simple-admin-core/service/admin/api/internal/logic/dictionary"
	"github.com/suyuan32/simple-admin-core/service/admin/api/internal/svc"
	"github.com/suyuan32/simple-admin-core/service/admin/api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

// swagger:route post /dictionary dictionary GetDictionaryById
//
// Get Dictionary by ID | 通过ID获取字典
//
// Get Dictionary by ID | 通过ID获取字典
//
// Parameters:
//  + name: body
//    require: true
//    in: body
//    type: IDReq
//
// Responses:
//  200: DictionaryInfoResp

func GetDictionaryByIdHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.IDReq
		if err := httpx.Parse(r, &req, true); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := dictionary.NewGetDictionaryByIdLogic(r.Context(), svcCtx)
		resp, err := l.GetDictionaryById(&req)
		if err != nil {
			err = svcCtx.Trans.TransError(r.Context(), err)
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
