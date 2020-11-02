class Solution {
public:
    string removeDuplicateLetters(string s) {
        stack<char> stack;
        map<char,int> count;
        map<char,bool> hasInStack;
        for (auto ch:s){
            count[ch]++;
        }
        for (auto ch:s){
            if (hasInStack[ch]){
                count[ch]--;
                continue;
            }
            if (stack.empty()){
                stack.push(ch);
                hasInStack[ch]=true;
                count[ch]--;
                continue;
            }
            while (!stack.empty() && stack.top()>ch && count[stack.top()]!=0){
                hasInStack[stack.top()]=false;
                stack.pop();
            }
            stack.push(ch);
            hasInStack[ch]=true;
            count[ch]--;
        }
        string result = "";
        while(!stack.empty()){
            result=stack.top() + result;
            stack.pop();
        }
        return result;
    }
};

/*
    知识点
    1. stack 的使用

    总结
    1. 这题需要分情况讨论。 大神链接: https://leetcode-cn.com/problems/remove-duplicate-letters/solution/you-qian-ru-shen-dan-diao-zhan-si-lu-qu-chu-zhong-/
    2. 代码有待整理。
*/