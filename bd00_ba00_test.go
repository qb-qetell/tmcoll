package tmcoll
import "testing"
import "fmt"
import "github.com/qb-qetell/combGUID"
import "strconv"
import "time"

func Test_Ba00 (t *testing.T) {
	trckTray := TrckTray_Estb ("00")
	trckTray.Pplt (
		Trck_Estb ("00.01", "Tr_01", trck0001, map[string]interface {} {"00.01": "aa"}),
		[]string  {"00.02", "00.03"},
		[]string  {"cb00",  "cc00"},
	)
	trckTray.Pplt (
		Trck_Estb ("00.02", "Tr_02", trck0002, map[string]interface {} {"00.02": "bb"}),
		[]string  {"00.01", "00.03"},
		[]string  {"cb00",  "cc00"},
	)
	trckTray.Pplt (
		Trck_Estb ("00.03", "Tr_03", trck0003, map[string]interface {} {"00.03": "cc"}),
		[]string  {"00.01", "00.02"},
		[]string  {"cb00",  "cc00"},
	)
	clap, flap := trckTray.Mngg ()
	
	go func () {
		for i200 := 1000001; i200 <= 1000005; i200 ++ {
			Mssg_Estb (
				"00" + ".xx.nn",
				"00.01.nn",
				[]string {
					combGUID.CombGUID_Estb ("", 16).SmplFrmt (),
					strconv.Itoa (i200),
				},
			).Send (clap)
		}
		for i300 := 1000001; i300 <= 1000005; i300 ++ {
			Mssg_Estb (
				"00" + ".xx.nn",
				"00.03.nn",
				[]string {
					combGUID.CombGUID_Estb ("", 16).SmplFrmt (),
					strconv.Itoa (i300),
				},
			).Send (clap)
		}
		time.Sleep (time.Second * 12)
		_ba00 := []string {
			combGUID.CombGUID_Estb ("", 16).SmplFrmt (),
			"cb00",
		}
		Mssg_Estb ("00", "00!", _ba00).Send (clap)
	} ()
	
	for {
		_ba00 := <- flap
		fmt.Println ("----: HOST:", _ba00)
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
	Mssg_Estb (iddd, hostIddd + "!", _ba00).Send (flap)
	
	go func () {
		for i200 := 1000001; i200 <= 1000005; i200 ++ {
			Mssg_Estb (
				iddd + ".xx.nn",
				"00",
				[]string {
					combGUID.CombGUID_Estb ("", 16).SmplFrmt (),
					strconv.Itoa (i200),
				},
			).Send (flap)
		}
		for i300 := 1000001; i300 <= 1000005; i300 ++ {
			Mssg_Estb (
				iddd + ".xx.nn",
				"00.02.nn",
				[]string {
					combGUID.CombGUID_Estb ("", 16).SmplFrmt (),
					strconv.Itoa (i300),
				},
			).Send (flap)
		}
		time.Sleep (time.Second * 40000000)
		_ba00 := []string {
			combGUID.CombGUID_Estb ("", 16).SmplFrmt (),
			"cc00",
		}
		Mssg_Estb (iddd, hostIddd + "!", _ba00).Send (flap)
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
			Mssg_Estb (iddd, hostIddd + "!", _bd00).Send (flap)
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
	Mssg_Estb (iddd, hostIddd + "!", _ba00).Send (flap)
	
	go func () {
		time.Sleep (time.Second * 1)
		for i200 := 1000001; i200 <= 1000001; i200 ++ {
			Mssg_Estb (
				iddd + ".xx.nn",
				"00!",
				[]string {
					combGUID.CombGUID_Estb ("", 16).SmplFrmt (),
					"bt00",
				},
			).Send (flap)
		}
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
			Mssg_Estb (iddd, hostIddd + "!", _bd00).Send (flap)
			break
		}
		//time.Sleep (time.Second * 1)
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
	Mssg_Estb (iddd, hostIddd + "!", _ba00).Send (flap)
	//return
	/*
	time.Sleep (time.Second * 4)
	_ba20 := combGUID.CombGUID_Estb ("", 16).SmplFrmt ()
	_ba25 := []string {
		_ba20,
		"bt00",
	}
	Mssg_Estb (iddd, hostIddd + "!", _ba25).Send (flap)
	fmt.Println ("tr02:", _ba20)
	*/
	go func () {
		time.Sleep (time.Second * 8)
		_ba50 := []string {
			combGUID.CombGUID_Estb ("", 16).SmplFrmt (),
			"bz00",
		}
		Mssg_Estb (iddd, hostIddd + "!", _ba50).Send (flap)
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
			Mssg_Estb (iddd, hostIddd + "!", _bd00).Send (flap)
			break
		}
	}
}
