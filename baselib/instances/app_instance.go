package instances

import "k8s.io/klog/v2"

type GAppInstance interface {
	RunLoop()
	Install() error
	Destroy()
}

func NewGAppInstance(instance GAppInstance) {
	klog.V(3)
}
