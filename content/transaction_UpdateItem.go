package content

import (
	"github.com/benpate/derp"
	"github.com/davecgh/go-spew/spew"
)

type UpdateItemTransaction struct {
	ItemID int                    `json:"itemId" form:"itemId"`
	Data   map[string]interface{} `json:"data"   form:"data"`
	Hash   string                 `json:"hash"   form:"hash"`
}

func (txn UpdateItemTransaction) Execute(content *Content) error {

	spew.Dump(txn)

	err := content.UpdateItem(txn.ItemID, txn.Data, txn.Hash)
	return derp.Wrap(err, "content.UpdateItemTransaction", "Error updating item")
}

func (txn UpdateItemTransaction) Description() string {
	return "Update Item"
}