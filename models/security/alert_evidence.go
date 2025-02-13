package security

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// AlertEvidence 
type AlertEvidence struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // The createdDateTime property
    createdDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // The remediationStatus property
    remediationStatus *EvidenceRemediationStatus
    // The remediationStatusDetails property
    remediationStatusDetails *string
    // The roles property
    roles []string
    // The tags property
    tags []string
    // The type property
    type_escaped *string
    // The verdict property
    verdict *EvidenceVerdict
}
// NewAlertEvidence instantiates a new alertEvidence and sets the default values.
func NewAlertEvidence()(*AlertEvidence) {
    m := &AlertEvidence{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    odatatypeValue := "#microsoft.graph.security.alertEvidence";
    m.SetType(&odatatypeValue);
    return m
}
// CreateAlertEvidenceFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateAlertEvidenceFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
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
                    case "#microsoft.graph.security.analyzedMessageEvidence":
                        return NewAnalyzedMessageEvidence(), nil
                    case "#microsoft.graph.security.cloudApplicationEvidence":
                        return NewCloudApplicationEvidence(), nil
                    case "#microsoft.graph.security.deviceEvidence":
                        return NewDeviceEvidence(), nil
                    case "#microsoft.graph.security.fileEvidence":
                        return NewFileEvidence(), nil
                    case "#microsoft.graph.security.ipEvidence":
                        return NewIpEvidence(), nil
                    case "#microsoft.graph.security.mailboxEvidence":
                        return NewMailboxEvidence(), nil
                    case "#microsoft.graph.security.mailClusterEvidence":
                        return NewMailClusterEvidence(), nil
                    case "#microsoft.graph.security.oauthApplicationEvidence":
                        return NewOauthApplicationEvidence(), nil
                    case "#microsoft.graph.security.processEvidence":
                        return NewProcessEvidence(), nil
                    case "#microsoft.graph.security.registryKeyEvidence":
                        return NewRegistryKeyEvidence(), nil
                    case "#microsoft.graph.security.registryValueEvidence":
                        return NewRegistryValueEvidence(), nil
                    case "#microsoft.graph.security.securityGroupEvidence":
                        return NewSecurityGroupEvidence(), nil
                    case "#microsoft.graph.security.urlEvidence":
                        return NewUrlEvidence(), nil
                    case "#microsoft.graph.security.userEvidence":
                        return NewUserEvidence(), nil
                }
            }
        }
    }
    return NewAlertEvidence(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *AlertEvidence) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetCreatedDateTime gets the createdDateTime property value. The createdDateTime property
func (m *AlertEvidence) GetCreatedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.createdDateTime
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *AlertEvidence) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
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
    res["remediationStatus"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseEvidenceRemediationStatus)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRemediationStatus(val.(*EvidenceRemediationStatus))
        }
        return nil
    }
    res["remediationStatusDetails"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRemediationStatusDetails(val)
        }
        return nil
    }
    res["roles"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                res[i] = *(v.(*string))
            }
            m.SetRoles(res)
        }
        return nil
    }
    res["tags"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                res[i] = *(v.(*string))
            }
            m.SetTags(res)
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
    res["verdict"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseEvidenceVerdict)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetVerdict(val.(*EvidenceVerdict))
        }
        return nil
    }
    return res
}
// GetRemediationStatus gets the remediationStatus property value. The remediationStatus property
func (m *AlertEvidence) GetRemediationStatus()(*EvidenceRemediationStatus) {
    if m == nil {
        return nil
    } else {
        return m.remediationStatus
    }
}
// GetRemediationStatusDetails gets the remediationStatusDetails property value. The remediationStatusDetails property
func (m *AlertEvidence) GetRemediationStatusDetails()(*string) {
    if m == nil {
        return nil
    } else {
        return m.remediationStatusDetails
    }
}
// GetRoles gets the roles property value. The roles property
func (m *AlertEvidence) GetRoles()([]string) {
    if m == nil {
        return nil
    } else {
        return m.roles
    }
}
// GetTags gets the tags property value. The tags property
func (m *AlertEvidence) GetTags()([]string) {
    if m == nil {
        return nil
    } else {
        return m.tags
    }
}
// GetType gets the @odata.type property value. The type property
func (m *AlertEvidence) GetType()(*string) {
    if m == nil {
        return nil
    } else {
        return m.type_escaped
    }
}
// GetVerdict gets the verdict property value. The verdict property
func (m *AlertEvidence) GetVerdict()(*EvidenceVerdict) {
    if m == nil {
        return nil
    } else {
        return m.verdict
    }
}
// Serialize serializes information the current object
func (m *AlertEvidence) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteTimeValue("createdDateTime", m.GetCreatedDateTime())
        if err != nil {
            return err
        }
    }
    if m.GetRemediationStatus() != nil {
        cast := (*m.GetRemediationStatus()).String()
        err := writer.WriteStringValue("remediationStatus", &cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("remediationStatusDetails", m.GetRemediationStatusDetails())
        if err != nil {
            return err
        }
    }
    if m.GetRoles() != nil {
        err := writer.WriteCollectionOfStringValues("roles", m.GetRoles())
        if err != nil {
            return err
        }
    }
    if m.GetTags() != nil {
        err := writer.WriteCollectionOfStringValues("tags", m.GetTags())
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
    if m.GetVerdict() != nil {
        cast := (*m.GetVerdict()).String()
        err := writer.WriteStringValue("verdict", &cast)
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
func (m *AlertEvidence) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetCreatedDateTime sets the createdDateTime property value. The createdDateTime property
func (m *AlertEvidence) SetCreatedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    if m != nil {
        m.createdDateTime = value
    }
}
// SetRemediationStatus sets the remediationStatus property value. The remediationStatus property
func (m *AlertEvidence) SetRemediationStatus(value *EvidenceRemediationStatus)() {
    if m != nil {
        m.remediationStatus = value
    }
}
// SetRemediationStatusDetails sets the remediationStatusDetails property value. The remediationStatusDetails property
func (m *AlertEvidence) SetRemediationStatusDetails(value *string)() {
    if m != nil {
        m.remediationStatusDetails = value
    }
}
// SetRoles sets the roles property value. The roles property
func (m *AlertEvidence) SetRoles(value []string)() {
    if m != nil {
        m.roles = value
    }
}
// SetTags sets the tags property value. The tags property
func (m *AlertEvidence) SetTags(value []string)() {
    if m != nil {
        m.tags = value
    }
}
// SetType sets the @odata.type property value. The type property
func (m *AlertEvidence) SetType(value *string)() {
    if m != nil {
        m.type_escaped = value
    }
}
// SetVerdict sets the verdict property value. The verdict property
func (m *AlertEvidence) SetVerdict(value *EvidenceVerdict)() {
    if m != nil {
        m.verdict = value
    }
}
