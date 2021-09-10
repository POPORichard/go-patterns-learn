package factory_method_patterns

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)
//工厂功能的测试
func TestGenerateWithdraw(t *testing.T) {
	c,err := GenerateWithdraw(FIL,100)
	assert.IsType(t, &FILCoin{}, c,"Error wrong type")
	assert.NoError(t, err, "Should not have error")

	c,err = GenerateWithdraw(-1,100)
	assert.IsType(t, nil, c,"Error wrong type")
	assert.Error(t, err, "Error should be : no such cash")
}

//实际业务的测试
func TestBTCCoin_WithdrawToRMB(t *testing.T) {
	//输入币种和金额，返回业务实体
	c, err := GenerateWithdraw(BTC,123.123)
	assert.NoError(t, err, "cash type wrong")
	//实体中提供了提现方法
	RMB,err := c.WithdrawToRMB()
	assert.NoError(t, err, "cash input illegal")
	fmt.Println(RMB)
}