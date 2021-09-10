package factory_method_patterns

import (
	"errors"
	"math/rand"
)

/*
	工厂模式：
		对于执行某项事务，输入的种类可能是多种多样的，但对于各种事务其抽象出的处理方式最终是一样的。
		类似于一个垃圾处理场，将输入的各种垃圾进行处理，最终都是让垃圾无害化。
 */

//本例子实现了一个虚拟货币的结算方法，各种不同的货币最后都按汇率转化为现实货币

//设置币种
type cash int
const (
	BTC cash = 1<<iota
	ETH
	FIL
)

//对于各种输入的统一事务
type Withdraw interface {
	WithdrawToRMB() (float64,error)
	//这里可以添加其他事务
}

//定义各种输入
type BTCCoin struct {
	Quantity float64
	Rate float64
}

type ETHCoin struct {
	Quantity float64
	Rate float64
}

type FILCoin struct {
	Quantity float64
	Rate float64
}


//实现统一事务的每种输入（连接这种事物与各种输入）
func (wdBTC *BTCCoin)WithdrawToRMB() (float64,error){
	if wdBTC.Quantity<=0{return -1,errors.New("illegal input")}
	return wdBTC.Rate*wdBTC.Quantity,nil
}

func (wdETC *ETHCoin)WithdrawToRMB() (float64,error){
	if wdETC.Quantity<=0{return -1,errors.New("illegal input")}
	return wdETC.Rate*wdETC.Quantity,nil
}

func (wdFIL *FILCoin)WithdrawToRMB() (float64,error){
	if wdFIL.Quantity<=0{return -1,errors.New("illegal input")}
	return wdFIL.Rate*wdFIL.Quantity,nil
}


//实现工厂方法，及这种方法可以实现各种输入（情况）下的某一种或几种相同的事务（处理）
func GenerateWithdraw(c cash, quantity float64)(Withdraw, error){
	switch c {
	case BTC:
		cash := new(BTCCoin)
		cash.Quantity = quantity
		cash.Rate = generateRate()
		return cash, nil
	case ETH:
		cash := new(ETHCoin)
		cash.Quantity = quantity
		cash.Rate = generateRate()
		return cash, nil
	case FIL:
		cash := new(FILCoin)
		cash.Quantity = quantity
		cash.Rate = generateRate()
		return cash, nil
	default:
		return nil, errors.New("no such cash")
	}
}

//生成汇率（娱乐用）
func generateRate()float64{
	return rand.Float64()+rand.Float64()
}
