// Forward declaration of isBadVersion API.
bool isBadVersion(int version);

class Solution {
public:
    int firstBadVersion(int n) {
        int l=1,r=n;
        while(l<=r){
            int mid = l+(r-l)/2;
            if (isBadVersion(mid)){
                r = mid - 1;
            }else{
                l = mid + 1;
            }
        }
        return l;
    }
};

/*
    总结
    1. 这题目要找的是第一个错误的版本，此时在循环外应该返回l。
        如果要找最后一个正确的版本，那么应该返回r.
*/