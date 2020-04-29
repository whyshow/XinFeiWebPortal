<!doctype html>
<html  class="x-admin-sm">
<head>
	<meta charset="UTF-8">
	<title>信飞门户管理登录</title>
    <link rel="icon" type="image/png" href="../../static/admin/images/xin.png">
    <meta http-equiv="X-UA-Compatible" content="IE=edge,chrome=1">
    <meta name="viewport" content="width=device-width,user-scalable=yes, minimum-scale=0.4, initial-scale=0.8,target-densitydpi=low-dpi" />
    <meta http-equiv="Cache-Control" content="no-siteapp" />
    <link rel="stylesheet" href="../../static/admin/css/font.css">
    <link rel="stylesheet" href="../../static/admin/css/login.css">
    <link rel="stylesheet" href="../../static/admin/css/xadmin.css">
    <!--[if lt IE 9]>
      <script src="https://cdn.staticfile.org/html5shiv/r29/html5.min.js"></script>
      <script src="https://cdn.staticfile.org/respond.js/1.4.2/respond.min.js"></script>
    <![endif]-->
</head>
<body class="login-bg">
    
    <div class="login layui-anim layui-anim-up">
        <div class="message">信飞门户管理登录</div>
        <div id="darkbannerwrap"></div>
        
        <form>
            <input id="username" placeholder="用户名" type="text" lay-verify="required" class="layui-input" >
            <hr class="hr15">
            <input id="password" lay-verify="required" placeholder="密码"  type="password" class="layui-input">
            <hr class="hr15">
            <input value="登录" onclick="login()" style="width:100%;" type="button">
            <hr class="hr20" >
        </form>
    </div>
    <footer class="footer">
        <div class="text-center">© 2019 信飞软件工程研发中心  豫ICP备19021719号-3</div>
    </footer>
    <script type="text/javascript" src="../../static/admin/js/jquery.min.js"></script>
    <script src="../../static/admin/lib/layer/layer.js"></script>
    <script>
        function login(){
            $.ajax({
                url:"/manage/request_login",
                dataType : 'json',
                type:'POST',
                data:{
                    Admin_name:$("#username").val(),
                    Admin_password:$("#password").val()
                },
                success:function (data) {
                    if (data.code ==1){
                        layer.open({
                            icon: 1,
                            content:data.message,
                            skin: 'dayer-ext-moon',
                            time:1500,
                            end:function(){
                                window.location.replace (data.result);
                            }
                        })
                    }else {
                        layer.open({
                            icon: 2,
                            content:data.message,
                            skin: 'dayer-ext-moon',
                            time:1500
                        })
                    }
                },
            })
        }
    </script>
</body>
</html>