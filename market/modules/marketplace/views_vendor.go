package marketplace

import (
	"fmt"
	"strconv"

	"github.com/gocraft/web"
	// "github.com/go-playground/validator/v10"

	"argomarket/market/modules/util"
)

func (c *Context) ViewVendorDashboard(w web.ResponseWriter, r *web.Request) {
	fmt.Println("vendor_dashboard")

	c.ViewPackages = FindPackagesByVendorUuid(c.ViewUser.Uuid)
	util.RenderTemplate(w, "vendor_dashboard", c)
}

func (c *Context) ViewCreatePhysical(w web.ResponseWriter, r *web.Request) {
	c.Categories, _ = GetTopCategories()

	util.RenderTemplate(w, "create_physical", c)
}

func (c *Context) CreatePhysicalPost(w web.ResponseWriter, r *web.Request) {
	// type FormData struct {
	// 	Name string `json:"name" validate:"required"`
	// }

	totalAmountStr := r.FormValue("total_amount")
	totalAmount, _ := strconv.ParseFloat(totalAmountStr, 64)
	
	package_ := &Package{
		Uuid:             util.GenerateUuid(),
		Name:       r.FormValue("name"),
		Type:   "Physical",
		CategoryUuid: r.FormValue("category"),
		VendorUuid:         c.ViewUser.Uuid,
		Description:	r.FormValue("description"),
		RefundPolicy: r.FormValue("refund_policy"),
		TotalAmount: totalAmount,
		Measure: r.FormValue("measure"),
		PayMethod: r.FormValue("paymethod"),
		AllowedCrypto: r.FormValue("currency"),
		ShippingFrom: r.FormValue("ships_from"),
		ShippingTo: r.FormValue("ships_to"),
	}

	validationError := package_.Validate()
	if validationError != nil {
		return 
	}

	package_.Save()

	Price1, _ := strconv.ParseFloat(r.FormValue("price1"), 64)
	SubQuantity1, _ := strconv.ParseFloat(r.FormValue("subquantity1"), 64)
	package_price := &PackagePrice{
		Uuid:	util.GenerateUuid(),
		PackageUuid:	package_.Uuid,
		Price:  Price1,
		Amount:	SubQuantity1,
	}

	validationError = package_price.Validate()
	if validationError != nil {
		package_.Remove()
		return 
	}

	package_price.Save()

	if r.FormValue("price2") != "" && r.FormValue("subquantity2") != "" {
		Price2, _ := strconv.ParseFloat(r.FormValue("price2"), 64)
		SubQuantity2, _ := strconv.ParseFloat(r.FormValue("subquantity2"), 64)
		package_price = &PackagePrice{
			Uuid:	util.GenerateUuid(),
			PackageUuid:	package_.Uuid,
			Price:  Price2,
			Amount:	SubQuantity2,
		}

		package_price.Save()
	}

	if r.FormValue("price3") != "" && r.FormValue("subquantity3") != "" {
		Price3, _ := strconv.ParseFloat(r.FormValue("price3"), 64)
		SubQuantity3, _ := strconv.ParseFloat(r.FormValue("subquantity3"), 64)
		package_price = &PackagePrice{
			Uuid:	util.GenerateUuid(),
			PackageUuid:	package_.Uuid,
			Price:  Price3,
			Amount:	SubQuantity3,
		}
		
		package_price.Save()
	}

	if r.FormValue("price4") != "" && r.FormValue("subquantity4") != "" {
		Price4, _ := strconv.ParseFloat(r.FormValue("price4"), 64)
		SubQuantity4, _ := strconv.ParseFloat(r.FormValue("subquantity4"), 64)
		package_price = &PackagePrice{
			Uuid:	util.GenerateUuid(),
			PackageUuid:	package_.Uuid,
			Price:  Price4,
			Amount:	SubQuantity4,
		}

		package_price.Save()
	}

	if r.FormValue("price5") != "" && r.FormValue("subquantity5") != "" {
		Price5, _ := strconv.ParseFloat(r.FormValue("price5"), 64)
		SubQuantity5, _ := strconv.ParseFloat(r.FormValue("subquantity5"), 64)
		package_price = &PackagePrice{
			Uuid:	util.GenerateUuid(),
			PackageUuid:	package_.Uuid,
			Price:  Price5,
			Amount:	SubQuantity5,
		}

		package_price.Save();
	}

	Days1, _ := strconv.ParseInt(r.FormValue("days1"), 10, 64)
	ShipesPrice1, _ := strconv.ParseFloat(r.FormValue("ships_price1"), 64)
	ships_method := &ShippingMethod{
		Uuid:	util.GenerateUuid(),
		PackageUuid:	package_.Uuid,
		Method:  r.FormValue("method1"),
		Days:  int(Days1),
		Price:  ShipesPrice1,
	}

	validationError = ships_method.Validate()
	if validationError != nil {
		package_.Remove()
		return 
	}

	ships_method.Save()

	if r.FormValue("method2") != "" && r.FormValue("days2") != "" && r.FormValue("ships_price2") != "" {
		Days2, _ := strconv.ParseInt(r.FormValue("days2"), 10, 64)
		ShipesPrice2, _ := strconv.ParseFloat(r.FormValue("ships_price2"), 64)
		ships_method = &ShippingMethod{
			Uuid:	util.GenerateUuid(),
			PackageUuid:	package_.Uuid,
			Method:  r.FormValue("method2"),
			Days:  int(Days2),
			Price:  ShipesPrice2,
		}

		ships_method.Save()
	}

	if r.FormValue("method3") != "" && r.FormValue("days3") != "" && r.FormValue("ships_price3") != "" {
		Days3, _ := strconv.ParseInt(r.FormValue("days3"), 10, 64)
		ShipesPrice3, _ := strconv.ParseFloat(r.FormValue("ships_price3"), 64)
		ships_method = &ShippingMethod{
			Uuid:	util.GenerateUuid(),
			PackageUuid:	package_.Uuid,
			Method:  r.FormValue("method3"),
			Days:  int(Days3),
			Price:  ShipesPrice3,
		}

		ships_method.Save()
	}

	util.RenderTemplate(w, "create_physical", c)
}

func (c *Context) ViewCreateDigital(w web.ResponseWriter, r *web.Request) {
	util.RenderTemplate(w, "create_digital", c)
}
