package models

type ResultadosHistorico struct {
	Colaborador ColaboradorInformacionCompleta
	Resultados  []ResultadoEvaluacion
}
type ResultadoEvaluacion struct {
	IdEvaluacion             int     `gorm:"column:idEvaluacion"`
	idEvaluacionAnual        int     `gorm:"column:idEvaluacionAnual"`
	Anio                     int     `gorm:"column:Anio"`
	CumplimientoPorcentaje60 float32 `gorm:"column:CumplimientoPorcentaje60"`
	CumplimientoEnBaseA60    float32 `gorm:"column:CumplimientoEnBaseA60"`
	CumplimientoPorcentaje40 float32 `gorm:"column:CumplimientoPorcentaje40"`
	CumplimientoEnBaseA40    float32 `gorm:"column:CumplimientoEnBaseA40"`
}
