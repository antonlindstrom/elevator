# Elevator

This is a program for simulating an Elevator system written in Golang.

To build this library, run `go build .` in the repository root. If you'd like
to see some output, run `./elevator`.

The program is very simple scheduling algorithm which checks for the direction
of elevators and which one is nearest to the pickup location. To improve this
one could take in account where people are going as well to be able to
collocate people going to the same place. This requires changing the
interface to not only know if the person is going up or down but also which
floor.
