package tmcoll

type Trck struct {
	Iddd string
	Name string
	Clap chan *Mssg
	Flap chan *Mssg
	Code func (hostIddd, iddd, name string, clap <- chan *Mssg, flap chan <- *Mssg,
		seed map[string]interface{})
	Seed map[string]interface{}
}
	func Trck_Estb (iddd, name string, code func (hostIddd, iddd, name string, clap <- chan
		*Mssg, flap chan <- *Mssg, seed map[string]interface{}), seed ...
		map[string]interface{}) (*Trck) {
		trck := &Trck {}
		trck.Iddd = iddd
		trck.Name = name
		trck.Clap = make (chan *Mssg)
		trck.Flap = make (chan *Mssg)
		trck.Code = code
		trck.Seed = make (map[string]interface{})
		if len (seed) != 0 && seed [0] != nil {
			trck.Seed = seed [0]
		}
		return trck
	}
	func (objc *Trck) Runn (hostIddd string) {
		defer recover ()
		objc.Code (hostIddd, objc.Iddd, objc.Name, objc.Clap, objc.Flap, objc.Seed)
	}
