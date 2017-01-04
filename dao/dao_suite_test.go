package dao_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
	"github.com/YanshuoH/youkonger/conf"
	"github.com/YanshuoH/youkonger/dao"
)

func TestDao(t *testing.T) {
	RegisterFailHandler(Fail)
	// load conf
	c, err := conf.Setup("../conf/conf_test.gcfg")
	if err != nil {
		panic(err)
	}
	// connect mysql test db
	dao.Connect(c.DbConf.Dsn)
	if err != nil {
		panic(err)
	}
	// migration tables
	dao.AutoMigration()

	RunSpecs(t, "Dao Suite")
}
