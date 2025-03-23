## Tasks
- Provide binaries in Releases of Github page for easy installation
- Change name to "nn"
- Add command to open logs in vscode for easy viewing
- Add command to re-run the last command
- Add ProcessID to info.yml
- Create a user config, nnn.config.json which can have a notification method on completion
- Notification methods:
    - Deskt
- Add command to self update from github release
- Script to auto release

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

### How to install

Releases -https://github.com/rohanraja/nnn-runner/releases/tag/v1.0.0


```sh
wget https://github.com/rohanraja/nnn-runner/releases/download/v1.0.0/nnn_amd64_win.exe -o nn.exe
wget https://github.com/rohanraja/nnn-runner/releases/download/v1.0.0/run_wrap_amd64_win.exe -o run_wrap.exe
```



### Installing on Windows

To install the latest version on Windows, run the following command in the command prompt:

```sh
powershell -ExecutionPolicy Bypass -Command "Invoke-WebRequest -Uri https://raw.githubusercontent.com/rohanraja/nnn-runner/master/scripts/install_win.ps1 -OutFile install_win.ps1; .\install_win.ps1"
```

This command downloads the `install_win.ps1` script from the repository and executes it to install the tool.