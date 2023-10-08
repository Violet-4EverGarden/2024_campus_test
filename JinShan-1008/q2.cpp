#include<bits/stdc++.h>

using namespace std;

struct aaa{
    int num,f;
}p[1000010];

int n,x,a[100010],pre[100010],next[100010],num,ans[100010],now,ind;

bool cmp(const aaa &a,const aaa &b){
    return a.num>b.num;
}

void mmm(int x){
    if (pre[x]!=0) next[pre[x]]=next[x];
    if (next[x]!=n+1) pre[next[x]]=pre[x];

}

int main(){
    scanf("%d%d",&n,&x);
    for (int i=1;i<=n;i++){
        scanf("%d",&a[i]);
        p[i].num=a[i];
        p[i].f=i;
        pre[i]=i-1;
        next[i]=i+1;
    }
    sort(p+1,p+n+1,cmp);
    now=3;
    for (int i=1;i<=n;i++){
        if (ans[p[i].f]==0){
            ind=p[i].f;
            ans[ind]=now;
            mmm(ind);
            num++;
            for (int j=pre[ind],o=1;j!=0&&o<=x;o++,j=pre[j]){
                ans[j]=now;
                mmm(j);
                num++;
            }

            for (int j=next[ind],o=1;j!=n+1&&o<=x;o++,j=next[j]){
                ans[j]=now;
                mmm(j);
                num++;
            }
            now=now^1;
        }
    }
    for (int i=1;i<=n;i++){
        if (ans[i]==3) printf("A");
        else printf("B");
    }
    return 0;
}