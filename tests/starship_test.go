package tests

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"testing"
	"testing-in-golang/framework/swapi"
)

type StarshipSuite struct {
	suite.Suite
	shipNumber int
}

func (suite *StarshipSuite) SetupTest() {
	suite.shipNumber = 9
}

func (suite *StarshipSuite) Test_Name() {
	ship := swapi.GetStarship(suite.shipNumber)
	assert.Equal(suite.T(), ship.Name, "Death Star")
}

func (suite *StarshipSuite) Test_Model() {
	ship := swapi.GetStarship(suite.shipNumber)
	assert.Equal(suite.T(), ship.Model, "DS-2 Orbital Battle Station")
}

func Test_GetStarship_200(t *testing.T) {
	t.Parallel()
	suite.Run(t, new(StarshipSuite))
}
