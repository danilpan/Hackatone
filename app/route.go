package app

func (a app) Routes() {
	v1 := a.fwk.Group("/api/v1")

	v1.GET("/establishment-types", a.hnd.GetEstablishmentTypes)
	v1.GET("/establishments", a.hnd.GetEstablishments)
}
