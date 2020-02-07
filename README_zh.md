# 验证中国身份证号码
![](https://travis-ci.org/yzy613/Verify-Chinese-ID-Number.svg?branch=master)

[README](https://github.com/yzy613/Verify-Chinese-ID-Number/blob/master/README.md)|[中文文档](https://github.com/yzy613/Verify-Chinese-ID-Number/blob/master/README_zh.md)

验证中国身份证号码是否正确。程序采用 Go 语言编写，支持带参运行。身份证号码最后一位校验码采用 MOD11-2 的标准校验，此程序同理。

## 用法

    1. 直接运行

    2. -id string
    
        输入18位身份证号码