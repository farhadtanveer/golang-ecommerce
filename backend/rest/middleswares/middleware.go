package middleware

import "ecommerce/config"

type Middleswares struct {
	cnf *config.Config
}

func NewMiddlewares(cnf *config.Config) *Middleswares {
	return &Middleswares{cnf: cnf}
}