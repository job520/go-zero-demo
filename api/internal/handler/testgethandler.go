package handler

import (
	"net/http"

	"demo/api/internal/logic"
	"demo/api/internal/svc"
	"demo/api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func testGetHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.Req
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := logic.NewTestGetLogic(r.Context(), svcCtx)
		resp, err := l.TestGet(&req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
