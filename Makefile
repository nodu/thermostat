deploy:
	cd api; GOOS=linux GOARCH=arm go build -o build/main main.go
	rsync -rltgD --recursive --human-readable --progress -v ~/Code/thermostat/api/build/ thermo:~/thermostat/api/build --delete
	cd ui; npm run build
	rsync -rltgD --recursive --human-readable --progress -v ~/Code/thermostat/ui/build/ thermo:/var/www/thermo/html --delete
	rsync -rltgD --recursive --human-readable --progress -v ~/Code/thermostat/hw/ thermo:~/thermostat/hw --delete

