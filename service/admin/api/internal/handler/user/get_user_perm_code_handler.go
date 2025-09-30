package user

import (
	"net/http"

	"github.com/suyuan32/simple-admin-core/service/admin/api/internal/logic/user"
	"github.com/suyuan32/simple-admin-core/service/admin/api/internal/svc"
	"github.com/zeromicro/go-zero/rest/httpx"
)

// swagger:route get /user/perm user GetUserPermCode
//
// Get user's permission code | 获取用户权限码
//
// Get user's permission code | 获取用户权限码
//
// Responses:
//  200: PermCodeResp

func GetUserPermCodeHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := user.NewGetUserPermCodeLogic(r.Context(), svcCtx)
		resp, err := l.GetUserPermCode()
		if err != nil {
			err = svcCtx.Trans.TransError(r.Context(), err)
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
