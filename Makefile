# set up
migrate-up:
	migrate -path "/workspaces/gorm-performance/migration" -database "mysql://root:password@tcp(mysql:3306)/db" up

migrate-down:
	migrate -path "/workspaces/gorm-performance/migration" -database "mysql://root:password@tcp(mysql:3306)/db" down

clean:
	mysql -h mysql -u dev-user -p < /workspaces/gorm-performance/devutil/clean.sql

# others
mysql:
	mysql -h mysql -D db -u dev-user -p

lint:
	cd /workspaces/gorm-performance/app \
	&& golangci-lint run