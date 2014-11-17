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


        <div class="row">
            <div class="col-xs-12">
                <div class="table-header orange">
                    <i class="icon-exclamation-sign"></i>
                    请谨慎操作此列表中的功能
                </div>

                <div class="table-responsive">
                    <table id="table_admin_list" class="table table-striped table-bordered table-hover">
                        <thead>
                            <tr>
                                <th>
                                    <label>
                                        <input type="checkbox" class="ace" />
                                        <span class="lbl"></span>
                                    </label>
                                </th>
                                <th>Account</th>
                                <th>Role</th>
                                <th>Email</th>
                                <th>Create Time</th>
                                <th>Update Time</th>
                                <th>Login Time</th>
                                <th>Status</th>
                                <th>Operation</th>
                            </tr>
                        </thead>

                        <tbody>
                        </tbody>
                    </table>
                </div>
            </div>
        </div>


        <!-- PAGE CONTENT ENDS -->
    </div>
</div><!-- /.row -->


<script type="text/javascript">
    var Somi = {
        //data table
        dataTable: function() {
            var config = DataTableConfig;
            config.sAjaxSource = "/admin/list?rand=" + Math.random();
            config.aoColumns = [
                    {"bSortable":false, "bSearchable":false},
                    {"bSortable":false},
                    null,
                    {"bSortable":false,}, 
                    {"bSearchable":false}, 
                    {"bSearchable":false},
                    {"bSearchable":false},
                    {"bSearchable":false},
                    {"bSortable":false, "bSearchable":false}
                ];
            config.fnServerData = function(sSource, aoData, fnCallback){
                aoData.push( { "name": "bao2dan", "value": "hahahaha" } );
                $.ajax( {
                    type: "POST",
                    url: sSource,
                    data: aoData,
                    dataType: "json",
                    cache: false,
                    timeout: 5000,
                    success: function(data) {
                        fnCallback(data);
                        Somi.bindOp(); //绑定操作事件
                    }
                });
            }
            var oTable = $('#table_admin_list').dataTable(config);
            DataTableSearchBind(oTable);
        },

        //绑定操作事件
        bindOp: function(){
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

            //删除账号
            $("#table_admin_list .action-buttons").on("click", ".delBtn", function(){
                var _this = this;
                var account = $(_this).closest(".action-buttons").attr("account");
                if (!account) {
                  alert("参数不能空")
                  return false;
                }
                var url = "/admin/del"

                bootbox.confirm("您确定要删除吗?", function(result) {
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
                              $(_this).closest("tr").remove();
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

            //查看账号信息
            $("#table_admin_list .action-buttons").on("click", ".updateBtn", function(){
                var _this = this;
                var account = $(_this).closest(".action-buttons").attr("account");
                if (!account) {
                  alert("参数不能空");
                  return false;
                }
                window.location.href = "/admin/update?account="+account+"&rand="+Math.random();
            });
        }
    }
    $(function() {
        Somi.dataTable();
    });
</script>