package controllers

import (
	"WallesterAssigment/dto"
	"WallesterAssigment/interfaces"
	"WallesterAssigment/models"
	"encoding/json"
	"github.com/go-chi/chi"
	"github.com/gorilla/schema"
	"github.com/monoculum/formam"
	"html/template"
	"net/http"
	"strconv"
)

type CustomerController struct {
	interfaces.ICustomerService
}

type ListView struct {
	Title   string
	Message string
	Error   string
}

func (c CustomerController) List(w http.ResponseWriter, r *http.Request) {
	err := r.URL.Query().Get("err")
	msg := r.URL.Query().Get("msg")
	data := &ListView{
		Title:   "Customers",
		Message: getMessage(msg),
		Error:   getMessage(err),
	}
	view := view("./templates/list.html")
	if err := view.Execute(w, data); err != nil {
		panic(err)
	}
}

type FormView struct {
	Title        string
	SubmitAction string
	DeleteAction string
	Customer     models.Customer
	Genders      map[string]string
	Errors       error
}

func (c CustomerController) Create(w http.ResponseWriter, r *http.Request) {
	data := &FormView{
		Title:        "Add new customer",
		SubmitAction: "/create",
		Genders:      getGenders(),
	}
	view := view("./templates/customer_form.html")
	if err := view.Execute(w, data); err != nil {
		panic(err)
	}
}

func (c CustomerController) Store(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	input := dto.CreateCustomerRequest{}
	dec := formam.NewDecoder(&formam.DecoderOptions{TagName: "formam"})
	err := dec.Decode(r.Form, &input)
	if err != nil {
		panic(err)
	}
	_, err = c.ICustomerService.Create(r.Context(), input)
	if err != nil {
		data := &FormView{
			Title:    "Add new customer",
			Errors:   err,
			Customer: input.ToEntity(),
			Genders:  getGenders(),
		}
		if err := view("./templates/customer_form.html").Execute(w, data); err != nil {
			panic(err)
		}
	}
	http.Redirect(w, r, "/?msg=create_ok", http.StatusSeeOther)
}

func (c CustomerController) Edit(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	customerId, err := strconv.Atoi(id)
	if err != nil {
		http.Redirect(w, r, "/?err=not_found", http.StatusSeeOther)
		return
	}

	customer, err := c.ICustomerService.GetById(customerId)
	if err != nil {
		http.Redirect(w, r, "/?err=not_found", http.StatusSeeOther)
		return
	}
	data := &FormView{
		Title:        "Customer details",
		SubmitAction: "/edit/" + id,
		DeleteAction: "/delete/" + id,
		Customer:     customer,
		Genders:      getGenders(),
	}
	if err := view("./templates/customer_form.html").Execute(w, data); err != nil {
		panic(err)
	}

}

func (c CustomerController) Update(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	customerId, err := strconv.Atoi(id)
	if err != nil {
		http.Redirect(w, r, "/?err=not_found", http.StatusSeeOther)
		return
	}
	r.ParseForm()

	input := dto.CreateCustomerRequest{}
	dec := formam.NewDecoder(&formam.DecoderOptions{TagName: "formam"})
	err = dec.Decode(r.Form, &input)
	if err != nil {
		panic(err)
	}
	_, err = c.ICustomerService.Update(r.Context(), customerId, input)
	if err != nil {
		customer, _ := c.ICustomerService.GetById(customerId)
		data := &FormView{
			Title:    "Customer details",
			Errors:   err,
			Customer: customer,
			Genders:  getGenders(),
		}
		if err := view("./templates/customer_form.html").Execute(w, data); err != nil {
			panic(err)
		}
	}

	http.Redirect(w, r, "/?msg=update_ok", http.StatusSeeOther)
}

func (c CustomerController) Delete(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	customerId, err := strconv.Atoi(id)
	if err != nil {
		http.Redirect(w, r, "/?err=not_found", http.StatusSeeOther)
		return
	}
	r.ParseForm()

	input := dto.CreateCustomerRequest{}
	dec := formam.NewDecoder(&formam.DecoderOptions{TagName: "formam"})
	err = dec.Decode(r.Form, &input)
	if err != nil {
		panic(err)
	}
	err = c.ICustomerService.Delete(r.Context(), customerId, input.Hash)
	if err != nil {
		http.Redirect(w, r, "/?err=delete_failed", http.StatusSeeOther)
		return
	}
	http.Redirect(w, r, "/?msg=delete_ok", http.StatusSeeOther)

}

func (c CustomerController) Search(w http.ResponseWriter, r *http.Request) {
	var search dto.SearchRequest
	schema.NewDecoder().Decode(&search, r.URL.Query())
	search = search.SetOrderDirection()
	search = search.SetOrderColumn(getColumns())
	customers, pagination, err := c.ICustomerService.Search(r.Context(), search)
	if err != nil {
		panic(err)
	}
	response := make(map[string]interface{})
	response["draw"] = search.Draw
	response["recordsTotal"] = pagination.Total
	response["recordsFiltered"] = pagination.Filtered
	response["data"] = customers

	w.Header().Set("Content-Tyoe", "application/json")
	json.NewEncoder(w).Encode(response)

}

func view(file string) *template.Template {
	layout := template.Must(template.ParseFiles("./templates/layout.html"))
	files, err := layout.ParseFiles(file)
	if err != nil {
		panic(err)
	}
	return files
}

func getMessage(code string) string {
	switch code {
	case "not_found":
		return "Customer not found"
	case "update_ok":
		return "Customer updated"
	case "create_ok":
		return "Customer created"
	case "delete_ok":
		return "Customer deleted"
	case "delete_failed":
		return "Customer delete failed"
	default:
		return ""
	}
}

func getGenders() map[string]string {
	var genders = make(map[string]string)
	genders["Male"] = "Male"
	genders["Female"] = "Female"
	return genders
}

func getColumns() map[int]string {
	var columns = make(map[int]string)
	columns[0] = "lastname"
	columns[1] = "firstname"
	columns[2] = "gender"
	columns[3] = "email"
	return columns
}
