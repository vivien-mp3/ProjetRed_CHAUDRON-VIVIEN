//ajout d'une économie permetant d'echanger ou acheter des chose plus tard

package economie

import (
	"fmt"
	"projet/src/player"
)

// func permetant le rajout de monnaie
func AjoutMonnaie(p *player.Character, montant int )  {
		p.MONNAIE += montant
		fmt.Println("Vous avez reussie la transaction")
	}

// func permettant le retrait en cas d'echange
func RetraitMonnaie(p *player.Character, retrait int ){
	if p.MONNAIE >= retrait {
		p.MONNAIE -= retrait
	}else{
		fmt.Println("vous avez pas assez d'argent ")
	}
}
	
