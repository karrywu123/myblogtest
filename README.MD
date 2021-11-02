## Beego框架简介和程序流程分析

**@author：Davie**
**版权所有：北京千锋互联科技有限公司**

从今天开始，我们开始进入一个新的阶段的学习。从今天开始我们开始一个新的框架知识的学习--GoWeb开发之Beego实战项目。

* beego官方文档：[https://beego.me/](https://beego.me/)

#### 一、beego简介：
1.beego是一个使用Go语言来开发WEB引用的GoWeb框架，该框架起始于2012年，由一位中国的程序员编写并进行公开，其目的就是为大家提供一个高效率的web应用开发框架。该框架采用模块封装，使用简单，容易学习。方便技术开发者快速学习并进行实际开发。对程序员来说，beego掌握起来非常简单，只需要关注业务逻辑实现即可，框架自动为项目需求提供不同的模块功能。

在对beego框架有了一个简单基本的了解之后，我们给大家介绍一下beego的一些特性。

2.Beego框架的主要特性：

1）**简单化：**RESTful支持，MVC模型；可以使用bee工具来提高开发效率，比如监控代码修改进行热编译，自动化测试代码，以及自动化打包部署等丰富的开发调试功能。

2）**智能化：**beego框架封装了路由模块，支持智能路由，智能监控，并可以监控内存消耗，CPU使用以及goroutine的运行状况，方便开发者对线上应用进行监控分析。

3）**模块化：**beego根据功能对代码进行节耦封装，形成了Session，Cache，Log，配置解析，性能监控，上下文操作，ORM等独立的模块，方便开发者进行使用。

4）**高性能：**beego采用Go原生的http请求，goroutine的并发效率应付大流量的Web应用和API应用。

大家需要记住beego的几个特性，一方面在面试过程中，如果有面试官问起大家关于goweb开发的相关知识的时候，有可能问beego框架有优势或者有什么特点，这个时候需要大家能够对beego框架这些特点做一个介绍；另一方面，咱们在后面实际的学习和编码过程中，大家感受一下这些特性在实际项目中为我们带来的便利和高效。

#### 二、beego安装
在对beego的特点做了简单了解之后，我们就真正开始来搭建beego的环境，开始进入实际的编码阶段。

首先我们进行beego源码的安装。说到beego的源码安装，这里需要我们使用一个命令，这个命令之前大家应该都接触过吧？我们使用go get命令来进行beego的安装。

**【注意】：**在安装的时候，有一点大家要注意，beego框架要求Go 语言版本1.1+以上，这个我们大家的版本应该至少也是1.9的吧，应该都没有。或者大家现在在命令行中，执行go version查看一下自己的go语言版本，如下所示：
![http://7xtcwd.com1.z0.glb.clouddn.com/WX20190515-154819@2x.png](http://7xtcwd.com1.z0.glb.clouddn.com/WX20190515-154819@2x.png)

上图显示，本地环境go语言版本就是1.10.4版本。

现在，我们来执行安装命令：go get github.com/astaxie/beego
等待go将源代码下载安装完毕，我们就可以在我们自己设置的GOPATH下面的src目录下找到beego框架源码。

那么，我们首先来看一个最最简单的例子beego项目的例子，这个例子就一句话：
* 1、编写第一个beego示例项目

    ```
    package main
    import "github.com/astaxie/beego"
    func main() {
        beego.Info("第一个beego案例")
        beego.Run("localhost:8080")
    }
    ```
* 2、编译
    ```
    go build -o hello hello.go
    ```
    
* 3、执行
    ```
    ./hello
    ```
* 4、效果查询
    ```
    访问浏览器http://localhost:8080
    ```
我们在浏览器当中进行访问:[http://localhost:8080](http://localhost:8080)，可以看到，在我们的浏览器当中出现了404，代表着页面没找到；这个很好理解，因为没有指定默认的页面，所以后台项目就不知道默认指向哪个页面路径，因此就提示404。但是，看后台打出了日志，这就说明，前端的请求确实到了后台main方法里面进行执行。

### 二、命令行工具Bee
现在，尝试来进行对代码进行修改，就简单的添加一行打印日志：beego.Info("代码修改")。这个时候，如果我们要想让代码生效，必须把原来正在执行的程序停止，然后重新执行编码和运行命令，我们来实际尝试一下。

经过我们的尝试，我们可以发现，确实第二次的重新启动的程序打印出了两条后台日志，这说明我们修改的代码生效了。

现在向大家引出beego的一个工具：**bee**。

* **bee简介**  
bee是一个开发工具，是协助Beego框架开发项目时进行创建项目，运行项目，热部署等相关的项目管理的工具。beego是源码，负责开发，bee是工具，负责构建和管理项目。介绍beego的特性的时候说过其中一点beego支持代码热部署吧。所谓热部署就是，当修改代码的时候，可以不用停止服务重新启动，内置的beego就能够实时感知源代码程序编码，并进行实时生效，这就是热部署，这样就大大方便了开发人员和维护人员要频繁的进行服务停止，服务重启的过程。现在，来看看bee都有哪些功能，如何进行操作？

* **bee安装**
如同beego一样，如果要使用bee，首先要进行安装，同样的道理，先使用go get命令安装bee工具。bee安装命令：
```
go get github.com/beego/bee
```

* **bee功能命令**
安装结束以后，进入到**GOPATH**对应的目录中，然后进入到bin目录下，可以看到有bee这个可执行文件，这就是安装的bee工具的可执行文件，在该目录中，我们在终端中执行以下bee命令，然后回车，会发现列出很多关于bee的用法，如下图：


    * **new命令：**  
    
        ```
        bee new ProjectName
        ```
    该命令表示新建一个全新的web项目，有一点需要注意：该命令必须在src目录下执行，才能生效，自动生成web项目的目录结构。如果在其他目录下面执行bee new命令，也同样会是在src目录下面生成对应新项目的目录结构，这是bee工具在构建项目的时候默认寻找创建的目录。
    
    * **api命令：**
    
        ```
        bee api ProjectNames
        ```
        该命令表示用来创建开发API应用。很多用户写前端的后台接口也是通过go来写，因此bee专门提供了一个写api接口应用的命令。通过目录结构可以看到，和Web项目相比，API项目少了static和views目录，多了一个test目录，test是用来进行写测试用例代码的。
        
    * **run命令：**
    
        ```
        bee run 
        ```
        命令用来运行项目，并且能够通过监控文件系统，实时进行代码的热部署更新。也就是，我们代码逻辑进行了修改，不用停止应用，在前端页面就能看到改变。
        
    * **pack命令：**
        ```
        bee pack
        ```
        pack命令用来发布应用的时候的大包操作，该命令会把项目大包成zip包，然后我们就可以在部署的时候直接把大包后的项目上传到服务器，然后解压进行部署。
    
    * **version命令：**
        ```
        bee version
        ```
        version命令来查看当前bee，beego，go的版本。
    

### 三、使用Bee工具

常见的几个bee工具命令就是上面的几个，立即来练习使用一下。打开命令行终端，进入到GOPATH所对应的目录，然后进入到src目录中。

* **1、创建项目**
使用bee run命令来新建一个案例项目，比如我们新建一个BeegoDemo2的项目，我们执行命令：bee new BeegoDemo2，命令执行效果如下：
![http://7xtcwd.com1.z0.glb.clouddn.com/bee_new.png](http://7xtcwd.com1.z0.glb.clouddn.com/bee_new.png)
可以看到，终端中输出了bee的图标和版本，并且打印出了很多日志，创建了很多文件和目录。这就是bee new命令的执行效果，表示新建一个项目。

* **2、项目目录结构**
使用我们的开发公局goland来打开新建的BeegoDemo2项目，并查看一下项目目录组织结构：
![http://7xtcwd.com1.z0.glb.clouddn.com/project_cont.png](http://7xtcwd.com1.z0.glb.clouddn.com/project_cont.png)

* **3、运行项目**
现在使用bee命令来运行一下案例，看一下效果。在开发工具下方的terminal中，打开terminal，输入命令：**bee run**，运行效果如下：
![http://7xtcwd.com1.z0.glb.clouddn.com/bee_run.png](http://7xtcwd.com1.z0.glb.clouddn.com/bee_run.png)

可以看到输出日志说，http服务已经在http://localhost:8080端口运行，现在我们去浏览器中验证一下，效果如下：
![http://7xtcwd.com1.z0.glb.clouddn.com/project_show.png](http://7xtcwd.com1.z0.glb.clouddn.com/project_show.png)




### 执行流程分析
Go语言执行的时候是执行main包下面的init函数，main函数依次执行。因此，先找到main.go文件。
![http://7xtcwd.com1.z0.glb.clouddn.com/main_file.png](http://7xtcwd.com1.z0.glb.clouddn.com/main_file.png)

如上图就是main.go文件内容，那么代码如何执行呢？首先，import导入了两个包，一个是routers，一个是beego。而在routers包前面，可以看到有一个“_”，这表明是引入routers包，并执行init方法。这里涉及一个知识点，就是Go语言的执行过程，这里给大家一张程序执行流程图：
![http://7xtcwd.com1.z0.glb.clouddn.com/main_init.png](http://7xtcwd.com1.z0.glb.clouddn.com/main_init.png)

根据上图，可以得知程序首先到routers包下执行init方法。到router.go里面看一下具体的代码：
![http://7xtcwd.com1.z0.glb.clouddn.com/router.png](http://7xtcwd.com1.z0.glb.clouddn.com/router.png)

可以看到在router.go文件中有一个init方法，看到beego.Router()这句代码。router表示的是路由的意思，这个函数的功能是映射 URL 到 controller，第一个参数是 URL (用户请求的地址)，这里注册的是 /，也就是访问的不带任何参数的 URL，第二个参数是对应的 Controller，即将把请求分发到那个控制器来执行相应的逻辑，现在去这里设置的MainController中去看一下有什么方法：

![http://7xtcwd.com1.z0.glb.clouddn.com/controller.png](http://7xtcwd.com1.z0.glb.clouddn.com/controller.png)

MainController结构体及函数声明在default.go文件中。而这里就看到一个Get方法，方法中有三行代码。

上文在浏览器中访问的是：[http://localhost:8080](http://localhost:8080)，这是一个get请求，请求到了后台以后，什么请求参数都没有，因此，就会被“/”拦截，执行到MainController中的代码，因为是get请求，所以这里自动找到Get函数并进行执行。

在get函数里面，有三句代码，前两句c.Data[]= ""表示设置返回的数据字段及内容，最后一句c.TplName表示设置处理该请求指向某个模板文件，这里指向了index.tpl，那么index.tpl文件在哪里呢？可以查看项目目录结构，在views下面，views下面存放一些模板文件。

* **模板文件**  

    简单解释一下，通常的页面都是使用静态的html+css+js等这些静态代码来进行页面的布局，页面效果控制等，而把页面的数据使用变量表示，这样，在进行页面展示的时候，就能够自动的填充页面里面的变量的值；这些静态的代码文件统称为模板文件。每个模板文件就是像一个模板一样，样式效果都固定，只是根据数据不一样进行渲染和展示。
    
###### **beego.Run()逻辑**
init方法分析完毕后，程序会继续往下执行，就到了main函数，在main函数中执行：beego.Run()代码。分析一下Run代码的逻辑，在Run方法内部，主要做了几件事：
* 1）解析配置文件，也就是我们的app.conf文件，比如端口，应用名称等信息。
* 2）检查是否开启session，如果开启session，就会初始化一个session对象。
* 3）是否编译模板，beego框架会在项目启动的时候根据配置把views目录下的所有模板进行预编译，然后存放在map中，这样可以有效的提高模板运行的效率，不需要进行多次编译。
* 4）监听服务端口。根据app.conf文件中的端口配置，启动监听。


## Beego框架总结及项目初始化配置
**@author：Davie**
**版权所有：北京千锋互联科技有限公司**

### **一、Beego项目组织架构**
上节课程内容对beego的案例代码进行了一个简单的分析，总结一下beego项目的组织结构，总结如下：

* **conf：**项目配置文件所在的目录，项目中有一些全局的配置都可以放在此目录下。默认的app.conf文件中默认指定了三个配置：

    * **1）appname = BeegoDemo2：** 指定项目名称。
    * **2）httpport = 8080：** 指定项目服务监听端口。
    * **3）runmode = dev：** 指定执行模式。
    
* **controllers：**该目录是存放控制器文件的目录，所谓控制器就是控制应用调用哪些业务逻辑，由controllers处理完http请求以后，并负责返回给前端调用者。

* **models：**models层可以解释为实体层或者数据层，在models层中实现和用户和业务数据的处理，主要和数据库表相关的一些操作会在这一目录中实现，然后将执行后的结果数据返回给controller层。比如向数据库中插入新数据，删除数据库表数据，修改某一条数据，从数据库中查询业务数据等都是在models层实现。

* **routers：**该层是路由层。所谓路由就是分发的意思，当前端浏览器进行一个http请求达到后台web项目时，必须要让程序能够根据浏览器的请求url进行不同的业务处理，从接收到前端请求到判断执行具体的业务逻辑的过程的工作，就由routers来实现。

* **static：**在static目录下，存放的是web项目的静态资源文件，主要有：css文件，img，js，html这几类文件。html中会存放应用的静态页面文件。

* **views：**views中存放的就是应用中存放html模版页面的目录。所谓模版，就是页面框架和布局是已经使用html写好了的，只需要在进行访问和展示时，将获取到的数据动态填充到页面中，能够提高渲染效率。因此，模版文件是非常常见的一种方式。

**综上，这个整个的项目架构就是MVC的运行模式。常见的架构模式大家应该熟悉比如：MVC，MVP，MVVM等架构。这里我们采用MVC，比较清晰易懂。**


### 二、路由设置
在beego框架中，支持四种路由设置，他们分别是：**基础路由**，**固定路由**，**正则路由**和**自动路由**。

* **1、基础路由：**  
    
    直接通过beego.Get,beego.POST,beego.Head,beego.Delete等方法来进行路由的映射，我们通过代码给大家来进行演示。我们知道，常见的http请求方法操作有：GET，HEAD，PUT，POST，DELETE，OPTIONS等。
    1.首先是基础的GET路由。下面我们通过一个例子来向大家进行展示。beego.GET("",func)
2.基础的POSt路由：beego.POST("",func)
3.除此之外，还支持Patch，Head，Delete等基础路由。

以上这种请求和对应找到请求方法类型的方式就是RESTful形式，我们最常用的也是RESTful形式
RESTful是目前开发API中非常常用的一种形式，其实就是RESTful形式就是用户是get请求时就自动执行Get方法，Post请求就执行Post方法。
    
* **2、固定路由：**
    诸如以下此种代码形式：
    
    ```
    beego.Router("/",controller);
    ```
    
    Get请求就会对应到Get方法，Post就会对应到Post方法，Delete就会对应到Delete方法，Header方法就会对应Header方法。

* **3、正则路由：**
    正则路由是指可以在进行固定路由的基础上，支持匹配一定格式的正则表达式。比如:id,:username,自定义正则,file的路径和后缀切换以及全匹配等操作。
    
* **4、自定义路由**
上面两种路由都是默认的根据请求的类型，Get就执行Get方法，Post执行Post方法，比较局限。因为在开发的时候大部分都是使用固定匹配想要直接执行对应的逻辑控制方法，因此beego提供给我们了可以自定义的自定义路由配置。方式如下：
```
beego.Router("/",&IndexController{},"")
```
可以用的HTTP Method：
* "*"：包含以下所有的函数
* "get"：GET 请求
* "post"：POST 请求
* "put"：PUT 请求
* "delete"：DELETE 请求
* "patch"：PATCH 请求
* "options"：OPTIONS 请求
* "head"：HEAD 请求

上面这些对应的请求方法，在beego.Controller中，定义了很多的方法，比如Init，Prepare，Post，Get，Head，Delete等方法。以上分析的是beego框架的Router以及Controller运行逻辑。

### 三、model层分析
在Web开发过程中，开发者操作数据，存储数据最常用的就是数据库操作，数据库操作是独立的一个模块，这个在后面具体的项目中会带着大家实现。

### 四、view编写
这里的实战项目静态的html模版文件以及一些其他的静态资源，是已经写好提供给开发者直接可以拿来放在项目中进行使用。在实际的开发过程中，要求大家需要掌握html，js语法以及简单的css语法要能看懂，能够时间简单的使用。更加复杂的前端页面或者静态页面，需要和前端开发人员进行对接。开发重心是在后台服务端，因此，只需要学会使用和调用调试即可。

### 五、静态文件的设置
在goweb项目中，上面说过是有一些静态资源文件的，如果用户要访问静态资源文件，我们应该也是能够访问的，这需要在我们的项目中进行静态资源设置。
设置方式为：
```
beego.SetStaticPath("/down1","download1")
```
这里的download目录是指的非goweb项目的static目录下目录，而是开发者重新新建的另外的目录。


### 六、实战项目介绍
在本系列课程中，我们一起使用Beego框架开发实现一个博客系统。如下图所示：
![http://7xtcwd.com1.z0.glb.clouddn.com/WX20190515-124022@2x.png](http://7xtcwd.com1.z0.glb.clouddn.com/WX20190515-124022@2x.png)

![http://7xtcwd.com1.z0.glb.clouddn.com/WX20190515-124056@2x.png](http://7xtcwd.com1.z0.glb.clouddn.com/WX20190515-124056@2x.png)

![http://7xtcwd.com1.z0.glb.clouddn.com/WX20190515-124145@2x.png](http://7xtcwd.com1.z0.glb.clouddn.com/WX20190515-124145@2x.png)

![http://7xtcwd.com1.z0.glb.clouddn.com/WX20190515-124339@2x.png](http://7xtcwd.com1.z0.glb.clouddn.com/WX20190515-124339@2x.png)



## 数据库配置及连接测试

**@author：Davie**
**版权所有：北京千锋互联科技有限公司**

在上节内容中总结了beego框架，本节课继续学习数据库配置和ORM映射关系介绍

### mysql数据库安装
mysql官方下载网站：[https://dev.mysql.com/downloads/](https://dev.mysql.com/downloads/)
我们使用的是5.7版本，下载链接：[https://dev.mysql.com/downloads/mysql/5.7.html#downloads](https://dev.mysql.com/downloads/mysql/5.7.html#downloads)

![http://7xtcwd.com1.z0.glb.clouddn.com/mysql.png](http://7xtcwd.com1.z0.glb.clouddn.com/mysql.png)

选择下载安装文件，并在本地进行安装。
安装过程中，若出现初始密码，要记住该密码，首次登陆时要使用。

**【注意】windows安装说明：[https://jingyan.baidu.com/article/09ea3ededa53b7c0aede39b5.html](https://jingyan.baidu.com/article/09ea3ededa53b7c0aede39b5.html)**

安装完毕后，将mysql的bin目录路径添加配置到环境变量，以便能够在终端命令行中进行使用登陆mysql。

在终端中登陆mysql的命令：
```
mysql -u root -p
```

输入安装时或者自己设置的密码即可登陆成功。如下所示：
![http://7xtcwd.com1.z0.glb.clouddn.com/mysql_login_1.png](http://7xtcwd.com1.z0.glb.clouddn.com/mysql_login_1.png)
如上图所示表示登陆成功。为了使用自己的密码，可以自己设置一个自己熟悉的密码，方便自己记忆。

### mysql数据库常用命令
mysql数据库的一些命令使用：
* 查看数据库：

    ```
    show databases;
    ```
* 使用某个数据库：

    ```
    use databaseName;
    ```

* 展示某个数据库表格列表：

    ```
    show tables;
    ```
    
* mysql中的一些其他命令：

    ![http://7xtcwd.com1.z0.glb.clouddn.com/mysql_command.png](http://7xtcwd.com1.z0.glb.clouddn.com/mysql_command.png)
    
以上mysql数据操作都是命令行终端形式，为了方便我们日常操作，我们可以使用图形化界面工具，在这里我们使用的是navicat工具。我们现在来看如何进行安装。

### Navicat安装
navicat工具下载地址：[https://www.navicat.com/en/download/navicat-for-mysql](https://www.navicat.com/en/download/navicat-for-mysql)

在上面的链接中，选择自己的系统版本，然后下载安装文件，进行安装，一直默认选下一步即可，最后安装成功。

安装好了以后，在mysql数据库中创建一个新的数据库，比如名字为：**myblog**;创建完数据库以后，我们来看一看在beego项目中如何使用mysql数据库。

### 数据库驱动

我们之前给大家说过，数据库编程属于beego中的Models层，也称ORM模块。

在beego中，目前支持三种数据库驱动，分别是：

* **MySQL：** [github.com/go-sql-driver/mysql](github.com/go-sql-driver/mysql)
* **PostgreSQL：**[github.com/lib/pq](github.com/lib/pq)
* **Sqlite3：**[github.com/mattn/go-sqlite3](github.com/mattn/go-sqlite3)

beego中的ORM所具备的几个特性：

* **1）支持Go语言的所有类型存储**
* **2）CRUD操作简单**
* **3）自动Join关联表**
* **4）允许直接使用SQL查询**

### beego项目中使用mysql

* **1、导入对应的数据库驱动**
比如mysql：

    ```
    import _ "github.com/go-sql-driver/mysql"
    ```
![http://7xtcwd.com1.z0.glb.clouddn.com/driver_mysql.png](http://7xtcwd.com1.z0.glb.clouddn.com/driver_mysql.png)

* **2、注册驱动，连接数据库**
通过如下两句话注册数据库驱动，以及连接数据库的操作：
    ```
    orm.RegisterDriver("mysql",orm.DRMySQL)
    orm.RegisterDataBase(aliasName,driverName,dbConn)
    ```
详细代码如下：
![http://7xtcwd.com1.z0.glb.clouddn.com/conn_mysql.png](http://7xtcwd.com1.z0.glb.clouddn.com/conn_mysql.png)

* **3、创建数据库并执行程序**
![http://7xtcwd.com1.z0.glb.clouddn.com/WX20190515-150049@2x.png](http://7xtcwd.com1.z0.glb.clouddn.com/WX20190515-150049@2x.png)

连接数据库代码如下：

```
package models

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"BlogProject/MysqlDemo/util"
	//切记：导入驱动包
	_ "github.com/go-sql-driver/mysql"
)

func init() {

	driverName := beego.AppConfig.String("driverName")

	//注册数据库驱动
	orm.RegisterDriver(driverName, orm.DRMySQL)

	//数据库连接
	user := beego.AppConfig.String("mysqluser")
	pwd := beego.AppConfig.String("mysqlpwd")
	host := beego.AppConfig.String("host")
	port := beego.AppConfig.String("port")
	dbname := beego.AppConfig.String("dbname")

	//dbConn := "root:yu271400@tcp(127.0.0.1:3306)/cmsproject?charset=utf8"
	dbConn := user + ":" + pwd + "@tcp(" + host + ":" + port + ")/" + dbname + "?charset=utf8"

	err := orm.RegisterDataBase("default", driverName, dbConn)
	if err != nil {
		util.LogError("连接数据库出错")
		return
	}
	util.LogInfo("连接数据库成功")

}

```
![http://7xtcwd.com1.z0.glb.clouddn.com/WX20190515-152119@2x.png](http://7xtcwd.com1.z0.glb.clouddn.com/WX20190515-152119@2x.png)





## 项目搭建以及注册用户信息功能
**@author：Davie**
**版权所有：北京千锋互联科技有限公司**

## 一、搭建项目
首先打开终端进入gopath下的src目录，然后执行以下命令，创建一个beego项目：

```shell
bee new myblog
```

运行效果如下：

![http://7xtcwd.com1.z0.glb.clouddn.com/WX20190519-230051@2x.png](http://7xtcwd.com1.z0.glb.clouddn.com/WX20190519-230051@2x.png)

然后通过goland打开该项目：
![http://7xtcwd.com1.z0.glb.clouddn.com/WX20190520-062806@2x.png](http://7xtcwd.com1.z0.glb.clouddn.com/WX20190520-062806@2x.png)

我们打开conf包下的配置文件：app.conf文件，修改端口号为8080：

```go
appname = myblogweb
httpport = 8080
runmode = dev
```

然后在终端中进入该项目目录，然后运行项目：
![http://7xtcwd.com1.z0.glb.clouddn.com/WX20190520-063001@2x.png](http://7xtcwd.com1.z0.glb.clouddn.com/WX20190520-063001@2x.png)

我们可以看到，项目已经运行了，并且监听在8080的端口上。接下来我们打开浏览器输入网址：http://127.0.0.1:8080/，然后可以看到欢迎界面：
![http://7xtcwd.com1.z0.glb.clouddn.com/WX20190520-063119@2x.png](http://7xtcwd.com1.z0.glb.clouddn.com/WX20190520-063119@2x.png)

## 二、注册功能

### 2.1model

首先我们先在创建一个数据库：
![http://7xtcwd.com1.z0.glb.clouddn.com/WX20190520-063304@2x.png](http://7xtcwd.com1.z0.glb.clouddn.com/WX20190520-063304@2x.png)

我们先创建一个工具包utils，然后创建一个go文件，用于做mysql的工具类，里面提供连接数据库和创建表的功能。

首先提供一个初始化方法：
```
func InitMysql() {

	fmt.Println("InitMysql....")
	driverName := beego.AppConfig.String("driverName")

	//注册数据库驱动
	orm.RegisterDriver(driverName, orm.DRMySQL)

	//数据库连接
	user := beego.AppConfig.String("mysqluser")
	pwd := beego.AppConfig.String("mysqlpwd")
	host := beego.AppConfig.String("host")
	port := beego.AppConfig.String("port")
	dbname := beego.AppConfig.String("dbname")

	//dbConn := "root:yu271400@tcp(127.0.0.1:3306)/cmsproject?charset=utf8"
	dbConn := user + ":" + pwd + "@tcp(" + host + ":" + port + ")/" + dbname + "?charset=utf8"

	//dbConn := "root:yu271400@tcp(127.0.0.1:3306)/cmsproject?charset=utf8"
	dbConn := user + ":" + pwd + "@tcp(" + host + ":" + port + ")/" + dbname + "?charset=utf8"

	db, _ = sql.Open(driverName, dbConn)
```


然后设计一下数据表，我们需要用户的id，并且作为主键，用户名username和密码password，还有状态status，这个用于标记该用户是否被删除，0表示正常状态，1表示删除。还要就是注册时间，我们用整型的时间戳来表示即可，所以sql语句如下：可以直接操作数据库创建，也可以通过代码进行创建：

```go
//创建用户表
func CreateTableWithUser() {
	sql := `CREATE TABLE IF NOT EXISTS users(
		id INT(4) PRIMARY KEY AUTO_INCREMENT NOT NULL,
		username VARCHAR(64),
		password VARCHAR(64),
		status INT(4),
		createtime INT(10)
		);`

	ModifyDB(sql)
}
```

然后再提供一个方法，用于执行sql语句：

```go
//操作数据库
func ModifyDB(sql string, args ...interface{}) (int64, error) {
	result, err := db.Exec(sql, args...)
	if err != nil {
		log.Println(err)
		return 0, err
	}
	count, err := result.RowsAffected()
	if err != nil {
		log.Println(err)
		return 0, err
	}
	return count, nil
}
```



再提供一个方法，用于之后的数据库操作：

```go

//查询
func QueryRowDB(sql string) *sql.Row{
	return db.QueryRow(sql)
}
```

然后再models中创建一个model文件：

```go
package models

import (
	"myblogweb/utils"
	"fmt"
)

type User struct {
	Id         int
	Username   string
	Password   string
	Status     int // 0 正常状态， 1删除
	Createtime int64
}

//--------------数据库操作-----------------

//插入
func InsertUser(user User)(int64, error){
	return utils.ModifyDB("insert into users(username,password,status,createtime) values (?,?,?,?)",
		user.Username,user.Password,user.Status,user.Createtime)
}

//按条件查询
func QueryUserWightCon(con string)int{
	sql := fmt.Sprintf("select id from users %s",con)
	fmt.Println(sql)
	row:=utils.QueryRowDB(sql)
	id :=0
	row.Scan(&id)
	return id
}

//根据用户名查询id
func QueryUserWithUsername(username string) int{
	sql := fmt.Sprintf("where username='%s'",username)
	return QueryUserWightCon(sql)
}

//根据用户名和密码，查询id
func QueryUserWithParam(username ,password string)int{
	sql:=fmt.Sprintf("where username='%s' and password='%s'",username,password)
	return QueryUserWightCon(sql)
}
```

### 2.2 view
我们先在views包下，创建一个html页面：register.html。
```html
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <title>注册</title>
    <link rel="stylesheet" type="text/css" href="../static/css/lib/login.css">
    <link rel="stylesheet" type="text/css" href="../static/css/blogsheet.css">
    <script src="../static/js/lib/jquery-3.3.1.min.js"></script>
    <script src="../static/js/lib/jquery.url.js"></script>
    <script src="../static/js/blog.js"></script>
</head>
<body>
    <div id="nav">
        <div id="nav-login">
            <ul>
                <li><a href="/login">登录</a></li>
                <li><a href="/register">注册</a></li>

            </ul>

        </div>
    </div>


    <div class="htmleaf-container">
        <div class="wrapper">
            <!--注册表单-->
            <div class="container">
                <h1>Welcome</h1>
                <form id="register-form" class="form">
                    <input type="text" name="username" placeholder="Username">
                    <input type="password" name="password" placeholder="Password" id="register-password">
                    <input type="password" name="repassword" placeholder="rePassword">
                    <br>
                    <button type="submit" id="login-button">Register</button>
                </form>
            </div>

            {{/*背景动画*/}}
            <ul class="bg-bubbles">
                <li></li>
                <li></li>
                <li></li>
                <li></li>
                <li></li>
                <li></li>
                <li></li>
                <li></li>
                <li></li>
                <li></li>
            </ul>
        </div>
    </div>
</body>
</html>
```

页面效果：
![http://7xtcwd.com1.z0.glb.clouddn.com/WX20190520-064918@2x.png](http://7xtcwd.com1.z0.glb.clouddn.com/WX20190520-064918@2x.png)

添加表单验证：

我们使用jquery来实现js部分，在js目录下，创建子目录lib，在里面放jquery框架文件，并在页面中引入。

然后重新创建一个js文件：blog.js

首先添加表单验证：

```js
$(document).ready(function () {
    //注册表单验证
    $("register-from").validate({
        rules:{
            username:{
                required:true,
                rangelength:[5,10]
            },
            password:{
                required:true,
                rangelength:[5,10]
            },
            repassword:{
                required:true,
                rangelength:[5,10],
                equalTo:"#register-password"
            }
        },
        messages:{
            username:{
                required:"请输入用户名",
                rangelength:"用户名必须是5-10位"
            },
            password:{
                required:"请输入密码",
                rangelength:"密码必须是5-10位"
            },
            repassword:{
                required:"请确认密码",
                rangelength:"密码必须是5-10位",
                equalTo:"两次输入的密码必须相等"
            }
        },
        submitHandler:function (form) {
            var urlStr = "/register";
            // alert("urlStr:"+urlStr)
            $(form).ajaxSubmit({
                url:urlStr,
                type:"post",
                dataType:"json",
                success:function (data,status) {
                    alert("data:"+data.message)
                    if (data.code == 1){
                        setTimeout(function () {
                            window.location.href="/login"
                        },1000)
                    }
                },
                err:function (data,status) {
                    alert("err:"+data.message+":"+status)
                }
            })
        }
    })
})
```

当用户点击表单的submit按钮，会跳转到/register路径，而因为设置为post请求，所以我们可以在post中完成表单的注册。如果注册成功，再跳转到/login路径。

### 2.3 controller

我们在controllers包下创建一个新的controller，用来处理用户的注册：

```go
package controllers

import "github.com/astaxie/beego"

type RegisterController struct {
	beego.Controller
}

func (this *RegisterController) Get(){
	this.TplName = "register.html"
}

func (this *RegisterController) Post() {

	//获取表单信息
	username := this.GetString("username")
	password := this.GetString("password")
	repassword := this.GetString("repassword")
	fmt.Println(username, password, repassword)
	log.INFO(username, password, repassword)

	//注册之前先判断该用户名是否已经被注册，如果已经注册，返回错误
	id := models.QueryUserWithUsername(username)
	fmt.Println("id:", id)
	if id > 0 {
		this.Data["json"] = map[string]interface{}{"code": 0, "message": "用户名已经存在"}
		this.ServeJSON()
		return
	}

	//注册用户名和密码
	//存储的密码是md5后的数据，那么在登录的验证的时候，也是需要将用户的密码md5之后和数据库里面的密码进行判断
	password = utils.MD5(password)
	fmt.Println("md5后：", password)

	user := models.User{0, username, password, 0, time.Now().Unix()}
	_, err := models.InsertUser(user)
	if err != nil {
		this.Data["json"] = map[string]interface{}{"code": 0, "message": "注册失败"}
	} else {
		this.Data["json"] = map[string]interface{}{"code": 1, "message": "注册成功"}
	}
	this.ServeJSON()

}

```

接下来，我们需要新注册一个路由，修改router.go文件：
```go
func init() {
    beego.Router("/", &controllers.MainController{})
    beego.Router("/register", &controllers.RegisterController{})
}
```

然后我们创建一个Post()方法，用于处理post的请求：
```go

//处理注册
func (this *RegisterController) Post() {
	//获取表单信息
	username := this.GetString("username")
	password := this.GetString("password")
	repassword := this.GetString("repassword")
	fmt.Println(username, password, repassword)

	//注册之前先判断该用户名是否已经被注册，如果已经注册，返回错误
	id := models.QueryUserWithUsername(username)
	fmt.Println("id:",id)
	if id > 0 {
		this.Data["json"] = map[string]interface{}{"code":0,"message":"用户名已经存在"}
		this.ServeJSON()
		return
	}

	//注册用户名和密码
	//存储的密码是md5后的数据，那么在登录的验证的时候，也是需要将用户的密码md5之后和数据库里面的密码进行判断
	password = utils.MD5(password)
	fmt.Println("md5后：",password)

	user := models.User{0,username,password,0,time.Now().Unix()}
	_,err :=models.InsertUser(user)
	if err != nil{
		this.Data["json"] = map[string]interface{}{"code":0,"message":"注册失败"}
	}else{
		this.Data["json"] = map[string]interface{}{"code":1,"message":"注册成功"}
	}
	this.ServeJSON()
}
```

思路就是先接收表单信息，然后判断数据库中是否已经存在了该用户名，如果已经存在，那么无法进行注册，我们可以通过json返回信息：用户名已经存在。

否则再进行注册，出于对密码安全的考虑，我们可以在数据库中存储密码的md5加密数据。

所以在工具包中，再添加一个工具类：myUtils.go

```go
package utils

import (
	"fmt"
	"crypto/md5"
)

//传入的数据不一样，那么MD5后的32位长度的数据肯定会不一样
func MD5(str string) string{
	md5str:=fmt.Sprintf("%x",md5.Sum([]byte(str)))
	return md5str
}

```

### 2.3 运行

我们打开终端，进入项目所在的目录，执行命令：

```shell
bee run
```

![http://7xtcwd.com1.z0.glb.clouddn.com/WX20190520-065447@2x.png](http://7xtcwd.com1.z0.glb.clouddn.com/WX20190520-065447@2x.png)

项目启动后，监听在8080端口上。

并且执行了InitMysql()方法，初始化数据库，我们打开数据库进行刷新，已经创建好了一张数据表user：

![http://7xtcwd.com1.z0.glb.clouddn.com/WX20190520-083914@2x.png](http://7xtcwd.com1.z0.glb.clouddn.com/WX20190520-083914@2x.png)

打开浏览器，输入以下网址：[http://localhost:8080/register](http://localhost:8080/register) ,然后输入用户名和密码进行注册。

![http://7xtcwd.com1.z0.glb.clouddn.com/WX20190520-084747@2x.png](http://7xtcwd.com1.z0.glb.clouddn.com/WX20190520-084747@2x.png)

![http://7xtcwd.com1.z0.glb.clouddn.com/WX20190520-084800@2x.png](http://7xtcwd.com1.z0.glb.clouddn.com/WX20190520-084800@2x.png)

此处，大家可以反复的试验，比如用户名长度不够，重复密码不一致，相同用户名等。





# 登录功能实现

上一章节我们已经实现了注册用户名和密码。接下来我们实现登录。


## 一、定义Controller

先创建一个新的go文件，用来定义新的Controller：

```go

type LoginController struct {
	beego.Controller
}

func (this *LoginController) Get() {
	this.TplName = "login.html"
}
```

然后注册一个新的路由：

```go
func init() {
    beego.Router("/", &controllers.MainController{})
    beego.Router("/register", &controllers.RegisterController{})
    beego.Router("/login", &controllers.LoginController{})
}
```

然后添加一个Post方法：

```go

func (this *LoginController) Post() {
	username := this.GetString("username")
	password := this.GetString("password")
	fmt.Println("username:", username, ",password:", password)

	id := models.QueryUserWithParam(username, utils.MD5(password))
	fmt.Println("id:",id)
	if id > 0 {
		this.Data["json"] = map[string]interface{}{"code": 1, "message": "登录成功"}
	} else {
		this.Data["json"] = map[string]interface{}{"code": 0, "message": "登录失败"}
	}
	this.ServeJSON()
}

```



## 二、Model操作

在之前的Post方法中，我们已经接收页面传来的用户名和密码，然后user_model.go中添加一个方法，用于根据用户名和密码来查询id，表示用户登录的用户名和密码是否正确。

```go
//根据用户名和密码，查询id
func QueryUserWithParam(username ,password string)int{
	sql:=fmt.Sprintf("where username='%s' and password='%s'",username,password)
	return QueryUserWightCon(sql)
}
```

## 三、View页面

我们重新再views包下创建一个html页面，login.html，内容如下：

```html
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <title>登录</title>
    <link rel="stylesheet" type="text/css" href="../static/css/lib/login.css">
    <link rel="stylesheet" type="text/css" href="../static/css/blogsheet.css">
    <script src="../static/js/lib/jquery-3.3.1.min.js"></script>
    <script src="../static/js/lib/jquery.url.js"></script>
    <script src="../static/js/blog.js"></script>
</head>
<body>
<div id="nav">
    <div id="nav-login">
        <ul>
            <li><a href="/login">登录</a></li>
            <li><a href="/register">注册</a></li>

        </ul>

    </div>
</div>


<div class="htmleaf-container">
    <div class="wrapper">
    {{/*注册表单*/}}
        <div class="container">
            <h1>Welcome</h1>
            <form id="login-form" class="form">
                <input type="text" name="username" placeholder="请输入用户名">
                <input type="password" name="password" placeholder="请输入密码" id="login-password">
                <br>
                <button type="submit" id="login-button">登录</button>
            </form>
        </div>

    {{/*背景动画*/}}
        <ul class="bg-bubbles">
            <li></li>
            <li></li>
            <li></li>
            <li></li>
            <li></li>
            <li></li>
            <li></li>
            <li></li>
            <li></li>
            <li></li>
        </ul>
    </div>
</div>

</body>
</html>
```

> 因为内容和注册页类似，可以将注册页的内容复制过来，但是要注意标签id等的不同。



接下来我们实现js部分，在blog.js中添加代码，实现登录部分：

```js

    //登录
    $("#login-form").validate({
        rules:{
            username:{
                required:true,
                rangelength:[5,10]
            },
            password:{
                required:true,
                rangelength:[5,10]
            }
        },
        messages:{
            username:{
                required:"请输入用户名",
                rangelength:"用户名必须是5-10位"
            },
            password:{
                required:"请输入密码",
                rangelength:"密码必须是5-10位"
            }
        },
        submitHandler:function (form) {
            var urlStr ="/login"
            alert("urlStr:"+urlStr)
            $(form).ajaxSubmit({
                url:urlStr,
                type:"post",
                dataType:"json",
                success:function (data,status) {
                    alert("data:"+data.message+":"+status)
                    if(data.code == 1){
                        setTimeout(function () {
                            window.location.href="/"
                        },1000)
                    }
                },
                error:function (data,status) {
                    alert("err:"+data.message+":"+status)
                }
            });
        }
    });
```

## 四、运行

启动项目后，打开浏览器输入以下访问路径：http://127.0.0.1:8080/login

然后输入用户名和密码，并且登录，页面返回信息依次如下：

![http://7xtcwd.com1.z0.glb.clouddn.com/WX20190520-090352@2x.png](http://7xtcwd.com1.z0.glb.clouddn.com/WX20190520-090352@2x.png)

![http://7xtcwd.com1.z0.glb.clouddn.com/WX20190520-090526@2x.png](http://7xtcwd.com1.z0.glb.clouddn.com/WX20190520-090526@2x.png)

![http://7xtcwd.com1.z0.glb.clouddn.com/WX20190520-090536@2x.png](http://7xtcwd.com1.z0.glb.clouddn.com/WX20190520-090536@2x.png)






    
## 项目实战处理Session

**@author：Davie**
**版权所有：千锋互联科技有限公司**

我们希望用户登录后可以有写博客的功能，当然也允许用户退出。实现该功能的操作需要使用session。

![http://7xtcwd.com1.z0.glb.clouddn.com/WX20190521-104344@2x.png](http://7xtcwd.com1.z0.glb.clouddn.com/WX20190521-104344@2x.png)

如果用户没有登录，直接访问首页，只可以查看文章，标签，相册等，但是没有写博客的功能：

![http://7xtcwd.com1.z0.glb.clouddn.com/WX20190521-104814@2x.png](http://7xtcwd.com1.z0.glb.clouddn.com/WX20190521-104814@2x.png)

## 一、Controller

首先我们在登录的时候添加session。

先修改配置文件打开session，打开conf包下的app.conf文件，添加一行：

```go
appname = myblog
httpport = 8080
runmode = dev

sessionon = true
```

或者在main.go中打开session：
```go
func main() {
	utils.InitMysql()
	beego.BConfig.WebConfig.Session.SessionOn = true // 打开session
	beego.Run()
}
```

> 以上两个操作，选择一个即可。

然后修改登录的Post方法，在登录的时候，我们设置session：
```go
func (this *LoginController) Post() {
	...
	if id > 0 {
		/*
		设置了session后悔将数据处理设置到cookie，然后再浏览器进行网络请求的时候回自动带上cookie
		因为我们可以通过获取这个cookie来判断用户是谁，这里我们使用的是session的方式进行设置
		 */
		this.SetSession("loginuser", username)
		this.Data["json"] = map[string]interface{}{"code": 1, "message": "登录成功"}
	} else {
		this.Data["json"] = map[string]interface{}{"code": 0, "message": "登录失败"}
	}
	this.ServeJSON()
}
```

接下来，我们新添加一个路由：
```go
func init() {
    beego.Router("/", &controllers.HomeController{})
    beego.Router("/register", &controllers.RegisterController{})
    beego.Router("/login", &controllers.LoginController{})
}
```

当访问/，这个根路径的时候，我们表示访问首页。

接下来创建一个新的Controller，HomeController，用于控制首页。但是在这之前，我们先设置一个父Controller用于获取session，查看用户是否登录。

先创建一个go文件：base_controller.go
```go
type BaseController struct {
	beego.Controller
	IsLogin   bool
	Loginuser interface{}
}
```

先定义个BaseController，里面额外设置两个字段，IsLogin表示用户是否登录，Loginuser表示用户名。

```go
//判断是否登录
func (this *BaseController) Prepare() {
	loginuser := this.GetSession("loginuser")
	fmt.Println("loginuser---->", loginuser)
	if loginuser != nil {
		this.IsLogin = true
		this.Loginuser = loginuser
	} else {
		this.IsLogin = false
	}
	this.Data["IsLogin"] = this.IsLogin
}
```

接下类，重写Prepare()方法，用于获取session。

```go
//判断是否登录
func (this *BaseController) Prepare() {
	loginuser := this.GetSession("loginuser")
	fmt.Println("loginuser---->", loginuser)
	if loginuser != nil {
		this.IsLogin = true
		this.Loginuser = loginuser
	} else {
		this.IsLogin = false
	}
	this.Data["IsLogin"] = this.IsLogin
}
```

这个函数主要是为了用户扩展用的，这个函数会在下面定义的这些 Method 方法之前执行，用户可以重写这个函数实现类似用户验证之类。然后再创建一个home_controller.go文件。

```go
type HomeController struct {
	//beego.Controller
	BaseController
}
```

在Get()方法 中，渲染home页面即可。

```go
func (this *HomeController)Get(){
	fmt.Println("IsLogin:",this.IsLogin,this.Loginuser)
	this.TplName="home.html"
}
```

## 二、View

我们需要如下几个功能：
![http://7xtcwd.com1.z0.glb.clouddn.com/WX20190521-111948@2x.png](http://7xtcwd.com1.z0.glb.clouddn.com/WX20190521-111948@2x.png)

后续的每个页面都有这几个功能。所以我们可以让页面嵌套。

现在views目录下创建一个子目录block，里面创建一个html页面：nav.html，内容如下：

```html
<div id="nav">
    <div id="nav-write-article">
        <ul>
        {{/*如果已经登录，才会显示"写博客"*/}}
        {{if .IsLogin}}
            <li><a href="/article/add">写博客</a></li>
        {{end}}
        </ul>
    </div>


    <div id="nav-menu">
        <ul>
            <li><a href="/">首页</a></li>
            <li><a href="/tags">标签</a></li>
            <li><a href="/album">相册</a></li>
            <li><a href="/#">关于我</a></li>
        </ul>
    </div>


    <div id="nav-login">
        <ul>
        {{if .IsLogin}}
            <li><a href="">退出</a></li>
        {{else}}
            <li><a href="/login">登录</a></li>
            <li><a href="/register">注册</a></li>
        {{end}}
        </ul>
    </div>

</div>
```

当然这里需要判断用户是否登录，如果登录，显示写博客功能，以及可以退出。否则显示登录或注册。

接下来，我们创建一个新的页面：**home.html**，内容如下：
```html
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <title>首页</title>
    <link href="../static/css/blogsheet.css" rel="stylesheet">
</head>
<body>
{{template "block/nav.html" .}}


</body>
</html>
```

通过页面的嵌套，我们可以将标签显示到home页面上。

> 特别注意，{{template "block/nav.html" .}}，后面的`.`，这是传递当前参数到子模板

同样，我们将之前的login.html页面和register.html页面，也将nav嵌套进来，代码如下：

login.html
```html
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <title>登录</title>
    <link rel="stylesheet" type="text/css" href="../static/css/lib/login.css">
    <link rel="stylesheet" type="text/css" href="../static/css/blogsheet.css">
    <script src="../static/js/lib/jquery-3.3.1.min.js"></script>
    <script src="../static/js/lib/jquery.url.js"></script>
    <script src="../static/js/blog.js"></script>
</head>
<body>
<!--
<div id="nav">
    <div id="nav-login">
        <ul>
            <li><a href="/login">登录</a></li>
            <li><a href="/register">注册</a></li>

        </ul>

    </div>
</div>
-->
{{template "block/nav.html"}}

<div class="htmleaf-container">
    <div class="wrapper">
    {{/*注册表单*/}}
        <div class="container">
            <h1>Welcome</h1>
            <form id="login-form" class="form">
                <input type="text" name="username" placeholder="请输入用户名">
                <input type="password" name="password" placeholder="请输入密码" id="login-password">
                <br>
                <button type="submit" id="login-button">登录</button>
            </form>
        </div>

    {{/*背景动画*/}}
        <ul class="bg-bubbles">
            <li></li>
            <li></li>
            <li></li>
            <li></li>
            <li></li>
            <li></li>
            <li></li>
            <li></li>
            <li></li>
            <li></li>
        </ul>
    </div>
</div>

</body>
</html>
```

register.html

```html
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <title>注册</title>
    <link rel="stylesheet" type="text/css" href="../static/css/lib/login.css">
    <link rel="stylesheet" type="text/css" href="../static/css/blogsheet.css">
    <script src="../static/js/lib/jquery-3.3.1.min.js"></script>
    <script src="../static/js/lib/jquery.url.js"></script>
    <script src="../static/js/blog.js"></script>
</head>
<body>
<!--
<div id="nav">
    <div id="nav-login">
        <ul>
            <li><a href="/login">登录</a></li>
            <li><a href="/register">注册</a></li>

        </ul>

    </div>
</div>
-->

{{template "block/nav.html"}}

<div class="htmleaf-container">
    <div class="wrapper">
    {{/*注册表单*/}}
        <div class="container">
            <h1>Welcome</h1>
            <form id="register-form" class="form">
                <input type="text" name="username" placeholder="Username">
                <input type="password" name="password" placeholder="Password" id="register-password">
                <input type="password" name="repassword" placeholder="rePassword">
                <br>
                <button type="submit" id="login-button">Register</button>
            </form>
        </div>

    {{/*背景动画*/}}
        <ul class="bg-bubbles">
            <li></li>
            <li></li>
            <li></li>
            <li></li>
            <li></li>
            <li></li>
            <li></li>
            <li></li>
            <li></li>
            <li></li>
        </ul>
    </div>
</div>

</body>
</html>
```

## 三、退出

我们先在nav.html页面中，添加退出的脚本：

```html

    <div id="nav-login">
        <ul>
        {{if .IsLogin}}
            <li><a href='javascript:if(confirm("确定退出吗？")){location="/exit"}'>退出</a></li>
        {{else}}
            <li><a href="/login">登录</a></li>
            <li><a href="/register">注册</a></li>
        {{end}}
        </ul>
    </div>
```

然后重新注册一个路由，用于处理退出：
```go
func init() {
    beego.Router("/", &controllers.HomeController{})
    //注册
    beego.Router("/register", &controllers.RegisterController{})
    //登录
    beego.Router("/login", &controllers.LoginController{})
    //退出
	beego.Router("/exit", &controllers.ExitController{})

}
```

然后再创建一个go文件：exit_controller.go

```go
package controllers

type ExitController struct {
	BaseController
}

func (this *ExitController)Get(){
	//清除该用户登录状态的数据
	this.DelSession("loginuser")

	this.Redirect("/",302)
}

```
清除session后，重定位到/路径上。

## 四、运行效果

我们重启项目后，打来浏览器输入网址：http://127.0.0.1:8081/

浏览器如下：
![http://7xtcwd.com1.z0.glb.clouddn.com/WX20190521-113931@2x.png](http://7xtcwd.com1.z0.glb.clouddn.com/WX20190521-113931@2x.png)

服务器端运行如下：
![http://7xtcwd.com1.z0.glb.clouddn.com/WX20190521-114207@2x.png](http://7xtcwd.com1.z0.glb.clouddn.com/WX20190521-114207@2x.png)

点击退出按钮：
![http://7xtcwd.com1.z0.glb.clouddn.com/WX20190521-135635@2x.png](http://7xtcwd.com1.z0.glb.clouddn.com/WX20190521-135635@2x.png)

后台程序执行效果：
![http://7xtcwd.com1.z0.glb.clouddn.com/WX20190521-135733@2x.png](http://7xtcwd.com1.z0.glb.clouddn.com/WX20190521-135733@2x.png)








## 博客项目实战之写文章功能
**@author：Davie**
**版权所有：北京千锋互联科技有限公司**

点击写文章，我们进入写文章的页面，当用户点击按钮进行提交后，应该将文章的数据存入到数据库中。

## 一、数据库

首先我们先设计数据库，用户提交的文章，包含标题，标签，简介，内容， 创建时间等。

在mysqlUtils.go文件中，添加article表的操作：

```go
//创建文章表
func CreateTableWithArticle(){
	sql:=`create table if not exists article(
		id int(4) primary key auto_increment not null,
		title varchar(30),
		author varchar(20),
		tags varchar(30),
		short varchar(255),
		content longtext,
		createtime int(10)
		);`
	ModifyDB(sql)
}
```
## 二、Model

在model目录下创建一个go文件：article_model.go

```go
package models

import "myblog/utils"

type Article struct {
	Id         int
	Title      string
	Tags       string
	Short      string
	Content    string
	Author     string
	Createtime int64
	//Status int //Status=0为正常，1为删除，2为冻结
}



//---------数据处理-----------
func AddArticle(article Article) (int64, error) {
	i, err := insertArticle(article)
	return i, err
}



//-----------数据库操作---------------

//插入一篇文章
func insertArticle(article Article) (int64, error) {
	return utils.ModifyDB("insert into article(title,tags,short,content,author,createtime) values(?,?,?,?,?,?)",
		article.Title, article.Tags, article.Short, article.Content, article.Author, article.Createtime)
}
```

目前我们只是写文章，所以需要的是添加数据。


## 三、Controller

首先创建一个controller文件，add_article_controller.go。

```go
package controllers

import (
	"fmt"
	"myblog/models"
	"time"
)

type AddArticleController struct {
	BaseController
}

/*
当访问/add路径的时候回触发AddArticleController的Get方法
响应的页面是通过TpName
 */
func (this *AddArticleController) Get() {
	this.TplName = "write_article.html"
}

//通过this.ServerJSON()这个方法去返回json字符串
func (this *AddArticleController) Post() {

	//获取浏览器传输的数据，通过表单的name属性获取值
	title := this.GetString("title")
	tags := this.GetString("tags")
	short := this.GetString("short")
	content := this.GetString("content")
	fmt.Printf("title:%s,tags:%s\n", title, tags)

	//实例化model，将它出入到数据库中

	art := models.Article{0, title, tags, short, content, "千锋教育", time.Now().Unix()}
	_, err := models.AddArticle(art)

	//返回数据给浏览器
	var response map[string]interface{}
	if err == nil {
		//无误
		response = map[string]interface{}{"code": 1, "message": "ok"}
	} else {
		response = map[string]interface{}{"code": 0, "message": "error"}
	}

	this.Data["json"] = response
	this.ServeJSON()

}
```

如果用户请求写文章路径，会展示write_article.html页面。添加完信息后，点击提交按钮，进行提交数据。

然后注册一个新的路由：
```go
//写文章
beego.Router("/article/add", &controllers.AddArticleController{})
```


## 四、View

我们在views目录下创建一个html文件(write_article.html)，用于写文章。

```html
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <title>写文章</title>
    <link href="../static/css/blogsheet.css" rel="stylesheet">
    <script src="../static/js/lib/jquery-3.3.1.min.js"></script>
    <script src="../static/js/lib/jquery.url.js"></script>
    <script src="../static/js/blog.js"></script>
</head>
<body>

{{template "block/nav.html" .}}

<div id="main">
    <form id="write-art-form" method="post">
        <div>标题</div>
        <input type="text" placeholder="请输入标题" name="title" >
        <div>标签</div>
        <input type="text" placeholder="请输入标签" name="tags" >
        <div>简介</div>
        <textarea placeholder="请输入简介" name="short"></textarea>
        <div>内容</div>
        <textarea id="content" placeholder="请输入内容" name="content"></textarea>
        <input id="write-article-id" hidden name="id" >
        <br>
        <button type="button" onclick="history.back()">返回</button>
        <button type="submit" id="write-art-submit">提交</button>
    </form>
</div>

</body>
</html>
```

接下来写js脚本文件，打开static/js目录下的blog.js文件。

```js
	//添加文章的表单
    $("#write-art-form").validate({
        rules: {
            title: "required",
            tags: "required",
            short: {
                required: true,
                minlength: 2
            },
            content: {
                required: true,
                minlength: 2
            }
        },
        messages: {
            title: "请输入标题",
            tags: "请输入标签",
            short: {
                required: "请输入简介",
                minlength: "简介内容最少两个字符"
            },
            content: {
                required: "请输入文章内容",
                minlength: "文章内容最少两个字符"
            }
        },
        submitHandler: function (form) {
            var urlStr = "/article/add";
            alert("urlStr:" + urlStr);
            $(form).ajaxSubmit({
                url: urlStr,
                type: "post",
                dataType: "json",
                success: function (data, status) {
                    alert(":data:" + data.message);
                    setTimeout(function () {
                        window.location.href = "/"
                    }, 1000)
                },
                error: function (data, status) {
                    alert("err:" + data.message + ":" + status)
                }
            });
        }
    })
```

## 五、运行

运行项目后，打开浏览器，登录后进入首页，点击写文章：
![http://7xtcwd.com1.z0.glb.clouddn.com/WX20190521-145856@2x.png](http://7xtcwd.com1.z0.glb.clouddn.com/WX20190521-145856@2x.png)

然后进入写文章页面：
![http://7xtcwd.com1.z0.glb.clouddn.com/WX20190521-150025@2x.png](http://7xtcwd.com1.z0.glb.clouddn.com/WX20190521-150025@2x.png)

点击按钮进行提交，然后查询数据库，数据已经被插入进去了：
![http://7xtcwd.com1.z0.glb.clouddn.com/WechatIMG238.png](http://7xtcwd.com1.z0.glb.clouddn.com/WechatIMG238.png)






## 首页设计
在上节课内容中，我们学习实现了写文章的功能。本节课，我们要实现首页功能的开发。首页是用户登录后要展示的页面。最终想实现的功能是：点击首页，就会自动查询数据库，显示文章；如果文章较多，我们可以实现分页。

## 一、查询文章功能

### 1.1 Controller

我们首先修改home_controller.go文件，在Get()方法中，先查询所有的文章，并显示到页面上。因为文章可能很多，为了更好的用户体验，我们需要分页查询。默认查询第一页。

```go
func (this *HomeController) Get() {
	page, _ := this.GetInt("page")
	if page <= 0 {
		page = 1
	}
	var artList []models.Article
	artList, _ = models.FindArticleWithPage(page)
	this.Data["PageCode"] = 1
	this.Data["HasFooter"] = true

	fmt.Println("IsLogin:", this.IsLogin, this.Loginuser)
	this.Data["Content"] = models.MakeHomeBlocks(artList, this.IsLogin)

	this.TplName = "home.html"
}
```

### 1.2 Model

我们现在article_model.go文件中，添加文章的查询，而且要分页查询：

```go

//-----------查询文章---------

//根据页码查询文章
func FindArticleWithPage(page int) ([]Article, error) {
	//从配置文件中获取每页的文章数量
	num, _ := beego.AppConfig.Int("articleListPageNum")
	page--
	fmt.Println("---------->page", page)
	return QueryArticleWithPage(page, num)
}

/**
分页查询数据库
limit分页查询语句，
	语法：limit m，n

	m代表从多少位开始获取，与id值无关
	n代表获取多少条数据

注意limit前面咩有where
 */
func QueryArticleWithPage(page, num int) ([]Article, error) {
	sql := fmt.Sprintf("limit %d,%d", page*num, num)
	return QueryArticlesWithCon(sql)
}

func QueryArticlesWithCon(sql string) ([]Article, error) {
	sql = "select id,title,tags,short,content,author,createtime from article " + sql
	rows, err := utils.QueryDB(sql)
	if err != nil {
		return nil, err
	}
	var artList []Article
	for rows.Next() {
		id := 0
		title := ""
		tags := ""
		short := ""
		content := ""
		author := ""
		var createtime int64
		createtime = 0
		rows.Scan(&id, &title, &tags, &short, &content, &author, &createtime)
		art := Article{id, title, tags, short, content, author, createtime}
		artList = append(artList, art)
	}
	return artList, nil
}

```

在models目录下创建一个go文件，用来控制首页显示内容：

```go
type HomeBlockParam struct {
	Id         int
	Title      string
	Tags       [] TagLink
	Short      string
	Content    string
	Author     string
	CreateTime string
	//查看文章的地址
	Link string

	//修改文章的地址
	UpdateLink string
	DeleteLink string

	//记录是否登录
	IsLogin bool
}

type TagLink struct {
	TagName string
	TagUrl  string
}
```

我们需要将从数据库中查询出来的数据，转为对应的结构体对象，所以先设计结构体，这里我们需要考虑如果用户是登录状态，那么是可以修改或删除某一篇文章。当然，如果没有登录，那么只能查看。所以在设计结构体的时候，我们直接创建了修改和删除的链接字段。

接下来，我们添加一个方法，用于将文章中的内容，显示到页面上：

```go
//----------首页显示内容---------
func MakeHomeBlocks(articles []Article, isLogin bool) template.HTML {
	htmlHome := ""
	for _, art := range articles {
		//将数据库model转换为首页模板所需要的model
		homeParam := HomeBlockParam{}
		homeParam.Id = art.Id
		homeParam.Title = art.Title
		homeParam.Tags = createTagsLinks(art.Tags)
		fmt.Println("tag-->", art.Tags)
		homeParam.Short = art.Short
		homeParam.Content = art.Content
		homeParam.Author = art.Author
		homeParam.CreateTime = utils.SwitchTimeStampToData(art.Createtime)
		homeParam.Link = "/article/" + strconv.Itoa(art.Id)
		homeParam.UpdateLink = "/article/update?id=" + strconv.Itoa(art.Id)
		homeParam.DeleteLink = "/article/delete?id=" + strconv.Itoa(art.Id)
		homeParam.IsLogin = isLogin

		//处理变量
		//ParseFile解析该文件，用于插入变量
		t, _ := template.ParseFiles("views/block/home_block.html")
		buffer := bytes.Buffer{}
		//就是将html文件里面的比那两替换为穿进去的数据
		t.Execute(&buffer, homeParam)
		htmlHome += buffer.String()
	}
	fmt.Println("htmlHome-->",htmlHome)
	return template.HTML(htmlHome)
}
```

额外还需要一个方法：

```go

//将tags字符串转化成首页模板所需要的数据结构
func createTagsLinks(tags string) []TagLink {
	var tagLink [] TagLink
	tagsPamar := strings.Split(tags, "&")
	for _, tag := range tagsPamar {
		tagLink = append(tagLink, TagLink{tag, "/?tag=" + tag})
	}
	return tagLink
}
```


### 1.3 View

接下来我们设计一下页面，刚刚在model的MakeHomeBlocks()方法中，就是需要使用模板填充格式化html页面内容，所以我们在views/block下再创建一个html页面：home_block.html，内容如下：

```html
<div id="home-block-item">
    <h2><a href="{{.Link}}">{{.Title}}</a></h2>
    <div>
        <span>{{.CreateTime}}</span>
        <span>
        {{range .Tags}}
            <a href="{{.TagUrl}}">&nbsp{{.TagName}}</a>
        {{end}}
        </span>
    </div>
    <p><a href={{.Link}}>{{.Short}}</a></p>
{{if .IsLogin}}
    <div class="home-block-item-udpate">
        <a href='javascript:if(confirm("确定删除吗？")){location="{{.DeleteLink}}"}'>删除</a>
        <a href={{.UpdateLink}}>修改</a>
    </div>
{{end}}
</div>
```

我们现实了从数据中查询出的文章的数据，如果用户是登录状态，那么我们现实删除和修改，因为用户有这两个权限，否则就不显示。


### 1.4 运行

我们在数据库中插入10条数据：
![http://7xtcwd.com1.z0.glb.clouddn.com/WX20190522-113706@2x.png](http://7xtcwd.com1.z0.glb.clouddn.com/WX20190522-113706@2x.png)

接下来我们设置一下配置文件，每页显示6条(也可以8条，10条。。)，

修改conf目录下的app.conf文件：

```
appname = myblog
httpport = 8080
runmode = dev

#mysql配置
driverName = mysql
mysqluser = root
mysqlpwd = yu271400
host = 127.0.0.1
port = 3306
dbname = myblog

#Session
sessionon = true
sessionprovider = "file"
sessionname = "qianfengjiaoyumyblog"
sessiongcmaxlifetime = 1800
sessionproviderconfig = "./tmp"
sessioncookielifetime = 1800

articleListPageNum = 6
```

然后启动项目，打开浏览器输入网址：[http://127.0.0.1:8080/](http://127.0.0.1:8080/)

![http://7xtcwd.com1.z0.glb.clouddn.com/WX20190522-141320@2x.png](http://7xtcwd.com1.z0.glb.clouddn.com/WX20190522-141320@2x.png)

用户虽然没有登录，但是也是可以查看的，接下来我们点击登录按钮进行登录：

![http://7xtcwd.com1.z0.glb.clouddn.com/WX20190522-141446@2x.png](http://7xtcwd.com1.z0.glb.clouddn.com/WX20190522-141446@2x.png)

登录后，用户就可以有删除和修改的功能了。

## 二、分页功能

截止，我们已经已经能够显示出第一页的内容了，接下来我们添加上一页和下一页的功能。

### 2.1 Model

首先在home_model.go中添加一个分页的结构体对象：

```go
type HomeFooterPageCode struct {
	HasPre bool
	HasNext bool
	ShowPage string
	PreLink string
	NextLink string
}
```

接下来添加方法：

```go


//-----------翻页-----------
//page是当前的页数
func ConfigHomeFooterPageCode(page int) HomeFooterPageCode {
	pageCode := HomeFooterPageCode{}
	//查询出总的条数
	num := GetArticleRowsNum()
	//从配置文件中读取每页显示的条数
	pageRow, _ := beego.AppConfig.Int("articleListPageNum")
	//计算出总页数
	allPageNum := (num-1)/pageRow + 1

	pageCode.ShowPage = fmt.Sprintf("%d/%d", page, allPageNum)

	//当前页数小于等于1，那么上一页的按钮不能点击
	if page <= 1 {
		pageCode.HasPre = false
	} else {
		pageCode.HasPre = true
	}

	//当前页数大于等于总页数，那么下一页的按钮不能点击
	if page >= allPageNum {
		pageCode.HasNext = false
	} else {
		pageCode.HasNext = true
	}

	pageCode.PreLink = "/?page=" + strconv.Itoa(page-1)
	pageCode.NextLink = "/?page=" + strconv.Itoa(page+1)
	return pageCode

}

```



这段代码需要查询出数据库中所有文章的总量，所以我们要先在article_model.go文件中，加入查询总数据量的方法：

```go

//------翻页------

//存储表的行数，只有自己可以更改，当文章新增或者删除时需要更新这个值
var artcileRowsNum = 0

//只有首次获取行数的时候采取统计表里的行数
func GetArticleRowsNum() int {
	if artcileRowsNum == 0 {
		artcileRowsNum = QueryArticleRowNum()
	}
	return artcileRowsNum
}

//查询文章的总条数
func QueryArticleRowNum() int {
	row := utils.QueryRowDB("select count(id) from article")
	num := 0
	row.Scan(&num)
	return num
}

```



我们还要考虑一个问题，就是当新增或删除文章的时候，数据总量会发生改变，所以还要修改增加文章的方法：

先新增一个方法用于设置总页数：

```go
//设置页数
func SetArticleRowsNum(){
	artcileRowsNum = QueryArticleRowNum()
}
```

然后修改增加文章的方法：

```go
//---------添加文章-----------
func AddArticle(article Article) (int64, error) {
	i, err := insertArticle(article)
	SetArticleRowsNum()
	return i, err
}
```



### 2.2 Controller

修改home_controller.go的Get()方法：

```go
func (this *HomeController) Get() {
	...
	artList, _ = models.FindArticleWithPage(page)
	this.Data["PageCode"] = models.ConfigHomeFooterPageCode(page)
	this.Data["HasFooter"] = true
	...
}

```



### 2.3 View

修改home.html页面：

```html
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <title>首页</title>
    <link href="../static/css/blogsheet.css" rel="stylesheet">
</head>
<body>
{{template "block/nav.html" .}}


<div id="main">
{{.Content}}



{{if .HasFooter}}
    <div id="home-footer">
        <a {{if .PageCode.HasPre}}href="{{.PageCode.PreLink}}" {{else}} class="disable" {{end}}>上一页</a>
        <span>{{.PageCode.ShowPage}}页</span>
        <a {{if .PageCode.HasNext}}href="{{.PageCode.NextLink}}" {{else}} class="disable" {{end}}>下一页</a>
    </div>
{{end}}


</div>

</body>
</html>
```



增添上一页，下一页的链接。



### 2.4 运行

首先我们再向数据库中插入5条数据：

![http://7xtcwd.com1.z0.glb.clouddn.com/WX20190522-141654@2x.png](http://7xtcwd.com1.z0.glb.clouddn.com/WX20190522-141654@2x.png)

然后修改配置文件，每页显示5条。

![http://7xtcwd.com1.z0.glb.clouddn.com/WX20190522-141839@2x.png](http://7xtcwd.com1.z0.glb.clouddn.com/WX20190522-141839@2x.png)

接下来，我们新增加一篇文章，点击写博客：
![http://7xtcwd.com1.z0.glb.clouddn.com/WX20190522-142111@2x.png](http://7xtcwd.com1.z0.glb.clouddn.com/WX20190522-142111@2x.png)

最后一页显示为我们刚刚增加的一篇文章，截止到现在我们可以很完美的显示页码了。




    
## 文章详情功能开发
**@author：Davie**
**版权所有：北京千锋互联科技有限公司**

上节课内容我们实现了项目主页的文章列表功能和分页功能。本节课继续进行开发实现，当点击文章的时候，应该显示文章的详细内容。

## 一、Controller

首先先设置路由：

```go
func init() {
    ...

	//写文章
	beego.Router("/article/add", &controllers.AddArticleController{})
	//显示文章内容
	beego.Router("/article/:id", &controllers.ShowArticleController{})

}
```


然后再controllers目录下创建一个go文件，show_article_controller.go：

```go
type ShowArticleController struct {
	//beego.Controller
	BaseController
}

func (this *ShowArticleController) Get() {
	idStr := this.Ctx.Input.Param(":id")
	id, _ := strconv.Atoi(idStr)
	fmt.Println("id:", id)


	//获取id所对应的文章信息
	art := models.QueryArticleWithId(id)

	this.Data["Title"] = art.Title
	this.Data["Content"] = art.Content
	//this.Data["Content"] = utils.SwitchMarkdownToHtml(art.Content)
	this.TplName="show_article.html"
}
```

## 二、Model

接下来在article_model.go文件中，添加方法，根据id查询文章：

```go

//----------查询文章-------------

func QueryArticleWithId(id int) Article {
	row := utils.QueryRowDB("select id,title,tags,short,content,author,createtime from article where id=" + strconv.Itoa(id))
	title := ""
	tags := ""
	short := ""
	content := ""
	author := ""
	var createtime int64
	createtime = 0
	row.Scan(&id, &title, &tags, &short, &content, &author, &createtime)
	art := Article{id, title, tags, short, content, author, createtime}
	return art
}
```

## 三、View

接下来我们创建view，在views目录下，新建html页面文件，show_article.html：

```html
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <title>{{.Title}}</title>
    <link href="../static/css/blogsheet.css" rel="stylesheet">
</head>
<body>

{{template "block/nav.html" .}}

<div id="main">
    <h1>{{.Title}}</h1>
    <div>{{.Content}}</div>
</div>

</body>
</html>
```

## 四、运行

接下来我们重启项目，并刷新页面：

![http://7xtcwd.com1.z0.glb.clouddn.com/WX20190522-143748@2x.png](http://7xtcwd.com1.z0.glb.clouddn.com/WX20190522-143748@2x.png)

点击一篇文章：

![http://7xtcwd.com1.z0.glb.clouddn.com/WX20190522-144626@2x.png](http://7xtcwd.com1.z0.glb.clouddn.com/WX20190522-144626@2x.png)

## 五、Markdown

虽然页面能够显示文章内容，但是看着很不舒服，一锅粥一样，我们通过markdown语法格式显示。

我们先了解一下，在进行项目开发前先熟悉下我们需要用到的库。包括：

- 转换Markdown语法：russross/blackfriday
- 查找Document的内容：PuerkitoBio/goquery
- 语法高亮：sourcegraph/syntaxhighlight
- 插入模块：html/template
- 执行外部命令：os/exec
- 文件操作：path/filepath
- 创建Web服务器：SimpleHTTPServer
- 解析.yml配置文件：gopkg.in/yaml.v2

首先需要安装markdown的安装包：

打开终端输入以下命令：

```shell
go get github.com/russross/blackfriday
go get github.com/PuerkitoBio/goquery
go get github.com/sourcegraph/syntaxhighlight
```

安装之后也可以到src目录下查看:
![http://7xtcwd.com1.z0.glb.clouddn.com/WX20190522-144232@2x.png](http://7xtcwd.com1.z0.glb.clouddn.com/WX20190522-144232@2x.png)

### 5.1 语法简介

#### 5.1.1 russross/blackfriday包

第三方库russross/blackfriday用于在golang中使用markdown语法。

**markdown**: 是一种可以使用普通文本编辑器编写的标记语言，通过简单的标记语法，它可以使普通文本内容具有一定的格式。 

Markdown具有一系列衍生版本，用于扩展Markdown的功能（如表格、脚注、内嵌HTML等等），这些功能原初的Markdown尚不具备，它们能让Markdown转换成更多的格式，例如LaTeX，Docbook。Markdown增强版中比较有名的有Markdown Extra、MultiMarkdown、 Maruku等。这些衍生版本要么基于工具，如Pandoc；要么基于网站，如GitHub和Wikipedia，在语法上基本兼容，但在一些语法和渲染效果上有改动。

test.md：

```markdown
## 一、russross/blackfriday包
```

示例代码：

```go
func main() {
	fileread, _ := ioutil.ReadFile("extra/blackfriday转换markdown/test.md")
	//转换Markdown语法，如将"#"转换为"<h1></h1>"
	subHtml := blackfriday.MarkdownCommon(fileread)
	subHtmlStr := string(subHtml)
	fmt.Println(subHtmlStr)
}
```

输出：

```
<h2>一、russross/blackfriday包</h2>
```

#### 5.1.2 PuerkitoBio/goquery包

做过 Web 开发的，应该都用过或听过 jQuery，它提供了方便的操作 DOM 的 API。使用 Go 语言做服务器端开发，有时候需要解析 HTML 文件，比如抓取网站内容、写一个爬虫等。这时候如果有一个类似 jQuery 的库可以使用，操作 DOM 会很方便，而且，上手也会很快。`PuerkitoBio/goquery` 这个库就实现了类似 jQuery 的功能，让你能方便的使用 Go 语言操作 HTML 文档。

该库提供的类型很少，但方法却很多，我们不可能一个个方法讲解。这里通过模拟几个使用场景来讲解该库的使用。

##### 1. Document

Document 代表一个将要被操作的 HTML 文档，不过，和 jQuery 不同，它装载的是 DOM 文档的一部分。

```go
type Document struct {
	*Selection
	Url *url.URL
	rootNode *html.Node // 文档的根节点
}
```

因为 Document 中内嵌了一个 Selection 类型，因此，Document 可以直接使用 Selection 类型的方法。提供有五种方法获取一个 Document 实例。

##### 2. Selection

Selection 代表符合特定条件的节点集合。

```go
type Selection struct {
	Nodes []*html.Node
	document *Document
	prevSel *Selection
}
```

一般地，得到了 Document 实例后，通过 Dcoument.Find 方法获取一个 Selection 实例，然后像 jQuery 一样使用链式语法和方法操作它。

Selection 类型提供的方法可以分为如下几大类

- 类似函数的位置操作
- 扩大 Selection 集合（增加选择的节点）
- 过滤方法，减少节点集合
- 循环遍历选择的节点
- 修改文档
- 检测或获取节点属性值
- 查询或显示一个节点的身份
- 在文档树之间来回跳转（常用的查找节点方法）

##### 3. 基本用法：

1、创建文档

```go
1. d,e := goquery.NewDocumentFromReader(reader io.Reader)
2. d,e := goquery.NewDocument(url string)
```

2、查找内容

```go
1. ele.Find("#title") //根据id查找
2. ele.Find(".title") //根据class查找
3. ele.Find("h2").Find("a") //链式调用
```

3、获取内容

```go
1. ele.Html()
2. ele.Text()
```

4、获取属性

```go
1. ele.Attr(“href”)
2. ele.AttrOr(“href”, “”)
```

5、遍历

```go
1. ele.Find(“.item”).Each(func(index int, ele *goquery.Selection){
2. 
3. })
```

示例：

```go
func main() {
	doc, err := goquery.NewDocument("http://studygolang.com/topics")
	if err != nil {
		log.Fatal(err)
	}

	doc.Find(".topic").Each(func(i int, contentSelection *goquery.Selection) {
		title := contentSelection.Find(".title a").Text()
        //Find(".title a")与Find(".title").Find("a")一样
		fmt.Println("第", i+1, "个帖子的标题：", title)
		//ret,_ := contentSelection.Html()
		//fmt.Printf("\n\n\n%v", ret)
		//fmt.Println(contentSelection.Text())
	})
    //最终输出为 html 文档：
	//new, err := doc.Html()
}
```

其中Find中的输入字符串是CSS selector，其语法风格参照 http://www.w3school.com.cn/cssref/css_selectors.asp 。如：

| 语法            | 表述                                              |
| --------------- | ------------------------------------------------- |
| #firstname      | 选择 id="firstname" 的所有元素。                  |
| *               | 选择所有元素。                                    |
| p               | 选择所有 <p> 元素。                               |
| div,p           | 选择所有 <div> 元素和所有 <p> 元素。              |
| div p           | 选择 <div> 元素内部的所有 <p> 元素。              |
| div>p           | 选择父元素为 <div> 元素的所有 <p> 元素。          |
| div+p           | 选择紧接在 <div> 元素之后的所有 <p> 元素。        |
| [target]        | 选择带有 target 属性所有元素。                    |
| [target=_blank] | 选择 target="_blank" 的所有元素。                 |
| a[src*=”abc”]   | 选择其 src 属性中包含 “abc” 子串的每个 <a> 元素。 |
| a[src$=”.pdf”]  | 选择其 src 属性以 “.pdf” 结尾的所有 <a> 元素。    |

#### 5.1.3 sourcegraph/syntaxhighlight包

syntaxhighlight包提供代码的语法高亮显示。 它目前使用独立于语言的词法分析器，
并在JavaScript，Java，Ruby，Python，Go和C上表现出色。

主要的AsHTML(src []byte) ([]byte, error)函数，输出就是HTML
与google-code-prettify相同的CSS类，因此任何样式表也应该适用于此包。

```go
func main() {
   src := []byte(`
/* hello, world! */
var a = 3;

// b is a cool function
function b() {
  return 7;
}`)

   highlighted, err := syntaxhighlight.AsHTML(src)
   if err != nil {
      fmt.Println(err)
      os.Exit(1)
   }

   fmt.Println(string(highlighted))
}
```

输出

```
<span class="com">/* hello, world! */</span>
<span class="kwd">var</span> <span class="pln">a</span> <span class="pun">=</span> <span class="dec">3</span><span class="pun">;</span>

<span class="com">// b is a cool function</span>
<span class="kwd">function</span> <span class="pln">b</span><span class="pun">(</span><span class="pun">)</span> <span class="pun">{</span>
  <span class="kwd">return</span> <span class="dec">7</span><span class="pun">;</span>
<span class="pun">}</span>
```

通过如下规则就行的转换

```go
var DefaultHTMLConfig = HTMLConfig{
	String:        "str",
	Keyword:       "kwd",
	Comment:       "com",
	Type:          "typ",
	Literal:       "lit",
	Punctuation:   "pun",
	Plaintext:     "pln",
	Tag:           "tag",
	HTMLTag:       "htm",
	HTMLAttrName:  "atn",
	HTMLAttrValue: "atv",
	Decimal:       "dec",
	Whitespace:    "",
}
```





### 5.2 修改项目代码

首先在show_article.html页面上导入样式包：

```html
<!DOCTYPE html>
<html lang="en">
<head>
    ...
    <link href="../static/css/lib/highlight.css" rel="stylesheet">
</head>

```



接下来在utils目录下，myUtils.go文件中添加方法：

```go
func SwitchMarkdownToHtml(content string) template.HTML {
	markdown := blackfriday.MarkdownCommon([]byte(content))

	//获取到html文档
	doc, _ := goquery.NewDocumentFromReader(bytes.NewReader(markdown))
	/**
	对document进程查询，选择器和css的语法一样
	第一个参数：i是查询到的第几个元素
	第二个参数：selection就是查询到的元素
	 */
	doc.Find("code").Each(func(i int, selection *goquery.Selection) {
		light, _ := syntaxhighlight.AsHTML([]byte(selection.Text()))
		selection.SetHtml(string(light))
		fmt.Println(selection.Html())
		fmt.Println("light:", string(light))
		fmt.Println("\n\n\n")
	})
	htmlString, _ := doc.Html()
	return template.HTML(htmlString)
}
```



最后修改controller中的Get()方法：

```go
func (this *ShowArticleController) Get() {
	...
	this.Data["Title"] = art.Title
	//this.Data["Content"] = art.Content
	this.Data["Content"] = utils.SwitchMarkdownToHtml(art.Content)
	this.TplName="show_article.html"
}
```

### 5.3 运行结果

重启项目后，然后刷新页面：

![http://7xtcwd.com1.z0.glb.clouddn.com/WX20190522-143848@2x.png](http://7xtcwd.com1.z0.glb.clouddn.com/WX20190522-143848@2x.png)

可以看到，效果就比较好看了，这就实现了我们对页面的美化效果。



## 修改文章功能开发
**@author：Davie**
**版权所有：北京千锋互联科技有限公司**

上节课内容我们学习了查看文章详情的功能开发。本节课我们事先修改文章的功能。

## 一、Controller

首先我们先添加一个路由：

```go
func init() {
    ...
	//写文章
	beego.Router("/article/add", &controllers.AddArticleController{})
	//显示文章内容
	beego.Router("/article/:id", &controllers.ShowArticleController{})
	//更新文章
	beego.Router("/article/update", &controllers.UpdateArticleController{})
}
```

然后新创建一个Controller(update_article_controller.go)，用于控制更新操作：

```go
type UpdateArticleController struct {
	BaseController
}

//当访问/update路径的时候回触发Get()方法，响应的页面是通过TplName这个属性指定返回给客户端的页面
func (this *UpdateArticleController) Get() {
	id, _ := this.GetInt("id")
	fmt.Println(id)

	//获取id所对应的文章信息
	art := models.QueryArticleWithId(id)
	this.Data["Title"] = art.Title
	this.Data["Tags"] = art.Tags
	this.Data["Short"] = art.Short
	this.Data["Content"] = art.Content
	this.Data["Id"] = art.Id
	this.TplName = "write_article.html"

}
```

当点击修改按钮触发路由之后，通过get请求，进入Get()方法，其实还是现实写文章的页面，只不过里面已经填好了要修改的数据。

要想在页面上显示要修改的文章的数据，就需要根据id查询这篇文章，然后将数据传入到前端页面。

当用户点击提交按钮，触发post请求，进入Post()方法：

```go
//修改文章
func (this *UpdateArticleController) Post() {
	id, _ := this.GetInt("id")
	fmt.Println("postid:", id)

	//获取浏览器传输的数据，通过表单的name属性获取值
	title := this.GetString("title")
	tags := this.GetString("tags")
	short := this.GetString("short")
	content := this.GetString("content")

	//实例化model，修改数据库
	art := models.Article{id, title, tags, short, content, "", 0}
	_, err := models.UpdateArticle(art)

	//返回数据给浏览器
	if err == nil {
		this.Data["json"] = map[string]interface{}{"code": 1, "message": "更新成功"}
	} else {
		this.Data["json"] = map[string]interface{}{"code": 0, "message": "更新失败"}
	}
	this.ServeJSON()
}
```

## 二、View

修改文章和写文章，我们用同一个页面即可，只不过需要修改write_article.html页面的数据：

```html
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <title>写文章</title>
    <link href="../static/css/blogsheet.css" rel="stylesheet">
    <script src="../static/js/lib/jquery-3.3.1.min.js"></script>
    <script src="../static/js/lib/jquery.url.js"></script>
    <script src="../static/js/blog.js"></script>
</head>
<body>

{{template "block/nav.html" .}}

<div id="main">
    <form id="write-art-form" method="post">
        <div>标题</div>
        <input type="text" placeholder="请输入标题" name="title" value="{{.Title}}">
        <div>标签</div>
        <input type="text" placeholder="请输入标签" name="tags" value="{{.Tags}}">
        <div>简介</div>
        <textarea placeholder="请输入简介" name="short">{{.Short}}</textarea>
        <div>内容</div>
        <textarea id="content" placeholder="请输入内容" name="content">{{.Content}}</textarea>
        <input id="write-article-id" hidden name="id" value="{{.Id}}">
        <button type="button" onclick="history.back()">返回</button>
        <button type="submit" id="write-art-submit">提交</button>
    </form>
</div>

</body>
</html>
```

之前点击提交按钮，是为了添加文章。现在也有可能是修改文章。到底是修改还是添加，其实要看页面上id的值，如果是添加使用默认值0，如果是修改是从数据库中查询出的真实数值，所以需要修改js脚本文件：

```js
submitHandler: function (form) {
    alert("hello")
    var urlStr = "/article/add";
    //判断文章id确定提交的表单的服务器地址
    //若id大于零，说明是修改文章
    var artId = $("#write-article-id").val();
    alert("artId:" + artId);
    if (artId > 0) {
        urlStr = "/article/update"
    }
    alert("urlStr:" + urlStr);
    $(form).ajaxSubmit({
        url: urlStr,
        type: "post",
        dataType: "json",
        success: function (data, status) {
            alert(":data:" + data.message);
            setTimeout(function () {
                window.location.href = "/"
            }, 1000)
        },
        error: function (data, status) {
            alert("err:" + data.message + ":" + status)
        }
    });
```

## 三、Model

然后我们就可以去修改数据库了，在article_model.go文件中，添加方法：

```go
//----------修改数据----------

func UpdateArticle(article Article) (int64, error) {
	//数据库操作
	return utils.ModifyDB("update article set title=?,tags=?,short=?,content=? where id=?",
		article.Title, article.Tags, article.Short, article.Content, article.Id)
}
```

## 四、运行

重启项目后，我们打来浏览器输入网址：[http://127.0.0.1:8080/](http://127.0.0.1:8080/)

![http://7xtcwd.com1.z0.glb.clouddn.com/WX20190522-155033@2x.png](http://7xtcwd.com1.z0.glb.clouddn.com/WX20190522-155033@2x.png)

我们就修改第一篇文章：

![http://7xtcwd.com1.z0.glb.clouddn.com/WX20190522-155245@2x.png](http://7xtcwd.com1.z0.glb.clouddn.com/WX20190522-155245@2x.png)

提交之后：

![http://7xtcwd.com1.z0.glb.clouddn.com/WX20190522-155431@2x.png](http://7xtcwd.com1.z0.glb.clouddn.com/WX20190522-155431@2x.png)

也可以查询数据库，也已经更新完毕：

![http://7xtcwd.com1.z0.glb.clouddn.com/WX20190522-155627@2x.png](http://7xtcwd.com1.z0.glb.clouddn.com/WX20190522-155627@2x.png)



## 删除文章功能开发
**@author：Davie**
**版权所有：北京千锋互联科技有限公司**

当用户登录后，是可以有删除功能的，点击删除，就可以删除一篇文章。删除文章功能不需要view页面，直接弹框确认即可，现在我们来实现model和controller部分。

### 一、Model

在article_model.go文件中，添加 删除文章的方法：

```go
//----------删除文章---------
func DeleteArticle(artID int) (int64, error) {
	i, err := deleteArticleWithArtId(artID)
	SetArticleRowsNum()
	return i, err
}

func deleteArticleWithArtId(artID int) (int64, error) {
	return utils.ModifyDB("delete from article where id=?", artID)
}
```

删除的时候，肯定是根据文章的ID来删除，当点击删除按钮的时候，可以传过来文章ID。

```html
{{if .IsLogin}}
    <div class="home-block-item-udpate">
        <a href='javascript:if(confirm("确定删除吗？")){location="{{.DeleteLink}}}"}'>删除</a>
        <a href={{.UpdateLink}}>修改</a>
    </div>
{{end}}
```

### 二、Controller

接下来，我们实现controller部分，在controllers目录下，创建一个新的go文件：delete_article_controller.go

```go
package controllers

import (
	"fmt"
	"myblogweb/models"
	"log"
)

type DeleteArticleController struct {
	BaseController
}

//点击删除后重定向到首页
func (this *DeleteArticleController) Get() {
	artID, _ := this.GetInt("id")
	fmt.Println("删除 id:", artID)

	_, err := models.DeleteArticle(artID)
	if err != nil {
		log.Println(err)
	}
	this.Redirect("/", 302)
}
```

先接收要删除的文章ID，然后操作sql语句进行删除数据，自后重定向到首页上。接下来注册个新的路由：

```go
// 删除文章
beego.Router("/article/delete", &controllers.DeleteArticleController{})
```

### 三、运行

重启项目后，打开浏览器输入：[http://127.0.0.1:8080/](http://127.0.0.1:8080/)，并且要登录，才能有删除功能。

![http://7xtcwd.com1.z0.glb.clouddn.com/WX20190522-161409@2x.png](http://7xtcwd.com1.z0.glb.clouddn.com/WX20190522-161409@2x.png)

我们删除第一篇文章，点击删除按钮。然后弹出确认框，选择确定。就会删除第一篇文章。然后跳转到首页。

![http://7xtcwd.com1.z0.glb.clouddn.com/WX20190522-161420@2x.png](http://7xtcwd.com1.z0.glb.clouddn.com/WX20190522-161420@2x.png)

我们发现首页已经没有第一篇文章了，此时也可以查询一下数据库：

![http://7xtcwd.com1.z0.glb.clouddn.com/WX20190522-161457@2x.png](http://7xtcwd.com1.z0.glb.clouddn.com/WX20190522-161457@2x.png)

我们发现数据库中id为1的文章已经被删除。


## 标签功能开发
**@author：Davie**
**版权所有：北京千锋互联科技有限公司**

上节课内容我们实现了删除文章的操作，本节课我们继续进行相关功能开发。当我们点击标签按钮的时候，应该要显示每个标签对应了多少篇文章。

## 一、Model

当点击标签的时候，需要查询出数据库中所有的标签，然后再统计出每个标签的文章总量，显示到前端页面上。

所以我们应该先查询出数据库中所有的标签。

在article_model.go文件中，加入以下代码，先查询出所有的标签。

```go
//查询标签，返回一个字段的列表
func QueryArticleWithParam(param string) []string {
	rows, err := utils.QueryDB(fmt.Sprintf("select %s from article", param))
	if err != nil {
		log.Println(err)
	}
	var paramList []string
	for rows.Next() {
		arg := ""
		rows.Scan(&arg)
		paramList = append(paramList, arg)
	}
	return paramList
}
```

然后再models目录下，创建一个新的model文件：tags_model.go
```go
package models

import "strings"

func HandleTagsListData(tags []string) map[string]int {
	var tagsMap = make(map[string]int)
	for _, tag := range tags {
		tagList := strings.Split(tag, "&")
		for _, value := range tagList {
			tagsMap[value]++
		}
	}
	return tagsMap
}
```

## 二、Controller

接下来，我们在controllers目录下新建一个controller的go文件，tags_controller.go。

```go
package controllers

import (
	"myblogweb/models"
	"fmt"
)

type TagsController struct {
	BaseController
}

func (this *TagsController) Get() {
	tags := models.QueryArticleWithParam("tags")
	fmt.Println(models.HandleTagsListData(tags))
	this.Data["Tags"] = models.HandleTagsListData(tags)
	this.TplName = "tags.html"
}
```

并且在router.go中注册新的路由：
```go
//标签
beego.Router("/tags", &controllers.TagsController{})
```

## 三、Views

最后我们去写前端页面，来接收数据。

在views包下，新建一个html文件，tags.html：

```html
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <title>标签</title>
    <link href="../static/css/blogsheet.css" rel="stylesheet">
</head>
<body>
{{template "block/nav.html" .}}
<div id="main">
    <h1>标签</h1>
    <div id="tags-list">
    {{range $k,$v := .Tags}}
        <div><a href="/?tag={{$k}}"><span class="global-color">{{$k}}</span> 有{{$v}}篇文章</a></div>{{end}}
    </div>
</div>
</body>
</html>
```

## 四、运行

重新启动项目后，在浏览器中点击标签按钮：

![http://7xtcwd.com1.z0.glb.clouddn.com/WX20190522-165548@2x.png](http://7xtcwd.com1.z0.glb.clouddn.com/WX20190522-165548@2x.png)



## 首页功能扩展开发
**@author：Davie**
**版权所有：北京千锋互联科技有限公司**

我们最初只是让首页显示第一页的文章数据，如果在访问首页的时候，后面传入相应的参数，那么我们应该根据参数显示对应的内容。

如果浏览器输入首页网址：[http://127.0.0.1:8080/](http://127.0.0.1:8080/)，那么我们显示第一页的文章数据

如果浏览器输入网址是：[http://127.0.0.1:8080/?page=3](http://127.0.0.1:8080/?page=3)，那么我们直接显示第三页的文章数据

如果浏览器输入网址是：[http://127.0.0.1:8080/?tag=go](http://127.0.0.1:8080/?tag=go)，那么我们应该显示所有标签为go的文章，同理如果tag=web，我们应该显示所有标签为web的文章。

接下来我们就去改进首页的功能。

### 一、Controller

我们首先修改home_controller.go文件中的处理，修改代码如下：

```go
package controllers

import (
	"fmt"
	"myblogweb/models"
)

type HomeController struct {
	//beego.Controller
	BaseController
}

func (this *HomeController) Get() {

	tag := this.GetString("tag")
	fmt.Println("tag:", tag)
	page, _ := this.GetInt("page")
	var artList []models.Article

	if len(tag) > 0 {
		//按照指定的标签搜索
		artList, _ = models.QueryArticlesWithTag(tag)
		this.Data["HasFooter"] = false

	} else {
		if page <= 0 {
			page = 1
		}
		artList, _ = models.FindArticleWithPage(page)
		this.Data["PageCode"] = models.ConfigHomeFooterPageCode(page)
		this.Data["HasFooter"] = true
	}

	fmt.Println("IsLogin:", this.IsLogin, this.Loginuser)
	this.Data["Content"] = models.MakeHomeBlocks(artList, this.IsLogin)

	this.TplName = "home.html"
}

```

我们先接收tag的值和page的值：

那么有三种情况：

 	1. 如果tag有值，page就不会有值。比如：[http://127.0.0.1:8080/?pag=web](http://127.0.0.1:8080/?pag=web)
 	2. 如果page有值，那么tag就不会有值。比如：[http://127.0.0.1:8080/?page=3](http://127.0.0.1:8080/?page=3)
 	3. 就是用户直接访问首页，没有传入tag也没有传入page：[http://127.0.0.1:8080](http://127.0.0.1:8080)

如果是第一种情况，那么我们应该按照tag进行搜索。如果是第二种情况就按照页数查询。如果是第三种就默认按照分页查询，只是查询第一页。

### 二、model
接下来我们去实现按照tag进行查询，这里有个复杂的地方，就是我们在写文章的手，标签可能有一个，也肯能有多个，如果是多个，我们是按照&连接的，比如：go&socket。那么在按照标签查询的时候，就要分为4种情况：

```
通过标签查询首页的数据
有四种情况
   1.左右两边有&符和其他符号
   2.左边有&符号和其他符号，同时右边没有任何符号
   3.右边有&符号和其他符号，同时左边没有任何符号
   4.左右两边都没有符号

通过%去匹配任意多个字符，至少是一个
```

我们在article_model.go文件中，添加一个方法：

```go
//--------------按照标签查询--------------
func QueryArticlesWithTag(tag string) ([]Article, error) {

	sql := " where tags like '%&" + tag + "&%'"
	sql += " or tags like '%&" + tag + "'"
	sql += " or tags like '" + tag + "&%'"
	sql += " or tags like '" + tag + "'"
	fmt.Println(sql)
	return QueryArticlesWithCon(sql)
}
```

### 三、运行
前端页面不用修改，我们打开浏览器输入以下网址：[http://127.0.0.1:8080?tag=web](http://127.0.0.1:8080?tag=web)

![http://7xtcwd.com1.z0.glb.clouddn.com/WX20190524-103729@2x.png](http://7xtcwd.com1.z0.glb.clouddn.com/WX20190524-103729@2x.png)

可以看到查询出web标签的三篇文章，重新输入网址：[http://127.0.0.1:8080?tag=http](http://127.0.0.1:8080?tag=http)

![http://7xtcwd.com1.z0.glb.clouddn.com/WX20190524-104605@2x.png](http://7xtcwd.com1.z0.glb.clouddn.com/WX20190524-104605@2x.png)

socket标签下的两篇文章也被查询出来了，下面我们输入以下页码：[http://127.0.0.1:8080?page=3](http://127.0.0.1:8080?page=3)

![http://7xtcwd.com1.z0.glb.clouddn.com/WX20190524-104754@2x.png](http://7xtcwd.com1.z0.glb.clouddn.com/WX20190524-104754@2x.png)


## 文件上传和图片展示功能
**@author：Davie**
**版权所有：北京千锋互联科技有限公司**

上节内容我们学习了首页扩功能，本节课我们继续开发实现新的功能--文件上传和图片展示功能。在实际案例中，提供有相册展示功能，相册功能展示的图片需要用到文件上传，我们来实际开发实现。

### 一、创建数据表
首先我们要在数据库中创建数据表，要想存储图片，我们需要图片的id，图片的路径位置，图片的名字等等。所在在utils工具包下的mysqlUtil.go文件中，添加创建数据表的方法：

```go

//--------图片--------
func CreateTableWithAlbum() {
	sql := `create table if not exists album(
		id int(4) primary key auto_increment not null,
		filepath varchar(255),
		filename varchar(64),
		status int(4),
		createtime int(10)
		);`
	ModifyDB(sql)
}
```

并且在初始化数据库的方法中进行调用：
```go

func InitMysql() {
	fmt.Println("InitMysql....")
	if db == nil {
		db, _ = sql.Open("mysql", "root:hanru1314@tcp(127.0.0.1:3306)/myblogweb")
		CreateTableWithUser()
		CreateTableWithArticle()
		CreateTableWithAlbum()
	}
}
```

### 二、上传文件
###### 2.1 AlbumController
然后我们创建一个新的go文件，album_controller.go

```go
package controllers

import (
	"myblog/models"
	"github.com/opentracing/opentracing-go/log"
)

type AlbumController struct {
	BaseController
}

func (this *AlbumController) Get() {
	this.TplName="album.html"
}
```

##### 2.2 View
接下来我们去写html页面，在views目录下创建一个新的html页面：

```html
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <title>相册</title>

    <script src="../static/js/lib/jquery-3.3.1.min.js"></script>
    <script src="../static/js/lib/jquery.url.js"></script>
    <script src="../static/js/blog.js"></script>
    <link href="../static/css/blogsheet.css" rel="stylesheet">
</head>
<body>
{{template "block/nav.html" .}}
<div id="main">
    <form method="post">
        <input type="file" id="album-upload-file" name="upload">
        <input type="button" id="album-upload-button" value="提交文件">
    </form>
</div>

</body>
</html>
```

然后我们需要编写点击按钮之后的脚本代码，在blog.js文件中继续添加脚本：
```go
//文件
    $("#album-upload-button").click(function () {
        var filedata = $("#album-upload-file").val();
        if (filedata.length <= 0) {
            alert("请选择文件!");
            return
        }
        //文件上传通过Formdata去储存文件的数据
        var data = new FormData()
        data.append("upload", $("#album-upload-file")[0].files[0]);
        alert(data)
        var urlStr = "/upload"
        $.ajax({
            url: urlStr,
            type: "post",
            dataType: "json",
            contentType: false,
            data: data,
            processData: false,
            success: function (data, status) {
                alert(":data:" + data.message);
                if (data.code == 1) {
                    setTimeout(function () {
                        window.location.href = "/album"
                    }, 1000)
                }
            },
            error: function (data, status) {
                alert("err:" + data.message + ":" + status)
            }
        })
    })
```

###### 2.3 UploadController

此时我们需要新建一个文件上传的controller，upload_controller.go

```go
package controllers

import (
	"fmt"
	"time"
	"path/filepath"
	"os"
	"io"
	"myblog/models"
)

type UploadController struct {
	BaseController
}

func (this *UploadController) Post() {
	fmt.Println("fileupload...")
	fileData, fileHeader, err := this.GetFile("upload")
	if err != nil {
		this.responseErr(err)
		return
	}
	fmt.Println("name:", fileHeader.Filename, fileHeader.Size)
	fmt.Println(fileData)
	now := time.Now()
	fmt.Println("ext:", filepath.Ext(fileHeader.Filename))
	fileType := "other"
	//判断后缀为图片的文件，如果是图片我们才存入到数据库中
	fileExt := filepath.Ext(fileHeader.Filename)
	if fileExt == ".jpg" || fileExt == ".png" || fileExt == ".gif" || fileExt == ".jpeg" {
		fileType = "img"
	}
	//文件夹路径
	fileDir := fmt.Sprintf("static/upload/%s/%d/%d/%d", fileType, now.Year(), now.Month(), now.Day())
	//ModePerm是0777，这样拥有该文件夹路径的执行权限
	err = os.MkdirAll(fileDir, os.ModePerm)
	if err != nil {
		this.responseErr(err)
		return
	}

	//文件路径
	timeStamp := time.Now().Unix()
	fileName := fmt.Sprintf("%d-%s", timeStamp, fileHeader.Filename)
	filePathStr := filepath.Join(fileDir, fileName)
	desFile, err := os.Create(filePathStr)
	if err != nil {
		this.responseErr(err)
		return
	}

	//将浏览器客户端上传的文件拷贝到本地路径的文件里面
	_, err = io.Copy(desFile, fileData)
	if err != nil {
		this.responseErr(err)
		return
	}

	if fileType == "img" {
		album := models.Album{0, filePathStr, fileName, 0, timeStamp}
		models.InsertAlbum(album)
	}

	this.Data["json"] = map[string]interface{}{"code": 1, "message": "上传成功"}
	this.ServeJSON()

}

func (this *UploadController) responseErr(err error) {
	this.Data["json"] = map[string]interface{}{"code": 0, "message": err}
	this.ServeJSON()
}

```

最后要注册路由：

```go
//相册
beego.Router("/album", &controllers.AlbumController{})

//文件上传
beego.Router("/upload", &controllers.UploadController{})
```

###### 2.4 Model

接下来我们创建一个新的model，album_model.go文件：

```go
type Album struct {
	Id         int
	Filepath   string
	Filename   string
	Status     int
	Createtime int64
}
```

我们需要在album_model.go文件中，添加插入数据的方法：

```go
//-------插入图片---------------
func InsertAlbum(album Album) (int64, error) {
	return utils.ModifyDB("insert into album(filepath,filename,status,createtime)values(?,?,?,?)",
		album.Filepath, album.Filename, album.Status, album.Createtime)
}
```

###### 2.5 运行
重启项目后，点击图片按钮，然后选择一张图片：
![http://7xtcwd.com1.z0.glb.clouddn.com/WX20190524-114125@2x.png](http://7xtcwd.com1.z0.glb.clouddn.com/WX20190524-114125@2x.png)

点击提交文件按钮后，可以上传图片了。

![http://7xtcwd.com1.z0.glb.clouddn.com/WX20190524-114242@2x.png](http://7xtcwd.com1.z0.glb.clouddn.com/WX20190524-114242@2x.png)

我们可以查询一下数据库：
![http://7xtcwd.com1.z0.glb.clouddn.com/WX20190524-114318@2x.png](http://7xtcwd.com1.z0.glb.clouddn.com/WX20190524-114318@2x.png)

### 三、查看图片

通过上面的操作已经能够上传文件了，那么如何显示文件呢？当点击图片标签的时候，除了有上传操作，还应该可以显示已经存储的图片。

###### 3.1 Model

我们现在album_model.go文件中，添加查找图片数据的方法：

```go
//--------查询图片----------
func FindAllAlbums() ([]Album, error) {
	rows, err := utils.QueryDB("select id,filepath,filename,status,createtime from album")
	if err != nil {
		return nil, err
	}
	var albums []Album
	for rows.Next() {
		id := 0
		filepath := ""
		filename := ""
		status := 0
		var createtime int64
		createtime = 0
		rows.Scan(&id, &filepath, &filename, &status, &createtime)
		album := Album{id, filepath, filename, status, createtime}
		albums = append(albums, album)
	}
	return albums, nil
}
```

###### 3.2 Controller

接下来，我们修改album_controller.go文件中的Get()方法，

```go
func (this *AlbumController) Get() {
	albums,err := models.FindAllAlbums()
	if err !=nil{
		log.Error(err)
	}
	this.Data["Album"] = albums
	this.TplName="album.html"
}
```

###### 3.3 View

最后我们还要修改一下html页面，在album.html中，添加以下内容：

```html
<div id="album-box">
    {{range .Album}}
        <div class="album-item" style='background-image: url("{{.Filepath}}");'></div>
    {{end}}
</div>
```

###### 3.4 运行

我们上传几张图片后，刷新页面：
![http://7xtcwd.com1.z0.glb.clouddn.com/WX20190524-114301@2x.png](http://7xtcwd.com1.z0.glb.clouddn.com/WX20190524-114301@2x.png)



## 关于我功能开发及项目总结
**@author：Davie**
**版权所有：北京千锋互联科技有限公司**

本节课内容我们来开发完成本系列课程的最后一个功能--关于我页面的功能开发，同时对本系列课程内容做一个回顾与总结。

### 一、Controller

首先创建一个新的go文件，aboutme_controller.go

```go
package controllers

type AboutMeController struct {
	BaseController
}

func (c *AboutMeController) Get() {
	c.Data["wechat"] = "微信：13167582311"
	c.Data["qq"] = "QQ：861574834"
	c.Data["tel"] = "Tel：13167582311"
	c.TplName = "aboultme.html"
}
```

然后不要忘记注册路由：

```go
//关于我
beego.Router("/aboutme", &controllers.AboutMeController{})
```

### 二、View

然后我们要写一个html页面，aboultme.html：
```html
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <title>{{.Title}}</title>
    <link href="../static/css/blogsheet.css" rel="stylesheet">
    <link href="../static/css/lib/highlight.css" rel="stylesheet">
</head>
<body>

{{template "block/nav.html" .}}

<div id="aboutme">
    <div>
        <img src="http://7xtcwd.com1.z0.glb.clouddn.com/wechat_avator.png">

        <div id="myt">
            <p class="myp">
                <a href="http://{{.wechat}}">{{.wechat}}</a>
            </p>
            <p class="myp">
                <a href="http://{{.qq}}">{{.qq}}</a>
            </p>
            <p class="myp">
                <a href="http://{{.tel}}">{{.tel}}</a>
            </p>
        </div>
    </div>
</div>
</body>
</html>
```

### 三、课程总结
经过了16节的课程内容，我们完成了使用Beego框架开发完成了一个博客系统。通过该项目，以期大家能够掌握Beego框架的使用方法。现在我们来回顾一下Beego框架开发一个web项目所需要掌握的一些重要知识点。

* **Beego框架的组成：**八大模块分别是cache，config，context，httplibs，logs，orm，session，toolbox等模块组合而成。模块之间高度解耦，依赖性低。

* **Bee工具：**Beego框架的项目管理工具Bee工具的使用，可以方便开发者管理、调试、打包项目，自动生成项目目录结构等。

* **Beeg框架执行流程：**Beego框架的程序执行流程如下图所示。
![http://7xtcwd.com1.z0.glb.clouddn.com/WX20190527-103955@2x.png](http://7xtcwd.com1.z0.glb.clouddn.com/WX20190527-103955@2x.png)

* **数据库操作：**
    
    * 在实际开发过程中，根据实际情况做数据库的选择。一般情况下，现在选择mysql数据库的企业原来越多。
    
    * 数据库连接：
        * 驱动：beego中的orm支持MySQL，Sqlite3、PostgreSQL。
        * 用户名、密码、主机、端口号、数据库名称等的配置。
    
    * 数据库操作：
        * sql语句
        * 条件查询
        * 统计功能
        * 增加（insert）、删除（delete）、修改（update）、查询（select）
        * 模糊查询
        * 表关联（未使用到）

* **Beego项目架构：**
    
    * **mvc模式：**m（model）模型层，v（view）视图层，c（controller）控制器层。
    
    * **路由解析：**
        * 默认路由：beego模块自带Post、Put、Delete、Head、Options、Patch等网络请求类型的对应方法
        * 自动路由（自动映射到Get、Post、Put、Delete、Head、Options、Patch等方法。
        * 正则表达式路由（支持正则表达式及变量）
        * 自定义路由：开发者自己指定路由解析方法

* **Session处理：**
    * 两种管理方式：
        * 配置文件配置session相关配置项
        * 在程序中通过SessionConfig中进行配置
    
    * 操作方法：
        * SetSession：设置session值
        * GetSession：获取session值
        * DelSession：删除session值

* **Views视图模板：**
    * views目录存放视图模板文件
    
    * controller.TplName 指定渲染的页面模板文件全称
    
    * 模板文件中通过**{{.param}}**实现变量的使用
    
    * controller.Data["param"]为模板页面的变量赋值



