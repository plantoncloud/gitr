# gitr

Open the following in the SCM Web Insterface of the repo in the browser right from the command line.

* branches
* prs
* commits
* issues
* pipelines
* releases

Did you ever hate searching for clone URL for a repo and then use your mouse pointer to copy the link?

`gitr` solves that problem by cloning git repositories using `browser urls`

### Install

```
brew tap swarupdonepudi/homebrew-gitr
brew install gitr
```

### Examples

Clone repo using Browser URL

```
gitr clone https://github.com/silvermullet/hyper-kube-config/issues
```

> Note that the URL is not a typical git url. This URL is copied from browser's url bar

Open the repo on SCM Web Interface. Run this from a local git repo.

```
gitr rem
```

Open the Pull Requests on SCM Web Interface. Run this from a local git repo.

```
gitr prs
```

Open the Branches on SCM Web Interface. Run this from a local git repo.

```
gitr branches
```

Open the Commits on SCM Web Interface. Run this from a local git repo.

```
gitr commits
```

Open the Issues on SCM Web Interface. Run this from a local git repo.

```
gitr issues
```

Open the Pipelines on SCM Web Interface. Run this from a local git repo.

```
gitr pipelines
```

Open the Releases on SCM Web Interface. Run this from a local git repo.

```
gitr releases
```


### Cleanup

```
brew uninstall gitr
brew untap swarupdonepudi/homebrew-gitr
```