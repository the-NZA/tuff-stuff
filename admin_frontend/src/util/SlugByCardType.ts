import {CardType} from "../types/CardType";

export default function SlugByCardType(cardType: CardType): string {
	switch (cardType) {
		case CardType.Shopping:
			return `shopping-card`
		case CardType.Commission:
			return `commission-card`
		case CardType.HowWorks:
			return `how-works-card`
		case CardType.WhyUs:
			return `why-us-card`
		default:
			throw new Error(`Unknown card type: ${cardType}`);
	}
}