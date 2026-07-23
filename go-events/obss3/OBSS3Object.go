package obs

import (
	"fmt"
)

type OBSS3Object struct {
	// object-key
	Key string `json:"key"`

	// object eTag
	ETag string `json:"etag"`

	// object-size in bytes
	Size int64 `json:"size"`

	// version id
	VersionId string `json:"versionId"`

	// a string representation of a hexadecimal value used to determine event sequence, only used with PUTs and DELETEs
	Sequencer string `json:"sequencer"`
}

func (b *OBSS3Object) String() string {
	return fmt.Sprintf(`OBSS3Object{
                                 key=%v,
                                 etag=%v,
                                 size=%v,
																 versionId=%v,
																 sequencer=%v
                               }`,
		b.Key,
		b.ETag,
		b.Size,
		b.VersionId,
		b.Sequencer)
}
