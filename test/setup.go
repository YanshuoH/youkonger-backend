package test

import (
	"github.com/YanshuoH/youkonger/conf"
	"github.com/YanshuoH/youkonger/dao"
	"github.com/YanshuoH/youkonger/models"
	"github.com/gin-gonic/gin"
	"net/http/httptest"
	"net/http"
	"io/ioutil"
	"encoding/json"
	"github.com/YanshuoH/youkonger/utils"
	"path"
	"runtime"
	"strings"
)

func Setup() {
	// pwd of test/setup.go
	_, thisFile, _, _ := runtime.Caller(0)
	// load conf
	c, err := conf.Setup(path.Join(thisFile, "../../conf/conf_test.toml"))
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
		DropTableIfExists(&models.EventUnavailable{}).
		DropTableIfExists(&models.ParticipantUser{})

	// migration tables
	dao.AutoMigration()
	//dao.Conn.LogMode(true)
}

func Teardown() {
	dao.Conn.Close()
}

// given url, http handler and http method, should return the response (recorder) for assertion
func PerformRequest(method, url string, engine *gin.Engine, datas ...string) *httptest.ResponseRecorder {
	var req *http.Request
	if len(datas) > 0 {
		req, _ = http.NewRequest(method, url, strings.NewReader(datas[0]))
	} else {
		req, _ = http.NewRequest(method, url, nil)
	}
	if method == http.MethodPost || method == http.MethodPut {
		req.Header.Set("Content-Type", "application/json")
	}

	w := httptest.NewRecorder()
	engine.ServeHTTP(w, req)
	return w
}

// read uniformed json response
func ReadJsonResponse(w *httptest.ResponseRecorder) utils.JSONResponse {
	byt, err := ioutil.ReadAll(w.Body)
	if err != nil {
		panic(err)
	}

	jsResp := utils.JSONResponse{}
	err = json.Unmarshal(byt, &jsResp)
	if err != nil {
		panic(err)
	}

	return jsResp
}
