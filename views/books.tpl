<!DOCTYPE html>
<html data-ng-app="bookApp">
<meta name="viewport" content="width=device-width, initial-scale=1">
<link rel="stylesheet" href="css/bootstrap.min.css">
<link rel="stylesheet" href="css/bootstrap-theme.min.css">
<link rel="stylesheet" href="css/books.css">
<base href="/">
<body>
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
<script type="text/javascript" src="js/lib/angular-sanitize.js"></script>
<script type="text/javascript" src="js/app.js"></script>
<script type="text/javascript" src="js/controllers.js"></script>
<script type="text/javascript" src="js/services.js"></script>
</body>
</html>