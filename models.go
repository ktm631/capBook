package capBook

type Publisher struct {
	PublisherID int    `gorm:"primary_key";"AUTO_INCREMENT"`
	Name        string `gorm:"name"`
}

type Author struct {
	AuthorID int    `gorm:"primary_key";"AUTO_INCREMENT"`
	Name     string `gorm:"name"`
	Surname  string `gorm:"surname"`
}

type Book struct {
	BookID         int    `gorm:"primary_key";"AUTO_INCREMENT"`
	LocationID     int    `gorm:"foreign_key";"location_id"`
	OwnerID        int    `gorm:"foreign_key";"owner_id"`
	AuthorID       int    `gorm:"foreign_key";"author_id"`
	PublisherID    int    `gorm:"foreign_key";"publisher_id"`
	Title          string `gorm:"title"`
	Isbn           string `gorm:"isbn"`
	Edition        int    `gorm:"edition"`
	IsFree         bool   `gorm:"is_free"`
	DescriptionURL string `gorm:"description_url"`
}

type Location struct {
	LocationID int    `gorm:"primary_key";"AUTO_INCREMENT"`
	Building   string `gorm:"building"`
	Room       string `gorm:"room"`
}

type Rental struct {
	RentalID        int    `gorm:"primary_key";"AUTO_INCREMENT"`
	BookID          int    `gorm:"foreign_key";"book_id"`
	UserID          int    `gorm:"foreign_key";"user_id"`
	StartDate       string `gorm:"start_date"`
	ExpectedEndDate string `gorm:"expected_end_date"`
	EndDate         string `gorm:"end_date"`
}

type User struct {
	UserID   int    `gorm:"primary_key";"AUTO_INCREMENT"`
	Name     string `gorm:"name"`
	Surname  string `gorm:"surname"`
	Email    string `gorm:"email"`
	Password string `gorm:"password"`
	IsAdmin  bool   `gorm:"is_admin"`
}
