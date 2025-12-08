# About

Runny is a terminal-based application launcher for Linux, built in Go with [BubbleTea](https://github.com/charmbracelet/bubbletea) framework.

# Installation

## From Source

```shell
    git clone https://github.com/b-swist/runny.git
    cd runny
```

Install system-wide:
```shell
    sudo make install
```

Alternatively, Install for a single user:
```shell
    make install PREFIX=~/.local
```

> [!IMPORTANT]
> Ensure that the directory is in `$PATH` (example for `~/.local/bin`)
> ```shell
>   export PATH="$HOME/.local/bin:$PATH"
> ```

# Inspirations

- [rofi](https://github.com/davatorium/rofi)
- [clipse](https://github.com/savedra1/clipse)

# License

Runny is released under the MIT License.
