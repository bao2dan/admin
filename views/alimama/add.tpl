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
                新建分类
            </small>
        </h1>
    </div><!-- /.page-header -->

    <div class="row">
        <div class="col-xs-12">
        <!-- PAGE CONTENT BEGINS -->


            <form class="form-horizontal" id="formBox" role="form">
                <div class="form-group">
                  <label class="col-sm-3 control-label no-padding-right"> 分类 </label>
                  <div class="col-sm-9">
                    <input type="text" id="fCategory" value="" placeholder="选择父分类" class="col-xs-10 col-sm-5" readonly />
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
                    <button class="btn btn-sm" id="cancel" type="button">
                      <i class="icon-undo bigger-110"></i> 取消 </button>
                  </div>
                </div>
            </form>


        <!-- PAGE CONTENT ENDS -->
        </div>
    </div><!-- /.row -->
</div><!-- /.page-content -->


<script type="text/javascript">
var treeJsonData = [
    {name: 'For Sale', type: 'folder', additionalParameters:{
      'children' : [
          {name: 'Appliances', type: 'item'},
          {name: 'Arts & Crafts', type: 'item'},
          {name: 'Clothing', type: 'item'},
          {name: 'Computers', type: 'item', data:"Computers1"},
          {name: 'Jewelry', type: 'item', data:"jewelry1"},
          {name: 'Office & Business', type: 'item'},
          {name: 'Sports & Fitness', type: 'folder', additionalParameters:{
              'children' : [
                  {name: 'Appliances', type: 'item', data:"Appliances1"}
              ]
          }}
        ]
      }} ,
    {name: 'For Sale', type: 'folder', additionalParameters:{
      'children' : [
          {name: 'Appliances', type: 'item'},
          {name: 'Arts & Crafts', type: 'item'},
          {name: 'Clothing', type: 'item'},
          {name: 'Computers', type: 'item'},
          {name: 'Jewelry', type: 'item', data:"jewelry1"},
          {name: 'Office & Business', type: 'item'},
          {name: 'Sports & Fitness', type: 'folder', additionalParameters:{
              'children' : [
                  {name: 'Appliances', type: 'item'}
              ]
          }}
        ]
      }} ,
    {name: 'For Sale', type: 'folder', additionalParameters:{
      'children' : [
          {name: 'Appliances', type: 'item'},
          {name: 'Arts & Crafts', type: 'item'},
          {name: 'Clothing', type: 'item'},
          {name: 'Computers', type: 'item'},
          {name: 'Jewelry', type: 'item', data:"jewelry1"},
          {name: 'Office & Business', type: 'item'},
          {name: 'Sports & Fitness', type: 'folder', additionalParameters:{
              'children' : [
                  {name: 'Appliances', type: 'item'}
              ]
          }}
        ]
      }} ,
    {name: 'Real Estate', type: 'folder', data:"Estate1"} ,
    {name: 'Pets', type: 'folder', data:"Pets1"} ,
    {name: 'Tickets', type: 'item', data:"Tickets1"} ,
    {name: 'Services', type: 'item', data:"services1"} ,
    {name: 'Personals', type: 'item', data:"personals1"}
];


$(function() {
    //选择父分类
    $('#formBox').on('click', '#fCategory', function(){
      //分类树初始化
      var treeDataSource = new DataSourceTree({data: treeJsonData});
      var treeHtml = $("<div>").attr({"id":"treeCategory", "class":"tree"}).css({"width":"100%", "height":"400px", "overflow":"auto"});
      $(treeHtml).ace_tree({
          'dataSource': treeDataSource,
          'multiSelect': false,
          'loadingHTML': '<div class="tree-loading"><i class="icon-refresh icon-spin blue"></i></div>',
          'open-icon': 'icon-minus',
          'close-icon': 'icon-plus',
          'selectable': true,
          'selected-icon': 'icon-ok',
          'unselected-icon': 'icon-remove'
        });

      $(treeHtml).on('selected', function (evt, data) {
          var obj = $(this).tree('selectedItems');
          var catId = obj[0].data;
          console.log(catId)
      });


      bootbox.dialog({
        message: treeHtml,
        buttons: {
          "button": {
            "label" : "取消",
            "className" : "btn-sm"
          }
        }
      });


        
        /*bootbox.confirm(treeHtml, function(result) {
            if(result) {
                
            }
        });*/
    })

    //取消
    $("#formBox").on("click", "#cancel", function(){
        location.href = "/category/list";
    });

    //提交
    $("#formBox").on("click", "#submit", function(){
        var account = $("#account").val();
        var passwd = $("#passwd").val();
        var name = $("#name").val();
        var phone = $("#phone").val();
        var email = $("#email").val();
        var sex = $("#formBox input[name='sex']:checked").val();
        var role = [];
        $("#formBox input[name='role']:checked").each(function(k, v) {
              role.push($(v).val());
        });

        if ("" == account || !Somi.isEmail(account)) {
          $('#errorMsg').text('账号有误');
          return false;
        }

        if ("" != passwd && !Somi.isPasswd(passwd)) {
          $('#errorMsg').text('密码有误');
          return false;
        }

        if ("" == name || name.length < 2 || name.length > 5) {
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

        var url = "/admin/update"
        var data = {
          "account": account,
          "passwd": passwd,
          "name": name,
          "phone": phone,
          "email": email,
          "sex": sex,
          "role": role.join(","),
        }
        Somi.ajax(url, data, function(data){
            if(data.succ){
              $('#errorMsg').text("修改成功");
            } else {
              $('#errorMsg').text(data.msg);
            }
        });
    });
});
</script>