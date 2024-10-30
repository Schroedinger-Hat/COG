<div align='center'>
  
[![Contributors][contributors-shield]][contributors-url]
[![Forks][forks-shield]][forks-url]
[![Stargazers][stars-shield]][stars-url]
[![Issues][issues-shield]][issues-url]

</div>

<!-- PROJECT LOGO -->
<br />
<div align="center">
  <h1>COG</h1>
  
  <a href="https://github.com/schroedinger-Hat/cog">
    <img src="public/sh.png" alt="Logo" width="80" height="80">
  </a>

  <p align="center">
    <br />
    <a href="https://github.com/schroedinger-Hat/cog/blob/main/README.md"><strong>Explore the docs »</strong></a>
    <br />
    <br />
    <a href="https://github.com/schroedinger-Hat/COG/issues/new?assignees=&labels=bug&projects=&template=bug_report.yml">Report Bug</a>
    ·
    <a href="https://github.com/schroedinger-Hat/COG/issues/new?assignees=&labels=feature+request%2Cneeds+discussion&projects=&template=feature_request.yml">Request Feature</a>
  </p>
</div>

<!-- TABLE OF CONTENTS -->
<details>
  <summary>Table of Contents</summary>
  <ol>
    <li>
      <a href="#about-the-project">About The Project</a>
      <ul>
        <li><a href="#tech-stack">Built With</a></li>
      </ul>
    </li>
    <li><a href="#usage">Usage</a></li>
    <li><a href="#contributing">Contributing</a></li>
    <li><a href="#contact">Contact</a></li>
  </ol>
</details>

<!-- ABOUT THE PROJECT -->

## About The Project

This is a project that we use to automate issues creation for vote talks of [Open Source Day](osday.dev). We have the call for speakers on sessionize, where we can download a csv file with talk+speaker name and to let the community vote we need all of them into issues.

The name *COG* is meant to lead back to the fact that this project is nothing more than a cog in a larger mechanism.

### Tech stack

The main language is [GO](https://go.dev/) and the main graphical library is [bubbletea](https://github.com/charmbracelet/bubbletea), nothing more.

## Usage

At the moment the fastest way to run this is to run:

```bash
go mod tidy
export GHTOKEN=<your-token>
go run main.go -csv <path-to-csv> -gh-user <user-repository> -gh-repository <repository>
```

The csv must have a similar structure of [this template](./template.csv):

```csv
name,description,labels
issue name,issue description,good first issue;bug;question
another issue name,another issue description,question
```

Is very important to split the labels with the separator *;* otherwise something unexpected might happen.

Last thing, the most important one: you need a [Github Personal Access Token](https://github.com/settings/tokens) no specific scope are needed.

## Contributing

Contributions are what make the open source community such an amazing place to learn, inspire, and create. Any contributions you make are **greatly appreciated**.

If you have a suggestion that would make this better, please fork the repo and create a pull request. You can also simply open an issue with the tag "enhancement".
Don't forget to give the project a star! Thanks again!

<!-- CONTACT -->

## Contact

Schrödinger's Hat Team - [@schroedinger_hat](mailto:dev@schroedinger-hat.org)

Project Link: [https://github.com/schroedinger-Hat/cog](https://github.com/schroedinger-Hat/cog)

<!-- MARKDOWN LINKS & IMAGES -->
<!-- https://www.markdownguide.org/basic-syntax/#reference-style-links -->

[contributors-shield]: https://img.shields.io/github/contributors/schroedinger-Hat/cog.svg?style=for-the-badge
[contributors-url]: https://github.com/schroedinger-Hat/cog/graphs/contributors
[forks-shield]: https://img.shields.io/github/forks/schroedinger-Hat/cog.svg?style=for-the-badge
[forks-url]: https://github.com/schroedinger-Hat/cog/network/members
[stars-shield]: https://img.shields.io/github/stars/schroedinger-Hat/cog?style=for-the-badge
[stars-url]: https://github.com/schroedinger-Hat/cog/stargazers
[issues-shield]: https://img.shields.io/github/issues/schroedinger-Hat/cog.svg?style=for-the-badge
[issues-url]: https://github.com/schroedinger-Hat/cog/issues
