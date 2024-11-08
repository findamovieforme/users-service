package models

// User represents a user with favorite genres and liked movies
type User struct {
	ID             int      `json:"id"`
	FavoriteGenres []Genre  `json:"favorite_genres"` // Use Genre type here
	LikedMovies    []string `json:"liked_movies"`    // List of IMDb IDs
}

// Genre represents a fixed set of genres
type Genre string

const (
	Action      Genre = "Action"
	Adventure   Genre = "Adventure"
	Animation   Genre = "Animation"
	Biography   Genre = "Biography"
	Comedy      Genre = "Comedy"
	Crime       Genre = "Crime"
	Documentary Genre = "Documentary"
	Drama       Genre = "Drama"
	Family      Genre = "Family"
	Fantasy     Genre = "Fantasy"
	History     Genre = "History"
	Horror      Genre = "Horror"
	Music       Genre = "Music"
	Musical     Genre = "Musical"
	Mystery     Genre = "Mystery"
	Romance     Genre = "Romance"
	SciFi       Genre = "Sci-Fi"
	Sport       Genre = "Sport"
	Thriller    Genre = "Thriller"
	War         Genre = "War"
	Western     Genre = "Western"
	Unknown     Genre = "Unknown"
)
