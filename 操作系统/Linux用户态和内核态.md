# Linux用户态和内核态

[toc]

参考：https://blog.csdn.net/m0_37199770/article/details/113482312

1. 用户态和内核态的本质
    ring 0被叫做内核态，完全在操作系统内核中运行
    ring 3被叫做用户态，在应用程序中运行

    shell和库函数（操作硬件的）就是为了屏蔽底层cpu指令集复杂实现细节，减轻程序员的负担，防止出错。

2. 用户态切换到内核态

- 系统调用：用户态进程主动切换到内核态的方式，用户态进程通过系统调用向操作系统申请资源完成工作，例如 fork（）就是一个创建新进程的系统调用；
- 异常：当 CPU 在执行用户态的进程时，发生了一些没有预知的异常，这时当前运行进程会切换到处理此异常的内核相关进程中，也就是切换到了内核态，如缺页异常；
- 外围设备中断：当 CPU 在执行用户态的进程时，外围设备完成用户请求的操作后，会向 CPU 发出相应的中断信号，这时 CPU 会暂停执行下一条即将要执行的指令，转到与中断信号对应的处理程序去执行，也就是切换到了内核态。如硬盘读写操作完成，系统会切换到硬盘读写的中断处理程序中执行后边的操作等；
