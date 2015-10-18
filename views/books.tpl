<!DOCTYPE html>
<html data-ng-app="bookApp">
<meta name="viewport" content="width=device-width, initial-scale=1">
<link rel="stylesheet" href="css/bootstrap.min.css">
<link rel="stylesheet" href="css/bootstrap-theme.min.css">
<link rel="stylesheet" href="css/books.css">
<base href="/">
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
<div class="container">
	<div class="row top-buffer">
		<div class="col-xs-8 col-xs-offset-2">
			<div ui-view></div>
		</div>
	</div>
</div>
<script type="text/javascript" src="js/lib/jquery-1.11.3.min.js"></script>
<script type="text/javascript" src="js/lib/bootstrap.min.js"></script>
<script type="text/javascript" src="js/lib/angular.min.js"></script>
<script type="text/javascript" src="js/lib/angular-ui-router.min.js"></script>
<script type="text/javascript" src="js/lib/angular-resource.min.js"></script>
<script type="text/javascript" src="js/app.js"></script>
<script type="text/javascript" src="js/controllers.js"></script>
<script type="text/javascript" src="js/services.js"></script>
</body>
</html>