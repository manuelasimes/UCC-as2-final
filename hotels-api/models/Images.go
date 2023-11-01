package model

type Image struct {
	Url string `bson:"image_url"`
}

type Images []Image