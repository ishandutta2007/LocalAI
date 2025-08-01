package localai

import (
	"github.com/gofiber/fiber/v2"
	"github.com/mudler/LocalAI/core/config"
	"github.com/mudler/LocalAI/core/gallery"
	"github.com/mudler/LocalAI/core/http/utils"
	"github.com/mudler/LocalAI/core/services"
	"github.com/mudler/LocalAI/internal"
	"github.com/mudler/LocalAI/pkg/model"
)

func WelcomeEndpoint(appConfig *config.ApplicationConfig,
	cl *config.BackendConfigLoader, ml *model.ModelLoader, opcache *services.OpCache) func(*fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		backendConfigs := cl.GetAllBackendConfigs()
		galleryConfigs := map[string]*gallery.ModelConfig{}

		for _, m := range backendConfigs {
			cfg, err := gallery.GetLocalModelConfiguration(ml.ModelPath, m.Name)
			if err != nil {
				continue
			}
			galleryConfigs[m.Name] = cfg
		}

		modelsWithoutConfig, _ := services.ListModels(cl, ml, config.NoFilterFn, services.LOOSE_ONLY)

		// Get model statuses to display in the UI the operation in progress
		processingModels, taskTypes := opcache.GetStatus()

		summary := fiber.Map{
			"Title":             "LocalAI API - " + internal.PrintableVersion(),
			"Version":           internal.PrintableVersion(),
			"BaseURL":           utils.BaseURL(c),
			"Models":            modelsWithoutConfig,
			"ModelsConfig":      backendConfigs,
			"GalleryConfig":     galleryConfigs,
			"ApplicationConfig": appConfig,
			"ProcessingModels":  processingModels,
			"TaskTypes":         taskTypes,
		}

		if string(c.Context().Request.Header.ContentType()) == "application/json" || len(c.Accepts("html")) == 0 {
			// The client expects a JSON response
			return c.Status(fiber.StatusOK).JSON(summary)
		} else {
			// Render index
			return c.Render("views/index", summary)
		}
	}
}
