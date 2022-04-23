package handler

import (
	"net/http"

	"Admonitor-api/server/internal/logic"
	"Admonitor-api/server/internal/svc"
	"Admonitor-api/server/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func GetStatisticsHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.GetStatisticsRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := logic.NewGetStatisticsLogic(r.Context(), svcCtx)
		resp, err := l.GetStatistics(&req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
