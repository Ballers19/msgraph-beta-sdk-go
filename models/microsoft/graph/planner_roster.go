package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
type PlannerRoster struct {
    Entity
    // Retrieves the members of the plannerRoster.
    members []PlannerRosterMember;
    // Retrieves the plans contained by the plannerRoster.
    plans []PlannerPlan;
}
// Instantiates a new plannerRoster and sets the default values.
func NewPlannerRoster()(*PlannerRoster) {
    m := &PlannerRoster{
        Entity: *NewEntity(),
    }
    return m
}
// Gets the members property value. Retrieves the members of the plannerRoster.
func (m *PlannerRoster) GetMembers()([]PlannerRosterMember) {
    if m == nil {
        return nil
    } else {
        return m.members
    }
}
// Gets the plans property value. Retrieves the plans contained by the plannerRoster.
func (m *PlannerRoster) GetPlans()([]PlannerPlan) {
    if m == nil {
        return nil
    } else {
        return m.plans
    }
}
// The deserialization information for the current model
func (m *PlannerRoster) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["members"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewPlannerRosterMember() })
        if err != nil {
            return err
        }
        res := make([]PlannerRosterMember, len(val))
        for i, v := range val {
            res[i] = *(v.(*PlannerRosterMember))
        }
        m.SetMembers(res)
        return nil
    }
    res["plans"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewPlannerPlan() })
        if err != nil {
            return err
        }
        res := make([]PlannerPlan, len(val))
        for i, v := range val {
            res[i] = *(v.(*PlannerPlan))
        }
        m.SetPlans(res)
        return nil
    }
    return res
}
func (m *PlannerRoster) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
func (m *PlannerRoster) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetMembers()))
        for i, v := range m.GetMembers() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("members", cast)
        if err != nil {
            return err
        }
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetPlans()))
        for i, v := range m.GetPlans() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("plans", cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// Sets the members property value. Retrieves the members of the plannerRoster.
// Parameters:
//  - value : Value to set for the members property.
func (m *PlannerRoster) SetMembers(value []PlannerRosterMember)() {
    m.members = value
}
// Sets the plans property value. Retrieves the plans contained by the plannerRoster.
// Parameters:
//  - value : Value to set for the plans property.
func (m *PlannerRoster) SetPlans(value []PlannerPlan)() {
    m.plans = value
}
