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


          <form role="form">
            <div class="form-group">
              <label for="exampleInputEmail1">Email address</label>
              <input type="email" class="form-control" id="exampleInputEmail1" placeholder="Enter email">
            </div>
            <div class="form-group">
              <label for="exampleInputPassword1">Password</label>
              <input type="password" class="form-control" id="exampleInputPassword1" placeholder="Password">
            </div>
            <div class="form-group">
              <label for="exampleInputFile">File input</label>
              <input type="file" id="exampleInputFile">
              <p class="help-block">Example block-level help text here.</p>
            </div>
            <div class="checkbox">
              <label>
                <input type="checkbox"> Check me out
              </label>
            </div>
            <button type="submit" class="btn btn-default">Submit</button>
          </form>


        <!-- PAGE CONTENT ENDS -->
    </div>
</div><!-- /.row -->


<script type="text/javascript">
$(function() {
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