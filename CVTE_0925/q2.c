// 题目要求C语言实现字符串反转
#include<stdio.h>
#include<string.h>

void reverse(char* s){
  int n = strlen(s);  // 使用strlen获取s的长度（遇到结束符\0截止）
  if (n>1) {
  	char c = *s;
    *s = *(s + n - 1);  // 将s的第一个字符替换为最后一个字符，并将最后一个字符设为字符串结束符
    *(s + n - 1) = '\0';
    reverse(s+1);  // 递归地对去掉第一个字符的子字符串进行反转操作
    *(s + n - 1) = c;  // 将原来存储在c的第一个字符放回到原字符串的最后一个位置
  }
}