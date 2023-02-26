#!/usr/bin/env python3

from sys import argv
from gpiozero import Servo
from time import sleep

myGPIO = 18

myCorrection = 0.45
maxPW = (2.0+myCorrection)/1000
minPW = (1.0-myCorrection)/1000

servo = Servo(myGPIO, min_pulse_width=minPW, max_pulse_width=maxPW)

# servo.value = .3
arg1 = 0 if len(argv) == 1 else argv[1]
servo.value = float(arg1)
sleep(0.5)
servo.value = None
print("Set: ", arg1)
