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

    <!-- HTML5 shim and Respond.js for IE8 support of HTML5 elements and media queries -->
    <!--[if lt IE 9]>
      <script src="http://cdn.bootcss.com/html5shiv/3.7.0/html5shiv.js"></script>
      <script src="http://cdn.bootcss.com/respond.js/1.4.2/respond.min.js"></script>
    <![endif]-->
  </head>

  <body>

    <div class="container-fluid">

      <form class="form-signin">
        <h2 class="form-signin-heading">SOMI管理系统</h2>
        <input type="email" class="form-control" id="account" value="584143515@qq.com" placeholder="Email address" required autofocus>
        <input type="password" class="form-control" id="passwd" value="123123123" placeholder="Password" required>
        <div class="checkbox">
          <label>
            <input type="checkbox" value="remember-me"> Remember me
          </label>
        </div>
        <div id="resError" class="error"></div>
        <button class="btn btn-lg btn-primary btn-block" id="loginSubmit" type="button">登陆</button>
      </form>

    </div> <!-- /container -->


    <!-- Bootstrap core JavaScript
    ================================================== -->
    <!-- Placed at the end of the document so the pages load faster -->
    <script src="/static/js/base/jquery-2.0.3.min.js"></script>
    <script src="/static/js/base/bootstrap.min.js"></script>
  </body>
</html>

<script type="text/javascript">
//登陆方法
var loginFunc = function(){
  var _this = this;
  var account = $("#account").val();
  var passwd = $("#passwd").val();
  if (!account || !passwd) {
    $('#resError').html("请输入账号和密码");
    return false;
  }
  var passwdReg = /^[A-Za-z0-9_]+$/
  if(!passwdReg.test(passwd)){
    $('#resError').html("密码只能使用字母、数字、下划线");
    return false;
  }
  if (passwd.length < 8){
    $('#resError').html("密码长度至少8位");
    return false;
  }

  $(_this).addClass("disabled");

  $.ajax({
    type: "POST",
    url: "/site/login",
    data: {"account":account, "passwd":passwd},
    dataType: "json",
    cache: false,
    timeout: 5000,
    success:function(data){
      if(data.succ){
        $('#resError').html("");
        location.href = "/";
      }else{
        $('#resError').html(data.msg);
        $(_this).removeClass("disabled");
      }
    },
    error:function(){
      $('#resError').html("网络连接超时");
      $(_this).removeClass("disabled");
    }
  });
}


$(function(){
  //登陆按钮事件
  $(".form-signin").on("click", "#loginSubmit", loginFunc);

  //当input获取焦点时候按回车键提交
  $(".form-signin").keydown(function(e){
    if(e.keyCode == 13){
      $('#loginSubmit').click();
    }
    return;
  });
});
</script>
