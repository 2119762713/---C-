#define _CRT_SECURE_NO_WARNINGS 1
#include <stdio.h>




//=================================================
void main()
{
	int a = 10;
	int b = 20;
	int c = 30;
	int d = 40;
	int* arr[] = {&a,&b,&c,&d};
	int i = 0;
	for (i = 0; i < 4; i++)
	{
		printf("%d\n", *arr[i]);
	}
}
//=================================================
//void main()
//{
//	int arr1[10] = { 0 };
//	char arr2[5] = { 0 };
//	int* ch1[4] = {0};//�������ָ�������-ָ������
//	char* ch2[5] = {0};//����ַ�ָ�������-ָ������ 
//}
////=================================================
//void main()
//{
//	char arr1[] = "abcdef";
//	char arr2[] = "abcdef";
//	char* p1 = "abcdef";
//	char* p2 = "abcdef";//��������ΪϵͳΪ�˽�ʡ�ռ�,����ͬ�ĳ�������һ��,�����õĵ�ַҲ����һ����
//	if (p1==p2)
//	{
//		printf("hehe\n");
//	}
//	else
//	{
//		printf("haha\n");
//	}
//}