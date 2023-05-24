package types

import (
	"github.com/go-gorm/gorm"
)

//go:generate mockgen -source=$GOFILE -destination=./mock/mock_db.go -package=mock
type IReader interface {
	GetMechanismInfo(name string) (*Mechanism, error)
}

type IWriter interface {
	GetSession() *gorm.Session
	GetEngine() *gorm.Engine
	CommitWithSession(db IDB, executeFunc func(*gorm.Session) error) (err error)

	InsertAgent(itf gorm.Interface, agent *AgentRecord) (err error)
	InsertCard(itf gorm.Interface, card *AgentCardInfo) (err error)
	InsertProfitRule(itf gorm.Interface, rule *AgentProfitInfo) (err error)

	UpdateCard(itf gorm.Interface, agentName string, password string) (err error)

	//InsertTransfer(itf xorm.Interface, transfer *TransferRecord) (err error)
	//InsertWithdraw(itf xorm.Interface, withdraw *Withdraw) (err error)
	//InsertMechanism(itf xorm.Interface, mechanism *Mechanism) (err error)
}

type IDB interface {
	IReader
	IWriter
}
