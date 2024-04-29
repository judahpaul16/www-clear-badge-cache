## Systemd Service File 📜
```bash
[Unit]
Description=Docker Container for Clear-Badge_Cache.com Go App
After=docker.service network-online.target
Requires=docker.service
Wants=network-online.target

[Service]
Type=simple
User=root
# Ensure any existing container is removed before starting
ExecStartPre=-/usr/bin/docker stop clear-badge-cache
ExecStartPre=-/usr/bin/docker rm clear-badge-cache
ExecStart=/usr/bin/docker run --rm --name clear-badge-cache --env-file /opt/www-clear-badge-cache/.env -p 8080:8080 -v /opt/www-clear-badge-cache:/mnt clear-badge-cache:latest
ExecStop=/usr/bin/docker stop clear-badge-cache

Restart=always
RestartSec=5s

[Install]
WantedBy=multi-user.target
```

## NGINX Configuration 🌐
```nginx
# Redirect HTTP to HTTPS
server {
    server_name "clear-badge-cache.com";
    add_header X-Frame-Options "ALLOW-FROM clear-badge-cache.com";
    add_header Content-Security-Policy "frame-ancestors clear-badge-cache.com";

    listen [::]:443 ssl ipv6only=on; # managed by Certbot
    listen 443 ssl; # managed by Certbot
    ssl_certificate /etc/letsencrypt/live/clear-badge-cache.com/fullchain.pem; # managed by Certbot
    ssl_certificate_key /etc/letsencrypt/live/clear-badge-cache.com/privkey.pem; # managed by Certbot
    include /etc/letsencrypt/options-ssl-nginx.conf; # managed by Certbot
    ssl_dhparam /etc/letsencrypt/ssl-dhparams.pem; # managed by Certbot

    # Static files
    location /static/ {
	alias /opt/www-clear-badge-cache/static/;
    }

    # Proxy pass configuration
    location / {
	proxy_intercept_errors on;
	proxy_set_header Host $http_host;
	proxy_set_header X-Real-IP $remote_addr;
	proxy_set_header X-Forwarded-For $proxy_add_x_forwarded_for;
	proxy_set_header X-Forwarded-Proto $scheme;
	proxy_connect_timeout   600;
	proxy_send_timeout      600;
	proxy_read_timeout      600;
        proxy_pass http://127.0.0.1:8080;
    }
}

server {
    if ($host = clear-badge-cache.com) {
        return 301 https://$host$request_uri;
    } # managed by Certbot


    listen 80;
    listen [::]:80;
    server_name "clear-badge-cache.com";
    return 404; # managed by Certbot
}
```

## Deployment Script 🚀
```bash
#!/bin/bash

# Stop the service
sudo systemctl stop clear-badge-cache

# Define the repository and container details
REPO_URL="https://github.com/judahpaul16/www-clear-badge-cache"
REPO_FOLDER="www-clear-badge-cache"
CONTAINER_NAME="clear-badge-cache"
IMAGE_NAME="clear-badge-cache"
PORT_MAPPING="8080:8080"

# Remove existing repository folder if it exists
if [ -d "${REPO_FOLDER}" ]; then
    echo "Deleting existing repository folder..."
    rm -rf ${REPO_FOLDER}
fi

# Clone the repository
echo "Cloning repository..."
git clone ${REPO_URL}
echo "Repository cloned."

# Copy the .env file into the repository folder
echo "Copying .env file..."
touch .env
sudo cp .env ${REPO_FOLDER}/.env
echo ".env file copied."

cd ${REPO_FOLDER}

# Clean up Docker system to free up space
echo "Cleaning up Docker system..."
docker system prune -f
echo "Docker system cleaned."

# Remove any existing container with the same name
echo "Checking for existing container named ${CONTAINER_NAME}..."
if docker ps -q -f name=^/${CONTAINER_NAME}$; then
    echo "Stopping and removing existing container..."
    docker stop ${CONTAINER_NAME} && docker rm --force ${CONTAINER_NAME}
    echo "Existing container stopped and removed."
fi

# Ensure port 8080 is free by killing any processes using it
echo "Ensuring port 8080 is free..."
sudo fuser -k 8080/tcp

# Build the Docker image
echo "Building Docker image..."
docker build -t ${IMAGE_NAME} .

# Run the Docker container
echo "Running Docker container..."
docker run -d --name ${CONTAINER_NAME} \
           --env-file .env \
           -p ${PORT_MAPPING} \
           -v "$(pwd)":/app \
           ${IMAGE_NAME}

echo "Docker container ${CONTAINER_NAME} is now running."

# Start the service
sudo systemctl start clear-badge-cache
```