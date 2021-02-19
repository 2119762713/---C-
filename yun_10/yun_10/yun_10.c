#define _CRT_SECURE_NO_WARNINGS 1
#include <stdio.h>





//==================================
//qsort -可以排序任意类型的数据
void main()
{
		
}
//==================================
//void menu()
//{
//	printf("*****************************\n");
//	printf("****    1.add    2.sub   ****\n");
//	printf("****    3.mul    4.div   ****\n");
//	printf("****        0.exit       ****\n");
//	printf("*****************************\n");
//}
//
//int add(int x,int y)
//{
//	return x + y;
//}
//
//int sub(int x, int y)
//{
//	return x - y;
//}
//
//int mul(int x, int y)
//{
//	return x * y;
//}
//
//int div(int x, int y)
//{
//	return x / y;
//}
//
//void calc(int (*pu)(int,int))
//{
//	int a = 0;
//	int a1 = 0;
//	printf("请输入两个操作数:>");
//	int b = scanf("%d%d", &a, &a1);
//	printf("%d\n", pu(a, a1));
//}
//
//void main()//计算器
//{
//	int b1 = 0;
//	do
//	{
//		menu();
//		printf("请选择:>");
//		int bb = scanf("%d", &b1);
//			switch (b1)
//			{
//			case 1:
//				calc(add);
//				break;
//			case 2:
//				calc(sub);
//				break;
//			case 3:
//				calc(mul);
//				break;
//			case 4:
//				calc(div);
//				break;
//			case 0:
//				printf("退出\n");
//				break;
//			default:
//				printf("输入错误，请重新输入\n");
//				break;
//			}
//	} while (b1);
//}

//void main()//计算器
//{
//	int a = 0;
//	int a1 = 0;
//	//pfArr是一个函数指针数组---通常叫转移表
//	int b1 = 0;
//	int (*put[5])(int, int) = {0,add,sub,mul,div};
//	do
//	{
//		menu();
//		printf("请选择:>");
//		int bb = scanf("%d", &b1);
//		if (b1 >= 1&&b1<=4)
//		{
//			printf("请输入两个操作数:>");
//			int b = scanf("%d%d", &a, &a1);
//			int ret = put[b1](a, a1);
//			printf("%d\n", ret);
//		}
//		else if (b1 == 0)
//		{
//			printf("退出\n");
//			return;
//		}
//		else
//		{
//			printf("输入错误，请重新输入:>");
//		}
//		
//	} while (b1);
//}


//void main()//计算器
//{
//	int a = 0;
//	int a1 = 0;
//	int b1 = 0;
//	do
//	{
//		menu();
//		printf("请选择:>");
//		int bb = scanf("%d", &b1);
//		if (b1 == 0) 
//		{
//			printf("退出\n");
//			break;
//		}
//			printf("请输入两个操作数:>");
//			int b = scanf("%d%d", &a, &a1);
//			switch (b1)
//			{
//			case 1:
//				printf("%d + %d = %d\n", a, a1, add(a, a1));
//				break;
//			case 2:
//				printf("%d - %d = %d\n", a, a1, sub(a, a1));
//				break;
//			case 3:
//				printf("%d * %d = %d\n", a, a1, mul(a, a1));
//				break;
//			case 4:
//				printf("%d / %d = %d\n", a, a1, div(a, a1));
//				break;
//			default:
//				printf("输入错误，请重新输入\n");
//				break;
//			}
//	} while (b1);
//}