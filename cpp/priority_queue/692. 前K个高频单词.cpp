class Solution {
public:
    vector<string> topKFrequent(vector<string>& words, int k) {
        Cmp cmp;
        unordered_map<string,int> count;
        for(auto word:words){
            count[word]++;
        }

    vector<wordCount> wordCounts;
    for(auto it : count) {
        wordCounts.push_back(wordCount(it.first,it.second));
    }
    priority_queue<wordCount,vector<wordCount>,Cmp> queue;
    for(auto wc : wordCounts){
        if (queue.size()<k){
            queue.push(wc);
            continue;
        }
        if (cmp(wc,queue.top())){
            queue.pop();
            queue.push(wc);
        }
    }
    vector<string> result;
    while(!queue.empty()){
        wordCount wc = queue.top();
        queue.pop();
        result.push_back(wc.word);
    }
    reverse(result.begin(),result.end());
    return result;

}
struct wordCount{
    string word;
    int count; 

    wordCount(string word,int count) {
        this->word = word;
        this->count = count;
    }
};
struct Cmp {
    bool operator() (const wordCount& w1, const wordCount& w2) const{
        if (w1.count==w2.count){
            return w1.word<w2.word;
        }
        return w1.count>w2.count;
    }
};

/*
	知识点
 		1. 优先队列自定义比较。
		2. Map 的使用。
		3. For 迭代的使用。
*/


