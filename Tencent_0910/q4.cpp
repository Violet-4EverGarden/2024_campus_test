#include<bits/stdc++.h>

using namespace std;

int ttt,n,k;
long long ans,h=1e9+7;
long long _2[64];

int main(){
    scanf("%d",&ttt);
    for (int tt=1;tt<=ttt;tt++){
        scanf("%d%d",&n,&k);
        _2[0]=1;
        for (int i=1;i<=63;i++){
            _2[i]=_2[i-1]*2%h;
        }
        int t=n-k;
        ans=1;
        int o=0;
        while (t>0){
            o++;
            if (t%2==1){
                ans=ans*_2[o]%h;
            }
            t=t/2;
        }
        printf("%d\n",(ans+h-1)%h);
    }
    return 0;
}