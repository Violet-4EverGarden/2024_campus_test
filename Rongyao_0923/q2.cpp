#include<bits/stdc++.h>

using namespace std;

int n,k,a,b;
int c[100];

// 注意：结果需要按字典序排列
int dfs(int x){
    if (x==0){
        for (int i=n+k;i>=1;i--){
            printf("%c",c[i]);
        }
        printf("\n");
        return 0;
    }
    if (a<n){
        a++;
        c[x]='*';
        dfs(x-1);
        a--;
    }
    if (b<k){
        b++;
        c[x]='|';
        dfs(x-1);
        b--;
    }
    return 0;
}

int main(){
    scanf("%d%d",&n,&k);
    k--;
    long long ans=1;
    for(int i=n+k;i>n;i--){
        ans=ans*i;
    }
    for (int i=1;i<=k;i++){
        ans=ans/i;
    }
    printf("%lld\n",ans);
    a=0,b=0;
    dfs(n+k);
    return 0;
}