function dragNdrop() {
  var isAdvancedUpload = (function () {
    var div = document.createElement("div");
    return (
      "dragable" in div ||
      ("ondragstart" in div &&
        "ondrop" in div &&
        "FormData" in window &&
        "FileReader" in window)
    );
  })();

  var $form = $(".box");
  var $input = $form.find('input[type="file"]'),
    $label = $form.find("label"),
    showFiles = function (files) {
      $label.text(
        files.length > 1
          ? ($input.attr("data-multiple-caption") || "").replace(
              "{count}",
              files.length
            )
          : files[0].name
      );
      readURL(files);
    };
  $form.on("drop", function (e) {
    droppedFiles = e.originalEvent.dataTransfer.files; // the files that were dropped
    showFiles(droppedFiles);
  });

  $input.on("change", function (e) {
    showFiles(e.target.files);
  });

  console.log(isAdvancedUpload);
  if (isAdvancedUpload) {
    $form.addClass("hasAdvancedUpload");

    var droppedFiles = false;

    $form
      .on("drag dragstart dragend dragover dragenter dragleave drop", function (
        e
      ) {
        e.preventDefault();
        e.stopPropagation();
      })
      .on("dragover dragenter", function () {
        $form.addClass("is-dragover");
      })
      .on("dragleave dragend drop", function () {
        $form.removeClass("is-dragover");
      })
      .on("drop", function (e) {
        droppedFiles = e.originalEvent.dataTransfer.files;
      });

    $form.on("submit", function (e) {
      if ($form.hasClass("is-uploading")) return false;
      $form.addClass("is-uploading").removeClass("is-error");
    });
  }
}

function readURL(input) {
  console.log(input);
  if (input && input[0]) {
    console.log("it works");
    var reader = new FileReader();

    reader.onload = function (e) {
      $("#foto").attr("src", e.target.result);
    };
    reader.readAsDataURL(input[0]);
  }

  //   if (input.files && input.files[0]) {
  //     console.log('it works')
  //     var reader = new FileReader();

  //     reader.onload = function (e) {
  //       $("#foto").attr("src", e.target.result);
  //     };
  //     reader.readAsDataURL(input.files[0]);
  //   }
}

// $("#file").change(function () {
//   readURL(this);
// });

function Show(id) {
  document.getElementById("section_Model").style.display = "none";
  document.getElementById("section_Function").style.display = "none";
  document.getElementById("section_Task").style.display = "none";
  document.getElementById("section_Note").style.display = "none";
  document.getElementById("section_Palette").style.display = "none";
  var elem = document.getElementById(id);
  elem.style.display = "block";
  elem.scrollIntoView();
}

function init() {
  dragNdrop();
  Recolor();
  $("#ProjectForm").on("submit", function (event) {
    event.preventDefault();
    pr = document.querySelectorAll("input[name=lastName]")[0].value;
    url = "/Data/"+pr;
    alert("hi")
    $.ajax({
      url: url,
      method: "PUT",
      data: $("#ProjectForm").serialize(),
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

/* 
 function deleteProyect(){
        var response = confirm("Deseas eliminar " + $(this).val() + "?")
        if (response == true){
            alert("delete")
            $.ajax({
                url: '/Crud/'+ $(this).val(),
                method: 'DELETE',
                contentType: 'application/json',
                success: function(result) {
                    alert("Eliminado Correctamente")
                    setTimeout(
                        function() 
                        {
                           location.reload();
                        }, 0001); 
                },
                error: function(request,msg,error) {
                    alert("Error al Eliminar")
                }
            });
        }
    }
*/

function addFields() {
  form = document
    .getElementById("MethodsForm")
    .getElementsByTagName("fieldset")[0];
  submit = form.querySelectorAll("input[type=submit]")[0];
  formfields = document.getElementsByClassName("modelsFields");
  form.insertBefore(formfields[0].cloneNode(true), submit);
}

function EditModel(id) {
  v = document.getElementById(id).getElementsByClassName("data")[0];
  values = v.getElementsByClassName("structureCard")[0];
  document.getElementById("modelsTitleForm").value = values.getElementsByClassName("name")[0].innerHTML;
  document.getElementById("model_id").value = values.getElementsByClassName("id")[0].innerHTML;
  form = document
    .getElementById("MethodsForm")
    .getElementsByTagName("fieldset")[0];
  formfields = document.getElementsByClassName("modelsFields");
  fields = values.getElementsByClassName("fields");
  submit = form.querySelectorAll("input[type=submit]")[0];
  console.log(submit);
  if (formfields.length > fields.length) {
    while (formfields.length > fields.length) {
      form.removeChild(formfields[0]);
    }
  }
  if (formfields.length < fields.length) {
    while (formfields.length < fields.length) {
      console.log(form.lastChild);
      form.insertBefore(formfields[0].cloneNode(true), submit);
    }
  }
  for (var i = 0; i < formfields.length; i++) {
    var data = formfields[i];
    data.children[1].children[0].value = fields[i].getElementsByClassName(
      "fieldName"
    )[0].innerHTML;
    data.children[3].children[0].setAttribute(
      "value",
      fields[i].getElementsByClassName("fieldType")[0].innerHTML
    );
  }
  $("#MethodsForm").on("submit", function (event) {
    // function PutModel(event){
    event.preventDefault();
    pr = document.querySelectorAll("input[name=lastName]")[0].value;
    url = "/Crud/Modelo";
    console.log($("#MethodsForm").serialize());
    // alert("PUTTING MODEL" + url);
    $.ajax({
      url: url,
      method: "PUT",
      // data: JSON.stringify($('#MethodsForm').serializeArray()),
      data: $("#MethodsForm").serialize(),
      // contentType: 'application/json',
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
    // }
  });
}

function DeleteModel(id) {
  var response = confirm("Deseas Eliminar ?");
  if (response == true) {
    pr = document.querySelectorAll("input[name=lastName]")[0].value;
    // alert("delete at /Crud/Modelo/" + id);
    $.ajax({
      url: "/Crud/Modelo/" + id,
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

function EditFunction(id) {
  values = document.getElementById(id);
  // form = document.getElementById('FunctionsForm')
  document.getElementById("functionCallForm").value = values.getElementsByClassName("functionCall")[0].innerHTML;
  document.getElementById("functionReturnForm").value = values.getElementsByClassName("functionReturn")[0].innerHTML;
  document.getElementById("functionDescriptionForm").value = values.getElementsByClassName("functionDescription")[0].innerHTML;
  document.getElementById("functionCodeForm").value = values.getElementsByClassName("functionCode")[0].innerHTML;
  document.getElementById("function_id").value = values.getElementsByClassName("id")[0].innerHTML;
  $("#FunctionsForm").on("submit", function (event) {
    event.preventDefault();
    pr = document.querySelectorAll("input[name=lastName]")[0].value;
    url = "/Crud/Funcion";
    $.ajax({
      url: url,
      method: "PUT",
      data: $("#FunctionsForm").serialize(),
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

function DeleteFunction(id) {
  var response = confirm("Deseas Eliminar ?");
  if (response == true) {
    pr = document.querySelectorAll("input[name=lastName]")[0].value;
    // alert("delete at /Crud/Funcion/" + id);
    $.ajax({
      url: "/Crud/Funcion/" + id,
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

function EditTask(id) {
  values = document.getElementById(id);

  document.getElementById("taskTextForm").value = values.getElementsByClassName("text")[0].innerHTML;
  document.getElementById("task_id").value = values.getElementsByClassName("id")[0].innerHTML;
  if (values.getElementsByClassName("taskImage")[0].alt == "Done") {
    document.getElementById("taskDoneForm").checked = true;
    document.getElementById("taskUndoneForm").checked = false;
  } else {
    document.getElementById("taskDoneForm").checked = false;
    document.getElementById("taskUndoneForm").checked = true;
  }
  $("#TaskForm").on("submit", function (event) {
    event.preventDefault();
    pr = document.querySelectorAll("input[name=lastName]")[0].value;
    url = "/Crud/Tarea";
    $.ajax({
      url: url,
      method: "PUT",
      data: $("#TaskForm").serialize(),
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

function DeleteTask(id) {
  var response = confirm("Deseas Eliminar ?");
  if (response == true) {
    pr = document.querySelectorAll("input[name=lastName]")[0].value;
    // alert("delete at /Crud/Modelo/" + id);
    $.ajax({
      url: "/Crud/Tarea/" + id,
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

function EditNote(id) {
  values = document.getElementById(id);
  document.getElementById("notesTitleForm").value = values.getElementsByClassName("noteTitle")[0].innerHTML;
  document.getElementById("note_id").value = values.getElementsByClassName("id")[0].innerHTML;
  document.getElementById("notesTextForm").value = values.getElementsByClassName("noteText")[0].getElementsByTagName("p")[0].innerHTML;
  $("#NotesForm").on("submit", function (event) {
    event.preventDefault();
    pr = document.querySelectorAll("input[name=lastName]")[0].value;
    url = "/Crud/Notas";
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

function DeleteNote(id) {
  var response = confirm("Deseas Eliminar ?");
  if (response == true) {
    pr = document.querySelectorAll("input[name=lastName]")[0].value;
    // alert("delete at /Crud/Modelo/" + id);
    $.ajax({
      url: "/Crud/Notas/" + id,
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

function Recolor() {
  document.getElementById("testText").style.color = document.getElementById(
    "texto_color"
  ).value;
  document.getElementById("testCode").style.color = document.getElementById(
    "codigo_color"
  ).value;
  document.getElementById("testButton").style.color = document.getElementById(
    "texto_color"
  ).value;
  document.getElementById(
    "testButton"
  ).style.backgroundColor = document.getElementById("enfasis_color").value;
  var color = document.getElementById("enfasis_color").value;
  var rgbaCol =
    "rgba(" +
    parseInt(color.slice(-6, -4), 16) +
    "," +
    parseInt(color.slice(-4, -2), 16) +
    "," +
    parseInt(color.slice(-2), 16) +
    "," +
    0.5 +
    ")";
  document.getElementById("testEnfasis").style.backgroundColor = rgbaCol;
}
