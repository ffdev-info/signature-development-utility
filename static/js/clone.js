/* clone.js */

const btnAdd = "#btnAdd";
const btnDel = "#btnDel";
const pathAdd = "#pathAdd";
const seqAdd = "#seqAdd";
const triggerOnly = "#trigger-only";

$(document).ready(function() {

    // Setup a standard signature form.
    cloneAndAddStandardSequences();

    // Setup a container signature form.
    cloneAndAddPath();

    // Add static components to pages.
    cloneAndAddGuide()
    cloneAndAddAbout()

    $(btnAdd).click(function() {
        cloneAndAdd();
    })

    $(btnDel).click(function() {
        removeClone();
    })

    $(pathAdd).click(function() {
        cloneAndAddPath();
    })

    $(seqAdd).click(function() {
        cloneAndAddContainerSequences();
    })

    $(pathDel).click(function() {
        deletePath();
    })

    $(seqDel).click(function() {
        deleteSequence();
    })

    $(triggerOnly).click(function() {
      updateSequenceFields(triggerOnly);
    });

});

// cloneeAndAddPath ...
function cloneAndAddStandardSequences() {
    const standardSequencesSelector = "#standard-signatures-container"
    data = $.parseHTML( standardSignatureSequences ),
    $(data).appendTo(standardSequencesSelector);
}

// disableRemove will enable and or disable the remove button from the
// UI when the number of elements are at the correct threshold.
function disableRemove() {

    const initialMin = 1;
    const minToRemove = 2;

    len = $('.cloned-input').length;
    if (len == minToRemove) {
        $(btnDel).removeAttr("disabled");
    }
    if (len == initialMin) {
        $(btnDel).attr('disabled', 'disabled');
    }
}

// _makeElementName returns to us an element name nicely formatted.
function _makeElementName(elemName, idNo) {
    return "#".concat(elemName, idNo)
}

// _setIDandName in a single bound.
function _setIDAndName(elem, value) {
    elem.attr("name", value);
    elem.attr("id", value);
}

// cloneAndAdd will clone the signature input boxes on a page and
// append them to the DOM for submission via the DOM.
function cloneAndAdd() {

    const maxInputs = 3;
    const clonedInput = ".cloned-input"
    const cloned = "cloned-input-";
    const sig = "signature-input-";
    const rel = "signature-relativity-";
    const off = "offset-";
    const max = "max-offset-";
    const disable = "disabled";

    // This loosely represents a PDF signature which means the utility
    // can also be used for demonstration purposes out of the box.
    placeholder = [
        "255044462D312E34",
        "2525454F46(0A|0D|0D0A)",
        "000000",
    ]

    // Vary the anchors as we create each new element to mirror
    // something like a signature development workflow.
    anchor = [
        "BOFoffset",
        "EOFoffset",
        "VARoffset"
    ]

    len = $(clonedInput).length;
    newIDX = len++;

    newID = "".concat(cloned, newIDX);
    newSig = "".concat(sig, newIDX);
    newRel = "".concat(rel, newIDX);
    newOff = "".concat(off, newIDX);
    newMax = "".concat(max, newIDX);

    clone = $(_makeElementName(cloned, "0")).clone();
    clone.attr("id", newID);

    signature = clone.find(_makeElementName(sig, "0"));
    relativity = clone.find(_makeElementName(rel, "0"));
    offset = clone.find(_makeElementName(off, "0"));
    maxOffset = clone.find(_makeElementName(max, "0"));

    signature.attr("placeholder", placeholder[newIDX]);
    signature.attr("value", placeholder[newIDX]);
    relativity.val(anchor[newIDX]).attr("selected", "selected");

    clone.appendTo("#inputs");

    // Check whether or not we need to disable the append button.
    len = $(clonedInput).length;
    if (len == maxInputs) {
        $(btnAdd).attr(disable, disable);
        return false;
    }

    disableRemove();
}

// removeClone will remove the cloned elements from the DOM.
function removeClone() {

    const cloned = "cloned-input-";
    len = $('.cloned-input').length;

    var removable = "".concat("#", cloned, len-1);

    $(removable).remove();
    $(btnAdd).removeAttr("disabled");

    disableRemove();
}

const fileSelector = "#files-container"
const containerSelector = "#container-path"
const containerPathID = "container-path-";
const containerSequencesID = "container-sequences-";

var currentContainerIndex = -1;
var currentSequenceCount = -1;
var allContainerSequenceCount = -1;

// cloneeAndAddPath ...
function cloneAndAddPath() {
    updateContainerIndex();
    data = $.parseHTML( newContainerFile ),
    $(data).find(containerSelector).attr("name", getContainerPathID());
    $(data).attr("id", getContainerPathID())
    $(data).appendTo(fileSelector);
    disablePath();

    // Sequence count is only used for the current path.
    resetSequenceCount();
}

// disablePath() will disable or enable path based on how many files
// are associated with the container signature.
function disablePath() {

    // Zero-based index, so 3 max. Arbitrary figure, which is largely
    // to help manage complexity. We might be able to remove this with
    // ease after more testing.
    const pathMax = 2
    const pathMin = 0
    const pathAdd = "#pathAdd"
    const pathDel = "#pathDel"

    if (currentContainerIndex < pathMax) {
        $(pathAdd).removeAttr("disabled");
    }
    if (currentContainerIndex == pathMax) {
        $(pathAdd).attr('disabled', 'disabled');
    }
    if (currentContainerIndex == pathMin) {
        $(pathDel).attr('disabled', 'disabled');
    }
    if (currentContainerIndex > pathMin) {
        $(pathDel).removeAttr("disabled");
    }
}

// deletePath will delete a path entry from the DOM and update the
// different counts that we rely on.
function deletePath() {
    var removable = getContainerPathSelector()
    $(removable).remove();
    decreaseContainerIndex()
    disablePath()
}

const sequenceSelector = "#container-signature"
const maxOffsetSelector = "#container-max-offset"
const offsetSelector = "#container-offset"
const relativitySelector = "#container-signature-relativity"

const sequenceField = "signature-"
const maxOffsetField = "max-offset-"
const offsetField = "offset-"
const relativityField = "relativity-"

// cloneAndAddSequence ...
function cloneAndAddContainerSequences() {

    // Updating the sequence count provides a helper to limit complexity
    // on this form.
    updateSequenceCount();

    data = $.parseHTML( newSignatureSequences ),

    // Create an ID to enable deleting sequences from the DOM.
    $(data).attr("id", getContainerSeuqneceID())

    // Find the ID within the new block of HTML and then dynamically
    // set the name to help with form processing.
    $(data).find(sequenceSelector).attr("name", makeContainerFieldName(sequenceField));
    $(data).find(offsetSelector).attr("name", makeContainerFieldName(offsetField));
    $(data).find(maxOffsetSelector).attr("name", makeContainerFieldName(maxOffsetField));
    $(data).find(relativitySelector).attr("name", makeContainerFieldName(relativityField));

    // Amend the DOM with our new data entry fields.
    $(data).appendTo(getContainerPathSelector());
}

// disableSequence enables deletion of sequences sequentially not per
// block. If there is at least one sequence on the page it can be
// deleted. This seems like it might be a little over-zealous, so it
// can be improved with more sophisticated heuristics.
function disableSequence() {
    const maxSequence = 2;
    const minSequence = -1;
    const seqAdd = "#seqAdd";
    const seqDel = "#seqDel";

    if (allContainerSequenceCount == minSequence) {
        $(seqDel).attr('disabled', 'disabled');
    }
    if (allContainerSequenceCount > minSequence) {
        $(seqDel).removeAttr("disabled");
    }
    if (currentSequenceCount == maxSequence) {
        $(seqAdd).attr('disabled', 'disabled');
    }
    if (currentSequenceCount == minSequence) {
        $(seqAdd).removeAttr("disabled");
    }
}

function deleteSequence() {
    var removable = getContainerSeuqneceSelector()
    $(removable).remove();
    decreaseSequenceIndex()
}

// updateContainerIndex ...
function updateContainerIndex() {
    currentContainerIndex++
    updateContainerCountfField()
}

// updateContainerIndex ...
function decreaseContainerIndex() {
    currentContainerIndex--
    updateContainerCountfField()
}

// updateContainerCountField ...
function updateContainerCountfField() {
    const containerCountFieldSelector = "#container-file-count"
    $(containerCountFieldSelector).attr("value", (currentContainerIndex)+1);
}

// updateSequenceCount ...
function updateSequenceCount() {
    currentSequenceCount++
    allContainerSequenceCount++
    disableSequence();
}

// decreaseSequenceIndex ...
function decreaseSequenceIndex() {
    currentSequenceCount--
    allContainerSequenceCount--
    disableSequence();
}

// resetSequenceCount ...
function resetSequenceCount() {
    currentSequenceCount = -1
    disableSequence();
}

// getSelector ...
function getContainerPathID() {
    return "".concat(containerPathID, currentContainerIndex)
}

// getContainerPathSelector ...
function getContainerPathSelector() {
    return "#".concat(containerPathID, currentContainerIndex)
}

// makeContainerFieldName will create a name that can be retrieved and
// parsed by the server into a signature file.
function makeContainerFieldName(field) {
    return "".concat(field, currentContainerIndex)
}

// getContainerSequenceID ...
function getContainerSeuqneceID() {
    return "".concat(containerSequencesID, allContainerSequenceCount)
}

// getContainerSeuqneceSelector ...
function getContainerSeuqneceSelector() {
    return "#".concat(containerSequencesID, allContainerSequenceCount)
}

// updateSequenceFields is called when we only want to output a
// signature file to support container identification.
function updateSequenceFields(trigger) {
    const clonedInput = "#inputs";
    if (!$(trigger).checked) {
        $(clonedInput).toggle(100);
    } else {
        $(clonedInput).toggle(100);
    }
}
