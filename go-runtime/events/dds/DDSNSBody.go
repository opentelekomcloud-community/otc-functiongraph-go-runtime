package dds

import (
	"fmt"
)

type DDSNSBody struct {
	DB   string `json:"db"`
	Coll string `json:"coll"`
}

func (b *DDSNSBody) String() string {
	return fmt.Sprintf(`DDSNSBody{
                                 db=%v,
                                 coll=%v
                               }`, b.DB, b.Coll)
}
