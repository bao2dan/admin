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
                <div class="table-header">
                    <i class="icon-double-angle-right"></i>
                    Results for "Latest Registered Domains"
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

                            <tr style="display:none;">
                                <td class="center">
                                    <label>
                                        <input type="checkbox" class="ace" />
                                        <span class="lbl"></span>
                                    </label>
                                </td>

                                <td>584143515@qq.com</td>
                                <td>root</td>
                                <td class="hidden-480">584143515@qq.com</td>
                                <td class="hidden-480">2014-11-10 12:12:10</td>
                                <td class="hidden-480">2014-11-11 12:12:20</td>
                                <td>2014-11-13 14:12:10</td>

                                <td>
                                    <span class="label label-sm label-warning">lock</span>
                                </td>

                                <td>
                                    <div class="visible-md visible-lg hidden-sm hidden-xs action-buttons">
                                        <a class="blue" href="#">
                                            <i class="icon-lock bigger-130"></i>
                                        </a>
                                        <a class="green" href="#">
                                            <i class="icon-pencil bigger-130"></i>
                                        </a>
                                        <a class="red" href="#">
                                            <i class="icon-trash bigger-130"></i>
                                        </a>
                                    </div>
                                </td>
                            </tr>

                        </tbody>
                    </table>
                </div>
            </div>
        </div>


        <!-- PAGE CONTENT ENDS -->
    </div>
</div><!-- /.row -->

<!-- page specific plugin scripts -->
<script src="/static/js/base/jquery.dataTables.min.js"></script>
<script src="/static/js/base/jquery.dataTables.bootstrap.js"></script>

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
                    success: fnCallback
                });
            }
            var oTable = $('#table_admin_list').dataTable(config);
            DataTableSearchBind(oTable);
        }
    }
    $(function() {
        //Somi.getAdminList();
        Somi.dataTable();
    });
</script>