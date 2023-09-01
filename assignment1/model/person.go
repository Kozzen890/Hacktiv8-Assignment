package person

type Participant struct {
	Id string `json:"id"`
	Student_code string `json:"student_code"`
	Student_name string `json:"student_name"`
	Student_address string `json:"student_address"`
	Student_occupation string `json:"student_occupation"`
	Joining_reason string `json:"joining_reason"`
}

type People struct {
	People []Participant `json:"participants"`
}
