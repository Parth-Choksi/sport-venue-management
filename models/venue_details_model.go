package models

type VenueDetails struct {
	VenueName          string   `json:"venue_name" validate:"required"`
	VenueLocation      string   `json:"venue_location" validate:"required"`
	VenueAvailableDays []int    `json:"venue_available_days" validate:"required"`
	VenuePrice         int16    `json:"venue_price" validate:"required"`
	VenueDescription   string   `json:"venue_description" validate: "required"`
	VenueImages        []string `json:"venue_images" validate:"required"`
	SportsCategory     []string `json:"sports_category" validate:"required"`
}
