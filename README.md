chmod +x thermostat.py

sudo apt-get install pigpio python3-pigpio
sudo systemctl start pigpiod
sudo systemctl enable pigpiod


GPIOZERO_PIN_FACTORY=pigpio python3 thermostat.py

https://abyz.me.uk/rpi/pigpio/examples.html#Python_code/DHT.py
The above library saved the project!


cp-rsync ~/Code/thermostat thermo:~/
