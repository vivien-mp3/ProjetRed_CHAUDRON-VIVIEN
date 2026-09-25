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

type AI struct {
	name       string
	atk        int
	pv         int
	pvmax      int
	coinReward int
}

var dummy AI

func (e *AI) initEnemy(n string, h int, a int) {
	e.name = n
	e.pvmax = h
	e.pv = e.pvmax
	e.atk = a
}

func StartBattle(p *player.Character, e AI, n string, h int, a int) bool {
	if n == "Le chat" {
		common.PlaySound("cattheme")
	} else {
		common.PlaySound("battletheme")
	}

	Enemy := &e
	Enemy.initEnemy(n, h, a)
	win := Enemy.battle(p)
	return win
}

func spell(sn string, p *player.Character) int {
	switch sn {
	case "Charge. [8 PR]":
		if p.PR >= 8 {
			p.PR -= 8
			return p.ATK + (p.ATK / 2)
		}
	case "Papier tranchant. [10 PR]":
		if p.PR >= 10 {
			p.PR -= 10
			return p.ATK * 2
		}
	case "Vampirisme. [16 PR]":
		if p.PR >= 16 {
			p.PR -= 16
			p.PV += p.ATK
			if p.PV > p.PVMAX {
				p.PV = p.PVMAX
			}
			return p.ATK * 2
		}
	case "Triangle de lumière. [32 PR]":
		if p.PR >= 16 {
			p.PR -= 16
			return 65
		}
	}
	return -1
}

func (e *AI) battle(p *player.Character) bool {
	fmt.Println(e)
	p.PR = 100
	tour := 1
	for true {
		playerTurn := bubble.StartChoice([]string{"Attaquer.", "Spécial.", "Défendre."}, true)
		plrDefTurn := 0
		switch playerTurn {
		case 0: //Le joueur attaque
			e.pv -= p.ATK
		case 1: //Le joueur fait une attaque spéciale
			fmt.Printf("\n\tVous avez: %d / %d\n", p.PR, p.PRMAX)
			specials := p.SPELLS
			specials = append(specials, "Retour.")
			s := bubble.StartChoice(specials, false)
			if specials[s] == "Retour." {
				continue
			} else {
				spelldmg := spell(specials[s], p)
				fmt.Println(p)
				if spelldmg < 0 {
					continue
				}
				e.pv -= spelldmg
				fmt.Println(p)
			}
		case 2: //Le joueur se défend
			plrDefTurn = p.DEF
			p.PR += p.DEF * 3
		case 3: //Le joueur accède à l'inventaire
			player.AccesInventory(&p.INV)
		}

		if (e.atk - p.DEF) > 0 {
			p.PV -= e.atk - plrDefTurn
		}

		fmt.Println(e.pv, e.atk, plrDefTurn, p.PV)
		common.DisplayBattle(tour, e.name, e.pv, e.pvmax, p.NAME, p.PV, p.PVMAX, p.PR, p.PRMAX)

		if p.PV <= 0 {
			common.DisplayNarration("Le combat est perdu...")
			if lives <= 0 || e.name != "Le chat" {
				common.DisplayNarration("Dans ce monde irréel, vous êtes succombé à ce douloureux combat. Vous vous ne réveillez plus jamais.")
				src.Startmenu()
			}
			lives--
			return false
		} else if e.pv <= 0 {
			common.DisplayNarration("Le combat est gagné.")
			return true
		}
		tour++
	}
	return false
}
