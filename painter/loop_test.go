package painter

import (
	"image"
	"image/color"
	"image/draw"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"golang.org/x/exp/shiny/screen"
)

// MockScreen is a mock of shiny screen
type MockScreen struct {
	mock.Mock
}

func (_ *MockScreen) NewBuffer(size image.Point) (screen.Buffer, error) {
	return nil, nil
}

func (_ *MockScreen) NewWindow(opts *screen.NewWindowOptions) (screen.Window, error) {
	return nil, nil
}

func (mockScreen *MockScreen) NewTexture(size image.Point) (screen.Texture, error) {
	args := mockScreen.Called(size)
	return args.Get(0).(screen.Texture), args.Error(1)
}

// MockTexture is a mock of shiny screen Texture
type MockTexture struct {
	mock.Mock
}

func (mockTexture *MockTexture) Release() {
	mockTexture.Called()
}

func (mockTexture *MockTexture) Upload(dp image.Point, src screen.Buffer, sr image.Rectangle) {
	mockTexture.Called(dp, src, sr)
}

func (mockTexture *MockTexture) Bounds() image.Rectangle {
	args := mockTexture.Called()
	return args.Get(0).(image.Rectangle)
}

func (mockTexture *MockTexture) Fill(dr image.Rectangle, src color.Color, op draw.Op) {
	mockTexture.Called(dr, src, op)
}

func (mockTexture *MockTexture) Size() image.Point {
	args := mockTexture.Called()
	return args.Get(0).(image.Point)
}

// MockReceiver is a mock of Receiver struct from loop.go
type MockReceiver struct {
	mock.Mock
}

func (mockReceiver *MockReceiver) Update(texture screen.Texture) {
	mockReceiver.Called(texture)
}

// MockOperation is a mock of Operation interface from op.go
type MockOperation struct {
	mock.Mock
}

func (mockOperation *MockOperation) Do(t screen.Texture) bool {
	args := mockOperation.Called(t)
	return args.Bool(0)
}

func TestLoop_Post_Failure(t *testing.T) {
	textureMock := new(MockTexture)
	receiverMock := new(MockReceiver)
	screenMock := new(MockScreen)

	texture := image.Pt(800, 800)
	screenMock.On("NewTexture", texture).Return(textureMock, nil)
	receiverMock.On("Update", textureMock).Return()
	loop := Loop{
		Receiver: receiverMock,
	}

	loop.Start(screenMock)

	operationOne := new(MockOperation)
	textureMock.On("Bounds").Return(image.Rectangle{})
	operationOne.On("Do", textureMock).Return(false)

	assert.Empty(t, loop.Mq.Ops)
	loop.Post(operationOne)
	time.Sleep(1 * time.Second)
	assert.Empty(t, loop.Mq.Ops)

	operationOne.AssertCalled(t, "Do", textureMock)
	receiverMock.AssertNotCalled(t, "Update", textureMock)
	screenMock.AssertCalled(t, "NewTexture", image.Pt(800, 800))
}

func TestLoop_Post_Success(t *testing.T) {
	textureMock := new(MockTexture)
	receiverMock := new(MockReceiver)
	screenMock := new(MockScreen)

	texture := image.Pt(800, 800)
	screenMock.On("NewTexture", texture).Return(textureMock, nil)
	receiverMock.On("Update", textureMock).Return()
	loop := Loop{
		Receiver: receiverMock,
	}

	loop.Start(screenMock)

	operationOne := new(MockOperation)
	textureMock.On("Bounds").Return(image.Rectangle{})
	operationOne.On("Do", textureMock).Return(true)

	assert.Empty(t, loop.Mq.Ops)
	loop.Post(operationOne)
	time.Sleep(1 * time.Second)
	assert.Empty(t, loop.Mq.Ops)

	operationOne.AssertCalled(t, "Do", textureMock)
	receiverMock.AssertCalled(t, "Update", textureMock)
	screenMock.AssertCalled(t, "NewTexture", image.Pt(800, 800))
}

func TestLoop_Post_Multiple_Success(t *testing.T) {
	textureMock := new(MockTexture)
	receiverMock := new(MockReceiver)
	screenMock := new(MockScreen)

	texture := image.Pt(800, 800)
	screenMock.On("NewTexture", texture).Return(textureMock, nil)
	receiverMock.On("Update", textureMock).Return()
	loop := Loop{
		Receiver: receiverMock,
	}

	loop.Start(screenMock)

	operationOne := new(MockOperation)
	operationTwo := new(MockOperation)
	textureMock.On("Bounds").Return(image.Rectangle{})
	operationOne.On("Do", textureMock).Return(true)
	operationTwo.On("Do", textureMock).Return(true)

	assert.Empty(t, loop.Mq.Ops)
	loop.Post(operationOne)
	loop.Post(operationTwo)
	time.Sleep(1 * time.Second)
	assert.Empty(t, loop.Mq.Ops)

	operationOne.AssertCalled(t, "Do", textureMock)
	operationTwo.AssertCalled(t, "Do", textureMock)
	receiverMock.AssertCalled(t, "Update", textureMock)
	screenMock.AssertCalled(t, "NewTexture", image.Pt(800, 800))
}
