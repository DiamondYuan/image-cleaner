# image-cleaner


清理全部的 Docker 镜像，可以使用正则编写白名单。


## 使用方法

| 参数名称     | 参数类型 | 参数效果             | 默认    |
| -------- | ---- | ---------------- | ----- |
| notWait  | bool | 是否停止删除镜像之前的等待    | false |
| dryRun   | bool | 是否只列出需要删除的镜像而不删除 | false |
| infinite | bool | 程序是否循环定时永久执行     | false |


### 1 列出全部需要删除的镜像
```bash
docker run --rm  -v /var/run/docker.sock:/var/run/docker.sock diamondyuan/image-cleaner
```

### 2 添加白名单

docker 运行参数里面增加  `-v ~/temp/whiteList:/whiteList` ,其中 `~/temp/whiteList` 是个人定义的白名单的地址。

### 3 立刻删除
```Bash
docker run --rm  -v /var/run/docker.sock:/var/run/docker.sock -v ~/whiteList:/whiteList diamondyuan/image-cleaner -notWait
```




























