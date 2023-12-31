#include<bits/stdc++.h>

using namespace std;

int n,a;

// 推公式：首先连续子区间，所以最大最小值在两侧，得最大j最小i的区间贡献等价于a[i]*a[j]*(j-i)
// 相当于计算两个求和函数：i从1到n、j从i到n，计算极差之和。
// 求和函数可通过前缀和的方式优化，实现O(n)解法

int main(){
    scanf("%d",&n);
    long long ans=0,h=1e9+7,suma=0,sumaa=0;
    for (int i=1;i<=n;i++){
        scanf("%d",&a);
        ans=(ans+sumaa*a)%h;
        suma=(suma+a)%h;
        sumaa=(sumaa+suma)%h;
    }
    printf("%lld\n",ans);
    return 0;
}