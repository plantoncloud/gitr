# gitr(git rapid): the missing link between git cli and web browsers

**tl;dr** in short, gitr does this:

if you are in iterm inside a subdirectory of a git repo, and you want to see the pull requests on that repo?

*before gitr:*

1. open web browser
2. go to gitlab.com
3. search for the repo or visually locate the repo on the home page
4. click on the repo
5. after the home page is loaded then locate the icon on the page to go to the pull requests section and click on it

*with gitr:*

in iterm, from any subdirectory under any git repo

1. gitr prs

> same for tags, releases, pipelines, issues, branches, commits etc...

`gitr` reduces a ton of clicking and waiting for the git repo web pages to load by **taking you directly** to the page
that you would like to see right from the command line. You may think that it's not a lot of waiting but trust me, you
will notice the difference once you start using `gitr`.

`gitr` relies on the contents of `.git` folder and combines it with the provider(think of gitlab, github and bitbucket)
knowledge that is built into it to smartly navigate you to the right web page of your repo right from the command line.

`gitr` can run on linux, windows and mac systems

### install

#### mac

```
brew tap swarupdonepudi/homebrew-gitr
brew install gitr
```

#### windows & linux

`gitr` binaries can be downloaded directly from
the [releases section of this repo](https://github.com/swarupdonepudi/gitr/releases)

### scm providers

`gitr` currently, supports the following scm providers. We can always add support for more if there is demand.

* github
* gitlab
* bitbucket

### features

`gitr` has two features

1. gitr web - open a repo and different parts of a repo in web browser from command line
2. gitr clone - this features makes it possible to organize git repos cloned from different scm providers and also
   retain their hierarchy on the scm provider on laptops which is not possible with the default `git clone <clone-url>`.
   This is particularly useful for gitlab as it supports a nested hierarchy.

#### gitr web

`gitr` opens the following parts of your git repo in the web browser right from the command line

> note: the below commands will work only when executed from root or subdirectory of any git repo

| command         |  description                                                    |
|-----------------|-----------------------------------------------------------------|
| gitr web        |  opens the home page of the repo in the browser                 |
| gitr rem        |  opens the local checkout branch of the repo in the browser     |
| gitr prs        |  opens the prs/mrs of the repo in the browser                   |
| gitr pipe       |  opens the pipelines/actions of the repo in the browser         |
| gitr issues     |  opens the issues of the repo in the browser                    |
| gitr releases   |  opens the releases of the repo in the browser                  |
| gitr tags       |  opens the tags of the repo in the browser                      |
| gitr commits    |  opens the commits of the local branch of repo in the browser   |
| gitr branches   |  opens the branches of the repo in the browser                  |

#### gitr clone

This may not seem like a very useful feature at first. If you would like every git repo that you clone to land in a
deterministic location then use `gitr clone <clone-url>` instead of `git clone <clone-url>`.

By default `gitr clone` does the same exact thing as `git clone`.

To mimic the folder structure that your repo is on the scm then use `-c` flag.

Running `gitr clone git@gitlab.mycompany.net:parent/subgroup1/subgroup2/repo.git -c` command from `~/scm` folder, the
repo is cloned to `~/scm/subgroup1/subgroup2/repo` location.

Adding the below configuration to `~/.gitr.yaml`, every repo cloned using `gitr` will be cloned to the location that
mimics scm provider.

```yaml
clone:
  scmHome: /Users/swarup/scm
  includeHostForCreDir: false
  alwaysCreDir: true
## more config
```

With above config `gitr` will clone all repos to `scmHome` location, regardless of where you
run `gitr clone <clone-url>` command.

Because `includeHostForCreDir` is set to `true`, `gitr` will clone the repo to a folder with the name of the hostname of
the scm under `scmHome`.

Because `alwaysCreDir` is set to `true`, `gitr` will clone the repo to the same path as that is in the `<clone-url>`.

example:

```shell
gitr clone git@gitlab.mycompany.net:parent/subgroup1/subgroup2/repo.git
```

The repo gets cloned to `{scmHome}/{scmHostname}/{repoPath}/{repoName}`
i.e `/Users/swarup/scm/gitlab.mycompany.net/parent/subgroup1/subgroup2/repo` location.

Below is a snapshot of a nicely organized directory structure for all repos from different scm systems.

```shell
> tree ~/scm

gitlab.com
└── swarup
    └── group1
        └── app-1
            ├── Dockerfile
            ├── Makefile
            ├── README.md
            ├── main.go
    └── group2
        └── sub-group
            └──app-2
               ├── Dockerfile
               ├── Makefile
               ├── README.md
               ├── main.go       
github.com
└── swarupdonepudi
    ├── gitr
    │ ├── go.mod
    │ ├── lib
    │ ├── lib_test
    │ └── main.go
    └── homebrew-gitr
        ├── Formula
        └── README.md
    ahmetb
    └── kubectx
        ├── kubens
        ├── kubectx
        └── README.md
```

### .gitr.yaml config file

Below is the config options supported in `~/.gitr.yaml`

| config                      |  default  | description                                                                                                                             |
|-----------------------------|-----------|-----------------------------------------------------------------------------------------------------------------------------------------|
| clone.scmHome               |     ""    |  if this value is set, then gitr clone will always clone the repos to this path, regardles of where you run `gitr clone` command from   |
| clone.alwaysCreDir          |     false |  if this is set to true, then gitr clone will always create the directories present in the clone url                                    |
| clone.includeHostForCreDir  |     false |  if this is set to true, then gitr clone will always prefix the hostname to the clone path                                              |
| scmSystems.[].hostname      |     ""    |  hostname of the on-prem deployment of scm system                                                                                       |
| scmSystems.[].provider      |     ""    |  provider of the on-prem deployment. allowed values are github, gitlab and bitbucket                                                    |
| scmSystems.[].defaultBranch |     ""    |  this is the value of the default branch configured on the scm                                                                          |

#### example config file

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

### dry Run

Currently `gitr` only supports `github`, `gitlab` and `bitbucket`. Each provider has slight differences in the urls to
access different parts of a repo. `gitr` could just work out of the box for a different cloud provider ex: AWS
CodeCommit or it may not behave as expected.

So, if `gitr` does not work as you expect it to, both for supported providers and other providers, you will be able to
see what urls `gitr` produces but using `--dry` option. This option is available for both web and clone features.

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

### supports on-prem deployments

`gitr` can work with on-prem deployments of the supported scm providers i.e github, gitlab and bitbucket.

You need to help `gitr` figure out what SCM system you are using. You can do so by simply creating `~/.gitr.yaml` file
and adding *your* SCM hostname and SCM Provider to the config file.

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

### cleanup

```
brew uninstall gitr
brew untap swarupdonepudi/homebrew-gitr
```

### contributions

I built `gitr` to share my passion for extreme productivity with other productivity geeks. Life is too short, and I
ain't wasting time clicking and typing around that does not return any value. For those of you who share the same
passion, I hope you find this project both useful and interesting. I am also pretty sure that productivity geeks are
never content and will always look for more. So, if you see opportunities to improve this, I will be a happy man to see
new issues and pull-requests.
