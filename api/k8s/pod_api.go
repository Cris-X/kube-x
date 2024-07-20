package k8s

import (
	"context"
	"fmt"
	"go.uber.org/zap"
	"net/http"

	"github.com/gin-gonic/gin"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	"kube-x/global"
)

type PodApi struct {
}

// GetPodList get pod list
func (*PodApi) GetPodList(c *gin.Context) {
	ctx := context.TODO()
	list, err := global.KubeConfigSet.CoreV1().Pods("").List(ctx, metav1.ListOptions{})
	if err != nil {
		global.Logger.Error("get pod list error", zap.Error(err))
	}
	for _, item := range list.Items {
		fmt.Println(item.Namespace, item.Name)
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "pong",
	})
}
