class Solution {
public:
    bool isValid(string s) {
        stack<char> stack;
        map<char,char> leftSymbol;
        leftSymbol[')']='(';
        leftSymbol['}']='{';
        leftSymbol[']']='[';
        for (auto ch : s) {
            if (ch =='(' || ch=='['||ch=='{'){
                stack.push(ch);
            }else{
                if (stack.empty() || leftSymbol[ch] != stack.top()){
                    return false;
                }
                stack.pop();
            }
        }
        return stack.empty();
    }
};

/*
    知识点:
    1. stack 的使用。
*/