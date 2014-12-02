<div class="breadcrumbs" id="breadcrumbs">
    <script type="text/javascript">
        try {ace.settings.check('breadcrumbs', 'fixed');} catch (e) {}
    </script>
    <ul class="breadcrumb">
        <li>
            <i class="icon-home home-icon"></i>管理员
        </li>
        <li class="active">编辑</li>
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
                <div class="form-group">
                  <label class="col-sm-3 control-label no-padding-right" for="account"> 账号 </label>
                  <div class="col-sm-9">
                    <input type="text" id="account" value="[[.Info.account]]" placeholder="这里显示的是账号" class="col-xs-10 col-sm-5" readonly />
                  </div>
                </div>

                <div class="space-4"></div>
                <div class="form-group">
                  <label class="col-sm-3 control-label no-padding-right" for="passwd"> 密码 </label>
                  <div class="col-sm-9">
                    <input type="password" id="passwd" value="" placeholder="若不修改，则必须为空" class="col-xs-10 col-sm-5" />
                  </div>
                </div>

                <div class="space-4"></div>
                <div class="form-group">
                  <label class="col-sm-3 control-label no-padding-right" for="name"> 姓名 </label>
                  <div class="col-sm-9">
                    <input type="text" id="name" value="[[.Info.name]]" placeholder="请输入姓名" class="col-xs-10 col-sm-5" />
                  </div>
                </div>

                <div class="space-4"></div>
                <div class="form-group">
                  <label class="col-sm-3 control-label no-padding-right" for="phone"> 手机 </label>
                  <div class="col-sm-9">
                    <input type="text" id="phone" value="[[.Info.phone]]" placeholder="请输入手机号码" class="col-xs-10 col-sm-5" />
                  </div>
                </div>

                <div class="space-4"></div>
                <div class="form-group">
                  <label class="col-sm-3 control-label no-padding-right" for="email"> 邮箱 </label>
                  <div class="col-sm-9">
                    <input type="text" id="email" value="[[.Info.email]]" placeholder="请输入邮箱" class="col-xs-10 col-sm-5" />
                  </div>
                </div>

                <div class="space-4"></div>
                <div class="form-group">
                  <label class="col-sm-3 control-label no-padding-right"> 性别 </label>
                  <div class="col-sm-9">
                      <label>
                        <input name="sex" value="1" type="radio" class="ace" [[if .Info.sex]][[if eq .Info.sex "1"]] checked [[end]][[end]] />
                        <span class="lbl">男</span>
                      </label>
                      <label>
                        <input name="sex" value="0" type="radio" class="ace" [[if .Info.sex]][[if eq .Info.sex "0"]] checked [[end]] [[else]] checked [[end]]/>
                        <span class="lbl">女</span>
                      </label>
                  </div>
                </div>

                [[if .IsAdmin]]
                <div class="space-4"></div>
                <div class="form-group">
                  <label class="col-sm-3 control-label no-padding-right"> 权限 </label>
                  <div class="col-sm-9">
                    [[.RoleHtml]]
                  </div>
                </div>
                [[end]]

                <!--下拉框-->
                <div class="form-group" style="display:none;">
                  <label class="col-sm-3 control-label no-padding-right" for="role_bak"> 权限 </label>
                  <div class="col-sm-9">
                    <select class="col-xs-10 col-sm-5" id="role_bak" data-placeholder="请选择权限">
                      <option value="guest">游客</option>
                      <option value="admin1">一级管理员</option>
                      <option value="admin2">二级管理员</option>
                    </select>
                  </div>
                </div>

                <div class="form-group">
                  <label class="col-sm-3 control-label no-padding-right"></label>
                  <div class="col-sm-9">
                    <span class="text-danger" id="errorMsg">&nbsp;</span>
                  </div>
                </div>

                <div class="clearfix form-actions">
                  <div class="col-md-offset-3 col-md-9">
                    <button class="btn btn-sm btn-success" id="submit" type="button">
                      <i class="icon-ok bigger-110"></i> 确定 </button>
                    &nbsp; &nbsp; &nbsp;
                    [[if .IsAdmin]]
                      <button class="btn btn-sm" id="callback" type="button">
                        <i class="icon-undo bigger-110"></i> 返回 </button>
                    [[end]]
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
    $("#formBox").on("click", "#callback", function(){
        location.href = "/admin/list";
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
              $('#errorMsg').text(data.msg).css({"color":"#00cc66"});
            } else {
              $('#errorMsg').text(data.msg).css({"color":"#f00"});
            }
        });
    });
});
</script>