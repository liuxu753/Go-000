#学习笔记
1. 基于 errgroup 实现一个 http server 的启动和关闭 ，以及 linux signal 信号的注册和处理，要保证能够一个退出，全部注销退出。

作业核心在于使用errgroup.WithContext()
WithContext()会调用context.withCancel()，获取cancel
当某个goroutine返回error!=nil时，errgroup内部会记录这个错误(并使用sync.once包保证只会记录第一次错误信息）
并调用cancel()方法告诉其他goroutine，需要安全退出,其他goroutine只要保证select接收ctx.Done()这个chan的信号就可以实现了