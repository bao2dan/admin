<div class="breadcrumbs" id="breadcrumbs">
    <script type="text/javascript">
        try {ace.settings.check('breadcrumbs', 'fixed');} catch (e) {}
    </script>
    <ul class="breadcrumb">
        <li>
            <i class="icon-home home-icon"></i>管理员
        </li>
        <li class="active">管理员列表</li>
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


            <div class="row">
                <div class="col-xs-12">
                    <div class="table-header white">
                        <i class="icon-exclamation-sign"></i>
                        只列出了部分信息
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
                                    <th>账号</th>
                                    <th>角色</th>
                                    <th>姓名</th>
                                    <th>手机号</th>
                                    <th>创建时间</th>
                                    <th>登陆时间</th>
                                    <th>状态</th>
                                    <th>操作</th>
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
</div><!-- /.page-content -->


<script type="text/javascript">
    var Obj = {
        //data table
        dataTable: function() {
            var config = DataTableConfig;
            config.sAjaxSource = "/admin/list?rand=" + Math.random();
            config.aoColumns = [
                    {"bSortable":false, "bSearchable":false},
                    {"bSortable":false},
                    null,
                    {"bSortable":false},
                    {"bSortable":false},
                    {"bSearchable":false},
                    {"bSearchable":false},
                    {"bSearchable":false},
                    {"bSortable":false, "bSearchable":false}
                ];
            config.fnServerData = function(sSource, aoData, fnCallback){
                aoData.push( { "name": "bao2dan", "value": "hahahaha" } );
                Somi.ajax(sSource, aoData, function(data) {
                    fnCallback(data);
                    Obj.bindOp(); //绑定操作事件
                });
            }
            var oTable = $('#table_admin_list').dataTable(config);
            Somi.dataTableSearchBind(oTable);
        },

        //绑定操作事件
        bindOp: function(){
            //锁定、解锁
            $("#table_admin_list .action-buttons").on("click", ".unLockBtn", function(){
                var _this = this;
                var account = $(_this).closest(".action-buttons").attr("account");
                if (!account) {
                  Somi.gritter('error', "参数不能空");
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
                        var data = {"account":account};
                        Somi.ajax(url, data, function(data){
                            if(data.succ){
                              if(op == "lock"){
                                $(_this).find("i").removeClass("icon-unlock").addClass("icon-lock");
                                $(_this).closest("tr").find(".status").removeClass("label-success").addClass("label-warning").html("已锁定");
                              }else{
                                $(_this).find("i").removeClass("icon-lock").addClass("icon-unlock");
                                $(_this).closest("tr").find(".status").removeClass("label-warning").addClass("label-success").html("已激活");
                              }
                            }else{
                              Somi.gritter('error', data.msg);
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
                  Somi.gritter('error', "参数不能空");
                  return false;
                }
                var url = "/admin/del"

                bootbox.confirm("您确定要删除吗?", function(result) {
                    if(result) {
                        var data = {"account":account};
                        Somi.ajax(url, data, function(data){
                            if(data.succ){
                              $(_this).closest("tr").remove();
                              Somi.gritter('success', "删除成功");
                            }else{
                              Somi.gritter('error', data.msg);
                            }
                        });
                    }
                });
            });

            //编辑账号
            $("#table_admin_list .action-buttons").on("click", ".updateBtn", function(){
                var _this = this;
                var account = $(_this).closest(".action-buttons").attr("account");
                if (!account) {
                  Somi.gritter('error', "参数不能空");
                  return false;
                }
                window.location.href = "/admin/update?account="+account+"&rand="+Math.random();
            });
        }
    }
    $(function() {
        Obj.dataTable();
    });
</script>