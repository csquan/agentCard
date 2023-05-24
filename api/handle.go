package api

import (
	"github.com/ethereum/agentCard/types"
	"github.com/gin-gonic/gin"
	"github.com/go-xorm/xorm"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"github.com/tidwall/gjson"
	"net/http"
)

const ADDRLEN = 42

const Ok = 0

func checkAddr(addr string) error {
	if addr[:2] != "0x" {
		return errors.New("addr must start with 0x")
	}
	if len(addr) != ADDRLEN {
		return errors.New("addr len wrong ,must 40")
	}
	return nil
}

// 录入代理商信息
func (a *ApiService) initAgent(c *gin.Context) {
	buf := make([]byte, 1024)
	n, _ := c.Request.Body.Read(buf)
	data1 := string(buf[0:n])
	res := types.HttpRes{}

	isValid := gjson.Valid(data1)
	if isValid == false {
		logrus.Error("Not valid json")
		res.Code = http.StatusBadRequest
		res.Message = "Not valid json"
		c.SecureJSON(http.StatusBadRequest, res)
		return
	}

	agentName := gjson.Get(data1, "agentName")
	mobileNo := gjson.Get(data1, "mobileNo")
	certType := gjson.Get(data1, "certType")
	certNo := gjson.Get(data1, "certNo")
	name := gjson.Get(data1, "name")
	mailbox := gjson.Get(data1, "mailbox")
	addr := gjson.Get(data1, "addr")
	certPhoto := gjson.Get(data1, "certPhoto")
	certOther := gjson.Get(data1, "certOther")
	password := "12345" //默认密码

	AgentRecord := types.AgentRecord{
		AgentName: agentName.String(),
		MobileNo:  mobileNo.String(),
		CertType:  certType.String(),
		CertNo:    certNo.String(),
		Name:      name.String(),
		Mailbox:   mailbox.String(),
		Addr:      addr.String(),
		CertPhoto: certPhoto.String(),
		CertOther: certOther.String(),
		Password:  password,
	}

	err = a.db.CommitWithSession(a.db, func(s *xorm.Session) error {
		if err := a.db.InsertAgent(s, &AgentRecord); err != nil {
			logrus.Errorf("insert AgentRecord transaction task error:%v tasks:[%v]", err, AgentRecord)
			return err
		}
		return nil
	})

	res.Code = Ok
	res.Message = "success"
	res.Data = ""
	c.SecureJSON(http.StatusOK, res)
}

// 导入卡号
func (a *ApiService) initCard(c *gin.Context) {
	buf := make([]byte, 1024)
	n, _ := c.Request.Body.Read(buf)
	data1 := string(buf[0:n])
	res := types.HttpRes{}

	isValid := gjson.Valid(data1)
	if isValid == false {
		logrus.Error("Not valid json")
		res.Code = http.StatusBadRequest
		res.Message = "Not valid json"
		c.SecureJSON(http.StatusBadRequest, res)
		return
	}

	agentName := gjson.Get(data1, "agentName")
	cardNo := gjson.Get(data1, "cardNo")

	//首先根据卡号 查询得到对应的地址addr
	addr := "xxx"

	AgentCardInfo := types.AgentCardInfo{
		AgentName:  agentName.String(),
		CardNo:     cardNo.String(),
		CardRebate: false,
		Address:    addr,
	}

	err = a.db.CommitWithSession(a.db, func(s *gorm.Session) error {
		if err := a.db.InsertCard(s, &AgentCardInfo); err != nil {
			logrus.Errorf("insert AgentCardInfo transaction task error:%v tasks:[%v]", err, AgentCardInfo)
			return err
		}
		return nil
	})

	res.Code = Ok
	res.Message = "success"
	res.Data = ""
	c.SecureJSON(http.StatusOK, res)
}

func (a *ApiService) updatePassword(c *gin.Context) {
	buf := make([]byte, 1024)
	n, _ := c.Request.Body.Read(buf)
	data1 := string(buf[0:n])
	res := types.HttpRes{}

	isValid := gjson.Valid(data1)
	if isValid == false {
		logrus.Error("Not valid json")
		res.Code = http.StatusBadRequest
		res.Message = "Not valid json"
		c.SecureJSON(http.StatusBadRequest, res)
		return
	}

	agentName := gjson.Get(data1, "agentName")
	password := gjson.Get(data1, "password")

	err = a.db.CommitWithSession(a.db, func(s *gorm.Session) error {
		if err := a.db.UpdateCard(s, agentName.String(), password.String()); err != nil {
			logrus.Errorf("update password error:%v name:%v name:%v", err, agentName, password)
			return err
		}
		return nil
	})

	res.Code = Ok
	res.Message = "success"
	res.Data = ""
	c.SecureJSON(http.StatusOK, res)
}

func (a *ApiService) setProfitRule(c *gin.Context) {
	buf := make([]byte, 1024)
	n, _ := c.Request.Body.Read(buf)
	data1 := string(buf[0:n])
	res := types.HttpRes{}

	isValid := gjson.Valid(data1)
	if isValid == false {
		logrus.Error("Not valid json")
		res.Code = http.StatusBadRequest
		res.Message = "Not valid json"
		c.SecureJSON(http.StatusBadRequest, res)
		return
	}

	agentName := gjson.Get(data1, "agentName")
	settlementRule := gjson.Get(data1, "settlementRule")
	activeRebate := gjson.Get(data1, "activeRebate")
	withdrawThreshold := gjson.Get(data1, "withdrawThreshold")
	grade1proportion := gjson.Get(data1, "grade1proportion")
	grade2proportion := gjson.Get(data1, "grade2proportion")
	grade3proportion := gjson.Get(data1, "grade3proportion")

	AgentProfitInfo := types.AgentProfitInfo{
		AgentName:         agentName.String(),
		SettlementRule:    settlementRule.String(),
		ActiveRebate:      activeRebate.String(),
		WithdrawThreshold: withdrawThreshold.String(),
		Grade1proportion:  grade1proportion.String(),
		Grade2proportion:  grade2proportion.String(),
		Grade3proportion:  grade3proportion.String(),
	}

	err = a.db.CommitWithSession(a.db, func(s *gorm.Session) error {
		if err := a.db.InsertProfitRule(s, &AgentProfitInfo); err != nil {
			logrus.Errorf("insert AgentProfitInfo transaction task error:%v tasks:[%v]", err, AgentProfitInfo)
			return err
		}
		return nil
	})

	res.Code = Ok
	res.Message = "success"
	res.Data = ""
	c.SecureJSON(http.StatusOK, res)
}

// 批量返佣--首先看传入的卡是否已经佣金结算，若没有佣金结算，更新卡的返佣状态
func (a *ApiService) cardsRebate(c *gin.Context) {
	addr := c.Param("contractAddr")
	res := types.HttpRes{}

	err := checkAddr(addr)
	if err != nil {
		logrus.Error(err)
		res.Code = http.StatusBadRequest
		res.Message = err.Error()
		c.SecureJSON(http.StatusBadRequest, res)
		return
	}
	res.Code = Ok
	res.Message = "success"
	res.Data = ""
	c.SecureJSON(http.StatusOK, res)
}

func (a *ApiService) withdraw(c *gin.Context) {
	addr := c.Param("contractAddr")
	res := types.HttpRes{}

	err := checkAddr(addr)
	if err != nil {
		logrus.Error(err)
		res.Code = http.StatusBadRequest
		res.Message = err.Error()
		c.SecureJSON(http.StatusBadRequest, res)
		return
	}
	res.Code = Ok
	res.Message = "success"
	res.Data = ""
	c.SecureJSON(http.StatusOK, res)
}

func (a *ApiService) activeCardInfo(c *gin.Context) {
	addr := c.Param("contractAddr")
	res := types.HttpRes{}

	err := checkAddr(addr)
	if err != nil {
		logrus.Error(err)
		res.Code = http.StatusBadRequest
		res.Message = err.Error()
		c.SecureJSON(http.StatusBadRequest, res)
		return
	}
	res.Code = Ok
	res.Message = "success"
	res.Data = ""
	c.SecureJSON(http.StatusOK, res)
}

func (a *ApiService) agentCardInfo(c *gin.Context) {
	addr := c.Param("contractAddr")
	res := types.HttpRes{}

	err := checkAddr(addr)
	if err != nil {
		logrus.Error(err)
		res.Code = http.StatusBadRequest
		res.Message = err.Error()
		c.SecureJSON(http.StatusBadRequest, res)
		return
	}
	res.Code = Ok
	res.Message = "success"
	res.Data = ""
	c.SecureJSON(http.StatusOK, res)
}
func (a *ApiService) profitHistory(c *gin.Context) {
	addr := c.Param("contractAddr")
	res := types.HttpRes{}

	err := checkAddr(addr)
	if err != nil {
		logrus.Error(err)
		res.Code = http.StatusBadRequest
		res.Message = err.Error()
		c.SecureJSON(http.StatusBadRequest, res)
		return
	}
	res.Code = Ok
	res.Message = "success"
	res.Data = ""
	c.SecureJSON(http.StatusOK, res)
}
func (a *ApiService) GetUnRebate(c *gin.Context) {
	addr := c.Param("contractAddr")
	res := types.HttpRes{}

	err := checkAddr(addr)
	if err != nil {
		logrus.Error(err)
		res.Code = http.StatusBadRequest
		res.Message = err.Error()
		c.SecureJSON(http.StatusBadRequest, res)
		return
	}
	res.Code = Ok
	res.Message = "success"
	res.Data = ""
	c.SecureJSON(http.StatusOK, res)
}
