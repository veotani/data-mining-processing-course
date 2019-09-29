package doc

import (
	"strings"
	"testing"
)

func TestStudentsTableParsing(t *testing.T) {
	p := Parser{"../data/test_case.doc"}
	students, err := p.ExtractStudents()

	if err != nil {
		t.Error(err)
	}

	if !strings.Contains(students[1], "Абдуллина Диана Марселевна") {
		t.Errorf("first student had to be \"Абдуллина Диана Марселевна\" but got \"%v\"", students[1])
	}

	if !strings.Contains(students[len(students)-1], "Ярославцев Владислав Сергеевич") {
		t.Errorf(
			"first student had to be \"Ярославцев Владислав Сергеевич\" but got \"%v\"",
			students[len(students)-1],
		)
	}

	if len(students) != 64+1 {
		t.Errorf(
			"invalid number of students: %v, had to be 64+1 (because of first line which is header",
			len(students),
		)
	}
}
func TestDocumentParsing(t *testing.T) {
	p := Parser{"../data/test.doc"}
	text, err := p.ExtractAllText()
	if err != nil {
		t.Error(err)
	}

	lines := strings.Split(text, "\n")

	if lines[2][:9] != "TestStart" {
		t.Errorf("first line is invalid: %v", lines[2])
	}

	if lines[len(lines)-5][:7] != "TestEnd" {
		t.Errorf("last line is invalid: %v", lines[len(lines)-5])
	}
}
