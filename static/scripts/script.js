// https://css-tricks.com/drag-and-drop-file-uploading/

function move(fill) {
  var elem = document.getElementById("Fill");
  var width = 1;
  var id = setTimeout(setInterval(frame, 10), 500);
  function frame() {
    if (width >= fill) {
      clearInterval(id);
    } else {
      width++;
      elem.style.width = width + "%";
      document.getElementById("value").innerHTML = width * 1 + "%";
    }
  }
}

function maximize(id) {
  var modal = document.getElementById("myModal");
  var modalImg = document.getElementById("img01");
  var img = document.getElementById(id);
  var captionText = document.getElementById("caption");
  modal.style.display = "block";
  console.log(this);
  modalImg.src = img.src;
  captionText.innerHTML = img.alt;
}

// Get the <span> element that closes the modal
var span = document.getElementsByClassName("close")[0];

// When the user clicks on <span> (x), close the modal
function hideModal() {
  var modal = document.getElementById("myModal");
  modal.style.display = "none";
}


function AddLink(){
  icon = document.getElementById("LinkForm")
  clon = icon.cloneNode(true)
  icon.querySelector(".nameCol input[type='text']").value=""
  icon.querySelector(".nameCol .icon").src=""
  $(clon).removeAttr("id")
  clon.RemoveChild(clon.lastChild)
  icon.parentElement.insertBefore(clon,icon)
}

function addField() {
  //https://www.sanwebcorner.com/2017/02/dynamically-generate-form-fields-using.html
  var max_fields = 5;
  var wrapper = $(document.getElementById("lenguage_container"));
  var x = 1;
  if (x < max_fields) {
    x++;
    wrapper.append(
      '<input type="text" name="lenguaje[]" placeholder="Lenguaje">'
    );
  } else {
    alert("limit reachead");
  }
}

function something(formID) {
  dragNdrop();
  // $(formID).on("submit", function (event) {
  console.log("A");
  SetVariables();
  console.log("B");
}
