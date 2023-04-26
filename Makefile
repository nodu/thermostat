deploy:
	cd api; GOOS=linux GOARCH=arm go build -o build/main main.go
	rsync -rltgD --recursive --human-readable --progress -v ./api/build/ thermo:~/thermostat/api/build --delete
	cd ui; npm run build
	rsync -rltgD --recursive --human-readable --progress -v ./ui/build/ thermo:/var/www/thermo/html --delete
	rsync -rltgD --recursive --human-readable --progress -v ./hw/ thermo:~/thermostat/hw --delete

