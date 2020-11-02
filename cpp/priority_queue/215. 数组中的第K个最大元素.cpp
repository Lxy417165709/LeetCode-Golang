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
	知识点
	1. 新建最小堆: priority_queue<int,vector<int>,greater<int> > queue;  
	2. 队列取头元素、移出头元素、推入元素: .top()、.pop()、push() 
	
	
	总结
	1. 这题除了用最小堆做，还能使用随机选择算法。 
*/
