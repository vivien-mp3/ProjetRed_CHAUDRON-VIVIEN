package battle

import (
	"fmt"
	"projet/src/bubble"
	"projet/src/player"
)

var chr player.Character


type AI struct{
	name string
	atk int
	pv int
	pvmax int
	coinReward int
}
var dummy AI


func (e *AI) initEnemy(n string, h int, a int) {
	e.name = n
	e.pvmax = h
	e.pv = e.pvmax
	e.atk = a
}

func setUpEnemies(e AI, n string, h int, a int) {
	e.initEnemy(n, h, a)
}

func StartBattle(p *player.Character, e AI, n string, h int, a int) {
	setUpEnemies(e, n, h, a)
	e.battle(p)
}

func (e *AI) battle(p *player.Character) {
	for true {
		playerTurn := bubble.StartChoice([]string{"Attaquer.", "Défendre."}, true)
		plrDefTurn := 0
		switch playerTurn{
		case 0: //Le joueur attaque
		e.pv -= p.ATK
		case 1: //Le joueur se défend
			plrDefTurn = p.DEF
		case 2: //Le joueur accède à l'inventaire
			
		}

		if (e.atk - p.DEF) > 0 {
			p.PV -= e.atk - plrDefTurn
		}
		
		fmt.Println(e.pv, e.atk, plrDefTurn, p.PV)

		if p.PV <= 0 || e.pv <= 0 {
			return
		}
	}
}