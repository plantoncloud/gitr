# gitr(Git Rapid): The missing link between Git CLI and SCM Providers

Tool to navigate to important features of SCM efficiently right from the command line.

### Supported Platforms

`gitr` can run on linux, windows and mac systems

### Install

#### mac

```
brew tap swarupdonepudi/homebrew-gitr
brew install gitr
```

#### windows

Coming Soon.

### Features

`gitr` has two features

1. Open a repo and different parts of a repo in web browser from command line
2. Clone a repository by creating the directories in the clone url

#### Examples for Opening repo in web browser

You can open the following features of your git repo on SCM Web Interface right from the command line

* web - open the repo home page on web
* rem - open the local branch on web
* prs - open prs/mrs on web
* branches - open branches on web
* commits - open the commits of the local branch on web
* issues - open issues on web
* pipelines - open pipelines/actions on web
* releases - open releases on web
* tags - open tags on web

> The below commands will only work when executed from inside the git repo folder

Open the home page of the repo on SCM Web Interface

```
gitr web
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
gitr pipe
```

Open the Releases on SCM Web Interface

```
gitr releases
```

Open the Tags on SCM Web Interface

```
gitr tags
```

#### Examples for cloning repo

This might not be very useful for repositories hosted on github but is very handy for repositories hosted on gitlab.

clone a repo without creating the directories in the url path

```shell
gitr clone clone git@gitlab.mycompany.net:parent/subgroup1/subgroup2/repo.git
```

clone a repo creating the directories in the url path

```shell
gitr clone clone git@gitlab.mycompany.net:parent/subgroup1/subgroup2/repo.git -c
```

### Support for Enterprise Editions

`gitr` can work with enterprise deployments of Github, Gitlab and Bitbucket(Datacenter) editions as well. 

You need to help `gitr` figure out what SCM system you are using. You can do so by simply creating `~/.gitr.yaml` file and adding *your* SCM hostname and SCM Provider to the config file.

Example:

```yaml
scmSystems:
  - hostname: gitlab.mycompany.net
    provider: gitlab
    defaultBranch: main
```

If you are working with different SCM enterprise deployments, you can add all of them to `~/.gitr.yaml` file

```yaml
scmSystems:
  - hostname: github.mycompany.com
    provider: github
    defaultBranch: master
  - hostname: bitbucket.mycompany.com
    provider: bitbucket
    defaultBranch: master
  - hostname: gitlab.mycompany.com
    provider: gitlab
    defaultBranch: main
```

Below is the list of valid values for `scmSystems[].scm` in `~/.gitr.yaml` 

* gitlab
* github
* bitbucket

Config also support few additional options for cloning repos

| config                     | default | description                                                                                          |
|----------------------------|---------|------------------------------------------------------------------------------------------------------|
| clone.scmHome              |  ""     | if this value is set, then gitr clone will always clone the repos to this path                |
| clone.alwaysCreDir         |  false  | if this is set to true, then gitr clone will always create the directories present in the clone url  |
| clone.includeHostForCreDir |  false  | if this is set to true, then gitr clone will always prefix the hostname to the clone path            |

```yaml
clone:
  scmHome: /Users/swarup/scm
  alwaysCreDir: true
  includeHostForCreDir: false
scmSystems:
  - hostname: github.mycompany.com
    provider: github
    defaultBranch: master
  - hostname: gitlab.mycompany.com
    provider: gitlab
    defaultBranch: main
```


### Cleanup

```
brew uninstall gitr
brew untap swarupdonepudi/homebrew-gitr
```