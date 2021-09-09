package builder_patterns

/*
	建造者模式：
		将复杂的对象通过简单的每一步构造出来
*/

type Tank struct {
	Guns       int     //火炮
	MachineGun int     //机枪
	AirDefense int     //防空
	Power      string  //动力
	Armor      float64 //装甲厚度
}

//对于我们的坦克，我们要进行参数的设置
//先写出接口，以明确有那些参数
type Builder interface {
	SetGuns() Builder
	SetMachineGun() Builder
	SetAirDefense() Builder
	SetPower() Builder
	SetArmor() Builder
	MakeTank() Tank //建造坦克，最后送出的是辆坦克，而不是我们的工程师
}

//设计我们的超级坦克，当然要符合坦克的基本特征，所以会从坦克类中“继承”
type SuperTank struct {
	Tank Tank
}

//下面实现我们接口中的功能
func (superTank *SuperTank) SetGuns() Builder {
	superTank.Tank.Guns = 2
	return superTank
}

func (superTank *SuperTank) SetMachineGun() Builder {
	superTank.Tank.MachineGun = 10
	return superTank
}

func (superTank *SuperTank) SetAirDefense() Builder {
	superTank.Tank.AirDefense = 20
	return superTank
}

func (superTank *SuperTank) SetPower() Builder {
	superTank.Tank.Power = "Nuclear"
	return superTank
}

func (superTank *SuperTank) SetArmor() Builder {
	superTank.Tank.Armor = 10000.66
	return superTank
}

func (superTank *SuperTank) MakeTank() Tank {
	return superTank.Tank
}

//调用工程师的命令
type Order struct{
	builder Builder
}

//构造能完成我们坦克的工程师
func (order *Order)Construct(){
	order.builder.SetGuns().SetMachineGun().SetAirDefense().SetPower().SetArmor()
}

//安置工程师
func (order *Order)SetBuilder(builder Builder){
	order.builder = builder
}
