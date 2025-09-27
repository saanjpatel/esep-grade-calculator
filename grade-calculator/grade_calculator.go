package esepunittests

type GradeCalculator struct {
	// one list
	grade_types []Grade
}

type GradeType int

const (
	Assignment GradeType = iota
	Exam
	Essay
)

var gradeTypeName = map[GradeType]string{
	Assignment: "assignment",
	Exam:       "exam",
	Essay:      "essay",
}

func (gt GradeType) String() string {
	return gradeTypeName[gt]
}

type Grade struct {
	Name  string
	Grade int
	Type  GradeType
}

func NewGradeCalculator() *GradeCalculator {
	return &GradeCalculator{
		// one list
		grade_types: make([]Grade, 0),
	}
}

func (gc *GradeCalculator) GetFinalGrade() string {
	numericalGrade := gc.calculateNumericalGrade()

	if numericalGrade >= 90 {
		return "A"
	} else if numericalGrade >= 80 {
		return "B"
	} else if numericalGrade >= 70 {
		return "C"
	} else if numericalGrade >= 60 {
		return "D"
	}

	return "F"
}

func (gc *GradeCalculator) AddGrade(name string, grade int, gradeType GradeType) {
	// append to one list
	gc.grade_types = append(gc.grade_types, Grade{
		Name:  name,
		Grade: grade,
		Type:  gradeType,
	})
}

func (gc *GradeCalculator) calculateNumericalGrade() int {
	// add string to check type in computeAverage
	assignment_average := computeAverage(gc.grade_types, "assignment")
	exam_average := computeAverage(gc.grade_types, "exam")
	essay_average := computeAverage(gc.grade_types, "essay") // change so it calculates essay average

	weighted_grade := float64(assignment_average)*.5 + float64(exam_average)*.35 + float64(essay_average)*.15

	return int(weighted_grade)
}

func computeAverage(grade_types []Grade, gradeType string) int {
	sum := 0
	// intialize list for specific type
	specific_grades := make([]int, 0)
	// fix syntax based on documentation
	// if it is type that is being looked for add to list
	for _, grade := range grade_types {
		if grade.Type.String() == gradeType {
			specific_grades = append(specific_grades, grade.Grade)
		}

	}
	// now take average of that list
	for _, grade := range specific_grades {
		sum += grade
	}

	return sum / len(specific_grades)
}

// add function like getFinalGrade but for pass/fail
func (gc *GradeCalculator) GetPassFailFinalGrade() string {
	numericalGrade := gc.calculateNumericalGrade()

	if numericalGrade >= 70 {
		return "Pass"
	} else {
		return "Fail"
	}
}
