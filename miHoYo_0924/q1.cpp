#include<bits/stdc++.h>

using namespace std;

// 相加异或：直接暴力，求出a数组总和及b数组总和后，依次减去对应数组的每个元素，遍历以便即可
int ans=0;
int a[100010],b[100010],n;

int main(){
    scanf("%d",&n);
    for (int i=1;i<=n;i++){
        scanf("%d",&a[i]);
        a[0]+=a[i];
    }
    for (int i=1;i<=n;i++){
        scanf("%d",&b[i]);
        b[0]+=b[i];
    }
    for (int i=1;i<=n;i++){
        ans=max(ans,(a[0]-a[i])^b[0]);
        ans=max(ans,(b[0]-b[i])^a[0]);
    }
    printf("%d\n",ans);
    return 0;
}