function addOption() {
	const config = document.querySelector("[data-option-config]")
	const newNode = config.cloneNode(true);
	config.parentNode.appendChild(newNode)
}


function removeOption(e) {
	if(document.querySelectorAll("[data-option-config]").length > 1) {
		e.parentNode.remove();
	}
}
