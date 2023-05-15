package main

import (
	"encoding/binary"
	"flag"
	"fmt"
	"io/ioutil"
	"strings"
)

func gasboy(inputdata []byte, option int) (outdata []byte) {
	gas1 := binary.BigEndian.Uint32(inputdata[1232:1236])
	gas2 := binary.BigEndian.Uint32(inputdata[1260:1264])
	gas3 := binary.BigEndian.Uint32(inputdata[1288:1292])
	gas4 := binary.BigEndian.Uint32(inputdata[1316:1320])

	gastong := []byte{}
	if (option == 1 && (gas1 == 1 || gas1 == 3 || gas1 == 6 || gas1 == 7)) || (option == 2 && (gas1 == 2 || gas1 == 6)) || (option == 3 && gas1 == 4) {
		gastong = inputdata[1224:1248]
	} else if (option == 1 && (gas2 == 1 || gas2 == 3 || gas2 == 6 || gas2 == 7)) || (option == 2 && (gas2 == 2 || gas2 == 6)) || (option == 3 && gas2 == 4) {
		gastong = inputdata[1252:1276]
	} else if (option == 1 && (gas3 == 1 || gas3 == 3 || gas3 == 6 || gas3 == 7)) || (option == 2 && (gas3 == 2 || gas3 == 6)) || (option == 3 && gas3 == 4) {
		gastong = inputdata[1280:1304]
	} else if (option == 1 && (gas4 == 1 || gas4 == 3 || gas4 == 6 || gas4 == 7)) || (option == 2 && (gas4 == 2 || gas4 == 6)) || (option == 3 && gas4 == 4) {
		gastong = inputdata[1304:1328]
	} else {
		fmt.Println("no target")
		outdata = nil
		return
	}

	gasfix := 0
	if (gas4 >= 1 && gas4 <= 7) || gas4 == 512 {
		gasfix = 28
	} else if (gas3 >= 1 && gas3 <= 7) || gas3 == 512 {
	} else if (gas2 >= 1 && gas2 <= 7) || gas2 == 512 {
		gasfix = -28
	} else if (gas1 >= 1 && gas1 <= 7) || gas1 == 512 {
		gasfix = -56
	}

	gastong_start := binary.BigEndian.Uint32(gastong[0:4])
	gastong_size := binary.BigEndian.Uint32(gastong[4:8])
	hanaread := inputdata[1308+gasfix : 1312+gasfix]

	if hanaread[0] == byte('H') && hanaread[1] == byte('A') && hanaread[2] == byte('N') && hanaread[3] == byte('A') {
		outdata = inputdata[gastong_start : gastong_start+gastong_size]
		fmt.Println("OK")
		return
	} else {
		fmt.Println("HANB pass")
		outdata = nil
		return
	}
}

func main() {
	fmt.Println("GAS-BOY V1 (public)")
	file := flag.String("file", "00000.KY2", "Input file")
	option := flag.Int("option", 1, "target option (1=DREAM, 2=88, 3=8820)")
	flag.Parse()
	dat, err := ioutil.ReadFile(*file)
	if err != nil {
		fmt.Println(err.Error())
		//fmt.Println("File not found")
	} else {
		kycdata := gasboy(dat, *option)
		if kycdata != nil {
			filename := *file
			err := ioutil.WriteFile(filename[:len(filename)-3]+strings.ToUpper(filename[len(filename)-3:len(filename)-1]+"C"), kycdata, 0755)
			if err != nil {
				fmt.Println(err.Error())
				return
			}
		}
	}
}
