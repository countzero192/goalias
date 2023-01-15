<div align="center">

<h1> Goalias </h1>

Make shell alias quickly.

<br>

</div>

## Requirements
golang compilator;
# Build
```bash
# clone repository
git clone https://github.com/FejkFejkov25/goalias
cd goalias
# build
make
# install
make install
```

# Usage
```bash
goal -s <your_shell> -n <name_of_alias> -c <command>
```
```bash
goal -s bash -n publicip -c "curl ifconfig.me"
```
## License
[MIT](./LICENSE)
