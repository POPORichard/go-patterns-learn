package builder_patterns

import (
	"fmt"
	"testing"
)

func TestBuilderPattern(t *testing.T){
	//获得召集工程师的命令
	order := Order{}

	//获得建造图纸
	superTank := &SuperTank{}

	//把图纸给工程师
	order.SetBuilder(superTank)

	//工程师获得建造能力
	order.Construct()

	//建造超级坦克
	tank := order.builder.MakeTank()

	//交付展示
	fmt.Println(tank)
}