# gitr(Git Rapid): The missing link between Git CLI and SCM Providers

`gitr` reduces a ton of clicking and waiting for the git repo web pages to load by **taking you directly** to the page that you would like to see right from the command line. You may think that it's not a lot of waiting but trust me, you will notice the difference once you start using `gitr`.

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

### Dry Run

Currently `gitr` only supports `github`, `gitlab` and `bitbucket`. Each provider has slight differences in the urls to access different parts of a repo. `gitr` could just work out of the box for a different cloud provider ex: AWS CodeCommit or it may not behave as expected.

So, if `gitr` does not work as you expect it to, both for supported providers and other providers, you will be able to see what urls `gitr` produces but using `--dry` option. This option is available for both web and clone features.

> Note: When `--dry` or `-d` flags are used, gitr will not open up the repo or clone the repo. It simply displays the info to the console that it would use in non-dry mode

Here is an example of the output when `--dry or -d` flag is passed to gitr

```shell
> gitr rem -d

+---------------+-------------------------------------------------------+
| remote        | git@github.com:swarupdonepudi/gitr.git                |
+---------------+-------------------------------------------------------+
| provider      | github                                                |
+---------------+-------------------------------------------------------+
| host          | github.com                                            |
+---------------+-------------------------------------------------------+
| repo-path     | swarupdonepudi/gitr                                   |
+---------------+-------------------------------------------------------+
| repo-name     | gitr                                                  |
+---------------+-------------------------------------------------------+
| branch        | master                                                |
+---------------+-------------------------------------------------------+
| url-web       | https://github.com/swarupdonepudi/gitr                |
+---------------+-------------------------------------------------------+
| url-remote    | https://github.com/swarupdonepudi/gitr/tree/master    |
+---------------+-------------------------------------------------------+
| url-commits   | https://github.com/swarupdonepudi/gitr/commits/master |
+---------------+-------------------------------------------------------+
| url-branches  | https://github.com/swarupdonepudi/gitr/branches       |
+---------------+-------------------------------------------------------+
| url-tags      | https://github.com/swarupdonepudi/gitr/tags           |
+---------------+-------------------------------------------------------+
| url-releases  | https://github.com/swarupdonepudi/gitr/releases       |
+---------------+-------------------------------------------------------+
| url-pipelines | https://github.com/swarupdonepudi/gitr/actions        |
+---------------+-------------------------------------------------------+
```

Here is a sample output of the `--dry` options passed to clone command

```shell
> gitr clone git@github.com:swarupdonepudi/gitr.git -d

+------------+---------------------------------------------------+
| remote     | git@github.com:swarupdonepudi/gitr.git            |
+------------+---------------------------------------------------+
| provider   | github                                            |
+------------+---------------------------------------------------+
| host       | github.com                                        |
+------------+---------------------------------------------------+
| repo-name  | gitr                                              |
+------------+---------------------------------------------------+
| create-dir | true                                              |
+------------+---------------------------------------------------+
| clone-path | /Users/swarupd/scm/github.com/swarupdonepudi/gitr |
+------------+---------------------------------------------------+
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

### Cleanup

```
brew uninstall gitr
brew untap swarupdonepudi/homebrew-gitr
```

### Contributions

I built `gitr` to share my passion for extreme productivity with other productivity geeks. Life is too short, and I ain't wasting time clicking and typing around that does not return any value. For those of you who share the same passion, I hope you find this project both useful and interesting. I am also pretty sure that productivity geeks are never content and will always look for more. So, if you see opportunities to improve this, I will be a happy man to see new issues and pull-requests.