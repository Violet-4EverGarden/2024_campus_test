#include<bits/stdc++.h>

using namespace std;

int n,x,a[1000001],h=1e9+7;
long long ans=0;

int main(){
    scanf("%d%d",&n,&x);
    for (int i=1;i<=n;i++){
        scanf("%d",&a[i]);
    }
    for (int i=n;i>=1;i--){
        ans=(ans*x+a[i])%h;
    }
    ans=(ans*x)%h;
    printf("%lld\n",ans);
    return 0;
}