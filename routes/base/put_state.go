package base

import (
	"github.com/Parker-Yang/def-braveTroops/consts/status"
	"github.com/Parker-Yang/srv-braveTroops/utils"
	"reflect"

	"github.com/Parker-Yang/def-braveTroops/proto"
	"github.com/Parker-Yang/srv-braveTroops/global"
	"github.com/Parker-Yang/srv-braveTroops/modules/base"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func Put(ctx *gin.Context) {
	data := &proto.PutState{}
	err := ctx.BindJSON(data)
	if err != nil {
		utils.FailedResp(ctx, status.BadRequest, nil)
		return
	}
	logrus.Debugln(reflect.TypeOf(data))
	controller := base.NewController(global.Conf.Client)
	state, err := controller.PutState(data)
	if err != nil {
		utils.UnWarpFailedResp(ctx,err)
		return
	}
	utils.SuccessResp(ctx, status.Ok, state)
}
