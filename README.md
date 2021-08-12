# kubectl 插件 kubeimg
## 练习用

## 参照源


[Cobra + Client-go实现K8s 自定义插件开发](https://juejin.cn/post/6983324056502140964 "原理参考资料")
### 参考资料问题
1. 源代码不全
2. 部分图片无法查看
---

[Github:hejianlai/kubeimg](https://github.com/hejianlai/kubeimg "代码补充参考")

### 参考资料问题
1. 部分依赖太多，主要是指间接依赖
2. 没有原理讲解，配合上面的参考资料刚好合适

## 主要改动
1. 对homedir的依赖，改成本地库了
2. 对client-go的版本做了修改

## 后续改动
1. 由于本地client-go的间接依赖错误，改用线上依赖

[安装使用说明](./doc/install-guid.md "安装使用说明")







