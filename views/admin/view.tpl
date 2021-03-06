<div class="breadcrumbs" id="breadcrumbs">
    <script type="text/javascript">
        try {ace.settings.check('breadcrumbs', 'fixed');} catch (e) {}
    </script>
    <ul class="breadcrumb">
        <li class="active">
            <i class="icon-home home-icon"></i>管理员
        </li>
        <li class="active">个人资料</li>
    </ul><!-- .breadcrumb -->
</div><!-- .breadcrumbs -->

<div class="page-content">
    <div class="page-header">
        <h1>
            <small>
                <i class="icon-hand-right icon-animated-hand-pointer orange"></i>
                请谨慎操作
            </small>
        </h1>
    </div><!-- /.page-header -->

    <div class="row">
        <div class="col-xs-12">
        <!-- PAGE CONTENT BEGINS -->


          <form class="form-horizontal" id="formBox" role="form">
              <div class="space-4"></div>
              <div class="form-group">
                <label class="col-sm-3 control-label no-padding-right" for="account"> 账号 </label>
                <div class="col-sm-9">
                  <span>[[.Info.account]]</span>
                </div>
              </div>

              <div class="space-4"></div>
              <div class="form-group">
                <label class="col-sm-3 control-label no-padding-right" for="name"> 姓名 </label>
                <div class="col-sm-9">
                  <span>[[.Info.name]]</span>
                </div>
              </div>

              <div class="space-4"></div>
              <div class="form-group">
                <label class="col-sm-3 control-label no-padding-right" for="phone"> 手机 </label>
                <div class="col-sm-9">
                  <span>[[.Info.phone]]</span>
                </div>
              </div>

              <div class="space-4"></div>
              <div class="form-group">
                <label class="col-sm-3 control-label no-padding-right" for="email"> 邮箱 </label>
                <div class="col-sm-9">
                  <span>[[.Info.email]]</span>
                </div>
              </div>

              <div class="space-4"></div>
              <div class="form-group">
                <label class="col-sm-3 control-label no-padding-right"> 性别 </label>
                <div class="col-sm-9">
                      [[if .Info.sex]]
                          [[if eq .Info.sex "1"]]
                            <span class="lbl">男</span>
                          [[else]]
                            <span class="lbl">女</span>
                          [[end]]
                      [[else]]
                          <span class="lbl">未知</span>
                      [[end]]
                </div>
              </div>

              <div class="space-4"></div>
              <div class="form-group">
                <label class="col-sm-3 control-label no-padding-right"> 权限 </label>
                <div class="col-sm-9">
                    [[.RoleHtml]]
                  </div>
              </div>

              <div class="space-4"></div>
              <div class="form-group">
                <label class="col-sm-3 control-label no-padding-right" for="phone"> 注册时间 </label>
                <div class="col-sm-9">
                  <span>[[.Info.addTime]]</span>
                </div>
              </div>

              <div class="space-4"></div>
              <div class="form-group">
                <label class="col-sm-3 control-label no-padding-right" for="phone"> 最后登陆时间 </label>
                <div class="col-sm-9">
                  <span>[[.Info.loginTime]]</span>
                </div>
              </div>
          </form>


        <!-- PAGE CONTENT ENDS -->
        </div>
    </div><!-- /.row -->
</div><!-- /.page-content -->


<script type="text/javascript">
$(function() {
    //下拉框样式初始化
    $("#formBox select").chosen()
    .next('.chosen-container').each(function(){
      $(this).addClass("col-xs-10 col-sm-5").css({}).css({"padding":"0px"});
      $(this).find('.chosen-drop').css({});
      $(this).find('.chosen-search input').css({});
    });

    //取消
    $("#formBox").on("click", "#cancel", function(){
        location.href = "/";
    });

    //提交
    $("#formBox").on("click", "#submit", function(){
        var account = $("#account").val();
        var passwd = $("#passwd").val();
        var name = $("#name").val();
        var phone = $("#phone").val();
        var email = $("#email").val();
        var sex = $("#formBox input[name='sex']:checked").val();
        var role = [];
        $("#formBox input[name='role']:checked").each(function(k, v) {
              role.push($(v).val());
        });

        if ("" == account || !Somi.isEmail(account)) {
          $('#errorMsg').text('账号有误');
          return false;
        }

        if ("" != passwd && !Somi.isPasswd(passwd)) {
          $('#errorMsg').text('密码有误');
          return false;
        }

        if ("" == name || name.length < 2 || name.length > 5) {
          $('#errorMsg').text('姓名有误');
          return false;
        }

        if ("" == phone || !Somi.isPhone(phone)) {
          $('#errorMsg').text('手机号有误');
          return false;
        }

        if (!email || !Somi.isEmail(email)) {
          $('#errorMsg').text('邮箱有误');
          return false;
        }

        var url = "/admin/update"
        var data = {
          "account": account,
          "passwd": passwd,
          "name": name,
          "phone": phone,
          "email": email,
          "sex": sex,
          "role": role.join(","),
        }
        Somi.ajax(url, data, function(data){
            if(data.succ){
              $('#errorMsg').text("修改成功");
            } else {
              $('#errorMsg').text(data.msg);
            }
        });
    });
});
</script>