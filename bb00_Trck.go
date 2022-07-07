package tmcoll
import "fmt"
import "github.com/qb-qetell/combGUID"

type Trck struct {
	Iddd string
	Name string
	Clap chan *Mssg
	Flap chan *Mssg
	Code func (hostIddd, iddd, name string, clap <- chan *Mssg, flap chan <- *Mssg,
		seed map[string]interface{})
	Seed map[string]interface{}
}
	func Trck_Estb (
			
		iddd string,
		name string,
		code func (hostIddd, iddd, name string, clap <- chan *Mssg, flap chan <- *Mssg,
			seed map[string]interface{}),
		seed ... map[string]interface{},
		) (
		*Trck,
		) {
		
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
		defer func (objc *Trck) {
			_ba00 := recover ()
			if _ba00 != nil {
				_ca00 := fmt.Sprintf (`Paniced. [%v]`, _ba00)
				_cb00 := []string {
					combGUID.CombGUID_Estb ("", 16).SmplFrmt (),
					"bm00",
					_ca00,
				}
				Mssg_Estb (objc.Iddd, hostIddd, _cb00).Send (objc.Flap)
			}
		} (objc)
		objc.Code (hostIddd, objc.Iddd, objc.Name, objc.Clap, objc.Flap, objc.Seed)
	}
