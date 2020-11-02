/**
 * Definition for a binary tree node.
 * struct TreeNode {
 *     int val;
 *     TreeNode *left;
 *     TreeNode *right;
 *     TreeNode(int x) : val(x), left(NULL), right(NULL) {}
 * };
 */
class Solution {
public:
    vector<vector<int>> levelOrderBottom(TreeNode* root) {
        queue<TreeNode*> queue;
        vector<vector<int>> layerVectors;
        queue.push(root);
        while(!queue.empty()){ 
            int size = queue.size();
            vector<int> layerVector;
            while(size--){
                TreeNode* top = queue.front();
                queue.pop();
                if (top == NULL){
                    continue;
                }
                layerVector.push_back(top->val);
                queue.push(top->left);
                queue.push(top->right);
            }
            if(layerVector.size()!=0){
                layerVectors.push_back(layerVector);
            }
        }
        reverse(layerVectors.begin(),layerVectors.end());
        return layerVectors;
    }
};

/*
    知识点
    1. 非空指针的判定。
    2. queue 的使用。
    3. vector 的使用。
*/