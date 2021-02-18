#include <stdio.h>



//==========================
//==========================
char* my_strcpy(char* dest, const char* src);
//1.写一个函数指针pf,能够指向my_ strcpy
//2.写一个函数指针数组pfArr. 能够存放4个my_ _strcpy函数的地址
void main()
{
	char* (*pf)(char*, const char*) = my_strcpy;
	char* (*pf[4])(char*, const char*) = { my_strcpy ,my_strcpy ,my_strcpy ,my_strcpy };
}
//==========================
////函数指针
//void main()
//{
//	void(* signal(int, void(*)(int)) )(int);
//	typedef void(*pun_t)(int);
//	pun_t signal(int, pun_t);//这个金和第一个是一样的意思
//}