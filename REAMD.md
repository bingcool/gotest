### go语言学习笔记

可以学习源码，推荐：ithub.com/sohaha/zlsgo

go 标准常用包
fmt 包：提供了输入输出和格式化的函数。  
math 包：提供了数学函数和常量。   
net 包：提供了网络编程的功能。   
time 包：提供了时间和日期的相关功能。   
strings 包：提供了字符串处理的功能。   
os 包：提供了操作系统相关的功能。  
flag: cli参数输入,输出处理   


第三方常用的包：
github.com/robfig/cron/v3: 封裝的cron的计划任务定时器的处理

github.com/spf13/cobra：构建console的应用，定义各种命令   

github.com/spf13/viper: 使用 Viper 来管理配置

github.com/cheggaaa/pb: 终端进度条管理-Terminal progress bar for Go

github.com/swaggo/gin-swagger: 自动生成接口文档

github.com/urfave/cli: 用于在Go中构建命令行工具

go-playground/validator: 参数验证

github.com/go-redis/redis/v8： redis操作库

go.uber.org/zap：Uber开源的日志库
gopkg.in/natefinch/lumberjack: 日志分割库包

github.com/go-playground/validator: web开发的参数验证包

github.com/samber/lo : Lo 是一个基于 Go 1.18 泛型的 Lodash 风格的 Go 库,是一个函数式编程库，提供了一些常用的函数，例如 map、filter、reduce、find、findIndex、some、every、sort、reverse、slice、concat、flatten、uniq、intersection、difference、union、zip、unzip、groupBy、countBy、findKey、findLastKey、forIn、forInRight、forOwn、forOwnRight、forEach、forEachRight、reduce、reduce

github.com/fsnotify/fsnotify: Go 语言库，用于在 Windows、Linux、macOS、BSD 和 illumos 上提供跨平台的文件系统通知功能

github.com/joho/godotenv: Go 语言库，用于加载环境变量

go fiber : Go 语言库，用于构建高性能、可扩展的 HTTP 服务器，它提供了许多方法来处理 HTTP 请求，例如处理静态文件，处理动态文件，处理路由，等等。

https://taskfile.dev/ : Taskfile 是一个用于管理 Go 项目的构建工具，它允许你定义一个 YAML 文件来定义你的构建任务，然后通过一个简单的命令行工具来运行这些任务。

https://github.com/golang-module/carbon : Carbon 是一个 Go 语言库，用于处理日期和时间，它提供了许多方法来处理日期和时间，例如格式化日期和时间，计算日期和时间差，等等。
    
https://github.com/gocolly/colly : Colly 是一个 Go 语言库，用于抓取网站数据，它提供了许多方法来抓取网站数据，例如抓取 HTML 页面，抓取 JSON 数据，抓取 XML 数据，等等。

https://github.com/redis/go-redis 这是一个很棒的、高度维护的Go redis数据库客户端。它与redis 6和7都兼容，并且有一个非常简单的设置过程。强烈推荐

https://mp.weixin.qq.com/s/n-d2HKiDhNYrTYSFRksf3w 实时可视化Go Runtime指标

https://github.com/qax-os/excelize excelize 是一个 Go 语言库，用于读取、写入和操作 Microsoft Excel 文件。它提供了许多方法来读取、写入和操作 Excel 文件，例如读取工作簿、读取工作表、读取单元格、写入单元格、写入工作表、写入工作簿、写入 Excel 文件等。

https://github.com/go-echarts/go-echarts 这是一个Go语言库，用于生成ECharts图表。它提供了许多方法来生成ECharts图表，例如生成折线图、柱状图、饼图、雷达图、散点图、地图等。

https://github.com/segmentio/kafka-go 这是一个Go语言库，用于与Kafka进行通信。它提供了许多方法来与Kafka进行通信，例如创建Topic、创建Partition、创建Consumer、创建Producer、发送消息、接收消息、消费消息等。

https://github.com/rabbitmq/amqp091-go 这是一个Go语言库，用于与RabbitMQ进行通信。它提供了许多方法来与RabbitMQ进行通信，例如创建Channel、创建Queue、创建Exchange、创建Binding、发送消息、接收消息、消费消息等。


