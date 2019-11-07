package routes

import (
	"github.com/gorilla/mux"
)

func Routes() *mux.Router {

	myRouter := mux.NewRouter()

	myRouter.HandleFunc("/prueba", PruebaHandler).Methods("GET")
	myRouter.HandleFunc("/login", LoginHandler)
	myRouter.HandleFunc("/GetColaboradoresEquipo/{idColaborador}", GetColaboradoresEquipo).Methods("GET")
	// myRouter.HandleFunc("/GetMetasColaborador/{codigoEmpleado}", GetMetasColaborador).Methods("GET")
	// myRouter.HandleFunc("/GetMedidasPredictivas/{codigoEmpleado}", GetMedidasPredictivas).Methods("GET")
	// myRouter.HandleFunc("/GetFrecuencias", GetFrecuencias)
	// myRouter.HandleFunc("/GetMediciones", GetMediciones)
	// myRouter.HandleFunc("/GetResultados/{idMP}/{mes}", GetResultados).Methods("GET")
	// myRouter.HandleFunc("/GetResultadosMCI/{idMCI}", GetResultadosMCI).Methods("GET")
	// myRouter.HandleFunc("/PeriodosPorMPNew", DBPeriodosPorMPInsert).Methods("POST")
	// myRouter.HandleFunc("/PeriodosPorMPUpdate", DBPeriodosPorMPUpdateHandler).Methods("POST")
	// myRouter.HandleFunc("/ResultadosUpdate", ResultadosUpdate).Methods("POST")
	// myRouter.HandleFunc("/ResultadoMCIAutorizar/{idResultado}", ResultadoMCIAutorizar).Methods("GET")
	// myRouter.HandleFunc("/ValorMCIAdd", ValorMCIAdd).Methods("POST")
	// myRouter.HandleFunc("/MetaMCIAdd", MetaMCIAdd).Methods("POST")
	// myRouter.HandleFunc("/TipoGraficos", TipoGraficosHandler).Methods("GET")
	// myRouter.HandleFunc("/GetPeriodicidadMCI", GetPeriodicidadMCI)
	// myRouter.HandleFunc("/GetResultadosGraficaMCI/{IdMCI}/{Anio}", GetResultadosGraficaMCI).Methods("GET")
	// myRouter.HandleFunc("/GraficoPorMCINew", GraficoPorMCINewHandler).Methods("POST")
	// myRouter.HandleFunc("/SubAreas", SubAreas).Methods("GET")
	// myRouter.HandleFunc("/ColaboradoresPorArea/{AreaId}", ColaboradoresPorArea).Methods("GET")
	// myRouter.HandleFunc("/PeriodosPorMCIAdd", PeriodosPorMCIAdd).Methods("POST")
	// myRouter.HandleFunc("/ResultadosMCICreate/{idMCI}", ResultadosMCICreate).Methods("GET")
	// myRouter.HandleFunc("/GetGraficoColaborador/{codigoEmpleado}", GetGraficoColaborador).Methods("GET")
	// myRouter.HandleFunc("/TableroColaborador/{codigoEmpleado}", TableroColaborador).Methods("GET")
	// myRouter.HandleFunc("/TableroColaborador/{codigoEmpleado}/{mesId}", TableroColaborador).Methods("GET")
	// myRouter.HandleFunc("/AutorizarResultado", AutorizarResultadoHandler).Methods("POST")
	// myRouter.HandleFunc("/GetMCIColaborador/{codigoEmpleado}", GetMCI).Methods("GET")
	// myRouter.HandleFunc("/GetColaboradoresSubArea/{idSubArea}", GetColaboradoresSubArea).Methods("GET")
	// myRouter.HandleFunc("/GetColaboradoresAdmins", GetColaboradoresAdmins).Methods("GET")
	// myRouter.HandleFunc("/BrujulaPorMPAdd", BrujulaPorMPCreate).Methods("POST")
	// myRouter.HandleFunc("/BrujulaPorMPUpdate", BrujulaPorMPUpdate).Methods("POST")
	// myRouter.HandleFunc("/BrujulasPorMP/{codigoEmpleado}/{idResultado}", BrujulasPorMPGet).Methods("GET")
	// myRouter.HandleFunc("/BrujulaEstados", BrujulaEstadosGet).Methods("GET")
	// myRouter.HandleFunc("/BrujulaActividadesPorMP/{idMP}/{mes}", BrujulaActividadesPorMP).Methods("GET")
	// myRouter.HandleFunc("/BrujulaActividadesPorColaborador/{codigoEmpleado}", BrujulaActividadesPorColaborador).Methods("GET")
	// myRouter.HandleFunc("/BrujulaActividadesPorColaborador/{codigoEmpleado}/{idEstado}", BrujulaActividadesPorColaborador).Methods("GET")
	// myRouter.HandleFunc("/BrujulaActividadesPorColaborador/{codigoEmpleado}/{idEstado}/{esLider}", BrujulaActividadesPorColaborador).Methods("GET")
	// myRouter.HandleFunc("/RegistrarEventoDelSistema", RegistrarEventoDelSistema).Methods("POST")
	// myRouter.HandleFunc("/SendEmail", SendEmailHandler).Methods("POST")
	// myRouter.HandleFunc("/SendEmailConfirmedAutorization/{empleadoCodigo}", SendEmailHandler).Methods("POST")
	// myRouter.HandleFunc("/ReunionMCINew", ReunionMCICreate).Methods("POST")
	// myRouter.HandleFunc("/DetalleReunionMCINew", DetalleReunionMCICreate).Methods("POST")
	// myRouter.HandleFunc("/GetReunionDelDia/{IdLider}", GetReunionDelDia).Methods("GET")
	// myRouter.HandleFunc("/UpdateTiempoReunion", UpdateTiempoReunion).Methods("POST")

	return myRouter
}
