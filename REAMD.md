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
https://github.com/shockerli/go-awesome?tab=readme-ov-file 

gorm.io/gorm : go的ORM数据库包
github.com/bingcool/gen : 基于gorm自动生成model和query代码  

github.com/robfig/cron/v3: 封裝的cron的计划任务定时器的处理

github.com/spf13/cobra：构建console的应用，定义各种命令   

github.com/spf13/viper: 使用 Viper 来管理配置

github.com/spf13/cast: 类型转化工具库cast详解

github.com/cheggaaa/pb: 终端进度条管理-Terminal progress bar for Go

github.com/swaggo/gin-swagger: 自动生成接口文档

github.com/urfave/cli: 用于在Go中构建命令行工具

github.com/duke-git/lancet/v2  （推荐）一个全面、高效、可复用的go语言工具函数库  

github.com/mitchellh/mapstructure 提供了更灵活的方式来将 map 转换为结构体，而无需经过 JSON 序列化/反序列化的中间步骤 

go-playground/validator: 参数验证

github.com/gogf/gf/v2: 高性能的Go语言框架，用于快速开发Web服务,里面很多工具包可以使用，gjson.gconv,gutil,gfile等

github.com/gookit/validate：参数验证   

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

https://github.com/mattn ： 大牛项目仓库

https://github.com/qmuntal/stateless go的状态机的包

https://github.com/looplab/fsm go的有限状态机

https://github.com/mattn/go-sqlite3     

https://github.com/shockerli/go-awesome

https://github.com/schollz/progressbar: 进度条库

https://github.com/pterm/pterm ： 终端美化库

golang.org/x/tools/cmd/goimports : 自动导入包规范包的位置

https://github.com/bytedance/sonic?tab=readme-ov-file ： 字节开源的json库

https://github.com/goccy/go-json :  Go 语言库，用于解析和序列化 JSON 数据,性能高

https://github.com/hashicorp/golang-lru ： LRU 缓存库

https://github.com/sourcegraph/conc ： 短小精悍的go协程库

github.com/go-resty/resty： curl的包，类似php的guzzle-client包，封装了http请求，支持json、xml、form、multipart、文件上传、文件下载、重试、超时、代理、认证、压缩、日志、监控等

github.com/gofrs/uuid： Go语言库，用于生成UUID

github.com/Shopify/sarama：Sarama 是一个纯 Go 编写的 Kafka 客户端库，支持 Apache Kafka 0.8 及以上版本。它提供了一个高级 API 用于简化消息的生产和消费，同时也提供了一个低级 API，用于更细粒度的控制。 支持多种Kafka版本，良好的文档和社区支持

github.com/confluentinc/confluent-kafka-go：Confluent 的 Kafka Go 客户端，它是对 C 库的封装，提供了与 Confluent Platform 版本的 Kafka 完全兼容的功能。与 Confluent Platform 的高度集成，支持 SASL/SSL 认证，可能在性能上有优势

github.com/segmentio/kafka-go：Segment 提供的 Kafka Go 客户端，设计简洁，易于使用，专注于性能和灵活性。简洁的 API，高性能，支持多种 Kafka 功能

github.com/fatih/color 是用于输出对应编码颜色的包。
github.com/schollz/progressbar 是用于为执行时间过久的任务创建进度条的包。
github.com/jimlawless/whereami 是用于捕获源代码的文件名、行号、函数等信息的包，这对于改正错误信息十分有用！
github.com/spf13/cobra 是用于更轻松地创建带有输入选项和相关文档的复杂脚本的包。



