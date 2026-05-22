package merchants

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	httpServer "github.com/xnxq1/mini-ledger-go/internal/http-server"
)

type MerchantRouter struct {
}

func (route *MerchantRouter) CreateMerchant(w http.ResponseWriter, r *http.Request) {
	var request CreateMerchantRequest
	err := httpServer.DecodeJson(r.Body, &request)
	httpServer.SetError(r, err)
}

func (route *MerchantRouter) GetMerchants(w http.ResponseWriter, r *http.Request) {
	e := json.NewEncoder(w)
	if err := e.Encode("Hello World"); err != nil {
		log.Fatal(err)
	}

}

func (route *MerchantRouter) GetMerchantsWithBalance(w http.ResponseWriter, r *http.Request) {

}

func (route *MerchantRouter) GetMerchantBalances(w http.ResponseWriter, r *http.Request) {

}
func (route *MerchantRouter) CreateBalance(w http.ResponseWriter, r *http.Request) {

}

func (route *MerchantRouter) Init() *chi.Mux {
	router := chi.NewRouter()
	router.Route("/merchants", func(r chi.Router) {
		r.Get("/", route.GetMerchants)
		r.Post("/", route.CreateMerchant)
		r.Get("/{merchant_name}", route.GetMerchantsWithBalance)
		r.Get("/{merchant_name}/balance", route.GetMerchantBalances)
		r.Post("/balance", route.CreateBalance)
	})
	return router
}
