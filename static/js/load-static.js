/* Load static pages dynamically */

guideContainerSelector = "#guide-container"
aboutContainerSelector = "#about-contaainer"

// cloneAndAddSequence ...
function cloneAndAddGuide() {
	fetch("./guide.htm")
	  .then(response => {
	    return response.text()
	  })
	  .then(data => {
	    $(data).insertAfter(guideContainerSelector)
	  });
}

// cloneAndAddSequence ...
function cloneAndAddAbout() {
	fetch("./about.htm")
	  .then(response => {
	    return response.text()
	  })
	  .then(data => {
	    $(data).insertAfter(aboutContainerSelector)
	  });
}
