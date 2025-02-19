/**
 * Definition for singly-linked list.
 * class ListNode {
 *     val: number
 *     next: ListNode | null
 *     constructor(val?: number, next?: ListNode | null) {
 *         this.val = (val===undefined ? 0 : val)
 *         this.next = (next===undefined ? null : next)
 *     }
 * }
 */

function mergeTwoLists(list1: ListNode | null, list2: ListNode | null): ListNode | null {
    const dummy = new ListNode(0);
    let cur = dummy;

    let next = getSmallerNode(list1, list2);
    while (next !== null) {
        cur.next = next;
        cur = cur.next;

        [list1, list2] = advanceLists(list1, list2, next);
        next = getSmallerNode(list1, list2);
    }

    return dummy.next;
}


function getSmallerNode(node1: ListNode | null, node2: ListNode | null): ListNode | null {
    if (node1 === null) {
        return node2;
    }
    if (node2 === null) {
        return node1;
    }
    return node1.val < node2.val ? node1 : node2;
}


function advanceLists(node1: ListNode | null, node2: ListNode | null, chosen: ListNode): [ListNode | null, ListNode | null] {
    if (chosen === node1) {
        return [node1?.next || null, node2];
    }
    return [node1, node2?.next || null];
}