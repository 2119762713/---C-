#include "game.h"




void menu()//�˵�����
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
	 *///�������Ҫ��ӡ������


void game()//��Ϸ�������㷨ʵ�ֺ���
{
	char ret = 0;
	//���� - ����������������Ϣ
	char board[ROW][COL] = {0};//����ô����������ȫ����ʼ���ո�
	//��ʼ������
	Init(board,ROW,COL);//�����̳�ʼ����
	//��ӡ����
	DisplayBoard(board, ROW, COL);//�����̴�ӡ����
	//����
	while (1)
	{
		//��������壬�ٵ�������
		//�������
		PlayerMove(board, ROW, COL);//������庯��
		DisplayBoard(board, ROW, COL);//�����̴�ӡ��������ҿ���ֱ���µ���������
		//�ж�����Ƿ�Ӯ
		ret = RetBOOT(board, ROW, COL);
		//�жϺ�����������ֵ��ʲô��ֻҪ����C������break����
		if (ret !='C')
		{
			break;
		}
		//��������
		ComputerMove(board, ROW, COL);
		DisplayBoard(board, ROW, COL);//�����̴�ӡ��������ҿ���ֱ���µ���������
		//�ж�������Ƿ�Ӯ
		ret = RetBOOT(board, ROW, COL);//�жϺ�����������ֵ��ʲô��ֻҪ����C������break����
		if (ret != 'C')
		{
			break;
		}
	}
	if (ret=='*')
	{
		printf("���Ӯ\n");
	}
	else if (ret=='#')
	{
		printf("����Ӯ\n");
	}
	else
	{
		printf("ƽ��\n");
	}
}

void test()
{
	int input = 0;
	srand((unsigned int)time(NULL));//��ʱ����������������������
	do
	{
		menu();//���ò˵�����
		printf("��ѡ��>��");
		int b=scanf("%d",&input);
		switch (input)
		{
		case 1:
			game();
			printf("������\n");
			break;
		case 0:
			printf("�˳���Ϸ\n");
			break;
		default:
			printf("ѡ�����������ѡ��!\n");
			break;
		}
	}while(input);
}

void main()
{
	test();//������Ϸ
}