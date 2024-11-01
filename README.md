<!-- PROJECT LOGO -->
<p align="center">
  <!-- <a href="https://github.com/berbyte/share">
    <img src="images/logo.png" alt="Logo" width="80" height="80">
  </a> -->

  <h1 align="center">Share</h1>

  <p align="center">

> [!NOTE]
>
>  A command-line tool written in Go to upload files to a Tigris bucket. If no file is provided, it reads from standard input and names the file using a timestamp. **The resulting URL is copied to the clipboard.**
>


<!-- USAGE EXAMPLES -->
> [!IMPORTANT]
>
> ## Usage
>
> ### Upload a File
> To upload a specific file to the bucket:
>
> ```sh
> share filename.txt
> ```
>
> ### Upload from Standard Input
> To upload from standard input:
>
> ```sh
> echo "your content here" | share
> ```
> This will create a file named with the current timestamp, e.g., `1730344643.txt`.
>
> ### Taco Bell Programming
> ```sh
> awk '{ print $1 }' access.log | sort | uniq -c | sort -nr | head -10 | share
> ```
>
> ### Output
> After uploading, the URL of the uploaded file will be copied to the clipboard, and also printed in the terminal.
>


<br />
<a href="https://github.com/berbyte/share"><strong>Explore the docs »</strong></a>
<br />
<br />
<a href="https://github.com/berbyte/share">View Demo</a>
·
<a href="https://github.com/berbyte/share/issues">Report Bug</a>
·
<a href="https://github.com/berbyte/share/issues">Request Feature</a>

  </p>
</p>

<hr />

[![Contributors][contributors-shield]][contributors-url]
[![Forks][forks-shield]][forks-url]
[![Stargazers][stars-shield]][stars-url]
[![Issues][issues-shield]][issues-url]
[![Apache-2.0 license][license-shield]][license-url]

<hr />


<!-- TABLE OF CONTENTS -->
<details>
  <summary>Table of Contents</summary>
  <ol>
    <li>
      <a href="#about-the-project">About The Project</a>
      <ul>
        <li><a href="#built-with">Built With</a></li>
      </ul>
    </li>
    <li><a href="#usage">Usage</a></li>
    <li>
      <a href="#getting-started">Getting Started</a>
      <ul>
        <li><a href="#prerequisites">Prerequisites</a></li>
        <li><a href="#installation">Installation</a></li>
      </ul>
    </li>
    <li><a href="#roadmap">Roadmap</a></li>
    <li><a href="#contributing">Contributing</a></li>
    <li><a href="#license">License</a></li>
    <li><a href="#contact">Contact</a></li>
    <li><a href="#acknowledgments">Acknowledgments</a></li>
  </ol>
</details>


<!-- ABOUT THE PROJECT -->
## About The Project

Share is a command-line tool that allows you to upload files to a Tigris bucket. If no file is provided, it reads from standard input and names the file using a timestamp. The resulting URL is copied to the clipboard.

### Features
- Upload files to a tigris bucket.
- Overwrite existing files.
- Automatically set expiration to 10 days.
- Copy URL to the clipboard after successful upload.

### Built With
- [Go](https://go.dev/)
- [AWS SDK for Go](https://aws.amazon.com/sdk-for-go/)



<!-- GETTING STARTED -->
## Getting Started

To get a local copy up and running follow these simple steps.

### Prerequisites

- Go installed (1.18+ recommended)
- Bucket credentials with appropriate permissions.
- Environment variables configured:
  - `BUCKET_NAME`
  - `AWS_REGION`
  - `AWS_ENDPOINT_URL_S3`
  - `AWS_ACCESS_KEY_ID`
  - `AWS_SECRET_ACCESS_KEY`
  - `DOMAIN`

### Installation

#### Build the Project
Clone the repository and build the project using the following commands:

```sh
git clone https://github.com/berbyte/share.git
cd share
go build -o share
```



<!-- CONTRIBUTING -->
## Contributing

Contributions are what make the open source community such an amazing place to learn, inspire, and create. Any contributions you make are **greatly appreciated**.

If you have a suggestion that would make this better, please fork the repo and create a pull request. You can also simply open an issue with the tag "enhancement".
Don't forget to give the project a star! Thanks again!

1. Fork the Project
2. Create your Feature Branch (`git checkout -b feature/AmazingFeature`)
3. Commit your Changes (`git commit -m 'Add some AmazingFeature'`)
4. Push to the Branch (`git push origin feature/AmazingFeature`)
5. Open a Pull Request


<!-- LICENSE -->
## License

Distributed under the Apache-2.0 License. See `LICENSE.txt` for more information.


<!-- CONTACT -->
## Contact

BER Team - [@berbyte](https://github.com/berbyte) - github@ber.sh

Project Link: [https://github.com/berbyte/share](https://github.com/berbyte/share)


<!-- ACKNOWLEDGMENTS -->
## Acknowledgments

- [othneildrew's Best-README-Template](https://github.com/othneildrew/Best-README-Template)
- [AWS SDK for Go](https://aws.amazon.com/sdk-for-go/)


<!-- MARKDOWN LINKS & IMAGES -->
[contributors-shield]: https://img.shields.io/github/contributors/berbyte/share.svg?style=for-the-badge
[contributors-url]: https://github.com/berbyte/share/graphs/contributors
[forks-shield]: https://img.shields.io/github/forks/berbyte/share.svg?style=for-the-badge
[forks-url]: https://github.com/berbyte/share/network/members
[stars-shield]: https://img.shields.io/github/stars/berbyte/share.svg?style=for-the-badge
[stars-url]: https://github.com/berbyte/share/stargazers
[issues-shield]: https://img.shields.io/github/issues/berbyte/share.svg?style=for-the-badge
[issues-url]: https://github.com/berbyte/share/issues
[license-shield]: https://img.shields.io/github/license/berbyte/share.svg?style=for-the-badge
[license-url]: https://github.com/berbyte/share/blob/master/LICENSE.txt
