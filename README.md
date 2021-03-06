<!--
<!-- ALL-CONTRIBUTORS-BADGE:START - Do not remove or modify this section -->
[![All Contributors](https://img.shields.io/badge/all_contributors-1-orange.svg?style=flat-square)](#contributors-)
<!-- ALL-CONTRIBUTORS-BADGE:END -->





<!-- PROJECT SHIELDS -->
<!--
*** I'm using markdown "reference style" links for readability.
*** Reference links are enclosed in brackets [ ] instead of parentheses ( ).
*** See the bottom of this document for the declaration of the reference variables
*** for contributors-url, forks-url, etc. This is an optional, concise syntax you may use.
*** https://www.markdownguide.org/basic-syntax/#reference-style-links
-->
[![Forks](https://img.shields.io/github/forks/LabKom-Unila/health-checker?label=Fork&style=social)](https://github.com/LabKom-Unila/health-checker/network/members)
[![Stargazers](https://img.shields.io/github/stars/LabKom-Unila/health-checker?style=social)](https://github.com/LabKom-Unila/health-checker/stargazers)
[![Issues](https://img.shields.io/github/issues/LabKom-Unila/health-checker?style=flat-square)](https://github.com/LabKom-Unila/health-checker/issues)



<!-- PROJECT LOGO -->
<br />
<p align="center">
  <h3 align="center">Health Checker CLI</h3>

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
4. Add to your path



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

Project Link: [https://github.com/LabKom-Unila/health-checker](https://github.com/LabKom-Unila/health-checker)


## Contributors ✨

Thanks goes to these wonderful people ([emoji key](https://allcontributors.org/docs/en/emoji-key)):

<!-- ALL-CONTRIBUTORS-LIST:START - Do not remove or modify this section -->
<!-- prettier-ignore-start -->
<!-- markdownlint-disable -->
<table>
  <tr>
    <td align="center"><a href="https://github.com/arham09"><img src="https://avatars0.githubusercontent.com/u/18678301?v=4" width="100px;" alt=""/><br /><sub><b>Arham Awal Abiyan</b></sub></a><br /><a href="https://github.com/LabKom-Unila/health-checker/commits?author=arham09" title="Code">💻</a></td>
  </tr>
</table>

<!-- markdownlint-enable -->
<!-- prettier-ignore-end -->
<!-- ALL-CONTRIBUTORS-LIST:END -->

This project follows the [all-contributors](https://github.com/all-contributors/all-contributors) specification. Contributions of any kind welcome!