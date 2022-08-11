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
	mngrIddd             string
	trck                 []*trckTray_trck
	hostMssgList         *list.List
	hostMssgListMtxx     *sync.Mutex
	systMssgList         *list.List
	systMssgListMtxx     *sync.Mutex
	shutDownInnnPrgsBool bool
	shutDownBool         bool
}
	func TrckTray_Estb (mngrIddd string) (*TrckTray) {
		return &TrckTray {
			mngrIddd:             mngrIddd,
			trck:                 []*trckTray_trck {},
			hostMssgList:             list.New (),
			hostMssgListMtxx:         &sync.Mutex  {},
			systMssgList:             list.New (),
			systMssgListMtxx:         &sync.Mutex  {},
			shutDownInnnPrgsBool: false,
			shutDownBool:         false,
		}
	}
	func (objc *TrckTray) Pplt (trck *Trck, whttList []string, prvl []string) {
		trckInst := &trckTray_trck {}
		trckInst.trck =             trck
		trckInst.whttList =         whttList
		trckInst.whttList = append (trckInst.whttList, objc.mngrIddd + "!")
		trckInst.whttList = append (trckInst.whttList, objc.mngrIddd)
		trckInst.prvl =             prvl
		trckInst.strtUpppBool =     false
		trckInst.strtUpppSccsBool = "undf"
		trckInst.strtUpppMssg =     ""
		trckInst.lifeBool =         false
		trckInst.mssgList =         list.New ()
		trckInst.mssgListMtxx =     &sync.Mutex {}
		objc.trck = append (objc.trck, trckInst)
	}
	func (objc *TrckTray) Mngg ()   (clap chan *Mssg, flap chan *Mssg) {
		clap = make (chan *Mssg)
		flap = make (chan *Mssg)
		// Start up [failed/succesful] + shutdown
		
		go func (objc *TrckTray, clap chan *Mssg, flap chan *Mssg) {
			// ~Step 1
			if len (objc.trck) == 0 {
				_ca00 := []string {
					combGUID.CombGUID_Estb ("", 16).SmplFrmt (),
					"00fa",
					"No track to manage.",
				}
				_cb00 := Mssg_Estb (objc.mngrIddd + "!", objc.mngrIddd, _ca00)
				flap <- _cb00
				_cc00 := []string {
					combGUID.CombGUID_Estb ("", 16).SmplFrmt (),
					"00ta",
				}
				_cd00 := Mssg_Estb (objc.mngrIddd + "!", objc.mngrIddd, _cc00)
				flap <- _cd00
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
					_cb00 := Mssg_Estb (objc.mngrIddd + "!", objc.mngrIddd,
						_ca00)
					flap <- _cb00
					_cc00 := []string {
						combGUID.CombGUID_Estb ("", 16).SmplFrmt (),
						"00ta",
					}
					_cd00 := Mssg_Estb (objc.mngrIddd + "!", objc.mngrIddd,
						_cc00)
					flap <- _cd00
					return
				}
				dscvTrck [_ba00.trck.Iddd] = "dscv"
			}
			for _,  _bb00 := range objc.trck {
				for _,  _ca00 := range _bb00.whttList {
					if dscvTrck [_ca00] != "dscv" && _ca00 !=
					(objc.mngrIddd + "!") && _ca00 != objc.mngrIddd {
						_da00 := []string {
							combGUID.CombGUID_Estb ("", 16,
								).SmplFrmt (),
							"00fa",
							"A track's whitelist contains a " +
								"non-existent track's id.",
						}
						_db00 := Mssg_Estb (objc.mngrIddd + "!",
							objc.mngrIddd, _da00)
						flap <- _db00
						_dc00 := []string {
							combGUID.CombGUID_Estb ("", 16,
								).SmplFrmt(),
							"00ta",
						}
						_dd00 := Mssg_Estb (objc.mngrIddd + "!",
							objc.mngrIddd, _dc00)
						flap <- _dd00
						return
					}
				}
			}
			for _,  _bb00 := range objc.trck {
				if _bb00.prvl == nil   { continue }
				for _,  _ca00 := range _bb00.prvl {
					if _ca00 != "cb00" && _ca00 != "cc00" {
						_da00 := []string {
							combGUID.CombGUID_Estb ("", 16,
								).SmplFrmt (),
							"00fa",
							"A track specified has an invalid " +
								"privilege.",
						}
						_db00 := Mssg_Estb (objc.mngrIddd + "!",
							objc.mngrIddd, _da00)
						flap <- _db00
						_dc00 := []string {
							combGUID.CombGUID_Estb ("", 16,
								).SmplFrmt (),
							"00ta",
						}
						_dd00 := Mssg_Estb (objc.mngrIddd + "!",
							objc.mngrIddd, _dc00)
						flap <- _dd00
						return
					}
				}
			}
			// ~Step 2
			go func (objc *TrckTray) {
				strtUpppSccsBool := true
				// ~Step 2.1: Start all tracks
				for _, _ba00 := range objc.trck {
					go _ba00.trck.Runn (objc.mngrIddd)
					for {
						if _ba00.strtUpppBool == false {
							time.Sleep (time.Millisecond * 1)
							continue
						}
						if _ba00.strtUpppSccsBool == "flss" {
							strtUpppSccsBool   =  false
							_cb00 := fmt.Sprintf (`Track "%s:%s" ` +
								`could not start. [%s]`,
								_ba00.trck.Iddd,
								_ba00.trck.Name,
								_ba00.strtUpppMssg,
							)
							_cc00 := []string {
								combGUID.CombGUID_Estb ("",
									16).SmplFrmt (),
								"00fa",
								_cb00 ,
							}
							_cd00 := Mssg_Estb (objc.mngrIddd + "!",
								objc.mngrIddd,
								_cc00)
							flap <- _cd00
							goto next
						} else {
							_cb00 := fmt.Sprintf (`Track "%s:%s" ` +
								`has started.`, _ba00.trck.Iddd,
								_ba00.trck.Name,
							)
							_cc00 := []string {
								combGUID.CombGUID_Estb ("",
									16).SmplFrmt (),
								"00ga",
								_cb00 ,
							}
							_cd00 := Mssg_Estb (objc.mngrIddd + "!",
								objc.mngrIddd,
								_cc00)
							flap <- _cd00
						}
						break
					}
				}
				// ~Step 2.2: If a track could not startup, send shutdown message
				next:
				if strtUpppSccsBool == false {
					_ce00 := []string {
						combGUID.CombGUID_Estb ("", 16).SmplFrmt (),
						"cb00",
					}
					_cf00 := Mssg_Estb (objc.mngrIddd + "!", objc.mngrIddd +
						"!", _ce00)
					objc.systMssgListMtxx.Lock ()
					objc.systMssgList.PushBack (_cf00)
					objc.systMssgListMtxx.Unlock ()
				} else {
					_ck00 := []string {
						combGUID.CombGUID_Estb ("", 16).SmplFrmt (),
						"00ma",
					}
					_cl00 := Mssg_Estb (objc.mngrIddd + "!", objc.mngrIddd,
						_ck00)
					flap <- _cl00
				}
			} (objc)
			// ~Step 3
			mssgTrck := []*trckTray_trck {}
			mssgTrck = append (mssgTrck, objc.trck...)
			systTrck :=   &trckTray_trck {
				trck:             Trck_Estb (objc.mngrIddd + "!", "", nil),
				mssgList:         objc.systMssgList,
				mssgListMtxx:     objc.systMssgListMtxx,
			}
			whttList := []string {}
			for _, _ak00:= range objc.trck {
				whttList = append (whttList, _ak00.trck.Iddd)
			}
			systTrck.whttList = whttList
			whttList = append (whttList, objc.mngrIddd + "!")
			whttList = append (whttList, objc.mngrIddd)
			mssgTrck = append (mssgTrck, systTrck)
			hostTrck :=   &trckTray_trck {
				trck:             Trck_Estb (objc.mngrIddd,       "", nil),
				whttList:         whttList,
				mssgList:         objc.hostMssgList,
				mssgListMtxx:     objc.hostMssgListMtxx,
			}
			hostTrck.trck.Clap = flap
			hostTrck.trck.Flap = clap
			mssgTrck = append (mssgTrck, hostTrck)
			for {
				wrkd := false
				// ~Step 3.1: Receiving messages
				for _, _ba00 := range mssgTrck {
					select {
					case _bb00 := <- _ba00.trck.Flap: {
						wrkd = true
						if strings.Index (_bb00.Sndr,
							_ba00.trck.Iddd) != 0 {
							goto nxt1
						}
						for _,  _ca00 := range mssgTrck {
							if strings.Index (_bb00.Rcpn,
								_ca00.trck.Iddd) == 0 {
								_ca00.mssgListMtxx.Lock ()
								_ca00.mssgList.PushBack (_bb00)
								_ca00.mssgListMtxx.Unlock ()
								break
							}
						}
					}
					default: {}
					}
					nxt1:
				}
				// ~Step 3.2: Pushing messages
				for _, _bc00 := range mssgTrck {
					if _bc00.trck.Iddd == (objc.mngrIddd + "!") { continue }
					for {
						_bc00.mssgListMtxx.Lock ()
						_ca00 := _bc00.mssgList.Front ()
						_bc00.mssgListMtxx.Unlock ()
						if _ca00 == nil { goto nxt2 }
						wrkd = true
						_bc00.mssgListMtxx.Lock ()
						_bc00.mssgList.Remove (_ca00)
						_bc00.mssgListMtxx.Unlock ()
						for _, _cb00 := range _bc00.whttList {
							if strings.Index (
							_ca00.Value.(*Mssg).Sndr, _cb00) == 0 {
								select {
								case _bc00.trck.Clap <-
									_ca00.Value.(*Mssg): {
									goto nxt2
								}
								default: {
									_bc00.mssgListMtxx.Lock (
									)
									_bc00.mssgList.PushFront(
										_ca00.Value.(
										*Mssg),
									)
								_bc00.mssgListMtxx.Unlock ()
									goto nxt2
								}
								}
							} else {
							}
						}
					}
					nxt2:
				}
				// ~Step 3.3
				_cb00 := trckTray_hndlAaaaMssg (objc, flap)
				if _cb00 == true { wrkd = true }
				// ~Step 3.4
				if objc.shutDownBool == true {
					_cc00 := []string {
						combGUID.CombGUID_Estb ("", 16).SmplFrmt (),
						"00ta",
					}
					_cd00 := Mssg_Estb (objc.mngrIddd + "!", objc.mngrIddd,
						_cc00)
					flap <- _cd00
					return
				}
				// ~Step 3.5
				if wrkd == false { time.Sleep (time.Millisecond * 1) }
			}
		} (objc, clap, flap)
		
		return
	}

type trckTray_trck struct {
	trck             *Trck
	whttList         []string
	prvl             []string
	strtUpppBool     bool
	strtUpppSccsBool string
	strtUpppMssg     string
	lifeBool         bool
	mssgList         *list.List
	mssgListMtxx     *sync.Mutex
}
func trckTray_hndlAaaaMssg (objc *TrckTray, flap chan <- *Mssg) (wrkd bool) {
	wrkd = false
	
	if objc.systMssgList.Len () == 0 { return }
	objc.systMssgListMtxx.Lock ()
	_aa50 := objc.systMssgList.Front ()
	objc.systMssgListMtxx.Unlock ()
	if _aa50 == nil { return }
	_ba00 := _aa50.Value.(*Mssg)
	if _ba00 == nil { return }
	objc.systMssgListMtxx.Lock ()
	objc.systMssgList.Remove (_aa50)
	objc.systMssgListMtxx.Unlock ()
	_bb00, _bc00 := _ba00.Core.([]string)
	if _bc00 ==  false { return }
	_bd00 := _bb00
	if len (_bd00) < 2 { return }
	if regexp.MustCompile (`^[a-z0-9]{16,16}$`).MatchString (_bd00 [0]) == false { return }
	if regexp.MustCompile (`^[a-z0-9]{4,4}$`  ).MatchString (_bd00 [1]) == false { return }
	wrkd = true
	
	if        _bd00 [1] == "bb00" {
	// Track: Start-up failed
		for _, _ca00 := range objc.trck {
			if _ba00.Sndr == _ca00.trck.Iddd {
				_ca00.strtUpppSccsBool = "flss"
				if len (_bd00) > 2 {
					_ca00.strtUpppMssg = _bd00 [2]
				}
				_ca00.strtUpppBool = true
				return
			}
		}
	} else if _bd00 [1] == "bc00" {
	// Track: Start-up successful
		for _, _ca00 := range objc.trck {
			if _ba00.Sndr == _ca00.trck.Iddd {
				_ca00.strtUpppSccsBool = "true"
				_ca00.lifeBool = true
				_ca00.strtUpppBool = true
				return
			}
		}
	} else if _bd00 [1] == "bm00" {
	// Track: Failed
		_ca00 := ""
		_ca25 := ""
		for _, _ca50 := range objc.trck {
			if _ba00.Sndr == _ca50.trck.Iddd {
				_ca00 =  _ca50.trck.Iddd
				_ca25 =  _ca50.trck.Name
				_ca50.lifeBool = false
				break
			}
		}
		if _ca00 == "" { return }
		_cb00 := ""
		if len (_bd00) > 2 { _cb00 = _bd00 [2] }
		_cc00 := []string {
			combGUID.CombGUID_Estb ("", 16).SmplFrmt (),
			"10aa",
			_ca00 ,
			_ca25 ,
			_cb00 ,
		}
		_cd00 := Mssg_Estb (objc.mngrIddd + "!", objc.mngrIddd, _cc00)
		trckTray_frwrToooPrvlTrck (objc, _cd00)
	} else if _bd00 [1] == "bt00" {
	// Track: How many messages do I have?
		for _, _ca00 := range objc.trck {
			if strings.Index (_ba00.Sndr, _ca00.trck.Iddd) == 0 {
				_ca00.mssgListMtxx.Lock ()
				_cb00 := []string {
					_bd00 [0],
					"20aa",
					strconv.Itoa (_ca00.mssgList.Len ()),
				}
				_cc00 := Mssg_Estb (objc.mngrIddd + "!", _ba00.Sndr, _cb00)
				_ca00.mssgList.PushFront (_cc00)
				_ca00.mssgListMtxx.Unlock ()
				break
			}
		}
	} else if _bd00 [1] == "bz00" {
	// Track: Shutdown
		_ca00 := ""
		_ca25 := ""
		for _, _ca50 := range objc.trck {
			if _ba00.Sndr == _ca50.trck.Iddd {
				_ca50.lifeBool = false
				_ca00 =  _ca50.trck.Iddd
				_ca25 =  _ca50.trck.Name
				_ca50.lifeBool = false
				break
			}
		}
		if _ca00 == "" { return }
		_cb00 := ""
		if len (_bd00) > 2 { _cb00 = _bd00 [2] }
		_cc00 := []string {
			combGUID.CombGUID_Estb ("", 16).SmplFrmt (),
			"30aa",
			_ca00 ,
			_ca25 ,
			_cb00 ,
		}
		_cd00 := Mssg_Estb (objc.mngrIddd + "!", objc.mngrIddd, _cc00)
		trckTray_frwrToooPrvlTrck (objc, _cd00)

	} else if _bd00 [1] == "cb00" {
	// Track Management: shutdown gracefully
		if _ba00.Sndr != objc.mngrIddd && _ba00.Sndr != (objc.mngrIddd + "!") { return }
		if objc.shutDownInnnPrgsBool == true { return }
		objc.shutDownInnnPrgsBool     = true
		go trckTray_hndlAaaaMssg_shutDownSyst (objc)
	} else if _bd00 [1] == "cc00" {
	// Track Management: shutdown imediately
		if _ba00.Sndr != objc.mngrIddd && _ba00.Sndr != (objc.mngrIddd + "!") { return }
                objc.shutDownBool = true
	}
	
	return
}
func trckTray_hndlAaaaMssg_shutDownSyst (objc *TrckTray) {
	for _ba00 := len (objc.trck); _ba00 >= 1; _ba00 -- {
		_bb00 := _ba00 - 1
		_bb50 := objc.trck [_bb00]
		_bc00 := []string {
			combGUID.CombGUID_Estb ("", 16).SmplFrmt (),
			"18aa",
		}
		_bd00 := Mssg_Estb (objc.mngrIddd + "!", _bb50.trck.Iddd, _bc00)
		_bb50.mssgListMtxx.Lock ()
		_bb50.mssgList.PushBack (_bd00)
		_bb50.mssgListMtxx.Unlock ()
		for {
			time.Sleep (time.Millisecond * 1)
			if _bb50.lifeBool == false { break }
		}
	}
	objc.shutDownBool = true
}
func trckTray_frwrToooPrvlTrck (objc *TrckTray, mssg *Mssg) {
	objc.hostMssgListMtxx.Lock ()
	objc.hostMssgList.PushBack (mssg)
	objc.hostMssgListMtxx.Unlock ()
}
