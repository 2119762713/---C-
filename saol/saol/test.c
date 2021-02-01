#include "game.h"


void menu()//菜单函数
{
	printf("***********************\n");
	printf("**** 1.play 0.exit ****\n");
	printf("***********************\n");
}

void game()
{
	printf("扫雷游戏\n");
	//11*11
	char board[HANS][LIES] = {0};//雷的信息存储//这里放常量变量名，以后就可以改常量就可以了，没必要去改执行算法里面的行和列
	char show[HANS][LIES] = {0};//排查雷的信息存储
	InitBoot(board, HANS, LIES,'0');//初始化数组
	InitBoot(show, HANS, LIES,'*');//初始化数组
	//打印棋盘
	//Display(board, HAN, LIE);
	Display(show, HAN, LIE);
	//布置雷--随机
	SetMine(board, HAN, LIE);
	//Display(board, HAN, LIE);////雷的棋盘不会打印，要不然直接被看了，这只是给你自己看的
	//扫雷---随便打印棋盘
	FindMine(board, show, HAN, LIE);
}

void test()
{
	int input = 0;
	srand((unsigned int)time(NULL));//把时间戳返回的类型强制转换成int型//随机数rand必须要的随机数生成器srand
	do//先执行一次游戏
	{
		menu();//菜单函数
		printf("请选择：>");
		int b = scanf("%d", &input);
		switch (input)
		{
		case 1:
			game();//实现游戏的函数
			break;
		case 0:
			printf("已退出游戏\n");
			break;
		default:
			printf("输入错误，请重新输入\n");
			break;
		}
	} while (input);
}

void main()
{
	test();//游戏函数实现的地方
}