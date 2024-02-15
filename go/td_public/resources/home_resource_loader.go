package resources

import (
	"encoding/json"
	"os"

	"github.com/Tracking-Detector/td-backend/go/td_public/model"
	"github.com/a-h/templ"
)

const (
	heroResourcePath              = "./assets/content/hero.json"
	featuresResourcePath          = "./assets/content/features.json"
	productsResourcePath          = "./assets/content/products.json"
	installationGuideResourcePath = "./assets/content/installation_guide.json"
	contactResourcePath           = "./assets/content/contact.json"
)

func LoadHomeResource() *model.Home {
	home := model.NewHome()
	hero := &model.Hero{}
	loadResource(heroResourcePath, hero)
	home.Hero = hero
	features := &model.Features{}
	loadResource(featuresResourcePath, features)
	home.Features = features
	products := &model.Products{}
	loadResource(productsResourcePath, products)
	home.Products = products
	installationGuide := &model.InstallationGuide{}
	loadResource(installationGuideResourcePath, installationGuide)
	home.InstallationGuide = installationGuide
	contact := &model.Contact{}
	loadResource(contactResourcePath, contact)
	home.Contact = contact
	home.Navbar = &model.Navbar{
		Links: []model.Link{
			{
				Text: "Features",
				URL:  templ.SafeURL("#" + home.Features.ID),
			},
			{
				Text: "Products",
				URL:  templ.SafeURL("#" + home.Products.ID),
			},
			{
				Text: "Install",
				URL:  templ.SafeURL("#" + home.InstallationGuide.ID),
			},
			{
				Text: "Contact",
				URL:  templ.SafeURL("#" + home.Contact.ID),
			},
		},
	}
	return home
}

func loadResource(filePath string, resource interface{}) {
	// Read file path and map to model.Hero
	file, err := os.Open(filePath)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	err = json.NewDecoder(file).Decode(resource)
	if err != nil {
		panic(err)
	}
}
