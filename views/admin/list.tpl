<div class="page-header">
    <h1>
        <small>
            <i class="icon-double-angle-right"></i>
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
                    <i class="icon-hand-right icon-animated-hand-pointer blue"></i>
                    Results for "Latest Registered Domains"
                </div>

                <div class="table-responsive">
                    <table id="table_admin_list" class="table table-striped table-bordered table-hover">
                        <thead>
                            <tr>
                                <th class="center">
                                    <label>
                                        <input type="checkbox" class="ace" />
                                        <span class="lbl"></span>
                                    </label>
                                </th>
                                <th>Account</th>
                                <th>Role</th>
                                <th class="hidden-480">Email</th>
                                <th class="hidden-480">Create Time</th>
                                <th class="hidden-480">Last Update Time</th>
                                <th>Last Login Time</th>
                                <th>status</th>
                                <th></th>
                            </tr>
                        </thead>

                        <tbody>

                            <tr>
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
                    {"bSortable": false},
                    {"bSortable": false},
                    null, 
                    {"bSortable": false}, 
                    {'bSearchable':false}, 
                    null, null, null,
                    {"bSortable": false}
                ];
            config.bProcessing = true;
            config.bServerSide = true;
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
            $('#table_admin_list').dataTable(config);
        },

        //format admin list(ajax)
        formatAdminList: function() {
            var liHtml = "";
            var liStr = '<tr>\
                            <td class="center">\
                                <label>\
                                    <input type="checkbox" class="ace" />\
                                    <span class="lbl"></span>\
                                </label>\
                            </td>\
                            <td>{account}</td>\
                            <td>{role}</td>\
                            <td class="hidden-480">{email}</td>\
                            <td class="hidden-480">{create_time}</td>\
                            <td class="hidden-480">{update_time}</td>\
                            <td>{login_time}</td>\
                            <td>\
                                <span class="label label-sm {status_class}">{status}</span>\
                            </td>\
                            <td>\
                                <div class="visible-md visible-lg hidden-sm hidden-xs action-buttons">\
                                    <a class="blue" href="#">\
                                        <i class="{btn_class} bigger-130"></i>\
                                    </a>\
                                    <a class="green" href="#">\
                                        <i class="icon-pencil bigger-130"></i>\
                                    </a>\
                                    <a class="red" href="#">\
                                        <i class="icon-trash bigger-130"></i>\
                                    </a>\
                                </div>\
                            </td>\
                        </tr>';

            for (i = 0; i < data.list.length; i++) {
                var account = data.list[i].account;
                var role = data.list[i].role;
                var email = data.list[i].email;
                var lock = data.list[i].lock;
                var create_time = data.list[i].create_time;
                var update_time = data.list[i].update_time;
                var login_time = data.list[i].login_time;
                var status = "lock", status_class = "label-warning", btn_class = "icon-lock";
                if ("0" === lock) {
                    status = "unlock";
                    status_class = "label-success";
                    btn_class = "icon-unlock";
                }
                liHtml += liStr.format({account: account, role: role, email: email, status_class: status_class, status: status, btn_class: btn_class, create_time: create_time, update_time: update_time, login_time: login_time});
            }
            $('#table_admin_list tbody').html(liHtml);

        },

        //get admin list(ajax)
        getAdminList2: function() {
            var liHtml = "";
            var liStr = '<tr>\
                            <td class="center">\
                                <label>\
                                    <input type="checkbox" class="ace" />\
                                    <span class="lbl"></span>\
                                </label>\
                            </td>\
                            <td>{account}</td>\
                            <td>{role}</td>\
                            <td class="hidden-480">{email}</td>\
                            <td class="hidden-480">{create_time}</td>\
                            <td class="hidden-480">{update_time}</td>\
                            <td>{login_time}</td>\
                            <td>\
                                <span class="label label-sm {status_class}">{status}</span>\
                            </td>\
                            <td>\
                                <div class="visible-md visible-lg hidden-sm hidden-xs action-buttons">\
                                    <a class="blue" href="#">\
                                        <i class="{btn_class} bigger-130"></i>\
                                    </a>\
                                    <a class="green" href="#">\
                                        <i class="icon-pencil bigger-130"></i>\
                                    </a>\
                                    <a class="red" href="#">\
                                        <i class="icon-trash bigger-130"></i>\
                                    </a>\
                                </div>\
                            </td>\
                        </tr>';

            $.ajax({
                type: "POST",
                url: "/admin/list?rand=" + Math.random(),
                data: {},
                dataType: "json",
                cache: false,
                timeout: 5000,
                success: function(data) {
                    if (data.succ) {
                        for (i = 0; i < data.list.length; i++) {
                            var account = data.list[i].account;
                            var role = data.list[i].role;
                            var email = data.list[i].email;
                            var lock = data.list[i].lock;
                            var create_time = data.list[i].create_time;
                            var update_time = data.list[i].update_time;
                            var login_time = data.list[i].login_time;
                            var status = "lock", status_class = "label-warning", btn_class = "icon-lock";
                            if ("0" === lock) {
                                status = "unlock";
                                status_class = "label-success";
                                btn_class = "icon-unlock";
                            }
                            liHtml += liStr.format({account: account, role: role, email: email, status_class: status_class, status: status, btn_class: btn_class, create_time: create_time, update_time: update_time, login_time: login_time});
                        }
                        $('#table_admin_list tbody').html(liHtml);
                        Somi.dataTable();
                    }
                },
                error: function() {
                    alert("网络连接超时")
                }
            });
        }
    }
    $(function() {
        //Somi.getAdminList();
        Somi.dataTable();
    });
</script>