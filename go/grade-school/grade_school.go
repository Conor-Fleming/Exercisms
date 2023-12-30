package school

type School struct {
	Grades []Grade
}

type Grade struct {
	grade    int
	students []student
}

type student struct {
	name string
}

func New() *School {
	return &School{
		Grades: []Grade{},
	}
}

func (s *School) Add(student string, grade int) {
	return
}

func (s *School) Grade(level int) []string {
	return nil
}

func (s *School) Enrollment() []Grade {
	return nil
}
