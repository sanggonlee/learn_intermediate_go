# [Learn Intermediate Go] Go Modules
Go’s dependency management has evolved over time like other languages. But ever since Go Modules was introduced in v1.11, the community has quickly adapted it as its primary dependency management system.
If you are relatively new to Go, you might have been using modules with some gap on how they really work. I will try my best to fill some of those gaps here.
In general, Go module keeps track of the dependencies in files named go.mod and go.sum. go.sum is automatically generated and derived from go.mod, so you typically never need to do anything directly to that file. If you're familiar with Node.js, go.mod is comparable to package.json and go.sum is comparable to package-lock.json.

## Creating a module
To have an example that we want to play around with, let’s start with creating a module. To create a module, you use go mod init :

Let’s pause for a second and understand what this really does.
go mod init part should be self-explanatory - it initializes the current directory as a Go module directory. In fact, all it does is create a file named go.mod in the current directory with the following content:
module github.com/sanggonlee/learn_intermediate_go/01_modules
go 1.15
Which are, the name of your module and the Go version.
Now let’s look at the later part: github.com/sanggonlee/learn_intermediate_go/01_modules
Your initial guess might be it’s a path or a URL. The answer to that is yes and no.
No, because it’s just the name of the module. You can name it whatever you want. BUT also
Yes, you should almost always name it as the repository path that you will upload your code to. This is because when other Go projects need to import your module, they will use this name to find your module.

## Adding dependencies
To install a specific dependency, use go get:

For version query, in addition to the five options there are also special strings: latest, upgrade, patch, and none.
none is especially special in that running with this will remove the dependency.

## When something looks off
go mod tidy command is your friend. It fixes most of inconsistencies in go.mod and go.sum files, and updates them in canonical form.
A common scenario is you modified the go.mod file (maybe because you added a new dependency) and your teammate also modified it, so you're in a situation having to resolve merge conflicts. Simply resolve the conflict in go.mod file, delete go.sum file, and run go mod tidy.

## Closer look into go.mod file
Now let’s have a look inside the go.mod file. There are actually only 5 directives you can use in go.mod file, so it's quite simple!

NOTE: retract is a new feature introduced in Go 1.16.

## Minimum Version Selection (MVS)
Dependency tree is simple to understand, but it’s less simple when considering all the versions of the dependencies.
Go uses Minimum Version Selection (MVS) to select exactly which dependency versions to use.
First, all the dependencies imported by your module are found recursively (i.e. dependencies of those dependencies, and dependencies of them, etc).
And then the latest version for each dependency is used to build the final graph.

In the above example, your module imports Dependency A v1.2 and B v2.1. But B v2.1 imports A v1.3, and A v1.3 imports C v3.4. So even though your module imports A v1.2 and B v2.1, in the final build list these are used (highlighted with blue colour):
A v1.3
B v2.1
C v3.4

## Vendoring
Vendoring is basically downloading all the dependencies’ source code into your (or whatever the build environment’s) file system. It’s similar to how node_modules works in Node.js.
By default, vendoring is disabled and the dependencies are downloaded to the module cache. You might want to use vendoring for cases like:
for interoperation with older versions of Go (remember the Go modules were introduced only in v1.11!)
your module imports private projects and you don’t want to inject your git credentials into the build environment (e.g. CI/CD)
if you feel it’s more reliable that way (it could be more reliable since you have everything you need in an earlier stage)
In order to vendor, run go mod vendor command, which will create a vendor directory and put all the dependencies source code there.
Your module will use the code in the vendor directory, so whenever the go.mod file changes, you should run go mod vendor to sync the vendor with the Go modules.

## Closing remark
I hope that was helpful. Thank you for reading! (Oh btw, did you know you can clap up to 50 times to a Medium article? 😄)