package tmcoll

type Trck struct {
	Iddd string
	Name string
	clap chan *Mssg
	Flap chan *Mssg
}
	func Estb_Trck (iddd, name string) (*Trck, chan<- *Mssg) {
		trck := &Trck {}
		trck.Iddd = iddd
		trck.Name = name
		trck.clap = make (chan *Mssg)
		trck.Flap = make (chan *Mssg)
		return trck, trck.clap
	}
