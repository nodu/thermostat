import DHT22 as DHT
import time

pin = 4 #e.g. if using GPIO pin 4 for sensor

while True:
  sensorarray = DHT.sensor.read(DHT.sensor(pigpio.pi(), pin))

  if sensorarray[2] == 0: #good sensor read
      temp_c = sensorarray[3]
      humidity = sensorarray[4]
      """insert rest of code that you want to use those values for"""
      print(temp_c)
      time.sleep(60) #delay between good reads

  else: #bad sensor read
    time.sleep(3) #wait a few seconds then try again
