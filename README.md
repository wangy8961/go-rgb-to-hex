# go-rgb-to-hex

将 https://a.xiumi.us/board/v5/251mJ/23297905 和 https://b.xiumi.us/board/v5/251mJ/23381856 中的 RGB 颜色字符串转换成十六进制格式的，因为像 https://www.processon.com/ 这类网站选择颜色都是十六进制格式的


> 如何使用？

下载我已经编译好的可执行程序 `RGBToHex`，或者你克隆代码后用 `go build -o RGBToHex` 自己编译

```bash
1. 查看帮助文档
[root@CentOS go-rgb-to-hex]# ./RGBToHex -h
Usage of ./RGBToHex:
  -h	显示帮助信息
  -r string
    	请指定包含 Red,Green,Blue 的字符串，颜色值用逗号隔开即可，可包含其它字符，比如 'rgb(255,97,171);' 也能正常解析

2. 使用
[root@CentOS go-rgb-to-hex]# ./RGBToHex -r "rgb(255,97,171);"
Hex conv of 'rgb(255,97,171);' is: #FF61AB

[root@CentOS go-rgb-to-hex]# ./RGBToHex -r "255,97,171"
Hex conv of '255,97,171' is: #FF61AB
```