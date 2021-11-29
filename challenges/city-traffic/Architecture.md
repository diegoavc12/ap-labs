Architecture- Traffic Simulator
====================================

Introduction
----------------------
This project consists in the creation of a simulator that will demonstrate how cars will move across a city
from a specific point to another following a certain path. During the transportation, the cars will encounter
stop lights and, depending on their state they will continue or stop. Also, they cars can find other cars on
the way, which means that they will have to slow down and let other cars move first before continuing their path.

Programming Language
----------------------
We decided to use the Go programming language for our implementation because, since this is a problem that requires 
concurrent programming, we concluded thatmany of the tools that the language provides such as structs, goroutines and 
channels would be fundamental to manage not only the movement ofthe cars, but also the communication between them as well 
as the behavior of stop lights.

Flow of the Program
----------------------
![City Traffic] (flowchart-city-traffic.png)
