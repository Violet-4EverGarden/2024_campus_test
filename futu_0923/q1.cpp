#include<bits/stdc++.h>

using namespace std;

int ttt,n,a,m;

int main(){
    scanf("%d",&ttt);
    for (int tt=1;tt<=ttt;tt++){
        scanf("%d%d",&n,&m);
        int t= n/m+(n%m!=0);
        printf("%d\n%d",t,n-(t-1)*m);
        for (int i=2;i<=t;i++){
            printf(" %d",m);
        }
        printf("\n");
    }
    return 0;
}