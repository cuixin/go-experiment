Goroutine的内存泄漏测试
======  

测试目的
------------
Golang目前不适合开很多Goroutine，另外Goroutine的确是hold住内存不释放。

测试步骤
------------
```
GODEBUG=gctrace=1 go run main.go
```  
```
ctrl+c
```

测试环境
------------
go 1.2  
darwin 64bit  
4g ram  

测试结果和结论
------------
Goroutine在开了40w个后，ctrl+c终止程序后，看到内存还是有100m不归还给操作系统，目前
这个估计1.3会有改变，你可以试试测试结果。

```
 goroutine-memory-leak git:(master) ✗ GODEBUG=gctrace=1 go run main.go
gc1(1): 0+0+0 ms, 0 -> 0 MB 16 -> 18 (19-1) objects, 0(0) handoff, 0(0) steal, 0/0/0 yields
gc2(1): 0+0+0 ms, 0 -> 0 MB 29 -> 29 (30-1) objects, 0(0) handoff, 0(0) steal, 0/0/0 yields
gc3(1): 0+0+0 ms, 0 -> 0 MB 864 -> 492 (865-373) objects, 0(0) handoff, 0(0) steal, 0/0/0 yields
gc4(1): 0+0+0 ms, 0 -> 0 MB 2468 -> 1480 (2841-1361) objects, 0(0) handoff, 0(0) steal, 0/0/0 yields
gc5(1): 0+0+0 ms, 0 -> 0 MB 5086 -> 2990 (6447-3457) objects, 0(0) handoff, 0(0) steal, 0/0/0 yields
gc6(1): 0+0+0 ms, 1 -> 0 MB 6188 -> 3135 (9645-6510) objects, 0(0) handoff, 0(0) steal, 0/0/0 yields
gc7(1): 0+0+0 ms, 1 -> 0 MB 8190 -> 2826 (14700-11874) objects, 0(0) handoff, 0(0) steal, 0/0/0 yields
gc8(1): 0+0+0 ms, 1 -> 0 MB 7118 -> 2857 (18992-16135) objects, 0(0) handoff, 0(0) steal, 0/0/0 yields
gc9(1): 0+0+0 ms, 1 -> 0 MB 8158 -> 3477 (24293-20816) objects, 0(0) handoff, 0(0) steal, 0/0/0 yields
gc10(1): 0+0+0 ms, 1 -> 0 MB 9096 -> 3630 (29912-26282) objects, 0(0) handoff, 0(0) steal, 0/0/0 yields
gc11(1): 0+0+0 ms, 1 -> 0 MB 9339 -> 4155 (35621-31466) objects, 0(0) handoff, 0(0) steal, 0/0/0 yields
gc1(1): 0+0+0 ms, 0 -> 0 MB 16 -> 18 (19-1) objects, 0(0) handoff, 0(0) steal, 0/0/0 yields
gc2(1): 0+0+0 ms, 0 -> 0 MB 34 -> 34 (35-1) objects, 0(0) handoff, 0(0) steal, 0/0/0 yields
gc3(3): 0+0+0 ms, 0 -> 0 MB 182 -> 157 (183-26) objects, 2(27) handoff, 6(27) steal, 51/10/0 yields
gc4(8): 0+0+0 ms, 0 -> 0 MB 991 -> 577 (1026-449) objects, 0(0) handoff, 17(57) steal, 262/27/0 yields
gc5(8): 0+0+0 ms, 1 -> 1 MB 5047 -> 2815 (5500-2685) objects, 0(0) handoff, 18(127) steal, 253/24/0 yields
gc6(8): 1+0+0 ms, 2 -> 2 MB 11286 -> 7051 (13973-6922) objects, 0(0) handoff, 36(285) steal, 245/18/0 yields
gc7(8): 4+0+1 ms, 4 -> 4 MB 23658 -> 15355 (30582-15227) objects, 0(0) handoff, 35(358) steal, 234/19/0 yields
gc8(8): 9+0+3 ms, 9 -> 9 MB 47588 -> 31471 (62817-31346) objects, 0(0) handoff, 40(811) steal, 254/19/0 yields
gc9(8): 20+0+6 ms, 18 -> 18 MB 93911 -> 62690 (125259-62569) objects, 0(0) handoff, 25(1010) steal, 228/18/0 yields
gc10(8): 41+0+12 ms, 36 -> 35 MB 183604 -> 123145 (246177-123032) objects, 0(0) handoff, 40(3113) steal, 262/32/0 yields
gc11(8): 81+0+23 ms, 71 -> 69 MB 364107 -> 243625 (487142-243517) objects, 6(239) handoff, 86(7706) steal, 295/45/0 yields
2014/01/16 11:46:58 You have created  400000 routines
^C2014/01/16 11:47:03 interrupt
gc12(8): 114+3+65 ms, 124 -> 115 MB 960025 -> 402531 (1203545-801014) objects, 25(770) handoff, 37(8966) steal, 318/24/0 yields
gc13(8): 109+1+29 ms, 115 -> 115 MB 402531 -> 400355 (1203545-803190) objects, 13(505) handoff, 108(8311) steal, 295/27/0 yields
scvg-1: 10 MB released
scvg-1: inuse: 116, idle: 10, sys: 127, released: 10, consumed: 116 (MB)
2014/01/16 11:47:09 Current Routines 365337
2014/01/16 11:47:09 You have created  800000 routines
gc14(8): 20+1+12 ms, 115 -> 115 MB 400360 -> 400350 (1203556-803206) objects, 29(776) handoff, 74(14043) steal, 356/161/127 yields
gc15(8): 17+1+15 ms, 115 -> 115 MB 400350 -> 400347 (1203556-803209) objects, 31(751) handoff, 53(6533) steal, 383/76/0 yields
scvg-1: 0 MB released
scvg-1: inuse: 117, idle: 10, sys: 127, released: 10, consumed: 117 (MB)
2014/01/16 11:47:11 Current Routines 4
gc16(8): 17+0+13 ms, 115 -> 115 MB 400349 -> 400347 (1203558-803211) objects, 23(534) handoff, 69(6992) steal, 263/87/14 yields
gc17(8): 16+1+12 ms, 115 -> 115 MB 400347 -> 400347 (1203558-803211) objects, 34(900) handoff, 66(7187) steal, 331/31/0 yields
scvg-1: 0 MB released
scvg-1: inuse: 117, idle: 10, sys: 127, released: 10, consumed: 117 (MB)
2014/01/16 11:47:13 Current Routines 4
gc18(8): 17+1+18 ms, 115 -> 115 MB 400349 -> 400347 (1203560-803213) objects, 27(654) handoff, 36(2057) steal, 353/42/0 yields
gc19(8): 18+0+13 ms, 115 -> 115 MB 400347 -> 400347 (1203560-803213) objects, 20(456) handoff, 66(10725) steal, 305/94/14 yields
scvg-1: 0 MB released
scvg-1: inuse: 117, idle: 10, sys: 127, released: 10, consumed: 117 (MB)
2014/01/16 11:47:15 Current Routines 4
gc20(8): 17+1+14 ms, 115 -> 115 MB 400349 -> 400347 (1203562-803215) objects, 18(906) handoff, 42(5250) steal, 313/96/14 yields
gc21(8): 17+1+14 ms, 115 -> 115 MB 400347 -> 400347 (1203562-803215) objects, 26(1316) handoff, 65(8977) steal, 300/94/7 yields
scvg-1: 0 MB released
scvg-1: inuse: 117, idle: 10, sys: 127, released: 10, consumed: 117 (MB)
2014/01/16 11:47:17 Current Routines 4
gc22(8): 16+1+12 ms, 115 -> 115 MB 400349 -> 400347 (1203564-803217) objects, 23(669) handoff, 67(5613) steal, 327/30/0 yields
gc23(8): 16+1+13 ms, 115 -> 115 MB 400347 -> 400347 (1203564-803217) objects, 39(1205) handoff, 57(9301) steal, 378/46/0 yields
scvg-1: 0 MB released
scvg-1: inuse: 117, idle: 10, sys: 127, released: 10, consumed: 117 (MB)
2014/01/16 11:47:19 Current Routines 4
gc24(8): 17+0+13 ms, 115 -> 115 MB 400349 -> 400347 (1203566-803219) objects, 18(414) handoff, 25(1721) steal, 331/28/0 yields
gc25(8): 19+0+12 ms, 115 -> 115 MB 400347 -> 400347 (1203566-803219) objects, 25(797) handoff, 76(16281) steal, 372/164/121 yields
scvg-1: 0 MB released
scvg-1: inuse: 117, idle: 10, sys: 127, released: 10, consumed: 117 (MB)
2014/01/16 11:47:21 Current Routines 4
gc26(8): 17+0+13 ms, 115 -> 115 MB 400349 -> 400347 (1203568-803221) objects, 17(575) handoff, 83(17648) steal, 303/92/14 yields
gc27(8): 17+0+12 ms, 115 -> 115 MB 400347 -> 400347 (1203568-803221) objects, 16(504) handoff, 18(566) steal, 362/40/0 yields
scvg-1: 0 MB released
scvg-1: inuse: 117, idle: 10, sys: 127, released: 10, consumed: 117 (MB)
2014/01/16 11:47:23 Current Routines 4
gc28(8): 16+1+13 ms, 115 -> 115 MB 400349 -> 400347 (1203570-803223) objects, 15(400) handoff, 55(9595) steal, 316/94/14 yields
gc29(8): 16+1+12 ms, 115 -> 115 MB 400347 -> 400347 (1203570-803223) objects, 32(762) handoff, 73(12985) steal, 354/101/14 yields
scvg-1: 0 MB released
scvg-1: inuse: 117, idle: 10, sys: 127, released: 10, consumed: 117 (MB)
2014/01/16 11:47:25 Current Routines 4
gc30(8): 16+0+12 ms, 115 -> 115 MB 400349 -> 400347 (1203572-803225) objects, 18(335) handoff, 46(3138) steal, 343/32/0 yields
gc31(8): 16+1+12 ms, 115 -> 115 MB 400347 -> 400347 (1203572-803225) objects, 27(812) handoff, 44(4996) steal, 351/93/7 yields
scvg-1: 0 MB released
scvg-1: inuse: 117, idle: 10, sys: 127, released: 10, consumed: 117 (MB)
2014/01/16 11:47:27 Current Routines 4
```