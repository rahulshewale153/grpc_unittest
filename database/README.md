Mysql Connection
`development:
    driver: mysql
    open: username:password@tcp(localhost:3306)/databasename?parseTime=true
   
    OR

production:
    driver: mysql
    open: username:password@tcp(localhost:3306)/databasename?parseTime=true
    
    `

Create SQl file for write query

    `goose -path . create  AddSomeTable sql`

Up migration (default use development env  "goose up")

    `goose -env production up`

Down Migration  (default use development env  "goose down")

    `goose -env production down`

redo
    Roll back the most recently applied migration, then run it again. (Delete all record and table again create table)

    `goose -env production redo`

status
    Get Migration status (default use development env  "goose status")

    `goose  -env production status `

dbversion
    Print the current version of the database
    (default use development env  "goose dbversion")

        ` goose -env production dbversion `

down-to
    Roll back migrations to a specific version.       

    `goose down-to 20170506082527`

up-to
    Migrate up to a specific version.

    `goose up-to 20170506082420`