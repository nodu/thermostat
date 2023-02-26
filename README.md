chmod +x thermostat.py

sudo apt-get install pigpio python3-pigpio
sudo systemctl start pigpiod
sudo systemctl enable pigpiod


GPIOZERO_PIN_FACTORY=pigpio python3 thermostat.py

GPIOZERO_PIN_FACTORY=pigpio python3 ~/thermostat/thermostat.py
python ~/thermostat/DHT.py 4

https://abyz.me.uk/rpi/pigpio/examples.html#Python_code/DHT.py
The above library saved the project!


cp-rsync ~/Code/thermostat thermo:~/

Range
servo.value=-.8 - 50
servo.value=.6 - 74?
servo.value=.2 - 69


curl -d '{"value":69}' -H "Content-Type: application/json" -X POST http://localhost:80/temperature
curl -H "Content-Type: application/json" http://localhost/temperature


go run .
npm start


GPIOZERO_PIN_FACTORY=pigpio python thermostat/hw/on.py
GPIOZERO_PIN_FACTORY=pigpio python thermostat/hw/onMid.py
GPIOZERO_PIN_FACTORY=pigpio python thermostat/hw/off.py

python thermostat/hw/DHT.py 4 

~/go/bin/air .
