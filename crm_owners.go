package hubspot

import (
	"fmt"
)

const (
	crmOwnersPath = "owners"
)

type CrmOwnersList struct {
	Results []*CrmOwner `json:"results,omitempty"`
}

type CrmOwner struct {
	Archived                *HsBool         `json:"archived,omitempty"`
	CreatedAt               *HsTime         `json:"createdAt,omitempty"`
	Email                   *HsStr          `json:"email,omitempty"`
	FirstName               *HsStr          `json:"firstName,omitempty"`
	Id                      *HsStr          `json:"id,omitempty"`
	LastName                *HsStr          `json:"lastName,omitempty"`
	Type                    *HsStr          `json:"type,omitempty"`
	UpdatedAt               *HsTime         `json:"updatedAt,omitempty"`
	UserID                  *HsStr          `json:"userId,omitempty"`
	UserIdIncludingInactive *HsStr          `json:"userIdIncludingInactive,omitempty"`
	Teams                   *[]CrmOwnerTeam `json:"teams,omitempty"`
}

type CrmOwnerModificationMeta struct {
	Archivable       *HsBool `json:"archivable,omitempty"`
	ReadOnlyDefition *HsBool `json:"readOnlyDefinition,omitempty"`
	ReadOnlyValue    *HsBool `json:"readOnlyValue,omitempty"`
	ReadOnlyOptions  *HsBool `json:"readOnlyOptions,omitempty"`
}

type CrmOwnerTeam struct {
	Id      *HsStr  `json:"id,omitempty"`
	Name    *HsStr  `json:"name,omitempty"`
	Primary *HsBool `json:"primary,omitempty"`
}

// CrmOwnersService is an interface of CRM owners endpoints of the HubSpot API.
// Reference: https://developers.hubspot.com/docs/api/crm/owners
type CrmOwnersService interface {
	List(objectType string) (*CrmOwnersList, error)
	Create(objectType string, reqData interface{}) (*CrmOwner, error)
	Get(objectType string, ownerName string) (*CrmOwner, error)
	Delete(objectType string, ownerName string) error
	Update(objectType string, ownerName string, reqData interface{}) (*CrmOwner, error)
}

// CrmOwnersServiceOp handles communication with the CRM owners endpoint.
type CrmOwnersServiceOp struct {
	client        *Client
	crmOwnersPath string
}

var _ CrmOwnersService = (*CrmOwnersServiceOp)(nil)

func (s *CrmOwnersServiceOp) List(objectType string) (*CrmOwnersList, error) {
	var resource CrmOwnersList
	path := fmt.Sprintf("%s/%s", s.crmOwnersPath, objectType)
	if err := s.client.Get(path, &resource, nil); err != nil {
		return nil, err
	}
	return &resource, nil
}

func (s *CrmOwnersServiceOp) Get(objectType, ownerName string) (*CrmOwner, error) {
	var resource CrmOwner
	path := fmt.Sprintf("%s/%s/%s", s.crmOwnersPath, objectType, ownerName)
	if err := s.client.Get(path, &resource, nil); err != nil {
		return nil, err
	}
	return &resource, nil
}

func (s *CrmOwnersServiceOp) Create(objectType string, reqData interface{}) (*CrmOwner, error) {
	var resource CrmOwner
	path := fmt.Sprintf("%s/%s", s.crmOwnersPath, objectType)
	if err := s.client.Post(path, reqData, &resource); err != nil {
		return nil, err
	}
	return &resource, nil
}

func (s *CrmOwnersServiceOp) Delete(objectType string, ownerName string) error {
	path := fmt.Sprintf("%s/%s/%s", s.crmOwnersPath, objectType, ownerName)
	return s.client.Delete(path, nil)
}

func (s *CrmOwnersServiceOp) Update(objectType string, ownerName string, reqData interface{}) (*CrmOwner, error) {
	var resource CrmOwner
	path := fmt.Sprintf("%s/%s/%s", s.crmOwnersPath, objectType, ownerName)
	if err := s.client.Patch(path, reqData, &resource); err != nil {
		return nil, err
	}
	return &resource, nil
}
