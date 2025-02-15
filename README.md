# Flean
### IMPORTANT: Flutter needs to be globally installed as Flean uses it internally for project creation
Flean is an open source flutter clean app architecture generation and management tool that helps developers manage their flutter
project seamlessly and enhance the flutter development experience.
It also facilitates managing verbose blocs and cubits within the project just with basic cli commands making project management less annoying.

## Go Libraries and Tools used
- [cobra-cli](https://github.com/spf13/cobra)
- [bubbletea](https://github.com/charmbracelet/bubbletea)
- [lipgloss](https://github.com/charmbracelet/lipgloss)

## Overview
- [Motivation](#motivation)
- [Installation Guide](#installation-guide)
- [Usage](#usage)
- [Future Plans](#future-plans)
- [Contribution](#contribution)

## Motivation
While making some Flutter projects in Clean Architecture with Blocs and Cubits for state management, I realized things can become verbose pretty quickly as the project grows.

Using app architectures makes the code more manageable and scalable but at the same time managing the architecture itself becomes a pain. Even though some code editor and IDE extensions exist, none of them provide a complete comprehensive and intuitive solution to managing flutter projects with ease.

That's when the idea to make a simple and intuitive CLI tool to manage my projects came into being!  
Flean was initially just a thought to manage Flutter Clean Architecture projects (which is where it gets it's name from) but this project definitely has more room for flexibility and growth
which can make it a one stop solution for all flutter project management needs.
## Installation Guide
Follow these steps to install flean:

### Windows
1. Download the latest binary from the [Releases](https://github.com/yourusername/yourproject/releases) page.
   
2. Create a new folder (e.g., `C:\flean`) and move the downloaded binary into this folder.

3. Add the folder to your system's environment variables:
   - Open **Start Menu** and search for `Environment Variables`.
   - Click **Edit the system environment variables**.
   - In the **System Properties** window, click **Environment Variables**.
   - Under **System variables**, find and select `Path`, then click **Edit**.
   - Click **New**, then enter the path to the folder you created (e.g., `C:\flean`).
   - Click **OK** on all windows to save the changes.
5. Verify the installation by opening **Command Prompt** and running:
   ```sh
   flean
   ```
### macOS
Follow these steps to install and set up the project on macOS:

1. Download the latest macOS binary from the [Releases](https://github.com/TanishqPokharia/flean/releases) page.
2. Move the binary to `/usr/local/bin`:
   ```sh
   mv flean-mac /usr/local/bin/flean
   chmod +x /usr/local/bin/flean
   ```
3. Verify the installation by opening **Terminal** and running:
   ```sh
   flean
   ```

### Linux
Follow these steps to install and set up the project on Linux:

1. Download the latest Linux binary from the [Releases](https://github.com/TanishqPokharia/yourproject/releases) page.
2. Move the binary to `/usr/local/bin`:
   ```sh
   mv flean-linux /usr/local/bin/flean
   chmod +x /usr/local/bin/flean
   ```
3. Verify the installation by opening **Terminal** and running:
   ```sh
   flean
   ```


## Usage
### Available Commands
  - [Bloc](#bloc)
  - [Create](#create)
  - [Cubit](#cubit)
  - [Feature](#feature)
  - [Help](#help)


#### Bloc
Manage BLoCs inside your application.
```sh
flean bloc add auth
flean bloc remove auth
```


https://github.com/user-attachments/assets/86828c5f-3981-4982-a843-590c0ea55f90


When working with feature first clean architecture, -f or -feature flag can be used to create a bloc inside a specific feature


https://github.com/user-attachments/assets/828a6739-5183-4ed7-90ee-226346d5ca27



#### Create
Create a new Flutter project with the desired architecture.
1. Select project architecture
2. Choose commonly used dependencies or skip by just pressing f
```sh
flean create my_app
```


https://github.com/user-attachments/assets/8b7863da-152e-4c33-9ced-95444cb234ce



#### Cubit
Manage Cubits inside your application.
This command works exactly like the bloc command but to manage cubits.
```sh
flean cubit add auth
flean cubit remove auth
```

#### Feature
Manage features in a feature-first architecture Flutter app.
```sh
flean feature add profile
flean feature remove profile
```




https://github.com/user-attachments/assets/47344184-ccd2-49ce-864c-87e8b6630544





#### Help
Display help about any command.
```sh
flean help [command]
```

Use `flean [command] --help` for more details on each command.

## Future Plans
1. Flean configurations can be customized using flean.yaml file which unlocks a possibility for highly customizable flutter project management
2. More state management tools other than bloc and cubits can be configured using Flean due to its flexibility

## Contribution
Contributions are welcomed whole-heartedly! Follow these steps to contribute:

1. Fork the repository.
2. Create a new branch:
   ```sh
   git checkout -b feature-branch
   ```
3. Commit your changes:
   ```sh
   git commit -m "Add new feature"
   ```
4. Push to your branch:
   ```sh
   git push origin feature-branch
   ```
5. Open a Pull Request.

Feel free to raise issues or contribute to help this project grow.
