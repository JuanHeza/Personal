 select  i.*, f.llamada, f.regresa, f.detalle, f.codigo,n.titulo, n.cuerpo, t.situacion, t.tarea
  from informacion_general as i
  join funciones as f
  on nombre = f.proyecto
  join notas as n
  on nombre = n.proyecto
	join tareas as t
	on nombre = t.proyecto where nombre='Personal';
	
  select modelo, campos, datos from modelos where proyecto='Personal';
  select * from informacion_general where nombre='Personal';

alter table notas alter column titulo type text;
alter table notas alter column cuerpo type text;
alter table tareas alter column tarea type text;
alter table tareas alter column situacion type boolean;

alter table notas add column id serial PRIMARY KEY;
alter table funciones add column id serial PRIMARY KEY;
alter table tareas add column id serial PRIMARY KEY;
alter table modelos add column id_modelos serial PRIMARY KEY;

alter table notas RENAME COLUMN id TO id_nota;
alter table funciones RENAME COLUMN id TO id_funcion;
alter table tareas RENAME COLUMN id TO id_tarea;

select * from notas;
select * from modelos;
delete from modelos where datos = '{"texto"}';

delete from funciones where proyecto='Personal';
select  campos, datos, proyecto from modelos where modelo='Projects' join  
insert into modelos (modelo, proyecto) values ('Test', 'Personal');

insert into funciones(proyecto, llamada, regresa, detalle, codigo) values 
('Personal','func getWakaTime(Project string)',' []string','Obtiene el tiempo trabajado en el proyecto utilizando la aplicacion WakaTime',' '),
('Personal','func ReadJSON()',' ','Lee el archivo JSON correspondiente al proyecto en cuestion',' ');

insert into tareas(proyecto, tarea, situacion) values
('Personal','Diseño Responsivo para todos los dispositivos',false),
('Personal','Implementar la base de datos SQL para migrar todos los datos de JSON a algo mejor manejado',true);

insert into notas(proyecto, titulo, cuerpo) values
('Personal','Go Templates','Durante todo el desarrollo de este proyecto he entendido a mayor detalle como funcionan las Templates en Go y lo utiles que pueden llegar a ser para organizar el codigo en las diferentes secciones que lo componen, asi como tambien se puede usar para llamar secciones que son constantes en todas las paginas como lo seria la cabecera y el pie de pagina, sin necesidad de escribirlos en cada pagina, corriendo el riesgo de cometer variar los datos'),
('Personal','Api','Tengo pensado implementar un api para administrar algunos dispositivos IoT o proyectos que se conecten a internet a traves de un servidor, como lo serian el proyecto Arrow con IoT y el proyecto Elements con el servidor');

select * from notas where proyecto='Personal' and titulo='Api';

UPDATE informacion_general set icon='/static/images/Personal/Icon.png', banner='/static/images/Personal/Banner.png' where nombre='Personal'