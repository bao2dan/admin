<div class="breadcrumbs" id="breadcrumbs">
    <script type="text/javascript">
        try {ace.settings.check('breadcrumbs', 'fixed');} catch (e) {}
    </script>
    <ul class="breadcrumb">
        <li>
            <i class="icon-home home-icon"></i>基础配置
        </li>
        <li class="active">分类列表</li>
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
                        <a href="/category/add" class="btn btn-success btn-sm pull-right" role="button" style="margin:4px 10px 0px 0px;">添加一级分类</a>
                    </div>

                    <div class="table-responsive">
                        <table id="table_category_list" class="table table-striped table-bordered table-hover">
                            <thead>
                                <tr>
                                    <th style="width:16%">名称</th>
                                    <th style="width:16%">分类ID</th>
                                    <th style="width:16%">父分类ID</th>
                                    <th style="width:6%">级数</th>
                                    <th style="width:8%">排序</th>
                                    <th style="width:14%">创建时间</th>
                                    <th style="width:14%">修改时间</th>
                                    <th style="width:10%">操作</th>
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
            config.sAjaxSource = "/category/list?rand=" + Math.random();
            config.aoColumns = [
                    {"bSortable":false, "sClass":"left"},
                    {"bSortable":false},
                    {"bSortable":false},
                    {"bSortable":false},
                    {"bSearchable":false},
                    {"bSearchable":false},
                    {"bSearchable":false},
                    {"bSortable":false, "bSearchable":false},
                ];
            config.fnServerData = function(sSource, aoData, fnCallback){
                aoData.push( {} );
                Somi.ajax(sSource, aoData, function(data) {
                    fnCallback(data);
                    Obj.bindOp(); //绑定操作事件
                });
            }
            config.iDisplayLength = 100; //每页默认显示100条(目的是不分页)
            config.bLengthChange = false; //不显示选择每页显示多少条
            config.bFilter = false; //不显示搜索框
            config.bPaginate = false; //不显示分页器
            var oTable = $('#table_category_list').dataTable(config);
            Somi.dataTableSearchBind(oTable);
        },

        //绑定操作事件
        bindOp: function(){
            //删除分类
            $("#table_category_list .action-buttons").on("click", ".delBtn", function(){
                var _this = this;
                var catid = $(_this).closest(".action-buttons").attr("catid");
                var name = $.trim($(_this).closest(".action-buttons").attr("name"));
                if (!catid || !name) {
                  Somi.gritter('error', "参数不能空");
                  return false;
                }
                var url = "/category/del"

                bootbox.confirm('您确定要删除商品“<font color="red">'+name+'</font>”吗?', function(result) {
                    if(result) {
                        var data = {"catid":catid};
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

            //编辑分类
            $("#table_category_list .action-buttons").on("click", ".updateBtn", function(){
                var _this = this;
                var catid = $(_this).closest(".action-buttons").attr("catid");
                if (!catid) {
                  Somi.gritter('error', "参数不能空");
                  return false;
                }
                window.location.href = "/category/update?catid="+catid+"&rand="+Math.random();
            });

            //添加子分类
            $("#table_category_list .action-buttons").on("click", ".addBtn", function(){
                var _this = this;
                var catid = $(_this).closest(".action-buttons").attr("catid");
                var name = $(_this).closest(".action-buttons").attr("name");
                var level = $(_this).closest(".action-buttons").attr("level");
                name = $.trim(name);
                if (!catid || !name || !level) {
                  Somi.gritter('error', "参数不能空");
                  return false;
                }
                window.location.href = "/category/add?fid="+catid+"&fname="+name+"&flevel="+level+"&rand="+Math.random();
            });
        }
    }
    $(function() {
        Obj.dataTable();
    });
</script>