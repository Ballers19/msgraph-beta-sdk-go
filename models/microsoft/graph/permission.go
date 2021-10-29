package graph

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
type Permission struct {
    Entity
    // A format of yyyy-MM-ddTHH:mm:ssZ of DateTimeOffset indicates the expiration time of the permission. DateTime.MinValue indicates there is no expiration set for this permission. Optional.
    expirationDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    // For user type permissions, the details of the users & applications for this permission. Read-only.
    grantedTo *IdentitySet;
    // For link type permissions, the details of the users to whom permission was granted. Read-only.
    grantedToIdentities []IdentitySet;
    // 
    grantedToIdentitiesV2 []SharePointIdentitySet;
    // 
    grantedToV2 *SharePointIdentitySet;
    // This indicates whether password is set for this permission, it's only showing in response. Optional and Read-only and for OneDrive Personal only.
    hasPassword *bool;
    // Provides a reference to the ancestor of the current permission, if it is inherited from an ancestor. Read-only.
    inheritedFrom *ItemReference;
    // Details of any associated sharing invitation for this permission. Read-only.
    invitation *SharingInvitation;
    // Provides the link details of the current permission, if it is a link type permissions. Read-only.
    link *SharingLink;
    // The type of permission, e.g. read. See below for the full list of roles. Read-only.
    roles []string;
    // A unique token that can be used to access this shared item via the **shares** API. Read-only.
    shareId *string;
}
// Instantiates a new permission and sets the default values.
func NewPermission()(*Permission) {
    m := &Permission{
        Entity: *NewEntity(),
    }
    return m
}
// Gets the expirationDateTime property value. A format of yyyy-MM-ddTHH:mm:ssZ of DateTimeOffset indicates the expiration time of the permission. DateTime.MinValue indicates there is no expiration set for this permission. Optional.
func (m *Permission) GetExpirationDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.expirationDateTime
    }
}
// Gets the grantedTo property value. For user type permissions, the details of the users & applications for this permission. Read-only.
func (m *Permission) GetGrantedTo()(*IdentitySet) {
    if m == nil {
        return nil
    } else {
        return m.grantedTo
    }
}
// Gets the grantedToIdentities property value. For link type permissions, the details of the users to whom permission was granted. Read-only.
func (m *Permission) GetGrantedToIdentities()([]IdentitySet) {
    if m == nil {
        return nil
    } else {
        return m.grantedToIdentities
    }
}
// Gets the grantedToIdentitiesV2 property value. 
func (m *Permission) GetGrantedToIdentitiesV2()([]SharePointIdentitySet) {
    if m == nil {
        return nil
    } else {
        return m.grantedToIdentitiesV2
    }
}
// Gets the grantedToV2 property value. 
func (m *Permission) GetGrantedToV2()(*SharePointIdentitySet) {
    if m == nil {
        return nil
    } else {
        return m.grantedToV2
    }
}
// Gets the hasPassword property value. This indicates whether password is set for this permission, it's only showing in response. Optional and Read-only and for OneDrive Personal only.
func (m *Permission) GetHasPassword()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.hasPassword
    }
}
// Gets the inheritedFrom property value. Provides a reference to the ancestor of the current permission, if it is inherited from an ancestor. Read-only.
func (m *Permission) GetInheritedFrom()(*ItemReference) {
    if m == nil {
        return nil
    } else {
        return m.inheritedFrom
    }
}
// Gets the invitation property value. Details of any associated sharing invitation for this permission. Read-only.
func (m *Permission) GetInvitation()(*SharingInvitation) {
    if m == nil {
        return nil
    } else {
        return m.invitation
    }
}
// Gets the link property value. Provides the link details of the current permission, if it is a link type permissions. Read-only.
func (m *Permission) GetLink()(*SharingLink) {
    if m == nil {
        return nil
    } else {
        return m.link
    }
}
// Gets the roles property value. The type of permission, e.g. read. See below for the full list of roles. Read-only.
func (m *Permission) GetRoles()([]string) {
    if m == nil {
        return nil
    } else {
        return m.roles
    }
}
// Gets the shareId property value. A unique token that can be used to access this shared item via the **shares** API. Read-only.
func (m *Permission) GetShareId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.shareId
    }
}
// The deserialization information for the current model
func (m *Permission) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["expirationDateTime"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        m.SetExpirationDateTime(val)
        return nil
    }
    res["grantedTo"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewIdentitySet() })
        if err != nil {
            return err
        }
        m.SetGrantedTo(val.(*IdentitySet))
        return nil
    }
    res["grantedToIdentities"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewIdentitySet() })
        if err != nil {
            return err
        }
        res := make([]IdentitySet, len(val))
        for i, v := range val {
            res[i] = *(v.(*IdentitySet))
        }
        m.SetGrantedToIdentities(res)
        return nil
    }
    res["grantedToIdentitiesV2"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewSharePointIdentitySet() })
        if err != nil {
            return err
        }
        res := make([]SharePointIdentitySet, len(val))
        for i, v := range val {
            res[i] = *(v.(*SharePointIdentitySet))
        }
        m.SetGrantedToIdentitiesV2(res)
        return nil
    }
    res["grantedToV2"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewSharePointIdentitySet() })
        if err != nil {
            return err
        }
        m.SetGrantedToV2(val.(*SharePointIdentitySet))
        return nil
    }
    res["hasPassword"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        m.SetHasPassword(val)
        return nil
    }
    res["inheritedFrom"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewItemReference() })
        if err != nil {
            return err
        }
        m.SetInheritedFrom(val.(*ItemReference))
        return nil
    }
    res["invitation"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewSharingInvitation() })
        if err != nil {
            return err
        }
        m.SetInvitation(val.(*SharingInvitation))
        return nil
    }
    res["link"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewSharingLink() })
        if err != nil {
            return err
        }
        m.SetLink(val.(*SharingLink))
        return nil
    }
    res["roles"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        res := make([]string, len(val))
        for i, v := range val {
            res[i] = v.(string)
        }
        m.SetRoles(res)
        return nil
    }
    res["shareId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetShareId(val)
        return nil
    }
    return res
}
func (m *Permission) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
func (m *Permission) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteTimeValue("expirationDateTime", m.GetExpirationDateTime())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("grantedTo", m.GetGrantedTo())
        if err != nil {
            return err
        }
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetGrantedToIdentities()))
        for i, v := range m.GetGrantedToIdentities() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("grantedToIdentities", cast)
        if err != nil {
            return err
        }
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetGrantedToIdentitiesV2()))
        for i, v := range m.GetGrantedToIdentitiesV2() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("grantedToIdentitiesV2", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("grantedToV2", m.GetGrantedToV2())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("hasPassword", m.GetHasPassword())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("inheritedFrom", m.GetInheritedFrom())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("invitation", m.GetInvitation())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("link", m.GetLink())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteCollectionOfStringValues("roles", m.GetRoles())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("shareId", m.GetShareId())
        if err != nil {
            return err
        }
    }
    return nil
}
// Sets the expirationDateTime property value. A format of yyyy-MM-ddTHH:mm:ssZ of DateTimeOffset indicates the expiration time of the permission. DateTime.MinValue indicates there is no expiration set for this permission. Optional.
// Parameters:
//  - value : Value to set for the expirationDateTime property.
func (m *Permission) SetExpirationDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.expirationDateTime = value
}
// Sets the grantedTo property value. For user type permissions, the details of the users & applications for this permission. Read-only.
// Parameters:
//  - value : Value to set for the grantedTo property.
func (m *Permission) SetGrantedTo(value *IdentitySet)() {
    m.grantedTo = value
}
// Sets the grantedToIdentities property value. For link type permissions, the details of the users to whom permission was granted. Read-only.
// Parameters:
//  - value : Value to set for the grantedToIdentities property.
func (m *Permission) SetGrantedToIdentities(value []IdentitySet)() {
    m.grantedToIdentities = value
}
// Sets the grantedToIdentitiesV2 property value. 
// Parameters:
//  - value : Value to set for the grantedToIdentitiesV2 property.
func (m *Permission) SetGrantedToIdentitiesV2(value []SharePointIdentitySet)() {
    m.grantedToIdentitiesV2 = value
}
// Sets the grantedToV2 property value. 
// Parameters:
//  - value : Value to set for the grantedToV2 property.
func (m *Permission) SetGrantedToV2(value *SharePointIdentitySet)() {
    m.grantedToV2 = value
}
// Sets the hasPassword property value. This indicates whether password is set for this permission, it's only showing in response. Optional and Read-only and for OneDrive Personal only.
// Parameters:
//  - value : Value to set for the hasPassword property.
func (m *Permission) SetHasPassword(value *bool)() {
    m.hasPassword = value
}
// Sets the inheritedFrom property value. Provides a reference to the ancestor of the current permission, if it is inherited from an ancestor. Read-only.
// Parameters:
//  - value : Value to set for the inheritedFrom property.
func (m *Permission) SetInheritedFrom(value *ItemReference)() {
    m.inheritedFrom = value
}
// Sets the invitation property value. Details of any associated sharing invitation for this permission. Read-only.
// Parameters:
//  - value : Value to set for the invitation property.
func (m *Permission) SetInvitation(value *SharingInvitation)() {
    m.invitation = value
}
// Sets the link property value. Provides the link details of the current permission, if it is a link type permissions. Read-only.
// Parameters:
//  - value : Value to set for the link property.
func (m *Permission) SetLink(value *SharingLink)() {
    m.link = value
}
// Sets the roles property value. The type of permission, e.g. read. See below for the full list of roles. Read-only.
// Parameters:
//  - value : Value to set for the roles property.
func (m *Permission) SetRoles(value []string)() {
    m.roles = value
}
// Sets the shareId property value. A unique token that can be used to access this shared item via the **shares** API. Read-only.
// Parameters:
//  - value : Value to set for the shareId property.
func (m *Permission) SetShareId(value *string)() {
    m.shareId = value
}
