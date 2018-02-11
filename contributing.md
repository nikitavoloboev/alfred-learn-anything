# Contributing
Thank you for wanting to make the workflow better. You can:
- Make suggestions and file bugs in [Issues](../../issues/).
- Fix issues and add features to the workflow with [Pull Requests](../../pulls/). Instructions below will provide the necessary information on how you can do that.
- Improve [Learn anything](https://learn-anything.xyz/) and its [curated lists](https://github.com/learn-anything/learn-anything/wiki/Curated-Lists) that the workflow searches over.

## Developing the workflow
If you want to add features and things to the workflow.

It is best to use [this Alfred CLI tool](https://godoc.org/github.com/jason0x43/go-alfred/alfred) which you can install by running:
`go install github.com/jason0x43/go-alfred/alfred`

You can then clone this repository and run: `alfred link` inside it. This will make a symbolic link of the [`workflow`](workflow) directory.

You can then make changes to the code and after run `alfred build` to build the go binary to `workflow` directory. Which you can then use from inside Alfred [script filters](https://www.alfredapp.com/help/workflows/inputs/script-filter/).

