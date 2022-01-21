# Area-China

中华人民共和国行政区划：省级（省份直辖市自治区）、 地级（城市）、 县级（区县）

## 指令flag

* `-year` 抓取的年份 默认2020
* `-size` 默认code的长度,不足则补0,最大长度为17位,默认长度为6位，最小长度为2位
* `filename` 导出的文件名称

## 数据格式

数据格式为`json`格式

```json
[
  {
    "code ": "编码",
    "name": "名称",
    "children": [ 
      {
        "code": "编码",
        "name": "名称",
        "children": []
      }
    ]
  }
]
```

# 数据来源

+ [统计用区划代码和城乡划分代码编制规则](http://www.stats.gov.cn/tjsj/tjbz/200911/t20091125_8667.html)
+ [统计用区划和城乡划分代码](http://www.stats.gov.cn/tjsj/tjbz/tjyqhdmhcxhfdm/)

# thanks

* [Administrative-divisions-of-China](https://github.com/modood/Administrative-divisions-of-China)
