package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

type TargetedManagedAppProtection struct {
    ManagedAppProtection
    assignments []TargetedManagedAppPolicyAssignment;
    isAssigned *bool;
    targetedAppManagementLevels *AppManagementLevel;
}
func NewTargetedManagedAppProtection()(*TargetedManagedAppProtection) {
    m := &TargetedManagedAppProtection{
        ManagedAppProtection: *NewManagedAppProtection(),
    }
    return m
}
func (m *TargetedManagedAppProtection) GetAssignments()([]TargetedManagedAppPolicyAssignment) {
    if m == nil {
        return nil
    } else {
        return m.assignments
    }
}
func (m *TargetedManagedAppProtection) GetIsAssigned()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.isAssigned
    }
}
func (m *TargetedManagedAppProtection) GetTargetedAppManagementLevels()(*AppManagementLevel) {
    if m == nil {
        return nil
    } else {
        return m.targetedAppManagementLevels
    }
}
func (m *TargetedManagedAppProtection) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.ManagedAppProtection.GetFieldDeserializers()
    res["assignments"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewTargetedManagedAppPolicyAssignment() })
        if err != nil {
            return err
        }
        res := make([]TargetedManagedAppPolicyAssignment, len(val))
        for i, v := range val {
            res[i] = *(v.(*TargetedManagedAppPolicyAssignment))
        }
        m.SetAssignments(res)
        return nil
    }
    res["isAssigned"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        m.SetIsAssigned(val)
        return nil
    }
    res["targetedAppManagementLevels"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseAppManagementLevel)
        if err != nil {
            return err
        }
        cast := val.(AppManagementLevel)
        m.SetTargetedAppManagementLevels(&cast)
        return nil
    }
    return res
}
func (m *TargetedManagedAppProtection) IsNil()(bool) {
    return m == nil
}
func (m *TargetedManagedAppProtection) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.ManagedAppProtection.Serialize(writer)
    if err != nil {
        return err
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetAssignments()))
        for i, v := range m.GetAssignments() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("assignments", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("isAssigned", m.GetIsAssigned())
        if err != nil {
            return err
        }
    }
    if m.GetTargetedAppManagementLevels() != nil {
        cast := m.GetTargetedAppManagementLevels().String()
        err = writer.WriteStringValue("targetedAppManagementLevels", &cast)
        if err != nil {
            return err
        }
    }
    return nil
}
func (m *TargetedManagedAppProtection) SetAssignments(value []TargetedManagedAppPolicyAssignment)() {
    m.assignments = value
}
func (m *TargetedManagedAppProtection) SetIsAssigned(value *bool)() {
    m.isAssigned = value
}
func (m *TargetedManagedAppProtection) SetTargetedAppManagementLevels(value *AppManagementLevel)() {
    m.targetedAppManagementLevels = value
}
