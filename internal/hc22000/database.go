package hc22000

type Database struct {
	Networks []Network
}

func (d *Database) Add(n Network) {
	d.Networks = append(d.Networks, n)
}
