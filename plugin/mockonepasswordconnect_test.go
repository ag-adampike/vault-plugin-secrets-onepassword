package onepassword

import (
	"github.com/1Password/connect-sdk-go/onepassword"
)

type TestClient struct {
	GetVaultsFunc      func() ([]onepassword.Vault, error)
	GetItemFunc        func(uuid string, vaultUUID string) (*onepassword.Item, error)
	GetItemsFunc       func(vaultUUID string) ([]onepassword.Item, error)
	GetItemByTitleFunc func(title string, vaultUUID string) (*onepassword.Item, error)
	CreateItemFunc     func(item *onepassword.Item, vaultUUID string) (*onepassword.Item, error)
	UpdateItemFunc     func(item *onepassword.Item, vaultUUID string) (*onepassword.Item, error)
	DeleteItemFunc     func(item *onepassword.Item, vaultUUID string) error
}

var (
	DoGetVaultsFunc      func() ([]onepassword.Vault, error)
	DoGetItemFunc        func(uuid string, vaultUUID string) (*onepassword.Item, error)
	DoGetItemByTitleFunc func(title string, vaultUUID string) (*onepassword.Item, error)
	DoCreateItemFunc     func(item *onepassword.Item, vaultUUID string) (*onepassword.Item, error)
	DoDeleteItemFunc     func(item *onepassword.Item, vaultUUID string) error
	DoGetItemsFunc       func(vaultUUID string) ([]onepassword.Item, error)
	DoUpdateItemFunc     func(item *onepassword.Item, vaultUUID string) (*onepassword.Item, error)
)

// Do is the mock client's `Do` func
func (m *TestClient) GetVaults() ([]onepassword.Vault, error) {
	return DoGetVaultsFunc()
}

func (m *TestClient) GetItem(uuid string, vaultUUID string) (*onepassword.Item, error) {
	return DoGetItemFunc(uuid, vaultUUID)
}

func (m *TestClient) GetItems(vaultUUID string) ([]onepassword.Item, error) {
	return DoGetItemsFunc(vaultUUID)
}

func (m *TestClient) GetItemByTitle(title string, vaultUUID string) (*onepassword.Item, error) {
	return DoGetItemByTitleFunc(title, vaultUUID)
}

func (m *TestClient) CreateItem(item *onepassword.Item, vaultUUID string) (*onepassword.Item, error) {
	return DoCreateItemFunc(item, vaultUUID)
}

func (m *TestClient) DeleteItem(item *onepassword.Item, vaultUUID string) error {
	return DoDeleteItemFunc(item, vaultUUID)
}

func (m *TestClient) UpdateItem(item *onepassword.Item, vaultUUID string) (*onepassword.Item, error) {
	return DoUpdateItemFunc(item, vaultUUID)
}