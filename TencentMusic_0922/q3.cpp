#include<bits/stdc++.h>

using namespace std;

char s[4000];
long long dp[4000];
int h=1e9+7,f[4000];

int main(){
    scanf("%s",s);
    dp[0]=1;
    int len=strlen(s);
    for (int i=1;i<len;i++){
        dp[i]=dp[i-1];
        f[i]=i+1;
        int now=i+1,pr=1;
        for (int j=i-1;j>=0;j--){
            while (now!=i+1){
                if (s[j]==s[now-1])break;
                now=f[i];
            }
            if (s[j]==s[now-1]){
                f[j]=now-1;
            } else {
                f[j]=i+1;
            }
            now=f[j];
            //printf("%d %d\n",j,now);
            if (i-j+1<(i-now+1)*2){
                now=f[now];
                //printf("--%d %d\n",j,now);
            }
            if (i-now+1>pr){
                pr=i-now+1;
                dp[i]=(dp[i]+dp[now-1])%h;
            }
        }
    }
    printf("%lld\n",dp[len-1]);
    return 0;
}