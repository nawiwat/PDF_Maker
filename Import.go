package main

type report struct {
	Report []struct{
		Grades []struct{
			CourseCode 	string 	`json:"courseCode"`
			Final_score string 	`json:"final_score"`
			Points 		float64 	`json:"points"`
			Short_title string 	`json:"short_title"`
			Title 		string 	`json:"title"`
			Units 		int 	`json:"units"`
		} 	`json:"grades"`
		Semester 	string 	`json:"semester"`
		Summary struct{
			Unitspassed float64 	`json:"units"`
			UnitsFact 	float64		`json:"units_fac"`
			FinalQpa 	float64 	`json:"qpa"`
			TotalPoints float64		`json:"points"`
		}	`json:"summary"`
		Cumulative struct{
			Unitspassed float64 	`json:"units"`
			UnitsFact 	float64		`json:"units_fac"`
			FinalQpa 	float64 	`json:"qpa"`
			TotalPoints float64		`json:"points"`
		}	`json:"cumulative"`
	} 	`json:"report"`
}