package tmcoll
import (
	_ "github.com/qb-qetell/errr"
)

type TrckTray struct {
	core []*trckTray_trck
}
	func TrckTray_Estb (mngrIddd string) (*TrckTray) { return &TrckTray {} }
	func (objc *TrckTray) Pplt (trck *Trck, whttList []string, prvlStts bool) {
		trckInst := &trckTray_trck {}
		trckInst.trck = trck
		trckInst.whttList = whttList
		trckInst.prvlStts = prvlStts
		objc.core = append (objc.core, trckInst)
	}
	func (objc *TrckTray) Mngg () (error) {
		return nil
	}

type trckTray_trck struct {
	trck *Trck
	whttList []string
	prvlStts bool
}
