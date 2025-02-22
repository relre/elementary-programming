package piscine

type NodeL struct {
	Data interface{}
	Next *NodeL
}

type List struct {
	Head *NodeL
	Tail *NodeL
}

// KOD BAŞLANGICI
func ListRemoveIf(l *List, data_ref interface{}) {
	temp := l.Head
	prev := l.Head

	for temp != nil && temp.Data == data_ref {
		l.Head = temp.Next
		temp = l.Head
	}
	for temp != nil {
		if temp.Data != data_ref {
			prev = temp
		}
		prev.Next = temp.Next
		temp = prev.Next
	}
}
