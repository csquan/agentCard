package api

import (
	"fmt"
	"github.com/ethereum/agentCard/config"
	"github.com/ethereum/agentCard/types"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"net/http"
)

type ApiService struct {
	db     types.IDB
	config *config.Config
}

func NewApiService(db types.IDB, cfg *config.Config) *ApiService {
	return &ApiService{
		db:     db,
		config: cfg,
	}
}

func (a *ApiService) Run() {
	r := gin.Default()

	corsConfig := cors.DefaultConfig()
	corsConfig.AllowCredentials = true
	corsConfig.AllowOrigins = []string{"*"}
	r.Use(func(ctx *gin.Context) {
		method := ctx.Request.Method
		ctx.Header("Access-Control-Allow-Origin", "*")
		ctx.Header("Access-Control-Allow-Headers", "*")
		// ctx.Header("Access-Control-Allow-Headers", "Content-Type,addr,GoogleAuth,AccessToken,X-CSRF-Token,Authorization,Token,token,auth,x-token")
		ctx.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS")
		ctx.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Content-Type")
		ctx.Header("Access-Control-Allow-Credentials", "true")
		if method == "OPTIONS" {
			ctx.AbortWithStatus(http.StatusNoContent)
			return
		}
		ctx.Next()
	})

	//验证token
	//r.Use(auth.MustExtractUser())

	//录入代理商注册信息
	r.POST("/initAgent", a.initAgent)
	//导入卡号
	r.POST("/initCard", a.initCard)
	//更新密码
	r.POST("/updatePassword", a.updatePassword)

	//设置收益结算规则
	r.POST("/setProfitRule", a.setProfitRule)

	//批量返佣
	r.POST("/cardsRebate", a.cardsRebate)

	//提现
	r.POST("/withdraw", a.withdraw)

	//卡激活信息查询
	r.Get("/activeCardInfo", a.activeCardInfo)

	//代理卡信息查询
	r.Get("/agentCardInfo", a.agentCardInfo)

	//收益发放历史
	r.Get("/profitHistory", a.profitHistory)

	//未返佣查的卡查询
	r.Get("/GetUnRebate", a.GetUnRebate)

	logrus.Info("agentCard run at " + a.config.Server.Port)

	err := r.Run(fmt.Sprintf(":%s", a.config.Server.Port))
	if err != nil {
		logrus.Fatalf("start http server err:%v", err)
	}
}
