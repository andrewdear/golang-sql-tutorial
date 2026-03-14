# Migrations- using goose

https://github.com/pressly/goose

## Install goose
`brew install goose`

## THings to talk about here:
Adding the add priority field migration.
If the project is live using a migration file you can create a new migration which not only gives you a reference to go back to to find out when database changes happened but allows you to role back easily if things go wrong.