# Alfred Awesome Lists [![Workflows](https://img.shields.io/badge/More%20Workflows-🎩-purple.svg)](https://github.com/learn-anything/alfred-workflows) [![Thanks](https://img.shields.io/badge/Say%20Thanks-💗-ff69b4.svg)](https://www.patreon.com/nikitavoloboev)
> [Alfred](https://www.alfredapp.com/) workflow to browse all [awesome lists](https://github.com/sindresorhus/awesome) inside Alfred

<img src="https://i.imgur.com/CDgfOKj.png" width="500" alt="img">

You can also press ⌃ to quickly go to readme of the Awesome list so you can quickly submit PR to it.

## Install
Download the workflow from [GitHub releases](https://github.com/nikitavoloboev/alfred-awesome-lists/releases/latest).

## Contributing
[Suggestions](https://github.com/nikitavoloboev/alfred-awesome-lists/issues) and pull requests are highly encouraged!

## Developing
If you want to add features and things to the workflow. It is best to use [this Alfred CLI tool](https://godoc.org/github.com/jason0x43/go-alfred/alfred) which you can install by running:

`go install github.com/jason0x43/go-alfred/alfred`

You can then clone this repo and run `alfred link` inside it. This will make a symbolic link of the [`workflow`](workflow) directory.

You can then make changes to the code and after run `alfred build` to build the go binary to `workflow` directory. Which you can then use from inside Alfred [script filters](https://www.alfredapp.com/help/workflows/inputs/script-filter/).

## Credits
The workflow uses [AwGo](https://github.com/deanishe/awgo) library for all the Alfred related things.

## Thank you 💜
You can support what I do on [Patreon](https://www.patreon.com/nikitavoloboev) or look into [other projects](https://nikitavoloboev.xyz/projects) I shared.

## License
MIT © [Nikita Voloboev](https://www.nikitavoloboev.xyz)