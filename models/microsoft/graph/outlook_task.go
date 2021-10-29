package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
type OutlookTask struct {
    OutlookItem
    // The name of the person who has been assigned the task in Outlook. Read-only.
    assignedTo *string;
    // The collection of fileAttachment, itemAttachment, and referenceAttachment attachments for the task.  Read-only. Nullable.
    attachments []Attachment;
    // The task body that typically contains information about the task. Note that only HTML type is supported.
    body *ItemBody;
    // The date in the specified time zone that the task was finished.
    completedDateTime *DateTimeTimeZone;
    // The date in the specified time zone that the task is to be finished.
    dueDateTime *DateTimeTimeZone;
    // Set to true if the task has attachments.
    hasAttachments *bool;
    // 
    importance *Importance;
    // 
    isReminderOn *bool;
    // The collection of multi-value extended properties defined for the task. Read-only. Nullable.
    multiValueExtendedProperties []MultiValueLegacyExtendedProperty;
    // 
    owner *string;
    // 
    parentFolderId *string;
    // 
    recurrence *PatternedRecurrence;
    // 
    reminderDateTime *DateTimeTimeZone;
    // 
    sensitivity *Sensitivity;
    // The collection of single-value extended properties defined for the task. Read-only. Nullable.
    singleValueExtendedProperties []SingleValueLegacyExtendedProperty;
    // 
    startDateTime *DateTimeTimeZone;
    // 
    status *TaskStatus;
    // 
    subject *string;
}
// Instantiates a new outlookTask and sets the default values.
func NewOutlookTask()(*OutlookTask) {
    m := &OutlookTask{
        OutlookItem: *NewOutlookItem(),
    }
    return m
}
// Gets the assignedTo property value. The name of the person who has been assigned the task in Outlook. Read-only.
func (m *OutlookTask) GetAssignedTo()(*string) {
    if m == nil {
        return nil
    } else {
        return m.assignedTo
    }
}
// Gets the attachments property value. The collection of fileAttachment, itemAttachment, and referenceAttachment attachments for the task.  Read-only. Nullable.
func (m *OutlookTask) GetAttachments()([]Attachment) {
    if m == nil {
        return nil
    } else {
        return m.attachments
    }
}
// Gets the body property value. The task body that typically contains information about the task. Note that only HTML type is supported.
func (m *OutlookTask) GetBody()(*ItemBody) {
    if m == nil {
        return nil
    } else {
        return m.body
    }
}
// Gets the completedDateTime property value. The date in the specified time zone that the task was finished.
func (m *OutlookTask) GetCompletedDateTime()(*DateTimeTimeZone) {
    if m == nil {
        return nil
    } else {
        return m.completedDateTime
    }
}
// Gets the dueDateTime property value. The date in the specified time zone that the task is to be finished.
func (m *OutlookTask) GetDueDateTime()(*DateTimeTimeZone) {
    if m == nil {
        return nil
    } else {
        return m.dueDateTime
    }
}
// Gets the hasAttachments property value. Set to true if the task has attachments.
func (m *OutlookTask) GetHasAttachments()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.hasAttachments
    }
}
// Gets the importance property value. 
func (m *OutlookTask) GetImportance()(*Importance) {
    if m == nil {
        return nil
    } else {
        return m.importance
    }
}
// Gets the isReminderOn property value. 
func (m *OutlookTask) GetIsReminderOn()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.isReminderOn
    }
}
// Gets the multiValueExtendedProperties property value. The collection of multi-value extended properties defined for the task. Read-only. Nullable.
func (m *OutlookTask) GetMultiValueExtendedProperties()([]MultiValueLegacyExtendedProperty) {
    if m == nil {
        return nil
    } else {
        return m.multiValueExtendedProperties
    }
}
// Gets the owner property value. 
func (m *OutlookTask) GetOwner()(*string) {
    if m == nil {
        return nil
    } else {
        return m.owner
    }
}
// Gets the parentFolderId property value. 
func (m *OutlookTask) GetParentFolderId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.parentFolderId
    }
}
// Gets the recurrence property value. 
func (m *OutlookTask) GetRecurrence()(*PatternedRecurrence) {
    if m == nil {
        return nil
    } else {
        return m.recurrence
    }
}
// Gets the reminderDateTime property value. 
func (m *OutlookTask) GetReminderDateTime()(*DateTimeTimeZone) {
    if m == nil {
        return nil
    } else {
        return m.reminderDateTime
    }
}
// Gets the sensitivity property value. 
func (m *OutlookTask) GetSensitivity()(*Sensitivity) {
    if m == nil {
        return nil
    } else {
        return m.sensitivity
    }
}
// Gets the singleValueExtendedProperties property value. The collection of single-value extended properties defined for the task. Read-only. Nullable.
func (m *OutlookTask) GetSingleValueExtendedProperties()([]SingleValueLegacyExtendedProperty) {
    if m == nil {
        return nil
    } else {
        return m.singleValueExtendedProperties
    }
}
// Gets the startDateTime property value. 
func (m *OutlookTask) GetStartDateTime()(*DateTimeTimeZone) {
    if m == nil {
        return nil
    } else {
        return m.startDateTime
    }
}
// Gets the status property value. 
func (m *OutlookTask) GetStatus()(*TaskStatus) {
    if m == nil {
        return nil
    } else {
        return m.status
    }
}
// Gets the subject property value. 
func (m *OutlookTask) GetSubject()(*string) {
    if m == nil {
        return nil
    } else {
        return m.subject
    }
}
// The deserialization information for the current model
func (m *OutlookTask) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.OutlookItem.GetFieldDeserializers()
    res["assignedTo"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetAssignedTo(val)
        return nil
    }
    res["attachments"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewAttachment() })
        if err != nil {
            return err
        }
        res := make([]Attachment, len(val))
        for i, v := range val {
            res[i] = *(v.(*Attachment))
        }
        m.SetAttachments(res)
        return nil
    }
    res["body"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewItemBody() })
        if err != nil {
            return err
        }
        m.SetBody(val.(*ItemBody))
        return nil
    }
    res["completedDateTime"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewDateTimeTimeZone() })
        if err != nil {
            return err
        }
        m.SetCompletedDateTime(val.(*DateTimeTimeZone))
        return nil
    }
    res["dueDateTime"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewDateTimeTimeZone() })
        if err != nil {
            return err
        }
        m.SetDueDateTime(val.(*DateTimeTimeZone))
        return nil
    }
    res["hasAttachments"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        m.SetHasAttachments(val)
        return nil
    }
    res["importance"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseImportance)
        if err != nil {
            return err
        }
        cast := val.(Importance)
        m.SetImportance(&cast)
        return nil
    }
    res["isReminderOn"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        m.SetIsReminderOn(val)
        return nil
    }
    res["multiValueExtendedProperties"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewMultiValueLegacyExtendedProperty() })
        if err != nil {
            return err
        }
        res := make([]MultiValueLegacyExtendedProperty, len(val))
        for i, v := range val {
            res[i] = *(v.(*MultiValueLegacyExtendedProperty))
        }
        m.SetMultiValueExtendedProperties(res)
        return nil
    }
    res["owner"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetOwner(val)
        return nil
    }
    res["parentFolderId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetParentFolderId(val)
        return nil
    }
    res["recurrence"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewPatternedRecurrence() })
        if err != nil {
            return err
        }
        m.SetRecurrence(val.(*PatternedRecurrence))
        return nil
    }
    res["reminderDateTime"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewDateTimeTimeZone() })
        if err != nil {
            return err
        }
        m.SetReminderDateTime(val.(*DateTimeTimeZone))
        return nil
    }
    res["sensitivity"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseSensitivity)
        if err != nil {
            return err
        }
        cast := val.(Sensitivity)
        m.SetSensitivity(&cast)
        return nil
    }
    res["singleValueExtendedProperties"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewSingleValueLegacyExtendedProperty() })
        if err != nil {
            return err
        }
        res := make([]SingleValueLegacyExtendedProperty, len(val))
        for i, v := range val {
            res[i] = *(v.(*SingleValueLegacyExtendedProperty))
        }
        m.SetSingleValueExtendedProperties(res)
        return nil
    }
    res["startDateTime"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewDateTimeTimeZone() })
        if err != nil {
            return err
        }
        m.SetStartDateTime(val.(*DateTimeTimeZone))
        return nil
    }
    res["status"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseTaskStatus)
        if err != nil {
            return err
        }
        cast := val.(TaskStatus)
        m.SetStatus(&cast)
        return nil
    }
    res["subject"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetSubject(val)
        return nil
    }
    return res
}
func (m *OutlookTask) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
func (m *OutlookTask) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.OutlookItem.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("assignedTo", m.GetAssignedTo())
        if err != nil {
            return err
        }
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetAttachments()))
        for i, v := range m.GetAttachments() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("attachments", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("body", m.GetBody())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("completedDateTime", m.GetCompletedDateTime())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("dueDateTime", m.GetDueDateTime())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("hasAttachments", m.GetHasAttachments())
        if err != nil {
            return err
        }
    }
    if m.GetImportance() != nil {
        cast := m.GetImportance().String()
        err = writer.WriteStringValue("importance", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("isReminderOn", m.GetIsReminderOn())
        if err != nil {
            return err
        }
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetMultiValueExtendedProperties()))
        for i, v := range m.GetMultiValueExtendedProperties() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("multiValueExtendedProperties", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("owner", m.GetOwner())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("parentFolderId", m.GetParentFolderId())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("recurrence", m.GetRecurrence())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("reminderDateTime", m.GetReminderDateTime())
        if err != nil {
            return err
        }
    }
    if m.GetSensitivity() != nil {
        cast := m.GetSensitivity().String()
        err = writer.WriteStringValue("sensitivity", &cast)
        if err != nil {
            return err
        }
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetSingleValueExtendedProperties()))
        for i, v := range m.GetSingleValueExtendedProperties() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("singleValueExtendedProperties", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("startDateTime", m.GetStartDateTime())
        if err != nil {
            return err
        }
    }
    if m.GetStatus() != nil {
        cast := m.GetStatus().String()
        err = writer.WriteStringValue("status", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("subject", m.GetSubject())
        if err != nil {
            return err
        }
    }
    return nil
}
// Sets the assignedTo property value. The name of the person who has been assigned the task in Outlook. Read-only.
// Parameters:
//  - value : Value to set for the assignedTo property.
func (m *OutlookTask) SetAssignedTo(value *string)() {
    m.assignedTo = value
}
// Sets the attachments property value. The collection of fileAttachment, itemAttachment, and referenceAttachment attachments for the task.  Read-only. Nullable.
// Parameters:
//  - value : Value to set for the attachments property.
func (m *OutlookTask) SetAttachments(value []Attachment)() {
    m.attachments = value
}
// Sets the body property value. The task body that typically contains information about the task. Note that only HTML type is supported.
// Parameters:
//  - value : Value to set for the body property.
func (m *OutlookTask) SetBody(value *ItemBody)() {
    m.body = value
}
// Sets the completedDateTime property value. The date in the specified time zone that the task was finished.
// Parameters:
//  - value : Value to set for the completedDateTime property.
func (m *OutlookTask) SetCompletedDateTime(value *DateTimeTimeZone)() {
    m.completedDateTime = value
}
// Sets the dueDateTime property value. The date in the specified time zone that the task is to be finished.
// Parameters:
//  - value : Value to set for the dueDateTime property.
func (m *OutlookTask) SetDueDateTime(value *DateTimeTimeZone)() {
    m.dueDateTime = value
}
// Sets the hasAttachments property value. Set to true if the task has attachments.
// Parameters:
//  - value : Value to set for the hasAttachments property.
func (m *OutlookTask) SetHasAttachments(value *bool)() {
    m.hasAttachments = value
}
// Sets the importance property value. 
// Parameters:
//  - value : Value to set for the importance property.
func (m *OutlookTask) SetImportance(value *Importance)() {
    m.importance = value
}
// Sets the isReminderOn property value. 
// Parameters:
//  - value : Value to set for the isReminderOn property.
func (m *OutlookTask) SetIsReminderOn(value *bool)() {
    m.isReminderOn = value
}
// Sets the multiValueExtendedProperties property value. The collection of multi-value extended properties defined for the task. Read-only. Nullable.
// Parameters:
//  - value : Value to set for the multiValueExtendedProperties property.
func (m *OutlookTask) SetMultiValueExtendedProperties(value []MultiValueLegacyExtendedProperty)() {
    m.multiValueExtendedProperties = value
}
// Sets the owner property value. 
// Parameters:
//  - value : Value to set for the owner property.
func (m *OutlookTask) SetOwner(value *string)() {
    m.owner = value
}
// Sets the parentFolderId property value. 
// Parameters:
//  - value : Value to set for the parentFolderId property.
func (m *OutlookTask) SetParentFolderId(value *string)() {
    m.parentFolderId = value
}
// Sets the recurrence property value. 
// Parameters:
//  - value : Value to set for the recurrence property.
func (m *OutlookTask) SetRecurrence(value *PatternedRecurrence)() {
    m.recurrence = value
}
// Sets the reminderDateTime property value. 
// Parameters:
//  - value : Value to set for the reminderDateTime property.
func (m *OutlookTask) SetReminderDateTime(value *DateTimeTimeZone)() {
    m.reminderDateTime = value
}
// Sets the sensitivity property value. 
// Parameters:
//  - value : Value to set for the sensitivity property.
func (m *OutlookTask) SetSensitivity(value *Sensitivity)() {
    m.sensitivity = value
}
// Sets the singleValueExtendedProperties property value. The collection of single-value extended properties defined for the task. Read-only. Nullable.
// Parameters:
//  - value : Value to set for the singleValueExtendedProperties property.
func (m *OutlookTask) SetSingleValueExtendedProperties(value []SingleValueLegacyExtendedProperty)() {
    m.singleValueExtendedProperties = value
}
// Sets the startDateTime property value. 
// Parameters:
//  - value : Value to set for the startDateTime property.
func (m *OutlookTask) SetStartDateTime(value *DateTimeTimeZone)() {
    m.startDateTime = value
}
// Sets the status property value. 
// Parameters:
//  - value : Value to set for the status property.
func (m *OutlookTask) SetStatus(value *TaskStatus)() {
    m.status = value
}
// Sets the subject property value. 
// Parameters:
//  - value : Value to set for the subject property.
func (m *OutlookTask) SetSubject(value *string)() {
    m.subject = value
}
