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

	//召集工程师把图纸给工程师
	order.SetBuilder(superTank)

	//召集工程师并使获得建造能力
	order.Construct()

	//召集工程师使他们开始建造超级坦克
	tank := order.builder.MakeTank()

	//交付展示
	fmt.Println(tank)
}