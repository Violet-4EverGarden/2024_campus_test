#include<bits/stdc++.h>

using namespace std;

int ttt, n;
struct numnum {
    int num, index;
} a[100010];

int pre[100010], next[100010], from[100010], b;

bool cmp(const numnum& x, const numnum& y) {
    return x.num < y.num;
}

void pr(int x) {
    printf("%d%s", x / 2, (x % 2 == 1) ? ".5" : "");
}

int main() {
    scanf("%d", &ttt);
    for (int tt = 1; tt <= ttt; tt++) {
        scanf("%d", &n);
        for (int i = 1; i <= n; i++) {
            scanf("%d", &a[i].num);
            a[i].index = i;
        }
        sort(a + 1, a + n + 1, cmp);
        int now = n / 2 + n % 2;
        for (int i = 1; i <= n; i++) {
            from[a[i].index] = i;
            pre[i] = i - 1;
            next[i] = i + 1;
        }
        pr(a[now].num + a[now + (n + 1) % 2].num);
        printf(" ");
        int tn = n;
        for (int i = 1; i <= n - 1; i++) {
            scanf("%d", &b);
            int d = from[b];
            if (d == now) {
                if (tn % 2 == 0) {
                    now = next[now];
                } else {
                    now = pre[now];
                }
            } else {
                if (d < now) {
                    if (tn % 2 == 0) {
                        now = next[now];
                    }
                }
                if (d > now) {
                    if (tn % 2 == 1) {
                        now = pre[now];
                    }
                }
            }
            next[pre[d]] = next[d];
            pre[next[d]] = pre[d];
            int ans2 = a[now].num;
            tn--;
            if (tn % 2 == 1) {
                ans2 += a[now].num;
            } else {
                ans2 += a[next[now]].num;
            }
            printf("%d%s", ans2 / 2, (ans2 % 2 == 1) ? ".5" : "");
            if (i != n - 1) {
                printf(" ");
            } else {
                printf("\n");
            }
        }
    }
    return 0;
}
/*
2
5
2 2 1 3 5
3 1 2 5
4
1 1 1 1
1 2 3
*/