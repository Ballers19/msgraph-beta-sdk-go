package graph

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

type ApprovalStep struct {
    Entity
    assignedToMe *bool;
    displayName *string;
    justification *string;
    reviewedBy *Identity;
    reviewedDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    reviewResult *string;
    status *string;
}
func NewApprovalStep()(*ApprovalStep) {
    m := &ApprovalStep{
        Entity: *NewEntity(),
    }
    return m
}
func (m *ApprovalStep) GetAssignedToMe()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.assignedToMe
    }
}
func (m *ApprovalStep) GetDisplayName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.displayName
    }
}
func (m *ApprovalStep) GetJustification()(*string) {
    if m == nil {
        return nil
    } else {
        return m.justification
    }
}
func (m *ApprovalStep) GetReviewedBy()(*Identity) {
    if m == nil {
        return nil
    } else {
        return m.reviewedBy
    }
}
func (m *ApprovalStep) GetReviewedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.reviewedDateTime
    }
}
func (m *ApprovalStep) GetReviewResult()(*string) {
    if m == nil {
        return nil
    } else {
        return m.reviewResult
    }
}
func (m *ApprovalStep) GetStatus()(*string) {
    if m == nil {
        return nil
    } else {
        return m.status
    }
}
func (m *ApprovalStep) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["assignedToMe"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        m.SetAssignedToMe(val)
        return nil
    }
    res["displayName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetDisplayName(val)
        return nil
    }
    res["justification"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetJustification(val)
        return nil
    }
    res["reviewedBy"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewIdentity() })
        if err != nil {
            return err
        }
        m.SetReviewedBy(val.(*Identity))
        return nil
    }
    res["reviewedDateTime"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        m.SetReviewedDateTime(val)
        return nil
    }
    res["reviewResult"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetReviewResult(val)
        return nil
    }
    res["status"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetStatus(val)
        return nil
    }
    return res
}
func (m *ApprovalStep) IsNil()(bool) {
    return m == nil
}
func (m *ApprovalStep) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteBoolValue("assignedToMe", m.GetAssignedToMe())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("displayName", m.GetDisplayName())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("justification", m.GetJustification())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("reviewedBy", m.GetReviewedBy())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("reviewedDateTime", m.GetReviewedDateTime())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("reviewResult", m.GetReviewResult())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("status", m.GetStatus())
        if err != nil {
            return err
        }
    }
    return nil
}
func (m *ApprovalStep) SetAssignedToMe(value *bool)() {
    m.assignedToMe = value
}
func (m *ApprovalStep) SetDisplayName(value *string)() {
    m.displayName = value
}
func (m *ApprovalStep) SetJustification(value *string)() {
    m.justification = value
}
func (m *ApprovalStep) SetReviewedBy(value *Identity)() {
    m.reviewedBy = value
}
func (m *ApprovalStep) SetReviewedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.reviewedDateTime = value
}
func (m *ApprovalStep) SetReviewResult(value *string)() {
    m.reviewResult = value
}
func (m *ApprovalStep) SetStatus(value *string)() {
    m.status = value
}
