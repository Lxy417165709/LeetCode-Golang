class Solution {
public:
    const int undefined = -1000000000;
    static bool cmp(vector<int> a,vector<int> b) {
        return a[0]<b[0];
    }
    vector<int> getIntervalVector(int start,int stop) {
        vector<int> element;
        element.push_back(start);
        element.push_back(stop);
        return element;
    }
    vector<vector<int>> merge(vector<vector<int>>& intervals) {
        vector<vector<int>> result;
        if (intervals.size()==0){
            return result;
        }
        sort(intervals.begin(),intervals.end(),cmp);
        int curStart = undefined;
        int curStop = undefined;
        for(auto interval : intervals){
            if (curStart==undefined && curStop==undefined){
                curStart = interval[0];
                curStop = interval[1];
                continue;
            }
            if (curStop>=interval[0]) {
                curStop = max(curStop,interval[1]);
            }else{
                result.push_back(getIntervalVector(curStart,curStop));
                curStart = interval[0];
                curStop = interval[1];
            }
        }
       if (curStart!=undefined && curStop!=undefined){
            result.push_back(getIntervalVector(curStart,curStop));
        }
        return result;
    }
};

/*
    知识点
    1. vector 自定义排序。
*/