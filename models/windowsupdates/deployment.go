package windowsupdates

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
)

// Deployment provides operations to manage the admin singleton.
type Deployment struct {
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Entity
    // Specifies the audience to which content is deployed.
    audience DeploymentAudienceable
    // Specifies what content to deploy. Cannot be changed. Returned by default.
    content DeployableContentable
    // The date and time the deployment was created. Returned by default. Read-only.
    createdDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // The date and time the deployment was last modified. Returned by default. Read-only.
    lastModifiedDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // Settings specified on the specific deployment governing how to deploy content. Returned by default.
    settings DeploymentSettingsable
    // Execution status of the deployment. Returned by default.
    state DeploymentStateable
}
// NewDeployment instantiates a new deployment and sets the default values.
func NewDeployment()(*Deployment) {
    m := &Deployment{
        Entity: *ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.NewEntity(),
    }
    return m
}
// CreateDeploymentFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateDeploymentFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewDeployment(), nil
}
// GetAudience gets the audience property value. Specifies the audience to which content is deployed.
func (m *Deployment) GetAudience()(DeploymentAudienceable) {
    if m == nil {
        return nil
    } else {
        return m.audience
    }
}
// GetContent gets the content property value. Specifies what content to deploy. Cannot be changed. Returned by default.
func (m *Deployment) GetContent()(DeployableContentable) {
    if m == nil {
        return nil
    } else {
        return m.content
    }
}
// GetCreatedDateTime gets the createdDateTime property value. The date and time the deployment was created. Returned by default. Read-only.
func (m *Deployment) GetCreatedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.createdDateTime
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *Deployment) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["audience"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateDeploymentAudienceFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAudience(val.(DeploymentAudienceable))
        }
        return nil
    }
    res["content"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateDeployableContentFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetContent(val.(DeployableContentable))
        }
        return nil
    }
    res["createdDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCreatedDateTime(val)
        }
        return nil
    }
    res["lastModifiedDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLastModifiedDateTime(val)
        }
        return nil
    }
    res["settings"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateDeploymentSettingsFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSettings(val.(DeploymentSettingsable))
        }
        return nil
    }
    res["state"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateDeploymentStateFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetState(val.(DeploymentStateable))
        }
        return nil
    }
    return res
}
// GetLastModifiedDateTime gets the lastModifiedDateTime property value. The date and time the deployment was last modified. Returned by default. Read-only.
func (m *Deployment) GetLastModifiedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.lastModifiedDateTime
    }
}
// GetSettings gets the settings property value. Settings specified on the specific deployment governing how to deploy content. Returned by default.
func (m *Deployment) GetSettings()(DeploymentSettingsable) {
    if m == nil {
        return nil
    } else {
        return m.settings
    }
}
// GetState gets the state property value. Execution status of the deployment. Returned by default.
func (m *Deployment) GetState()(DeploymentStateable) {
    if m == nil {
        return nil
    } else {
        return m.state
    }
}
// Serialize serializes information the current object
func (m *Deployment) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteObjectValue("audience", m.GetAudience())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("content", m.GetContent())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("createdDateTime", m.GetCreatedDateTime())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("lastModifiedDateTime", m.GetLastModifiedDateTime())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("settings", m.GetSettings())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("state", m.GetState())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAudience sets the audience property value. Specifies the audience to which content is deployed.
func (m *Deployment) SetAudience(value DeploymentAudienceable)() {
    if m != nil {
        m.audience = value
    }
}
// SetContent sets the content property value. Specifies what content to deploy. Cannot be changed. Returned by default.
func (m *Deployment) SetContent(value DeployableContentable)() {
    if m != nil {
        m.content = value
    }
}
// SetCreatedDateTime sets the createdDateTime property value. The date and time the deployment was created. Returned by default. Read-only.
func (m *Deployment) SetCreatedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    if m != nil {
        m.createdDateTime = value
    }
}
// SetLastModifiedDateTime sets the lastModifiedDateTime property value. The date and time the deployment was last modified. Returned by default. Read-only.
func (m *Deployment) SetLastModifiedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    if m != nil {
        m.lastModifiedDateTime = value
    }
}
// SetSettings sets the settings property value. Settings specified on the specific deployment governing how to deploy content. Returned by default.
func (m *Deployment) SetSettings(value DeploymentSettingsable)() {
    if m != nil {
        m.settings = value
    }
}
// SetState sets the state property value. Execution status of the deployment. Returned by default.
func (m *Deployment) SetState(value DeploymentStateable)() {
    if m != nil {
        m.state = value
    }
}
