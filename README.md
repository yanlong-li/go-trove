# Trove

[中文](./docs/README.zh-hans.md)

Trove is a package dependency manager for the Go language. Compiled with the latest `v1.12.9'SDK, the minimum requirement version is not tested.
Trove is an experimental project and has no funds to build a separate package management service. At present, it uses Git warehouse management to test compatibility with GitHub warehouse. It theoretically supports all GIT code warehouses (not tested) and saves the cost of individual servers.
Trove can introduce private warehouse code packages at will (with warehouse cloning privileges) and refer to private warehouse code at will without additional configuration and construction of private servers.
Trove is based on Git, so make sure that your runtime environment has Git installed and that you can use the GIT command globally
## Official website
Https://trove.daohang.dev
## Development Language
Since it's a package management tool for Go, it's natural to develop it in Go language.
## Open Source Protocol
Publish Open Source Based on [MulanPSL](http://license.coscl.org.cn/MulanPSL) License Agreement
## Usage
#### Initialization
    trove init
#### Reference dependency
    Trove require https://github.com/XXXXXX [commitId]
    By default, commit version control is equivalent to the following
    Trove require commit@https://github.com/XXXXXX
    [commitId] Optional commitId version, default to the latest version
    Can be set to tag tag control version
    Trove require tag@https://github.com/XXXXXX [tag]
    [tag] optional tag version, default to the latest version
#### Remove dependencies
    Trove remove [packageName]
    [packageName] package name
#### View the dependency list
    Trove --list
    Viewable project direct dependency list
    Trove -- List -- all
    You can view a list of all dependencies on a project, including direct and indirect dependencies
#### Install all dependency packages
    Trove install [packageName]
    When manually editing trove.json to increase dependencies, you can use this command to install the dependency packages that are not downloaded, or to download all dependency packages when initializing the project
    [packageName] Optional installation of separate dependency packages
#### Update dependency packages
    Trove update [packageName]
    Update all dependency packages to the latest version within version limits
    [packageName] Optionally updates only the specified package
#### Trove version
    Trove -V,--version
    Check Trove version number and final revision time
#### Help
    Trove -h,--help
    View Trove Help Information
## Development Team
> No ranking

[Yanlong-li] (https://github.com/yanlong-li)

## Update logs

    September 3, 2019
    Reporting Error Problem Caused by Initialization of Fixed Update Operations Not Assigned to Zero
    Repair require operation. lock file exists, but packages parameter is nil time error problem
    Fixed parameter error when writing dependency packages
    Adding Removal Recursive Removal Dependencies
    Adjust. lock to continue traovpackage to increase use count parameters
    Modified Regression Processing Unupdated Use Count Problem
    Optimize deletion of current project write.lock
    entry name
    Fix Dependency Item Not Written. Lock Problem
    Optimizing Citation Dependent Time Chain
    Updated version 0.0.1.13
    
    September 3, 2019
    Increasing recursive dependency processing
    Reorganization of install, update, require logic
    Remove the install command, and overlap the update function
    
    September 3, 2019
    Increase init command to initialize project
    
    September 3, 2019
    Optimizing version control structure
    Fixed the problem of not switching to the specified version when updating a single package
    
    September 2, 2019
    The new git version control support Tag tag@https://github.com/XXXXX/XXXXXX defaults to commit@https://github.com/XXXXXX
    Optimizing git@ssh import package
    
    September 2, 2019
    Enhanced command line
    Add require command to introduce git package
    New install installs packages that do not exist
    Added update update package command
    Add a list to display all incoming packages
    Add - v, - version output version number
    Add - h, - help output help hints
    New remove removal package