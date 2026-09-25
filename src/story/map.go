package story

import (
	"fmt"
	"projet/src"
	"projet/src/battle"
	"projet/src/bubble"
	"projet/src/common"
	"projet/src/economie"
	"projet/src/player"
	"time"
)


var plr *player.Character

var tookPile bool
var metCatPlace bool
var trashmanBeaten bool

func BeginAdventure(p *player.Character) {
	common.PlaySound("orangeforest")
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
	common.PlaySound("orangeforest")
	common.DisplayNarration("Bizarre, Vous descendez les yeux et cette forêt est très orange. Les arbres jaune avec un dégradé orange, le feuilles rouge.\n Vous remarquez l'herbe, jaune vif à vos pieds, dégradant à l'orange quelques mètres plus loin puis au rouge à l'horizon.\n Vous vous retournez pour regagner le vaisseau, espérant que le panel de contrôle marche toujours.\n Hélas pour vous, derrière vous se trouvait un lac, de l'eau violette,\n contrastant à la nature rouge, au ciel couleur pomme et aux fissures roses...")
	
	for true {
		c := bubble.StartChoice([]string{"Regarder le lac", "Partir en direction de la ville", "Aller voir le forgeron."}, false)
		switch c{
		case 0:
			common.DisplayNarration("Vous observez le lac, c'est bien, vous voyez 3 amis au loin, un chauve, un gros et un grand. Ils s'amusent.")
			continue
		case 1:
			common.DisplayNarration("Vous partez en direction de la ville, voyant les immeubles gigantesques, ils dépassent même le ciel.")
			town_entrance()
		case 2:
			common.DisplayNarration("Vous allez vers la forge, une forge en forêt ? Dangereux. Mais bon, pas le choix.")
			
		}
	}
}

func forest_cat_place() {
	common.PlaySound("orangeforest")
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

/*
	La ville
*/

func town_entrance() {
	common.PlaySound("towntheme")
	common.DisplayNarration("Vous rentrez en ville, les immeubles montent jusqu'au ciel.\n Vous decidiez de regarder à l'intérieur du rez-de-chaussez de l'un d'entre-eux..\n Vous voyez à l'intérieur le toit du bâtiment, avec des bancs et des buissons, très jolie..\n Vous continuez votre chemin et vous voyez une auberge.")

	for true {
		c := bubble.StartChoice([]string{"Rentrer dans l'auberge", "Explorer la ville", "Retourner dans la forêt"}, true)
		switch c{
		case 0:
			common.DisplayNarration("Vous vous dirigiez vers l'auberge, le chemin semble long pour 5 mètres de marche.")
			auberge()
		case 1:
			if !trashmanBeaten {
				common.DisplayNarration("Vous explorez la ville, Vous croisez une poubelle avec des bras et des jambes...\n Il semble vouloir votre peau, littéralment.")
				var trashman battle.AI
				if battle.StartBattle(plr, trashman, "Trash-Man", 75, 6) {
					common.DisplayNarration("Le combat a été rude, mais.. pas d'inquiétude, vous êtes encore vivant.")
				} else {
					common.DisplayNarration("Vous avez certes perdu mais.. après vous êtes réveillé, vous êtes encore vivant")
				}
				trashmanBeaten = true
			} else {
				common.DisplayNarration("Vous êtes en ville, c'est vide.. Mais continuez votre chemin.")
			}
			town_end()
		case 2:
			common.DisplayNarration("Vous faites demi-tour et parter vers la forêt, en levant les yeux en l'air.")
			forest_lake()
		case 3:
			player.AccesInventory(&plr.INV)
			continue
		}
	}
}

func town_end() {
	common.PlaySound("towntheme")

	common.DisplayNarration("Vous continuez plus loin dans la ville.\nVous voyez pas grand chose appart le chat, il vous remarque et miaule, que c'est mignon.\n Il va grimper en haut d'un grand arbre à chat.\n La ville est silencieuse et monotone.. Vous sentez mal dedans..")

	for true {
		c := bubble.StartChoice([]string{"Grimper l'arbre à chat.", "Faire demi tour"}, true)
		switch c{
		case 0:
			common.DisplayNarration("Vous grimper l'arbre à chat, C'est pénible, vous avez du mal à vous accrocher, normal. Vous n'êtes pas un chat.")
			treecat()
		case 1:
			common.DisplayNarration("Vous faites demi-tour... Cependant même si l'allée était plate, le retour est en pente...\n Pire. Vous êtes dans un trou et vous remontez avec la petite echelle de sécurité.")
			town_entrance()
		case 2:
			player.AccesInventory(&plr.INV)
			continue
		}
	}
}

func auberge() {
	common.PlaySound("lostintime")
	common.DisplayNarration("Vous rentrez dans l'auberge\n Le barman semble de nettoyez des verres, vous regardez plus en détails et vous voyez un ange déchu... à son 7eme verre.. \nEt un placard avec marquer <<Mégazin>> dessus...")
	
	for true {
			c := bubble.StartChoice([]string{"Allez voir le barman.", "Allez voir l'ange déchu.", "Allez voir ce <<magasin>>", "Quitter l'auberge."}, true)
			switch c{
			case 0:
				common.DisplayDialogue("?", "Barman", "Ah ! Jeune matelot ! Comment aimes-tu ma belle auberge !\n Yargh! Voudrais tu te reposer dans une piaule ?")
				c2 := bubble.StartChoice([]string{"Allez vous reposez.", "Refusez l'offre"}, true)
					switch c2{
					case 0:
						common.DisplayNarration("Vous decidez donc d'accepter l'offre.\n Vous rentrez dans la chambre et en allumant la lumière, vous appercevez que la pièce se plonge dans le noir,\n mais que les coins sont un peu éclairer, vous éteignez donc la lumière et la pièce deviens normale...\nAprès cette découverte, vous partez vous allonger dans le lit et vous regardez le réveil,\n il est minuit moins douze. Puis vous fermez les yeux.")
						plr.PV = plr.PVMAX
						fmt.Printf("\n\t Vous vous êtes bien reposé, votre vie a été soignée complètement: %d", plr.PV)
						continue
					case 1:
						common.DisplayNarration("Vous refusez l'offre et regarder ailleurs.")
						continue
					}
			case 1:
				common.DisplayNarration("Vous allez voir l'ange déchu.. Vous ne savez pas vraiment pourquoi.")
				common.DisplayDialogue("angel", "L'ange déchu", "Ce monde de barbarie. Personne n'a l'espoir d'atteindre le haut au final..")
				common.DisplayNarration("Vous hésitez à lui faire une remarque mais vaut mieux pour ne pas de le déranger.")
				continue
			case 2:
				common.DisplayNarration("Vous allez voir le marchand dans le placard... Une musique stupide se déroule..")
				common.PlaySound("stupidtheme")
				economie.Boutique(plr)
			case 3:
				common.DisplayNarration("Vous quittez l'auberge, vous entendez le barman commencer à pleurer de voir un client partir.. Cela vous fait de la peine mais bon.")
				town_entrance()
			case 4:
				player.AccesInventory(&plr.INV)
				continue
		}
	}
}

func treecat() {
	common.DisplayDialogue("cat", "Ce chat", "Miaou, miaou miaou miaou !")
	common.DisplayNarration("Vous revoyez ce chat, il semble vous apprécier, ces pupilles sont dilatés, que c'est mignon.\n Mais.\nVous vous rappelé que les chat dilates leur yeux quand ils voient une proie.")
	var cat battle.AI
	win := battle.StartBattle(plr, cat, "Le chat", 999, 35)
	if win {
		common.DisplayNarration("Tu as... survécu..? La pauvre bête est par terre, immobile.. Tu n'as aucun remord, enfin, c'est que tu te dis. Mais au fond de toi, t'aurais aimé que ce chat gagne.")
		common.DisplayDialogue("joueur", plr.NAME, "gg ez")
		credit()
		return
	} else {
		common.DisplayNarration("C'était à prévoir... Bravo, vous êtes à terre, le chat à gagné.")
		credit()
		return
	}
	
}

func credit() {
	common.DisplayTitle()
	fmt.Printf("\nUn jeu fait par:\n\t- RACHETER Lioris\n\t- JULIEN Mathieu\n\t- CHAUDRON Vivien")

	time.Sleep(10 * time.Second)

	src.Startmenu()
}

