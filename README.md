# blu
One click bluetooth reset for mac os x


![Package Mac](https://github.com/barelyhuman/blu/workflows/Package%20Mac/badge.svg)
[![Go Report Card](https://goreportcard.com/badge/github.com/barelyhuman/blu)](https://goreportcard.com/report/github.com/barelyhuman/blu)

## Motivation
I use a wired mouse and a bluetooth keyboard with a docked macbook which looses the bluetooth connectivity when any other iDevice interrupts the frequency having to undock, reset bluetooh and then restart the system, I use blueutil as a reset when the keyboard is accessible but then a GUI helps when the keyboard itself is inaccessible.

## Preview
![](assets/preview.gif)

## Releases 
Check the repositories [Releases](http://github.com/barelyhuman/blu/releases)



## Installation

### Prereq
- Make sure you have `homebrew` and `blueutil` installed 
```sh 
/bin/bash -c "$(curl -fsSL https://raw.githubusercontent.com/Homebrew/install/HEAD/install.sh)"

brew install blueutil
```

The DMG's are unsigned and hence need a few additional steps to install. 

1. Download the *.dmg from [releases](https://github.com/barelyhuman/blu/releases)
2. Mount the dmg , Copy the app to your Applications Folder.
3. Open System Preferences > Security and Privacy > General Tab 
4. Use spotlight or double click the apps from the Applications Folder.
5. Hit `cancel` on the security pop-up that asks you to delete the app. 
6. At this point you should see the the window from Step 3 have a small `Open Anyway` button under the `App Store and Identified Developer`, Click on it to authorize the app. 

Unfortunately, this has to be done everytime you replace or update the app from here till I get a apple developer account.


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

