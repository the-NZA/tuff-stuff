// HandleHomepageOnScrollAction adds or removes class based on scroll position
function HandleHomepageOnScrollAction(): void {
	// Check window location value in "/", "ru" or "en"
	const location = window.location.pathname;
	const isHomepage = location === "/" || location === "/ru" || location === "/en";

	// If not homepage, return
	if (!isHomepage) {
		return;
	}

	// Create observer options
	const observerOptions = {
		rootMargin: "0px",
		threshold: 0.1,
	};

	// Get main element with class "site-main"
	const mainElement = document.querySelector(".site-main");

	// Get first section in main element
	const firstSection = mainElement!.querySelector("section");

	// Create observer and observe main element
	const observer = new IntersectionObserver(HandleHomepageScrollCallback, observerOptions);
	observer.observe(firstSection as Element);
}

// Callback for homepage scroll
function HandleHomepageScrollCallback(entries: IntersectionObserverEntry[], _observer: IntersectionObserver): void {
	// Get first entry
	const firstEntry = entries[0];

	// Get header element with class "site-header"
	const headerElement = document.querySelector(".site-header");

	// If first entry is not intersecting add active class or remove otherwise
	if (!firstEntry.isIntersecting) {
		headerElement!.classList.add("site-header--active");
	} else {
		headerElement!.classList.remove("site-header--active");
	}
}

export default HandleHomepageOnScrollAction;