package handlers

// Container will hold all dependencies for your application.
type Container struct {
	KafkaCtrl *AppController
}

// NewContainer returns an empty or an initialized container for your handlers.
func NewContainer(ctrl *AppController) (Container, error) {
	c := Container{KafkaCtrl: ctrl}
	return c, nil
}
