package types

// UniversityTeachingStaff holds the data for university teaching staff.
type UniversityTeachingStaff struct {
	AssistantProfessors int    `json:"assistant_professors" fake:"{number:100,1000}"`
	AssociateProfessors int    `json:"associate_professors" fake:"{number:100,1000}"`
	FullProfessors      int    `json:"full_professors" fake:"{number:300,1500}"`
	Institution         string `json:"institution" fake:"{randomstring:[ΕΘΝΙΚΟ & ΚΑΠΟΔΙΣΤΡΙΑΚΟ ΠΑΝΕΠΙΣΤΗΜΙΟ ΑΘΗΝΩΝ,ΑΡΙΣΤΟΤΕΛΕΙΟ ΠΑΝΕΠΙΣΤΗΜΙΟ ΘΕΣΣΑΛΟΝΙΚΗΣ,ΕΘΝΙΚΟ ΜΕΤΣΟΒΙΟ ΠΟΛΥΤΕΧΝΕΙΟ]}"`
	Lecturers           int    `json:"lecturers" fake:"{number:10,100}"`
	PracticeLecturers   int    `json:"practice_lecturers" fake:"{number:0,15}"`
	PracticeProfessors  int    `json:"practice_professors" fake:"{number:0,5}"`
	Year                int    `json:"year" fake:"{year}"`
}

// StudentsBySchool holds the data for students by school.
type StudentsBySchool struct {
	District               string `json:"district" fake:"{randomstring:[ΔΙΕΥΘΥΝΣΗ Δ.Ε. ΚΟΖΑΝΗΣ,ΔΙΕΥΘΥΝΣΗ Π.Ε. ΕΒΡΟΥ,ΔΙΕΥΘΥΝΣΗ Δ.Ε. ΑΝΑΤΟΛΙΚΗΣ ΑΤΤΙΚΗΣ]}"`
	Jurisdiction           string `json:"jurisdiction" fake:"{randomstring:[ΠΕΡΙΦΕΡΕΙΑΚΗ Δ/ΝΣΗ Α/ΘΜΙΑΣ ΚΑΙ Β/ΘΜΙΑΣ ΕΚΠ/ΣΗΣ ΚΕΝΤΡΙΚΗΣ ΜΑΚΕΔΟΝΙΑΣ,ΠΕΡΙΦΕΡΕΙΑΚΗ Δ/ΝΣΗ Α/ΘΜΙΑΣ ΚΑΙ Β/ΘΜΙΑΣ ΕΚΠ/ΣΗΣ ΑΤΤΙΚΗΣ]}"`
	RegisteredStudentBoys  int    `json:"registered_student_boys" fake:"{number:1,200}"`
	RegisteredStudentGirls int    `json:"registered_student_girls" fake:"{number:1,200}"`
	SchoolClass            string `json:"school_class" fake:"{randomstring:[Νηπιαγωγεία,Δημοτικά Σχολεία,Γυμνάσια,Λύκεια]}"`
	SchoolName             string `json:"school_name" fake:"{randomstring:[4ο ΗΜΕΡΗΣΙΟ ΓΥΜΝΑΣΙΟ ΜΕΤΑΜΟΡΦΩΣΗΣ ΑΤΤΙΚΗΣ,7ο ΗΜΕΡΗΣΙΟ ΓΕΝΙΚΟ ΛΥΚΕΙΟ ΑΧΑΡΝΩΝ]}"`
	SchoolType             string `json:"school_type" fake:"{randomstring:[Ημερήσιο Γυμνάσιο,Ημερήσιο Γενικό Λύκειο,Ιδιωτικό Νηπιαγωγείο]}"`
	Year                   int    `json:"year" fake:"{year}"`
}

// AtlasInternshipStatistics holds the data for Atlas internship system.
type AtlasInternshipStatistics struct {
	Year          int    `json:"year" fake:"{year}"`
	Institution   string `json:"institution" fake:"{randomstring:[ΟΙΚΟΝΟΜΙΚΟ ΠΑΝΕΠΙΣΤΗΜΙΟ ΑΘΗΝΩΝ,ΠΑΝΕΠΙΣΤΗΜΙΟ ΑΙΓΑΙΟΥ,ΤΕΙ ΑΘΗΝΑΣ]}"`
	PrivateSector int    `json:"private_sector" fake:"{number:0,500}"`
	PublicSector  int    `json:"public_sector" fake:"{number:0,500}"`
	NGO           int    `json:"ngo" fake:"{number:0,1}"`
}

// EudoksosRequestsAndDeliveries holds the data for Eudoksos requests and deliveries.
type EudoksosRequestsAndDeliveries struct {
	Year                   int    `json:"year" fake:"{year}"`
	Period                 string `json:"period" fake:"{randomstring:[Χειμερινή,Εαρινή]}"`
	Institution            string `json:"institution" fake:"{randomstring:[ΟΙΚΟΝΟΜΙΚΟ ΠΑΝΕΠΙΣΤΗΜΙΟ ΑΘΗΝΩΝ,ΠΑΝΕΠΙΣΤΗΜΙΟ ΑΙΓΑΙΟΥ,ΤΕΙ ΑΘΗΝΑΣ]}"`
	Department             string `json:"department" fake:"{randomstring:[ΑΙΣΘΗΤΙΚΗΣ ΚΑΙ ΚΟΣΜΗΤΟΛΟΓΙΑΣ,ΦΥΣΙΚΟΘΕΡΑΠΕΙΑΣ,ΝΟΜΙΚΗΣ]}"`
	StudentsWithStatements int    `json:"studentswithstatements" fake:"{number:0,2500}"`
	StudentWithDeliveries  int    `json:"studentwithdeliveries" fake:"{number:0,2500}"`
}
