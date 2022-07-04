package tmcoll
import "container/list"
import "github.com/qb-qetell/errr"
import "regexp"
import "strings"
import "time"

type TrckTray struct {
	mngrIddd string
	trck []*trckTray_trck
	mssgList *list.List
	shutDownBool bool
}
	func TrckTray_Estb (mngrIddd string) (*TrckTray) {
		return &TrckTray {
			mngrIddd: mngrIddd,
			trck: []*trckTray_trck {},
			mssgList: list.New (),
			shutDownBool: false,
		}
	}
	func (objc *TrckTray) Pplt (trck *Trck, whttList []string, prvlBool bool) {
		trckInst := &trckTray_trck {}
		trckInst.trck = trck
		trckInst.whttList = whttList
		trckInst.prvlBool = prvlBool
		trckInst.strtUpppBool = false
		trckInst.strtUpppSccsBool = "undf"
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
				if _ba00.strtUpppBool == false {
				time.Sleep (time.Microsecond * 1)
				continue
				}
				if _ba00.strtUpppSccsBool == "flss" {
				goto step_22
				}
				break
			}
			}
			
			step_22:
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
			// ~Step 3.1: Receiving messages
			for _, _ba00 := range objc.trck {
				select {
				case _bb00 := <- _ba00.trck.Flap: {
					if strings.Index (_bb00.Sndr, _ba00.trck.Iddd) != 0 {
						continue
					}
					for _,  _ca00 := range objc.trck {
						if strings.Index (_bb00.Rcpn, _ca00.trck.Iddd) ==
							0 {
							_ba00.mssgList.PushBack (_bb00)
							break
						}
					}
					if _bb00.Rcpn == objc.mngrIddd {
						objc.mssgList.PushBack (_bb00)
					}
				}
				default: {}
				}
			}
			// ~Step 3.2: Pushing messages
			for _, _bc00 := range objc.trck {
				for {
					_ca00 := _bc00.mssgList.Front ()
					_bc00.mssgList.Remove (_ca00)
					for _, _cb00 := range _bc00.whttList {
						if strings.Index (_ca00.Value.(*Mssg).Sndr,
							_cb00) == 0 {
							select {
							case _bc00.trck.Clap <-
								_ca00.Value.(*Mssg): {
								goto next
							}
							default: {
								_bc00.mssgList.PushFront (_ca00)
								goto next
							}
							}
						}
					}
				}
				next:
			}
			// ~Step 3.3
			// xxxxxxxxx
		}
		} (objc, flap)
		
		return
	}

type trckTray_trck struct {
	trck *Trck
	whttList []string
	prvlBool bool
	strtUpppBool bool
	strtUpppSccsBool string
	lifeBool bool
	mssgList *list.List
}
func trckTray_hndlMssg (sndr string, mssg *Mssg) (bool) {

return false
}


/*
ba00: A component failed to start up
cb00: 

*/
