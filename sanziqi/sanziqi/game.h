#define _CRT_SECURE_NO_WARNINGS 1//�������˼���ǣ�test.c��game.c��Ҫ����stdio.h�ļ���������game.h�ļ���������stdio.h�ļ�,Ȼ��test.c��game.c����game.h�ļ��͵���test.c��game.c������stdio.h�ļ�����Լ�˴�����
#include <stdio.h>//�������˼���ǣ�test.c��game.c��Ҫ����stdio.h�ļ���������game.h�ļ���������stdio.h�ļ�,Ȼ��test.c��game.c����game.h�ļ��͵���test.c��game.c������stdio.h�ļ���������Լ�˴�����
#include <stdlib.h>
#include <time.h>
#define ROW 3
#define COL 3

//��������
void Init(char board[ROW][COL],int row,int col);
//��������
void DisplayBoard(char board[ROW][COL],int row,int col);
//��������
void PlayerMove(char board[ROW][COL], int row, int col);
//��������
void ComputerMove(char board[ROW][COL], int row, int col);
//��������
char RetBOOT(char board[ROW][COL], int row, int col);
//��������������Ϸ��״̬
//���Ӯ ---'*'
//����Ӯ --- '#'
//ƽ��   --- 'Q'
//����  ---'C'