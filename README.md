# CourseHub

General Course Inquiry Engine implements using Golang and Gin.

## Database
0. Supports MySQL and MariaDB (using GORM)
1. Create a configuration file call `config.toml`, an example file is provided.
2. setup `PATH_TO_CONFIG` in your system environment, points to the **Absolute directory** to the config file.
    - e.g. `PATH_TO_CONFIG=/home/ecwu/coursehub/config.toml`