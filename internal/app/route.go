package app

func (r *RESTApp) registerRoute() {
	// Stats
	r.echo.GET("/stats", r.statsDelivery.Get)

	v1 := r.echo.Group("v1")

	// Farm route
	v1.GET("/farm", r.farmDelivery.GetList)
	v1.GET("/farm/:id", r.farmDelivery.Get)
	v1.POST("/farm", r.farmDelivery.Create)
	v1.PUT("/farm/:id", r.farmDelivery.Update)
	v1.DELETE("/farm/:id", r.farmDelivery.Delete)

	// Pond route
	v1.GET("/pond", r.pondDelivery.GetList)
	v1.GET("/pond/:id", r.pondDelivery.Get)
	v1.POST("/pond", r.pondDelivery.Create)
	v1.PUT("/pond/:id", r.pondDelivery.Update)
	v1.DELETE("/pond/:id", r.pondDelivery.Delete)
}
