# mszostok's dotfiles

![](./docs/assets/terminal.png)

## Contents

What's in there?

- all my `brew` dependencies including: applications, fonts, etc. See [`Brewfile`](./Brewfile)
  - initial file was created via [Homebrew Bundle](https://github.com/Homebrew/homebrew-bundle)
- all my `macOS` configuration. See [`macos`](./macos/)
- all my shell configuration including my own theme See [`shell/`](./shell)
  and [`config/zshrc`](./config/zshrc)

## Installation

The [`dotbot`](https://github.com/anishathalye/dotbot/) is used to set things up. Steps:

1. Clone this repo with: `git clone --depth 1 https://github.com/mszostok/dotfiles dotfiles`
2. `cd dotfiles/`
3. Run: [`./install`](./install)

   >**NOTE:** To install macOS Applications export `APPS=true` environment variable.

## CLI

The [iTerm2](https://iterm2.com/) is used along with [`zsh`](https://github.com/zsh-users/zsh)
and [`oh-my-zsh`](https://github.com/robbyrussell/oh-my-zsh) as the main shell.

## Apps

The [`brew`](https://brew.sh/) is used to install all free apps for my Mac. I also sync apps from AppStore with `brew`
via [`mas`](https://formulae.brew.sh/formula/mas), so the
resulting [`Brewfile`](https://github.com/sobolevn/dotfiles/blob/master/Brewfile) contains everything.

## Local configuration

Some used tools requires local configuration. Such as `git` with username and email.

## License

Inspired by [sobolevn/dotfiles](https://github.com/sobolevn/dotfiles). The transitive license applies with one
additional requirement: don't make any harm.

[WTFPL](https://en.wikipedia.org/wiki/WTFPL): do the f* you want. Enjoy!
