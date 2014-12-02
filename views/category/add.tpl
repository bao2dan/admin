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
                <i class="icon-hand-right icon-animated-hand-pointer blue"></i>
                添加分类（添加一级分类时，父分类ID和名称必须为默认值0/无）
            </small>
        </h1>
    </div><!-- /.page-header -->

    <div class="row">
        <div class="col-xs-12">
        <!-- PAGE CONTENT BEGINS -->


            <form class="form-horizontal" id="formBox" role="form">
                <div class="form-group">
                  <label class="col-sm-3 control-label no-padding-right" for="fid"> 父分类ID </label>
                  <div class="col-sm-9">
                    <input type="text" id="fid" value="[[.Fid]]"  placeholder="" class="col-xs-10 col-sm-5" readonly />
                  </div>
                </div>

                <div class="space-4"></div>
                <div class="form-group">
                  <label class="col-sm-3 control-label no-padding-right" for="fname"> 父分类名称 </label>
                  <div class="col-sm-9">
                    <input type="text" id="fname" value="[[.Fname]]" placeholder="" class="col-xs-10 col-sm-5" readonly />
                  </div>
                </div>

                <div class="space-4"></div>
                <div class="form-group">
                  <label class="col-sm-3 control-label no-padding-right" for="flevel"> 父分类级数 </label>
                  <div class="col-sm-9">
                    <input type="text" id="flevel" value="[[.Flevel]]" placeholder="" class="col-xs-10 col-sm-5" readonly />
                  </div>
                </div>

                <div class="space-4"></div>
                <div class="form-group">
                  <label class="col-sm-3 control-label no-padding-right" for="name"> 分类名称 </label>
                  <div class="col-sm-9">
                    <input type="text" id="name" value="" placeholder="请输入分类名称" class="col-xs-10 col-sm-5" />
                  </div>
                </div>

                <div class="space-4"></div>
                <div class="form-group">
                  <label class="col-sm-3 control-label no-padding-right" for="sort"> 分类排序 </label>
                  <div class="col-sm-9">
                    <input type="text" id="sort" value="0" placeholder="请输入数字(正序)" class="col-xs-10 col-sm-5" />
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
                      <button class="btn btn-sm" id="callback" type="button">
                        <i class="icon-undo bigger-110"></i> 返回 </button>
                  </div>
                </div>
            </form>


        <!-- PAGE CONTENT ENDS -->
        </div>
    </div><!-- /.row -->
</div><!-- /.page-content -->


<script type="text/javascript">
$(function() {
    //取消
    $("#formBox").on("click", "#callback", function(){
        location.href = "/category/list";
    });

    //提交
    $("#formBox").on("click", "#submit", function(){
        var fid = $("#fid").val();
        var flevel = $("#flevel").val();
        var name = $("#name").val();
        var sort = $("#sort").val();

        if ("" === fid) {
          $('#errorMsg').text('父分类ID有误');
          return false;
        }

        if ("" === flevel || isNaN(flevel)) {
          $('#errorMsg').text('级数有误');
          return false;
        }

        if ("" == name) {
          $('#errorMsg').text('名称有误');
          return false;
        }

        if ("" == sort || isNaN(sort)) {
          $('#errorMsg').text('排序有误');
          return false;
        }

        var url = "/category/add"
        var data = {
          "fid": fid,
          "flevel": flevel,
          "name": name,
          "sort": sort
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