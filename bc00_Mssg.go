package tmcoll
import "github.com/qb-qetell/errr"
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
	func (objc *Mssg) Send (mssgBoxx chan <- *Mssg, waitDrtn ... time.Duration) (error) {
		if len (waitDrtn) != 0 {
			if waitDrtn [0] == (time.Nanosecond * 0) {
				select {
				case mssgBoxx <- objc: { return nil }
				default: {
					return errr.Errr_Estb (
						"ba00",
						"Could not send message. [Channel did not " +
						"allow message in.]",
					)
				}
				}
			} else {
				flap := make (chan bool)
				go func (slppDrtn time.Duration, flap chan <- bool) {
					time.Sleep (slppDrtn)
					select {
					case flap <- true: {}
					default: {}
					}
				} (waitDrtn [0], flap)
				select {
				case mssgBoxx <- objc: { return nil  }
				case _    =   <- flap: {
					return errr.Errr_Estb (
						"ba00",
						"Could not send message. [Channel did not " +
						"allow message in.]",
					)
				}
				}
			}
		} else {
			mssgBoxx <- objc
			return nil
		}
	}
