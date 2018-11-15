# ServiceComputing-NameSearch

## 1介绍
```
是基于cloudgo的一款在线信息查询程序，实现了在线查询功能。
您可以通过上传自己的信息让更多的人认识您，也可以通过此平台查询自己想要认识的人的信息。当然啦，知识基础信息，包括：姓名、邮箱、联系方式，如果想要保密，也可以只添部分。
该程序是在老师的博客基础上开发的，使用了他所创建的架构。连接附在下面。
```
[链接](http://blog.csdn.net/pmlpml/article/details/78404838)

-----------------------

## 2使用

### 2.1注册信息
```
介绍：当您需要在网上注册自己的信息，只需要登陆网站输入相关信息即可。
需要指令：regist
输入格式：curl -v http://localhost:9090/regist:[name]:[email]:[phone]
例如：curl -v http://localhost:9090/regist:wyj:180@163.com:122
如果成功，会有如下显示：
{
  "RESULT": "Success! Have regist!"
}
否则即为失败，即您已经注册过。
注：如哦您只想分享一部分信息，只要在不想分享的部分以empty替代即可，例如，若您不想分享你的电话，只需输入：
curl -v http://localhost:9090/regist:wyj:180@163.com:empty
```

----------------------

### 2.2更新信息

```
介绍：当您需要在网上更新自己的信息，只需要登陆网站输入相关信息即可。
需要指令：update
输入格式：curl -v http://localhost:9090/update:[name]:[email]:[phone]
例如：curl -v http://localhost:9090/check:wyj:1:1curl -v http://localhost:9090/update:wyj:122@1.com:999
如果成功，会有如下显示：
{
  "RESULT": "Update Success!"
}
否则即为失败，即您尚未注册。
注：如哦您只想更新一部分信息，只要在不想分享的部分以原有信息替代即可。
```

----------------------

### 2.3查询信息
```
介绍：当您需要在网上查询信息，只需要登陆网站输入相关信息即可。
需要指令：check
输入格式：curl -v http://localhost:9090/check:[name]:[0]:[0]
例如：curl -v http://localhost:9090/check:wyj:0:0
如果成功，会有如下显示：
{
  "RESULT": "Name: wyj Email: 1122@1.com Phone: 999"
}
否则即为失败，即该用户尚未注册，请您耐心等待！
```

---------------------

祝您使用愉快！
