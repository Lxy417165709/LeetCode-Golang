class Solution {
public:
    vector<int> intersection(vector<int>& nums1, vector<int>& nums2) {
        set<int> set1,set2,intersectionSet;
        for (auto num : nums1){
            set1.insert(num);
        }
        for (auto num : nums2){
            set2.insert(num);
        }
        set_intersection(set1.begin(),set1.end(),set2.begin(),set2.end(),inserter(intersectionSet, intersectionSet.begin()));
        vector<int> result;
        for (auto num : intersectionSet){
            result.push_back(num);
        }
        return result;
    }
};