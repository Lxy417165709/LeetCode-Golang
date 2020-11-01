# include<bits\stdc++.h>

using namespace std; 

class Solution {
public:
    int findKthLargest(vector<int>& nums, int k) {
        priority_queue<int,vector<int>,greater<int> > queue;
        for(int i=0;i<nums.size();i++){
            if (queue.size()<k){
                queue.push(nums[i]);
                continue;
            }
            if (nums[i]>=queue.top()){
                queue.pop();
                queue.push(nums[i]);
            }
        }
        return queue.top(); 
    }
};

/*
	֪ʶ��
	1. �½���С��: priority_queue<int,vector<int>,greater<int> > queue;  
	2. ����ȡͷԪ�ء��Ƴ�ͷԪ�ء�����Ԫ��: .top()��.pop()��push() 
	
	
	�ܽ�
	1. �����������С����������ʹ�����ѡ���㷨�� 
*/
