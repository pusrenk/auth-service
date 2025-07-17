#!/bin/bash
DOCKER_USERNAME="rayprastya"
IMAGE_NAME="pusrenk"
TAG="auth-service"

# Colors for output
RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
NC='\033[0m' # No Color

echo -e "${YELLOW}Building Docker image...${NC}"
docker build -t ${DOCKER_USERNAME}/${IMAGE_NAME}:${TAG} .

if [ $? -eq 0 ]; then
    echo -e "${GREEN}Build successful!${NC}"
    
    echo -e "${YELLOW}Tagging image...${NC}"
    docker tag ${DOCKER_USERNAME}/${IMAGE_NAME}:${TAG} ${DOCKER_USERNAME}/${IMAGE_NAME}:${TAG}
    
    echo -e "${YELLOW}Pushing to Docker Hub...${NC}"
    docker push ${DOCKER_USERNAME}/${IMAGE_NAME}:${TAG}
    
    if [ $? -eq 0 ]; then
        echo -e "${GREEN}Successfully pushed to Docker Hub!${NC}"
        echo -e "${GREEN}Image: ${DOCKER_USERNAME}/${IMAGE_NAME}:${TAG}${NC}"
    else
        echo -e "${RED}Failed to push to Docker Hub${NC}"
        exit 1
    fi
else
    echo -e "${RED}Build failed!${NC}"
    exit 1
fi 