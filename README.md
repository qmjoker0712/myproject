# myproject
server for myproject, include:
- statistical

## Project Structure

```text

┌── api:                api module
|__ build:              build dir
|__ cmd:                cli dir
|__ config:             config struct and file
|__ statistical:        statistical service
|__ vendor:             third pkg dir
|__ version:            version info
```

## Usage

### API

```shell
cd myproject

make api
# make

cd build/bin

# start api for test
./myproject-api start -c ../../config/app.toml

```

### statis

* build statis on linux&mac
``` shell
# call 
make statis
# or call
make all

# clean build
make clean 
```

* build statis on Windows
``` shell
# call
.\\buildall.bat statis
# or call
.\\buildall.bat all

# clean build
.\\buildall.bat clean
```

* config file

    The config file in myproject/build/statis.toml

* start statis
``` shell
cd build
# statis [config file]
./bin/statis ./statis.toml
```

