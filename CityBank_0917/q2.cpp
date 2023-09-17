#include <vector>

int solution(int X, int Y, vector<int> &A) {
    int N = A.size();
    int result = -1;
    int nX = 0;
    int nY = 0;
    for (int i = 0; i < N; i++) {
        if (A[i] == X)
            nX += 1;
        // else if (A[i] == Y)  原代码
        if (A[i] == Y)
            nY += 1;
        if (nX == nY)
            result = i;
    }
    return result;
}
