package verifysignature

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// VerifySignatureRequestBody provides operations to call the verifySignature method.
type VerifySignatureRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // The digest property
    digest []byte
    // The signature property
    signature []byte
    // The signingKeyId property
    signingKeyId *string
}
// NewVerifySignatureRequestBody instantiates a new verifySignatureRequestBody and sets the default values.
func NewVerifySignatureRequestBody()(*VerifySignatureRequestBody) {
    m := &VerifySignatureRequestBody{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateVerifySignatureRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateVerifySignatureRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewVerifySignatureRequestBody(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *VerifySignatureRequestBody) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetDigest gets the digest property value. The digest property
func (m *VerifySignatureRequestBody) GetDigest()([]byte) {
    if m == nil {
        return nil
    } else {
        return m.digest
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *VerifySignatureRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["digest"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetByteArrayValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDigest(val)
        }
        return nil
    }
    res["signature"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetByteArrayValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSignature(val)
        }
        return nil
    }
    res["signingKeyId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSigningKeyId(val)
        }
        return nil
    }
    return res
}
// GetSignature gets the signature property value. The signature property
func (m *VerifySignatureRequestBody) GetSignature()([]byte) {
    if m == nil {
        return nil
    } else {
        return m.signature
    }
}
// GetSigningKeyId gets the signingKeyId property value. The signingKeyId property
func (m *VerifySignatureRequestBody) GetSigningKeyId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.signingKeyId
    }
}
// Serialize serializes information the current object
func (m *VerifySignatureRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteByteArrayValue("digest", m.GetDigest())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteByteArrayValue("signature", m.GetSignature())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("signingKeyId", m.GetSigningKeyId())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteAdditionalData(m.GetAdditionalData())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAdditionalData sets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *VerifySignatureRequestBody) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetDigest sets the digest property value. The digest property
func (m *VerifySignatureRequestBody) SetDigest(value []byte)() {
    if m != nil {
        m.digest = value
    }
}
// SetSignature sets the signature property value. The signature property
func (m *VerifySignatureRequestBody) SetSignature(value []byte)() {
    if m != nil {
        m.signature = value
    }
}
// SetSigningKeyId sets the signingKeyId property value. The signingKeyId property
func (m *VerifySignatureRequestBody) SetSigningKeyId(value *string)() {
    if m != nil {
        m.signingKeyId = value
    }
}
