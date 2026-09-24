package battle

import (
	"fmt"
	"projet/src"
	"projet/src/bubble"
	"projet/src/common"
	"projet/src/item"
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

func StartBattle(p *player.Character, e AI, n string, h int, a int) {
	Enemy := &e
	Enemy.initEnemy(n, h, a)
	Enemy.battle(p)
}

func spell(sn string, p *player.Character) int {
	switch sn{
	case "Charge. [8 PR]":
		if p.PR >= 8 {
			p.PR -= 8
			return p.ATK + (p.ATK/2)
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
			if p.PV > p.PVMAX { p.PV = p.PVMAX } 
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

func (e *AI) battle(p *player.Character) {
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
			s := bubble.StartChoice(specials,false)
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
			if len(player.INV) == 0 {
				fmt.Println("\nVotre inventaire est vide !")
				continue
			}
			invOptions := []string{"Retour."}
			var itemKeys []string
			for k, v := range player.INV {
				invOptions = append(invOptions, fmt.Sprintf("%s (x%d)", k, v))
				itemKeys = append(itemKeys, k)
			}
			choixInv := bubble.StartChoice(invOptions, false)
			if choixInv == 0 {
				continue
			}
			nomItem := itemKeys[choixInv-1]
			switch nomItem {
			case item.Jus.Nom:
				p.PV += item.Jus.Soin
				if p.PV > p.PVMAX {
					p.PV = p.PVMAX
				}
				player.SupInventory(nomItem)
				fmt.Printf("\nVous buvez un %s. Vous récupérez +%d PV ! (PV : %d/%d)\n", nomItem, item.Jus.Soin, p.PV, p.PVMAX)
			case item.Pizza.Nom:
				e.pv -= item.Pizza.Degats
				player.SupInventory(nomItem)
				fmt.Printf("\nVous lancez une %s sur %s ! Il subit %d dégâts ! (PV restant : %d/%d)\n", nomItem, e.name, item.Pizza.Degats, e.pv, e.pvmax)
			default:
				fmt.Printf("\nVous ne pouvez pas utiliser %s en combat !\n", nomItem)
				continue
			}
		}

		if e.pv <= 0 {
			common.DisplayBattle(tour, e.name, e.pv, e.pvmax, p.NAME, p.PV, p.PVMAX, p.PR, p.PRMAX)
			common.DisplayNarration("Le combat est gagné.")
			return
		}

		if (e.atk - p.DEF) > 0 {
			p.PV -= e.atk - plrDefTurn
		}

		fmt.Println(e.pv, e.atk, plrDefTurn, p.PV)
		common.DisplayBattle(tour, e.name, e.pv, e.pvmax, p.NAME, p.PV, p.PVMAX, p.PR, p.PRMAX)

		if p.PV <= 0 {
			common.DisplayNarration("Le combat est perdu...")
			if lives <= 0 {
				common.DisplayNarration("Dans ce monde irréel, vous êtes succombé à ce douloureux combat. Vous vous ne réveillez plus jamais.")
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