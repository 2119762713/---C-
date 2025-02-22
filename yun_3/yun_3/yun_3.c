#include <stdio.h>

//================================================================
//以及他们所占存储空间的大小。类型的意义:
//1.使用这个类型开辟内存空间的大小(大小决定了使用范围)。
//2.如何看待内存空间的视角。

//原码、反码、补码
//计算机中的有符号数有三种表示方法，即原码、反码和补码。
//三种表示方法均有符号位和数值位两部分，符号位都是用0表示"正”，用1表示"负”，而数值位
//三种表示方法各不相同。
//原码直接将二进制按照正负数的形式翻译成二进制就可以。
//将原码的符号位不变，其他位依次按位取反就可以得到了。
//将原码的符号位不变，其他位依次按位取反就可以得到了。
void main()
{
	int a = 20;
	//存储内存里面是：14000000
	//为什么是这个呢，因为32位里面4位转换1位字节
	//20的原码是0000 0000 0000 0000 0000 0000 0001 0100
	//所以是      0   0    0    0    0    0    1    4
	//为什么是倒着的呢，因为什么大端小端: 
	//大端(存储)模式,是指数据的低位保存在内存的高地址中,而数据的高位,保存在内存的低地址中;
	//小端(存储)模式,是指数据的低位保存在内存的低地址中,而数据的高位, ,保存在内存的高地址中。
}