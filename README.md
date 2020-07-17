<!--
*** Thanks for checking out this README Template. If you have a suggestion that would
*** make this better, please fork the repo and create a pull request or simply open
*** an issue with the tag "enhancement".
*** Thanks again! Now go create something AMAZING! :D
-->





<!-- PROJECT SHIELDS -->
<!--
*** I'm using markdown "reference style" links for readability.
*** Reference links are enclosed in brackets [ ] instead of parentheses ( ).
*** See the bottom of this document for the declaration of the reference variables
*** for contributors-url, forks-url, etc. This is an optional, concise syntax you may use.
*** https://www.markdownguide.org/basic-syntax/#reference-style-links
-->
[![Contributors](https://img.shields.io/github/contributors/othneildrew/Best-README-Template.svg?style=flat-square)](https://github.com/LabKom-Unila/health-checker/graphs/contributors)
[![Forks](https://img.shields.io/github/forks/othneildrew/Best-README-Template.svg?style=flat-square)](https://github.com/LabKom-Unila/health-checker/network/members)
[![Stargazers](https://img.shields.io/github/stars/othneildrew/Best-README-Template.svg?style=flat-square)](https://github.com/LabKom-Unila/health-checker/stargazers)
[![Issues](https://img.shields.io/github/issues/othneildrew/Best-README-Template.svg?style=flat-square)](https://github.com/LabKom-Unila/health-checker/issues)



<!-- PROJECT LOGO -->
<br />
<p align="center">
  <h3 align="center">Service Checker CLI</h3>

  <p align="center">
    Easy to use CLI tools to check dependant services before running your project!
    <br />
    <a href="https://github.com/LabKom-Unila/health-checker"><strong>Explore the docs »</strong></a>
    <br />
    <br />
    <a href="https://github.com/LabKom-Unila/health-checker">View Demo</a>
    ·
    <a href="https://github.com/LabKom-Unila/health-checker/issues">Report Bug</a>
    ·
    <a href="https://github.com/LabKom-Unila/health-checker/issues">Request Feature</a>
  </p>
</p>



<!-- TABLE OF CONTENTS -->
## Table of Contents

* [About the Project](#about-the-project)
  * [Built With](#built-with)
* [Getting Started](#getting-started)
  * [Installation](#installation)
* [Usage](#usage)
* [Roadmap](#roadmap)
* [Contributing](#contributing)
* [Contact](#contact)



<!-- ABOUT THE PROJECT -->
## About The Project

Sometimes it's annoying to check the system status of the services needed by your project, this cli project will simplify your way to check the services you need in your project.

Currently we have:
* MySQL's checker
* Redis's checker
* PostgreSQL's checker

Of course, we will add more to that list, we're waiting for your help.

### Built With
This project is build with love and :
* [Go](https://golang.org/)
* [Cobra](https://github.com/spf13/cobra)



<!-- GETTING STARTED -->
## Getting Started

Before you run and install this project, your go version should be higher than 1.10

### Installation

1. Clone the repo
```sh
git clone https://github.com/LabKom-Unila/health-checker.git
```
2. Install GO packages
```sh
go mod download
```
3. Build your go app
```sh
go build .
```
3. Add to your path



<!-- USAGE EXAMPLES -->
## Usage

Use flag -h or --help to know all the list command.

```sh
checker -h
```



<!-- ROADMAP -->
## Roadmap

See the [open issues](https://github.com/LabKom-Unila/health-checker/issues) for a list of proposed features (and known issues).



<!-- CONTRIBUTING -->
## Contributing

Contributions are what make the open source community such an amazing place to be learn, inspire, and create. Any contributions you make are **greatly appreciated**.

1. Fork the Project
2. Create your Feature Branch (`git checkout -b feature/AmazingFeature`)
3. Commit your Changes (`git commit -m 'Add some AmazingFeature'`)
4. Push to the Branch (`git push origin feature/AmazingFeature`)
5. Open a Pull Request




<!-- CONTACT -->
## Contact

Arham Abiyan - [@arham_abiyan](https://twitter.com/arham_abiyan) - arham.abiyan@gmail.com

Project Link: [https://github.com/LabKom-Unila/health-checker](https://github.com/LabKom-Unila/health-checker)

