package dao_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
	"github.com/YanshuoH/youkonger/conf"
	"github.com/YanshuoH/youkonger/dao"
	"github.com/YanshuoH/youkonger/models"
)

func TestDao(t *testing.T) {
	RegisterFailHandler(Fail)

	// close when specs done
	defer

	RunSpecs(t, "Dao Suite")
}

var _ = BeforeSuite(func() {
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

	// drop tables
	dao.Conn.
		DropTableIfExists(&models.Event{}).
		DropTableIfExists(&models.EventDate{}).
		DropTableIfExists(&models.EventParticipant{}).
		DropTableIfExists(&models.EventUnavailable{})

	// migration tables
	dao.AutoMigration()
	dao.Conn.LogMode(true)
})

var _ = AfterSuite(func() {
	dao.Conn.Close()
})
