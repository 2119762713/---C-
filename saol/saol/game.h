#define _CRT_SECURE_NO_WARNINGS 1
#include <stdio.h>
#include <stdlib.h>
#include <time.h>
#pragma once
#define HAN 9
#define LIE 9
#define BUZL 10//�����׵ĺ���
#define HANS HAN+2//����+2��ɨ����Χ//Ϊʲô��HAN����ֱ������9������Ϊ�˱������Ҫ����Χ�ķ�Χ����Χ�ĵ�Ҳ������Զ������Χ�ķ�Χ2��������Ϊ�����Χ��������Χ������
#define LIES LIE+2//����+2��ɨ����Χ

//��������
void InitBoot(char board[HANS][LIES],int hans,int lies,char set);
//��������
void Display(char board[HANS][LIES], int han, int lie);
//��������
void SetMine(char board[HANS][LIES], int han, int lie);
//��������
void FindMine(char board[HANS][LIES], char show[HANS][LIES], int han, int lie);