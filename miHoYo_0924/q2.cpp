#include<bits/stdc++.h>

using namespace std;

int n,h,t,k,sum;
double dp[7000],dpt[7000];

int main(){
    scanf("%d%d",&n,&h);
    int tsum=0;
    for (int i=1;i<=n;i++){
        scanf("%d%d",&t,&k);
        if (t==1) tsum+=k;
        if (t==2) {
            h-=k;
            sum+=tsum;
            tsum=0;
        }
    }
    dp[0]=1;
    double ans=0;
    if (h<=6*sum)
        for (int i=1;i<=sum;i++){
            double s=0;
            for (int j=1;j<=6;j++){
                if (h-j>=0) s=s+dp[h-j];
                ans+=s/6;
            }
            for (int j=1;j<=min(h-1,6*sum);j++){
                s = 0;
                for (int k=1;k<=6;k++){
                    if (j-k<0) continue;
                    s=s+dp[j-k];
                }
                dpt[j]=s/6;
            }
            for (int j=0;j<=min(h-1,6*sum);j++){
                dp[j]=dpt[j];
            }
        }
    if (h<=0) ans=1;
    printf("%lf",ans);
    return 0;
}