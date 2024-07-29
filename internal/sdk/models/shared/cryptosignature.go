// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

type CryptoSignature struct {
	Excluded   any `json:"excluded,omitempty"`
	Signatures any `json:"signatures,omitempty"`
}

func (o *CryptoSignature) GetExcluded() any {
	if o == nil {
		return nil
	}
	return o.Excluded
}

func (o *CryptoSignature) GetSignatures() any {
	if o == nil {
		return nil
	}
	return o.Signatures
}
