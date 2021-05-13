package application

import (
	"github.com/amiremperor/bookStore-api/internal/config"
	"github.com/amiremperor/bookStore-api/pkg/translator/i18n"
)

func Run(cfg *config.Config) error {
	translator , err := i18n.New(cfg.I18n.BundlePath)
	if err != nil {
		return err
	}

	_ = translator
	return err
}