#include "game.h"


//函数定义
void Init(char board[ROW][COL], int row, int col)
{
	for (int i=0;i<row;i++)
	{
		//这是给一行初始化空格
		for (int j = 0; j < col; j++)
		{
			//这是给一列初始化空格
			board[i][j] = ' ';//现在就给数组初始化空格了
		}
	}
}

//void DisplayBoard(char board[ROW][COL], int row, int col)//给他优化优化一下
//{
//	int i = 0;
//	for (i = 0; i < row; i++)
//	{
//		//这是打印一行数组元素
//		printf(" %c | %c | %c \n",board[i][0], board[i][1], board[i][2]);
//		if (i<row-1)//这是让他只打印两行美观，而不是打印三行不美观
//		{
//			//再打分割行
//			printf("---|---|---\n");
//	}
//
//	/*
//	   |   |
//	---|---|---
//	   |   |
//	---|---|---
//	   |   |
//	 *///这就是我要打印的棋盘
//}



void DisplayBoard(char board[ROW][COL], int row, int col)
{
	int i = 0;
	for (i = 0; i < row; i++)
	{
		int j = 0;
		for (j = 0; j < col; j++)
		{
			printf(" %c ", board[i][j]);//这是打印一行数组元素+|
			//这里是打印 数组元素数值 
			if (j < col - 1)//这是让他只打印两行美观，而不是打印三行不美观
				//这里判断是打印   |   |   ，而不不平等是打印   |   |   |
			{
				printf("|");
				//这里是打印    |      |
			}
		}
		//上面两加起来循环3次是 数值 | 数值 | 数值 
		printf("\n");
		if (i < row - 1)//这是让他只打印两行美观，而不是打印三行不美观
			//                  |   |                             |   |     
			//这里判断是打印 ---|---|---这个二行而不是打印三行 ---|---|---不美观
			//                  |   |                             |   |     
			//               ---|---|---                       ---|---|---
			//                  |   |                             |   |    
			//                                                 ---|---|--- 
		{
			for (j = 0; j < col; j++)
			{
				  //再打分割行
				  printf("---");
				  //这里是打印 ---
				if(j<col-1)//这是让他只打印两列美观，而不是打印三列不美观
					//这里判断是打印   |   |   ，而不不平等是打印   |   |   |
				  printf("|");
				//这里是打印   |   |
			}
			printf("\n");
			//上面加起来打印是---|---|---
		}
	}
	//上面全部加起来打印   |   |   
	//                  ---|---|---
	//                     |   |   
    //                  ---|---|---
	//                     |   |   
}//是不是很麻痹为什么要这么优化，因为但 你要玩10*10的棋盘你把常量改成10就行了，不用来改这些算法


void PlayerMove(char board[ROW][COL], int row, int col)//玩家下棋操作实现函数
{
	int x = 0;
	int y = 0;
	printf("玩家走!\n");
	while (1)//while循环就是为了 玩家输入错误坐标可以重新输入
	{
		printf("请输入要下的坐标:>");
		int b = scanf("%d%d", &x, &y);//这玩家输入的坐标你不知道是不是合法有效的，我们就需要判断了
		if (x >= 1 && x <= row && y >= 1 && y <= col)//这是判断玩家输入的坐标只能是1-3，为什么不是0呢，因为玩家不是程序员，不知道数组下标是从0开始的，一般都是觉得从1开始的
		{
			if (board[x - 1][y - 1] == ' ')//这里就是坐标里面初始化是空格就代表没有被下过//玩家输入的是坐标不是下标，所以要-1
			{
				board[x - 1][y - 1] = '*';
				break;
			}
			else
			{
				printf("此坐标已经下过棋了,不能重复下棋！\n");
			}
		}
		else
		{
			printf("坐标非法，请重新输入!\n");
		}
	}
}

void ComputerMove(char board[ROW][COL], int row, int col)//电脑下棋操作实现函数
{
	int x = 0;
	int y = 0;
	printf("电脑走!\n");
	while(1)
	{
		x = rand() % row;//这里是随机生成电脑下棋的坐标,取模row就是3,让他的随机值范围是0-2之间
		y = rand() % col;//这里是随机生成电脑下棋的坐标,取模col就是3,让他的随机值范围是0-2之间
		//上面生成的坐标不可能是非法的坐标，因为他们的随机值永远是0-2的下标，所以不用-1
		if (board[x][y] == ' ')
		{
			board[x][y] = '#';
			break;
		}
	}
}

int  IsFull(char board[ROW][COL], int row, int col)//判断是平局还是继续游戏，棋盘满了返回一个1，美满返回0
{
	int i = 0;
	int j = 0;
	for (i=0; i < row; i++)
	{
		for (j = 0; j < col; j++)
		{
			if (board[i][j] == ' ')//循环判断每行每列有没有空格如果有就是棋盘没满返回0，否则返回1
			{
				return 0;
			}
		}
	}
	return 1;//满了返回1
}

char RetBOOT(char board[ROW][COL], int row, int col)//判断电脑还是玩家输赢
{
	int i = 0;
	for (i = 0; i < row; i++)
	{
		////这是判断横三行                  ***
		//                ***
		//                       ***
		if (board[i][0] == board[i][1] && board[i][1] == board[i][2] && board[i][1] !=' ')
		{
			return board[i][1];
		}
	}

	for (i = 0; i < col; i++)
	{
		////这是判断竖三列*    *    *
		//                *    *    *
		//                *    *    *
		if (board[0][i] == board[1][i] && board[0][i] == board[2][i] && board[1][i] !=' ')
		{
			return board[1][i];
		}
	}
	//这两个if判断是判断对角线的    *               
	//                                *           
	//                                  *       
	if (board[0][0] == board[1][1] && board[1][1] == board[2][2] && board[1][1] != ' ')
	{
		return board[1][1];
	}
	//这两个if判断是判断对角线的                 *
	//                                         *
	//                                       *
	if (board[0][2] == board[1][1] && board[1][1] == board[2][0] && board[1][1] != ' ')
	{
		return board[1][1];
	}
	//判断是否平局
	if (1 == IsFull(board, ROW, COL))
	{
		return 'Q';
	}
		return 'C';
}