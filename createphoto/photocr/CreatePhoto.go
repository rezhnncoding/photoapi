package photocr

type CreatePhoto struct {
	Title         string `form:"Title" validate:"required"`
	Description   string `form:"Description" validate:"required"`
	ImageName     string
	CreatorUserId string
}
type EditPhoto struct {
	Id            string
	Title         string `form:"Title" validate:"required"`
	Description   string `form:"Description" validate:"required"`
	ImageName     string
	CreatorUserId string
}
