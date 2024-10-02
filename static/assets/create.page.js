const selector = document.querySelector("[data-option-config]").cloneNode();

function addOption() {
	const config = selector;
	const newNode = config.cloneNode(true);
	config.parentNode.appendChild(newNode);
}

function removeOption(e) {
	if(document.querySelectorAll("[data-option-config]").length > 1) {
		e.parentNode.remove();
	}
}
