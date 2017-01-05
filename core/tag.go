// shared interfaces and utility functions
package core

// interface for an tag markup
type Tag interface {
	AddCode(code Code)
	AddSibling(tag Tag) error
	LastSibling() Tag
	Name() string
	Type() TagType
	Code
}

// The type of tag
type TagType int

const (
	ContainerTag TagType = iota
	LoopTag
	EndTag
	SiblingTag
	StandaloneTag
	NoopTag
)
