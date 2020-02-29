package itszwizard_objects

/*
The Session User is the data which is stored in the cookie.
It must be called gob.Register (itswizard_objects.SessionUser{})
*/
type SessionUser struct {
	Username            string
	UserID              uint
	FirstAuthentication bool
	Authenticated       bool
	TwoFac              bool
	Firstname           string
	Lastname            string
	Mobile              string
	IpAddress           string
	Institution         string
	School              string
	Email               string
	Information         string
	Admin               bool
	OrganisationID      uint
	InstitutionID       uint
}

/*
Creates a pointer to a SessionUser
*/
func NewSessionUser() *SessionUser {
	return new(SessionUser)
}
