package models

type ResultadosHistorico struct {
	Colaborador                  ColaboradorInformacionCompleta
	ResultadoEvaluacionPorMeta40 []ResultadoEvaluacion
	ResultadoEvaluacionGeneral60 []ResultadoEvaluacion
}
type ResultadoEvaluacion struct {
	IdEvaluacion           int     `gorm:"column:idEvaluacion"`
	Anio                   int     `gorm:"column:Anio"`
	IdGrado                int     `gorm:"column:idGrado"`
	Logrado                int     `gorm:"column:Logrado"`
	CumplimientoPorcentaje float32 `gorm:"column:CumplimientoPorcentaje"`
}
