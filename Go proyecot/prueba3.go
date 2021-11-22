/*/*package main*/
/*
import (
	/*"bufio"
	"fmt"
	"net"
	"os"
	"strconv"
	"strings")
*/
/*
var (
	arregloIPA                                                                [50]float32
	arregloH                                                                  [50]int
	arregloM                                                                  [50]int
	n                                                                         = 0
	aux                                                                       = 0
	extra, grandei, g, grandeH, grandeM, arr2, arr3, ind, auxH, auxM, i, h, m int
	auxIPA, arr1                                                              float32
	respIPA, respH, respM, resp                                               string
)*/
/*
func main() {
	dats() // Carga la información en los arreglos

	ln, err := net.Listen("tcp", "localhost:9000") // Modo escucha

	if err != nil {
		fmt.Println("Falla al resolver la direccion de red", err.Error())
		os.Exit(1) // Finaliza el programa con error
	}
	// Conexión lista
	defer ln.Close()

	con, err := ln.Accept()
	if err != nil {
		fmt.Println("Falla al resolver la direccion de red", err.Error())
		os.Exit(1) // Finaliza el programa con error
	}
	// Cionexión aceptada
	defer con.Close()

	bufferIn := bufio.NewReader(con)
	msg, err := bufferIn.ReadString('\n')
	if err != nil {
		fmt.Println("Falla al resolver la direccion de red", err.Error())
		os.Exit(1) // Finaliza el programa con error
	}
	msg = strings.TrimSpace(msg) // Recibe el número identificador

	ind, err = strconv.Atoi(msg) // Transforma el número identificador en int
	if err != nil {
		fmt.Println("Falla al resolver la direccion de red", err.Error())
		os.Exit(1) // Finaliza el programa con error
	}

	extra = ind - 1 // Resta el número identificador en uno (ya que el 0 también cuenta)

	mast() // Calcula y compara los promedios

	// Asigna una respuesta con  base en los promedios
	if grandei == 0 && grandeM == 0 {
		g = 0
	} else if grandei == 1 && grandeM == 0 {
		g = 1
	} else if grandei == 1 && grandeM == 1 {
		g = 2
	}

	envResp(resp) // Envía la respuesta

}

func dats() { // Carga los datos
	go ipa()
	go her()
	go mort()
}

func mast() { // Calcula y compara promedios
	go promH(extra)
	go promIPA(extra)
	go promM(extra)
}

func envResp(resp string) { // Responde 4 veces
	var i = 0

	for i < 4 {
		if i == 0 {
			resp = fmt.Sprintf("%f", arregloIPA[extra]) // Carga el IPA del identificador elegido
		} else if i == 1 {
			resp = strconv.Itoa(arregloH[extra]) // Carga el indice de heridos del identificador elegido
		} else if i == 2 {
			resp = strconv.Itoa(arregloM[extra]) // Carga el indice de mortalidad del identificador elegido
		} else if i == 3 {
			if g == 0 { // Carga si es o no recomendable la empresa prestaria
				resp = "La empresa es recomendable para su viaje, ya que su índice de participación en accidentes es menor al promedio, al igual que su índice de accidentes letales"
			} else if g == 1 {
				resp = "La empresa puede recomendable para su viaje, ya que su índice de participación en accidentes no es muy superior al promedio"
			} else if g == 2 {
				resp = "La empresa no es recomendable para su viaje debido a su gran índice de participación en accidentes"
			}
		}
		con, err := net.Dial("tcp", "localhost:9001")
		if err != nil {
			fmt.Println("Falla al resolver la direccion de red", err.Error())
			os.Exit(1) // Finaliza el programa con error
		}
		fmt.Fprintln(con, resp) // Envia las respuesta
		if i == 4 {
			break // Cierra el ciclo
		}
		i++
	}
}

func promIPA(extra int) { // Compara el promedio del IPA
	auxIPA = 1.9

	arr1 = arregloIPA[extra]
	if arregloIPA[extra] < auxIPA {
		grandei = 1 //		El indice de IPA elegido está por encima del promedio
	} else {
		grandei = 0 //		El indice de IPA elegido está dentro del promedio
	}
}

func promH(extra int) { // Compara el promedio de heridos
	auxH = 10

	arr2 = arregloH[extra]
	if arregloH[extra] < auxH {
		grandeH = 1 //		El indice de heridos elegido está por encima del promedio
	} else {
		grandeH = 0 //		El indice de heridos elegido está dentro del promedio
	}
}

func promM(extra int) { // Compara el promedio de mortalidad
	auxM = 2

	arr3 = arregloM[extra]
	if arregloM[extra] < auxM {
		grandeM = 1 //		El indice de mortalidad elegido está por encima del promedio
	} else {
		grandeM = 0 //		El indice de mortalidad elegido está dentro del promedio
	}
}

func ipa() { // Datos de IPA
	arregloIPA[0] = 2.44
	arregloIPA[1] = 3.40
	arregloIPA[2] = 0.90
	arregloIPA[3] = 0.01
	arregloIPA[4] = 0.77
	arregloIPA[5] = 1.81
	arregloIPA[6] = 3.45
	arregloIPA[7] = 0.71
	arregloIPA[8] = 3.47
	arregloIPA[9] = 0.59
	arregloIPA[10] = 1.01
	arregloIPA[11] = 2.01
	arregloIPA[12] = 1.52
	arregloIPA[13] = 1.04
	arregloIPA[14] = 1.29
	arregloIPA[15] = 0.68
	arregloIPA[16] = 1.43
	arregloIPA[17] = 1.27
	arregloIPA[18] = 1.05
	arregloIPA[19] = 0.90
	arregloIPA[20] = 0.81
	arregloIPA[21] = 1.10
	arregloIPA[22] = 0.76
	arregloIPA[23] = 1.42
	arregloIPA[24] = 0.93
	arregloIPA[25] = 0.61
	arregloIPA[26] = 1.00
	arregloIPA[27] = 0.63
	arregloIPA[28] = 1.87
	arregloIPA[29] = 0.85
	arregloIPA[30] = 1.19
	arregloIPA[31] = 1.29
	arregloIPA[32] = 15.61
	arregloIPA[33] = 1.50
	arregloIPA[34] = 1.50
	arregloIPA[35] = 1.79
	arregloIPA[36] = 1.56
	arregloIPA[37] = 1.06
	arregloIPA[38] = 1.23
	arregloIPA[39] = 2.95
	arregloIPA[40] = 1.03
	arregloIPA[41] = 5.38
	arregloIPA[42] = 3.20
	arregloIPA[43] = 0.85
	arregloIPA[44] = 4.61
	arregloIPA[45] = 1.07
	arregloIPA[46] = 0.75
	arregloIPA[47] = 0.31
	arregloIPA[48] = 8.01
	arregloIPA[49] = 0.30
}

func her() { // Datos de heridos
	arregloH[0] = 10
	arregloH[1] = 6
	arregloH[2] = 4
	arregloH[3] = 1
	arregloH[4] = 23
	arregloH[5] = 10
	arregloH[6] = 11
	arregloH[7] = 3
	arregloH[8] = 18
	arregloH[9] = 8
	arregloH[10] = 13
	arregloH[11] = 9
	arregloH[12] = 14
	arregloH[13] = 9
	arregloH[14] = 1
	arregloH[15] = 4
	arregloH[16] = 10
	arregloH[17] = 9
	arregloH[18] = 16
	arregloH[19] = 5
	arregloH[20] = 2
	arregloH[21] = 15
	arregloH[22] = 6
	arregloH[23] = 5
	arregloH[24] = 14
	arregloH[25] = 3
	arregloH[26] = 10
	arregloH[27] = 5
	arregloH[28] = 5
	arregloH[29] = 6
	arregloH[30] = 36
	arregloH[31] = 14
	arregloH[32] = 18
	arregloH[33] = 11
	arregloH[34] = 5
	arregloH[35] = 9
	arregloH[36] = 1
	arregloH[37] = 6
	arregloH[38] = 4
	arregloH[39] = 15
	arregloH[40] = 1
	arregloH[41] = 36
	arregloH[42] = 14
	arregloH[43] = 5
	arregloH[44] = 7
	arregloH[45] = 3
	arregloH[46] = 5
	arregloH[47] = 8
	arregloH[48] = 43
	arregloH[49] = 2
}

func mort() { // Datos de mortalidad
	arregloM[0] = 1
	arregloM[1] = 1
	arregloM[2] = 0
	arregloM[3] = 0
	arregloM[4] = 3
	arregloM[5] = 1
	arregloM[6] = 2
	arregloM[7] = 0
	arregloM[8] = 3
	arregloM[9] = 0
	arregloM[10] = 1
	arregloM[11] = 1
	arregloM[12] = 4
	arregloM[13] = 0
	arregloM[14] = 1
	arregloM[15] = 1
	arregloM[16] = 0
	arregloM[17] = 2
	arregloM[18] = 1
	arregloM[19] = 1
	arregloM[20] = 1
	arregloM[21] = 8
	arregloM[22] = 1
	arregloM[23] = 0
	arregloM[24] = 0
	arregloM[25] = 2
	arregloM[26] = 1
	arregloM[27] = 0
	arregloM[28] = 1
	arregloM[29] = 1
	arregloM[30] = 3
	arregloM[31] = 0
	arregloM[32] = 2
	arregloM[33] = 0
	arregloM[34] = 0
	arregloM[35] = 0
	arregloM[36] = 1
	arregloM[37] = 0
	arregloM[38] = 0
	arregloM[39] = 1
	arregloM[40] = 1
	arregloM[41] = 5
	arregloM[42] = 0
	arregloM[43] = 0
	arregloM[44] = 0
	arregloM[45] = 1
	arregloM[46] = 0
	arregloM[47] = 1
	arregloM[48] = 5
	arregloM[49] = 0
}
*/