package hubspot

const (
	ownerBasePath = "owners"
)

// OwnerService is an interface of owner endpoints of the HubSpot API.
// HubSpot owners store information about individuals.
// It can also be associated with other CRM objects such as deal and company.
// Reference: https://developers.hubspot.com/docs/api/crm/owners
type OwnerService interface {
	Get(ownerID string, owner interface{}, option *RequestQueryOption) (*ResponseResource, error)
	Create(owner interface{}) (*ResponseResource, error)
	Update(ownerID string, owner interface{}) (*ResponseResource, error)
	Delete(ownerID string) error
	AssociateAnotherObj(ownerID string, conf *AssociationConfig) (*ResponseResource, error)
}

// OwnerServiceOp handles communication with the product related methods of the HubSpot API.
type OwnerServiceOp struct {
	ownerPath string
	client    *Client
}

var _ OwnerService = (*OwnerServiceOp)(nil)

type Owner struct {
	Archived                *HsBool `json:"archived,omitempty"`
	CreatedAt               *HsTime `json:"createdAt,omitempty"`
	Email                   *HsStr  `json:"email,omitempty"`
	FirstName               *HsStr  `json:"firstName,omitempty"`
	Id                      *HsStr  `json:"id,omitempty"`
	LastName                *HsStr  `json:"lastName,omitempty"`
	Type                    *HsStr  `json:"type,omitempty"`
	UpdatedAt               *HsTime `json:"updatedAt,omitempty"`
	UserID                  *HsStr  `json:"userId,omitempty"`
	UserIdIncludingInactive *HsStr  `json:"userIdIncludingInactive,omitempty"`
}

var defaultOwnerFields = []string{
	"archived",
	"createdAt",
	"email",
	"firstName",
	"id",
	"lastName",
	"type",
	"updatedAt",
	"userId",
	"userIdIncludingInactive",
}

// Get gets a owner.
// In order to bind the get content, a structure must be specified as an argument.
// Also, if you want to gets a custom field, you need to specify the field name.
// If you specify a non-existent field, it will be ignored.
// e.g. &hubspot.RequestQueryOption{ Properties: []string{"custom_a", "custom_b"}}
func (s *OwnerServiceOp) Get(ownerID string, owner interface{}, option *RequestQueryOption) (*ResponseResource, error) {
	resource := &ResponseResource{Properties: owner}
	if err := s.client.Get(s.ownerPath+"/"+ownerID, resource, option.setupProperties(defaultOwnerFields)); err != nil {
		return nil, err
	}
	return resource, nil
}

// Create creates a new owner.
// In order to bind the created content, a structure must be specified as an argument.
// When using custom fields, please embed hubspot.Owner in your own structure.
func (s *OwnerServiceOp) Create(owner interface{}) (*ResponseResource, error) {
	req := &RequestPayload{Properties: owner}
	resource := &ResponseResource{Properties: owner}
	if err := s.client.Post(s.ownerPath, req, resource); err != nil {
		return nil, err
	}
	return resource, nil
}

// Update updates a owner.
// In order to bind the updated content, a structure must be specified as an argument.
// When using custom fields, please embed hubspot.Owner in your own structure.
func (s *OwnerServiceOp) Update(ownerID string, owner interface{}) (*ResponseResource, error) {
	req := &RequestPayload{Properties: owner}
	resource := &ResponseResource{Properties: owner}
	if err := s.client.Patch(s.ownerPath+"/"+ownerID, req, resource); err != nil {
		return nil, err
	}
	return resource, nil
}

// Delete deletes a owner.
func (s *OwnerServiceOp) Delete(ownerID string) error {
	return s.client.Delete(s.ownerPath+"/"+ownerID, nil)
}

// AssociateAnotherObj associates Owner with another HubSpot objects.
// If you want to associate a custom object, please use a defined value in HubSpot.
func (s *OwnerServiceOp) AssociateAnotherObj(ownerID string, conf *AssociationConfig) (*ResponseResource, error) {
	resource := &ResponseResource{Properties: &Owner{}}
	if err := s.client.Put(s.ownerPath+"/"+ownerID+"/"+conf.makeAssociationPath(), nil, resource); err != nil {
		return nil, err
	}
	return resource, nil
}
