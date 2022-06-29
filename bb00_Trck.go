package tmcoll

type Trck struct {
	iddd string
	name string
	clap chan *Mssg
	flap chan *Mssg
	code func (iddd, name string, clap <- chan *Mssg, flap chan <- *Mssg,
		seed map[string]interface{})
	seed map[string]interface{}
}
	func Trck_Estb (iddd, name string, code func (iddd, name string, clap <- chan *Mssg,
		flap chan <- *Mssg, seed map[string]interface{}), seed ...
		map[string]interface{}) (*Trck, chan<- *Mssg) {
		trck := &Trck {}
		trck.iddd = iddd
		trck.name = name
		trck.clap = make (chan *Mssg)
		trck.flap = make (chan *Mssg)
		trck.code = code
		trck.seed = make (map[string]interface{})
		if len (seed) != 0 && seed [0] != nil {
			trck.seed = seed [0]
		}
		return trck, trck.clap
	}
	func (objc *Trck) Runn () {
		objc.code (objc.iddd, objc.name, objc.clap, objc.flap, objc.seed)
	}
