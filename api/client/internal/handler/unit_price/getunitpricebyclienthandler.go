package unit_price

import (
	"net/http"

	"air-grating-pms-backend/api/client/internal/logic/unit_price"
	"air-grating-pms-backend/api/client/internal/svc"
	"air-grating-pms-backend/api/client/internal/types"
	"air-grating-pms-backend/common/response"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func GetUnitPriceByClientHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.GetUnitPriceByClientReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := unit_price.NewGetUnitPriceByClientLogic(r.Context(), svcCtx)
		resp, err := l.GetUnitPriceByClient(&req)
		response.Response(w, resp, err)
	}
}
