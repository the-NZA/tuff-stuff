// Import styles
import "@csstools/normalize.css"
import "./styles/index.css"

// Import header functions
import HandleHeaderOnScrollAction from "./scripts/header/header_on_scroll_action";

// Import home page functions
import HandleArrowDownOnClickAction from "./scripts/home/arrow_down_on_click_action";

window.addEventListener("DOMContentLoaded", function () {
	HandleHeaderOnScrollAction();
	HandleArrowDownOnClickAction();
})

