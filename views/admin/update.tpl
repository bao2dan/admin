<div class="page-header">
    <h1>
        <small>
            <i class="icon-hand-right icon-animated-hand-pointer blue"></i>
            管理员列表
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
                <input type="text" id="account"  placeholder="这里显示的是账号" class="col-xs-10 col-sm-5" readonly />
              </div>
            </div>

            <div class="space-4"></div>

            <div class="form-group">
              <label class="col-sm-3 control-label no-padding-right" for="passwd"> 密码 </label>
              <div class="col-sm-9">
                <input type="password" id="passwd" placeholder="请输入密码" class="col-xs-10 col-sm-5" />
              </div>
            </div>

            <div class="space-4"></div>

            <div class="form-group">
              <label class="col-sm-3 control-label no-padding-right" for="uname"> 姓名 </label>
              <div class="col-sm-9">
                <input type="text" id="uname"  placeholder="请输入姓名" class="col-xs-10 col-sm-5" />
              </div>
            </div>

            <div class="space-4"></div>

            <div class="form-group">
              <label class="col-sm-3 control-label no-padding-right" for="phone"> 手机 </label>
              <div class="col-sm-9">
                <input type="text" id="phone"  placeholder="请输入手机号码" class="col-xs-10 col-sm-5" />
              </div>
            </div>

            <div class="space-4"></div>

            <div class="form-group">
              <label class="col-sm-3 control-label no-padding-right" for="email"> 邮箱 </label>
              <div class="col-sm-9">
                <input type="text" id="email"  placeholder="请输入邮箱" class="col-xs-10 col-sm-5" />
              </div>
            </div>

            <div class="space-4"></div>

            <div class="form-group">
              <label class="col-sm-3 control-label no-padding-right"> 性别 </label>
              <div class="col-sm-9">
                  <label>
                    <input name="gender" value="1" type="radio" class="ace" checked />
                    <span class="lbl">男</span>
                  </label>
                  <label>
                    <input name="gender" value="0" type="radio" class="ace" />
                    <span class="lbl">女</span>
                  </label>
              </div>
            </div>

            <div class="space-4"></div>

            <div class="form-group">
              <label class="col-sm-3 control-label no-padding-right"> 权限 </label>
              <div class="col-sm-9">
                  <label>
                    <input name="role" value="admin1" type="checkbox" class="ace ace-checkbox-2" />
                    <span class="lbl">一级管理员</span>
                  </label>
                  <label>
                    <input name="role" value="admin2" type="checkbox" class="ace ace-checkbox-2" />
                    <span class="lbl">二级管理员</span>
                  </label>
                  <label>
                    <input name="role" value="guest" type="checkbox" class="ace ace-checkbox-2" />
                    <span class="lbl">游客</span>
                  </label>
                </div>
            </div>

            <div class="space-4"></div>

            <div class="form-group">
              <label class="col-sm-3 control-label no-padding-right"></label>
              <div class="col-sm-9">
                <span class="text-danger" id="errorMsg">&nbsp;</span>
              </div>
            </div>

            <div class="space-4"></div>

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


            <div class="clearfix form-actions">
              <div class="col-md-offset-3 col-md-9">
                <button class="btn btn-sm btn-success" id="submit" type="button">
                  <i class="icon-ok bigger-110"></i>
                  确定
                </button>

                &nbsp; &nbsp; &nbsp;
                <button class="btn btn-sm" id="cancel" type="button">
                  <i class="icon-undo bigger-110"></i>
                  取消
                </button>
              </div>
            </div>
          </form>


    <!-- PAGE CONTENT ENDS -->
    </div>
</div><!-- /.row -->


<script type="text/javascript">
$(function() {
    //下拉框样式初始化
    $("select").chosen()
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
        var uname = $("#uname").val();
        var phone = $("#phone").val();
        var email = $("#email").val();
        var gender = $("#formBox input[name='gender']:checked").val();
        var role = [];
        $("#formBox input[name='role']:checked").each(function(k, v) {
              role.push($(v).val());
        });

        if ("" == account || !Somi.isEmail(account)) {
          $('#errorMsg').text('账号有误');
          return false;
        }

        if ("" == passwd || !Somi.isPasswd(passwd)) {
          $('#errorMsg').text('密码有误');
          return false;
        }

        if ("" == uname || uname.length < 2 || uname.length > 5) {
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

        if (role.length == 0) {
          $('#errorMsg').text('请选择角色');
          return false;
        }

        var url = "/admin/update"
        var data = {
          "account": account,
          "passwd": passwd,
          "uname": uname,
          "phone": phone,
          "email": email,
          "gender": gender,
          "role": role,
        }
        Somi.ajax(url, data, function(data){
            if(data.succ){
              $('#errorMsg').text(data.msg);
            } else {
              $('#errorMsg').text(data.msg);
            }
        });
    });
});
</script>