// Function handles click on arrow down button and scrolls to the next section
function HandleArrowDownOnClickAction() {
	const arrowDown = document.querySelector('.hero__arrow_down');

	// If now arrow down button is found, return
	if (!arrowDown) {
		return;
	}

	// Handle on click event
	arrowDown.addEventListener('click', () => {
		// Scroll to the next section
		window.scrollTo({
			top: window.innerHeight,
			behavior: 'smooth',
		});
	});
}

export default HandleArrowDownOnClickAction;