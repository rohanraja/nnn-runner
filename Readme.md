## Tasks

## API

### Run command

nnn run 


## Folder structure


- runs
    - 1
        - stdout.txt
        - stderr.txt
        - info.yml
    - 2
        - stdout.txt
        - stderr.txt
        - info.yml

### Info.yml

CMD: <command that was run including params>
StartTime: <Datetime stamp>
EndTime: <Datetime stamp>
Status: Running/Completed/Error


## Dev info

### Building

```sh
# Mac/Linux
make build
make wrapper


# Windows
make -f Makefile.win wrapper
make -f Makefile.win build

# Add ./bin to PATH
```