package ApiSkeleton

type Runoob struct {
    // runoob_id 
	 RunoobId int64 `gorm:"runoob_id" json:"runoob_id"` 
   // runoob_title 
	 RunoobTitle string `gorm:"runoob_title" json:"runoob_title"` 
   // runoob_author 
	 RunoobAuthor string `gorm:"runoob_author" json:"runoob_author"` 
   // submission_date 
	 SubmissionDate string `gorm:"submission_date" json:"submission_date"` 

}

func (model *Runoob) TableName() string {
	return "runoob"
}