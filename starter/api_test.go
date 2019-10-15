package starter

import (
	"fmt"
	"testing"
	"time"

	"github.com/IMQS/log"
	"github.com/IMQS/nf/nftest"
	"github.com/IMQS/serviceauth"
	"github.com/IMQS/serviceauth/permissions"
)

const httpPort = 2000

func startServer(t *testing.T) *Service {
	s := NewService()
	s.config.HttpPort = httpPort
	s.config.DB = nftest.MakeDBConfig("starter")
	s.log = log.NewTesting(t)
	serviceauth.ActivateMockToken(1, "user@example.com", []int{permissions.PermEnabled, permissions.PermReportCreator})
	s.Initialize()
	go s.ListenAndServe()
	nftest.PingUntil200(t, time.Second, baseURL()+"/ping")
	return s
}

func baseURL() string {
	return fmt.Sprintf("http://localhost:%v", httpPort)
}

func TestFrogCRUD(t *testing.T) {
	startServer(t)
	nftest.POSTJson(t, baseURL()+"/frog/add", `[{"FrogTypeID": 123, "Description": "bullfrog"},{"FrogTypeID": 666, "Description": "toad"}]`, 200)
	nftest.GETJson(t, baseURL()+"/frog/list", `>>"FrogTypeID":123`)
	nftest.GETJson(t, baseURL()+"/frog/list", `>>"FrogTypeID":666`)
	nftest.GETJson(t, baseURL()+"/frog/list", `>>"Description":"bullfrog"`)
	nftest.GETJson(t, baseURL()+"/frog/list", `>>"Description":"toad"`)
	nftest.GETJson(t, baseURL()+"/frog/list?type=123", `!>>"Description":"toad"`)
	nftest.GETJson(t, baseURL()+"/frog/list?type=123", `>>"Description":"bullfrog"`)
	nftest.GETDump(t, baseURL()+"/frog/list?type=123")
}
