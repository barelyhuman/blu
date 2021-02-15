# blu
One click bluetooth reset for mac os x


## Motivation
I use a wired mouse and a bluetooth keyboard with a docked macbook which looses the bluetooth connectivity when any other iDevice interrupts the frequency having to undock, reset bluetooh and then restart the system, I use blueutil as a reset when the keyboard is accessible but then a GUI helps when the keyboard itself is inaccessible.

## Preview
![](assets/preview.gif)

## Releases 
Check the repositories [Releases](http://github.com/barelyhuman/blu/releases)


## Dev 
Make sure you have go mod support, so go version >= 1.3

- clone the repo

```sh
git clone https://github.com/barelyhuman/blu
```
- run the package to get the deps 
```sh
go run .
```
#### Build
```
go run build
```

### Roadmap 
- [x] reset using an installed blueutil binary
- [ ] add in a static blueutil binary as backup for edge cases 
- [ ] customize the Widgets to make it look better.

