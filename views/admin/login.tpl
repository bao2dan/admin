<!DOCTYPE html>
<html lang="zh-CN">
  <head>
    <meta charset="utf-8">
    <meta http-equiv="X-UA-Compatible" content="IE=edge">
    <meta name="viewport" content="width=device-width, initial-scale=1">
    <meta name="description" content="">
    <meta name="author" content="">
    <link rel="icon" href="/static/img/favicon.ico">

    <title>Somi admin login</title>

    <!-- Bootstrap core CSS -->
    <link href="/static/css/base/bootstrap.min.css" rel="stylesheet">

    <!-- Custom styles for this template -->
    <link href="/static/css/login.css" rel="stylesheet">

    <!-- Just for debugging purposes. Don't actually copy these 2 lines! -->
    <!--[if lt IE 9]><script src="/static/js/base/ie8-responsive-file-warning.js"></script><![endif]-->
    <script src="/static/js/base/ie-emulation-modes-warning.js"></script>

    <!-- HTML5 shim and Respond.js for IE8 support of HTML5 elements and media queries -->
    <!--[if lt IE 9]>
      <script src="http://cdn.bootcss.com/html5shiv/3.7.0/html5shiv.js"></script>
      <script src="http://cdn.bootcss.com/respond.js/1.4.2/respond.min.js"></script>
    <![endif]-->

    <style>
    .error{color: #f00;}
    </style>
  </head>

  <body>

    <div class="container">

      <form class="form-signin">
        <h2 class="form-signin-heading">Please sign in</h2>
        <input type="email" class="form-control" id="uname" placeholder="Email address" required autofocus>
        <input type="password" class="form-control" id="passwd" placeholder="Password" required>
        <div class="checkbox">
          <label>
            <input type="checkbox" value="remember-me"> Remember me
          </label>
        </div>
        <span id="resError" class="error">234234</span>
        <button class="btn btn-lg btn-primary btn-block" id="loginSubmit" type="button">Sign in</button>
      </form>

    </div> <!-- /container -->


    <!-- Bootstrap core JavaScript
    ================================================== -->
    <!-- Placed at the end of the document so the pages load faster -->
    <script src="http://cdn.bootcss.com/jquery/1.11.1/jquery.min.js"></script>
    <script src="/static/js/base/bootstrap.min.js"></script>
    <!-- IE10 viewport hack for Surface/desktop Windows 8 bug -->
    <script src="/static/js/base/ie10-viewport-bug-workaround.js"></script>
  </body>
</html>

<script type="text/javascript">
$(function(){
  $(".form-signin").on("click", "#loginSubmit", function(){
    var uname = $("#uname").val();
    var paswd = $("#passwd").val();
    if (!uname || !passwd) {
      alert("You must input the email and passwd!");
      return false;
    }

    $.ajax({
      type: "POST",
      url: "/admin/login",
      data: {uname:uname, passwd:passwd},
      dataType: "json",
      cache: false,
      timeout: 5000,
      success: function(data){
          if(data.succ){
            $('#resError').empty();
          }else{
            $('#resError').html(data.msg);
          }
      }
    });
  });
});
</script>
