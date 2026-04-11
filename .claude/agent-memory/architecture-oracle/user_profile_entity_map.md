---
name: User/Profile entity map
description: Complete mapping of User/Profile domain entities, GORM models, repositories, auth flow, DTOs, and DI wiring — captured before the User/Profile split refactor
type: project
---

## Domain entities

### User (domain)
`packages/site/features/user/domain/user.go`
Fields: ID (uuid), Name, Login, PrimaryEmail, Password (bcrypt hash), Verified, Emails []string
Constructor: `NewUser(name, login, password)` — bcrypt hashes password at construction
Methods: `SetPrimaryEmail(email)`

### Profile (domain) — STUB ONLY
`packages/site/features/user/domain/profile.go`
Fields: ProfileName, UserDisplayName — no ID, no UserID, no email/password. Not used anywhere.

## GORM models

### models.User
`packages/site/shared/db/models/user.go`
Fields: gorm.Model (embeds ID as uint — conflict with uuid PK below), ID uuid, Name, Email, Password, Verified, CreatedAt, UpdatedAt, DeletedAt

### models.Profile
`packages/site/shared/db/models/profile.go`
Fields: gorm.Model, ID uuid, UserID uuid, User (FK to User), Email, Password, Name
Relationship: belongs-to User via UserID

### models.Session
`packages/site/shared/db/models/session.go`
Fields: gorm.Model, ID uuid, UserID uuid, User FK, ProfileID uuid, Profile FK, ExpirationTime

## Auto-migrate registration
`packages/site/shared/db/gorm.go` lines 10-14
All three models (User, Session, Profile) are registered in `appModels` and migrated via `AutoMigrate` on startup.

## Repository

### UserRepository (interface)
`packages/site/features/user/domain/user_repository.go`
Methods: Save(ctx, User), FindByEmail(ctx, email) → User

### GormUserRepository (implementation)
`packages/site/features/user/infrastructure/gorm_user_repository.go`
- Save: maps domain.User → models.User (only ID and Email, TODO comment for profile creation), calls DB.Create
- FindByEmail: queries models.User WHERE email = ? AND deleted_at IS NULL, maps back to domain.User (populates ID, PrimaryEmail, Password, Verified, Name)

## Application services

### RegisterService
`packages/site/features/user/application/register_service.go`
DTO: RegisterReqDTO {Email, Password, Name}
ResponseDTO: UserRegisterResDTO {UserID}
Flow: NewUser(name, email, password) → UserService.CreateUser → UserRepository.Save

### UserService (domain service)
`packages/site/features/user/domain/user_service.go`
- CreateUser: checks for existing email via FindByEmail, returns DATA_CONFLICT if found, else Repository.Save
- VerifyUser: sets Verified=true, calls Repository.Save

## Authentication flow

### Session creation entry points
1. Site (HTMX form): POST /site/session — `packages/site/server/site/session/controller.go`
   Calls SessionAppService.Create(ctx, SessionCreateReq{User: email, Password: password})
2. API: POST /api/session/login — `packages/site/server/api/controllers/session/controller.go`
   Also calls SessionAppService.Create — but instantiates SessionAppService inline with s.Db (raw gorm.DB) and config

### SessionAppService.Create (BROKEN / INCOMPLETE)
`packages/site/features/session/application/session_create.go`
Known bugs:
1. Line 26: calls `s.SessionService.FindByEmail(ctx, email)` — FindByEmail does NOT exist on domain.SessionService (only on GormSessionRepository directly, and not exposed through the domain interface)
2. Line 31: `bcrypt.CompareHashAndPassword([]byte(profile.Password), ...)` — profile is declared as `var profile models.Profile` but never populated; always compares against empty string
3. Return type: declared `(err error)` but site controller expects `(res, err)` — site controller does `res.ID.String()` on the result

### SessionAppService.Delete (BROKEN)
`packages/site/features/session/application/session_delete.go`
Uses `s.DB.Delete(...)` but SessionAppService struct has no DB field (only SessionService and Config).

### Password verification intent
The design intent is: look up models.Profile by email, compare bcrypt hash on profile.Password against provided password, then build a Session via SessionBuilder.Build(profile).

### SessionBuilder.Build
`packages/site/features/session/domain/session_builder.go`
Takes models.Profile directly (cross-layer dependency — domain reaching into infrastructure models).
Reads profile.UserID and profile.ID to populate Session.

## DTOs

- RegisterReqDTO: Email, Password, Name — `packages/site/features/user/application/register_service.go` line 15
- UserRegisterResDTO: UserID — same file line 21
- SessionCreateReq: ID, User (email), Password — `packages/site/features/session/application/session_create.go` line 14
- SessionDto: ID, UserID, ExpirationTime, ProfileID — `packages/site/features/session/application/dto/session_dto.go`

## DI / Wiring modules

### UserFeatureModule
`packages/site/server/features/user/module.go`
Singletons: GormUserRepository → UserRepository, UserService
ExportSingletons: RegisterService (exports to consumers)

### UserApiModule
`packages/site/server/api/controllers/user/module.go`
Imports UserFeatureModule, registers UserController
Route: POST /api/user/register → UserController → RegisterService.Register

### SessionFeatureModule
`packages/site/server/features/session/module.go`
Singletons: GormSessionRepository → SessionRepository
NOTE: SessionAppService is NOT wired here — wired in SiteSessionModule

### SiteSessionModule
`packages/site/server/site/session/module.go`
Singletons: NewSessionAppService (takes domain.SessionService + config), NewJwtVerify
Controller: SiteSessionController
NOTE: domain.SessionService is a struct not interface; wiring works by value type

### ApiController (session)
`packages/site/server/api/controllers/session/controller.go`
Does NOT use the DI container — instantiates SessionAppService inline with raw gorm.DB and config.GetConfig()

## Why: captured for User→Profile split refactor
The refactor goal: User holds only id + lifecycle fields; Email/Name/Password move to Profile; auth looks up Profile not User; one User can have many Profiles.

## How to apply
Use this map to identify every touch point that must change: domain entities, GORM models, GormUserRepository mapper, RegisterService DTO/flow, session_create auth flow (which is currently broken and needs a real ProfileRepository lookup), SessionBuilder cross-layer dependency, auto-migrate list, and all DI wiring modules.
