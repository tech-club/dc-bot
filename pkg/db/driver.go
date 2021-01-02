package db

//Driver defined the used database
type Driver string

const (
	DriverMySQL Driver = "mysql"
)

//StringToDriver check if string contains valid driver name and returns as defined typ
func StringToDriver(driverString string) (Driver, error) {
	switch driverString {
	case "mysql":
		return DriverMySQL, nil
	default:
		return "", ErrUnknownDriver
	}
}
