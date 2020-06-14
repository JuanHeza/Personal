    function Active(clicked){	
        document.getElementById(clicked).addClass('active')
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
    
    {/* https://www.w3schools.com/code/tryit.asp?filename=GFIVUL17L9WX */}