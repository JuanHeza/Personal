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
function AddNote(){
  note = document.getElementById("note_block")
  clon = note.cloneNode(true)
  note.querySelector(".nameCol input[type='text']").value=""
  note.querySelector(".inputCol textarea").innerHTML=""
  $(clon).removeAttr("id")
  note.parentElement.insertBefore(clon,note)
}

function readURL(input) {
  console.log(input,"&&",input[0]);
  carta = document.getElementById("plantilla")
  clon = carta.cloneNode(true)
  $(clon).removeAttr("id")
  document.getElementById("imageBox").insertBefore(clon, carta);
  console.log(clon)
  foto = clon.querySelector("#foto")

  if (input && input[0]) {
    console.log("it works");
    var reader = new FileReader();

    reader.onload = function (e) {
      $(foto).attr("src", e.target.result);
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

function init() {
  dragNdrop();
  $("#ProjectForm").on("submit", function (event) {
    event.preventDefault();
    pr = document.querySelectorAll("input[name=project_id]")[0].value;
    url = "/Update/Project/"+pr;
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


