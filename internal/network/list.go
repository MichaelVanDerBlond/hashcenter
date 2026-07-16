package network

type List struct {
	Networks []Network
}

func (l *List) Add(n Network) {
	n.ID = len(l.Networks) + 1
	l.Networks = append(l.Networks, n)
}
