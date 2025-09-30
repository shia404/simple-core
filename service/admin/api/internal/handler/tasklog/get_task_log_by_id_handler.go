package tasklog

import (
	"net/http"

	"github.com/suyuan32/simple-admin-core/service/admin/api/internal/logic/tasklog"
	"github.com/suyuan32/simple-admin-core/service/admin/api/internal/svc"
	"github.com/suyuan32/simple-admin-core/service/admin/api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

// swagger:route post /task_log tasklog GetTaskLogById
//
// Get task log by ID | 通过ID获取任务日志
//
// Get task log by ID | 通过ID获取任务日志
//
// Parameters:
//  + name: body
//    require: true
//    in: body
//    type: IDReq
//
// Responses:
//  200: TaskLogInfoResp

func GetTaskLogByIdHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.IDReq
		if err := httpx.Parse(r, &req, true); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := tasklog.NewGetTaskLogByIdLogic(r.Context(), svcCtx)
		resp, err := l.GetTaskLogById(&req)
		if err != nil {
			err = svcCtx.Trans.TransError(r.Context(), err)
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
