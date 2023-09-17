//

int solution(string &S) {
    int len=S.length();
    int a=0,b=0,c=0,d=0;
    for (int i=0;i<len;i++){
        if (S[i]=='^') a++;
        if (S[i]=='<') b++;
        if (S[i]=='>') c++;
        if (S[i]=='v') d++;
    }
    int m=0;
    m=max(m,a);
    m=max(m,b);
    m=max(m,c);
    m=max(m,d);
    return a+b+c+d-m;
}