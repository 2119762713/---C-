#define _CRT_SECURE_NO_WARNINGS 1//这里的意思就是：test.c和game.c都要引入stdio.h文件，我们再game.h文件里面引入stdio.h文件,然后test.c和game.c引入game.h文件就等于test.c和game.c引入了stdio.h文件，节约了代码量
#include <stdio.h>//这里的意思就是：test.c和game.c都要引入stdio.h文件，我们再game.h文件里面引入stdio.h文件,然后test.c和game.c引入game.h文件就等于test.c和game.c引入了stdio.h文件，这样节约了代码量
#include <stdlib.h>
#include <time.h>
#define ROW 3
#define COL 3

//函数声明
void Init(char board[ROW][COL],int row,int col);
//函数声明
void DisplayBoard(char board[ROW][COL],int row,int col);
//函数声明
void PlayerMove(char board[ROW][COL], int row, int col);
//函数声明
void ComputerMove(char board[ROW][COL], int row, int col);
//函数声明
char RetBOOT(char board[ROW][COL], int row, int col);
//告诉我们四种游戏的状态
//玩家赢 ---'*'
//电脑赢 --- '#'
//平局   --- 'Q'
//继续  ---'C'