#include<bits/stdc++.h>

using namespace std;

int n;

struct lisan{
    int num,i;
}l[200010];

struct point{
    int fa,num,i,q;
}p[200010];
int deep[200010];
struct {
    int next,i;
}bian[400010];
int tail[200010],tot;

bool cmp(const lisan &a,const lisan &b){
    return a.num<b.num;
}
long long ans=0,o=0;
int h=1e9+7;
int dfs(int nowp,int d){
    p[nowp].q=deep[p[nowp].i];
    long long t=deep[p[nowp].i];
    o=max(o,t);
    ans=(ans+d-deep[p[nowp].i])%h;

    deep[p[nowp].i]=d;
    t=o;
    for (int i=nowp;bian[i].next!=-1;i=bian[i].next){
        int nowson=bian[i].i;
        dfs(nowson,d+1);
    }
    o=t;
    deep[p[nowp].i]=p[nowp].q;
}

int main(){
    scanf("%d",&n);
    p[1].fa=-1;
    tot=n;
    for (int i=1;i<=n;i++){
        tail[i]=i;
    }
    for (int i=2;i<=n;i++){
        scanf("%d",&p[i].fa);
        bian[tail[p[i].fa]].i=i;
        tot++;
        bian[tail[p[i].fa]].next=tot;
        tail[p[i].fa]=tot;
    }
    for (int i=1;i<=n;i++){
        bian[tail[i]].next=-1;
    }
    for (int i=1;i<=n;i++){
        scanf("%d",&p[i].num);
        l[i].num=p[i].num;
        l[i].i=i;
    }
    sort(l+1,l+1+n,cmp);
    int now=1;
    for (int i=1;i<=n;i++){
        if (l[i].num!=l[now].num){
            now=i;
        }
        p[l[i].i].i=now;
    }
    dfs(1,1);
    printf("%lld",ans);
    return 0;
}