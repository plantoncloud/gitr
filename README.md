# gitr

Tool to clone git repositories and locate git repo on SCM using Browser URLs

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