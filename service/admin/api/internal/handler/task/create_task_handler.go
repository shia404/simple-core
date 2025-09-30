package task

import (
	"net/http"

	"github.com/suyuan32/simple-admin-core/service/admin/api/internal/logic/task"
	"github.com/suyuan32/simple-admin-core/service/admin/api/internal/svc"
	"github.com/suyuan32/simple-admin-core/service/admin/api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

// swagger:route post /task/create task CreateTask
//
// Create task information | 创建Task
//
// Create task information | 创建Task
//
// Parameters:
//  + name: body
//    require: true
//    in: body
//    type: TaskInfo
//
// Responses:
//  200: BaseMsgResp

func CreateTaskHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.TaskInfo
		if err := httpx.Parse(r, &req, true); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := task.NewCreateTaskLogic(r.Context(), svcCtx)
		resp, err := l.CreateTask(&req)
		if err != nil {
			err = svcCtx.Trans.TransError(r.Context(), err)
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
