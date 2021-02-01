#pragma once
#include "game.h"


//��������
void InitBoot(char board[HANS][LIES], int hans, int lies,char set)//�����ʼ��
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

void Display(char board[HANS][LIES], int han, int lie)//��ӡ���̺���
{
	int i = 0;
	int j = 0;
	for (i = 0; i <= lie; i++)
	{
		printf("%d ",i);//�к�
	}
	printf("\n");
	for (i = 1; i <= han; i++)//Ϊʲô��i=1����0����Ϊ�����Ǵ�ӡ���̵���Χ������Χ����Χ���±�0-10,��Χ��1-9,�������Ǵӱ�1��ʼ
	{
		printf("%d ", i);//�к�
		for (j = 1; j <= lie; j++)
		{
			printf("%c ", board[i][j]);
		}
		printf("\n");
	}
}

void SetMine(char board[HANS][LIES], int han, int lie)//�����׵ĺ���
{
	int count = BUZL;//�׵�����
	while (count)
	{
		int x = rand()% han+1;//+1���������1-9����������1��������0-8����//�����׵������������
		int y = rand()% lie+1;//+1���������1-9����������1��������0-8����//�����׵������������
		if (board[x][y]=='0')
		{
			board[x][y] = '1';//�׵�ͼ�����
				count--;//�����Լ�����ŵ��ײ�ֹ10��
		}
	}
}

int get_mine_count(char board[HANS][LIES], int x, int y)//���ǲ���x��y������Χ�м����׵ĺ���
{
	//'0'-'1'==0,//�������ַ�0-�ַ�1==��ֵ1��ԭ������ASCLL�����������ַ������ֵ
	//'0'-'3'==3
	//'0'-'2'==2
	//'0'+'0'==0
	return board[x - 1][y] +board[x - 1][y - 1] +board[x][y - 1] +board[x + 1][y - 1] +board[x + 1][y] +board[x + 1][y + 1] +board[x][y + 1] +board[x - 1][y + 1] - 8 * '0';
	//���ǰ���������������Χ������û���ף������м���
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

void FindMine(char board[HANS][LIES], char show[HANS][LIES], int han, int lie)//������������Ų��׵�ʵ�ֺ���
{
	int x = 0;
	int y = 0;
	int win = 0;
	//9*9-10==71
	while (win<han*lie- BUZL)//�����ж���Ϊ�ж�������������������71��Ҳ�ʹ������Ų������е�λ�ã�����10����֮��
	{
		printf("�������Ų��׵�����:>");
		int b = scanf("%d%d", &x, &y);//��Χ��������1-9
		if (x >= 1 && x <= han && y >= 1 && y <= lie)
		{
			//����Ϸ���Ч
			if (board[x][y]=='1')//1.���ǲ���
			{
				printf("��Ǹ��ȵ����ˣ���Ϸ����\n");
				Display(board, han, lie);//������Ϸ����������̴�ӡ����//�����ӡ����û�б���ס������
				break;
			}
			else
			{
				
				//��������x��y��Χ�����м��������Ҵ�ӡ����
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
				    Display(show, han, lie);//�����Ǵ�ӡ����ȫ���Ǳ���ס������(�Ѿ���������겻�ᱻ��ס)
				//}
				win++;//������
			}
		}
		else
		{
			printf("��������Ƿ�������������\n");
		}
		if (win== han * lie - BUZL)//�����ж���Ϊ�˷�ֹ��������ߵ���
		{
			printf("��ϲ�㣬���׳ɹ���Ӯ����Ϸ\n");
			Display(board, han, lie);//�����ӡ�׵�������Ϊ�˸���ҿ�
		}
	}
}
