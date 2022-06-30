package tmcoll
import _ "github.com/qb-qetell/errr"

type TrckTray struct {
	Core []*TrckTray_trck
}
	func TrckTray_Estb (mngrIddd string) (*TrckTray) { return  &TrckTray {} }
	func (objc *TrckTray) Pplt (trck *Trck, whttList []string, prvlStts bool) {
		trckInst := &TrckTray_trck {}
		trckInst.trck = trck
		trckInst.whttList = whttList
		trckInst.prvlStts = prvlStts
		trckInst.lifeStts = false
		objc.Core = append (objc.Core, trckInst)
	}
	func (objc *TrckTray) Mngg () (error) {
		return nil
	}

type TrckTray_trck struct {
	trck *Trck
	whttList []string
	prvlStts bool
	lifeStts bool
}
