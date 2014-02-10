<!DOCTYPE html>

<html>
  	<head>
    	<title>Beego</title>
    	<meta http-equiv="Content-Type" content="text/html; charset=utf-8">

<!-- 最新 Bootstrap 核心 CSS 文件 -->
<link rel="stylesheet" href="static/css/bootstrap.css">

<!-- jQuery文件。务必在bootstrap.min.js 之前引入 -->
<script src="static/js/jquery-1.10.2.min.js"></script>

<!-- 最新的 Bootstrap 核心 JavaScript 文件 -->
<script src="static/js/bootstrap.js"></script>

		<style type="text/css">

		</style>
	</head>
  	
  	<body>
<div><p>{{.verror}}</p></div>
<form role="form" action="/register?wid={{.Wid}}" method="POST">
  <div class="form-group">
    <label for="HrCode">HR代码</label>
    <input type="text" class="form-control" id="HrCode" name="HrCode" placeholder="HR代码">
  </div>
  <div class="form-group">
    <label for="Name">姓名</label>
    <input type="text" class="form-control" id="Name" name="Name" placeholder="姓名">
  </div>

  <div class="form-group">
    <label for="Tail">身份证后六位</label>
    <input type="text" class="form-control" id="Tail"  name="Tail" placeholder="身份证后六位">
  </div>	

  <button type="submit" class="btn btn-default">确认</button>
</form>
	</body>
</html>
