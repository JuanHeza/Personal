{{define "Edit"}}
<head>
  <title>Juan Heza{{if .One}}/{{.One.Titulo}}{{end}}</title>
</head>
{{template "header" .}}
<script>
  $(document).ready(function() {
  $("#tags").select2({
  tags: true,
  placeholder: 'type now...',
      data: [{{range Lenguajes}}
      // {  id: '{{.ID}}', text: '{{.Titulo}}' }, //Este envia numeros
      {  id: '{{.Titulo}}', text: '{{.Titulo}}' },//Este envia nombres
      {{end}}],
      tokenSeparators: ['|'],
  createTag: function (params) {
    return {
      id: params.term,
      text: params.term,
      newOption: true
    }
  },
   templateResult: function (data) {
    var $result = $("<span></span>");

    $result.text(data.text);

    if (data.newOption) {
      $result.append(" <em>(new)</em>");
    }

    return $result;
  }
}).on('change', function() {
      console.log($(this).val());
    });
    });
</script>
{{if .Data.One }}
  <body onload="init()">
{{else}}
  <body onload="something()">
    <button onclick="projectForm(true)">Add Project</button>
    <table> 
      <caption>Projectos</caption>
      <tr>
        <th id="tableIcon">Icono</th>
        <th id="tableNime">Nombre</th>
        <th id="tableDescription">Introduccion</th>
        <th id="tableProgress">Progreso</th>
        <th id="tableLenguage">Lenguaje</th>
        <th id="tableEdit">Editar</th>
        <th id="TableDelete">Eliminar</th>
      </tr>
      {{range .Data.Many}}
      <tr>
        <td>Icono</td>
        <td>{{.Titulo}}</td>
        <td>{{.Detalle}}</td>
        <td>{{.Progreso}}</td>
        <td>Join .Lenguajes</td>
        <td>Editar</td>
        <td>Eliminar</td>
      </tr>
      {{end}}
    </table>
    <br />
{{end}}
    {{template "projectForm" .Data.One}}

{{if .Data.Many}}
    {{template "StaticInfo"}}

    {{template "PostEditor"}}
{{end}}
    {{template "footer"}}
   
  </body>

</html>
{{end}}