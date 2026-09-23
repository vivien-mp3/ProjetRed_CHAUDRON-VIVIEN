package battle

import (
	"fmt"
	"projet/src"
	"projet/src/bubble"
	"projet/src/common"
	"projet/src/player"
)

var chr player.Character

var lives int = 1

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


func StartBattle(p *player.Character, e AI, n string, h int, a int) {
	Enemy := &e
	Enemy.initEnemy(n, h, a)
	Enemy.battle(p)
}

func spell(sn string, pr *int) int {
	switch sn{
	case "Test":
		if *pr >= 10 {
			*pr -= 10
			return 20
		}
	case "ULTIMATE":
		if *pr >= 20 {
			*pr -= 20
			return 80
		}
	}
	return -1
}

func (e *AI) battle(p *player.Character) {
	fmt.Println(e)
	p.PR = 0
	tour := 1
	for true {
		playerTurn := bubble.StartChoice([]string{"Attaquer.","Spécial.", "Défendre."}, true)
		plrDefTurn := 0
		switch playerTurn{
		case 0: //Le joueur attaque
		e.pv -= p.ATK
		case 1: //Le joueur fait une attaque spéciale
			fmt.Printf("\n\tVous avez: %d / %d\n", p.PR, p.PRMAX)
			specials := []string{"Retour.", "Test - 10PR", "ULTIMATE - 20PR"}
			s := bubble.StartChoice(specials,false)
			if specials[s] == "Retour." {
				continue
			} else {
				spelldmg := spell(specials[s], &p.PR)
				if spelldmg < 0 {
					continue
				}
				e.pv -= spelldmg
			}
		case 2: //Le joueur se défend
			plrDefTurn = p.DEF
			p.PR += p.DEF * 3
		case 3: //Le joueur accède à l'inventaire
			
		}

		if (e.atk - p.DEF) > 0 {
			p.PV -= e.atk - plrDefTurn
		}
		
		fmt.Println(e.pv, e.atk, plrDefTurn, p.PV)
		common.DisplayBattle(tour,e.name, e.pv, e.pvmax, p.NAME, p.PV, p.PVMAX, p.PR, p.PRMAX)
		

		if p.PV <= 0 {
			common.DisplayNarration("Le combat est perdu...")
			if lives <= 0 {
				src.Startmenu()
			}
			lives--
			return
		} else if e.pv <= 0 {
			common.DisplayNarration("Le combat est gagné.")
			return
		}
		tour++
	}
}