package getmailtips

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
)

type GetMailTips struct {
    additionalData map[string]interface{};
    automaticReplies *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.AutomaticRepliesMailTips;
    customMailTip *string;
    deliveryRestricted *bool;
    emailAddress *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.EmailAddress;
    error *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.MailTipsError;
    externalMemberCount *int32;
    isModerated *bool;
    mailboxFull *bool;
    maxMessageSize *int32;
    recipientScope *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.RecipientScopeType;
    recipientSuggestions []i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.Recipient;
    totalMemberCount *int32;
}
func NewGetMailTips()(*GetMailTips) {
    m := &GetMailTips{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
func (m *GetMailTips) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
func (m *GetMailTips) GetAutomaticReplies()(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.AutomaticRepliesMailTips) {
    if m == nil {
        return nil
    } else {
        return m.automaticReplies
    }
}
func (m *GetMailTips) GetCustomMailTip()(*string) {
    if m == nil {
        return nil
    } else {
        return m.customMailTip
    }
}
func (m *GetMailTips) GetDeliveryRestricted()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.deliveryRestricted
    }
}
func (m *GetMailTips) GetEmailAddress()(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.EmailAddress) {
    if m == nil {
        return nil
    } else {
        return m.emailAddress
    }
}
func (m *GetMailTips) GetError()(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.MailTipsError) {
    if m == nil {
        return nil
    } else {
        return m.error
    }
}
func (m *GetMailTips) GetExternalMemberCount()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.externalMemberCount
    }
}
func (m *GetMailTips) GetIsModerated()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.isModerated
    }
}
func (m *GetMailTips) GetMailboxFull()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.mailboxFull
    }
}
func (m *GetMailTips) GetMaxMessageSize()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.maxMessageSize
    }
}
func (m *GetMailTips) GetRecipientScope()(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.RecipientScopeType) {
    if m == nil {
        return nil
    } else {
        return m.recipientScope
    }
}
func (m *GetMailTips) GetRecipientSuggestions()([]i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.Recipient) {
    if m == nil {
        return nil
    } else {
        return m.recipientSuggestions
    }
}
func (m *GetMailTips) GetTotalMemberCount()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.totalMemberCount
    }
}
func (m *GetMailTips) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["automaticReplies"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.NewAutomaticRepliesMailTips() })
        if err != nil {
            return err
        }
        m.SetAutomaticReplies(val.(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.AutomaticRepliesMailTips))
        return nil
    }
    res["customMailTip"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetCustomMailTip(val)
        return nil
    }
    res["deliveryRestricted"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        m.SetDeliveryRestricted(val)
        return nil
    }
    res["emailAddress"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.NewEmailAddress() })
        if err != nil {
            return err
        }
        m.SetEmailAddress(val.(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.EmailAddress))
        return nil
    }
    res["error"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.NewMailTipsError() })
        if err != nil {
            return err
        }
        m.SetError(val.(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.MailTipsError))
        return nil
    }
    res["externalMemberCount"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        m.SetExternalMemberCount(val)
        return nil
    }
    res["isModerated"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        m.SetIsModerated(val)
        return nil
    }
    res["mailboxFull"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        m.SetMailboxFull(val)
        return nil
    }
    res["maxMessageSize"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        m.SetMaxMessageSize(val)
        return nil
    }
    res["recipientScope"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.ParseRecipientScopeType)
        if err != nil {
            return err
        }
        cast := val.(i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.RecipientScopeType)
        m.SetRecipientScope(&cast)
        return nil
    }
    res["recipientSuggestions"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.NewRecipient() })
        if err != nil {
            return err
        }
        res := make([]i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.Recipient, len(val))
        for i, v := range val {
            res[i] = *(v.(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.Recipient))
        }
        m.SetRecipientSuggestions(res)
        return nil
    }
    res["totalMemberCount"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        m.SetTotalMemberCount(val)
        return nil
    }
    return res
}
func (m *GetMailTips) IsNil()(bool) {
    return m == nil
}
func (m *GetMailTips) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteObjectValue("automaticReplies", m.GetAutomaticReplies())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("customMailTip", m.GetCustomMailTip())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("deliveryRestricted", m.GetDeliveryRestricted())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("emailAddress", m.GetEmailAddress())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("error", m.GetError())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("externalMemberCount", m.GetExternalMemberCount())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("isModerated", m.GetIsModerated())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("mailboxFull", m.GetMailboxFull())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("maxMessageSize", m.GetMaxMessageSize())
        if err != nil {
            return err
        }
    }
    if m.GetRecipientScope() != nil {
        cast := m.GetRecipientScope().String()
        err := writer.WriteStringValue("recipientScope", &cast)
        if err != nil {
            return err
        }
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetRecipientSuggestions()))
        for i, v := range m.GetRecipientSuggestions() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err := writer.WriteCollectionOfObjectValues("recipientSuggestions", cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("totalMemberCount", m.GetTotalMemberCount())
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
func (m *GetMailTips) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
func (m *GetMailTips) SetAutomaticReplies(value *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.AutomaticRepliesMailTips)() {
    m.automaticReplies = value
}
func (m *GetMailTips) SetCustomMailTip(value *string)() {
    m.customMailTip = value
}
func (m *GetMailTips) SetDeliveryRestricted(value *bool)() {
    m.deliveryRestricted = value
}
func (m *GetMailTips) SetEmailAddress(value *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.EmailAddress)() {
    m.emailAddress = value
}
func (m *GetMailTips) SetError(value *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.MailTipsError)() {
    m.error = value
}
func (m *GetMailTips) SetExternalMemberCount(value *int32)() {
    m.externalMemberCount = value
}
func (m *GetMailTips) SetIsModerated(value *bool)() {
    m.isModerated = value
}
func (m *GetMailTips) SetMailboxFull(value *bool)() {
    m.mailboxFull = value
}
func (m *GetMailTips) SetMaxMessageSize(value *int32)() {
    m.maxMessageSize = value
}
func (m *GetMailTips) SetRecipientScope(value *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.RecipientScopeType)() {
    m.recipientScope = value
}
func (m *GetMailTips) SetRecipientSuggestions(value []i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.Recipient)() {
    m.recipientSuggestions = value
}
func (m *GetMailTips) SetTotalMemberCount(value *int32)() {
    m.totalMemberCount = value
}
