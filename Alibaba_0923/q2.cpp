#include<bits/stdc++.h>

using namespace std;

char s[100010];

int main(){
    scanf("%s",s);
    int pre = 1,now=0,ans=0;
    for (int i=0;i<strlen(s);i++){
        if (i==0){
            now++;
        } else {
            if (s[i]==s[i-1]) now++;
            else{
                pre=now;
                now=1;
            }
        }
        if (now+pre>=4){
            now--;
            ans++;
        }
    }
    printf("%d",ans);
    return 0;
}