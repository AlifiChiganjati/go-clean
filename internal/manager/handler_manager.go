package manager

type (
	HandlerManager interface{}
	handlerManager struct {
		useCaseManager UseCaseManager
	}
)

func NewHandlerManager(useCaseManager UseCaseManager) HandlerManager {
	return &handlerManager{useCaseManager: useCaseManager}
}

func (h *handlerManager) UserHandler() {}
