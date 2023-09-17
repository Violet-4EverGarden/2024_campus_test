#include<bits/stdc++.h>

using namespace std;

int n,a[200010];

int main(){
    scanf("%d",&n);
    for (int i=1;i<=n;i++){
        scanf("%d",&a[i]);
    }
    sort(a+1,a+n+1);
    int l=1,r=n;
    long long ans=0,now=0;
    while (l<=r){
        ans=ans+a[r]-now;
        now=a[l];
        l++;
        r--;
    }
    printf("%lld\n",ans);
    return 0;
}
