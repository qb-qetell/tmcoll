package tmcoll

type Mssg struct {
	Sndr string
	Rcpn string
	Core []byte
}
	func Estb_Mssg (sndr, rcpn string, core []byte) (*Mssg) {
		mssg := &Mssg {}
		mssg.Sndr = sndr
		mssg.Rcpn = rcpn
		mssg.Core = core
		return mssg
	}
	func (objc *Mssg) Send (mssgBoxx chan<- *Mssg, wait ... bool) (bool) {
	// wait time
		if len (wait) == 0 {
			select {
				case mssgBoxx <- objc: { return true }
				default: { return false }
			}
		} else {
			mssgBoxx <- objc
			return true
		}
	}
