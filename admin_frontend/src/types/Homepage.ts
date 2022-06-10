export type Homepage = {
	id: string;
	lang: string;
	content: HomepageContent;
}

export type HomepageContent = {
	hero_title: string;
	hero_subtitle: string;
	about_title: string;
	shopping_title: string;
	how_works_title: string;
	commission_title: string;
	why_us_title: string;
	about_text: string[];
}