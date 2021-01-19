 SELECT  i.*, f.llamada, f.regresa, f.detalle, f.codigo,n.titulo, n.cuerpo, t.situacion, t.tarea
  FROM informacion_general AS i
  JOIN funciones AS f
  ON nombre = f.proyecto
  JOIN notas AS n
  ON nombre = n.proyecto
	JOIN tareas AS t
	on nombre = t.proyecto WHERE nombre='Personal';
	
  SELECT modelo, campos, datos FROM modelos WHERE proyecto='Personal';
  SELECT * FROM informacion_general WHERE nombre='Personal';

ALTER TABLE notas ALTER COLUMN titulo TYPE TEXT;
ALTER TABLE notas ALTER COLUMN cuerpo TYPE TEXT;
ALTER TABLE tareas ALTER COLUMN tarea TYPE TEXT;
ALTER TABLE tareas ALTER COLUMN situacion TYPE BOOLEAN;

ALTER TABLE notas ADD COLUMN id SERIAL PRIMARY KEY;
ALTER TABLE funciones ADD COLUMN id SERIAL PRIMARY KEY;
ALTER TABLE tareas ADD COLUMN id SERIAL PRIMARY KEY;
ALTER TABLE modelos ADD COLUMN id_modelos SERIAL PRIMARY KEY;

ALTER TABLE notas RENAME COLUMN id TO id_nota;
ALTER TABLE funciones RENAME COLUMN id TO id_funcion;
ALTER TABLE tareas RENAME COLUMN id TO id_tarea;

SELECT * FROM notas;
SELECT * FROM modelos;
delete FROM modelos WHERE datos = '{"TEXTo"}';

  CREATE TABLE usuarios(
    id TEXT,
    password TEXT,
    admin BOOLEAN,
    PRIMARY KEY(id)
  );
  DROP TABLE galeria;
  CREATE TABLE galeria(
    id TEXT,
    proyect TEXT,
    dir TEXT,
    caption TEXT,
    PRIMARY KEY(id)
  );
  CREATE TABLE statics(
    introduccion TEXT,
    about TEXT,
    tutorial TEXT,
    contacto TEXT,
    link TEXT[][], /* array de [link, link image]*/
    avatar TEXT,
    banner TEXT
  );
  INSERT INTO statics(introduccion, about, tutorial, contacto, leng) VALUES(
    'Este sitio funciona como portafolio donde se pueden encontrar los proyectos que tengo completados, en desarrollo y en planeacion, principalmente de software pero tambien se pueden encontrar algo de electronica, IoT y algunas cosas mas',
    'Juan Enrique un desarrollador de software mexicano, siendo "Heza" una abreviacion de sus apellidos, con gusto por los retos ya que es una forma muy interesante de aprender cosas nuevas y usarlas en el momento, esta especializado en el lenguaje Go y enfocado en aprender el desarrollo web en el ambito de Backend',
    'Aqui puedes encontrar la lista completa de los proyectos, asi como tambien tutoriales y guias de como se hicieron algunas cosas en los diferentes proyectos.',
    'Informacion de contacto', 'es'
  );
  INSERT INTO statics(introduccion, about, tutorial, contacto, leng) VALUES(
    'This site works as a portafolio where you can found the proyects i have completed, in development or planning, mainly of software but also you can found some of electronic, IoT and some other things',
    'Juan Enrique is a mexican software developer, "Heza" its just an abbreviation of his last names, he likes the challenges since it`s a very interesting way to learn new things and use them at the moment, he is specialized in the Go language and focused on learning web development in the field of Backend',
    'You can found a complete list of proyects, tutorials and guides of how to do some things i have use in the proyects',
    'Contact Information', 'en'
  );
  UPDATE statics set link='{
    {"https://repl.it/@JuanHeza/","/static/stylesheets/Repl.it.png"},
    {"https://github.com/JuanHeza","/static/stylesheets/GitHub-Mark-Light-32px.png"},
    {"mailto:juanehza@hotmail.com","/static/stylesheets/envelope.png"}
    }' WHERE leng = 'en';
  UPDATE statics set link=ARRAY[
    ['https://repl.it/@JuanHeza/','/static/stylesheets/Repl.it.png'],
    ['https://github.com/JuanHeza','/static/stylesheets/GitHub-Mark-Light-32px.png'],
    ['mailto:juanehza@hotmail.com','/static/stylesheets/envelope.png']
    ] WHERE leng ='es';

ALTER TABLE links ADD COLUMN id SERIAL PRIMARY KEY;
DROP TABLE links;
CREATE TABLE links (
  link text,
  icon text,
  id SERIAL,
  PRIMARY KEY(id)
);
INSERT  INTO  links VALUES ('TEST','/static/stylesheets/config.png');
SELECT * FROM links ;
UPDATE links set icon = '/static/images/lisa' where id = 4;
    UPDATE statics SET avatar='/static/stylesheets/avatar.png',banner='/static/stylesheets/banner.png' WHERE leng IN ('es','en')
  DELETE FROM statics WHERE leng='es';
SELECT * FROM statics;
UPDATE links SET icon='/static/stylesheets/GitHub-Mark-32px.png' WHERE link='https://github.com/JuanHeza';
SELECT * FROM links;
  ALTER TABLE statics ADD COLUMN leng TEXT;
  ALTER TABLE statics DROP COLUMN avatar;
  ALTER TABLE statics DROP COLUMN banner;
  INSERT INTO usuarios(id, password, admin) VALUES ('Juan Heza', 'JHZ697heza', true);
  UPDATE usuarios SET id='JuanHeza' WHERE id='Juan Heza'
  SELECT * FROM usuarios WHERE id = 'JuanHeza' AND password = 'JHZ697heza';

delete FROM funciones WHERE proyecto='Personal';
SELECT  campos, datos, proyecto FROM modelos WHERE modelo='Projects' JOIN  
INSERT INTO modelos (modelo, proyecto) VALUES ('Test', 'Personal');

INSERT INTO funciones(proyecto, llamada, regresa, detalle, codigo) VALUES 
('Personal','func getWakaTime(Project string)',' []string','Obtiene el tiempo trabajado en el proyecto utilizando la aplicacion WakaTime',' '),
('Personal','func ReadJSON()',' ','Lee el archivo JSON correspondiente al proyecto en cuestion',' ');

INSERT INTO tareas(proyecto, tarea, situacion) VALUES
('Personal','Diseño Responsivo para todos los dispositivos',false),
('Personal','Implementar la base de datos SQL para migrar todos los datos de JSON a algo mejor manejado',true);

INSERT INTO notas(proyecto, titulo, cuerpo) VALUES
('Personal','Go Templates','Durante todo el desarrollo de este proyecto he entendido a mayor detalle como funcionan las Templates en Go y lo utiles que pueden llegar a ser para organizar el codigo en las diferentes secciones que lo componen, asi como tambien se puede usar para llamar secciones que son constantes en todas las paginas como lo seria la cabecera y el pie de pagina, sin necesidad de escribirlos en cada pagina, corriendo el riesgo de cometer variar los datos'),
('Personal','Api','Tengo pensado implementar un api para administrar algunos dispositivos IoT o proyectos que se conecten a internet a traves de un servidor, como lo serian el proyecto Arrow con IoT y el proyecto Elements con el servidor');

SELECT * FROM notas WHERE proyecto='Personal' and titulo='Api';

UPDATE informacion_general SET icon='/static/images/Personal/Icon.png', banner='/static/images/Personal/Banner.png' WHERE nombre='Personal'

DROP TABLE IF EXISTS  notes  CASCADE;

SELECT pr.project_id, pr.titulo, pr.detalle, lengs.titulo FROM proj_leng AS rl JOIN lenguajes AS ln ON rl.lenguaje_id = ln.lenguaje_id JOIN projects AS pr ON rl.project_id = pr.project_id JOIN proj_leng AS rels ON rels.project_id = pr.project_id JOIN lenguajes AS lengs ON lengs.lenguaje_id = rels.lenguaje_id WHERE ln.titulo = 'Go'
 SELECT  i.*, f.llamada, f.regresa, f.detalle, f.codigo,n.titulo, n.cuerpo, t.situacion, t.tarea
  FROM informacion_general AS i
  JOIN funciones AS f
  ON nombre = f.proyecto
  JOIN notas AS n
  ON nombre = n.proyecto
	JOIN tareas AS t
	on nombre = t.proyecto WHERE nombre='Personal';