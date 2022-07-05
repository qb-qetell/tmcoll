package tmcoll
import "github.com/qb-qetell/combGUID"
import "testing"
import "fmt"

func Test_ (t *testing.T) {
	//
}

func smth01 (hostIddd, iddd, name string, clap <- chan *Mssg, flap chan <- *Mssg,
	seed map[string]interface{}) {
	_ba00 := fmt.Sprintf ("%s:bc00", combGUID.CombGUID_Estb ("tr", 14).SmplFrmt ())
	Mssg_Estb (iddd, hostIddd, _ba00).Send (clap)
}
