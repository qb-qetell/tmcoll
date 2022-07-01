package tmcoll
import "time"

type Mssg struct {
	Sndr string
	Rcpn string
	Core interface {}
}
	func Mssg_Estb (sndr, rcpn string, core interface {}) (*Mssg) {
		mssg := &Mssg {}
		mssg.Sndr = sndr
		mssg.Rcpn = rcpn
		mssg.Core = core
		return mssg
	}
	func (objc *Mssg) Send (mssgBoxx chan <- *Mssg, waitDrtn ... time.Duration) (bool) {
		if len (waitDrtn) != 0 {
			if waitDrtn [0] == (time.Nanosecond * 0) {
				select {
				case mssgBoxx <- objc: { return true }
				default: { return false }
				}
			} else {
				flapXX := make (chan bool)
				go func (slppDrtn time.Duration, flap chan <- bool) {
					time.Sleep (slppDrtn)
					select {
					case flap <- true: {}
					default: {}
					}
				} (waitDrtn [0], flapXX)
				select {
				case mssgBoxx <- objc: { return true  }
				case _  =  <- flapXX:  { return false }
				}
			}
		} else {
			mssgBoxx <- objc
			return true
		}
	}
