package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ServicePlanInfo 
type ServicePlanInfo struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // The object the service plan can be assigned to. The possible values are:User - service plan can be assigned to individual users.Company - service plan can be assigned to the entire tenant.
    appliesTo *string
    // The provisioning status of the service plan. The possible values are:Success - Service is fully provisioned.Disabled - Service has been disabled.ErrorStatus - The service plan has not been provisioned and is in an error state.PendingInput - Service is not yet provisioned; awaiting service confirmation.PendingActivation - Service is provisioned but requires explicit activation by administrator (for example, Intune_O365 service plan)PendingProvisioning - Microsoft has added a new service to the product SKU and it has not been activated in the tenant, yet.
    provisioningStatus *string
    // The unique identifier of the service plan.
    servicePlanId *string
    // The name of the service plan.
    servicePlanName *string
}
// NewServicePlanInfo instantiates a new servicePlanInfo and sets the default values.
func NewServicePlanInfo()(*ServicePlanInfo) {
    m := &ServicePlanInfo{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateServicePlanInfoFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateServicePlanInfoFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewServicePlanInfo(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ServicePlanInfo) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetAppliesTo gets the appliesTo property value. The object the service plan can be assigned to. The possible values are:User - service plan can be assigned to individual users.Company - service plan can be assigned to the entire tenant.
func (m *ServicePlanInfo) GetAppliesTo()(*string) {
    if m == nil {
        return nil
    } else {
        return m.appliesTo
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ServicePlanInfo) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["appliesTo"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAppliesTo(val)
        }
        return nil
    }
    res["provisioningStatus"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetProvisioningStatus(val)
        }
        return nil
    }
    res["servicePlanId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetServicePlanId(val)
        }
        return nil
    }
    res["servicePlanName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetServicePlanName(val)
        }
        return nil
    }
    return res
}
// GetProvisioningStatus gets the provisioningStatus property value. The provisioning status of the service plan. The possible values are:Success - Service is fully provisioned.Disabled - Service has been disabled.ErrorStatus - The service plan has not been provisioned and is in an error state.PendingInput - Service is not yet provisioned; awaiting service confirmation.PendingActivation - Service is provisioned but requires explicit activation by administrator (for example, Intune_O365 service plan)PendingProvisioning - Microsoft has added a new service to the product SKU and it has not been activated in the tenant, yet.
func (m *ServicePlanInfo) GetProvisioningStatus()(*string) {
    if m == nil {
        return nil
    } else {
        return m.provisioningStatus
    }
}
// GetServicePlanId gets the servicePlanId property value. The unique identifier of the service plan.
func (m *ServicePlanInfo) GetServicePlanId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.servicePlanId
    }
}
// GetServicePlanName gets the servicePlanName property value. The name of the service plan.
func (m *ServicePlanInfo) GetServicePlanName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.servicePlanName
    }
}
// Serialize serializes information the current object
func (m *ServicePlanInfo) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("appliesTo", m.GetAppliesTo())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("provisioningStatus", m.GetProvisioningStatus())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("servicePlanId", m.GetServicePlanId())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("servicePlanName", m.GetServicePlanName())
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
func (m *ServicePlanInfo) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetAppliesTo sets the appliesTo property value. The object the service plan can be assigned to. The possible values are:User - service plan can be assigned to individual users.Company - service plan can be assigned to the entire tenant.
func (m *ServicePlanInfo) SetAppliesTo(value *string)() {
    if m != nil {
        m.appliesTo = value
    }
}
// SetProvisioningStatus sets the provisioningStatus property value. The provisioning status of the service plan. The possible values are:Success - Service is fully provisioned.Disabled - Service has been disabled.ErrorStatus - The service plan has not been provisioned and is in an error state.PendingInput - Service is not yet provisioned; awaiting service confirmation.PendingActivation - Service is provisioned but requires explicit activation by administrator (for example, Intune_O365 service plan)PendingProvisioning - Microsoft has added a new service to the product SKU and it has not been activated in the tenant, yet.
func (m *ServicePlanInfo) SetProvisioningStatus(value *string)() {
    if m != nil {
        m.provisioningStatus = value
    }
}
// SetServicePlanId sets the servicePlanId property value. The unique identifier of the service plan.
func (m *ServicePlanInfo) SetServicePlanId(value *string)() {
    if m != nil {
        m.servicePlanId = value
    }
}
// SetServicePlanName sets the servicePlanName property value. The name of the service plan.
func (m *ServicePlanInfo) SetServicePlanName(value *string)() {
    if m != nil {
        m.servicePlanName = value
    }
}
