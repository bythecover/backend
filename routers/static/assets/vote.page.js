const numberOfOptions = document.querySelectorAll("[data-carousel-card]").length;
let currentIndex = 0;
let currentSelection = document.querySelectorAll("[data-carousel-card]")[currentIndex];
document.querySelector("#" + currentSelection.getAttribute("for")).setAttribute("checked", "");

document.querySelectorAll("[data-carousel-card]").forEach((el) => {
	el.addEventListener("click", () => { el.scrollIntoView({ behavior: "smooth", inline: "center" }) });
});

function decrementSelection() {
	if(currentIndex <= 0) {
		return;
	}

	document.querySelector("#" + currentSelection.getAttribute("for")).removeAttribute("checked");
	currentIndex-=1;
	currentSelection = document.querySelectorAll("[data-carousel-card]")[currentIndex];
	document.querySelector("#" + currentSelection.getAttribute("for")).setAttribute("checked", "");
	currentSelection.scrollIntoView({ behavior: "smooth", inline: "center" });
}

function incrementSelection() {
	if(currentIndex >= numberOfOptions - 1) {
		return;
	}

	document.querySelector("#" + currentSelection.getAttribute("for")).removeAttribute("checked");
	currentIndex += 1;
	currentSelection = document.querySelectorAll("[data-carousel-card]")[currentIndex];
	document.querySelector("#" + currentSelection.getAttribute("for")).setAttribute("checked", "");
	currentSelection.scrollIntoView({ behavior: "smooth", inline: "center" });
}
