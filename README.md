# go-webskeleton
skeleton for go web projects

This skeleton project contains a basic go web application, using sqlite3 as database and labstack/echo for
writing controller and middleware stuff.

It contains the Entity "user" with the User "admin" preloaded in the database.

It also contains a login form and a simple menu. You can login with "admin/xyz".

---
## Installation
1. `go install github.com/devgek/webskeleton@latest`

---
## Bootstrap new Project
1. `cd {your-go-modules}`
2. `mkdir {project}`
3. `webskeleton bootstrap --type=[api|web|cli] --repository={repository} --user={user} --project={project} --title={projecttitle} --templatedir={GOPATH}\pkg\mod\github.com\devgek\webskeleton@v0.1.7` 
4. `cd {project}`
5. `go run main.go serve --config=_test/config-serve.yaml`
---
## Generate new Entities
1. Write models-file (e.g. account.go) in directory models
2. `webskeleton generate --type=db --path={repository/user/project}` or `webskeleton generate --path={repository/user/project}`

   --> generates files *models/generated/entity_factory_creator.go* and *models/generated/entity_types_impl.go*
   
   --> all *.go - files in directory *models* are taken for generation process
3. Modify *data/datastore* --> complete `db.automigrate(...)` with *account*
---
## Make golang html templates for new Entities (e.g. "account")
1. Modify models-file (e.g. account.go) in directory models, set *gui:yes* and *nav:yes*
2. `webskeleton generate --type=gui`

   --> generates files *web/app/template/templates/account.html*

   --> generates files *web/app/template/templates/account-edit.html*

   --> generates files *web/app/template/templates/entity-nav.html*

   --> all *.go - files in directory *models* are taken for generation process

3. Change and complete field names ("name", "short") to field names of *account.go*
4. Modify *web/app/env/msg/messages.yaml* --> create messages for "account" like for "user"
