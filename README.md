# golang-msvc
golang生成支持MSVC调用的dll

# 目录结构
## capi
golang导出cgo的package

libcapi.def: MSVC的模块定义文件, 用于编写导出的方法, 格式如下
```def
EXPORTS
    add
    md5sum
```

## call_capi
c++调用golang的capi示例, 调用 golang 导出 capi 的 add 和 md5sum 方法


# 步骤
## go 编译 c-archive
编译输出到 call_capi 项目的 libcapi/libcapi.a, 同时包含头文件 libcapi/libcapi.h
```sh
go build -buildmode=c-archive -ldflags "-s -w" -o ./call_capi/libcapi/libcapi.a ./capi
```

## 生成 dll
使用 MinGW 的 gcc, 根据 capi/libcapi.def 将 libcapi/libcapi.a 转换为 libcapi/libcapi.dll
```sh
gcc ./capi/libcapi.def ./call_capi/libcapi/libcapi.a -shared -lwinmm -lWs2_32 -o ./call_capi/libcapi/libcapi.dll
```

## 生成 MSVC 支持的 lib
使用 MSVC 的 lib 命令, 根据 capi/libcapi.def 在 call_capi 项目生成 libcapi/libcapi.lib
```sh
lib /def:./capi/libcapi.def /name:libcapi.dll /out:./call_capi/libcapi/libcapi.lib /MACHINE:X64
```
如果需要生成32位的lib, 将`/MACHINE:X64`替换为`/MACHINE:X86`

## 修改 libcapi/libcapi.h
golang 生成的头文件部分语法不支持 MSVC, 需要部分修改。

将`__SIZE_TYPE__`替换为`size_t`

将`float _Complex`替换为`_Fcomplex`, 将`double _Complex`替换为`_Dcomplex`, 并在代码之前加上`#include <complex.h>`

```c
...
typedef size_t GoUintptr;
...
#include <complex.h>
typedef _Fcomplex GoComplex64;
typedef _Dcomplex GoComplex128;
```

也可使用 sed 命令进行替换
```sh
sed -i "s/__SIZE_TYPE__/size_t/g" ./call_capi/libcapi/libcapi.h
sed -i "s/float\ _Complex/_Fcomplex/g" ./call_capi/libcapi/libcapi.h
sed -i "s/double\ _Complex/_Dcomplex/g" ./call_capi/libcapi/libcapi.h
sed -i "/GoComplex64/i #include\ <complex.h>" ./call_capi/libcapi/libcapi.h
```

## 编译 call_capi
```sh
mkdir out
cmake -B out ./call_capi
cmake --build out --config Debug
cmake --install out --config Debug
```

编译生成的 exe 文件默认安装至 `C:\Program Files (x86)\call_capi`

## 执行结果
```
hello
add(1, 3) = 4
md5sum("hello world") = 5eb63bbbe01eeed093cb22bb8f5acdc3
```
