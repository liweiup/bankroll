package api

import (
	"bankroll/global"
	"bankroll/service/api/requestParam"
	"bankroll/service/common/response"
	"bankroll/utils"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"go.uber.org/zap"
)

type BankrollApi struct {
}

// @Tags rollback
// @Summary 查询板块交易额
// @accept application/json
// @Produce application/json
// @Param data body requestParam.BankrollParam true "开始时间 结束时间 比对几天 周期"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"succ"}"
// @Router /api/getPlateBankrollData [post]
func (bk *BankrollApi) GetPlateBankrollData(c *gin.Context) {
	var backrollparam requestParam.BankrollParam
	_ = c.ShouldBindBodyWith(&backrollparam,binding.JSON)
	if err := utils.Verify(backrollparam, GetPlateBankrollData); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if data,err := DataInfo.GetPlateBankroll(backrollparam); err != nil {
		global.Zlog.Error("获取失败!", zap.Any("err", err))
		response.FailWithMessage("获取失败", c)
	} else {
		response.OkWithDetailed(data, "获取成功", c)
	}
}