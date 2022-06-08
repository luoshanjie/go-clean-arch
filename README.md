# go-clean-arch

## What is Clean Arch
Clean Arch from [Uncle Bob](https://8thlight.com/blog/uncle-bob/2012/08/13/the-clean-architecture.html), Here are some rules:

1. Independent of Frameworks
2. Testable
3. Independent of UI
4. Independent of Database
5. Independent of any external agency

### Independent of Frameworks
In the actual development process, we will use various frameworks。If our business and frameworks are deeply coupled, it can be very difficult to change frameworks. So we need a pattern to avoid business dependence on frameworks

### Testable
When our business is complex, testing will depend on outside input. It would be difficult to test a single component or business independently because the inputs and outputs are heavily dependent on other components. So we need to be testable

### Independent of UI


### Independent of Database


### Independent of any external agency



## Don't use Clean Arch
If we don't use clean arch, Code becomes difficult to maintain as it becomes more functional. The best practice is not to use Clean Arch, and then as functionality increases, the code becomes decayed and difficult to maintain. The Clean Arch approach was then used for refactoring, making it easy to maintain and extend  


## Thanks
bxcodec's project inspired me a lot(https://github.com/bxcodec/go-clean-arch), and I'd like to thank Uncle Bob for bringing us such a beautiful design. Thank you Uncle Bob and Bxcodec ^_^
