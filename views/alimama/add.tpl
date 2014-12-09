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
                编辑分类（编辑一级分类时，父分类ID和名称必须为默认值0/无）
            </small>
        </h1>
    </div><!-- /.page-header -->

    <div class="row">
        <div class="col-xs-12">
        <!-- PAGE CONTENT BEGINS -->


            <form class="form-horizontal" id="formBox" role="form">
                <div class="form-group">
                  <label class="col-sm-3 control-label no-padding-right"> 选择分类 </label>
                  <div class="col-sm-9">
                    <input type="text" id="fCategory" value="" title="点击选择分类" placeholder="点击选择分类" class="col-xs-10 col-sm-5" style="background-color:ghostwhite; cursor:pointer;" />
                  </div>
                </div>

                <div class="form-group">
                  <label class="col-sm-3 control-label no-padding-right"> 商品名称 </label>
                  <div class="col-sm-9">
                    <input type="text" id="name" value=""  placeholder="" class="col-xs-10 col-sm-5" />
                  </div>
                </div>

                <div class="space-4"></div>
                <div class="form-group">
                  <label class="col-sm-3 control-label no-padding-right"> 商品原价 </label>
                  <div class="col-sm-9">
                    <input type="text" id="old_price" value="" placeholder="" class="col-xs-10 col-sm-5" />
                  </div>
                </div>

                <div class="space-4"></div>
                <div class="form-group">
                  <label class="col-sm-3 control-label no-padding-right"> 商品现价 </label>
                  <div class="col-sm-9">
                    <input type="text" id="price" value="" class="col-xs-10 col-sm-5" />
                  </div>
                </div>

                <div class="space-4"></div>
                <div class="form-group">
                  <label class="col-sm-3 control-label no-padding-right"> 商品排序 </label>
                  <div class="col-sm-9">
                    <input type="text" id="sort" value="" class="col-xs-10 col-sm-5" />
                  </div>
                </div>

                <!--下拉框-->
                <div class="form-group">
                  <label class="col-sm-3 control-label no-padding-right"> 商品状态 </label>
                  <div class="col-sm-9">
                    <label>
                        <input name="status" value="1" type="radio" class="ace" checked />
                        <span class="lbl">下线</span>
                      </label>
                      <label>
                        <input name="status" value="0" type="radio" class="ace" />
                        <span class="lbl">上线</span>
                      </label>
                  </div>
                </div>

                <div class="space-4"></div>
                <div class="form-group">
                  <label class="col-sm-3 control-label no-padding-right"> 商品图片 </label>
                  <div class="col-sm-9">
                    <input type="text" id="img" value="" placeholder="" class="col-xs-10 col-sm-5" />
                  </div>
                </div>

                <div class="space-4"></div>
                <div class="form-group">
                  <label class="col-sm-3 control-label no-padding-right"> 商品链接 </label>
                  <div class="col-sm-9">
                    <input type="text" id="url" value="" class="col-xs-10 col-sm-5" />
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
var treeJsonData = $.parseJSON('[[.CategoryTree]]');

$(function() {
    //下拉框样式初始化
    $("#formBox select").chosen()
    .next('.chosen-container').each(function(){
      $(this).addClass("col-xs-10 col-sm-5").css({}).css({"padding":"0px"});
      $(this).find('.chosen-drop').css({});
      $(this).find('.chosen-search input').css({});
    });

    //父分类选择初始化
    var treeHtml = Somi.categoryTreeInit('#fCategory', treeJsonData, function(fid, fname, flevel){
        $("#fid").val(fid);
        $("#fCategory").val(fname);
        $("#flevel").val(flevel);
        var level = parseInt(flevel) + 1;
        $("#level").val(level);
    });

    //取消
    $("#formBox").on("click", "#callback", function(){
        location.href = "/category/list";
    });

    //提交
    $("#formBox").on("click", "#submit", function(){
        var fid = $("#fid").val();
        var catid = $("#catid").val();
        var name = $("#name").val();
        var sort = $("#sort").val();

        $('#errorMsg').text('请求中......');

        if ("" == fid) {
          $('#errorMsg').text('父分类ID有误');
          return false;
        }

        if ("" == catid) {
          $('#errorMsg').text('分类ID有误');
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

        $('#submit').attr('disabled', 'disabled');

        var url = "/category/update"
        var data = {
          "fid": fid,
          "catid": catid,
          "name": name,
          "sort": sort
        }
        Somi.ajax(url, data, function(data){
            if(data.succ){
              $('#errorMsg').text(data.msg).css({"color":"#00cc66"});
            } else {
              $('#errorMsg').text(data.msg).css({"color":"#f00"});
              $('#submit').removeAttr('disabled');
            }
        });
    });
});
</script>