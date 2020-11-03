class Solution {
public:
    vector<string> result;
    void swapChar(string& str,int i,int j){
        char tmp = str[i];
        str[i]=str[j];
        str[j]=tmp;
    }
    
    void formResult(string str,string curString) {
        if (str == ""){
            if (curString != ""){
                this->result.push_back(curString);
            }
            return;
        }
        map<char,bool> hasBeenHandled;
        for (int i=0;i<str.length();i++){
            if (hasBeenHandled[str[i]]){
                continue;
            }
            hasBeenHandled[str[i]]=true;
            swapChar(str,0,i);
            formResult(str.substr(1),curString+str[0]);
            swapChar(str,0,i);
        }
    }
    
    vector<string> Permutation(string str) {
        formResult(str,"");
        sort(this->result.begin(),this->result.end());
        return this->result;
    }
};

/*
    总结:
        1. 采用递归思维就能 AC 了～
        2. 要注意不要重复处理。
    链接: https://www.nowcoder.com/practice/fe6b651b66ae47d7acce78ffdd9a96c7?tpId=188&&tqId=36198&rp=1&ru=/activity/oj&qru=/ta/job-code-high-week/question-ranking
*/