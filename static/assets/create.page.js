function addOption() {
	const config = document.querySelector("[data-option-config]");
	const newNode = config.cloneNode(true);
	newNode.querySelector('input[type="text"]').value = '';
	newNode.querySelector('input[type="file"]').value = '';
	config.parentNode.appendChild(newNode);
}

function removeOption(e) {
	if(document.querySelectorAll("[data-option-config]").length > 1) {
		e.parentNode.remove();
	}
}
