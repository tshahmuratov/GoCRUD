<!DOCTYPE html>
<html lang="en">
<meta name="viewport" content="width=device-width, initial-scale=1">
<link rel="stylesheet" href="css/bootstrap.min.css">
<link rel="stylesheet" href="css/bootstrap-theme.min.css">
<link rel="stylesheet" href="css/books.css">
<script src="js/bootstrap.min.js"></script>
<body>

<h3 class='text-center'>Мои библиотеки</h3>
<div class="row">
	<div class="col-md-4"></div>
	<div class="col-md-4 text-center">
<select class="selectpicker hor-margin align-center">
	<option value="">Выберите библиотеку</option>
{{range .Libraries}}
    <option value="{{.Id}}">{{.Name}}</option>
{{end}}
</select>
	<button type="button" class="btn btn-success hor-margin">
					<span class="glyphicon glyphicon-plus"></span>
					Добавить библиотеку
				</button>
	</div>
	<div class="col-md-4"></div>
</div>
<table class="table table-striped table-hover force-width-80"  align="center" width="80%" >
	<thead>
		<tr>
			<th>Id</th>
			<th>Название</th>
			<th>Автор</th>
		</tr>
	</thead>
	<tbody>
{{range .Books}} 
		<tr>
			<td>{{.Id}}</td>
			<td>{{.Name}}</td>
			<td>{{.Author}}</td>
		</tr>
{{end}}
		<tr>
			<td colspan=3 class=' text-center'>
				<button type="button" class="btn btn-success text-center">
					<span class="glyphicon glyphicon-plus"></span>
					Добавить книгу в библиотеку
				</button>
			</td>
		</tr>
	</tbody>
</table>
</body>
</html>