package ReverseLinkedList

next = cur.next
cur.next = prev
prev = cur
cur = next
run until nil pointer is not encountered
