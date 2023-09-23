#include<bits/stdc++.h>

using namespace std;

int n,a[10],k;

int main(){
    scanf("%d",&n);
    for (int i=1;i<=n;i++){
        scanf("%d",&k);
        a[k]++;
    }
    int a3=min(a[3],a[6]);
    a[6]=a[6]-a3;
    int a4=min(a[2],a[4]);
    a[2]=a[2]-a4;
    int a2=min(a[2],a[6]);
    printf("%d\n",min(a[1],a2+a3+a4));
    return 0;
}