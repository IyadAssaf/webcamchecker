package webcamchecker

import (
	"context"
	"github.com/stretchr/testify/assert"
	"gocv.io/x/gocv"
	"testing"
)

func TestGetWebcamIsOn(t *testing.T) {
	wc, err := gocv.OpenVideoCapture(0)
	if err != nil {
		t.Fatal(err)
	}
	defer wc.Close()

	isOn, err := IsWebcamOn(context.Background())
	assert.Nil(t, err)
	assert.True(t, isOn)
}

func TestGetWebcamIsOff(t *testing.T) {

	// wait for previous test to turn off the webcam properly
	for {
		on, err := IsWebcamOn(context.Background())
		if err != nil {
			t.Fatal(err)
		}
		if on == false {
			break
		}
	}

	isOn, err := IsWebcamOn(context.Background())
	assert.Nil(t, err)
	assert.False(t, isOn)
}