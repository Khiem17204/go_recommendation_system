#!/bin/bash

# Set default values
REGISTRY="go-recommendation-system"
VERSION=$(git rev-parse --short HEAD)

# Colors for output
GREEN='\033[0;32m'
RED='\033[0;31m'
YELLOW='\033[1;33m'
NC='\033[0m' # No Color

# Function to print usage
usage() {
    echo -e "${YELLOW}Usage: $0 <path-to-dockerfile>${NC}"
    echo -e "Example: $0 /path/to/service/Dockerfile"
    exit 1
}

# Function to check if Docker is running
check_docker() {
    if ! docker info > /dev/null 2>&1; then
        echo -e "${RED}Docker is not running. Please start Docker first.${NC}"
        exit 1
    fi
}

# Function to validate dockerfile path
validate_dockerfile() {
    local dockerfile_path=$1
    if [ ! -f "$dockerfile_path" ]; then
        echo -e "${RED}Error: Dockerfile not found at '$dockerfile_path'${NC}"
        exit 1
    fi
}

# Function to build and push a service
build_service() {
    local dockerfile_path=$1
    local service_name=$(basename "$(dirname "$dockerfile_path")")
    local dockerfile_dir=$(dirname "$dockerfile_path")

    echo -e "${GREEN}Building ${service_name} service using ${dockerfile_path}...${NC}"

    cd "$(git rev-parse --show-toplevel 2>/dev/null || pwd)"

    local image_name="${REGISTRY}/${service_name}:${VERSION}"
    local image_tag="${REGISTRY}/${service_name}:latest"
    docker build -t "$image_name" -f "$dockerfile_path" .
    docker tag "$image_name" "$image_tag"
    if [ $? -ne 0 ]; then
        echo -e "${RED}Failed to build ${service_name} image${NC}"
        exit 1
    fi

    echo -e "${GREEN}Successfully built ${service_name} image${NC}"

    # Get current kubectl context
    CONTEXT=$(kubectl config current-context 2>/dev/null)

    echo -e "${YELLOW}Current kubectl context: ${CONTEXT}${NC}"

    case "$CONTEXT" in
        minikube)
            echo -e "${GREEN}Loading ${service_name} image into Minikube...${NC}"
            minikube image load "$image_name"
            if [ $? -eq 0 ]; then
                echo -e "${GREEN}Successfully loaded image into Minikube${NC}"
            else
                echo -e "${RED}Failed to load image into Minikube${NC}"
                exit 1
            fi
            ;;
        orbstack)
            echo -e "${YELLOW}OrbStack shares Docker daemon with host. No need to load image manually.${NC}"
            ;;
        *)
            echo -e "${YELLOW}Unknown or unsupported context: '${CONTEXT}'. Skipping image load.${NC}"
            ;;
    esac
}

# Main script
main() {
    if [ $# -ne 1 ]; then
        usage
    fi

    local dockerfile_path=$1

    check_docker
    validate_dockerfile "$dockerfile_path"
    build_service "$dockerfile_path"

    echo -e "${GREEN}Service has been built and processed successfully!${NC}"
}

main "$@"
