package product

import (
	"ecommerce/domain"
	"ecommerce/util"
	"net/http"
	"strconv"
)

type Pagination struct {
	Data []*domain.Product  `json:"data"`
	Limit int64 			`json:"limit"`
	Page  int64 			`json:"page"`
	Totalitems int64 			`json:"totalItems"`
	TotalPages int64 			`json:"totalPages"`
}

func (h *Handler) GetProducts(w http.ResponseWriter, r *http.Request) {
	// get query parameters
	// get page from query params
	// get limit from query params
	reqQuery := r.URL.Query()
	pageAsString := reqQuery.Get("page")
	limitAsString := reqQuery.Get("limit")

	page, _ := strconv.ParseInt(pageAsString, 10, 32)
	limit,_ := strconv.ParseInt(limitAsString, 10, 32)

	if page <= 0 {
		page = 1
	}

	if limit <= 0 {
		limit = 10
	}


	productList, err := h.svc.List(page, limit)
	if err != nil {
		http.Error(w, "Error fetching products", http.StatusInternalServerError)
		return
	}

	cnt, err := h.svc.Count()
	if err != nil {
		http.Error(w, "Error counting products", http.StatusInternalServerError)
		return
	}

	paginatedData := Pagination {
		Data: productList,
		Limit: limit,
		Page: page,
		Totalitems: cnt,
		TotalPages: cnt / limit,
	}

	util.SendData(w, http.StatusOK, paginatedData) // Send the product list as JSON response
}