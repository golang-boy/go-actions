package main

// func mul(a, b int) int {
// 	return a * b
// }

// func main() {
// 	mul(3, 4)

// }

func main() {

	var a1, a2, a3 int = 1, 2, 3
	A(a1, a2, a3)

}

func A(a, b, c int) int {
	sum := a + b + c
	return sum

}

//  栈结构
//    栈基地址
//     局部变量
//     返回值
//     参数
//    返回地址

// 老版的都在栈上传递，参数与返回值
// 新版参数9个以内的参数或返回值，由寄存器传递，超过9个的参数或返回值，由栈传递。返回地址由被调用者保存

// GOTRACEBACK=1 go run test3.go 生成core dump文件，使用dlv或者gdb调试

// 栈扩容,与栈转移

// go tool compile -S -N -L .\test3.go

// 禁止编译器优化和函数内联

//  main.mul STEXT nosplit size=40 args=0x10 locals=0x10 funcid=0x0 align=0x0
//          0x0000 00000 (C:/Users/86186/workspace/go-actions/go-actions/trace/test/test3.go:3)     TEXT    main.mul(SB), NOSPLIT|ABIInternal, $16-16
//          0x0000 00000 (C:/Users/86186/workspace/go-actions/go-actions/trace/test/test3.go:3)     PUSHQ   BP
//          0x0001 00001 (C:/Users/86186/workspace/go-actions/go-actions/trace/test/test3.go:3)     MOVQ    SP, BP
//          0x0004 00004 (C:/Users/86186/workspace/go-actions/go-actions/trace/test/test3.go:3)     SUBQ    $8, SP
//          0x0008 00008 (C:/Users/86186/workspace/go-actions/go-actions/trace/test/test3.go:3)     FUNCDATA        $0, gclocals·g2BeySu+wFnoycgXfElmcg==(SB)
//          0x0008 00008 (C:/Users/86186/workspace/go-actions/go-actions/trace/test/test3.go:3)     FUNCDATA        $1, gclocals·g2BeySu+wFnoycgXfElmcg==(SB)
//          0x0008 00008 (C:/Users/86186/workspace/go-actions/go-actions/trace/test/test3.go:3)     FUNCDATA        $5, main.mul.arginfo1(SB)
//          0x0008 00008 (C:/Users/86186/workspace/go-actions/go-actions/trace/test/test3.go:3)     MOVQ    AX, main.a+24(SP)
//          0x000d 00013 (C:/Users/86186/workspace/go-actions/go-actions/trace/test/test3.go:3)     MOVQ    BX, main.b+32(SP)
//          0x0012 00018 (C:/Users/86186/workspace/go-actions/go-actions/trace/test/test3.go:3)     MOVQ    $0, main.~r0(SP)
//          0x001a 00026 (C:/Users/86186/workspace/go-actions/go-actions/trace/test/test3.go:4)     IMULQ   BX, AX
//          0x001e 00030 (C:/Users/86186/workspace/go-actions/go-actions/trace/test/test3.go:4)     MOVQ    AX, main.~r0(SP)
//          0x0022 00034 (C:/Users/86186/workspace/go-actions/go-actions/trace/test/test3.go:4)     ADDQ    $8, SP
//          0x0026 00038 (C:/Users/86186/workspace/go-actions/go-actions/trace/test/test3.go:4)     POPQ    BP
//          0x0027 00039 (C:/Users/86186/workspace/go-actions/go-actions/trace/test/test3.go:4)     RET
//          0x0000 55 48 89 e5 48 83 ec 08 48 89 44 24 18 48 89 5c  UH..H...H.D$.H.\
//          0x0010 24 20 48 c7 04 24 00 00 00 00 48 0f af c3 48 89  $ H..$....H...H.
//          0x0020 04 24 48 83 c4 08 5d c3                          .$H...].
//  main.main STEXT nosplit size=42 args=0x0 locals=0x20 funcid=0x0 align=0x0
//          0x0000 00000 (C:/Users/86186/workspace/go-actions/go-actions/trace/test/test3.go:7)     TEXT    main.main(SB), NOSPLIT|ABIInternal, $32-0     //  当前栈帧被分配的字节数，0表示函数参数与返回值的字节数，因为没返回值，所以是0
//          0x0000 00000 (C:/Users/86186/workspace/go-actions/go-actions/trace/test/test3.go:7)     PUSHQ   BP    // 保存当前bp值, 返回的地址, bp是栈基地址
//          0x0001 00001 (C:/Users/86186/workspace/go-actions/go-actions/trace/test/test3.go:7)     MOVQ    SP, BP  // 把sp赋值给bp
//          0x0004 00004 (C:/Users/86186/workspace/go-actions/go-actions/trace/test/test3.go:7)     SUBQ    $24, SP   // sp减24，分配24字节栈空间, 当前sp在32字节处，分配24字节后，sp在8字节处, 从高地址往低地址分配
//          0x0008 00008 (C:/Users/86186/workspace/go-actions/go-actions/trace/test/test3.go:7)     FUNCDATA        $0, gclocals·g2BeySu+wFnoycgXfElmcg==(SB)
//          0x0008 00008 (C:/Users/86186/workspace/go-actions/go-actions/trace/test/test3.go:7)     FUNCDATA        $1, gclocals·g2BeySu+wFnoycgXfElmcg==(SB)
//          0x0008 00008 (C:/Users/86186/workspace/go-actions/go-actions/trace/test/test3.go:8)     MOVQ    $3, main.a+16(SP)
//          0x0011 00017 (C:/Users/86186/workspace/go-actions/go-actions/trace/test/test3.go:8)     MOVQ    $4, main.b+8(SP)
//          0x001a 00026 (C:/Users/86186/workspace/go-actions/go-actions/trace/test/test3.go:8)     MOVQ    $12, main.~r0(SP)   // 返回值放到了寄存器
//          0x0022 00034 (C:/Users/86186/workspace/go-actions/go-actions/trace/test/test3.go:8)     JMP     36         // 调用函数
//          0x0024 00036 (C:/Users/86186/workspace/go-actions/go-actions/trace/test/test3.go:10)    ADDQ    $24, SP    // 栈帧释放，栈收缩，sp加24，回到32字节处
//          0x0028 00040 (C:/Users/86186/workspace/go-actions/go-actions/trace/test/test3.go:10)    POPQ    BP
//          0x0029 00041 (C:/Users/86186/workspace/go-actions/go-actions/trace/test/test3.go:10)    RET
//          0x0000 55 48 89 e5 48 83 ec 18 48 c7 44 24 10 03 00 00  UH..H...H.D$....

//  (base) PS C:\Users\86186\workspace\go-actions\go-actions\trace\test> go tool compile -S -N -L .\test3.go
//  main.main STEXT nosplit size=87 args=0x0 locals=0x48 funcid=0x0 align=0x0
//          0x0000 00000 (C:/Users/86186/workspace/go-actions/go-actions/trace/test/test3.go:12)    TEXT    main.main(SB), NOSPLIT|ABIInternal, $72-0
//          0x0000 00000 (C:/Users/86186/workspace/go-actions/go-actions/trace/test/test3.go:12)    PUSHQ   BP
//          0x0001 00001 (C:/Users/86186/workspace/go-actions/go-actions/trace/test/test3.go:12)    MOVQ    SP, BP
//          0x0004 00004 (C:/Users/86186/workspace/go-actions/go-actions/trace/test/test3.go:12)    SUBQ    $64, SP
//          0x0008 00008 (C:/Users/86186/workspace/go-actions/go-actions/trace/test/test3.go:12)    FUNCDATA        $0, gclocals·g2BeySu+wFnoycgXfElmcg==(SB)
//          0x0008 00008 (C:/Users/86186/workspace/go-actions/go-actions/trace/test/test3.go:12)    FUNCDATA        $1, gclocals·g2BeySu+wFnoycgXfElmcg==(SB)
//          0x0008 00008 (C:/Users/86186/workspace/go-actions/go-actions/trace/test/test3.go:14)    MOVQ    $1, main.a1+48(SP)
//          0x0011 00017 (C:/Users/86186/workspace/go-actions/go-actions/trace/test/test3.go:14)    MOVQ    $2, main.a2+40(SP)
//          0x001a 00026 (C:/Users/86186/workspace/go-actions/go-actions/trace/test/test3.go:14)    MOVQ    $3, main.a3+32(SP)
//          0x0023 00035 (C:/Users/86186/workspace/go-actions/go-actions/trace/test/test3.go:15)    MOVQ    $1, main.a+56(SP)
//          0x002c 00044 (C:/Users/86186/workspace/go-actions/go-actions/trace/test/test3.go:15)    MOVQ    $2, main.b+24(SP)
//          0x0035 00053 (C:/Users/86186/workspace/go-actions/go-actions/trace/test/test3.go:15)    MOVQ    $3, main.c+16(SP)
//          0x003e 00062 (<unknown line number>)    NOP
//          0x003e 00062 (C:/Users/86186/workspace/go-actions/go-actions/trace/test/test3.go:20)    MOVQ    $6, main.sum+8(SP)
//          0x0047 00071 (C:/Users/86186/workspace/go-actions/go-actions/trace/test/test3.go:15)    MOVQ    $6, main.~r0(SP)
//          0x004f 00079 (C:/Users/86186/workspace/go-actions/go-actions/trace/test/test3.go:15)    JMP     81
//          0x0051 00081 (C:/Users/86186/workspace/go-actions/go-actions/trace/test/test3.go:17)    ADDQ    $64, SP
//          0x0055 00085 (C:/Users/86186/workspace/go-actions/go-actions/trace/test/test3.go:17)    POPQ    BP
//          0x0056 00086 (C:/Users/86186/workspace/go-actions/go-actions/trace/test/test3.go:17)    RET
//          0x0000 55 48 89 e5 48 83 ec 40 48 c7 44 24 30 01 00 00  UH..H..@H.D$0...
//          0x0010 00 48 c7 44 24 28 02 00 00 00 48 c7 44 24 20 03  .H.D$(....H.D$ .
//          0x0020 00 00 00 48 c7 44 24 38 01 00 00 00 48 c7 44 24  ...H.D$8....H.D$
//          0x0030 18 02 00 00 00 48 c7 44 24 10 03 00 00 00 48 c7  .....H.D$.....H.
//          0x0040 44 24 08 06 00 00 00 48 c7 04 24 06 00 00 00 eb  D$.....H..$.....
//          0x0050 00 48 83 c4 40 5d c3                             .H..@].
//  main.A STEXT nosplit size=54 args=0x18 locals=0x18 funcid=0x0 align=0x0
//          0x0000 00000 (C:/Users/86186/workspace/go-actions/go-actions/trace/test/test3.go:19)    TEXT    main.A(SB), NOSPLIT|ABIInternal, $24-24
//          0x0000 00000 (C:/Users/86186/workspace/go-actions/go-actions/trace/test/test3.go:19)    PUSHQ   BP
//          0x0001 00001 (C:/Users/86186/workspace/go-actions/go-actions/trace/test/test3.go:19)    MOVQ    SP, BP
//          0x0004 00004 (C:/Users/86186/workspace/go-actions/go-actions/trace/test/test3.go:19)    SUBQ    $16, SP
//          0x0008 00008 (C:/Users/86186/workspace/go-actions/go-actions/trace/test/test3.go:19)    FUNCDATA        $0, gclocals·g2BeySu+wFnoycgXfElmcg==(SB)
//          0x0008 00008 (C:/Users/86186/workspace/go-actions/go-actions/trace/test/test3.go:19)    FUNCDATA        $1, gclocals·g2BeySu+wFnoycgXfElmcg==(SB)
//          0x0008 00008 (C:/Users/86186/workspace/go-actions/go-actions/trace/test/test3.go:19)    FUNCDATA        $5, main.A.arginfo1(SB)
//          0x0008 00008 (C:/Users/86186/workspace/go-actions/go-actions/trace/test/test3.go:19)    MOVQ    AX, main.a+32(SP)
//          0x000d 00013 (C:/Users/86186/workspace/go-actions/go-actions/trace/test/test3.go:19)    MOVQ    BX, main.b+40(SP)
//          0x0012 00018 (C:/Users/86186/workspace/go-actions/go-actions/trace/test/test3.go:19)    MOVQ    CX, main.c+48(SP)
//          0x0017 00023 (C:/Users/86186/workspace/go-actions/go-actions/trace/test/test3.go:19)    MOVQ    $0, main.~r0(SP)
//          0x001f 00031 (C:/Users/86186/workspace/go-actions/go-actions/trace/test/test3.go:20)    LEAQ    (AX)(BX*1), DX
//          0x0023 00035 (C:/Users/86186/workspace/go-actions/go-actions/trace/test/test3.go:20)    LEAQ    (DX)(CX*1), AX
//          0x0027 00039 (C:/Users/86186/workspace/go-actions/go-actions/trace/test/test3.go:20)    MOVQ    AX, main.sum+8(SP)
//          0x002c 00044 (C:/Users/86186/workspace/go-actions/go-actions/trace/test/test3.go:21)    MOVQ    AX, main.~r0(SP)
//          0x0030 00048 (C:/Users/86186/workspace/go-actions/go-actions/trace/test/test3.go:21)    ADDQ    $16, SP
//          0x0034 00052 (C:/Users/86186/workspace/go-actions/go-actions/trace/test/test3.go:21)    POPQ    BP
//          0x0035 00053 (C:/Users/86186/workspace/go-actions/go-actions/trace/test/test3.go:21)    RET
