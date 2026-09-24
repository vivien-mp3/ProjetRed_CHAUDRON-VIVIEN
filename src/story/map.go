package story

import (
	"projet/src/bubble"
	"projet/src/common"
)

func BeginAdventure() {

	l1 := "Après que la machine s'est rappelé de vôtre nom. Elle s'en vient à s'écraser, le vaisseau en miettes...\n"
	l2 := "Mais, tout vas bien, enfin, vous vous sentiez mal, comme si vous étiez à moitié vivant. Vous regardez les environs et...\n"
	l3 := "Vous voyez le reste de votre ex-vaisseau, toujours accessible. Derrière vous se trouve des trâces de pattes de chat, \ns'en allant à l'opposé de vous. Vous n'êtes pas seul.\n"
	
	common.DisplayNarration(l1+l2+l3)

	c := bubble.StartChoice([]string{"Fouiller ce qu'il reste de <<chez-vous>>", "Suivre les trâces de pas.", "Contempler le ciel."}, false)
	switch c{
	case 0:
		common.DisplayNarration("Vous chercher dans votre vaisseau, soulever les plaques et pour une raison que vous ignorez, \nvous semblez chercher quelque chose de précis. \nVotre téléphone ? Des rations ? Votre nounours en peluche ? Non. vous l'avez trouvé, une pîle.\n Félicitations.")
	case 1:
		common.DisplayNarration("Vous suivez ses pas de chat. Vous vous demandez comment un chat peut survivre ici... Cet endroit à l'air d'être sauvage.")
		forest_cat_place()
	case 2:
		common.DisplayNarration("Le ciel... Pourquoi ? Tu es cloué sur place. Tu peux rêver de t'envoler, mais cette planète sera ta tombe.\n Même si.\n en regardant de plus près, le ciel est brisé. Des fissures recouvre le ciel.\n Dans ses fissures, de petites coccinelles charbonne pour recoudre le ciel.\n Tu passes de longues minutes à regarder le ciel, les coccinelles ont du mal à y arriver. Elle ne font aucun progrès.")	
		forest_lake()
	}

}

/*
	LA FORET AU COULEUR CHAUDE.
	Une plaine rempli d'arbres, l'ammbiance est chaude.. aux couleurs visant le rouge jusqu'au jaune.. Un bel endroit chaleureux pour un début... mécréant.
*/

func forest_lake() {
	common.DisplayNarration("Bizarre, Vous descendez les yeux et cette forêt est très orange. Les arbres jaune avec un dégradé orange, le feuilles rouge.\n Vous remarquez l'herbe, jaune vif à vos pieds, dégradant à l'orange quelques mètres plus loin puis au rouge à l'horizon.\n Vous vous retournez pour regagner le vaisseau, espérant que le panel de contrôle marche toujours.\n Hélas pour vous, derrière vous se trouvait un lac, de l'eau violette,\n contrastant à la nature rouge, au ciel couleur pomme et aux fissures roses...")
}

func forest_cat_place() {
	common.DisplayNarration("Vous êtes arrivé à la fin de ces empreintes.\n Mais, à vôtre surprise, un couloir, les arbres oranges aux feuilles rouges crée ce couloir.\n Et puis, ce chat. Il sort des arbres et s'assoit à quelques mètres de vous.")
	common.DisplayDialogue("chat", "Ce chat.", "Miaou miaou miaou miaou.")
	common.DisplayNarration("Bizarre... ce chat essaye de parler avec vous ? Vous marcher vers lui.")
}

