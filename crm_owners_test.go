package hubspot

import (
	"fmt"
	"os"
	"testing"
	"time"
)

func TestListCrmOwners(t *testing.T) {
	t.SkipNow()
	cli, _ := NewClient(SetPrivateAppToken(os.Getenv("PRIVATE_APP_TOKEN")))
	// Use crm_schemas:TestCreate() to generate this...
	res, err := cli.CRM.Owners.List()
	if err != nil {
		t.Error(err)
	}

	if len(res.Results) < 1 {
		t.Error("expected len(res.Results) to be > 1")
	}
}

func TestGetCrmOwner(t *testing.T) {
	t.SkipNow()

	cli, _ := NewClient(SetPrivateAppToken(os.Getenv("PRIVATE_APP_TOKEN")))
	// Use crm_schemas:TestCreate() to generate this...
	res, err := cli.CRM.Owners.Get("12345678")
	if err != nil {
		t.Error(err)
	}
	if *res.Id != "12345678" {
		t.Errorf("expected res.Id to be 12345678, got %s", res.Id)
	}
}

func TestCreateOwner(t *testing.T) {
	t.SkipNow()
	cli, _ := NewClient(SetPrivateAppToken(os.Getenv("PRIVATE_APP_TOKEN")))
	newOwner := &CrmOwner{
		Archived:                NewBoolean(false),
		CreatedAt:               NewTime(time.Now()),
		Email:                   NewString("nobody@example.com"),
		FirstName:               NewString("Mr"),
		Id:                      NewString("12345678"),
		LastName:                NewString("Nobody"),
		Type:                    NewString("PERSON"),
		UpdatedAt:               NewTime(time.Now()),
		UserID:                  NewInt(0),
		UserIdIncludingInactive: NewInt(0),
	}

	_, err := cli.CRM.Owners.Create(newOwner)
	if err != nil {
		t.Error(err)
		return
	}
}

func TestUpdateOwner(t *testing.T) {
	t.SkipNow()
	cli, _ := NewClient(SetPrivateAppToken(os.Getenv("PRIVATE_APP_TOKEN")))

	updateOwner := make(map[string]interface{})
	updateOwner["label"] = fmt.Sprintf("Updated Label %s", time.Now().String())

	res, err := cli.CRM.Owners.Update("12345678", &updateOwner)
	if err != nil {
		t.Error(err)
		return
	}

	if res.Id != updateOwner["label"] {
		t.Errorf("expected res.Label to be %s, got %s", updateOwner["label"], res.Id)
	}
}

func TestDeleteOwner(t *testing.T) {
	t.SkipNow()
	cli, _ := NewClient(SetPrivateAppToken(os.Getenv("PRIVATE_APP_TOKEN")))
	err := cli.CRM.Owners.Delete("12345678")
	if err != nil {
		t.Error(err)
	}
}
