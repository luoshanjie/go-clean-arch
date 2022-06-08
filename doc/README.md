# go-clean-arch

## 什么是干净的架构 
干净的架构来自于 [Uncle Bob](https://8thlight.com/blog/uncle-bob/2012/08/13/the-clean-architecture.html), 它包含以下几条规则:

- 独立的框架
- 可测试性
- 独立的用户界面
- 独立的数据库
- 独立于任何外部组件

### Independent of Frameworks
In the actual development process, we will use various frameworks。If our business and frameworks are deeply coupled, it can be very difficult to change frameworks. So we need a pattern to avoid business dependence on frameworks

### Testable
When our business is complex, testing will depend on outside input. It would be difficult to test a single component or business independently because the inputs and outputs are heavily dependent on other components. So we need to be testable

### Independent of UI


### Independent of Database


### Independent of any external agency

## Clean Arch
如果我们使用Clean Arch，我们可以这样分层
1. Domain Layer
2. Delivery Layer
3. Usercase Layer
4. Repository Layer
### Diagram
![golang clean architecture](https://github.com/luoshanjie/go-clean-arch/blob/main/doc/clean-arch.png)


## Don't use Clean Arch
If we don't use clean arch, Code becomes difficult to maintain as it becomes more functional. The best practice is not to use Clean Arch, and then as functionality increases, the code becomes decayed and difficult to maintain. The Clean Arch approach was then used for refactoring, making it easy to maintain and extend  

## 最佳实践
### First Target
We will complete a simple functional Web service that will have the following capabilities
1. Add a user through interface (http://localhost/api/v1/user/add)
2. Modify user information through interface (http://localhost/api/v1/user/:id/update)
3. Search user information through interface (http://localhost/api/v1/user/:id/select)
4. Delete a user through interface (http://localhost/api/v1/user/:id/delete)

### Quick Finish
We did this quickly and without using Clean Arch: [lesson_one](https://github.com/luoshanjie/go-clean-arch/raw/main/lesson_one)


## 感谢
bxcodec's 给了我很多启迪(https://github.com/bxcodec/go-clean-arch), 我要感谢Uncle Bob给我们带来这么漂亮的设计。谢谢你，Uncle Bob和Bxcodec ^_^
