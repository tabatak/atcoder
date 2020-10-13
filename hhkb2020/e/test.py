mod=10**9+7
h,w=map(int,input().split())
s=[list(input()) for i in range(h)]
k=0
for i in range(h):
    for j in range(w):
        if s[i][j]==".":
            k+=1
t=[["" for j in range(h)] for i in range(w)]
for i in range(h):
    for j in range(w):
        t[j][i]=s[i][j]
po=[0]*(k+1)
po[0]=1
for i in range(1,k+1):
    po[i]=po[i-1]*2
    po[i]%=mod
#行でつながっている部分
from itertools import groupby
check=[[0]*w for i in range(h)]
for i in range(h):
    now=0
    for key,group in groupby(s[i]):
        l=len(list(group))
        if key=="#":
            now+=l
        else:
            for j in range(now,now+l):
                check[i][j]+=l
            now+=l
for i in range(w):
    now=0
    for key,group in groupby(t[i]):
        l=len(list(group))
        if key=="#":
            now+=l
        else:
            for j in range(now,now+l):
                check[j][i]+=l
            now+=l
ans=0
for i in range(h):
    for j in range(w):
        #print(k,k-check[i][j])
        if check[i][j]!=0:
            ans+=(po[k]-po[k-check[i][j]+1])
        ans%=mod
print(ans)
