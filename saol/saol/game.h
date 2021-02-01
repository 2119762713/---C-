#define _CRT_SECURE_NO_WARNINGS 1
#include <stdio.h>
#include <stdlib.h>
#include <time.h>
#pragma once
#define HAN 9
#define LIE 9
#define BUZL 10//这是雷的函数
#define HANS HAN+2//这里+2是扫雷外围//为什么是HAN不是直接数字9，这是为了避免如果要改内围的范围后，外围的的也可以永远大于内围的范围2，不会因为你改内围而导致外围跟不上
#define LIES LIE+2//这里+2是扫雷外围

//函数声明
void InitBoot(char board[HANS][LIES],int hans,int lies,char set);
//函数声明
void Display(char board[HANS][LIES], int han, int lie);
//函数声明
void SetMine(char board[HANS][LIES], int han, int lie);
//函数声明
void FindMine(char board[HANS][LIES], char show[HANS][LIES], int han, int lie);