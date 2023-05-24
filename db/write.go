package db

import (
	"fmt"
	"github.com/ethereum/agentCard/types"
	"time"

	"github.com/ethereum/agentCard/config"
	"github.com/go-gorm/gorm"
	_ "github.com/go-sql-driver/mysql"
	"github.com/sirupsen/logrus"
	"xorm.io/core"
)

type Mysql struct {
	conf   *config.Config
	engine *xorm.Engine
}

func NewMysql(cfg *config.Config) (m *Mysql, err error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8", cfg.Db.Name, cfg.Db.Password, cfg.Db.Ip, cfg.Db.Port, cfg.Db.Database)
	engine, err := xorm.NewEngine("mysql", dsn)
	if err != nil {
		logrus.Errorf("create engine error: %v", err)
		return
	}
	engine.ShowSQL(false)
	engine.Logger().SetLevel(core.LOG_DEBUG)
	location, err := time.LoadLocation("UTC")
	if err != nil {
		return nil, err
	}
	engine.SetTZLocation(location)
	engine.SetTZDatabase(location)

	m = &Mysql{
		conf:   cfg,
		engine: engine,
	}

	return
}

func (m *Mysql) GetEngine() *gorm.Engine {
	return m.engine
}

func (m *Mysql) GetSession() *gorm.Session {
	return m.engine.NewSession()
}

func (m *Mysql) CommitWithSession(db types.IDB, executeFunc func(*gorm.Session) error) (err error) {
	session := db.GetSession()
	err = session.Begin()
	if err != nil {
		logrus.Errorf("begin session error:%v", err)
		return
	}

	defer session.Close()

	err = executeFunc(session)
	if err != nil {
		logrus.Errorf("execute func error:%v", err)
		err1 := session.Rollback()
		if err1 != nil {
			logrus.Errorf("session rollback error:%v", err1)
		}
		return
	}

	err = session.Commit()
	if err != nil {
		logrus.Errorf("commit session error:%v", err)
	}

	return
}

func (m *Mysql) InsertAgent(itf gorm.Interface, agent *types.AgentRecord) (err error) {
	_, err = itf.Insert(agent)
	if err != nil {
		logrus.Errorf("insert agent  error:%v, agent:%v", err, agent)
	}
	return
}

//
//func (m *Mysql) UpdateTransactionTask(itf xorm.Interface, task *types.TransactionTask) error {
//	_, err := itf.Table("t_transaction_task").Where("f_id = ?", task.ID).Update(task)
//	return err
//}
//func (m *Mysql) UpdateTransactionTaskMessage(taskID uint64, message string) error {
//	_, err := m.engine.Exec("update t_transaction_task set f_message = ? where f_id = ?", message, taskID)
//	return err
//}
