关于FA系统
============

> FA系统，全称 **Fecshop Analysis** ，是针对fecshop开源电商系统打造的一款商城用户
数据分析统计系统，是通过js打点和php发送数据的2种方式接收原始数据，然后通过一系列的mapreduce
计算，归并算法计算统计数据，然后呈现给营销人员的系统，
trace系统就像一个摄像头，时刻接收来自网站的数据，统计计算，帮助营销人员，
查看历史数据，调整营销策略，管理营销人员，
根据历史数据决定运营策略，提高你的电商网站的销售额

### FA系统简介

1.FA Github源码地址

Github Golang部分：https://github.com/fecshopsoft/fec-go

Github Vue部分：https://github.com/fecshopsoft/vue-element-admin

2.FA演示地址：

> test账户有一定的权限限制

Demo：http://trace.fecshop.com/

账户：test   密码：test123

**不要修改密码，否则将会关闭测试demo !!**


### 为什么要做FA系统

对于电商系统，很多的追踪，可以用`google analysis`追踪，如果做广告，某些大的广告平台
也会提供一些统计功能，为什么fecshop要做`数据统计分析系统`呢？原因如下：

1.`google analysis`针对的是所有的网站类型，虽然针对电商做了一些
升级功能，但是，对于电商来说，远远不够

2.这些平台的数据授权，对于业务数据的收集不够全面，它没有收集`用户注册` `登录`数据，
没有收集具体的`sku`，`分类搜索`，等数据，对于针对业务的数据，`不够全面`，
因此无法满足需要，譬如我想通过email搜索某个用户的行为数据（email是用户在商城
注册的email），是无法满足的

3.这些统计平台，原理是通过`js`的方式收集，也只能通过`页面加载js`的方式收集数据，
而对于一些`没有页面的数据`，是不能收集的，譬如登录，注册，产品加入购物车等，
而FA系统可通过api接收服务端传递的数据，这样收集的数据更多，更全面，更准确。

4.对于`订单数据`的收集，`google analysis`是通过订单成功页面进行的，
下面的情况会造成订单数据`不准确`：

4.1电商网站生成订单，跳转到第三方支付平台，支付完成后，用户直接关掉了页面，
并没有跳转回电商订单支付成功页面，因此没有加载支付成功页面的js，
进而`无法收集` 订单支付成功数据。

4.2对于跨境商城，有一些支付并不像`paypal`，`支付宝`这样，很快就可以支付完成，
而是需要等几十分钟，几个小时，因为这些支付渠道需要到相应的信用卡银行去扣款，
存在延迟，当支付成功后，支付通道会通过`IPN`消息的方式通知商城，支付成功，更改
订单状态为`支付成功状态`,而IPN发送的订单支付状态，是支付通道发给服务端的，是`没有浏览器
界面`的，因此，传统的js收集数据的方式`并不能`收集到订单支付成功数据
，而FA系统可以通过api接收商城传送的订单支付成功数据。

5.对于广告分析支持不够

作为公司的广告，每一个广告都需要花钱，
从老板的角度，就想对广告数据进行更加`详细的统计`,下面的针对广告做的`精细数据分析`，也是
`Trace系统`独有的

5.1数据统计：每个`广告`在每一天的数据报告，
`每个广告`每天带来多少pv，uv，生成了多少订单，新增了多少用户等等。

5.2数据统计：每个`广告员`的所有的广告汇总，统计这个广告员的具体的数据报告

5.3数据统计：每个`广告小组`的所有的广告汇总

5.4数据统计：每个`广告活动`的所有的广告汇总

5.5数据统计；每个`渠道`，譬如facebook, google ppc，的所有广告汇总

5.6数据统计：每个`子渠道`的广告汇总

5.7数据统计：针对`EDM`这类，`多链接广告`的统计，统计各个链接进入
商城的流量的实际情况。

上面只是说了一个大概，具体的统计数据参看系统里面的具体详细。

6.支持不够，对于`vue`这种`前后端彻底分离`的商城应用，支持明显不够。


7.`底层数据`决定`上层建筑`，`底层数据`的收集`不全面` `不准确`，
会造成后面的统计数据`不够准确`

8.最后，最重要的，是数据的掌控性，我本地有了数据，那么我可以根据业务需要进行二次开发，
进行数据分析，满足需要。

### FA系统的特点

1.双方式收集数据，保证收集的业务数据更全面，数据更准确

1.1`js`打点的方式收集数据，类似于google analysis，将js嵌入网站页面，收集浏览器，ip，设备和一些业务信息
，除了这些通用数据，还收集`产品sku`，`搜索关键词`，`分类名称`，等业务数据。


1.2服务端通过`api`发送给FA系统的方式手机数据，这些主要收集一些
`无页面数据`，譬如登录注册，加入购物车等，和订单数据，尤其针对前面第4部分提到
的订单数据不准确问题的解决。

2.支持的场景多

除了传统的pc和手机浏览器这种商城，还支持`vue`这种`前后端彻底分离`的场景

3.深入业务和业务紧密相连

> 因为收集的数据全面，准确，进而可以和业务结合起来，进行很多实用性很强的统计。

3.1通用部分数据统计

`站点统计`：统计用户下的各个商城，每一天的汇总

`App入口统计`：fecshop分为appfront，apphtml5，以及vue类的前后端彻底分离的应用，为各个入口做具体的统计

`来源统计`：通过第三方网站跳转到商城后，来源指的是第三方网站的域名，redirect指的是直接访问

`设备统计`：各种设备的统计

`国家统计`：各个国家的统计

`浏览器统计`：各个浏览器的统计

3.2业务数据统计

`Sku统计`：各个sku的访问次数，页面跳出率，加入购物车数，生成订单数，支付订单数等数据
，深入结合业务，为每一个产品做具体的分析。

`Sku Refer统计`：通过各个来源点击后的sku的统计，和上面的sku统计类似，不同的是，按照
来源进行区分统计

`搜索统计`：各个搜索词的搜索情况，以及跳出率，点击后生成的订单数据

`搜索语言统计`：以语言进行区分，各个搜索关键词的具体统计

`着陆url统计`：用户的着陆页的数据统计

`url统计`：各个url的统计

`分类统计`：各个分类的统计


3.3广告数据统计

也就是上面第5部分的讲述，针对广告做了更细粒度的统计分析，
帮助广告员优化各个广告。















