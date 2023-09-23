#include<bits/stdc++.h>

using namespace std;

int n,a;
long long dp[201][201][201];
int h=1e9+7;

long long dfs(int n,int a,int k){
    if (n<0)return 0;
    if (n==0) return 1;
    if (a*k<n) return 0;
    if (dp[n][a][k]!=0) return dp[n][a][k];
    if (a==1) return 1;
    long long ans=0;
    for (int i=0;i<=a;i++){
        ans=(ans+dfs(n-i*k,a-i,k-1))%h;
    }
    dp[n][a][k]=ans;
    return ans;
}

int main(){
    scanf("%d%d",&n,&a);
    printf("%lld",dfs(n,a,n));
    return 0;
}