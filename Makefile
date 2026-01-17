build:
	go build .

start-db:
	cp .env.example .env
	docker compose up -d postgres

stop-db:
	docker compose down postgres --volumes --remove-orphans

populate-db:
	./jobtracker configure
	./jobtracker migrate
	./jobtracker add --company "Facebook" --position "Software Engineer"
	./jobtracker add --company "Google" --position "Data Scientist"
	./jobtracker add --company "Apple" --position "Machine Learning Engineer"
	./jobtracker list
