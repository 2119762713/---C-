#include "game.h"


void menu()//�˵�����
{
	printf("***********************\n");
	printf("**** 1.play 0.exit ****\n");
	printf("***********************\n");
}

void game()
{
	printf("ɨ����Ϸ\n");
	//11*11
	char board[HANS][LIES] = {0};//�׵���Ϣ�洢//����ų������������Ժ�Ϳ��Ըĳ����Ϳ����ˣ�û��Ҫȥ��ִ���㷨������к���
	char show[HANS][LIES] = {0};//�Ų��׵���Ϣ�洢
	InitBoot(board, HANS, LIES,'0');//��ʼ������
	InitBoot(show, HANS, LIES,'*');//��ʼ������
	//��ӡ����
	//Display(board, HAN, LIE);
	Display(show, HAN, LIE);
	//������--���
	SetMine(board, HAN, LIE);
	//Display(board, HAN, LIE);////�׵����̲����ӡ��Ҫ��Ȼֱ�ӱ����ˣ���ֻ�Ǹ����Լ�����
	//ɨ��---����ӡ����
	FindMine(board, show, HAN, LIE);
}

void test()
{
	int input = 0;
	srand((unsigned int)time(NULL));//��ʱ������ص�����ǿ��ת����int��//�����rand����Ҫ�������������srand
	do//��ִ��һ����Ϸ
	{
		menu();//�˵�����
		printf("��ѡ��>");
		int b = scanf("%d", &input);
		switch (input)
		{
		case 1:
			game();//ʵ����Ϸ�ĺ���
			break;
		case 0:
			printf("���˳���Ϸ\n");
			break;
		default:
			printf("�����������������\n");
			break;
		}
	} while (input);
}

void main()
{
	test();//��Ϸ����ʵ�ֵĵط�
}