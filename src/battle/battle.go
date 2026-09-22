package battle

type AI struct{
	name string
	atk int
	pv int
	pvmax int
	coinReward int
}

func (e *AI) initEnemy(n string, p int, a int) {

}

func setUpEnemies() {
	
}

func battle() {
	var dummy AI
	dummy.initEnemy("dummy", 100, 5)

}