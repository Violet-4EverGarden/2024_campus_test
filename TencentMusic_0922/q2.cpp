#include<bits/stdc++.h>

using namespace std;


struct TreeNode {
 	int val;
 	struct TreeNode *left;
 	struct TreeNode *right;
 	//TreeNode(int x) : val(x), left(nullptr), right(nullptr){}
};
class Solution {
public:
    int ret[200010][2];
    long long qkpow(long long a, long long p, long long mod){
        long long t=1, tt=a%mod;
        while(p){
            if(p&1)t=t*tt%mod;
            tt=tt*tt%mod;
            p>>=1;
            return t;
        }
    }
    long long getinv(long long a, long long mod){
        return qkpow(a, mod-2, mod);
    }

    void dfs(TreeNode *trees,int deep){
        int a=-1,b=-1,dl,dr;
        int *k1,*k2;
        if (trees->left){
            dfs(trees->left,deep+1);
            k1=ret[deep+1];
            a=k1[0];
            dl=k1[1];
        } else{
            a=1;
            dl=deep;
        }
        if (trees->right){
            dfs(trees->right,deep+1);
            k2=ret[deep+1];
            b=k2[0];
            dr=k2[1];
        }else{
            b=1;
            dr=deep;
        }
        int retx=0,retd=0;
        if (dl==dr){
            retd=dl;
            retx=a+b;
        } else{
            if (dl>dr){
                retd=dl;
                retx=a;
            }else{
                retd=dr;
                retx=b;
            }
        }
        ret[deep][0]=retx;
        ret[deep][1]=retd;
    }

    int cntOfMethods(vector<TreeNode*>& trees) {
        long long a=1,b=0;
        int h=1e9+7;
        for (int i=0;i<trees.size();i++){
            dfs(trees[i],0);
            int *k =ret[1];
            a=a*k[0]%h;
            b=(b+getinv(k[0],h))%h;
        }
        return a*b%h;
    }
};

int main(){


    return 0;
}