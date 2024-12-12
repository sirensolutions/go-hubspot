package hubspot

import "fmt"

const (
	crmBasePath = "crm"

	objectsBasePath = "objects"
)

type CRM struct {
	Contact    ContactService
	Company    CompanyService
	Deal       DealService
	Owners     CrmOwnersService
	Imports    CrmImportsService
	Schemas    CrmSchemasService
	Properties CrmPropertiesService
	Tickets    CrmTicketsServivce
}

func newCRM(c *Client) *CRM {
	crmPath := fmt.Sprintf("%s/%s", crmBasePath, c.apiVersion)
	return &CRM{
		Contact: &ContactServiceOp{
			contactPath: fmt.Sprintf("%s/%s/%s", crmPath, objectsBasePath, contactBasePath),
			client:      c,
		},
		Company: &CompanyServiceOp{
			companyPath: fmt.Sprintf("%s/%s/%s", crmPath, objectsBasePath, companyBasePath),
			client:      c,
		},
		Deal: &DealServiceOp{
			dealPath: fmt.Sprintf("%s/%s/%s", crmPath, objectsBasePath, dealBasePath),
			client:   c,
		},
		Owners: &CrmOwnersServiceOp{
			crmOwnersPath: fmt.Sprintf("%s/%s", crmPath, crmOwnersPath),
			client:        c,
		},
		Imports: &CrmImportsServiceOp{
			crmImportsPath: fmt.Sprintf("%s/%s", crmPath, crmImportsBasePath),
			client:         c,
		},
		Schemas: &CrmSchemasServiceOp{
			crmSchemasPath: fmt.Sprintf("%s/%s", crmPath, crmSchemasPath),
			client:         c,
		},
		Properties: &CrmPropertiesServiceOp{
			crmPropertiesPath: fmt.Sprintf("%s/%s", crmPath, crmPropertiesPath),
			client:            c,
		},
		Tickets: &CrmTicketsServivceOp{
			crmTicketsPath: fmt.Sprintf("%s/%s/%s", crmPath, objectsBasePath, crmTicketsBasePath),
			client:         c,
		},
	}
}
