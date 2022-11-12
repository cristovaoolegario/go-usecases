package usecase

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

func TestSuite(t *testing.T) {
	suite.Run(t, new(CalculatePriceUseCaseTestSuite))
	suite.Run(t, new(GetTotalUseCaseTestSuite))
}
