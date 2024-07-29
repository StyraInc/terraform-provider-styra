// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

import (
	"github.com/StyraInc/terraform-provider-styra/internal/sdk/internal/utils"
)

type SystemsV1BundleDownloadConfig struct {
	// enabled delta bundles on bundle download
	DeltaBundles *bool `default:"false" json:"delta_bundles"`
}

func (s SystemsV1BundleDownloadConfig) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(s, "", false)
}

func (s *SystemsV1BundleDownloadConfig) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &s, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *SystemsV1BundleDownloadConfig) GetDeltaBundles() *bool {
	if o == nil {
		return nil
	}
	return o.DeltaBundles
}
