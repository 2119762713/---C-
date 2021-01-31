#include "game.h"




void menu()//菜单函数
{
	printf("***********************\n");
	printf("**** 1.play 0.exit ****\n");
	printf("***********************\n");
}


	/*  
	   |   |
	---|---|---
	   |   |
	---|---|---
	   |   |
	 *///这就是我要打印的棋盘


void game()//游戏的整个算法实现函数
{
	char ret = 0;
	//数组 - 存放你输入的棋盘信息
	char board[ROW][COL] = {0};//这怎么把数组里面全部初始化空格
	//初始化棋盘
	Init(board,ROW,COL);//把棋盘初始化了
	//打印棋盘
	DisplayBoard(board, ROW, COL);//把棋盘打印出来
	//下棋
	while (1)
	{
		//玩家先下棋，再电脑下棋
		//玩家下棋
		PlayerMove(board, ROW, COL);//玩家下棋函数
		DisplayBoard(board, ROW, COL);//把棋盘打印出来给玩家看看直接下的棋盘样子
		//判断玩家是否赢
		ret = RetBOOT(board, ROW, COL);
		//判断函数传过来的值是什么，只要不是C继续就break跳出
		if (ret !='C')
		{
			break;
		}
		//电脑下棋
		ComputerMove(board, ROW, COL);
		DisplayBoard(board, ROW, COL);//把棋盘打印出来给玩家看看直接下的棋盘样子
		//判断玩电脑是否赢
		ret = RetBOOT(board, ROW, COL);//判断函数传过来的值是什么，只要不是C继续就break跳出
		if (ret != 'C')
		{
			break;
		}
	}
	if (ret=='*')
	{
		printf("玩家赢\n");
	}
	else if (ret=='#')
	{
		printf("电脑赢\n");
	}
	else
	{
		printf("平局\n");
	}
}

void test()
{
	int input = 0;
	srand((unsigned int)time(NULL));//把时间戳用来当作随机数生成器
	do
	{
		menu();//调用菜单函数
		printf("请选择>：");
		int b=scanf("%d",&input);
		switch (input)
		{
		case 1:
			game();
			printf("三子棋\n");
			break;
		case 0:
			printf("退出游戏\n");
			break;
		default:
			printf("选择错误，请重新选择!\n");
			break;
		}
	}while(input);
}

void main()
{
	test();//调用游戏
}