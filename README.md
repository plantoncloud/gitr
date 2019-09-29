# gitr(Git Rapid): The missing link between Git CLI and SCM Providers

Tool to navigate to important features of SCM efficiently right from the command line.

Did you ever hate searching for clone URL for a repo and then use your mouse pointer to copy the link?

`gitr` solves that problem by cloning git repositories using `browser urls`

### Supported Platforms

`gitr` can be installed on any operating system and it is written in golang

### Install

#### Mac


```
brew tap swarupdonepudi/homebrew-gitr
brew install gitr
```

#### Windows

Coming Soon.

### Examples

Clone repo using Browser URL

```
gitr clone https://github.com/silvermullet/hyper-kube-config/issues
```

> Note that the URL is not a typical git url. This URL is copied from browser's url bar

You can also open the following features of your git repo on SCM Web Interface right from the command line

* branches
* prs
* commits
* issues
* pipelines
* releases

> The below commands will only work when executed from inside the git repo folder

Open the home page of the repo on SCM Web Interface

```
gitr rem
```

Open the Pull Requests on SCM Web Interface

```
gitr prs
```

Open the Branches on SCM Web Interface

```
gitr branches
```

Open the Commits on SCM Web Interface

```
gitr commits
```

Open the Issues on SCM Web Interface

```
gitr issues
```

Open the Pipelines on SCM Web Interface

```
gitr pipelines
```

Open the Releases on SCM Web Interface

```
gitr releases
```

### Support for Enterprise Editions

`gitr` can work with enterprise deployments of Github, Gitlab and Bitbucket(Datacenter) editions as well. 

You need to help `gitr` figure out what SCM system you are using. You can do so by simply creating `~/.gitr.yaml` file and adding *your* SCM hostname and SCM Provider to the config file.

Example:

```
scmSystems:
  - hostname: code.mycompany.net
    scm: gitlab
```

If you are working with differrent SCM enterprise deployments, you can add all of them to `~/.gitr.yaml` file

```
scmSystems:
  - hostname: github.mycompany.com
    scm: github
  - hostname: stash.code.mycompany.net
    scm: bitbucket
```

Below is the list of valid values for `scmSystems[].scm` in `~/.gitr.yaml` 

* gitlab
* github
* bitbucket

### Cleanup

```
brew uninstall gitr
brew untap swarupdonepudi/homebrew-gitr
```