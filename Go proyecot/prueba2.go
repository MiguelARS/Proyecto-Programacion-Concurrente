package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
	"strconv"

	"github.com/gorilla/mux"
)

var remotehost string
var id int

type empresa struct {
	Id      int     `json:Id`
	Empresa string  `json:Empresa`
	Ipa     float64 `json:Ipa`
	Her     int     `json:Her`
	Mort    int     `json:Mort`
	Result  string  `json:Result`
}

type allEmpresas []empresa

var (
	arregloIPA                                                                [50]float32
	arregloH                                                                  [50]int
	arregloM                                                                  [50]int
	n                                                                         = 0
	aux                                                                       = 0
	extra, grandei, g, grandeH, grandeM, arr2, arr3, ind, auxH, auxM, i, h, m int
	auxIPA, arr1                                                              float32
	respIPA, respH, respM, resp                                               string
	emterprise                                                                empresa
	err                                                                       error
)

var empresas = allEmpresas{
	{
		Id:      1,
		Empresa: "EMPRESA DE TRANSPORTES APOCALIPSIS S.A.",
		Ipa:     0,
		Her:     0,
		Mort:    0,
		Result:  "",
	},
	{
		Id:      2,
		Empresa: "TRANSPORTES Y TURISMO INTERNACIONAL CARLITOS S.A.C.",
		Ipa:     0,
		Her:     0,
		Mort:    0,
		Result:  "",
	},
	{
		Id:      3,
		Empresa: "INTERNACIONAL CHALLENGERS S.A.C.",
		Ipa:     0,
		Her:     0,
		Mort:    0,
		Result:  "",
	},
	{
		Id:      4,
		Empresa: "TRANSPORTES LINEA S.A.",
		Ipa:     0,
		Her:     0,
		Mort:    0,
		Result:  "",
	},
	{
		Id:      5,
		Empresa: "EMPRESA DE TRANSPORTES PERU BUS S.A.",
		Ipa:     0,
		Her:     0,
		Mort:    0,
		Result:  "",
	},
	{
		Id:      6,
		Empresa: "EMPE TRANSP SALAZAR EIRL ETRANSA",
		Ipa:     0,
		Her:     0,
		Mort:    0,
		Result:  "",
	},
	{
		Id:      7,
		Empresa: "EMPRESA DE TRANSPORTES GUTARRA S.A.",
		Ipa:     0,
		Her:     0,
		Mort:    0,
		Result:  ""},
	{
		Id:      8,
		Empresa: "TRANSPORTES ANITA E.I.R.L.",
		Ipa:     0,
		Her:     0,
		Mort:    0,
		Result:  "",
	},
	{
		Id:      9,
		Empresa: "EXPRESO POWER E.I.R.L.",
		Ipa:     0,
		Her:     0,
		Mort:    0,
		Result:  ""},
	{
		Id:      10,
		Empresa: "EXPRESO ANTEZANA HNOS. S.A.",
		Ipa:     0,
		Her:     0,
		Mort:    0,
		Result:  "",
	},
	{
		Id:      11,
		Empresa: "EMPRESA DE TRANSPORTES JOSE HUAPAYA SORIANO S.A.",
		Ipa:     0,
		Her:     0,
		Mort:    0,
		Result:  "",
	},
	{
		Id:      12,
		Empresa: "TRANSPORTE WARI S.A.C.",
		Ipa:     0,
		Her:     0,
		Mort:    0,
		Result:  ""},
	{
		Id:      13,
		Empresa: "TRANSPORTES Y TURISMO REYNA S.R.L.",
		Ipa:     0,
		Her:     0,
		Mort:    0, Result: ""},
	{
		Id:      14,
		Empresa: "EMPRESA DE TRANSPORTES TICLLAS S.A.C. - E.T.T.I.C.S.A.C.",
		Ipa:     0,
		Her:     0,
		Mort:    0,
		Result:  ""},
	{
		Id:      15,
		Empresa: "EMPRESA DE TRANSPORTES INTERPROVINCIAL DE PASAJEROS JULI BUSS SOCIEDAD ANONIMA CERRADA",
		Ipa:     0,
		Her:     0,
		Mort:    0,
		Result:  ""},
	{
		Id:      16,
		Empresa: "EMPRESA DE TRANSPORTES EL SOL S.A.",
		Ipa:     0,
		Her:     0,
		Mort:    0,
		Result:  ""},
	{
		Id:      17,
		Empresa: "EMPRESA DE TRANSPORTES RONCO PERU S.A.C.",
		Ipa:     0,
		Her:     0,
		Mort:    0,
		Result:  ""},
	{
		Id:      18,
		Empresa: "TRANSPORTES CROMOTEX S.A.C.",
		Ipa:     0,
		Her:     0,
		Mort:    0,
		Result:  ""},
	{
		Id:      19,
		Empresa: "TRANSMAR EXPRESS SAC",
		Ipa:     0,
		Her:     0,
		Mort:    0,
		Result:  ""},
	{
		Id:      20,
		Empresa: "EMPRESA DE TRANSPORTES GRUPO HORNA S.A.C.",
		Ipa:     0,
		Her:     0,
		Mort:    0,
		Result:  ""},
	{
		Id:      21,
		Empresa: "ANDORI�A TOURS S.R.L",
		Ipa:     0,
		Her:     0,
		Mort:    0,
		Result:  ""},
	{
		Id:      22,
		Empresa: "EMPRESA DE TRANSPORTES EXPRESO INTERNACIONAL PALOMINO S.A.C.",
		Ipa:     0,
		Her:     0,
		Mort:    0,
		Result:  ""},
	{
		Id:      23,
		Empresa: "MOVIL TOURS S.A.",
		Ipa:     0,
		Her:     0,
		Mort:    0,
		Result:  ""},
	{
		Id:      24,
		Empresa: "TRANSPORTES MERCEDES S.R.L.",
		Ipa:     0,
		Her:     0,
		Mort:    0,
		Result:  ""},
	{
		Id:      25,
		Empresa: "EMPRESA DE TRANSPORTES EXPRESO NACIONAL CERRO DE PASCO S.R.LTDA.",
		Ipa:     0,
		Her:     0,
		Mort:    0,
		Result:  ""},
	{
		Id:      26,
		Empresa: "EMPRESA DE TRANSPORTES AVE FENIX S.A.C.",
		Ipa:     0,
		Her:     0,
		Mort:    0,
		Result:  ""},
	{
		Id:      27,
		Empresa: "EXPRESO SANCHEZ S.R.L.",
		Ipa:     0,
		Her:     0,
		Mort:    0,
		Result:  ""},
	{
		Id:      28,
		Empresa: "PLUMA BUS S.R.L.",
		Ipa:     0,
		Her:     0,
		Mort:    0,
		Result:  ""},
	{
		Id:      29,
		Empresa: "TRANSPORTES MENDO EIRL",
		Ipa:     0,
		Her:     0,
		Mort:    0,
		Result:  ""},
	{
		Id:      30,
		Empresa: "CORPORACION CHASKIS INDUSTRIA COMERCIO Y SERVICIOS S.A.C.",
		Ipa:     0,
		Her:     0,
		Mort:    0,
		Result:  ""},
	{
		Id:      31,
		Empresa: "EMPRESA DE TRANSPORTES FLORES HerMANOS S.C.R.LTDA.",
		Ipa:     0,
		Her:     0,
		Mort:    0,
		Result:  ""},
	{
		Id:      32,
		Empresa: "TOURS ANGEL DIVINO S.A.C.",
		Ipa:     0,
		Her:     0,
		Mort:    0,
		Result:  ""},
	{
		Id:      33,
		Empresa: "ASESORES Y CONSULTORES VIA SEGURA S.A.C.",
		Ipa:     0,
		Her:     0,
		Mort:    0,
		Result:  ""},
	{
		Id:      34,
		Empresa: "EXPRESO INTERNACIONAL MOLINA S.A.C.",
		Ipa:     0,
		Her:     0,
		Mort:    0,
		Result:  ""},
	{
		Id:      35,
		Empresa: "EXPRESO TURISMO ANDINO S.A.",
		Ipa:     0,
		Her:     0,
		Mort:    0,
		Result:  ""},
	{
		Id:      36,
		Empresa: "EMPRESA DE TRANSPORTES EXPRESO ROMATISA E.I.R.L.",
		Ipa:     0,
		Her:     0,
		Mort:    0,
		Result:  ""},
	{
		Id:      37,
		Empresa: "EMPRESA DE TRANSPORTES HerOES DEL PACIFICO S.R.L.",
		Ipa:     0,
		Her:     0,
		Mort:    0,
		Result:  ""},
	{
		Id:      38,
		Empresa: "EXPRESO LOS CHANKAS S.R.L.",
		Ipa:     0,
		Her:     0,
		Mort:    0,
		Result:  ""},
	{
		Id:      39,
		Empresa: "EMPRESA TURISMO ATAHUALPA SERVICIOS GENERALES S.C.R.L.",
		Ipa:     0,
		Her:     0,
		Mort:    0, Result: ""},
	{
		Id:      40,
		Empresa: "EXPRESO MOLINA UNION S.A.C.",
		Ipa:     0,
		Her:     0,
		Mort:    0, Result: ""},
	{
		Id:      41,
		Empresa: "TORRES ASOCIADOS S.R.L.",
		Ipa:     0,
		Her:     0,
		Mort:    0, Result: ""},
	{
		Id:      42,
		Empresa: "EMPRESA DE TRANSPORTES DE PASAJEROS Y CARGA CAVASSA S.A.C.",
		Ipa:     0,
		Her:     0,
		Mort:    0, Result: ""},
	{
		Id:      43,
		Empresa: "SHALOM EXPRESS S.A.C.",
		Ipa:     0,
		Her:     0,
		Mort:    0, Result: ""},
	{
		Id:      44,
		Empresa: "ALLINBUS S.A.C.",
		Ipa:     0,
		Her:     0,
		Mort:    0, Result: ""},
	{
		Id:      45,
		Empresa: "EMPRESA DE TRANSPORTES KALE S.A.C.",
		Ipa:     0,
		Her:     0,
		Mort:    0, Result: ""},
	{
		Id:      46,
		Empresa: "EMPRESA DE TRANSPORTES Y TURISMO JUDITH E.I.R.L.",
		Ipa:     0,
		Her:     0,
		Mort:    0, Result: ""},
	{
		Id:      47,
		Empresa: "TRANSANI S.R.L.",
		Ipa:     0,
		Her:     0,
		Mort:    0, Result: ""},
	{
		Id:      48,
		Empresa: "AMERICA EXPRESS S.A.",
		Ipa:     0,
		Her:     0,
		Mort:    0, Result: ""},
	{
		Id:      49,
		Empresa: "EMPRESA DE TRANSPORTES Y REPRESENTACIONES TURISMO CENTRAL S.A.",
		Ipa:     0,
		Her:     0,
		Mort:    0, Result: ""},
	{
		Id:      50,
		Empresa: "EMPRESA DE TRANSPORTES TSP E.I.R.L.",
		Ipa:     0,
		Her:     0,
		Mort:    0, Result: ""},
}

func dats() { // Carga los datos
	go ipa()
	go her()
	go mort()
}

func mast() { // Calcula y compara promedios
	promH(extra)
	promIPA(extra)
	promM(extra)
}

func envResp(resp string, id empresa) empresa { // Responde 4 veces
	var i = 0
	for i < 4 {
		if i == 0 {
			resp = fmt.Sprintf("%f", arregloIPA[extra]) // Carga el IPA del identificador elegido
			id.Ipa, _ = strconv.ParseFloat(resp, 64)
		} else if i == 1 {
			resp = strconv.Itoa(arregloH[extra]) // Carga el indice de heridos del identificador elegido
			id.Her, _ = strconv.Atoi(resp)
		} else if i == 2 {
			resp = strconv.Itoa(arregloM[extra]) // Carga el indice de mortalidad del identificador elegido
			id.Mort, _ = strconv.Atoi(resp)
		} else if i == 3 {
			if g == 0 { // Carga si es o no recomendable la empresa prestaria
				resp = "La empresa es recomendable para su viaje, ya que su índice de participación en accidentes es menor al promedio, al igual que su índice de accidentes letales"
				id.Result = resp
			} else if g == 1 {
				resp = "La empresa podría ser recomendable para su viaje, ya que su índice de participación en accidentes no es muy superior al promedio"
				id.Result = resp
			} else if g == 2 {
				resp = "La empresa no es recomendable para su viaje debido a su gran índice de participación en accidentes"
				id.Result = resp
			}
		}

		i++
	}
	return id
}

func promIPA(extra int) { // Compara el promedio del IPA
	auxIPA = 1.9

	arr1 = arregloIPA[extra]
	if arregloIPA[extra] > auxIPA {
		grandei = 1 //		El indice de IPA elegido está por encima del promedio
	} else {
		grandei = 0 //		El indice de IPA elegido está dentro del promedio
	}
}

func promH(extra int) { // Compara el promedio de heridos
	auxH = 10

	arr2 = arregloH[extra]
	if arregloH[extra] > auxH {
		grandeH = 1 //		El indice de heridos elegido está por encima del promedio
	} else {
		grandeH = 0 //		El indice de heridos elegido está dentro del promedio
	}
}

func promM(extra int) { // Compara el promedio de mortalidad
	auxM = 2

	arr3 = arregloM[extra]
	if arregloM[extra] > auxM {
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

func Test(id empresa) empresa {
	dats() // Carga la información en los arreglos
	ind := id.Id
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
	id1 := envResp(resp, id) // Envía la respuesta
	return id1
}

func run(id empresa) empresa {
	id1 := Test(id)
	return id1
}
func getEmpresa(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)
	params := mux.Vars(r)
	w.Header().Set("Content-Type", "application/json")
	for _, item := range empresas {
		id, _ = strconv.Atoi(params["id"])
		if item.Id == id {
			id1 := run(item)
			item.Id = id1.Id
			item.Empresa = id1.Empresa
			item.Her = id1.Her
			item.Ipa = id1.Ipa
			item.Mort = id1.Mort
			item.Result = id1.Result
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	msgError := fmt.Sprint("Empresa no encontrada con Id", id)
	json.NewEncoder(w).Encode(msgError)
}

func getEmpresas(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(empresas)
}

func indexRoute(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)
	fmt.Fprintf(w, "Welcome to ")
}

func enableCors(w *http.ResponseWriter) {

	(*w).Header().Set("Access-Control-Allow-Origin", "*")
}

func main() {

	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/", indexRoute)
	router.HandleFunc("/empresas", getEmpresas).Methods("GET")
	router.HandleFunc("/empresas/{id}", getEmpresa).Methods("GET")
	log.Fatal(http.ListenAndServe(":9000", router))

}

func envNum(msg string) { // Envia msg
	con, err := net.Dial("tcp", "localhost:9000")
	if err != nil {
		fmt.Println("Falla al resolver la direccion de red", err.Error())
		os.Exit(1) // Finaliza el programa con error
	}

	fmt.Fprintln(con, msg) //Envia msg
}
