#include <stdio.h>



//==========================
//==========================
char* my_strcpy(char* dest, const char* src);
//1.дһ������ָ��pf,�ܹ�ָ��my_ strcpy
//2.дһ������ָ������pfArr. �ܹ����4��my_ _strcpy�����ĵ�ַ
void main()
{
	char* (*pf)(char*, const char*) = my_strcpy;
	char* (*pf[4])(char*, const char*) = { my_strcpy ,my_strcpy ,my_strcpy ,my_strcpy };
}
//==========================
////����ָ��
//void main()
//{
//	void(* signal(int, void(*)(int)) )(int);
//	typedef void(*pun_t)(int);
//	pun_t signal(int, pun_t);//�����͵�һ����һ������˼
//}