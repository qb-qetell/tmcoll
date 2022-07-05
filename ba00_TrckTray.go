package tmcoll
import "container/list"
import "fmt"
import "github.com/qb-qetell/combGUID"
import "regexp"
import "strconv"
import "strings"
import "sync"
import "time"

type TrckTray struct {
	mngrIddd string
	trck []*trckTray_trck
	mssgList *list.List
	mssgListMtxx *sync.Mutex
	shutDownInnnPrgsBool bool
	shutDownBool bool
}
	func TrckTray_Estb (mngrIddd string) (*TrckTray) {
		return &TrckTray {
			mngrIddd: mngrIddd,
			trck: []*trckTray_trck {},
			mssgList: list.New (),
			mssgListMtxx: &sync.Mutex {},
			shutDownInnnPrgsBool: false,
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
		trckInst.strtUpppMssg = ""
		trckInst.lifeBool = false
		trckInst.mssgList = list.New ()
		trckInst.mssgListMtxx = &sync.Mutex {}
		objc.trck = append (objc.trck, trckInst)
	}
	func (objc *TrckTray) Mngg () (clap chan *Mssg, flap chan *Mssg) {
		clap = make (chan *Mssg)
		flap = make (chan *Mssg)
		// Start up [failed/succesful] + shutdown
		
		go func (objc *TrckTray, flap chan <- *Mssg) {
		// ~Step 1
		if len (objc.trck) == 0 {
			_ca00 := []string {
				combGUID.CombGUID_Estb ("", 16).SmplFrmt (),
				"00fa",
				"No track to manage.",
			}
			_cb00 := Mssg_Estb (objc.mngrIddd, "", _ca00)
			flap <- _cb00
			return
		}
		dscvTrck := make (map[string]string)
		for _,  _ba00 := range objc.trck {
			if regexp.MustCompile (`^[a-z0-9]{1,}(\.[a-z0-9]{1,}){0,}$`,
				).MatchString (_ba00.trck.Iddd) == false {
				_ca00 := []string {
					combGUID.CombGUID_Estb ("", 16).SmplFrmt (),
					"00fa",
					"A track's ID is invalid.",
				}
				_cb00 := Mssg_Estb (objc.mngrIddd, "", _ca00)
				flap <- _cb00
				return
			}
			dscvTrck ["_bb00.trck.Iddd"] = "dscv"
		}
		for _,  _bb00 := range objc.trck {
			for _,  _ca00 := range _bb00.whttList {
				if dscvTrck [_ca00] != "dscv" {
					_da00 := []string {
						combGUID.CombGUID_Estb ("", 16).SmplFrmt (),
						"00fa",
						"A track specified a non-existent track on its "+
							"whitelist.",
						}
					_db00 := Mssg_Estb (objc.mngrIddd, "", _da00)
					flap <- _db00
					return
				}
			}
		}
		
		// ~Step 2
		go func (objc *TrckTray) {
			// ~Step 2.1: Start all tracks
			strtUpppSccsBool := true
			for _, _ba00 := range objc.trck {
				go _ba00.trck.Runn (objc.mngrIddd)
				for {
					if _ba00.strtUpppBool == false {
						time.Sleep (time.Microsecond * 1)
						continue
					}
					if _ba00.strtUpppSccsBool == "flss" {
						strtUpppSccsBool   = false
						_cb00 := fmt.Sprintf ("Could not start " +
							`track "%s : %s". [%s]`, _ba00.trck.Iddd,
							_ba00.trck.Name, _ba00.strtUpppMssg)
						_cc00 := []string {
							combGUID.CombGUID_Estb ("",
								16).SmplFrmt (),
							"00fa",
							_cb00 ,
						}
						_cd00 := Mssg_Estb (objc.mngrIddd, "", _cc00)
						flap <- _cd00
						goto next
					} else {
						_cb00 := fmt.Sprintf (`Track "%s : %s" could ` +
							`not start. [%s]`, _ba00.trck.Iddd,
							_ba00.trck.Name, _ba00.strtUpppMssg)
						_cc00 := []string {
							combGUID.CombGUID_Estb ("",
								16).SmplFrmt (),
							"00ga",
							_cb00 ,
						}
						_cd00 := Mssg_Estb (objc.mngrIddd, "", _cc00)
						flap <- _cd00
					}
					break
				}
			}
			
			next:
			// ~Step 2.2: If a track could not startup, send shutdown message
			if strtUpppSccsBool == false {
				_ce00 := []string {
					combGUID.CombGUID_Estb ("", 16).SmplFrmt (),
					"by00",
				}
				_cf00 := Mssg_Estb (objc.mngrIddd, objc.mngrIddd, _ce00)
				objc.mssgListMtxx.Lock ()
				objc.mssgList.PushBack (_cf00)
				objc.mssgListMtxx.Unlock ()
			} else {
				_ca00 := []string {
					combGUID.CombGUID_Estb ("", 16).SmplFrmt (),
					"00ma",
				}
				_cb00 := Mssg_Estb (objc.mngrIddd, "", _ca00)
				flap <- _cb00
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
							_ba00.mssgListMtxx.Lock ()
							_ba00.mssgList.PushBack (_bb00)
							_ba00.mssgListMtxx.Unlock ()
							break
						}
					}
					if _bb00.Rcpn == objc.mngrIddd {
						objc.mssgListMtxx.Lock ()
						objc.mssgList.PushBack (_bb00)
						objc.mssgListMtxx.Unlock ()
					}
				}
				default: {}
				}
			}
			// ~Step 3.2: Pushing messages
			for _, _bc00 := range objc.trck {
				for {
					_bc00.mssgListMtxx.Lock ()
					_ca00 := _bc00.mssgList.Front ()
					_bc00.mssgListMtxx.Unlock ()
					if _ca00 == nil { goto next }
					_bc00.mssgListMtxx.Lock ()
					_bc00.mssgList.Remove (_ca00)
					_bc00.mssgListMtxx.Unlock ()
					for _, _cb00 := range _bc00.whttList {
						if strings.Index (_ca00.Value.(*Mssg).Sndr,
							_cb00) == 0 {
							select {
							case _bc00.trck.Clap <-
								_ca00.Value.(*Mssg): {
								goto next
							}
							default: {
								_bc00.mssgListMtxx.Lock ()
								_bc00.mssgList.PushFront (_ca00)
								_bc00.mssgListMtxx.Unlock ()
								goto next
							}
							}
						}
					}
				}
				next:
			}
			
			// ~Step 3.3
			select {
			case _ca00 := <- clap: {
				_ca00 := _ca00.Core.([]string)
				if _ca00 != nil && len (_ca00) >= 2 {
					if        _ca00 [1] == "**30" {
						_ce00 := []string {
							combGUID.CombGUID_Estb ("",
								16).SmplFrmt (),
							"by00",
						}
						_cf00 := Mssg_Estb ("", objc.mngrIddd,
							_ce00)
						objc.mssgListMtxx.Lock ()
						objc.mssgList.PushBack (_cf00)
						objc.mssgListMtxx.Unlock ()
					} else if _ca00 [1] == "**60" {
						return
					}
				}
			}
			default: {}
			}
			
			// ~Step 3.4
			trckTray_hndlAaaaMssg (objc, flap)
			
			// ~Step 3.5
			if objc.shutDownBool == true {
				_cc00 := []string {
					combGUID.CombGUID_Estb ("", 16).SmplFrmt (),
					"00ta",
				}
				_cd00 := Mssg_Estb (objc.mngrIddd, "", _cc00)
				flap <- _cd00
				return
			}
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
	strtUpppMssg string
	lifeBool bool
	mssgList *list.List
	mssgListMtxx *sync.Mutex
}
func trckTray_hndlAaaaMssg (objc *TrckTray, flap chan <- *Mssg) {
	_ba00 := objc.mssgList.Front ().Value.(*Mssg)
	if _ba00 == nil    { return }
	_bb00, _bc00 := _ba00.Core.([]string)
	if _bc00 == false  { return }
	_bd00 := _bb00
	if len (_bd00) < 2 { return }
	if regexp.MustCompile (`^[a-z0-9]{16,16}$`).MatchString (_bd00 [0]) == false {
		return
	}
	if regexp.MustCompile (`^[a-z0-9]{4,4}$`  ).MatchString (_bd00 [1]) == false {
		return
	}
	
	if        _bd00 [1] == "bb00" {
	// Start-up failed
		for _, _ca00 := range objc.trck {
			if _ca00.trck.Iddd == _ba00.Sndr {
				_ca00.strtUpppSccsBool = "flss"
				_ca00.strtUpppMssg = _bd00 [2]
				_ca00.strtUpppBool = true
				break
			}
		}
	} else if _bd00 [1] == "bc00" {
	// Start-up successful
		for _, _ca00 := range objc.trck {
			if _ca00.trck.Iddd == _ba00.Sndr {
				_ca00.strtUpppSccsBool = "true"
				_ca00.lifeBool = true
				_ca00.strtUpppBool = true
				break
			}
		}
	} else if _bd00 [1] == "bm00" {
	// Track failed
		_ca00 := ""
		for _, _ca50 := range objc.trck {
			if _ca50.trck.Iddd == _ba00.Sndr {
				_ca00 = _ca50.trck.Iddd
				_ca50.lifeBool = false
				break
			}
		}
		_cb00 := ""
		if len (_bd00) > 2 { _cb00 = _bd00 [2] }
		_cc00 := []string {
			combGUID.CombGUID_Estb ("", 16).SmplFrmt (),
			"10aa",
			_ca00 ,
			_cb00 ,
		}
		_cd00 := Mssg_Estb (objc.mngrIddd, "", _cc00)
		trckTray_frwrToooPrvlTrck (objc, _cd00)
	} else if _bd00 [1] == "by00" {
	// Shutdown
		if objc.shutDownInnnPrgsBool == true { return }
		objc.shutDownInnnPrgsBool     = true
		go trckTray_hndlAaaaMssg_shutDownSyst (objc)
	} else if _bd00 [1] == "cb00" {
	// How many messages do I have?
		for _, _ca00 := range objc.trck {
			if _ca00.trck.Iddd == _ba00.Sndr {
				_cb00 := []string {
					_bd00 [0],
					"20aa",
					strconv.Itoa (_ca00.mssgList.Len ()),
				}
				_cc00 := Mssg_Estb (objc.mngrIddd, _ca00.trck.Iddd, _cb00)
				_ca00.mssgListMtxx.Lock ()
				_ca00.mssgList.PushBack (_cc00)
				_ca00.mssgListMtxx.Unlock ()
				break
			}
		}
	}
}
func trckTray_hndlAaaaMssg_shutDownSyst (objc *TrckTray) {
	for _ba00 := len (objc.trck); _ba00 >= 1; _ba00 -- {
		_bb00 := _ba00 - 1
		_bc00 := []string {
			combGUID.CombGUID_Estb ("", 16).SmplFrmt (),
			"18aa",
		}
		_bd00 := Mssg_Estb (objc.mngrIddd, objc.trck [_bb00].trck.Iddd, _bc00)
		objc.trck [_bb00].mssgListMtxx.Lock ()
		objc.trck [_bb00].mssgList.PushBack (_bd00)
		objc.trck [_bb00].mssgListMtxx.Unlock ()
		for {
			time.Sleep (time.Microsecond * 1)
			if objc.trck [_bb00].lifeBool == false { break }
		}
	}
	objc.shutDownBool = true
}
func trckTray_frwrToooPrvlTrck (objc *TrckTray, mssg *Mssg) {
	for _, _ca00 := range objc.trck {
		if _ca00.prvlBool == true {
			mssg.Rcpn = _ca00.trck.Iddd
			_ca00.mssgListMtxx.Lock ()
			_ca00.mssgList.PushBack (mssg)
			_ca00.mssgListMtxx.Unlock ()
		}
	}
}
