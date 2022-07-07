package tmcoll
import "testing"
import "fmt"
import "github.com/qb-qetell/combGUID"
import "strconv"
import "time"

func Test_Ba00 (t *testing.T) {
	trckTray := TrckTray_Estb ("00")
	trckTray.Pplt (
		Trck_Estb ("01", "Tr01", trck0001, map[string]interface {} {"01": "aa"}),
		[]string  {"02", "03"},
		false,
	)
	trckTray.Pplt (
		Trck_Estb ("02", "Tr02", trck0002, map[string]interface {} {"02": "bb"}),
		[]string  {"01", "03"},
		false,
	)
	trckTray.Pplt (
		Trck_Estb ("03", "Tr03", trck0003, map[string]interface {} {"03": "cc"}),
		[]string  {"01", "02"},
		true,
	)
	clap, flap := trckTray.Mngg ()
	
	go func () {
		time.Sleep (time.Second * 1)
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
	for i200 := 1000001; i200 <= 1000005; i200 ++ {
		Mssg_Estb (iddd + ".xx.nn", "02", []string {strconv.Itoa (i200)}).Send (flap)
	}
	for i300 := 1000001; i300 <= 1000005; i300 ++ {
		Mssg_Estb (iddd + ".xx.nn", "03", []string {strconv.Itoa (i300)}).Send (flap)
	}
	go func () {
		time.Sleep (time.Second * 8)
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
	go func () {
		time.Sleep (time.Second * 4)
		_ba50 := []string {
			combGUID.CombGUID_Estb ("", 16).SmplFrmt (),
			"by00",
		}
		Mssg_Estb (iddd, hostIddd, _ba50).Send (flap)
	} ()
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
	go func () {
		time.Sleep (time.Second * 2)
		_ba50 := []string {
			combGUID.CombGUID_Estb ("", 16).SmplFrmt (),
			"by00",
		}
		Mssg_Estb (iddd, hostIddd, _ba50).Send (flap)
	} ()
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
