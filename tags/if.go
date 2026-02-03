package tags

import (
	"fmt"
	"io"

	"github.com/acstech/liquid/core"
)

var (
	endIf = &End{"if"}
)

type IfSibling interface {
	Condition() core.Verifiable
	core.Code
}

func IfFactory(p *core.Parser, config *core.Configuration) (core.Tag, error) {
	condition, err := p.ReadConditionGroup()
	if err != nil {
		return nil, err
	}
	i := &If{
		NewCommon(),
		condition,
		make([]IfSibling, 0, 3),
		nil,
	}
	i.conditions = append(i.conditions, i)
	p.SkipPastTag()
	return i, nil
}

func ElseIfFactory(p *core.Parser, config *core.Configuration) (core.Tag, error) {
	condition, err := p.ReadConditionGroup()
	if err != nil {
		return nil, err
	}
	p.SkipPastTag()
	return &ElseIf{NewCommon(), condition}, nil
}

func ElseFactory(p *core.Parser, config *core.Configuration) (core.Tag, error) {
	p.SkipPastTag()
	return &Else{NewCommon(), new(core.TrueCondition)}, nil
}

func EndIfFactory(p *core.Parser, config *core.Configuration) (core.Tag, error) {
	return endIf, nil
}

type If struct {
	*Common
	condition   core.Verifiable
	conditions  []IfSibling
	lastSibling core.Tag
}

func (i *If) AddSibling(tag core.Tag) error {
	ifs, ok := tag.(IfSibling)
	if ok == false {
		return fmt.Errorf("%q does not belong inside of an if", tag.Name())
	}
	i.conditions = append(i.conditions, ifs)
	i.lastSibling = tag
	return nil
}

func (i *If) LastSibling() core.Tag {
	return i.lastSibling
}

func (i *If) Execute(writer io.Writer, data map[string]interface{}) core.ExecuteState {
	for index, condition := range i.conditions {
		if condition.Condition().IsTrue(data) {
			if index == 0 {
				return i.Common.Execute(writer, data)
			} else {
				return condition.Execute(writer, data)
			}
		}
	}
	return core.Normal
}

func (i *If) Name() string {
	return "if"
}

func (i *If) Type() core.TagType {
	return core.ContainerTag
}

func (i *If) Condition() core.Verifiable {
	return i.condition
}

type ElseIf struct {
	*Common
	condition core.Verifiable
}

func (e *ElseIf) AddSibling(tag core.Tag) error {
	panic("AddSibling should not have been called on a elseif")
}

func (e *ElseIf) LastSibling() core.Tag {
	return nil
}

func (e *ElseIf) Name() string {
	return "elseif"
}

func (e *ElseIf) Type() core.TagType {
	return core.SiblingTag
}

func (e *ElseIf) Condition() core.Verifiable {
	return e.condition
}

type Else struct {
	*Common
	condition core.Verifiable
}

func (e *Else) AddSibling(tag core.Tag) error {
	panic("AddSibling should not have been called on a else")
}

func (e *Else) LastSibling() core.Tag {
	return nil
}

func (e *Else) Name() string {
	return "else"
}

func (e *Else) Type() core.TagType {
	return core.SiblingTag
}

func (e *Else) Condition() core.Verifiable {
	return e.condition
}

func (e *Else) AddLeftValue(value core.Value) {
}
