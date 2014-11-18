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


        <form class="form-horizontal" role="form">
            <div class="form-group">
              <label class="col-sm-3 control-label no-padding-right" for="account"> 账号 </label>

              <div class="col-sm-9">
                <input type="text" id="account"  placeholder="这里显示您的账号" class="col-xs-10 col-sm-5" readonly />
              </div>
            </div>

            <div class="space-4"></div>

            <div class="form-group">
              <label class="col-sm-3 control-label no-padding-right" for="passwd"> 密码 </label>

              <div class="col-sm-9">
                <input type="password" id="passwd" placeholder="请输入密码" class="col-xs-10 col-sm-5" />
                <span class="help-inline col-xs-12 col-sm-7">
                  <span class="middle">Inline help text</span>
                </span>
              </div>
            </div>

            <div class="space-4"></div>

            <div class="form-group">
              <label class="col-sm-3 control-label no-padding-right" for="role"> 权限 </label>
              <div class="col-sm-9">
                <select class="col-xs-10 col-sm-5" id="role" data-placeholder="请选择权限">
                  <option value="guest">游客</option>
                  <option value="admin1">一级管理员</option>
                  <option value="admin2">二级管理员</option>
                </select>
              </div>
            </div>



            <div class="clearfix form-actions">
              <div class="col-md-offset-3 col-md-9">
                <button class="btn btn-sm btn-success" type="button">
                  <i class="icon-ok bigger-110"></i>
                  确定
                </button>

                &nbsp; &nbsp; &nbsp;
                <button class="btn btn-sm" type="button">
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


    //锁定、解锁
    $("#table_admin_list .action-buttons").on("click", ".unLockBtn", function(){
        var _this = this;
        var account = $(_this).closest(".action-buttons").attr("account");
        if (!account) {
          alert("参数不能空")
          return false;
        }

        var iClass = $(_this).find("i").attr("class")
        var op = "lock"
        var url = "/admin/lock"
        if(iClass.indexOf("icon-lock") >- 1){
            op = "unlock"
            url = "/admin/unlock"
        }

        bootbox.confirm("您确定要操作吗?", function(result) {
            if(result) {
                $.ajax({
                  type: "POST",
                  url: url,
                  data: {"account":account},
                  dataType: "json",
                  cache: false,
                  timeout: 5000,
                  success:function(data){
                    if(data.succ){
                      if(op == "lock"){
                        $(_this).find("i").removeClass("icon-unlock").addClass("icon-lock");
                        $(_this).closest("tr").find(".status").removeClass("label-success").addClass("label-warning").html("已锁定");
                      }else{
                        $(_this).find("i").removeClass("icon-lock").addClass("icon-unlock");
                        $(_this).closest("tr").find(".status").removeClass("label-warning").addClass("label-success").html("已激活");
                      }
                    }else{
                      alert(data.msg);
                    }
                  },
                  error:function(){
                    alert("网络连接超时");
                  }
                });
            }
        });
    });
});
</script>