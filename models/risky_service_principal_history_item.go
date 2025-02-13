package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// RiskyServicePrincipalHistoryItem provides operations to manage the collection of accessReviewDecision entities.
type RiskyServicePrincipalHistoryItem struct {
    RiskyServicePrincipal
    // The activity related to service principal risk level change.
    activity RiskServicePrincipalActivityable
    // The identifier of the actor of the operation.
    initiatedBy *string
    // The identifier of the service principal.
    servicePrincipalId *string
}
// NewRiskyServicePrincipalHistoryItem instantiates a new riskyServicePrincipalHistoryItem and sets the default values.
func NewRiskyServicePrincipalHistoryItem()(*RiskyServicePrincipalHistoryItem) {
    m := &RiskyServicePrincipalHistoryItem{
        RiskyServicePrincipal: *NewRiskyServicePrincipal(),
    }
    return m
}
// CreateRiskyServicePrincipalHistoryItemFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateRiskyServicePrincipalHistoryItemFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewRiskyServicePrincipalHistoryItem(), nil
}
// GetActivity gets the activity property value. The activity related to service principal risk level change.
func (m *RiskyServicePrincipalHistoryItem) GetActivity()(RiskServicePrincipalActivityable) {
    if m == nil {
        return nil
    } else {
        return m.activity
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *RiskyServicePrincipalHistoryItem) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.RiskyServicePrincipal.GetFieldDeserializers()
    res["activity"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateRiskServicePrincipalActivityFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetActivity(val.(RiskServicePrincipalActivityable))
        }
        return nil
    }
    res["initiatedBy"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetInitiatedBy(val)
        }
        return nil
    }
    res["servicePrincipalId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetServicePrincipalId(val)
        }
        return nil
    }
    return res
}
// GetInitiatedBy gets the initiatedBy property value. The identifier of the actor of the operation.
func (m *RiskyServicePrincipalHistoryItem) GetInitiatedBy()(*string) {
    if m == nil {
        return nil
    } else {
        return m.initiatedBy
    }
}
// GetServicePrincipalId gets the servicePrincipalId property value. The identifier of the service principal.
func (m *RiskyServicePrincipalHistoryItem) GetServicePrincipalId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.servicePrincipalId
    }
}
// Serialize serializes information the current object
func (m *RiskyServicePrincipalHistoryItem) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.RiskyServicePrincipal.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteObjectValue("activity", m.GetActivity())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("initiatedBy", m.GetInitiatedBy())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("servicePrincipalId", m.GetServicePrincipalId())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetActivity sets the activity property value. The activity related to service principal risk level change.
func (m *RiskyServicePrincipalHistoryItem) SetActivity(value RiskServicePrincipalActivityable)() {
    if m != nil {
        m.activity = value
    }
}
// SetInitiatedBy sets the initiatedBy property value. The identifier of the actor of the operation.
func (m *RiskyServicePrincipalHistoryItem) SetInitiatedBy(value *string)() {
    if m != nil {
        m.initiatedBy = value
    }
}
// SetServicePrincipalId sets the servicePrincipalId property value. The identifier of the service principal.
func (m *RiskyServicePrincipalHistoryItem) SetServicePrincipalId(value *string)() {
    if m != nil {
        m.servicePrincipalId = value
    }
}
