
<!DOCTYPE html>
<!--<?php
session_start();
$admin = $_SESSION['ul']->admin;
if ($admin != 2) {
} else {
  header("Location: Login.php");
}
?>-->
<html lang="es"><head><meta charset="UTF-8">
	<link href='https://fonts.googleapis.com/css?family=Open+Sans' rel='stylesheet' type='text/css'>
	<meta name="viewport" content="width=device-width, user-scalable=no, initial-scale=1.0, maximum-scale=1.0, minimum-scale=1.0">
	<link rel="stylesheet" href="css/bootstrap.min.css">
	<link rel="stylesheet" href="css/estilos.css">
	<script type="text/javascript" src="js/flipclock.js"></script>
	<script src="js/JSAM.js"></script>
	<script type="text/javascript" src="http://ajax.googleapis.com/ajax/libs/jquery/1.10.2/jquery.min.js"></script>
	<title> TITULO </title>

	<head>
		<!--	<link rel="icon" href="src/Amfime.ico" type="image/x-icon" />	-->
		<link rel="stylesheet" type="text/css" href="src/slider/engine1/style.css" />
		<script type="text/javascript" src="src/slider/engine1/jquery.js"></script>
	</head>
		
	<body >
		<section class="jumbotron jumbotron-ar">
		<header>
			<center><a href="Inicio.php"><img src="src/new-logo.png" width="30%" ></a></center>
		</header>
		</section>
		<div id="sidebar">
			<ul id="content" class="nav nav-pills nav-stacked" style="max-width: 250px;">
				<li id="Inicio" class="">	 <a href="Inicio.php" onclick="Inicio()">		<span class="glyphicon glyphicon-home"></span>  Inicio</a></li>
				<li id="Quienes" class="">	 <a href="Quienes.php" onclick="Quienes()">		<span class="glyphicon glyphicon-th-list"></span>  ¿Quienes la Integran?</a></li>
				<li id="Calendario" class=""><a href="Calendario.php" onclick="Calendarios()"><span class="glyphicon glyphicon-calendar"></span>  Calendario</a></li>
				<li id="Eventos" class="">	 <a href="Evento.php" onclick="Eventos()">		<span class="glyphicon glyphicon-comment"></span>  Eventos</a></li>
				<li id="Galeria" class="">	 <a href="CrearGaleria.php" onclick="Galeria()">		<span class="glyphicon glyphicon-picture"></span>  Galeria</a></li>
				<li id="Usuario" class="">	 <a href="Usuario.php" onclick="Usuarios()">		<span class="glyphicon glyphicon-user"></span>  Usuario</a></li>
				<li id="Contacto" class="">	 <a href="Contacto.php" onclick="Contacto()">		<span class="glyphicon glyphicon-send"></span>  Contacto</a></li>
				<li id="Soporte" class="">	 <a href="Soporte.php" onclick="Soporte()">		<span class="glyphicon glyphicon-asterisk"></span> Soporte</a></li>
				<li id="Creditos" class="">	 <a href="Credito.php" onclick="Creditos()">		<span class="glyphicon glyphicon-info-sign"></span>  Creditos</a></li>
				<li id="Cerrar" class="">	 <a href="LogIn.php">							<span class="glyphicon glyphicon-off"></span> Cerrar Sesion</a></li>
			</ul>
		</div>
			
		<style type="text/css">
			#Cuerpo{
				width: 56.5%;
			}
		</style>
			
		<div id="FecHora" style="text-align:middle;width:320px;padding:0.5em 0;">
			<div><center>
					<ul class="nav nav-pills nav-stacked" style="max-width: 250px; padding: 10px;overflow: auto;">
					<li id="admin" class="active" style="display:none"><a href="Admin.php" style=" background-color: #2eb82e;"><span class="glyphicon glyphicon-star"></span> Admin</a></li>
					</ul>
				</div>
			<center><iframe src="http://www.zeitverschiebung.net/clock-widget-iframe?language=es&timezone=America%2FMonterrey" width="70%" height="80%" frameborder="0" seamless></iframe></center>
		</div>
	</h3>
	<?php
    if ($admin == 1) {
      echo "<script>
      document.getElementById('admin').style.display = 'block'; </script>";
    }
    ?>
		<!--		-->
		</body>
</html>
