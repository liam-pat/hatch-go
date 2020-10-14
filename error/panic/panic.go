package main

//主要功能
//1. 挺会当前函数的执行
//2. 一直往上返回，执行每一层的 defer
//3. 如果没有遇见 recover 程序退出
