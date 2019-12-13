package cert

import "testing"

func TestValidCertData(t *testing.T)  {
	c, err := New("Golang", "Bob", "2018-05-11")
	if err != nil {
		t.Errorf("Cert data should be valid. err %v", err)
	}
	if c == nil {
		t.Errorf("Cert should be a valid reference got nil")
	}

	if c.Course != "GOLANG COURSE" {
		t.Errorf("Course name is not valid. Exp: 'GOLANG COURSE', got: '%v'", c.Course)
	}
}

func TestCourseEmptyValue(t *testing.T)  {
	_, err := New("", "bob", "2018-01-05")
	if err == nil {
		t.Errorf("Course must not be empty")
	}
}

func TestCourseTooLong(t *testing.T)  {
	course := "khsqdqshdjsdhsqdhqsldhqsdjqskdjqskjdqksmjdmkqsjdmqsjdmsjdmqjdmJDMDJMSDJQSMDJQSMDJQSMDJQDSSLDJSMQDJ"
	_, err := New(course, "bob", "2018-01-05")
	if err == nil {
		t.Errorf("Course is too long")
	}
}

func TestNameEmpty(t *testing.T)  {
	_, err := New("Test", "", "2018-01-05")
	if err == nil {
		t.Errorf("Name must not be empty")
	}
}

func TestNameTooLong(t *testing.T)  {
	name := "khsqdqshdjsdhsqdhqsldhqsdjqskdjqskjdqksmjdmkqsjdmqsjdmsjdmqjdmJDMDJMSDJQSMDJQSMDJQSMDJQDSSLDJSMQDJ"
	_, err := New("test", name, "2018-01-05")
	if err == nil {
		t.Errorf("Name is too long")
	}
}