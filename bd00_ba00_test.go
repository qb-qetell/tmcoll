package tmcoll
import "testing"
import "fmt"
import "github.com/qb-qetell/combGUID"
import "strconv"
import "time"

// Testing startup and shutdown
func Test_Ba00 (t *testing.T) {
	pack0001 := Trck_Estb ("01", "Tr01", trck0001, map[string]interface {} {"01": "aa"})
	pack0002 := Trck_Estb ("02", "Tr02", trck0002, map[string]interface {} {"02": "bb"})
	pack0003 := Trck_Estb ("03", "Tr03", trck0003, map[string]interface {} {"03": "cc"})
	trckMngr := TrckTray_Estb ("00")
	trckMngr.Pplt (pack0001, []string {"02", "03"}, true )
	trckMngr.Pplt (pack0002, []string {"01", "03"}, false)
	trckMngr.Pplt (pack0003, []string {"01", "02"}, false)
	clap, flap := trckMngr.Mngg ()
	
	go func () {
		time.Sleep (time.Minute * 16)
		_ba00 := []string {
			combGUID.CombGUID_Estb ("", 16).SmplFrmt (),
			"**30",
		}
		Mssg_Estb ("", "", _ba00).Send (clap)
	} ()
	
	for {
		_ba00 := <- flap
		fmt.Println ("----:", _ba00)
		_bb00 := _ba00.Core.([]string)
		if _bb00 [1] == "00ta" { break }
	}
}

func trck0001 (hostIddd, iddd, name string, clap <- chan *Mssg, flap chan <- *Mssg,
	seed map[string]interface{}) {
	fmt.Println ("hidd:", hostIddd)
	fmt.Println ("tidd:", iddd)
	fmt.Println ("name:", name)
	fmt.Println ("seed:", seed)
	_ba00 := []string {
		combGUID.CombGUID_Estb ("", 16).SmplFrmt (),
		"bc00",
	}
	Mssg_Estb (iddd, hostIddd, _ba00).Send (flap)
	for i200 := 1000001; i200 <= 9999999; i200 ++ {
		Mssg_Estb (iddd + ".xx.nn", "02", []string {strconv.Itoa (i200)}).Send (flap)
	}
	for i300 := 1000001; i300 <= 9999999; i300 ++ {
		Mssg_Estb (iddd + ".xx.nn", "03", []string {strconv.Itoa (i300)}).Send (flap)
	}
	go func () {
		time.Sleep (time.Minute * 12)
		_ba00 := []string {
			combGUID.CombGUID_Estb ("", 16).SmplFrmt (),
			"by00",
		}
		Mssg_Estb (iddd, hostIddd, _ba00).Send (flap)
	} ()
	for {
		_bb00 := <- clap
		_bc00 := _bb00.Core.([]string)
		fmt.Println ("tr01:", _bb00.Sndr, _bb00.Rcpn, _bc00)
		if len (_bc00) > 1 && _bc00 [1] == "18aa" {
			_bd00 := []string {
				combGUID.CombGUID_Estb ("", 16).SmplFrmt (),
				"bz00",
			}
			Mssg_Estb (iddd, hostIddd, _bd00).Send (flap)
			break
		}
	}
}
func trck0002 (hostIddd, iddd, name string, clap <- chan *Mssg, flap chan <- *Mssg,
	seed map[string]interface{}) {
	fmt.Println ("hidd:", hostIddd)
	fmt.Println ("tidd:", iddd)
	fmt.Println ("name:", name)
	fmt.Println ("seed:", seed)
	_ba00 := []string {
		combGUID.CombGUID_Estb ("", 16).SmplFrmt (),
		"bc00",
	}
	Mssg_Estb (iddd, hostIddd, _ba00).Send (flap)
	for {
		_bb00 := <- clap
		_bc00 := _bb00.Core.([]string)
		fmt.Println ("tr02:", _bb00.Sndr, _bb00.Rcpn, _bc00)
		if len (_bc00) > 1 && _bc00 [1] == "18aa" {
			_bd00 := []string {
				combGUID.CombGUID_Estb ("", 16).SmplFrmt (),
				"bz00",
			}
			Mssg_Estb (iddd, hostIddd, _bd00).Send (flap)
			break
		}
	}
}
func trck0003 (hostIddd, iddd, name string, clap <- chan *Mssg, flap chan <- *Mssg,
	seed map[string]interface{}) {
	fmt.Println ("hidd:", hostIddd)
	fmt.Println ("tidd:", iddd)
	fmt.Println ("name:", name)
	fmt.Println ("seed:", seed)
	_ba00 := []string {
		combGUID.CombGUID_Estb ("", 16).SmplFrmt (),
		"bc00",
	}
	Mssg_Estb (iddd, hostIddd, _ba00).Send (flap)
	for {
		_bb00 := <- clap
		_bc00 := _bb00.Core.([]string)
		fmt.Println ("tr03:", _bb00.Sndr, _bb00.Rcpn, _bc00)
		if len (_bc00) > 1 && _bc00 [1] == "18aa" {
			_bd00 := []string {
				combGUID.CombGUID_Estb ("", 16).SmplFrmt (),
				"bz00",
			}
			Mssg_Estb (iddd, hostIddd, _bd00).Send (flap)
			break
		}
	}
}
