# mszostok's dotfiles

## Contents

What's in there?

- all my `brew` dependencies including: applications, fonts, etc.
  See [`Brewfile`](https://github.com/sobolevn/dotfiles/blob/master/Brewfile)
    - initial file was created via [Homebrew Bundle](https://github.com/Homebrew/homebrew-bundle)
- all my `macOS` configuration. See [`macos`](https://github.com/sobolevn/dotfiles/blob/master/macos/)
- all my shell configuration including [my own `sobole`](https://github.com/sobolevn/sobole-zsh-theme) theme.
  See [`shell/`](https://github.com/sobolevn/dotfiles/tree/master/shell)
  and [`config/zshrc`](https://github.com/sobolevn/dotfiles/blob/master/config/zshrc)

## Installation

The [`dotbot`](https://github.com/anishathalye/dotbot/) is used to set things up. Steps:

1. Clone this repo with: `git clone --depth 1 https://github.com/mszostok/dotfiles dotfiles`
2. `cd dotfiles/`
3. Run: [`bash ./install`](https://github.com/sobolevn/dotfiles/blob/master/install)

## CLI

The [iTerm2](https://iterm2.com/) is used along with [`zsh`](https://github.com/zsh-users/zsh)
and [`oh-my-zsh`](https://github.com/robbyrussell/oh-my-zsh) as the main shell.
The [`zplug`](https://github.com/zplug/zplug) to manage
shell [plugins](https://github.com/sobolevn/dotfiles/blob/master/config/zplugrc). I also have a lot of tools / scripts /
aliases to make my working experience better.

## Apps

The [`brew`](https://brew.sh/) is used to install all free apps for my Mac. I also sync apps from AppStore with `brew`
via [`mas`](https://formulae.brew.sh/formula/mas), so the
resulting [`Brewfile`](https://github.com/sobolevn/dotfiles/blob/master/Brewfile) contains everything.

## Local configuration

Some used tools requires local configuration. Such as `git` with username and email.

Here's the full list:

1. `~/.gitconfig_local` to store any user-specific data
2. `~/.shell_env_local` to store local shell config, like: usernames, passwords, tokens, `gpg` keys and so on

## License

Inspired by (sobolevn/dotfiles)[https://github.com/sobolevn/dotfiles]. The transitive license applies with one additional requirement: don't make any harm. 

[WTFPL](https://en.wikipedia.org/wiki/WTFPL): do the fuck you want. Enjoy!
