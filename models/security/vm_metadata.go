package security

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// VmMetadata 
type VmMetadata struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // The cloudProvider property
    cloudProvider *VmCloudProvider
    // The resourceId property
    resourceId *string
    // The subscriptionId property
    subscriptionId *string
    // The vmId property
    vmId *string
}
// NewVmMetadata instantiates a new vmMetadata and sets the default values.
func NewVmMetadata()(*VmMetadata) {
    m := &VmMetadata{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateVmMetadataFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateVmMetadataFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewVmMetadata(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *VmMetadata) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetCloudProvider gets the cloudProvider property value. The cloudProvider property
func (m *VmMetadata) GetCloudProvider()(*VmCloudProvider) {
    if m == nil {
        return nil
    } else {
        return m.cloudProvider
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *VmMetadata) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["cloudProvider"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseVmCloudProvider)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCloudProvider(val.(*VmCloudProvider))
        }
        return nil
    }
    res["resourceId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetResourceId(val)
        }
        return nil
    }
    res["subscriptionId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSubscriptionId(val)
        }
        return nil
    }
    res["vmId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetVmId(val)
        }
        return nil
    }
    return res
}
// GetResourceId gets the resourceId property value. The resourceId property
func (m *VmMetadata) GetResourceId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.resourceId
    }
}
// GetSubscriptionId gets the subscriptionId property value. The subscriptionId property
func (m *VmMetadata) GetSubscriptionId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.subscriptionId
    }
}
// GetVmId gets the vmId property value. The vmId property
func (m *VmMetadata) GetVmId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.vmId
    }
}
// Serialize serializes information the current object
func (m *VmMetadata) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetCloudProvider() != nil {
        cast := (*m.GetCloudProvider()).String()
        err := writer.WriteStringValue("cloudProvider", &cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("resourceId", m.GetResourceId())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("subscriptionId", m.GetSubscriptionId())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("vmId", m.GetVmId())
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
func (m *VmMetadata) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetCloudProvider sets the cloudProvider property value. The cloudProvider property
func (m *VmMetadata) SetCloudProvider(value *VmCloudProvider)() {
    if m != nil {
        m.cloudProvider = value
    }
}
// SetResourceId sets the resourceId property value. The resourceId property
func (m *VmMetadata) SetResourceId(value *string)() {
    if m != nil {
        m.resourceId = value
    }
}
// SetSubscriptionId sets the subscriptionId property value. The subscriptionId property
func (m *VmMetadata) SetSubscriptionId(value *string)() {
    if m != nil {
        m.subscriptionId = value
    }
}
// SetVmId sets the vmId property value. The vmId property
func (m *VmMetadata) SetVmId(value *string)() {
    if m != nil {
        m.vmId = value
    }
}
