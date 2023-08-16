class Solution {
public:
    ListNode* rotateRight(ListNode* head, int k) {
        if(k == 0) {
            return head;
        }
        if(head == nullptr) {
            return head;
        }

        int size = 1;
        ListNode* tail = head;
        while(true) {
            if(tail->next == nullptr) {
                break;
            }
            tail = tail->next;
            size++;
        }
        tail->next = head;

        k = k % size;

        ListNode* previousHead = nullptr;
        ListNode* newHead = head;

        for(int i = 0; i < size - k; i++) {
            previousHead = newHead;
            newHead = newHead->next;
        }
        
        previousHead->next = nullptr;
        return newHead;
    }
};
