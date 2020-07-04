// https://css-tricks.com/drag-and-drop-file-uploading/

var map = {
        "Introduccion":"Card", 
        "Modelos":"Models",
        "Funciones":"Functions",
        "Notas":"Notes",
        "Imagenes":"Gallery",
        "Pendientes":"Tasks",
        "Miss":"Missing",
    }

    function Active(clicked, val){	
        for (i in map){
            var actual = document.getElementById(i);
            if (actual.className == val){
                $(actual).removeClass(val);
                document.getElementById(map[actual.id]).style.display = "none";
            }
        }

        var x = document.getElementById(map[clicked]);
        if (x == null ){
            console.log("FALLA", val);
            document.getElementById("Missing").style.display = "block";
        }else{
            document.getElementById("Missing").style.display = "none";
            $(document.getElementById(clicked)).addClass(val);
            x.style.display = "block";
            document.getElementById("sidebar").style.height = x.style.height;
        }
    }
   
    function myFunction() {
         var y = document.getElementById("div1")
         var x = document.getElementById("div2")
        if (x.style.display == "none") {
            x.style.display = "block"
           	y.style.display = "none"
        } else {
            x.style.display = "none"
           	y.style.display = "block"
        }
    }
    
    function Show(id){
        var list = document.getElementsByClassName("functionDescription")
        for (item in list){
            var x = list[item].id;
            if(x == id){
                if(list[item].style.display == "block"){
                    list[item].style.display = "none";
                }else{
                    list[item].style.display = "block";
                }
            }else{
                list[item].style.display = "none";
            }
        }
    }

    function move(fill) {
        var elem = document.getElementById("Fill");
        var width = 1;
        var id = setTimeout(setInterval(frame, 10), 500);
        function frame() {
          if (width >=  fill ) {
            clearInterval(id);
          } else {
            width++;
            elem.style.width = width + "%";
            document.getElementById("value").innerHTML = width * 1 + "%";
          }
        }
    }

    function maximize(){
        var modal = document.getElementById("myModal");
        var modalImg = document.getElementById("img01");
        var img = document.getElementById("img-1");
        var captionText = document.getElementById("caption");
        modal.style.display = "block";
        console.log(this)
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

    function deleteField(){
        var wrapper = $(document.getElementById("lenguage_container"));
        wrapper.on("click",".delete", function(e){
            e.preventDefault(); $(this).parent('div').remove(); x--;
        });
    }
    
    function deleteProyect(){
        var response = confirm("Deseas eliminar " + $(this).val() + "?")
        if (response == true){
            alert("delete")
            $.ajax({
                url: '/Data/'+ $(this).val(),
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

    function editProyect(){
        alert($(this).val())
        window.location= "Crud/"+$(this).val()
        // $.ajax({
        //     url: '/Data/'+ $(this).val(),
        //     contentType: 'application/json',
        //     success: function(result) {
        //         setTimeout(
        //             function() 
        //             {
        //                location.reload();
        //             }, 0001); 
        //     },
        //     error: function(request,msg,error) {
        //         alert("Error al Eliminar")
        //     }
        // });
    }

function something(){
        ProjectTable = document.querySelector("table")
        fetch("/Data").then(response => response.json()).then(proyectList => {
            proyectList.forEach(proyect => {
                row = document.createElement("tr")
                if (proyect.banner != null){
                    row.style.backgroundImage = "url('"+proyect.banner+"')"
                }

                
                icono = document.createElement("td")
                $(icono).addClass("icon")
                if (proyect.icon != null){
                    icono.style.backgroundImage = "url('"+proyect.icon+"')"
                }
                
                data = document.createElement("td")
                data.innerHTML = proyect.name
                
                introduccion = document.createElement("td")
                introduccion.innerHTML = proyect.introduccion
                
                progreso = document.createElement("td")
                progreso.innerHTML = proyect.progress
                
                lenguajes = document.createElement("td")
                lenguajes.innerHTML = proyect.language
                
                editbutton = document.createElement("td")
                edit = document.createElement("button")
                edit.innerHTML = "edit"
                $(edit).addClass("edit")
                edit.addEventListener("click", editProyect)
                $(edit).val(proyect.name)
                editbutton.appendChild(edit)
                
                deletebutton = document.createElement("td")
                deleteb = document.createElement("button")
                deleteb.innerHTML = "delete"
                $(deleteb).addClass("delete")
                deleteb.addEventListener("click", deleteProyect)
                $(deleteb).val(proyect.name)
                deletebutton.appendChild(deleteb)
                
                row.appendChild(icono)
                row.appendChild(data)
                row.appendChild(introduccion)
                row.appendChild(progreso)
                row.appendChild(lenguajes)
                row.appendChild(editbutton)
                row.appendChild(deletebutton)
                ProjectTable.appendChild(row)
            })
        })
}

function projectForm(value, values) {
     var x = document.getElementById("ProjectForm")
    if (value =+ true) {
        x.style.display = "block"

    } else {
        x.style.display = "none"
    }
}
    {/* https://www.w3schools.com/code/tryit.asp?filename=GFIVUL17L9WX */}