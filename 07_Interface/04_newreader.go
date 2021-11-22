// a simple html parser, implement io.Reader interface
package main

import (
	"The.Go.Programming.Language/localpkg/ehtml"
	"fmt"
	"golang.org/x/net/html"
)

func main() {
	r := ehtml.NewReader(text)
	doc, err := ehtml.Parse(r)
	if err != nil {
		fmt.Println(err)
	}
	myVisit(doc)
}

func myVisit(n *html.Node) {
	if n.Type == html.ElementNode {
		fmt.Println(n.Data)
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		myVisit(c)
	}
}

var text string = `
<!DOCTYPE html>
<html lang="zh-CN">
<head>
<meta charset="utf-8" />
<meta http-equiv="X-UA-Compatible" content="IE=edge" />
<meta name="viewport" content="width=device-width, initial-scale=1" />
<meta name="renderer" content="webkit" />
<meta name="force-rendering" content="webkit"/>
<meta name="applicable-device" content="pc,mobile" />
<meta name="MobileOptimized" content="width" />
<meta name="HandheldFriendly" content="true" />
<meta http-equiv="Cache-Control" content="no-transform" />
<meta http-equiv="Cache-Control" content="no-siteapp" />
<meta name="format-detection" content="telephone=no" />
<link rel="shortcut icon" href="/favicon.ico?v=1.6.69" />
<link href="/templets/new/style/common.css?v=1.6.69" rel="stylesheet" />
<title>Go语言copy()：切片复制（切片拷贝）</title>
<meta name="description" content="使用 Go 语言内建的 copy() 函数，可以迅速地将一个切片的数据复制到另外一个切片空间中，copy() 函数的使用格式如下： copy( destSlice, srcSlice []T) int srcSlice 为数据来源切片。 destSlice 为复" />
</head>
<body>
<div id="topbar" class="clearfix">
<ul id="product-type" class="left">
<li>
<a href="/"><span class="iconfont iconfont-home"></span>首页</a>
</li>
<li class="active">
<a href="/sitemap/" rel="nofollow"><span class="iconfont iconfont-book"></span>教程</a>
</li>
<li>
<a href="http://vip.biancheng.net/p/vip/show.php" rel="nofollow" target="_blank"><span class="iconfont iconfont-vip"></span>VIP会员</a>
</li>
<li>
<a href="http://vip.biancheng.net/p/q2a/show.php" rel="nofollow" target="_blank"><span class="iconfont iconfont-q2a"></span>一对一答疑</a>
</li>
<li>
<a href="http://fudao.biancheng.net/" rel="nofollow" target="_blank"><span class="iconfont iconfont-fudao"></span>辅导班</a>
</li>
</ul>
</div>
<div id="header" class="clearfix">
<a id="logo" class="left" href="/">
<img height="26" src="/templets/new/images/logo.png?v=1.6.69" alt="C语言中文网" />
</a>
<ul id="nav-main" class="hover-none left clearfix">
<li class="wap-yes"><a href="/">首页</a></li>
<li><a href="/c/">C语言教程</a></li>
<li><a href="/cplus/">C++教程</a></li>
<li><a href="/python/">Python教程</a></li>
<li><a href="/java/">Java教程</a></li>
<li><a href="/linux_tutorial/">Linux入门</a></li>
<li><a href="/sitemap/" title="网站地图">更多&gt;&gt;</a></li>
</ul>
<span id="sidebar-toggle" class="toggle-btn" toggle-target="#sidebar">目录 <span class="glyphicon"></span></span>
<a href="http://vip.biancheng.net/?from=topbar" class="user-info glyphicon glyphicon-user hover-none" target="_blank" rel="nofollow" title="用户中心"></a>
</div>
<div id="main" class="clearfix">
<div id="sidebar" class="toggle-target">
<div id="contents">
<dt><span class="glyphicon glyphicon-option-vertical" aria-hidden="true"></span><a href="/golang/">Go语言</a></dt>
<dd>
<span class="channel-num">1</span>
<a href='/golang/intro/'>Go语言简介</a>
</dd>
<dd>
<span class="channel-num">2</span>
<a href='/golang/syntax/'>Go语言基本语法</a>
</dd>
<dd class="this"> <span class="channel-num">3</span> <a href="/golang/container/">Go语言容器</a> </dd><dl class="dl-sub"><dd>3.1 <a href="/view/26.html">Go语言数组</a></dd><dd>3.2 <a href="/view/4117.html">Go语言多维数组</a></dd><dd>3.3 <a href="/view/27.html">Go语言切片</a></dd><dd>3.4 <a href="/view/28.html">使用append()为切片添加元素</a></dd><dd>3.5 <a href="/view/29.html">Go语言切片复制</a></dd><dd>3.6 <a href="/view/30.html">Go语言从切片中删除元素</a></dd><dd>3.7 <a href="/view/4118.html">Go语言range关键字</a></dd><dd>3.8 <a href="/view/4119.html">Go语言多维切片</a></dd><dd>3.9 <a href="/view/31.html">Go语言map（映射）</a></dd><dd>3.10 <a href="/view/32.html">Go语言遍历map</a></dd><dd>3.11 <a href="/view/33.html">map元素的删除和清空</a></dd><dd>3.12 <a href="/view/vip_7306.html">Go语言map的多键索引</a><span class="glyphicon glyphicon-usd"></span></dd><dd>3.13 <a href="/view/34.html">Go语言sync.Map</a></dd><dd>3.14 <a href="/view/35.html">Go语言list（列表）</a></dd><dd>3.15 <a href="/view/4776.html">Go语言nil：空值/零值</a></dd><dd>3.16 <a href="/view/vip_7307.html">Go语言make和new关键字的区别及实现原理</a><span class="glyphicon glyphicon-usd"></span></dd></dl>
<dd>
<span class="channel-num">4</span>
<a href='/golang/flow_control/'>流程控制</a>
</dd>
<dd>
<span class="channel-num">5</span>
<a href='/golang/func/'>Go语言函数</a>
</dd>
<dd>
<span class="channel-num">6</span>
<a href='/golang/struct/'>Go语言结构体</a>
</dd>
<dd>
<span class="channel-num">7</span>
<a href='/golang/interface/'>Go语言接口</a>
</dd>
<dd>
<span class="channel-num">8</span>
<a href='/golang/package/'>Go语言包（package）</a>
</dd>
<dd>
<span class="channel-num">9</span>
<a href='/golang/concurrent/'>Go语言并发</a>
</dd>
<dd>
<span class="channel-num">10</span>
<a href='/golang/reflect/'>Go语言反射</a>
</dd>
<dd>
<span class="channel-num">11</span>
<a href='/golang/102/'>Go语言文件处理</a>
</dd>
<dd>
<span class="channel-num">12</span>
<a href='/golang/build/'>Go语言编译与工具</a>
</dd>
</div>
</div>
<div id="article-wrap">
<div id="article">
<div class="arc-info">
<span class="position"><span class="glyphicon glyphicon-map-marker"></span> <a href="/">首页</a> &gt; <a href="/golang/">Go语言</a> &gt; <a href="/golang/container/">Go语言容器</a></span>
<span class="read-num">阅读：57,802</span>
</div>
<div id="ad-position-bottom"></div>
<h1>Go语言copy()：切片复制（切片拷贝）</h1>
<div class="pre-next-page clearfix">
<span class="pre left"><span class="icon">&lt;</span> <span class="text-brief text-brief-pre">上一页</span><a href="/view/28.html">使用append()为切片添加元素</a></span>
<span class="next right"><a href="/view/30.html">Go语言从切片中删除元素</a><span class="text-brief text-brief-next">下一页</span> <span class="icon">&gt;</span></span>
</div>
<div id="ad-arc-top"><p class="pic"></p><p class="text" adid="default"></p></div>
<div id="arc-body">Go语言的内置函数 copy() 可以将一个数组切片复制到另一个数组切片中，如果加入的两个数组切片不一样大，就会按照其中较小的那个数组切片的元素个数进行复制。<br />
<br />
copy() 函数的使用格式如下：
<p class="info-box">
copy( destSlice, srcSlice []T) int</p>
其中 srcSlice 为数据来源切片，destSlice 为复制的目标（也就是将 srcSlice 复制到 destSlice），目标切片必须分配过空间且足够承载复制的元素个数，并且来源和目标的类型必须一致，copy() 函数的返回值表示实际发生复制的元素个数。<br />
<br />
下面的代码展示了使用 copy() 函数将一个切片复制到另一个切片的过程：
<pre class="go">
slice1 := []int{1, 2, 3, 4, 5}
slice2 := []int{5, 4, 3}
copy(slice2, slice1) // 只会复制slice1的前3个元素到slice2中
copy(slice1, slice2) // 只会复制slice2的3个元素到slice1的前3个位置</pre>
虽然通过循环复制切片元素更直接，不过内置的 copy() 函数使用起来更加方便，copy() 函数的第一个参数是要复制的目标 slice，第二个参数是源 slice，两个 slice 可以共享同一个底层数组，甚至有重叠也没有问题。<br />
<br />
【示例】通过代码演示对切片的引用和复制操作后对切片元素的影响。<br />
<pre class="go">
package main

import &quot;fmt&quot;

func main() {

    // 设置元素数量为1000
    const elementCount = 1000

    // 预分配足够多的元素切片
    srcData := make([]int, elementCount)

    // 将切片赋值
    for i := 0; i &lt; elementCount; i++ {
        srcData[i] = i
    }

    // 引用切片数据
    refData := srcData

    // 预分配足够多的元素切片
    copyData := make([]int, elementCount)
    // 将数据复制到新的切片空间中
    copy(copyData, srcData)

    // 修改原始数据的第一个元素
    srcData[0] = 999

    // 打印引用切片的第一个元素
    fmt.Println(refData[0])

    // 打印复制切片的第一个和最后一个元素
    fmt.Println(copyData[0], copyData[elementCount-1])

    // 复制原始数据从4到6(不包含)
    copy(copyData, srcData[4:6])

    for i := 0; i &lt; 5; i++ {
        fmt.Printf(&quot;%d &quot;, copyData[i])
    }
}</pre>
代码说明如下：
<ul>
<li>
第 8 行，定义元素总量为 1000。</li>
<li>
第 11 行，预分配拥有 1000 个元素的整型切片，这个切片将作为原始数据。</li>
<li>
第 14～16 行，将 srcData 填充 0～999 的整型值。</li>
<li>
第 19 行，将 refData 引用 srcData，切片不会因为等号操作进行元素的复制。</li>
<li>
第 22 行，预分配与 srcData 等大（大小相等）、同类型的切片 copyData。</li>
<li>
第 24 行，使用 copy() 函数将原始数据复制到 copyData 切片空间中。</li>
<li>
第 27 行，修改原始数据的第一个元素为 999。</li>
<li>
第 30 行，引用数据的第一个元素将会发生变化。</li>
<li>
第 33 行，打印复制数据的首位数据，由于数据是复制的，因此不会发生变化。</li>
<li>
第 36 行，将 srcData 的局部数据复制到 copyData 中。</li>
<li>
第 38～40 行，打印复制局部数据后的 copyData 元素。</li>
</ul>
</div>
<div id="arc-append">
<p>关注微信公众号「站长严长生」，在手机上阅读所有教程，随时随地都能学习。本公众号由<a class="col-link" href="/view/8092.html" target="_blank" rel="nofollow">C语言中文网站长</a>运营，每日更新，坚持原创，敢说真话，凡事有态度。</p>
<p style="margin-top:12px; text-align:center;">
<img width="180" src="/templets/new/images/material/qrcode_weixueyuan_original.png?v=1.6.69" alt="魏雪原二维码"><br>
<span class="col-green">微信扫描二维码关注公众号</span>
</p>
</div>
<div class="pre-next-page clearfix">
<span class="pre left"><span class="icon">&lt;</span> <span class="text-brief text-brief-pre">上一页</span><a href="/view/28.html">使用append()为切片添加元素</a></span>
<span class="next right"><a href="/view/30.html">Go语言从切片中删除元素</a><span class="text-brief text-brief-next">下一页</span> <span class="icon">&gt;</span></span>
</div>
<div id="ad-arc-bottom"></div>
<div id="nice-arcs" class="box-bottom">
<h4>优秀文章</h4>
<ul class="clearfix">
<li><a href="/view/1085.html" title="Java泛型简明教程">Java泛型简明教程</a></li>
<li><a href="/view/vip_1800.html" title="C语言scanf的高级用法，原来scanf还有这么多新技能">C语言scanf的高级用法，原来scanf还有这么多新技能</a></li>
<li><a href="/view/1809.html" title="C语言?和:详解，C语言条件运算符详解">C语言?和:详解，C语言条件运算符详解</a></li>
<li><a href="/view/1875.html" title="Qt MDI及其使用方法（详解版）">Qt MDI及其使用方法（详解版）</a></li>
<li><a href="/view/2893.html" title="Shell模块化（source命令）">Shell模块化（source命令）</a></li>
<li><a href="/view/5188.html" title="Hibernate createCriteria方法：创建Criteria对象">Hibernate createCriteria方法：创建Criteria对象</a></li>
<li><a href="/view/5501.html" title="Qt程序的字符编码方式">Qt程序的字符编码方式</a></li>
<li><a href="/view/5963.html" title="JS DOMContentLoaded事件：DOM文档结构加载完毕">JS DOMContentLoaded事件：DOM文档结构加载完毕</a></li>
<li><a href="/view/7456.html" title="MySQL CROSS JOIN：交叉连接">MySQL CROSS JOIN：交叉连接</a></li>
<li><a href="/view/8278.html" title="如何将SQL语句映射为文件操作">如何将SQL语句映射为文件操作</a></li>
</ul>
</div>
</div>
</div>
</div>
<script type="text/javascript">
// 当前文章ID
window.arcIdRaw = 'a_' + 29;
window.arcId = "93210ueeHo+3zqoh0S8bFc36DaL9ZBtwBBUUddijuDK1MUO4pB1UXCzC";
window.typeidChain = "5|1";
</script>
<div id="footer" class="clearfix">
<div class="info left">
<p>精美而实用的网站，分享优质编程教程，帮助有志青年。千锤百炼，只为大作；精益求精，处处斟酌；这种教程，看一眼就倾心。</p>
<p>
<a href="/view/8066.html" target="_blank" rel="nofollow">关于网站</a> <span>|</span>
<a href="/view/8092.html" target="_blank" rel="nofollow">关于站长</a> <span>|</span>
<a href="/view/8097.html" target="_blank" rel="nofollow">如何完成一部教程</a> <span>|</span>
<a href="/view/8093.html" target="_blank" rel="nofollow">联系我们</a> <span>|</span>
<a href="/sitemap/" target="_blank" rel="nofollow">网站地图</a>
</p>
<p>Copyright ©2012-2021 biancheng.net, <a href="http://www.beian.miit.gov.cn/" target="_blank" rel="nofollow" style="color:#666;">陕ICP备15000209号</a></p>
</div>
<img class="right" src="/templets/new/images/logo_bottom.gif?v=1.6.69" alt="底部Logo" />
<span id="return-top"><b>↑</b></span>
</div>
<script type="text/javascript">
window.siteId = 4;
window.cmsTemplets = "/templets/new";
window.cmsTempletsVer = "1.6.69";
</script>
<script src="/templets/new/script/jquery1.12.4.min.js"></script>
<script src="/templets/new/script/common.js?v=1.6.69"></script>
<span style="display:none;"><script src="http://s19.cnzz.com/z_stat.php?id=1274766082&web_id=1274766082" type="text/javascript" defer="defer" async="async"></script></span>
</body>
</html>`
