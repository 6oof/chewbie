# Dealing with database:
We're using https://github.com/rubenv/sql-migrate for migrations,
to get help execute:
`sql-migrate help` 

We're using https://sqlc.dev/ as query builder. to get familiar with sqlc go to https://docs.sqlc.dev/en/latest/usage.html

# Dealing with templates:
make sure the first template in the slice has
`{{define "base"}}{{end}}`
if this is not the case it won't render properly.

# Killing the process when it fails
fuser -k 8080/tcp


implement session auth
