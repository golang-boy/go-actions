package integration

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type Suite struct {
	suite.Suite
}

func (s *Suite) SetupSuite() {

	s.T().Log("setup suite")
}

func (s *Suite) TestGet() {

	s.T().Log("test functions")

}

func (s *Suite) TearDownSuite() {
	s.T().Log("teardown suite")
}

func TestMysqlInsert(t *testing.T) {
	suite.Run(t, &Suite{})
}
