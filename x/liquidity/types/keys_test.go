package types_test

import (
	"testing"

	"github.com/tendermint/liquidity/x/liquidity/types"

	"github.com/stretchr/testify/suite"
)

type keysTestSuite struct {
	suite.Suite
}

func TestKeysTestSuite(t *testing.T) {
	suite.Run(t, new(keysTestSuite))
}

func (s *keysTestSuite) TestGetLiquidityPoolKey() {
	s.Require().Equal([]byte{0x11, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x9}, types.GetLiquidityPoolKey(9))
	s.Require().Equal([]byte{0x11, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0}, types.GetLiquidityPoolKey(0))
}

func (s *keysTestSuite) TestGetLiquidityPoolByReserveAccIndexKey() {
	s.Require().Equal([]byte{18, 116, 101, 115, 116}, types.GetLiquidityPoolByReserveAccIndexKey([]byte("test")))
	s.Require().Equal([]byte{18}, types.GetLiquidityPoolByReserveAccIndexKey(nil))
}

func (s *keysTestSuite) TestGetLiquidityPoolBatchIndexKey() {
	s.Require().Equal([]byte{0x21, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0xa}, types.GetLiquidityPoolBatchIndexKey(10))
}

func (s *keysTestSuite) TestGetLiquidityPoolBatchKey() {
	s.Require().Equal([]byte{0x22, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0xa, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0xc8}, types.GetLiquidityPoolBatchKey(10, 200))
	s.Require().Equal([]byte{0x22, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0}, types.GetLiquidityPoolBatchKey(0, 0))
}

func (s *keysTestSuite) TestGetLiquidityPoolBatchDepositMsgsPrefix() {
	s.Require().Equal([]byte{0x31, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0xa, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0xc8}, types.GetLiquidityPoolBatchDepositMsgsPrefix(10, 200))
	s.Require().Equal([]byte{0x31, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0}, types.GetLiquidityPoolBatchDepositMsgsPrefix(0, 0))
}

func (s *keysTestSuite) TestGetLiquidityPoolBatchWithdrawMsgsPrefix() {
	s.Require().Equal([]byte{0x32, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0xa, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0xc8}, types.GetLiquidityPoolBatchWithdrawMsgsPrefix(10, 200))
	s.Require().Equal([]byte{0x32, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0}, types.GetLiquidityPoolBatchWithdrawMsgsPrefix(0, 0))
}

func (s *keysTestSuite) TestGetLiquidityPoolBatchSwapMsgsPrefix() {
	s.Require().Equal([]byte{0x33, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0xa, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0xc8}, types.GetLiquidityPoolBatchSwapMsgsPrefix(10, 200))
	s.Require().Equal([]byte{0x33, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0}, types.GetLiquidityPoolBatchSwapMsgsPrefix(0, 0))
}

func (s *keysTestSuite) TestGetLiquidityPoolBatchDepositMsgIndex() {
	s.Require().Equal([]byte{0x31, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0xa, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0xc8, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0xa},
		types.GetLiquidityPoolBatchDepositMsgIndex(10, 200, 10))
	s.Require().Equal([]byte{0x31, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0},
		types.GetLiquidityPoolBatchDepositMsgIndex(0, 0, 0))
}

func (s *keysTestSuite) TestGetLiquidityPoolBatchWithdrawMsgIndex() {
	s.Require().Equal([]byte{0x32, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0xa, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0xc8, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0xa},
		types.GetLiquidityPoolBatchWithdrawMsgIndex(10, 200, 10))
	s.Require().Equal([]byte{0x32, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0},
		types.GetLiquidityPoolBatchWithdrawMsgIndex(0, 0, 0))
}

func (s *keysTestSuite) TestGetLiquidityPoolBatchSwapMsgIndex() {
	s.Require().Equal([]byte{0x33, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0xa, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0xc8, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0xa},
		types.GetLiquidityPoolBatchSwapMsgIndex(10, 200, 10))
	s.Require().Equal([]byte{0x33, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0},
		types.GetLiquidityPoolBatchSwapMsgIndex(0, 0, 0))
}
