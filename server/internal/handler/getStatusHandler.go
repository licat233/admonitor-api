package handler

import (
	"net/http"

	"Admonitor-api/server/internal/logic"
	"Admonitor-api/server/internal/svc"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func GetStatusHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := logic.NewGetStatusLogic(r.Context(), svcCtx)
		resp, err := l.GetStatus()
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
