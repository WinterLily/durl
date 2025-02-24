#### 🦷 DURL: Cross-platform CLI URL Defanger 🐍

A simple Go-based CLI utility to sanitize and defang URLs for safe sharing in SOC environments.

This cleans up and IOC-ifies any URL and adds it to your clipboard 

##### Installation

1.  **Prerequisites:**
    *   Go installed and configured, if you want/need to build it yourself, otherwise, binaries for Linux and Windows are included. 

2.  **Clone the repository:**

    ```bash
    git clone <repository_url>
    cd durl
    ```

3.  **Build the project: (Optional)**

    Built executables are included for Windows and Linux, but if you want or need to build the project yourself: 

    ```bash
    go build -o durl main.go
    ```
4.  ** Install **

    To make the app globally available as a CLI utility, run the install scripts provided for Windows or Linux respectively. 

    Restart your terminal and you should be able to access the utility with **durl**


##### Usage

**Single URL**
Defangs a single URL
```bash
./durl <url>
```

**Multi-Mode**
Defangs multiple URLs (space-separated)
```bash
./durl <url1> <url2> <url3>
```

