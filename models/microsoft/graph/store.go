package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
type Store struct {
    Entity
    // Default language of the term store.
    defaultLanguageTag *string;
    // Collection of all groups available in the term store.
    groups []Group;
    // List of languages for the term store.
    languageTags []string;
    // Collection of all sets available in the term store.
    sets []Set;
}
// Instantiates a new store and sets the default values.
func NewStore()(*Store) {
    m := &Store{
        Entity: *NewEntity(),
    }
    return m
}
// Gets the defaultLanguageTag property value. Default language of the term store.
func (m *Store) GetDefaultLanguageTag()(*string) {
    if m == nil {
        return nil
    } else {
        return m.defaultLanguageTag
    }
}
// Gets the groups property value. Collection of all groups available in the term store.
func (m *Store) GetGroups()([]Group) {
    if m == nil {
        return nil
    } else {
        return m.groups
    }
}
// Gets the languageTags property value. List of languages for the term store.
func (m *Store) GetLanguageTags()([]string) {
    if m == nil {
        return nil
    } else {
        return m.languageTags
    }
}
// Gets the sets property value. Collection of all sets available in the term store.
func (m *Store) GetSets()([]Set) {
    if m == nil {
        return nil
    } else {
        return m.sets
    }
}
// The deserialization information for the current model
func (m *Store) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["defaultLanguageTag"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetDefaultLanguageTag(val)
        return nil
    }
    res["groups"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewGroup() })
        if err != nil {
            return err
        }
        res := make([]Group, len(val))
        for i, v := range val {
            res[i] = *(v.(*Group))
        }
        m.SetGroups(res)
        return nil
    }
    res["languageTags"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        res := make([]string, len(val))
        for i, v := range val {
            res[i] = v.(string)
        }
        m.SetLanguageTags(res)
        return nil
    }
    res["sets"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewSet() })
        if err != nil {
            return err
        }
        res := make([]Set, len(val))
        for i, v := range val {
            res[i] = *(v.(*Set))
        }
        m.SetSets(res)
        return nil
    }
    return res
}
func (m *Store) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
func (m *Store) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("defaultLanguageTag", m.GetDefaultLanguageTag())
        if err != nil {
            return err
        }
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetGroups()))
        for i, v := range m.GetGroups() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("groups", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteCollectionOfStringValues("languageTags", m.GetLanguageTags())
        if err != nil {
            return err
        }
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetSets()))
        for i, v := range m.GetSets() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("sets", cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// Sets the defaultLanguageTag property value. Default language of the term store.
// Parameters:
//  - value : Value to set for the defaultLanguageTag property.
func (m *Store) SetDefaultLanguageTag(value *string)() {
    m.defaultLanguageTag = value
}
// Sets the groups property value. Collection of all groups available in the term store.
// Parameters:
//  - value : Value to set for the groups property.
func (m *Store) SetGroups(value []Group)() {
    m.groups = value
}
// Sets the languageTags property value. List of languages for the term store.
// Parameters:
//  - value : Value to set for the languageTags property.
func (m *Store) SetLanguageTags(value []string)() {
    m.languageTags = value
}
// Sets the sets property value. Collection of all sets available in the term store.
// Parameters:
//  - value : Value to set for the sets property.
func (m *Store) SetSets(value []Set)() {
    m.sets = value
}
