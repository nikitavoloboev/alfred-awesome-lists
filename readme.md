# Alfred Awesome Lists [![Workflows](https://img.shields.io/badge/More%20Workflows-🎩-purple.svg)](https://github.com/learn-anything/alfred-workflows) [![Thanks](https://img.shields.io/badge/Say%20Thanks-💗-ff69b4.svg)](https://www.patreon.com/nikitavoloboev)
> [Alfred](https://www.alfredapp.com/) workflow to browse all [awesome lists](https://github.com/sindresorhus/awesome) inside Alfred

<img src="https://i.imgur.com/2ZezJsz.png" width="500" alt="img">

## Installing
Download the workflow from [GitHub releases](https://github.com/nikitavoloboev/alfred-awesome-lists/releases/latest).

## Contributing
[Suggestions](https://github.com/nikitavoloboev/alfred-awesome-lists/issues) and pull requests are highly encouraged!

## Developing
If you want to add features and things to the workflow. It is best to use [this Alfred CLI tool](https://godoc.org/github.com/jason0x43/go-alfred/alfred) which you can install by running:

`go install github.com/jason0x43/go-alfred/alfred`

You can then run `alfred link` inside this repo you cloned. This will make a symbolic link of the [`workflow`](workflow) directory. 

You can then make your changes to the code and after run `alfred build` to build the go binary to `workflow` directory. Which you can then use from inside Alfred [script filters](https://www.alfredapp.com/help/workflows/inputs/script-filter/).

The workflow is built using [Deanishe](https://github.com/deanishe)'s amazing [AwGo](https://github.com/deanishe/awgo) library.

## Thank you 💜
You can support what I do on [Patreon](https://www.patreon.com/nikitavoloboev) or look into [other repositories](https://my.mindnode.com/ZKGETDkUaQUsL3q8q9z788CxG84oEHgDiT79GuzX#-143.5,-902.6,0) I shared. 

## License
MIT © [Nikita Voloboev](https://www.nikitavoloboev.xyz)