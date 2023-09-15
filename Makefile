IMAGE_NAME=pkuphysu/meetplan
IMAGE_TAG=v1.0.0

docker:
	docker build . -t $(IMAGE_NAME):$(IMAGE_TAG)