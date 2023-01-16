# flap

<p align="center">
<img width="150" height="150" src="https://raw.githubusercontent.com/1l0/flap/master/art/logo.png"><br/>
Working proof of the Go server running inside Flutter
</p>

[Video in action](https://www.youtube.com/watch?v=4DPL5F2DVdc)

## Prerequisites

- [Flutter](https://flutter.dev) >2.0
- [Go](https://golang.org) >1.16

## Add platforms

```
flutter create --platforms=ios,macos .
```

## Build a server

```
pushd go && make && popd
```

## Run the app

```
flutter run -d macos
```

## Known issues

- Hot reload doesn't work
  - Todo: Investigate
  - Workaround: Run the server independently (detached) in a dev phase
