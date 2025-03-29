package main

type ProblemInfo struct {
	VolumenInicial uint8
	MaximaCapacidadBombeo uint8
	CosteDecisionesPorUnidad uint8
	CosteAlmacenamientoPorUnidad float64
	CosteDemandaNoSatisfechaPorUnidad uint8
}


func GetDefaultProblemContext() ProblemInfo {
	return ProblemInfo{
	VolumenInicial : 1,
	MaximaCapacidadBombeo : 2,
	CosteDecisionesPorUnidad : 1,
	CosteAlmacenamientoPorUnidad : 0.1,
	CosteDemandaNoSatisfechaPorUnidad : 2,
}
}

func optimize(info *ProblemInfo, week *[]uint8) {

}

func main() {

}
