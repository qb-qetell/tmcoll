package tmcoll
import "testing"
import "fmt"
import "github.com/qb-qetell/combGUID"
import "time"

// Testing startup and shutdown
func Test_Ba00 (t *testing.T) {
	pack0001 := Trck_Estb ("01", "Track 01", trck0001, map[string]interface {} {"01": "aa"})
	pack0002 := Trck_Estb ("02", "Track 02", trck0001, map[string]interface {} {"02": "bb"})
	trckMngr := TrckTray_Estb ("00")
	trckMngr.Pplt (pack0001, []string {"02"}, false)
	trckMngr.Pplt (pack0002, []string {"01"}, false)
	clap, flap := trckMngr.Mngg ()
	go func () {
		for {
			_ba00 := <- flap
			fmt.Println ("xx:", _ba00)
		}
	} ()
	time.Sleep (time.Second * 20)
	for {
		_ba00 := []string {
			combGUID.CombGUID_Estb ("", 16).SmplFrmt (),
			"**30",
		}
		Mssg_Estb ("", "", _ba00).Send (clap)
	}
}

func trck0001 (hostIddd, iddd, name string, clap <- chan *Mssg, flap chan <- *Mssg,
	seed map[string]interface{}) {
	fmt.Println (hostIddd)
	fmt.Println (iddd)
	fmt.Println (name)
	fmt.Println (seed)
	_ba00 := []string {
		combGUID.CombGUID_Estb ("", 16).SmplFrmt (),
		"bc00",
	}
	Mssg_Estb (iddd, hostIddd, _ba00).Send (flap)
	for {
		_bb00 := <- clap
		_bc00 := _bb00.Core.([]string)
		fmt.Println ("tr01", _bb00.Sndr, _bb00.Rcpn, _bc00)
		if _bc00 [1] == "18aa" {
			_bd00 := []string {
				combGUID.CombGUID_Estb ("", 16).SmplFrmt (),
				"by00",
			}
			Mssg_Estb (iddd, hostIddd, _bd00).Send (flap)
		}
	}

}
func trck0002 (hostIddd, iddd, name string, clap <- chan *Mssg, flap chan <- *Mssg,
	seed map[string]interface{}) {
	fmt.Println (hostIddd)
	fmt.Println (iddd)
	fmt.Println (name)
	fmt.Println (seed)
	_ba00 := []string {
		combGUID.CombGUID_Estb ("", 16).SmplFrmt (),
		"bb00",
	}
	Mssg_Estb (iddd, hostIddd, _ba00).Send (flap)
	for {
		_bb00 := <- clap
		_bc00 := _bb00.Core.([]string)
		fmt.Println ("tr02", _bb00.Sndr, _bb00.Rcpn, _bc00)
		if _bc00 [1] == "18aa" {
			_bd00 := []string {
				combGUID.CombGUID_Estb ("", 16).SmplFrmt (),
				"by00",
			}
			Mssg_Estb (iddd, hostIddd, _bd00).Send (flap)
		}
	}
}
