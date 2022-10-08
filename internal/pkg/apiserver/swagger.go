package apiserver

import (
	"github.com/go-openapi/spec"
)

func enrichSwaggerObject(swo *spec.Swagger) {
	swo.Info = &spec.Info{
		InfoProps: spec.InfoProps{
			Title:       "Mariner ApiServer",
			Description: "Mariner Daemon ApiServer",
			Contact: &spec.ContactInfo{
				ContactInfoProps: spec.ContactInfoProps{
					Name:  "YinYongYou",
					Email: "yinyongyou@gmail.com",
					URL:   "https://github.com/mariner-group/marinerd",
				},
			},
			License: &spec.License{
				LicenseProps: spec.LicenseProps{
					Name: "Apache 2.0",
					URL:  "http://www.apache.org/licenses/LICENSE-2.0.html",
				},
			},
			Version: "1.0.0",
		},
	}
	swo.Tags = []spec.Tag{
		{
			TagProps: spec.TagProps{
				Name:        "users",
				Description: "Managing users",
			},
		},
	}
	swo.SecurityDefinitions = map[string]*spec.SecurityScheme{
		"jwt": spec.APIKeyAuth("Authorization", "header"),
	}
	enrichSwaggerObjectSecurity(swo)
}

func enrichSwaggerObjectSecurity(swo *spec.Swagger) {
	swo.Security = []map[string][]string{
		{
			"jwt": {},
		},
	}
}
