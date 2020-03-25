# OJI

golang 版本 [oji](https://github.com/xxczaki/oji)

```
? 选择左胳膊: ╰
? 选择右胳膊: ╯
? 选择左身体: (
? 选择右身体: )
? 选择左腮红: .
? 选择右腮红: .
? 选择左眼睛: ^
? 选择右眼睛: ^
? 选择鼻子嘴巴: ω
! 你的颜文字造好啦: ╰(.^ω^.)╯
? 复制到剪贴板: Yes
```

## 下载

[Github releases](https://github.com/KeJunMao/oji/releases)

或者

```shell script
go install github.com/KeJunMao/oji
```

## 用法

```shell script
oji help
```

## 自定义部件

```shell script
oji init -c config.json # 输出默认配置
vim config.json # 自定义部件
oji -c config.json # 使用自定义部件
```