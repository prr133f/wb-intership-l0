package v1

func (r *Route) Data() {
	(*r.Group).Get("/dataById/:id", r.View.GetDataByID)
}
