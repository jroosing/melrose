package melrose

import (
	"bytes"
	"fmt"
	"log"
	"strings"
)

type Sequence struct {
	Notes [][]Note
}

func (s Sequence) Size() int {
	sum := 0
	s.NotesDo(func(n Note) {
		sum++
	})
	return sum
}

func (s Sequence) Length() int {
	return len(s.Notes)
}

func (s Sequence) At(i int) []Note {
	return s.Notes[i]
}

// SequenceJoin returns s + t
func (s Sequence) SequenceJoin(t Sequence) Sequence {
	return Sequence{append(s.Notes, t.Notes...)}
}

func (s Sequence) NotesCollect(transform func(Note) Note) Sequence {
	notes := make([][]Note, len(s.Notes))
	for i, eachGroup := range s.Notes {
		group := make([]Note, len(eachGroup))
		for j, eachNote := range eachGroup {
			note := transform(eachNote)
			group[j] = note
		}
		notes[i] = group
	}
	return Sequence{Notes: notes}
}

func (s Sequence) NotesDo(block func(Note)) {
	for _, eachGroup := range s.Notes {
		for _, eachNote := range eachGroup {
			block(eachNote)
		}
	}
}

// BuildSequence creates a Sequence from a slice of Note
func BuildSequence(notes []Note) Sequence {
	groups := [][]Note{}
	for _, each := range notes {
		groups = append(groups, []Note{each})
	}
	return Sequence{Notes: groups}
}

var S = MustParseSequence

func MustParseSequence(input string) Sequence {
	if s, err := ParseSequence(input); err != nil {
		log.Fatal("MustParseSequence failed:", err.Error())
		return s
	} else {
		return s
	}
}

// ParseSequence creates a Sequence by reading the format "Note* [Note Note*]* Note*"
func ParseSequence(input string) (Sequence, error) {
	m := Sequence{}
	// hack to keep scanning simple, TODO
	splitable := strings.Replace(input, "[", " [ ", -1)
	splitable = strings.Replace(splitable, "]", " ] ", -1)
	parts := strings.Fields(splitable)
	ingroup := false
	var group []Note
	for _, each := range parts {
		if "[" == each {
			ingroup = true
			group = []Note{}
		} else if "]" == each {
			ingroup = false
			m.Notes = append(m.Notes, group)
		} else {
			next, err := ParseNote(each)
			if err != nil {
				return m, err
			}
			if ingroup {
				group = append(group, next)
			} else {
				m.Notes = append(m.Notes, []Note{next})
			}
		}
	}
	return m, nil
}

func (s Sequence) S() Sequence {
	return s
}

// Conversion

// Storex returns the command line expression that creates the receiver
func (s Sequence) Storex() string {
	return fmt.Sprintf("sequence('%s')", s.String())
}

func (s Sequence) String() string {
	return s.PrintString(PrintAsSpecified)
}

func (s Sequence) PrintString(sharpOrFlatKey int) string {
	var buf bytes.Buffer
	s.writeNotesOn(&buf, (Note).printOn, sharpOrFlatKey)
	return buf.String()
}

func (s Sequence) writeNotesOn(
	buf *bytes.Buffer,
	printer func(n Note, buf *bytes.Buffer, sharpOrFlatKey int),
	sharpOrFlatKey int) {

	for i, each := range s.Notes {
		if i > 0 {
			buf.WriteString(" ")
		}
		if len(each) > 1 {
			buf.WriteString("[")
		}
		for j, other := range each {
			if j > 0 {
				buf.WriteString(" ")
			}
			printer(other, buf, sharpOrFlatKey)
		}
		if len(each) > 1 {
			buf.WriteString("]")
		}
	}
}
