function Edit(block, id) {
  var values = document.getElementById(id);
  document.getElementById(
    "notesTitleForm"
  ).value = values.getElementsByClassName("noteTitle")[0].innerHTML;
  document.getElementById("note_id").value = values.getElementsByClassName(
    "id"
  )[0].innerHTML;
  document.getElementById(
    "notesTextForm"
  ).value = values
    .getElementsByClassName("noteText")[0]
    .getElementsByTagName("p")[0].innerHTML;
  $("#NotesForm").on("submit", function (event) {
    event.preventDefault();
    var pr = document.querySelectorAll("input[name=lastName]")[0].value;
    var url = "/Update/Note/" + id;
    $.ajax({
      url: url,
      method: "PUT",
      data: $("#NotesForm").serialize(),
      success: function (data) {
        alert("Correctamente");
        setTimeout(function () {
          location.reload();
        }, 0001);
      },
      error: function (request, msg, error) {
        alert("Error");
      },
    });
  });
}

function Delete(url, id) {
  var response = confirm("Estas seguro que deseas eliminar?");
  if (response == true) {
    $.ajax({
      url: url + id,
      method: "DELETE",
      contentType: "application/json",
      success: function (result) {
        alert("Eliminado Correctamente");
        setTimeout(function () {
          location.reload();
        }, 0001);
      },
      error: function (request, msg, error) {
        alert("Error al Eliminar");
      },
    });
  }
}

function formUpdate(id, url) {
  $("#" + id).on("submit", function (event) {
    event.preventDefault();
    $.ajax({
      url: url,
      method: "PUT",
      data: $("#" + id).serialize(),
      success: function (data) {
        alert("Correctamente");
        setTimeout(function () {
          location.reload();
        }, 0001);
      },
      error: function (request, msg, error) {
        alert("Error");
      },
    });
  });
}

function DeleteLink(id) {
  var response = confirm("Deseas Eliminar ?");
  if (response == true) {
    $.ajax({
      url: "/Delete/Link/" + id,
      method: "DELETE",
      contentType: "application/json",
      success: function (result) {
        alert("Eliminado Correctamente");
        document.getElementById("L-" + id).style.display = "none";
      },
      error: function (request, msg, error) {
        alert("Error al Eliminar");
      },
    });
  }
}

function EditLink(index, id){
  var content = "application/x-www-form-urlencoded"
  var box = document.getElementById(id)
  var text = box.querySelector(".nameCol input[type='text']")
  var file = box.querySelector(".nameCol input[type='file']")
  var aceptar = box.querySelector(".nameCol .aceptar")
  var edit = box.querySelector(".nameCol .editbutton")
  edit.style.display = "none";
  aceptar.style.display = "block";
  text.disabled = false
  file.disabled = false
  $(box).on("submit", function (event) {
    event.preventDefault();
    var data = new FormData();
    text = box.querySelector(".nameCol input[type='text']")
    file = box.querySelector(".nameCol input[type='file']")
    console.log(text.value)
    console.log(file.value)
    data.append("text", text.value)
    //data.append("file", file.value)
    if (file.value == "" || file.value == null){
      content ="multipart/form-data"
    }
  $.ajax({
    url: "/Update/Link/"+index,
    method:"PUT",
    contentType:false,
    processData : false ,
    data: new FormData(this),//$(text).serialize(),//+'&'+$(file).serialize(),
    success: function (dats) {
      alert("exito")
      edit.style.display = "block";
      aceptar.style.display = "none";
      text.disabled = true
      file.disabled = true
    },
    error: function(request, msg, error){
      alert("Error")
    }
  });
});
}
