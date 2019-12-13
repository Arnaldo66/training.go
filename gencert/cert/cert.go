package cert

import (
	"fmt"
	"strings"
	"time"
)

var maxLenCourse int = 20
var maxLenName int = 30

type Cert struct {
	Course string
	Name string
	Date time.Time

	LabelTitle string
	LabelCompletion string
	LabelPresented string
	LabelParticipation string
	LabelDate string
}

type Saver interface {
	Save(c Cert) error
}

func New(course, name, date string) (*Cert, error)  {
	c, err := validateCourse(course)
	if err != nil {
		return nil, err
	}
	n, err := validateStrings(name, maxLenName)
	if err != nil {
		return nil, err
	}
	d, err := parseDate(date)
	
	cert := &Cert{
		Course:             c,
		Name:               strings.ToTitle(n),
		Date:               d,
		LabelTitle:         fmt.Sprintf("%v Certificate - %v", c, n),
		LabelCompletion:    "Certificate of Completion",
		LabelPresented:     "This certificate is presented to",
		LabelParticipation: fmt.Sprintf("For participation in the %v", c),
		LabelDate:          fmt.Sprintf("Date: %v", d.Format("02/01/2006")),
	}

	return cert, nil
}

func validateCourse(course string) (string, error)  {
	c, err := validateStrings(course, maxLenCourse)
	if err != nil {
		return c, err
	}
	if !strings.HasSuffix(c, "course") {
		c = c + " course"
	}

	return strings.ToTitle(c), nil
}

func validateStrings(course string, maxLen int)  (string, error){
	c := course
	if len(strings.TrimSpace(c)) <= 0 {
		return c, fmt.Errorf("Invalid string")
	}

	if len(c) > maxLen {
		return c, fmt.Errorf("Course is too long")
	}

	return c, nil
}

func parseDate(date string) (time.Time, error) {
	t, err := time.Parse("2006-01-02", date)
	if err != nil {
		return t, err
	}

	return t, nil
}
