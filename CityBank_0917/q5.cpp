// you can use includes, for example:
// #include <algorithm>

// you can write to stdout for debugging purposes, e.g.
// cout << "this is a debug message" << endl;

int solution(vector<int> &A, vector<int> &B, int N) {
    int m=A.size();
    int x[110]={0};
    for (int i=0;i<m;i++){
        x[A[i]]++;
        x[B[i]]++;
    }
    int ans=0;
    for (int i=0;i<m;i++){
        ans=max(x[A[i]]+x[B[i]]-1,ans);
    }
    return ans;
}

/*
([1, 2, 3, 3], [2, 3, 1, 4], 4)
*/