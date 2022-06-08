# go-clean-arch

## 什么是干净的架构 
干净的架构来自于 [Uncle Bob](https://8thlight.com/blog/uncle-bob/2012/08/13/the-clean-architecture.html), 它包含以下几条规则:

- 独立的框架
- 可测试性
- 独立的用户界面
- 独立的数据库
- 独立于任何外部组件

### 独立的框架
在实际的开发过程中，我们将使用各种框架。如果我们的业务和框架是深度耦合的，那么更改框架可能会非常困难。因此，我们需要一种模式来避免对框架的业务依赖

### 可测试性
当我们的业务很复杂时，测试将依赖于外部输入。很难独立地测试单个组件或业务，因为输入和输出严重依赖于其他组件。所以我们需要是可测试的

### Independent of UI


### Independent of Database


### Independent of any external agency

## Clean Arch
当我们使用clean arch ，可以这样来分层:
1. Domain Layer
2. Delivery Layer
3. Usercase Layer
4. Repository Layer
### Diagram
![golang clean architecture](https://github.com/luoshanjie/go-clean-arch/blob/main/doc/clean-arch.png)


## 不使用Clean Arch
如果我们不使用干净的arch，代码就会变得难以维护。最好的做法是不使用Clean Arch，然后随着功能的增加，代码就会退化，难以维护。然后，将Clean Arch用于重构，展示出使其易于维护和扩展的特点  

## 最佳实践
### 第一个目标
We will complete a simple functional Web service that will have the following capabilities
1. Add a user through interface (http://localhost/api/v1/user/add)
2. Modify user information through interface (http://localhost/api/v1/user/:id/update)
3. Search user information through interface (http://localhost/api/v1/user/:id/select)
4. Delete a user through interface (http://localhost/api/v1/user/:id/delete)

### Quick Finish
We did this quickly and without using Clean Arch: [lesson_one](https://github.com/luoshanjie/go-clean-arch/raw/main/lesson_one)


## 感谢
bxcodec's 给了我很多启迪(https://github.com/bxcodec/go-clean-arch), 我要感谢Uncle Bob给我们带来这么漂亮的设计。谢谢你，Uncle Bob和Bxcodec ^_^
