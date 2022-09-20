# Pyroscope Practice

## Architecture
![picture 1](https://i.imgur.com/swf3yQv.png)

## Start Pyroscope Server

**Option 1. docker**
```bash
docker run -it -p 4040:4040 pyroscope/pyroscope:latest server
```

**Option 2. macos**
```bash
brew install pyroscope-io/brew/pyroscope
brew services start pyroscope
```

## Run app
```bash
make run
```

## poke-fibo (CPU intensive job) and see profiling results

produce loading:
```bash
poke-fibo
```
open server on browser:
`http://localhost:4040/`


![picture 2](https://i.imgur.com/W8gRqZK.png)
