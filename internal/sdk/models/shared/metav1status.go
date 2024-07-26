// Code generated by Speakeasy (https://speakeasyapi.com). DO NOT EDIT.

package shared

import (
	"github.com/StyraInc/terraform-provider-styra/internal/sdk/internal/utils"
	"time"
)

type MetaV1Status struct {
	Code      string    `json:"code"`
	Message   string    `json:"message"`
	Timestamp time.Time `json:"timestamp"`
}

func (m MetaV1Status) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(m, "", false)
}

func (m *MetaV1Status) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &m, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *MetaV1Status) GetCode() string {
	if o == nil {
		return ""
	}
	return o.Code
}

func (o *MetaV1Status) GetMessage() string {
	if o == nil {
		return ""
	}
	return o.Message
}

func (o *MetaV1Status) GetTimestamp() time.Time {
	if o == nil {
		return time.Time{}
	}
	return o.Timestamp
}
