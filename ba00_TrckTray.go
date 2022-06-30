package tmcoll
import "container/list"
import "github.com/qb-qetell/errr"
import "time"

type TrckTray struct {
	MngrIddd string
	Trck []*TrckTray_trck
}
	func TrckTray_Estb (mngrIddd string) (*TrckTray) {
		return &TrckTray { MngrIddd: mngrIddd }
	}
	func (objc *TrckTray) Pplt (trck *Trck, whttList []string, prvlBool bool) {
		trckInst := &TrckTray_trck {}
		trckInst.Trck = trck
		trckInst.WhttList = whttList
		trckInst.PrvlBool = prvlBool
		trckInst.StrtBool = false
		trckInst.LifeBool = false
		trckInst.MssgList = list.New ()
		objc.Trck = append (objc.Trck, trckInst)
	}
	func (objc *TrckTray) Mngg () (error) {
		go func () {
			for _, _ba00 := range objc.Trck {
			go _ba00.Trck.Runn (objc.MngrIddd)
			for {
				if _ba00.StrtBool == false {
					time.Sleep (time.Nanosecond * 1000)
					continue
				}
				break
			}
			}
		} ()
		
		
return errr.Errr_Estb ("", "")
	}

type TrckTray_trck struct {
	Trck *Trck
	WhttList []string
	PrvlBool bool
	StrtBool bool
	LifeBool bool
	MssgList *list.List
}
func trckTray_plcc (sndr string, mssg *Mssg) {

}
