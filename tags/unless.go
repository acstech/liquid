package tags

import (
	"errors"
	"fmt"
	"io"

	"github.com/acstech/liquid/core"
)

var (
	endUnless = &End{"unless"}
)

func UnlessFactory(p *core.Parser, config *core.Configuration) (core.Tag, error) {
	condition, err := p.ReadConditionGroup()
	if err != nil {
		return nil, err
	}
	p.SkipPastTag()
	condition.Inverse()
	return &Unless{NewCommon(), condition, nil, nil}, nil
}

func EndUnlessFactory(p *core.Parser, config *core.Configuration) (core.Tag, error) {
	return endUnless, nil
}

type Unless struct {
	*Common
	condition     core.Verifiable
	elseCondition *Else
	lastSibling   core.Tag
}

func (u *Unless) AddSibling(tag core.Tag) error {
	e, ok := tag.(*Else)
	if ok == false {
		return errors.New(fmt.Sprintf("%q does not belong as a sibling of an unless"))
	}
	u.elseCondition = e
	u.lastSibling = tag
	return nil
}

func (u *Unless) LastSibling() core.Tag {
	return u.lastSibling
}

func (u *Unless) Execute(writer io.Writer, data map[string]interface{}) core.ExecuteState {
	if u.condition.IsTrue(data) {
		return u.Common.Execute(writer, data)
	}
	if u.elseCondition != nil {
		return u.elseCondition.Execute(writer, data)
	}
	return core.Normal
}

func (u *Unless) Name() string {
	return "unless"
}

func (u *Unless) Type() core.TagType {
	return core.ContainerTag
}
