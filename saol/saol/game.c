#pragma once
#include "game.h"


//函数定义
void InitBoot(char board[HANS][LIES], int hans, int lies,char set)//数组初始化
{
	int i = 0;
	int j = 0;
	for (i=0;i<hans;i++)
	{
		for (j = 0; j < lies; j++)
		{
			board[i][j] = set;
		}
	}
}

void Display(char board[HANS][LIES], int han, int lie)//打印棋盘函数
{
	int i = 0;
	int j = 0;
	for (i = 0; i <= lie; i++)
	{
		printf("%d ",i);//列号
	}
	printf("\n");
	for (i = 1; i <= han; i++)//为什么是i=1不是0，因为我们是打印棋盘的内围不是外围，外围是下标0-10,内围是1-9,所以我们从标1开始
	{
		printf("%d ", i);//行号
		for (j = 1; j <= lie; j++)
		{
			printf("%c ", board[i][j]);
		}
		printf("\n");
	}
}

void SetMine(char board[HANS][LIES], int han, int lie)//布置雷的函数
{
	int count = BUZL;//雷的数量
	while (count)
	{
		int x = rand()% han+1;//+1是随机生成1-9的数，不加1就是生成0-8的数//布置雷的随机坐标生成
		int y = rand()% lie+1;//+1是随机生成1-9的数，不加1就是生成0-8的数//布置雷的随机坐标生成
		if (board[x][y]=='0')
		{
			board[x][y] = '1';//雷的图标代表
				count--;//必须自减否则放的雷不止10个
		}
	}
}

int get_mine_count(char board[HANS][LIES], int x, int y)//这是查找x和y坐标周围有几个雷的函数
{
	//'0'-'1'==0,//这里是字符0-字符1==数值1，原因在于ASCLL表里面他们字符代表的值
	//'0'-'3'==3
	//'0'-'2'==2
	//'0'+'0'==0
	return board[x - 1][y] +board[x - 1][y - 1] +board[x][y - 1] +board[x + 1][y - 1] +board[x + 1][y] +board[x + 1][y + 1] +board[x][y + 1] +board[x - 1][y + 1] - 8 * '0';
	//这是把玩家输入的坐标周围计算有没有雷，并且有几个
}

//int xiaochu(char board[HANS][LIES], char show[HANS][LIES], int x, int y)
//{
//	int b = 0;
//	if (0==(board[x - 1][y] + board[x - 1][y - 1] + board[x][y - 1] + board[x + 1][y - 1] + board[x + 1][y] + board[x + 1][y + 1] + board[x][y + 1] + board[x - 1][y + 1] - 8 * '0'))
//	{
//
//		b = xiaochu(board, show, x-1, y);//1
//			if (b == 1)
//			{
//				show[x - 1][y] = ' ';
//			}
//		b = xiaochu(board, show, x - 1, y-1);//2
//			if (b == 1)
//			{
//				show[x - 1][y-1] = ' ';
//			}
//		b = xiaochu(board, show, x, y-1);//3
//			if (b == 1)
//			{
//				show[x][y - 1] = ' ';
//			}
//
//		//b = xiaochu(board, show, x+1, y - 1);//4
//		//if (y >= 1 && y <= 9)
//		//{
//		//	if (b == 1)
//		//	{
//		//		show[x+1][y - 1] = ' ';
//		//	}
//		//}
//		//b = xiaochu(board, show, x+1, y );//5
//		//if (y >= 1 && y <= 9)
//		//{
//		//	if (b == 1)
//		//	{
//		//		show[x+1][y] = ' ';
//		//	}
//		//}
//		//b = xiaochu(board, show, x+1, y + 1);//6
//		//if (y >= 1 && y <= 9)
//		//{
//		//	if (b == 1)
//		//	{
//		//		show[x+1][y + 1] = ' ';
//		//	}
//		//}
//		//b = xiaochu(board, show, x, y+1);//7
//		//if (y >= 1 && y <= 9)
//		//{
//		//	if (b == 1)
//		//	{
//		//		show[x][y + 1] = ' ';
//		//	}
//		//}
//		//b = xiaochu(board, show, x-1, y + 1);//8
//		//if (y >= 1 && y <= 9)
//		//{
//		//	if (b == 1)
//		//	{
//		//		show[x-1][y + 1] = ' ';
//		//	}
//		//}
//		return 1;
//	}
//	else
//	{
//		int count = get_mine_count(board, x, y);
//		show[x][y] = count + '0';
//		//Display(show, han, lie);
//		return 0;
//	}
//}

void FindMine(char board[HANS][LIES], char show[HANS][LIES], int han, int lie)//玩家输入坐标排查雷的实现函数
{
	int x = 0;
	int y = 0;
	int win = 0;
	//9*9-10==71
	while (win<han*lie- BUZL)//这里判断是为判断你输入的坐标如果等于71了也就代表你排查了所有的位置，除了10个雷之外
	{
		printf("请输入排查雷的坐标:>");
		int b = scanf("%d%d", &x, &y);//内围的坐标是1-9
		if (x >= 1 && x <= han && y >= 1 && y <= lie)
		{
			//坐标合法有效
			if (board[x][y]=='1')//1.这是踩雷
			{
				printf("抱歉你踩到雷了，游戏结束\n");
				Display(board, han, lie);//这是游戏结束后把棋盘打印出来//这里打印的是没有被蒙住的棋盘
				break;
			}
			else
			{
				
				//计算你想x，y周围的雷有几个，并且打印出来
				int count=get_mine_count(board,x,y);
				/*int b = 0;
				if (count==0)
				{
					b=xiaochu(board, show, x, y);
					if (b==1)
					{
						show[x][y] = ' ';
					}
					Display(show, han, lie);
				}
				else
				{*/
				    show[x][y] = count + '0';
				    Display(show, han, lie);//这里是打印那种全部是被蒙住的棋盘(已经输入的坐标不会被蒙住)
				//}
				win++;//计数器
			}
		}
		else
		{
			printf("输入坐标非法，请重新输入\n");
		}
		if (win== han * lie - BUZL)//这里判断是为了防止玩家输了走到这
		{
			printf("恭喜你，排雷成功，赢得游戏\n");
			Display(board, han, lie);//这里打印雷的棋盘是为了给玩家看
		}
	}
}
