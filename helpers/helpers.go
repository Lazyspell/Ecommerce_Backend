package helpers

import (
	"github.com/lazyspell/Ecommerce_Backend/config"
)

var app *config.AppConfig

func NewHelpers(a *config.AppConfig) {
	app = a
}
