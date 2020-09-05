# golang-spa-demo
这个简单的程序演示了一个单页应用的主体结构和前后端交互过程。
后端使用golang编写，只依赖于标准库的http和json库，提供一个非常简单的/api/add接口，将输入的两个整数相加并返回结果。
golang后端同时提供对html文件的静态访问服务。
后端只有一个主程序文件cmd/main.go

前端使用html+javascript编写，两个输入框，用户点击计算以后，javascript向后端发出请求，获得结果以后更细界面元素来呈现。
前端代码都在web/index.html中

# 使用方法

cd cmd
go build main.go
./main

打开浏览器,输入http://localhost:8080/web/index.html
在页面内输入两个数字，然后点击Calculate，结果会在下面显示

# 作业
使用C#或是C++（MFC）编写一个客户端程序，通过调用golang的/api/add接口来实现一个加法程序，讨论并说明CS和BS结构的异同。