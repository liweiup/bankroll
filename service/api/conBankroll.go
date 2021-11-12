package api

import (
	"bankroll/global"
	"bankroll/service/common/response"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type BankrollApi struct {
}

// @Summary 创建基础api
// @accept application/json
// @Produce application/json
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /api/getPlateBankrollData [post]
func (bk *BankrollApi) GetPlateBankrollData(c *gin.Context) {
	if data,err := bankrollModel.GetPlateBankroll("2021-11-05","2021-11-05"); err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Any("err", err))
		response.FailWithMessage("获取失败", c)
	} else {
		response.OkWithDetailed(data, "获取成功", c)
	}
}