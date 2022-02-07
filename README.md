# Area-China

中华人民共和国行政区划：省级（省份直辖市自治区）、 地级（城市）、 县级（区县）

## 指令flag

* `-year` 抓取的年份 默认2020
* `-size` 默认code的长度,不足则补0,最大长度为17位,默认长度为6位，最小长度为2位
* `-filename` 导出的文件名称

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
* [全国行政区划信息查询平台](http://xzqh.mca.gov.cn/map)

# thanks

* [Administrative-divisions-of-China](https://github.com/modood/Administrative-divisions-of-China)
