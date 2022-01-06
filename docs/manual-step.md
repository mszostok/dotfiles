# Manual configuration

### IntelliJ

_Just login to JetBrains account, as data is saved there._

1. Enable: `Use "CamelHumps" words`.
2. [Restore Local Changes view](https://coderedirect.com/questions/498036/cant-find-git-local-changes-in-intellij-idea-2020-1)
   ![](local-changes.png)
3. Change mappings:
  - `cmd + 1` - Project
  - `cmd + 2` - Terminal
  - `cmd + 3` - Git
  - `shift + cmd + o` - Open Recent
4. Set Tab placement to `None`.
5. Install the [Wakatime](https://wakatime.com/intellij-idea) plugin.

### Wakatime (terminal)

> **NOTE:** Time spend on editor is tracked by a dedicated plugins e.g. by plugin for IntelliJ. To enable it in terminal
the [wakatime-cli](https://github.com/wakatime/wakatime-cli) needs to be installed and proper shell plugin configured.

The Wakatime CLI is defined in [Brewfile](../Brewfile). The only thing is to make sure your API Key is in
your `~/.wakatime.cfg` file.
