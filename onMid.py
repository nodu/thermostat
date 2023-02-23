from gpiozero import Servo
from time import sleep

myGPIO=18

myCorrection=0.45
maxPW=(2.0+myCorrection)/1000
minPW=(1.0-myCorrection)/1000

servo = Servo(myGPIO,min_pulse_width=minPW,max_pulse_width=maxPW)

servo.value=.45
sleep(0.5)
servo.value=None
