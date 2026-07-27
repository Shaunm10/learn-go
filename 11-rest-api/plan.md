# Rest Api

Required routes:

|Http Verb|Route|Description|Authentication Required| Only Creator|Notes|
|---------|------|----------|--|------|-----|
| GET |     `/events`|          Get a list of available events||
| GET |     `/events/<id>`|     Get a event||
| POST |    `/events/`|         Create a new bookable event|YES||
| PUT |     `/events/<id>`|     Update an event|YES|YES|Requires JWT|
| DELETE |  `/events/<id>`|     Delete an event|Yes|YES|Requires JWT|
| POST |    `/signup`|          Create new user||
| POST |    `/login`|           Authenticate user|||Returns JWT|
| POST |    `/events/<id>/register`|   Register user for event|Yes||Requires JWT|
| DELETE |  `/events/<id>/register`|   Cancel registration|Yes||Requires JWT|


## Password hashing in Go:

Recommended API: `golang.org/x/crypto/bcrypt`

Industry-standard, purpose-built for password hashing (unlike general-purpose hashes like SHA-256, it's deliberately slow and includes a per-password salt automatically).
Already resolvable in this module: golang.org/x/crypto v0.54.0 is present in 11-rest-api/go.mod as an indirect dependency (pulled in transitively), so no new external dependency is introduced — only the bcrypt subpackage import needs to be added, and go mod tidy will promote it from // indirect to a direct requirement.
API shape:

```hashedBytes, err := bcrypt.GenerateFromPassword([]byte(plainPassword), bcrypt.DefaultCost)
// store string(hashedBytes) in the DB

err := bcrypt.CompareHashAndPassword([]byte(storedHash), []byte(candidatePassword))
// err == nil means match
```

Changes to 11-rest-api/models/user.go
Add import "golang.org/x/crypto/bcrypt".
In Save(), before building/executing the INSERT: hash user.Password with bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost), return a wrapped error (matching existing errors.Join style) if hashing fails, and use the resulting hash string (instead of user.Password) in stmt.Exec(...).
Add a new exported helper for future login use:
func (user *User) ValidatePassword(candidate string) error {
    return bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(candidate))
}
Doc comment: caller should have user.Password populated with the hash fetched from the DB (e.g., via a future GetUserByEmail); returns nil on match, a bcrypt error otherwise.
Dependency step
Run go get golang.org/x/crypto/bcrypt (or just go mod tidy) inside 11-rest-api/ after adding the import, so go.mod/go.sum correctly list it as a direct dependency.

Verification
go build ./... in 11-rest-api/ to confirm it compiles.
go vet ./... for a quick sanity check.
Optionally run the existing signup endpoint (per recent commit "finished sign up endpoint") and inspect the users table (e.g. via sqlite3 CLI) to confirm the password column now contains a $2a$... bcrypt hash rather than plaintext.