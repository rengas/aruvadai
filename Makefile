migrate-create:
	migrate create  -ext sql -seq  -dir migrations $(filename) 