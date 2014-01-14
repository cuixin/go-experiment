Http的内存泄漏测试
======  

测试目的
------------
来源于stackoverflow的问题：
http://stackoverflow.com/questions/21080642/memory-leak-in-go-http-standard-library

测试步骤
------------
1. 先进行200并发的测试
```
go build
GODEBUG=gctrace=1 ./http-memory-leak
ab -k -n10000000 -c200 http://localhost:1234/hello
```  
2. 4000并发
```
go build
GODEBUG=gctrace=1 ./http-memory-leak
ab -k -n10000000 -c200 http://localhost:1234/hello
```  

测试环境
------------
go 1.2  
darwin 64bit  
4g ram  

测试结果和结论
------------
1. 在并发量在200以内的时候(ab -k -n10000000 -c200 http://localhost:1234/hello),有1m内存不释放。
```
gc372(8): 0+0+0 ms, 1 -> 1 MB 1003 -> 1003 (13203323-13202320) objects, 15(225) handoff, 23(583) steal, 314/41/0 yields
scvg-1: inuse: 1, idle: 5, sys: 7, released: 5, consumed: 1 (MB)
2014/01/14 10:41:58 Current Routines 6
```  
2. 在并发量在4000以内的时候(ab -k -n10000000 -c4000 http://localhost:1234/hello),10m内存始终释放不了。
```
gc40(8): 0+0+0 ms, 10 -> 10 MB 10052 -> 9942 (13550303-13540361) objects, 14(411) handoff, 39(4279) steal, 347/46/0 yields
scvg-1: 1 MB released
scvg-1: inuse: 15, idle: 88, sys: 104, released: 88, consumed: 15 (MB)
2014/01/14 10:44:18 Current Routines 6
```
3. 随着goroutine的增多，内存也在变大，这部分内存应该是routine的缓存，如果你超过这个数量，内存还会增加，从一个侧面说明goroutine并不随着GC的释放而释放内存。 
4. 有同学说看看netstat的状态，是不是epoll或者kqueue，没释放，准确的告诉你释放了，这个测试你们可以自己做做。  
5. 这个测试并不表明golang的http模块存在内存泄漏，而是goroutine的机制并不是用完就释放的，或者说一直重用的状态。  