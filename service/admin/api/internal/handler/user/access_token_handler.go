package user

import (
	"net/http"

	"github.com/suyuan32/simple-admin-core/service/admin/api/internal/logic/user"
	"github.com/suyuan32/simple-admin-core/service/admin/api/internal/svc"
	"github.com/zeromicro/go-zero/rest/httpx"
)

// swagger:route get /user/access_token user AccessToken
//
// Access token | 获取短期 token
//
// Access token | 获取短期 token
//
// Responses:
//  200: RefreshTokenResp

func AccessTokenHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := user.NewAccessTokenLogic(r.Context(), svcCtx)
		resp, err := l.AccessToken()
		if err != nil {
			err = svcCtx.Trans.TransError(r.Context(), err)
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
