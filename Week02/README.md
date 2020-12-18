# 学习笔记
## 作业
### 问题：
我们在数据库操作的时候，比如 dao 层中当遇到一个 sql.ErrNoRows 的时候，是否应该 Wrap 这个 error，抛给上层。为什么，应该怎么做请写出代码？
### 解答：
当遇到err时，最底层需要warp，中间层如果不需要做降级处理，则直接将error向上抛，最上层统一处理error即可
sql.ErrNoRows这个错误，dao层是业务层的最底层，因此需要wrap这个error，并向上抛，上层service层直接透传即可，api层进行统一的处理（日志记录及包装错误，接口返回错误码）
### 运行示例如下
```
query user failed,param：999,cause:get user failed id=999: sql ErrNoRows,err:=sql ErrNoRows
get user failed id=999
main/dao.(*Dao).QueryUserById
        F:/workspace/go进阶/go进阶/笔记/Go-000/Week02/dao/user.go:20
main/service.(*Service).QueryUserById
        F:/workspace/go进阶/go进阶/笔记/Go-000/Week02/service/user.go:21
main/api.QueryUserById
        F:/workspace/go进阶/go进阶/笔记/Go-000/Week02/api/user.go:13
main.main
        F:/workspace/go进阶/go进阶/笔记/Go-000/Week02/main.go:9
runtime.main
        C:/Go/src/runtime/proc.go:204
runtime.goexit
        C:/Go/src/runtime/asm_amd64.s:1374
Process finished with exit code 0

```
