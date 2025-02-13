package windowsupdates

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// DeploymentSettings 
type DeploymentSettings struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // Settings governing conditions to monitor and automated actions to take.
    monitoring MonitoringSettingsable
    // Settings governing how the content is rolled out.
    rollout RolloutSettingsable
    // Settings governing safeguard holds on offering content.
    safeguard SafeguardSettingsable
    // The type property
    type_escaped *string
}
// NewDeploymentSettings instantiates a new deploymentSettings and sets the default values.
func NewDeploymentSettings()(*DeploymentSettings) {
    m := &DeploymentSettings{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    odatatypeValue := "#microsoft.graph.windowsUpdates.deploymentSettings";
    m.SetType(&odatatypeValue);
    return m
}
// CreateDeploymentSettingsFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateDeploymentSettingsFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    if parseNode != nil {
        mappingValueNode, err := parseNode.GetChildNode("@odata.type")
        if err != nil {
            return nil, err
        }
        if mappingValueNode != nil {
            mappingValue, err := mappingValueNode.GetStringValue()
            if err != nil {
                return nil, err
            }
            if mappingValue != nil {
                mappingStr := *mappingValue
                switch mappingStr {
                    case "#microsoft.graph.windowsUpdates.windowsDeploymentSettings":
                        return NewWindowsDeploymentSettings(), nil
                }
            }
        }
    }
    return NewDeploymentSettings(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *DeploymentSettings) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *DeploymentSettings) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["monitoring"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateMonitoringSettingsFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMonitoring(val.(MonitoringSettingsable))
        }
        return nil
    }
    res["rollout"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateRolloutSettingsFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRollout(val.(RolloutSettingsable))
        }
        return nil
    }
    res["safeguard"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateSafeguardSettingsFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSafeguard(val.(SafeguardSettingsable))
        }
        return nil
    }
    res["@odata.type"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetType(val)
        }
        return nil
    }
    return res
}
// GetMonitoring gets the monitoring property value. Settings governing conditions to monitor and automated actions to take.
func (m *DeploymentSettings) GetMonitoring()(MonitoringSettingsable) {
    if m == nil {
        return nil
    } else {
        return m.monitoring
    }
}
// GetRollout gets the rollout property value. Settings governing how the content is rolled out.
func (m *DeploymentSettings) GetRollout()(RolloutSettingsable) {
    if m == nil {
        return nil
    } else {
        return m.rollout
    }
}
// GetSafeguard gets the safeguard property value. Settings governing safeguard holds on offering content.
func (m *DeploymentSettings) GetSafeguard()(SafeguardSettingsable) {
    if m == nil {
        return nil
    } else {
        return m.safeguard
    }
}
// GetType gets the @odata.type property value. The type property
func (m *DeploymentSettings) GetType()(*string) {
    if m == nil {
        return nil
    } else {
        return m.type_escaped
    }
}
// Serialize serializes information the current object
func (m *DeploymentSettings) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteObjectValue("monitoring", m.GetMonitoring())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("rollout", m.GetRollout())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("safeguard", m.GetSafeguard())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("@odata.type", m.GetType())
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
func (m *DeploymentSettings) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetMonitoring sets the monitoring property value. Settings governing conditions to monitor and automated actions to take.
func (m *DeploymentSettings) SetMonitoring(value MonitoringSettingsable)() {
    if m != nil {
        m.monitoring = value
    }
}
// SetRollout sets the rollout property value. Settings governing how the content is rolled out.
func (m *DeploymentSettings) SetRollout(value RolloutSettingsable)() {
    if m != nil {
        m.rollout = value
    }
}
// SetSafeguard sets the safeguard property value. Settings governing safeguard holds on offering content.
func (m *DeploymentSettings) SetSafeguard(value SafeguardSettingsable)() {
    if m != nil {
        m.safeguard = value
    }
}
// SetType sets the @odata.type property value. The type property
func (m *DeploymentSettings) SetType(value *string)() {
    if m != nil {
        m.type_escaped = value
    }
}
