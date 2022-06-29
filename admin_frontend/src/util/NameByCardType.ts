import {CardType} from "../types/CardType";

export default function NameByCardType(cardType: CardType): string {
	switch (cardType) {
		case CardType.Commission:
			return "комиссионного ресурса";
		case CardType.Shopping:
			return "персонального шоппинга";
		case CardType.WhyUs:
			return "почему выбирают нас";
		case CardType.HowWorks:
			return "как работает сервис";
		default:
			throw new Error(`Unknown card type: ${cardType}`);
	}
}