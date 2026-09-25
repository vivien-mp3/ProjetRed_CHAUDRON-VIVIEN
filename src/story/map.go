package story

import (
	"projet/src/battle"
	"projet/src/bubble"
	"projet/src/common"
	"projet/src/player"
)


var plr *player.Character

var tookPile bool
var metCatPlace bool

func BeginAdventure(p *player.Character) {
	plr = p

	l1 := "Après que la machine s'est rappelé de vôtre nom. Elle s'en vient à s'écraser, le vaisseau en miettes...\n"
	l2 := "Mais, tout vas bien, enfin, vous vous sentiez mal, comme si vous étiez à moitié vivant. Vous regardez les environs et...\n"
	l3 := "Vous voyez le reste de votre ex-vaisseau, toujours accessible. Derrière vous se trouve des trâces de pattes de chat, \ns'en allant à l'opposé de vous. Vous n'êtes pas seul.\n"
	
	common.DisplayNarration(l1+l2+l3)
	for true {
		c := bubble.StartChoice([]string{"Fouiller ce qu'il reste de <<chez-vous>>", "Suivre les trâces de pas.", "Contempler le ciel."}, false)
		switch c{
		case 0:
			if !tookPile {
				common.DisplayNarration("Vous chercher dans votre vaisseau, soulever les plaques et pour une raison que vous ignorez, \nvous semblez chercher quelque chose de précis. \nVotre téléphone ? Des rations ? Votre nounours en peluche ? Non. vous l'avez trouvé, une pîle.\n Félicitations.")
				tookPile = true
				continue
			} else {
				common.DisplayNarration("Vous aviez déjà chercher dans vôtre vaisseau.. et rien d'important s'y trouve...")
				continue
			}
		case 1:
			common.DisplayNarration("Vous suivez ses pas de chat. Vous vous demandez comment un chat peut survivre ici... Cet endroit à l'air d'être sauvage.")
			forest_cat_place()
		case 2:
			common.DisplayNarration("Le ciel... Pourquoi ? Tu es cloué sur place. Tu peux rêver de t'envoler, mais cette planète sera ta tombe.\n Même si.\n en regardant de plus près, le ciel est brisé. Des fissures recouvre le ciel.\n Dans ses fissures, de petites coccinelles charbonne pour recoudre le ciel.\n Tu passes de longues minutes à regarder le ciel, les coccinelles ont du mal à y arriver. Elle ne font aucun progrès.")	
			forest_lake()
		}
	}
	
}


/*
	LA FORET AU COULEUR CHAUDE.
	Une plaine rempli d'arbres, l'ammbiance est chaude.. aux couleurs visant le rouge jusqu'au jaune.. Un bel endroit chaleureux pour un début... mécréant.
*/

func forest_lake() {
	common.DisplayNarration("Bizarre, Vous descendez les yeux et cette forêt est très orange. Les arbres jaune avec un dégradé orange, le feuilles rouge.\n Vous remarquez l'herbe, jaune vif à vos pieds, dégradant à l'orange quelques mètres plus loin puis au rouge à l'horizon.\n Vous vous retournez pour regagner le vaisseau, espérant que le panel de contrôle marche toujours.\n Hélas pour vous, derrière vous se trouvait un lac, de l'eau violette,\n contrastant à la nature rouge, au ciel couleur pomme et aux fissures roses...")
	var dummy battle.AI
	battle.StartBattle(plr, dummy, "dummy", 100, 10)
}

func forest_cat_place() {
	if !metCatPlace {
		common.DisplayNarration("Vous êtes arrivé à la fin de ces empreintes.\n Mais, à vôtre surprise, un couloir, les arbres oranges aux feuilles rouges crée ce couloir.\n Et puis, ce chat. Il sort des arbres et s'assoit à quelques mètres de vous.")
		common.DisplayDialogue("chat", "Ce chat.", "Miaou miaou miaou miaou.")
		common.DisplayNarration("Bizarre... ce chat essaye de parler avec vous ? Vous marcher vers lui.")
		common.DisplayDialogue("chat", "Ce chat.", "Miaou miaou, miaou miaou miaou miaou!")
		common.DisplayNarration("Ce chat vous donne un objet, un jus de banane..? Ce jus de banane à l'air d'être une source pour vous guérir")
		player.AddInventory("Jus de banane", 1, &plr.INV)

		common.DisplayNarration("On peut remercier ce chat, mais ce chat part à nouveau, en direction de la ville.")
	} else {
		common.DisplayNarration("Ici ce n'est juste qu'un couloir d'arbres, rien d'intéressant...")

		c := bubble.StartChoice([]string{"Regarder derrière un arbre", "Continuez de marcher."}, false)
		switch c{
		case 0:
			common.DisplayNarration("Vous regardez derrière l'arbre et il y a un homme, il tient un oeuf dans sa main...\n Vous decidez de quand même continuez.. Et vous avez soudainement oublié qu'est-ce qu'il y avait derrière cet arbre.")
		case 1:
			common.DisplayNarration("Vous faites demi-tour, et vous regardez le ciel un court instant.")
			forest_lake()
	}
}
	
	for true {
		c := bubble.StartChoice([]string{"Partir en ville", "Faire demi-tour"}, true)
		switch c{
		case 0:
			common.DisplayNarration("Vous partez en ville, vous suivez le chat mais rapidement il part derrière ")
			town_entrance()
		case 1:
			common.DisplayNarration("Vous faites demi-tour, et vous regardez le ciel un court instant.")
			forest_lake()
		case 2:
			player.AccesInventory(&plr.INV)
			continue
		}
	}
	

	
}

func town_entrance() {
	common.DisplayNarration("Vous rentrez en ville, les immeubles montent jusqu'au ciel.\n Vous decidiez de regarder à l'intérieur du rez-de-chaussez de l'un d'entre-eux..\n Vous voyez à l'intérieur le toit du bâtiment, avec des bancs et des buissons, très jolie..\n Vous continuez votre chemin et vous voyez une auberge.")

		for true {
		c := bubble.StartChoice([]string{"Rentrer dans l'auberge", "Explorer la ville"}, true)
		switch c{
		case 0:
			common.DisplayNarration("Vous partez en ville, vous suivez le chat mais rapidement il part derrière ")
			auberge()
		case 1:
			common.DisplayNarration("Vous explorez la ville,")
			forest_lake()
		case 2:
			player.AccesInventory(&plr.INV)
			continue
		}
	}
}

func auberge() {
	
}



