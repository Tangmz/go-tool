### 读取配置

* 一个读取自定义配置文件的工具.目前支持.conf格式


### 不定时更新

 1. 目前只支持通过文件路径读取配置文件,之后将支持通过url读取配置文件

 2. 之后将支持配置文件中引入配置文件.

 3. 之后将支持配置项重写,即一个配置文件引入多个配置文件,支持配置项优先级


---
有配置文件config.conf,内容如下
```
# this is a config test file

int=1
int32=2
int64=3
float32=1.1
float64=2.2
string=halo

judge=judge #判断这个

 # 过滤这个

## 过滤掉这个


 val = hey###### val = hey
```

读取配置文件config.conf
```

    confPath := "config.conf"
    conf := NewConfig()
    err = conf.Config(confPath)
    if err != nil {
        return
    }
```
