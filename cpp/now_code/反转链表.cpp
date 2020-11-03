/*
struct ListNode {
	int val;
	struct ListNode *next;
	ListNode(int x) :
			val(x), next(NULL) {
	}
};*/
class Solution {
public:
    ListNode* ReverseList(ListNode* pHead) {
        ListNode* pre = NULL;
        ListNode* next = NULL;
        while(pHead!=NULL){
            next = pHead->next;
            pHead->next = pre;
            pre = pHead;
            pHead = next;
        }
        return pre;
    }
};

/*
    题目链接: https://www.nowcoder.com/practice/75e878df47f24fdc9dc3e400ec6058ca?tpId=188&&tqId=36164&rp=1&ru=/activity/oj&qru=/ta/job-code-high-week/question-ranking
    总结:
        1. 反转链表很简单～按着步骤走就可以了。
*/