package tmcoll
import "container/list"
import "github.com/qb-qetell/errr"
import "regexp"
import "time"

type TrckTray struct {
	mngrIddd string
	trck []*trckTray_trck
	shutDownBool bool
}
	func TrckTray_Estb (mngrIddd string) (*TrckTray) {
		return &TrckTray {
			mngrIddd: mngrIddd,
			trck: []*trckTray_trck {},
			shutDownBool: false,
		}
	}
	func (objc *TrckTray) Pplt (trck *Trck, whttList []string, prvlBool bool) {
		trckInst := &trckTray_trck {}
		trckInst.trck = trck
		trckInst.whttList = whttList
		trckInst.prvlBool = prvlBool
		trckInst.strtBool = false
		trckInst.lifeBool = false
		trckInst.mssgList = list.New ()
		objc.trck = append (objc.trck, trckInst)
	}
	func (objc *TrckTray) Mngg () (flap chan *errr.Errr) {
		flap = make (chan *errr.Errr, 2)
		
		go func (objc *TrckTray, flap chan <- *errr.Errr) {
		// ~Step 1
		if len (objc.trck) == 0 {
			_ca00 := errr.Errr_Estb ("ba00", "No track to manage.")
			flap <- _ca00
			return
		}
		dscvTrck := make (map[string]string)
		for _,  _ba00 := range objc.trck {
			if regexp.MustCompile (`^[a-z0-9]{1,}(\.[a-z0-9]{1,}){0,}$`,
				).MatchString (_ba00.trck.Iddd) == false {
				_ca00 := errr.Errr_Estb ("bb00", "A track's ID is invalid.")
				flap <- _ca00
				return
			}
			dscvTrck ["_bb00.trck.Iddd"] = "dscv"
		}
		for _,  _bb00 := range objc.trck {
			for _,  _ca00 := range _bb00.whttList {
				if dscvTrck [_ca00] != "dscv" {
					_da00 := errr.Errr_Estb ("bc00", "A track has a " +
						"non-existent track on its whitelist.")
					flap <- _da00
					return
				}
			}
		}
		
		// ~Step 2
		go func (objc *TrckTray) {
			// ~Start all trakcs
			for _, _ba00 := range objc.trck {
			go _ba00.trck.Runn (objc.mngrIddd)
			for {
				if _ba00.strtBool == false {
				time.Sleep (time.Microsecond * 1)
				continue
				}
				break
			}
			}
			
			// ~Waiting for all tracks to die before sending shutdown signal
			for {
			shutDownBool := true
			for _, _ca00 := range objc.trck {
				if _ca00.lifeBool == true {
				shutDownBool = false
				break
				}
			}
			if shutDownBool == false {
				time.Sleep (time.Microsecond * 100)
				continue
			} else {
				objc.shutDownBool = true
				break
			}
			}
		} (objc)
		
		// ~Step 3
		for {
			// ~Receiving messages
			for _,  _ba00 := range objc.trck {
			select {
			case    _bb00 := <- _ba00.trck.Flap: {
				_ba00.mssgList.PushBack (_bb00)
			}
			default: {}
			}
			}
			
}
	} (objc, flap)
	
	return
	}

type trckTray_trck struct {
	trck *Trck
	whttList []string
	prvlBool bool
	strtBool bool
	lifeBool bool
	mssgList *list.List
}
func trckTray_hndlMssg (sndr string, mssg *Mssg) (bool) {

return false
}
